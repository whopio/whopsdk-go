package live_test

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	whopsdk "github.com/whopio/whopsdk-go/v2"
	client "github.com/whopio/whopsdk-go/v2/client"
	core "github.com/whopio/whopsdk-go/v2/core"
	live "github.com/whopio/whopsdk-go/v2/live"
	oidc "github.com/whopio/whopsdk-go/v2/oidc"
	option "github.com/whopio/whopsdk-go/v2/option"
	stats "github.com/whopio/whopsdk-go/v2/stats"
)

const apiVersion = "2026-01-01"

type answer = *stats.RetrieveTimeSeriesResponse

func timeSeries() *stats.RetrieveTimeSeriesRequest {
	return &stats.RetrieveTimeSeriesRequest{Metric: "gmv", From: "2026-11-01", To: "2026-11-30"}
}

func answerJSON(timestamp int) map[string]any {
	return map[string]any{"data": map[string]any{"points": []any{map[string]any{"timestamp": timestamp, "value": 1.5}}}}
}

func snapshot(timestamp int) string {
	data := answerJSON(timestamp)
	data["event"] = live.SnapshotEvent
	return fmt.Sprintf("id: %d\nevent: %s\ndata: %s\n\n", timestamp, live.SnapshotEvent, jsonOf(data))
}

func closed(reason string, status int) string {
	fields := map[string]any{"reason": reason, "status": nil}
	if status != 0 {
		fields["status"] = status
	}
	return fmt.Sprintf("event: %s\ndata: %s\n\n", live.ClosedEvent, jsonOf(fields))
}

func jsonOf(value any) []byte {
	encoded, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return encoded
}

const keepAlive = ": keep-alive\n\n"

// reply is one scripted answer to a connection: an event stream of frames, or a status with a JSON
// body. After its frames a stream ends, is cut off (abort), or stays open (hold) sending a
// keep-alive every keepAlive, if set, until the client goes away.
type reply struct {
	status    int
	body      any
	frames    []string
	abort     bool
	hold      bool
	keepAlive time.Duration
}

func events(frames ...string) reply { return reply{frames: frames} }

func held(frames ...string) reply { return reply{frames: frames, hold: true} }

func status(code int) reply {
	return reply{status: code, body: map[string]any{"error": map[string]any{"message": fmt.Sprintf("HTTP %d", code)}}}
}

// server answers each connection with the next scripted reply over a real local socket, and records
// its request and when the client went away.
type server struct {
	*httptest.Server
	t       *testing.T
	replies []reply

	mu       sync.Mutex
	requests []*http.Request
	gone     chan int
}

func serve(t *testing.T, replies ...reply) *server {
	t.Helper()
	s := &server{t: t, replies: replies, gone: make(chan int, len(replies))}
	s.Server = httptest.NewServer(http.HandlerFunc(s.handle))
	t.Cleanup(s.Close)
	return s
}

func (s *server) handle(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	index := len(s.requests)
	s.requests = append(s.requests, r.Clone(r.Context()))
	s.mu.Unlock()
	if index >= len(s.replies) {
		s.t.Errorf("connection %d has no scripted reply", index+1)
		http.Error(w, "unscripted", http.StatusGone)
		return
	}
	scripted := s.replies[index]
	if scripted.frames == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(scripted.status)
		_, _ = w.Write(jsonOf(scripted.body))
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.WriteHeader(http.StatusOK)
	flusher, _ := w.(http.Flusher)
	for _, frame := range scripted.frames {
		_, _ = io.WriteString(w, frame)
		flusher.Flush()
	}
	switch {
	case scripted.abort:
		connection, _, err := http.NewResponseController(w).Hijack()
		if err == nil {
			_ = connection.Close()
		}
	case scripted.hold:
		var ticks <-chan time.Time
		if scripted.keepAlive > 0 {
			ticker := time.NewTicker(scripted.keepAlive)
			defer ticker.Stop()
			ticks = ticker.C
		}
		for {
			select {
			case <-r.Context().Done():
				s.gone <- index
				return
			case <-ticks:
				_, _ = io.WriteString(w, keepAlive)
				flusher.Flush()
			}
		}
	}
}

