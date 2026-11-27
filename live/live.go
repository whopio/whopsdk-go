// Package live reads a live read's answer now, then again each time it changes.
//
// A live read is a GET the API also answers as an event stream: the same request with
// Accept: text/event-stream gets the read's whole answer now and again each time it changes. Its
// client carries Watch and WaitFor beside the read. They take exactly the read's parameters, then
// their own options:
//
//	request := &stats.RetrieveTimeSeriesRequest{Metric: "gmv", From: "2026-11-01", To: "2026-11-30"}
//	stream := whop.Stats.TimeSeries.Watch(ctx, request, live.WithMaxBackoff(10*time.Second))
//	defer stream.Close()
//	for series := range stream.Answers() {
//		fmt.Println(series.GetData().GetPoints())
//	}
//	if err := stream.Err(); err != nil {
//		return err
//	}
//
//	series, err := whop.Stats.TimeSeries.WaitFor(ctx, request, func(series *stats.RetrieveTimeSeriesResponse) bool {
//		return len(series.GetData().GetPoints()) == 30
//	}, live.WithTimeout(2*time.Minute))
//
// Answers yields each answer as the read's own answer type and closes when the stream ends: when
// the read stops being live (Err is nil), when it fails for good (Err says why), when ctx ends (Err
// is ctx.Err()), or when Close is called from any goroutine (Err is nil). Close returns once the
// connection is closed. Close a stream you stop reading, or its connection stays open.
//
// Each connection sends the request the read itself builds: the read runs against a stand-in for
// the HTTP client that records the request instead of sending it, and that request is sent with
// Accept: text/event-stream through the client's HTTP client. The request options passed with
// WithRequestOptions therefore apply to every connection; retries do not, since each connection is
// sent once and the stream reconnects itself. An *http.Client's Timeout would end every connection
// after it, so a stream sends through a copy without one, and drops a connection that receives
// nothing for 45 seconds instead (the server sends a keep-alive every 15).
//
// The stream reconnects after a drop, and each connection starts with a fresh answer, so nothing
// is replayed:
//   - a stream.closed event with the reason not_live and no status or a 2xx status ends it: the
//     last answer stands and will not change.
//   - not_live with any other status, or unreadable, fails it with a *StreamClosedError.
//   - any other close reconnects.
//   - a stream that ends without a close is a drop: WithOnDisconnect's callback gets a
//     *StreamEndedError, then it reconnects. So is a connection idle for 45 seconds
//     (*StreamIdleError).
//   - a client error fails it with the error the read returns for it, as the server ends a read
//     on one; 408, 429, a server error or a network error is a drop.
//   - a success that is not an event stream is the read's one answer: Answers yields it, then the
//     stream ends.
//
// Reconnects back off: the first waits a random 0 to 1 second (at most the maximum backoff), each
// after it 1 second doubling up to the maximum backoff, plus up to 25% jitter. The count starts
// over only after a connection that stayed open at least 30 seconds, not after an answer, since
// every connection opens with one.
package live

import (
	"bufio"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/whopio/whopsdk-go/v2/core"
	"github.com/whopio/whopsdk-go/v2/oidc"
	"github.com/whopio/whopsdk-go/v2/option"
)

// The events of a live read's stream, and the media type that asks for it.
const (
	SnapshotEvent = "stream.snapshot"
	ClosedEvent   = "stream.closed"
	EventStream   = "text/event-stream"
)

// The defaults of WithMaxBackoff and WithTimeout, and how long a connection may receive nothing
// before the stream drops it.
const (
	DefaultMaxBackoff     = 30 * time.Second
	DefaultWaitForTimeout = 30 * time.Second
	IdleTimeout           = 45 * time.Second
)

const steadyConnection = 30 * time.Second

// Read is a live read's generated method with its parameters bound: it sends the read with the
// request options it is given.
type Read[T any] func(ctx context.Context, opts ...option.RequestOption) (T, error)

