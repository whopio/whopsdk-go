package live_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	client "github.com/whopio/whopsdk-go/v2/client"
	live "github.com/whopio/whopsdk-go/v2/live"
	option "github.com/whopio/whopsdk-go/v2/option"
)

// checkLiveRead holds a live read that takes a request to its Watch and WaitFor (live_reads_test.go
// lists every live read): its type parameters bind to the read's, so Watch and WaitFor compile here
// only when they take the read's request and answer the read's answer type. With a sample for every
// field of the request, the read, Watch and WaitFor then send the same request through one
// recording server, but Accept, and fail with the same error.
func checkLiveRead[Resource, Request, Answer any](
	t *testing.T,
	resource func(*client.Whop) Resource,
	read func(Resource, context.Context, Request, ...option.RequestOption) (Answer, error),
	watch func(Resource, context.Context, Request, ...live.WatchOption) *live.Stream[Answer],
	waitFor func(Resource, context.Context, Request, func(Answer) bool, ...live.WaitForOption) (Answer, error),
) {
	t.Helper()
	request := sample[Request]()
	checkLiveReadWithoutParameters(
		t,
		resource,
		func(r Resource, ctx context.Context, opts ...option.RequestOption) (Answer, error) {
			return read(r, ctx, request, opts...)
		},
		func(r Resource, ctx context.Context, opts ...live.WatchOption) *live.Stream[Answer] {
			return watch(r, ctx, request, opts...)
		},
		func(r Resource, ctx context.Context, until func(Answer) bool, opts ...live.WaitForOption) (Answer, error) {
			return waitFor(r, ctx, request, until, opts...)
		},
	)
}

// checkLiveReadWithoutParameters is checkLiveRead for a live read that takes no parameters.
func checkLiveReadWithoutParameters[Resource, Answer any](
	t *testing.T,
	resource func(*client.Whop) Resource,
	read func(Resource, context.Context, ...option.RequestOption) (Answer, error),
	watch func(Resource, context.Context, ...live.WatchOption) *live.Stream[Answer],
	waitFor func(Resource, context.Context, func(Answer) bool, ...live.WaitForOption) (Answer, error),
) {
	t.Helper()
	var (
		mu       sync.Mutex
		requests []recorded
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mu.Lock()
		requests = append(requests, recorded{method: r.Method, url: r.URL.String(), body: string(body), header: r.Header.Clone()})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"error":{"message":"recorded"}}`)
	}))
	t.Cleanup(server.Close)
	whop := client.NewWhop(option.WithBaseURL(server.URL), option.WithToken("sk_test"), option.WithHTTPClient(server.Client()))
	requestOptions := []option.RequestOption{
		option.WithHTTPHeader(http.Header{"X-Trace": {"abc"}}),
		option.WithQueryParameters(url.Values{"extra": {"1"}}),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, readErr := read(resource(whop), ctx, requestOptions...)
	stream := watch(resource(whop), ctx, live.WithRequestOptions(requestOptions...))
	for range stream.Answers() {
		t.Fatal("watch answered a read that answers 404")
	}
	_, waitForErr := waitFor(resource(whop), ctx, func(Answer) bool { return true }, live.WithRequestOptions(requestOptions...))

	require.Error(t, readErr)
	for _, err := range []error{stream.Err(), waitForErr} {
		require.Equal(t, reflect.TypeOf(readErr), reflect.TypeOf(err))
		require.EqualError(t, err, readErr.Error())
	}
	mu.Lock()
	defer mu.Unlock()
	require.Len(t, requests, 3)
	sent := requests[0]
	require.Equal(t, "abc", sent.header.Get("X-Trace"))
	require.Contains(t, sent.url, "extra=1")
	for _, streamed := range requests[1:] {
		require.Equal(t, live.EventStream, streamed.header.Get("Accept"))
		require.Equal(t, sent.withoutAccept(), streamed.withoutAccept())
	}
}

type recorded struct {
	method string
	url    string
	body   string
	header http.Header
}

func (r recorded) withoutAccept() recorded {
	header := r.header.Clone()
	header.Del("Accept")
	return recorded{method: r.method, url: r.url, body: r.body, header: header}
}

// sample is a T with every exported field set, as deep as a request goes.
func sample[T any]() T {
	var value T
	fill(reflect.ValueOf(&value).Elem(), "sample", 0)
	return value
}

func fill(value reflect.Value, name string, depth int) {
	if depth > 8 {
		return
	}
	switch value.Kind() {
	case reflect.Pointer:
		value.Set(reflect.New(value.Type().Elem()))
		fill(value.Elem(), name, depth+1)
	case reflect.Struct:
		if value.Type() == reflect.TypeOf(time.Time{}) {
			value.Set(reflect.ValueOf(time.Date(2026, 11, 21, 12, 30, 0, 0, time.UTC)))
			return
		}
		for index := 0; index < value.NumField(); index++ {
			if field := value.Type().Field(index); field.IsExported() {
				fill(value.Field(index), field.Name, depth+1)
			}
		}
	case reflect.Slice:
		items := reflect.MakeSlice(value.Type(), 2, 2)
		fill(items.Index(0), name, depth+1)
		fill(items.Index(1), name+" again", depth+1)
		value.Set(items)
	case reflect.String:
		value.SetString(name + " 1/2")
	case reflect.Bool:
		value.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(7)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(7)
	case reflect.Float32, reflect.Float64:
		value.SetFloat(1.5)
	}
}