func (s *server) whop(opts ...option.RequestOption) *client.Whop {
	version := apiVersion
	return client.NewWhop(append([]option.RequestOption{
		option.WithBaseURL(s.URL + "/api/v1"),
		option.WithToken("sk_test"),
		option.WithAPIVersionDate(&version),
		option.WithHTTPClient(s.Client()),
	}, opts...)...)
}

func (s *server) sent() []*http.Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]*http.Request(nil), s.requests...)
}

// wentAway waits for the client to close the connection the server holds open.
func (s *server) wentAway(t *testing.T) {
	t.Helper()
	select {
	case <-s.gone:
	case <-time.After(2 * time.Second):
		t.Fatal("the client kept the connection open")
	}
}

func drain(t *testing.T, stream *live.Stream[answer]) []int {
	t.Helper()
	var timestamps []int
	deadline := time.After(10 * time.Second)
	for {
		select {
		case series, ok := <-stream.Answers():
			if !ok {
				return timestamps
			}
			timestamps = append(timestamps, series.GetData().GetPoints()[0].GetTimestamp())
		case <-deadline:
			t.Fatal("the stream did not end")
		}
	}
}

func timestampOf(series answer) int {
	return series.GetData().GetPoints()[0].GetTimestamp()
}

type drops struct {
	mu     sync.Mutex
	errors []error
}

func (d *drops) record(err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.errors = append(d.errors, err)
}

func (d *drops) statuses() []int {
	d.mu.Lock()
	defer d.mu.Unlock()
	var statuses []int
	for _, err := range d.errors {
		var apiError *core.APIError
		if errors.As(err, &apiError) {
			statuses = append(statuses, apiError.StatusCode)
		}
	}
	return statuses
}

func (d *drops) all() []error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return append([]error(nil), d.errors...)
}

func TestYieldsEachSnapshotAsTheReadsAnswer(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), keepAlive, snapshot(2), closed("not_live", 200)))
	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries())

	var answers []answer
	for series := range stream.Answers() {
		answers = append(answers, series)
	}

	require.NoError(t, stream.Err())
	require.Len(t, answers, 2)
	require.Equal(t, []int{1, 2}, []int{timestampOf(answers[0]), timestampOf(answers[1])})
	require.Empty(t, answers[0].GetExtraProperties(), "the snapshot's event name is not part of the answer")
	require.Len(t, s.sent(), 1)
}

func TestSendsTheReadsRequestAskingForTheEventStream(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), closed("not_live", 0)))
	request := timeSeries()
	account := "biz_1"
	request.AccountID = &account

	drain(t, s.whop().Stats.TimeSeries.Watch(context.Background(), request))

	sent := s.sent()[0]
	require.Equal(t, http.MethodGet, sent.Method)
	require.Equal(t, "/api/v1/stats/time_series/gmv", sent.URL.Path)
	require.Equal(t, url.Values{"from": {"2026-11-01"}, "to": {"2026-11-30"}, "account_id": {"biz_1"}}, sent.URL.Query())
	require.Equal(t, []string{live.EventStream}, sent.Header.Values("Accept"))
	require.Equal(t, "Bearer sk_test", sent.Header.Get("Authorization"))
	require.Equal(t, apiVersion, sent.Header.Get("Api-Version-Date"))
}

func TestSendsTheRequestOptionsOnEveryConnectionButNotTheirRetries(t *testing.T) {
	t.Parallel()
	s := serve(t, status(http.StatusServiceUnavailable), events(snapshot(1), closed("not_live", 0)))
	var dropped drops

	stream := s.whop().Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithRequestOptions(
			option.WithHTTPHeader(http.Header{"X-Trace": {"abc"}}),
			option.WithQueryParameters(url.Values{"extra": {"1"}}),
			option.WithMaxAttempts(5),
		),
		live.WithMaxBackoff(0),
		live.WithOnDisconnect(dropped.record),
	)

	require.Equal(t, []int{1}, drain(t, stream))
	require.NoError(t, stream.Err())
	sent := s.sent()
	require.Len(t, sent, 2)
	for _, request := range sent {
		require.Equal(t, "abc", request.Header.Get("X-Trace"))
		require.Equal(t, "1", request.URL.Query().Get("extra"))
	}
	require.Equal(t, []int{http.StatusServiceUnavailable}, dropped.statuses())
}