// StreamClosedError is a stream that closed for good: not_live with the status the read now
// answers, or unreadable. Status is 0 when the server sent none.
type StreamClosedError struct {
	Reason string
	Status int
}

func (e *StreamClosedError) Error() string {
	if e.Status == 0 {
		return "live: the stream closed: " + e.Reason
	}
	return fmt.Sprintf("live: the stream closed: %s (%d)", e.Reason, e.Status)
}

// StreamEndedError is a stream that ended without a stream.closed event: WithOnDisconnect's
// callback gets it, then the stream reconnects.
type StreamEndedError struct{}

func (*StreamEndedError) Error() string {
	return "live: the stream ended without a " + ClosedEvent + " event"
}

// StreamIdleError is a connection that received nothing for Timeout: WithOnDisconnect's callback
// gets it, then the stream reconnects.
type StreamIdleError struct {
	Timeout time.Duration
}

func (e *StreamIdleError) Error() string {
	return fmt.Sprintf("live: the stream received nothing for %s", e.Timeout)
}

// WaitForTimeoutError is WaitFor's timeout passing before an answer matched.
type WaitForTimeoutError struct {
	Timeout time.Duration
}

func (e *WaitForTimeoutError) Error() string {
	return fmt.Sprintf("live: no answer matched within %s", e.Timeout)
}

// WaitForFailedError is an answer WithFailWhen's predicate accepted.
type WaitForFailedError[T any] struct {
	Answer T
}

func (*WaitForFailedError[T]) Error() string {
	return "live: an answer matched WithFailWhen"
}

type settings struct {
	requestOptions []option.RequestOption
	maxBackoff     time.Duration
	onDisconnect   func(error)
	timeout        time.Duration
	failWhen       any
	idleTimeout    time.Duration
	now            func() time.Time
	random         func() float64
	pause          func(ctx context.Context, wait time.Duration) bool
}

func defaults() *settings {
	return &settings{
		maxBackoff:  DefaultMaxBackoff,
		timeout:     DefaultWaitForTimeout,
		idleTimeout: IdleTimeout,
		now:         time.Now,
		random:      rand.Float64,
		pause:       pause,
	}
}

// WaitForOption configures WaitFor. Every WatchOption is one.
type WaitForOption interface {
	applyWaitFor(s *settings)
}

// WatchOption configures Watch and WaitFor.
type WatchOption interface {
	WaitForOption
	applyWatch(s *settings)
}

type watchOption func(*settings)

func (apply watchOption) applyWatch(s *settings) { apply(s) }

func (apply watchOption) applyWaitFor(s *settings) { apply(s) }

type waitForOption func(*settings)

func (apply waitForOption) applyWaitFor(s *settings) { apply(s) }

// WithRequestOptions sends every connection with these request options, as the read is sent with
// them, except that retries do not apply.
func WithRequestOptions(opts ...option.RequestOption) WatchOption {
	return watchOption(func(s *settings) { s.requestOptions = append(s.requestOptions, opts...) })
}

// WithMaxBackoff caps the wait between reconnects; DefaultMaxBackoff by default.
func WithMaxBackoff(maxBackoff time.Duration) WatchOption {
	return watchOption(func(s *settings) { s.maxBackoff = maxBackoff })
}

// WithOnDisconnect observes each drop before the reconnect, on the goroutine reading the stream.
func WithOnDisconnect(onDisconnect func(error)) WatchOption {
	return watchOption(func(s *settings) { s.onDisconnect = onDisconnect })
}

// WithTimeout ends WaitFor with a *WaitForTimeoutError once it passes; DefaultWaitForTimeout by
// default, and 0 waits until ctx ends.
func WithTimeout(timeout time.Duration) WaitForOption {
	return waitForOption(func(s *settings) { s.timeout = timeout })
}

// WithFailWhen ends WaitFor with a *WaitForFailedError on the first answer failWhen accepts,
// checked before WaitFor's own predicate. Its answer type must be the read's.
func WithFailWhen[T any](failWhen func(T) bool) WaitForOption {
	return waitForOption(func(s *settings) { s.failWhen = failWhen })
}