func TestReconnectsAfterAReconnectCloseWithoutReportingADrop(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), closed("reconnect", 0)), events(snapshot(2), closed("not_live", 200)))
	var dropped drops

	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(dropped.record))

	require.Equal(t, []int{1, 2}, drain(t, stream))
	require.Len(t, s.sent(), 2)
	require.Empty(t, dropped.all())
}

func TestCountsAStreamThatEndsWithoutACloseAsADrop(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1)), events(snapshot(2), closed("not_live", 0)))
	var dropped drops

	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(dropped.record))

	require.Equal(t, []int{1, 2}, drain(t, stream))
	require.Len(t, dropped.all(), 1)
	var ended *live.StreamEndedError
	require.ErrorAs(t, dropped.all()[0], &ended)
}

func TestReadsLinesEndedByALoneCRThroughOneThatEndsTheBody(t *testing.T) {
	t.Parallel()
	crEnded := func(frame string) string { return strings.ReplaceAll(frame, "\n", "\r") }
	s := serve(t, events(crEnded(snapshot(1)), "event: "+live.ClosedEvent+"\r\ndata: {\"reason\":\"not_live\",\"status\":null}\r\n\r"))
	var dropped drops

	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(dropped.record))

	require.Equal(t, []int{1}, drain(t, stream))
	require.NoError(t, stream.Err())
	require.Empty(t, dropped.all())
	require.Len(t, s.sent(), 1)
}

func TestReconnectsAfterADroppedConnection(t *testing.T) {
	t.Parallel()
	s := serve(t, reply{frames: []string{snapshot(1)}, abort: true}, events(snapshot(2), closed("not_live", 0)))
	var dropped drops

	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(dropped.record))

	require.Equal(t, []int{1, 2}, drain(t, stream))
	require.Len(t, dropped.all(), 1)
	require.ErrorIs(t, dropped.all()[0], io.ErrUnexpectedEOF)
}

func TestReconnectsAfterAResponseItCannotRead(t *testing.T) {
	t.Parallel()
	listener, err := (&net.ListenConfig{}).Listen(context.Background(), "tcp", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })
	go func() {
		for {
			connection, err := listener.Accept()
			if err != nil {
				return
			}
			go func(connection net.Conn) {
				defer func() { _ = connection.Close() }()
				reader := bufio.NewReader(connection)
				for {
					line, err := reader.ReadString('\n')
					if err != nil || line == "\r\n" {
						break
					}
				}
				_, _ = connection.Write([]byte("HTTP/1.1 garbage\r\n\r\n"))
			}(connection)
		}
	}()
	var dropped drops
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	version := apiVersion
	whop := client.NewWhop(option.WithBaseURL("http://"+listener.Addr().String()), option.WithToken("sk_test"), option.WithAPIVersionDate(&version))

	stream := whop.Stats.TimeSeries.Watch(ctx, timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(func(err error) {
		dropped.record(err)
		if len(dropped.all()) == 2 {
			cancel()
		}
	}))

	require.Empty(t, drain(t, stream))
	require.Len(t, dropped.all(), 2)
	require.ErrorIs(t, stream.Err(), context.Canceled)
}

func TestReconnectsAfterAConnectionTheServerRefuses(t *testing.T) {
	t.Parallel()
	refused := httptest.NewServer(http.NotFoundHandler())
	refused.Close()
	var dropped drops
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	version := apiVersion
	whop := client.NewWhop(option.WithBaseURL(refused.URL), option.WithToken("sk_test"), option.WithAPIVersionDate(&version))

	stream := whop.Stats.TimeSeries.Watch(ctx, timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(func(err error) {
		dropped.record(err)
		if len(dropped.all()) == 2 {
			cancel()
		}
	}))

	require.Empty(t, drain(t, stream))
	require.ErrorIs(t, stream.Err(), context.Canceled)
	require.Len(t, dropped.all(), 2)
}

// clock answers its times in turn, then its last one: the stream reads it when a connection opens
// and again when it drops.
type clock struct {
	mu    sync.Mutex
	times []time.Duration
}

func (c *clock) now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	at := c.times[0]
	if len(c.times) > 1 {
		c.times = c.times[1:]
	}
	return time.Unix(0, 0).Add(at)
}

type waits struct {
	mu    sync.Mutex
	waits []time.Duration
}

func (w *waits) record(_ context.Context, wait time.Duration) bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.waits = append(w.waits, wait)
	return true
}

func TestBacksOffUntilAConnectionStaysOpen30SecondsAnswersOrNot(t *testing.T) {
	t.Parallel()
	s := serve(t,
		status(http.StatusServiceUnavailable),
		status(http.StatusBadGateway),
		status(http.StatusInternalServerError),
		events(snapshot(1), closed("reconnect", 0)),
		events(snapshot(2), closed("reconnect", 0)),
		events(snapshot(3), closed("not_live", 0)),
	)
	var waited waits
	opened := &clock{times: []time.Duration{0, 30 * time.Second, 30 * time.Second}}

	stream := s.whop().Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithPause(waited.record),
		live.WithClock(opened.now),
		live.WithRandom(func() float64 { return 0.5 }),
	)

	require.Equal(t, []int{1, 2, 3}, drain(t, stream))
	ms := time.Millisecond
	require.Equal(t, []time.Duration{500 * ms, 1125 * ms, 2250 * ms, 500 * ms, 1125 * ms}, waited.waits)
}

func TestKeepsBackingOffAfterAConnectionThatDroppedWithin30Seconds(t *testing.T) {
	t.Parallel()
	s := serve(t,
		status(http.StatusServiceUnavailable),
		events(snapshot(1), closed("reconnect", 0)),
		events(snapshot(2), closed("not_live", 0)),
	)
	var waited waits
	opened := &clock{times: []time.Duration{0, 29900 * time.Millisecond}}

	stream := s.whop().Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithPause(waited.record),
		live.WithClock(opened.now),
		live.WithRandom(func() float64 { return 0 }),
	)

	require.Equal(t, []int{1, 2}, drain(t, stream))
	require.Equal(t, []time.Duration{0, time.Second}, waited.waits)
}

func TestEndsQuietlyWhenTheReadStopsBeingLive(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), closed("not_live", 0)))
	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries())

	require.Equal(t, []int{1}, drain(t, stream))
	require.NoError(t, stream.Err())
}

func TestFailsWhenTheStreamClosesForGood(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		reason string
		status int
	}{{"not_live", http.StatusForbidden}, {"unreadable", 0}} {
		test := test
		t.Run(test.reason, func(t *testing.T) {
			t.Parallel()
			s := serve(t, events(snapshot(1), closed(test.reason, test.status)))
			stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries())

			require.Equal(t, []int{1}, drain(t, stream))
			var closedForGood *live.StreamClosedError
			require.ErrorAs(t, stream.Err(), &closedForGood)
			require.Equal(t, live.StreamClosedError{Reason: test.reason, Status: test.status}, *closedForGood)
		})
	}
}

func TestYieldsAPlainAnswerOnceThenEnds(t *testing.T) {
	t.Parallel()
	s := serve(t, reply{status: http.StatusOK, body: answerJSON(7)})
	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries())

	require.Equal(t, []int{7}, drain(t, stream))
	require.NoError(t, stream.Err())
	require.Len(t, s.sent(), 1)
}