// Stream is a live read's answer now and each time it changes; see the package documentation.
type Stream[T any] struct {
	answers chan T
	cancel  context.CancelFunc

	mu         sync.Mutex
	err        error
	closed     bool
	connection chan struct{}
}

// Answers yields the read's answer now and each time it changes, and closes when the stream ends.
func (s *Stream[T]) Answers() <-chan T {
	return s.answers
}

// Err is why the stream ended, once Answers is closed: nil when the read stopped being live or the
// stream was closed, ctx.Err() when ctx ended, and otherwise the error that ended it.
func (s *Stream[T]) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

// Close ends the stream from any goroutine and returns once its connection is closed.
func (s *Stream[T]) Close() error {
	s.mu.Lock()
	s.closed = true
	connection := s.connection
	s.mu.Unlock()
	s.cancel()
	if connection != nil {
		<-connection
	}
	return nil
}

// Watch reads read's answer now and each time it changes, sending each connection through client
// (the HTTP client the read's client was built with) unless WithRequestOptions names another.
func Watch[T any](ctx context.Context, client core.HTTPClient, read Read[T], opts ...WatchOption) *Stream[T] {
	s := defaults()
	for _, opt := range opts {
		opt.applyWatch(s)
	}
	return start(ctx, &connector[T]{client: client, read: read, settings: s})
}

// WaitFor returns the first of read's answers until accepts, the current one included. Its
// connection is closed as it returns.
func WaitFor[T any](
	ctx context.Context,
	client core.HTTPClient,
	read Read[T],
	until func(T) bool,
	opts ...WaitForOption,
) (T, error) {
	var none T
	s := defaults()
	for _, opt := range opts {
		opt.applyWaitFor(s)
	}
	failWhen, err := failWhenOf[T](s.failWhen)
	if err != nil {
		return none, err
	}
	if s.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeoutCause(ctx, s.timeout, errWaitForTimeout)
		defer cancel()
	}
	stream := start(ctx, &connector[T]{client: client, read: read, settings: s})
	defer func() { _ = stream.Close() }()
	for answer := range stream.Answers() {
		if failWhen != nil && failWhen(answer) {
			return none, &WaitForFailedError[T]{Answer: answer}
		}
		if until(answer) {
			return answer, nil
		}
	}
	if errors.Is(context.Cause(ctx), errWaitForTimeout) {
		return none, &WaitForTimeoutError{Timeout: s.timeout}
	}
	if err := stream.Err(); err != nil {
		return none, err
	}
	return none, &StreamClosedError{Reason: "not_live"}
}

var errWaitForTimeout = errors.New("live: WaitFor timed out")

func failWhenOf[T any](failWhen any) (func(T) bool, error) {
	if failWhen == nil {
		return nil, nil
	}
	predicate, ok := failWhen.(func(T) bool)
	if !ok {
		var answer T
		return nil, fmt.Errorf("live: WithFailWhen takes a func(%T) bool for this read, not a %T", answer, failWhen)
	}
	return predicate, nil
}

func start[T any](ctx context.Context, c *connector[T]) *Stream[T] {
	ctx, cancel := context.WithCancel(ctx)
	stream := &Stream[T]{answers: make(chan T), cancel: cancel}
	go stream.run(ctx, c)
	return stream
}

func (s *Stream[T]) run(ctx context.Context, c *connector[T]) {
	err := c.watch(ctx, s)
	s.mu.Lock()
	switch {
	case s.closed:
		s.err = nil
	case ctx.Err() != nil:
		s.err = ctx.Err()
	default:
		s.err = err
	}
	s.mu.Unlock()
	s.cancel()
	close(s.answers)
}

func (s *Stream[T]) deliver(ctx context.Context, answer T) bool {
	select {
	case s.answers <- answer:
		return true
	case <-ctx.Done():
		return false
	}
}