func TestFailsWithTheReadsOwnErrorForAStatusAReconnectWouldGetAgain(t *testing.T) {
	t.Parallel()
	for code, expected := range map[int]error{
		http.StatusBadRequest:          &whopsdk.BadRequestError{},
		http.StatusUnauthorized:        &whopsdk.UnauthorizedError{},
		http.StatusForbidden:           &whopsdk.ForbiddenError{},
		http.StatusNotFound:            &whopsdk.NotFoundError{},
		http.StatusGone:                &core.APIError{},
		http.StatusUnprocessableEntity: &core.APIError{},
	} {
		code, expected := code, expected
		t.Run(http.StatusText(code), func(t *testing.T) {
			t.Parallel()
			s := serve(t, status(code), status(code))
			var dropped drops
			_, readErr := s.whop().Stats.TimeSeries.Retrieve(context.Background(), timeSeries())

			stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithOnDisconnect(dropped.record))

			require.Empty(t, drain(t, stream))
			require.Equal(t, reflect.TypeOf(expected), reflect.TypeOf(stream.Err()))
			require.Equal(t, reflect.TypeOf(readErr), reflect.TypeOf(stream.Err()))
			require.EqualError(t, stream.Err(), readErr.Error())
			var apiError *core.APIError
			require.ErrorAs(t, stream.Err(), &apiError)
			require.Equal(t, code, apiError.StatusCode)
			require.Len(t, s.sent(), 2)
			require.Empty(t, dropped.all())
		})
	}
}

func TestReconnectsAfterAClientErrorTheServerRetries(t *testing.T) {
	t.Parallel()
	for _, code := range []int{http.StatusRequestTimeout, http.StatusTooManyRequests} {
		code := code
		t.Run(http.StatusText(code), func(t *testing.T) {
			t.Parallel()
			s := serve(t, status(code), events(snapshot(1), closed("not_live", 0)))
			var dropped drops

			stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithMaxBackoff(0), live.WithOnDisconnect(dropped.record))

			require.Equal(t, []int{1}, drain(t, stream))
			require.Equal(t, []int{code}, dropped.statuses())
		})
	}
}

func TestWaitForReturnsTheFirstAnswerUntilAccepts(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), snapshot(2), snapshot(3), closed("not_live", 0)))

	found, err := s.whop().Stats.TimeSeries.WaitFor(context.Background(), timeSeries(), func(series answer) bool {
		return timestampOf(series) >= 2
	})

	require.NoError(t, err)
	require.Equal(t, 2, timestampOf(found))
}

func TestWaitForFailsOnTheFirstAnswerFailWhenAccepts(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), snapshot(2), closed("not_live", 0)))

	_, err := s.whop().Stats.TimeSeries.WaitFor(
		context.Background(),
		timeSeries(),
		func(series answer) bool { return timestampOf(series) == 2 },
		live.WithFailWhen(func(series answer) bool { return timestampOf(series) == 1 }),
	)

	var failed *live.WaitForFailedError[answer]
	require.ErrorAs(t, err, &failed)
	require.Equal(t, 1, timestampOf(failed.Answer))
}

func TestWaitForRefusesAFailWhenForAnotherAnswerType(t *testing.T) {
	t.Parallel()
	s := serve(t)

	_, err := s.whop().Stats.TimeSeries.WaitFor(
		context.Background(),
		timeSeries(),
		func(answer) bool { return true },
		live.WithFailWhen(func(string) bool { return true }),
	)

	require.ErrorContains(t, err, "WithFailWhen takes a func(*stats.RetrieveTimeSeriesResponse) bool for this read, not a func(string) bool")
	require.Empty(t, s.sent())
}

func TestWaitForTimesOut(t *testing.T) {
	t.Parallel()
	s := serve(t, held(snapshot(1)))

	_, err := s.whop().Stats.TimeSeries.WaitFor(
		context.Background(),
		timeSeries(),
		func(answer) bool { return false },
		live.WithTimeout(200*time.Millisecond),
	)

	var timedOut *live.WaitForTimeoutError
	require.ErrorAs(t, err, &timedOut)
	require.Equal(t, 200*time.Millisecond, timedOut.Timeout)
	s.wentAway(t)
}

func TestWaitForFailsWhenTheReadStopsBeingLiveBeforeAMatch(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), closed("not_live", 0)))

	_, err := s.whop().Stats.TimeSeries.WaitFor(context.Background(), timeSeries(), func(answer) bool { return false })

	var closedForGood *live.StreamClosedError
	require.ErrorAs(t, err, &closedForGood)
	require.Equal(t, "not_live", closedForGood.Reason)
}

func TestWaitForFailsWithTheReadsOwnError(t *testing.T) {
	t.Parallel()
	s := serve(t, status(http.StatusUnauthorized))

	_, err := s.whop().Stats.TimeSeries.WaitFor(context.Background(), timeSeries(), func(answer) bool { return true })

	var unauthorized *whopsdk.UnauthorizedError
	require.ErrorAs(t, err, &unauthorized)
}

func TestWaitForClosesItsConnectionAsItReturns(t *testing.T) {
	t.Parallel()
	s := serve(t, held(snapshot(1)))

	found, err := s.whop().Stats.TimeSeries.WaitFor(context.Background(), timeSeries(), func(answer) bool { return true })

	require.NoError(t, err)
	require.Equal(t, 1, timestampOf(found))
	s.wentAway(t)
	require.Len(t, s.sent(), 1)
}

func TestCloseFromAnotherGoroutineEndsALoopReadingTheSocket(t *testing.T) {
	t.Parallel()
	for name, keepAlive := range map[string]time.Duration{"keep-alive": 100 * time.Millisecond, "silent": 0} {
		keepAlive := keepAlive
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			s := serve(t, reply{frames: []string{snapshot(1)}, hold: true, keepAlive: keepAlive})
			var dropped drops
			stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithOnDisconnect(dropped.record))

			first := <-stream.Answers()
			time.AfterFunc(300*time.Millisecond, func() { _ = stream.Close() })
			started := time.Now()
			rest := drain(t, stream)

			require.Equal(t, 1, timestampOf(first))
			require.Empty(t, rest)
			require.Less(t, time.Since(started), 2*time.Second)
			require.NoError(t, stream.Err())
			s.wentAway(t)
			require.Len(t, s.sent(), 1)
			require.Empty(t, dropped.all())
		})
	}
}

func TestCloseReturnsOnceTheConnectionIsClosed(t *testing.T) {
	t.Parallel()
	s := serve(t, held(snapshot(1)))
	stream := s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries())
	<-stream.Answers()

	require.NoError(t, stream.Close())

	_, open := <-stream.Answers()
	require.False(t, open)
	require.NoError(t, stream.Err())
	s.wentAway(t)
}

func TestCancellingTheContextEndsTheStreamAndClosesTheConnection(t *testing.T) {
	t.Parallel()
	s := serve(t, held(snapshot(1)))
	ctx, cancel := context.WithCancel(context.Background())
	stream := s.whop().Stats.TimeSeries.Watch(ctx, timeSeries())
	<-stream.Answers()

	cancel()

	require.Empty(t, drain(t, stream))
	require.ErrorIs(t, stream.Err(), context.Canceled)
	s.wentAway(t)
}

func TestCloseFromOnDisconnectEndsTheStream(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1)))
	var stream *live.Stream[answer]
	ready := make(chan struct{})
	stream = s.whop().Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithOnDisconnect(func(error) {
		<-ready
		_ = stream.Close()
	}))
	close(ready)

	require.Equal(t, []int{1}, drain(t, stream))
	require.NoError(t, stream.Err())
	require.Len(t, s.sent(), 1)
}

func TestDropsAConnectionThatReceivesNothingForTheIdleTimeout(t *testing.T) {
	t.Parallel()
	s := serve(t, held(snapshot(1)), events(snapshot(2), closed("not_live", 0)))
	var dropped drops

	stream := s.whop().Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithIdleTimeout(200*time.Millisecond),
		live.WithMaxBackoff(0),
		live.WithOnDisconnect(dropped.record),
	)

	require.Equal(t, []int{1, 2}, drain(t, stream))
	require.Len(t, dropped.all(), 1)
	require.Equal(t, &live.StreamIdleError{Timeout: 200 * time.Millisecond}, dropped.all()[0])
	s.wentAway(t)
}