// open records a connection Close waits for; it refuses once the stream is closed.
func (s *Stream[T]) open() (chan struct{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return nil, false
	}
	s.connection = make(chan struct{})
	return s.connection, true
}

func (s *Stream[T]) release(connection chan struct{}) {
	s.mu.Lock()
	s.connection = nil
	s.mu.Unlock()
	close(connection)
}

// backoff is the wait before reconnect attempt: a random 0 to 1 second (at most maxBackoff) for
// attempt 0, then 1 second doubling up to maxBackoff, plus up to 25%.
func backoff(attempt int, maxBackoff time.Duration, fraction float64) time.Duration {
	if attempt <= 0 {
		return time.Duration(float64(min(maxBackoff, time.Second)) * fraction)
	}
	base := math.Min(float64(maxBackoff), float64(time.Second)*math.Pow(2, float64(attempt-1)))
	return time.Duration(base + base*0.25*fraction)
}

func pause(ctx context.Context, wait time.Duration) bool {
	if wait <= 0 {
		return ctx.Err() == nil
	}
	timer := time.NewTimer(wait)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

type connector[T any] struct {
	client   core.HTTPClient
	read     Read[T]
	settings *settings
}

// drop is a connection lost in a way a reconnect may not be: the stream reconnects after it.
type drop struct {
	err error
}

func (d *drop) Error() string { return d.err.Error() }

func (d *drop) Unwrap() error { return d.err }

var errReconnect = errors.New("live: the server asked for a reconnect")

func (c *connector[T]) watch(ctx context.Context, s *Stream[T]) error {
	attempt := 0
	for {
		opened, err := c.connect(ctx, s)
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err == nil {
			return nil
		}
		var dropped *drop
		switch {
		case errors.Is(err, errReconnect):
		case errors.As(err, &dropped):
			if c.settings.onDisconnect != nil {
				c.settings.onDisconnect(dropped.err)
			}
		default:
			return err
		}
		if !opened.IsZero() && c.settings.now().Sub(opened) >= steadyConnection {
			attempt = 0
		}
		if !c.settings.pause(ctx, backoff(attempt, c.settings.maxBackoff, c.settings.random())) {
			return ctx.Err()
		}
		attempt++
	}
}

// connect opens one connection and reads it to its end: nil when the stream ends for good
// quietly, errReconnect, a *drop, or the error that ends the stream. opened is when the event
// stream opened, if it did.
func (c *connector[T]) connect(ctx context.Context, s *Stream[T]) (opened time.Time, err error) {
	connection, cancel := context.WithCancel(ctx)
	defer cancel()
	var idle atomic.Bool
	var heardAt atomic.Int64
	heardAt.Store(time.Now().UnixNano())
	// Reset can race a firing that has already begun, so the watchdog drops the connection only when
	// nothing has arrived for the whole idle timeout.
	watchdog := time.AfterFunc(c.settings.idleTimeout, func() {
		if time.Since(time.Unix(0, heardAt.Load())) < c.settings.idleTimeout {
			return
		}
		idle.Store(true)
		cancel()
	})
	defer watchdog.Stop()
	lost := func(err error) error {
		switch {
		case idle.Load():
			return &drop{&StreamIdleError{Timeout: c.settings.idleTimeout}}
		case ctx.Err() != nil:
			return ctx.Err()
		default:
			return err
		}
	}

	request, err := c.capture(connection)
	if err != nil {
		return opened, lost(err)
	}
	open, ok := s.open()
	if !ok {
		return opened, context.Canceled
	}
	defer s.release(open)

	response, err := send(c.httpClient(), request)
	if err != nil {
		if !failsAgainOnReconnect(err) {
			err = &drop{err}
		}
		return opened, lost(err)
	}
	defer func() { _ = response.Body.Close() }()

	if !isEventStream(response) {
		status := response.StatusCode
		answer, failed := c.replay(connection, response)
		if failed != nil {
			if reconnectsAfter(status) {
				failed = &drop{failed}
			}
			return opened, lost(failed)
		}
		if !s.deliver(ctx, answer) {
			return opened, ctx.Err()
		}
		return opened, nil
	}

	opened = c.settings.now()
	body, err := decompressed(response)
	if err != nil {
		return opened, lost(&drop{err})
	}
	lines := lineReader{reader: bufio.NewReader(&heard{reader: body, watchdog: watchdog, heardAt: &heardAt, idleTimeout: c.settings.idleTimeout})}
	var events parser
	for {
		line, err := lines.next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = &StreamEndedError{}
			}
			return opened, lost(&drop{err})
		}
		name, data, ok := events.feed(line)
		if !ok {
			continue
		}
		switch name {
		case ClosedEvent:
			return opened, closedBy(data)
		case SnapshotEvent:
			answer, err := decode[T](data)
			if err != nil {
				return opened, err
			}
			watchdog.Stop()
			if !s.deliver(ctx, answer) {
				return opened, ctx.Err()
			}
			heardAt.Store(time.Now().UnixNano())
			watchdog.Reset(c.settings.idleTimeout)
		}
	}
}