func TestKeepAlivesHoldOffTheIdleTimeout(t *testing.T) {
	t.Parallel()
	s := serve(t, reply{frames: []string{snapshot(1)}, hold: true, keepAlive: 50 * time.Millisecond})
	var dropped drops
	stream := s.whop().Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithIdleTimeout(200*time.Millisecond),
		live.WithOnDisconnect(dropped.record),
	)
	<-stream.Answers()

	time.Sleep(600 * time.Millisecond)
	require.NoError(t, stream.Close())

	require.Empty(t, dropped.all())
	require.Len(t, s.sent(), 1)
}

func TestHoldsOffTheIdleTimeoutWhileTheReaderIsBusy(t *testing.T) {
	t.Parallel()
	s := serve(t, held(snapshot(1), snapshot(2)))
	var dropped drops
	stream := s.whop().Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithIdleTimeout(200*time.Millisecond),
		live.WithOnDisconnect(dropped.record),
	)
	first := <-stream.Answers()

	time.Sleep(500 * time.Millisecond)
	second := <-stream.Answers()
	require.NoError(t, stream.Close())

	require.Equal(t, []int{1, 2}, []int{timestampOf(first), timestampOf(second)})
	require.Empty(t, dropped.all())
}

func TestStreamsPastAnHTTPClientTimeout(t *testing.T) {
	t.Parallel()
	s := serve(t, reply{frames: []string{snapshot(1)}, hold: true, keepAlive: 50 * time.Millisecond})
	var dropped drops
	timed := s.Client()
	timed.Timeout = 100 * time.Millisecond
	stream := s.whop(option.WithHTTPClient(timed)).Stats.TimeSeries.Watch(context.Background(), timeSeries(), live.WithOnDisconnect(dropped.record))
	<-stream.Answers()

	time.Sleep(400 * time.Millisecond)
	require.NoError(t, stream.Close())

	require.Empty(t, dropped.all())
	require.Len(t, s.sent(), 1)
}

func TestSendsThroughTheHTTPClientTheRequestOptionsName(t *testing.T) {
	t.Parallel()
	s := serve(t, events(snapshot(1), closed("not_live", 0)))
	var used []string
	recording := doer(func(request *http.Request) (*http.Response, error) {
		used = append(used, request.Header.Get("Accept"))
		return s.Client().Do(request)
	})
	unreachable := doer(func(*http.Request) (*http.Response, error) {
		t.Error("the stream sent through the client's own HTTP client, not the one its request options name")
		return nil, errors.New("unreachable")
	})

	stream := s.whop(option.WithHTTPClient(unreachable)).Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithRequestOptions(option.WithHTTPClient(recording)),
	)

	require.Equal(t, []int{1}, drain(t, stream))
	require.Equal(t, []string{live.EventStream}, used)
}

func TestFailsWithAnOIDCExchangeErrorInsteadOfReconnecting(t *testing.T) {
	t.Parallel()
	refused := &oidc.Error{Status: http.StatusForbidden, Code: "oidc_exchange_ip_not_allowed", Message: "IP not allowed"}
	var dropped drops
	exchange := doer(func(*http.Request) (*http.Response, error) { return nil, refused })

	stream := serve(t).whop(option.WithHTTPClient(exchange)).Stats.TimeSeries.Watch(
		context.Background(),
		timeSeries(),
		live.WithMaxBackoff(0),
		live.WithOnDisconnect(dropped.record),
	)

	require.Empty(t, drain(t, stream))
	require.ErrorIs(t, stream.Err(), refused)
	require.Empty(t, dropped.all())
}

type doer func(*http.Request) (*http.Response, error)

func (do doer) Do(request *http.Request) (*http.Response, error) { return do(request) }

func TestOnlyWaitForOptionsConfigureWaitForAlone(t *testing.T) {
	t.Parallel()
	for _, waitForOption := range []live.WaitForOption{live.WithTimeout(time.Second), live.WithFailWhen(func(answer) bool { return false })} {
		_, isWatchOption := waitForOption.(live.WatchOption)
		require.False(t, isWatchOption, reflect.TypeOf(waitForOption).String())
	}
}