type capture struct {
	request *http.Request
}

var errCaptured = errors.New("live: captured the read's request")

func (c *capture) Do(request *http.Request) (*http.Response, error) {
	c.request = request
	return nil, errCaptured
}

type replay struct {
	response *http.Response
}

func (r replay) Do(*http.Request) (*http.Response, error) {
	return r.response, nil
}

func (c *connector[T]) options(client core.HTTPClient) []option.RequestOption {
	opts := make([]option.RequestOption, 0, len(c.settings.requestOptions)+2)
	opts = append(opts, c.settings.requestOptions...)
	return append(opts, option.WithoutRetries(), option.WithHTTPClient(client))
}

// capture runs the read against a stand-in for its HTTP client and returns the request it built,
// asking for the event stream.
func (c *connector[T]) capture(ctx context.Context) (*http.Request, error) {
	captured := &capture{}
	_, err := c.read(ctx, c.options(captured)...)
	if captured.request == nil {
		if err == nil {
			err = errors.New("live: the read returned without building a request")
		}
		return nil, err
	}
	captured.request.Header.Set("Accept", EventStream)
	return captured.request, nil
}

// replay hands the read a response already received, so the read itself parses its answer or its
// error.
func (c *connector[T]) replay(ctx context.Context, response *http.Response) (T, error) {
	return c.read(ctx, c.options(replay{response: response})...)
}

func (c *connector[T]) httpClient() core.HTTPClient {
	if client := core.NewRequestOptions(c.settings.requestOptions...).HTTPClient; client != nil {
		return client
	}
	if c.client != nil {
		return c.client
	}
	return http.DefaultClient
}

func send(client core.HTTPClient, request *http.Request) (*http.Response, error) {
	switch limited := client.(type) {
	case *http.Client:
		if limited.Timeout > 0 {
			unlimited := *limited
			unlimited.Timeout = 0
			return unlimited.Do(request)
		}
	case *oidc.HTTPClient:
		return limited.WithoutTimeout().Do(request)
	}
	return client.Do(request)
}

// reconnectsAfter is whether a reconnect may get another answer than this status: not after a
// success or a client error, as the server ends a read on one, but after a timeout or too many
// requests.
func reconnectsAfter(status int) bool {
	switch {
	case status >= 200 && status < 300:
		return false
	case status == http.StatusRequestTimeout || status == http.StatusTooManyRequests:
		return true
	default:
		return status < 400 || status >= 500
	}
}

func isEventStream(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 300 &&
		strings.Contains(response.Header.Get("Content-Type"), EventStream)
}

// A server whose certificate fails verification fails it again on every reconnect, and an OIDC token
// exchange that failed fails the stream as it fails any request; any other failure to get an answer
// (a reset, a protocol error, a refused connection) may not.
func failsAgainOnReconnect(err error) bool {
	var verification *tls.CertificateVerificationError
	var exchange *oidc.Error
	return errors.As(err, &verification) || errors.As(err, &exchange)
}

func decompressed(response *http.Response) (io.Reader, error) {
	if response.Uncompressed || !strings.EqualFold(response.Header.Get("Content-Encoding"), "gzip") {
		return response.Body, nil
	}
	return gzip.NewReader(response.Body)
}

// heard holds off the idle watchdog while bytes arrive, keep-alive comments included.
type heard struct {
	reader      io.Reader
	watchdog    *time.Timer
	heardAt     *atomic.Int64
	idleTimeout time.Duration
}

func (h *heard) Read(buffer []byte) (int, error) {
	n, err := h.reader.Read(buffer)
	if n > 0 {
		h.heardAt.Store(time.Now().UnixNano())
		h.watchdog.Reset(h.idleTimeout)
	}
	return n, err
}

type parser struct {
	name string
	data []string
}

// feed takes one line of the event stream and returns the event a blank line ends, if it has data.
func (p *parser) feed(line string) (name, data string, ok bool) {
	if line == "" {
		name, data, ok = p.name, strings.Join(p.data, "\n"), len(p.data) > 0
		if name == "" {
			name = "message"
		}
		p.name, p.data = "", nil
		return name, data, ok
	}
	if strings.HasPrefix(line, ":") {
		return "", "", false
	}
	field, value, _ := strings.Cut(line, ":")
	value = strings.TrimPrefix(value, " ")
	switch field {
	case "event":
		p.name = value
	case "data":
		p.data = append(p.data, value)
	}
	return "", "", false
}

func closedBy(data string) error {
	var closed struct {
		Reason *string `json:"reason"`
		Status *int    `json:"status"`
	}
	reason, status := "unknown", 0
	if json.Unmarshal([]byte(data), &closed) == nil {
		if closed.Reason != nil {
			reason = *closed.Reason
		}
		if closed.Status != nil {
			status = *closed.Status
		}
	}
	switch {
	case reason == "not_live" && (status == 0 || status >= 200 && status < 300):
		return nil
	case reason == "not_live" || reason == "unreadable":
		return &StreamClosedError{Reason: reason, Status: status}
	default:
		return errReconnect
	}
}

// decode reads a snapshot as the read's answer: the event's data is the answer plus its event name.
func decode[T any](data string) (T, error) {
	var answer T
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte(data), &fields); err != nil {
		return answer, fmt.Errorf("live: a %s event is not a JSON object: %w", SnapshotEvent, err)
	}
	if fields == nil {
		return answer, fmt.Errorf("live: a %s event is not a JSON object: %s", SnapshotEvent, data)
	}
	delete(fields, "event")
	body, err := json.Marshal(fields)
	if err != nil {
		return answer, fmt.Errorf("live: a %s event: %w", SnapshotEvent, err)
	}
	if err := json.Unmarshal(body, &answer); err != nil {
		return answer, fmt.Errorf("live: a %s event is not a %T: %w", SnapshotEvent, answer, err)
	}
	return answer, nil
}

// The lines of an event stream, each ended by CRLF, LF or CR. A CR ends its line at once, so a
// line ended by a lone CR is not held for the next byte: an LF right after it is skipped instead.
// A line the body ends before its ending is dropped.
type lineReader struct {
	reader       *bufio.Reader
	skipLineFeed bool
}

func (l *lineReader) next() (string, error) {
	var line []byte
	for {
		b, err := l.reader.ReadByte()
		if err != nil {
			return "", err
		}
		if l.skipLineFeed {
			l.skipLineFeed = false
			if b == '\n' {
				continue
			}
		}
		switch b {
		case '\n':
			return string(line), nil
		case '\r':
			l.skipLineFeed = true
			return string(line), nil
		}
		line = append(line, b)
	}
}
