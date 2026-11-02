package oidc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const roleID = "aurl_xxxxxxxxxxxxx"

func exchangeServer(t *testing.T, tokens ...string) (*httptest.Server, *[]map[string]string) {
	t.Helper()
	var mu sync.Mutex
	var bodies []map[string]string
	var next int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/access_tokens/exchange" {
			http.NotFound(w, r)
			return
		}
		var body map[string]string
		_ = json.NewDecoder(r.Body).Decode(&body)
		mu.Lock()
		bodies = append(bodies, body)
		mu.Unlock()
		token := tokens[atomic.AddInt32(&next, 1)-1]
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": token, "token_type": "Bearer", "expires_in": 3600})
	}))
	t.Cleanup(server.Close)
	return server, &bodies
}

type reply struct {
	status int
	header map[string]string
	body   string
}

func granted(token string, lifetime int) reply {
	return reply{status: http.StatusOK, body: fmt.Sprintf(`{"access_token":%q,"expires_in":%d}`, token, lifetime)}
}

func refused(status int, header map[string]string) reply {
	return reply{status: status, header: header, body: `{"error":{"message":"The token exchange is unavailable.","code":"unavailable"}}`}
}

func scriptedServer(t *testing.T, replies ...reply) (*httptest.Server, *int32) {
	t.Helper()
	var exchanges int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := int(atomic.AddInt32(&exchanges, 1))
		if n > len(replies) {
			t.Errorf("unexpected exchange %d", n)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for key, value := range replies[n-1].header {
			w.Header().Set(key, value)
		}
		w.WriteHeader(replies[n-1].status)
		_, _ = w.Write([]byte(replies[n-1].body))
	}))
	t.Cleanup(server.Close)
	return server, &exchanges
}

func epoch(at time.Time) string {
	return strconv.FormatInt(at.Unix(), 10)
}

func staticIdentity(token string) func() (string, error) {
	return func() (string, error) { return token, nil }
}

func freezeClock(t *testing.T) *time.Time {
	t.Helper()
	clock := time.Date(2026, 10, 23, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return clock }
	t.Cleanup(func() { now = time.Now })
	return &clock
}

func pinRandom(t *testing.T, value float64) {
	t.Helper()
	random = func() float64 { return value }
	t.Cleanup(func() { random = rand.Float64 })
}

func shortenHTTPTimeout(t *testing.T) {
	t.Helper()
	httpTimeout = 100 * time.Millisecond
	t.Cleanup(func() { httpTimeout = 10 * time.Second })
}

func shortenRenewalTimeout(t *testing.T) {
	t.Helper()
	renewalTimeout = 100 * time.Millisecond
	t.Cleanup(func() { renewalTimeout = 30 * time.Second })
}

func stalledServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	release := make(chan struct{})
	var requests int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&requests, 1) == 1 {
			select {
			case <-release:
			case <-r.Context().Done():
			}
			return
		}
		handler(w, r)
	}))
	t.Cleanup(server.Close)
	t.Cleanup(func() { close(release) })
	return server
}

func sourceAndToken(t *testing.T, options Options) (*tokenSource, func() (string, error)) {
	t.Helper()
	httpClient, err := NewHTTPClient(options)
	if err != nil {
		t.Fatal(err)
	}
	source := httpClient.source
	t.Cleanup(func() { awaitRenewal(t, source) })
	return source, func() (string, error) { return source.token(context.Background()) }
}

func tokenFunc(t *testing.T, options Options) func() (string, error) {
	t.Helper()
	_, token := sourceAndToken(t, options)
	return token
}

func awaitRenewal(t *testing.T, source *tokenSource) {
	t.Helper()
	source.mu.Lock()
	call := source.inFlight
	source.mu.Unlock()
	if call == nil {
		return
	}
	select {
	case <-call.done:
	case <-time.After(5 * time.Second):
		t.Error("the renewal is still running")
	}
}

func TestTokenExchangesTheIdentityTokenForAnAccessTokenAssumingTheRole(t *testing.T) {
	server, bodies := exchangeServer(t, "whop-token-1")

	token, err := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL + "/api/v1/"})()
	if err != nil {
		t.Fatal(err)
	}
	if token != "whop-token-1" {
		t.Fatalf("token = %q", token)
	}
	want := map[string]string{
		"grant_type":         "urn:ietf:params:oauth:grant-type:token-exchange",
		"subject_token_type": "urn:ietf:params:oauth:token-type:id_token",
		"subject_token":      "id-token",
		"role_id":            roleID,
	}
	got := (*bodies)[0]
	for key, value := range want {
		if got[key] != value {
			t.Fatalf("body[%s] = %q, want %q", key, got[key], value)
		}
	}
}

func TestTokenRenewsAtARandomPointTwoToTenMinutesBeforeExpiryWithAFreshIdentityToken(t *testing.T) {
	server, bodies := exchangeServer(t, "whop-token-1", "whop-token-2")
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	identities := []string{"id-token-1", "id-token-2"}
	identity := func() (string, error) {
		next := identities[0]
		identities = identities[1:]
		return next, nil
	}
	source, token := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: identity, BaseURL: server.URL + "/api/v1"})

	first, _ := token()
	*clock = clock.Add(54*time.Minute - time.Second)
	reused, _ := token()
	if first != "whop-token-1" || reused != "whop-token-1" || len(*bodies) != 1 {
		t.Fatalf("first = %q, reused = %q, exchanges = %d", first, reused, len(*bodies))
	}

	*clock = clock.Add(time.Second)
	renewing, _ := token()
	awaitRenewal(t, source)
	renewed, _ := token()
	if renewing != "whop-token-1" || renewed != "whop-token-2" || (*bodies)[1]["subject_token"] != "id-token-2" {
		t.Fatalf("renewing = %q, renewed = %q, subject = %q", renewing, renewed, (*bodies)[1]["subject_token"])
	}
}

func TestTokenRenewsAShortLivedTokenHalfwayThroughItsLife(t *testing.T) {
	var exchanges int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"access_token":"whop-token-` + strconv.Itoa(int(atomic.AddInt32(&exchanges, 1))) + `","expires_in":120}`))
	}))
	t.Cleanup(server.Close)
	clock := freezeClock(t)
	pinRandom(t, 0)
	source, token := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	_, _ = token()
	*clock = clock.Add(time.Minute - time.Second)
	reused, _ := token()
	*clock = clock.Add(time.Second)
	renewing, _ := token()
	awaitRenewal(t, source)
	renewed, _ := token()
	if reused != "whop-token-1" || renewing != "whop-token-1" || renewed != "whop-token-2" || atomic.LoadInt32(&exchanges) != 2 {
		t.Fatalf("reused = %q, renewing = %q, renewed = %q, exchanges = %d", reused, renewing, renewed, atomic.LoadInt32(&exchanges))
	}
}

func TestTokenSharesOneExchangeBetweenConcurrentCallers(t *testing.T) {
	server, bodies := exchangeServer(t, "whop-token-1", "unexpected")
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL + "/api/v1"})

	var wg sync.WaitGroup
	results := make([]string, 5)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i], _ = token()
		}(i)
	}
	wg.Wait()

	for _, result := range results {
		if result != "whop-token-1" {
			t.Fatalf("results = %v", results)
		}
	}
	if len(*bodies) != 1 {
		t.Fatalf("exchanges = %d", len(*bodies))
	}
}

func TestTokenSurfacesTheAPIErrorWithoutAnotherExchangeUntilItsRetryTime(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":{"type":"bad_request","message":"No trust linked to this role accepts this token.","code":"oidc_exchange_no_match"}}`))
			return
		}
		_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
	}))
	t.Cleanup(server.Close)
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	_, err := token()
	var oidcErr *Error
	if !errors.As(err, &oidcErr) || oidcErr.Status != 400 || oidcErr.Code != "oidc_exchange_no_match" || oidcErr.Message != "No trust linked to this role accepts this token." || !oidcErr.RetryAt.Equal(clock.Add(2*time.Second)) {
		t.Fatalf("err = %#v", err)
	}
	*clock = clock.Add(2*time.Second - time.Millisecond)
	for i := 0; i < 5; i++ {
		if _, again := token(); again != err {
			t.Fatalf("again = %#v", again)
		}
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("exchanges = %d", got)
	}

	*clock = clock.Add(time.Millisecond)
	if retried, err := token(); err != nil || retried != "whop-token-1" {
		t.Fatalf("retried = %q, err = %v", retried, err)
	}
}

func TestAFailedRenewalKeepsServingTheValidTokenAndRetriesAfterTheBackoff(t *testing.T) {
	server, exchanges := scriptedServer(t, granted("whop-token-1", 3600), refused(http.StatusInternalServerError, nil), refused(http.StatusInternalServerError, nil), granted("whop-token-2", 3600))
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	source, token := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	if first, err := token(); err != nil || first != "whop-token-1" {
		t.Fatalf("first = %q, err = %v", first, err)
	}
	steps := []struct {
		advance   time.Duration
		exchanges int32
	}{
		{advance: 54 * time.Minute, exchanges: 2},
		{advance: 2*time.Second - time.Millisecond, exchanges: 2},
		{advance: time.Millisecond, exchanges: 3},
		{advance: 4*time.Second - time.Millisecond, exchanges: 3},
		{advance: time.Millisecond, exchanges: 4},
	}
	for i, step := range steps {
		*clock = clock.Add(step.advance)
		got, err := token()
		awaitRenewal(t, source)
		if err != nil || got != "whop-token-1" || atomic.LoadInt32(exchanges) != step.exchanges {
			t.Fatalf("step %d: token = %q, err = %v, exchanges = %d", i, got, err, atomic.LoadInt32(exchanges))
		}
	}
	if renewed, err := token(); err != nil || renewed != "whop-token-2" {
		t.Fatalf("renewed = %q, err = %v", renewed, err)
	}
}

func TestAFailedRenewalServesTheTokenUntilThirtySecondsBeforeItExpiresAndThenTheCachedError(t *testing.T) {
	server, exchanges := scriptedServer(t, granted("whop-token-1", 120), refused(http.StatusServiceUnavailable, map[string]string{"Retry-After": "60"}), granted("whop-token-2", 3600))
	clock := freezeClock(t)
	start := *clock
	pinRandom(t, 0)
	source, token := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	_, _ = token()
	*clock = start.Add(time.Minute)
	got, err := token()
	awaitRenewal(t, source)
	if err != nil || got != "whop-token-1" {
		t.Fatalf("at renewal: token = %q, err = %v", got, err)
	}
	*clock = start.Add(90*time.Second - time.Millisecond)
	if got, err := token(); err != nil || got != "whop-token-1" {
		t.Fatalf("31 seconds before expiry: token = %q, err = %v", got, err)
	}

	*clock = start.Add(90 * time.Second)
	_, err = token()
	var oidcErr *Error
	if !errors.As(err, &oidcErr) || oidcErr.Status != http.StatusServiceUnavailable || !oidcErr.RetryAt.Equal(start.Add(2*time.Minute)) || atomic.LoadInt32(exchanges) != 2 {
		t.Fatalf("30 seconds before expiry: err = %#v, exchanges = %d", err, atomic.LoadInt32(exchanges))
	}
	*clock = start.Add(2 * time.Minute)
	if got, err := token(); err != nil || got != "whop-token-2" || atomic.LoadInt32(exchanges) != 3 {
		t.Fatalf("after the retry time: token = %q, err = %v, exchanges = %d", got, err, atomic.LoadInt32(exchanges))
	}
}

func TestARenewalThatFailsInsideTheLastThirtySecondsReturnsTheErrorWithItsRetryTime(t *testing.T) {
	server, _ := scriptedServer(t, granted("whop-token-1", 40), refused(http.StatusInternalServerError, nil))
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	_, _ = token()
	*clock = clock.Add(20 * time.Second)
	got, err := token()
	var oidcErr *Error
	if got != "" || !errors.As(err, &oidcErr) || oidcErr.Status != http.StatusInternalServerError || !oidcErr.RetryAt.Equal(clock.Add(2*time.Second)) {
		t.Fatalf("token = %q, err = %#v", got, err)
	}
}

func TestAFailedExchangeWaitsForTheServersRetryHintUpToAMinute(t *testing.T) {
	rateLimited := `{"message":"You are being rate limited"}`
	cases := []struct {
		name     string
		status   int
		header   func(start time.Time) map[string]string
		body     string
		wantWait time.Duration
	}{
		{name: "Retry-After in seconds", status: http.StatusTooManyRequests, header: func(time.Time) map[string]string { return map[string]string{"Retry-After": "37"} }, wantWait: 37 * time.Second},
		{name: "Retry-After as an HTTP date", status: http.StatusTooManyRequests, header: func(start time.Time) map[string]string {
			return map[string]string{"Retry-After": start.Add(45 * time.Second).Format(http.TimeFormat)}
		}, wantWait: 45 * time.Second},
		{name: "Retry-After on a 503", status: http.StatusServiceUnavailable, header: func(time.Time) map[string]string { return map[string]string{"Retry-After": "20"} }, wantWait: 20 * time.Second},
		{name: "Retry-After before RateLimit-Reset", status: http.StatusTooManyRequests, header: func(start time.Time) map[string]string {
			return map[string]string{"Retry-After": "20", "RateLimit-Reset": epoch(start.Add(50 * time.Second))}
		}, wantWait: 20 * time.Second},
		{name: "RateLimit-Reset in epoch seconds", status: http.StatusTooManyRequests, header: func(start time.Time) map[string]string {
			return map[string]string{"RateLimit-Reset": epoch(start.Add(37 * time.Second))}
		}, wantWait: 37 * time.Second},
		{name: "unreadable Retry-After falls back to RateLimit-Reset", status: http.StatusTooManyRequests, header: func(start time.Time) map[string]string {
			return map[string]string{"Retry-After": "soon", "RateLimit-Reset": epoch(start.Add(30 * time.Second))}
		}, wantWait: 30 * time.Second},
		{name: "Retry-After over a minute", status: http.StatusTooManyRequests, header: func(time.Time) map[string]string { return map[string]string{"Retry-After": "300"} }, wantWait: time.Minute},
		{name: "RateLimit-Reset over a minute away", status: http.StatusTooManyRequests, header: func(start time.Time) map[string]string {
			return map[string]string{"RateLimit-Reset": epoch(start.Add(5 * time.Minute))}
		}, wantWait: time.Minute},
		{name: "RateLimit-Reset already passed", status: http.StatusTooManyRequests, header: func(start time.Time) map[string]string {
			return map[string]string{"RateLimit-Reset": epoch(start.Add(-10 * time.Second))}
		}, wantWait: 2 * time.Second},
		{name: "trust limit without a hint", status: http.StatusTooManyRequests, header: func(time.Time) map[string]string { return nil }, body: `{"error":{"message":"Too many token exchanges for this trust. Try again in a minute.","code":"oidc_exchange_rate_limited"}}`, wantWait: 2 * time.Second},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			clock := freezeClock(t)
			start := *clock
			pinRandom(t, 0.5)
			body := tc.body
			if body == "" {
				body = rateLimited
			}
			server, exchanges := scriptedServer(t, reply{status: tc.status, header: tc.header(start), body: body}, granted("whop-token-1", 3600))
			token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

			_, err := token()
			var oidcErr *Error
			if !errors.As(err, &oidcErr) || oidcErr.Status != tc.status || !oidcErr.RetryAt.Equal(start.Add(tc.wantWait)) {
				t.Fatalf("err = %#v", err)
			}
			*clock = start.Add(tc.wantWait - time.Millisecond)
			if _, again := token(); again != err || atomic.LoadInt32(exchanges) != 1 {
				t.Fatalf("again = %#v, exchanges = %d", again, atomic.LoadInt32(exchanges))
			}
			*clock = start.Add(tc.wantWait)
			if retried, err := token(); err != nil || retried != "whop-token-1" || atomic.LoadInt32(exchanges) != 2 {
				t.Fatalf("retried = %q, err = %v, exchanges = %d", retried, err, atomic.LoadInt32(exchanges))
			}
		})
	}
}

func TestTheBackoffDoublesFromTwoSecondsToAMinuteAndRestartsAfterASuccess(t *testing.T) {
	waits := []time.Duration{2 * time.Second, 4 * time.Second, 8 * time.Second, 16 * time.Second, 32 * time.Second, time.Minute, time.Minute}
	var replies []reply
	for range waits {
		replies = append(replies, refused(http.StatusInternalServerError, nil))
	}
	server, _ := scriptedServer(t, append(replies, granted("whop-token-1", 40), refused(http.StatusInternalServerError, nil))...)
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	for i, wait := range waits {
		_, err := token()
		var oidcErr *Error
		if !errors.As(err, &oidcErr) || !oidcErr.RetryAt.Equal(clock.Add(wait)) {
			t.Fatalf("failure %d: err = %#v, want a retry in %v", i+1, err, wait)
		}
		*clock = clock.Add(wait)
	}
	if got, err := token(); err != nil || got != "whop-token-1" {
		t.Fatalf("token = %q, err = %v", got, err)
	}
	*clock = clock.Add(20 * time.Second)
	_, err := token()
	var oidcErr *Error
	if !errors.As(err, &oidcErr) || !oidcErr.RetryAt.Equal(clock.Add(2*time.Second)) {
		t.Fatalf("after a success: err = %#v", err)
	}
}

func TestTheBackoffVariesByTenPercent(t *testing.T) {
	for _, tc := range []struct {
		random float64
		wait   time.Duration
	}{{random: 0, wait: 1800 * time.Millisecond}, {random: 1, wait: 2200 * time.Millisecond}} {
		t.Run(fmt.Sprint(tc.random), func(t *testing.T) {
			server, _ := scriptedServer(t, refused(http.StatusInternalServerError, nil))
			clock := freezeClock(t)
			pinRandom(t, tc.random)

			_, err := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})()
			var oidcErr *Error
			if !errors.As(err, &oidcErr) || !oidcErr.RetryAt.Equal(clock.Add(tc.wait)) {
				t.Fatalf("err = %#v", err)
			}
		})
	}
}

func TestAFailureFromOutsideThePackageIsWrappedWithItsRetryTime(t *testing.T) {
	unreachable := errors.New("metadata server unreachable")
	var identityCalls int32
	identity := func() (string, error) {
		if atomic.AddInt32(&identityCalls, 1) == 1 {
			return "", unreachable
		}
		return "id-token", nil
	}
	server, _ := scriptedServer(t, granted("whop-token-1", 3600))
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: identity, BaseURL: server.URL})

	_, err := token()
	var oidcErr *Error
	if !errors.As(err, &oidcErr) || !errors.Is(err, unreachable) || oidcErr.Message != unreachable.Error() || !oidcErr.RetryAt.Equal(clock.Add(2*time.Second)) {
		t.Fatalf("err = %#v", err)
	}
	*clock = clock.Add(2*time.Second - time.Millisecond)
	if _, again := token(); again != err || atomic.LoadInt32(&identityCalls) != 1 {
		t.Fatalf("again = %#v, identity calls = %d", again, atomic.LoadInt32(&identityCalls))
	}
	*clock = clock.Add(time.Millisecond)
	if got, err := token(); err != nil || got != "whop-token-1" {
		t.Fatalf("token = %q, err = %v", got, err)
	}
}

func TestAStalledExchangeTimesOutWithACustomHTTPClientAndTheNextOneStarts(t *testing.T) {
	shortenHTTPTimeout(t)
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	server := stalledServer(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
	})
	httpClient, err := NewHTTPClient(Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL, HTTPClient: &http.Client{}})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = httpClient.source.token(ctx)
	var oidcErr *Error
	if !errors.As(err, &oidcErr) || oidcErr.Message != "The token exchange timed out after 0.1 seconds." || !oidcErr.RetryAt.Equal(clock.Add(2*time.Second)) {
		t.Fatalf("err = %#v", err)
	}
	*clock = clock.Add(2 * time.Second)
	if token, err := httpClient.source.token(ctx); err != nil || token != "whop-token-1" {
		t.Fatalf("token = %q, err = %v", token, err)
	}
}

func TestConcurrentCallersShareOneFailedExchange(t *testing.T) {
	var exchanges int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&exchanges, 1)
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"error":{"message":"No trust linked to this role accepts this token.","code":"oidc_exchange_no_match"}}`))
	}))
	t.Cleanup(server.Close)
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	var wg sync.WaitGroup
	errs := make([]error, 8)
	for i := range errs {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, errs[i] = token()
		}(i)
	}
	wg.Wait()

	if got := atomic.LoadInt32(&exchanges); got != 1 {
		t.Fatalf("exchanges = %d", got)
	}
	for _, err := range errs {
		var oidcErr *Error
		if err != errs[0] || !errors.As(err, &oidcErr) || oidcErr.Code != "oidc_exchange_no_match" || oidcErr.RetryAt.IsZero() {
			t.Fatalf("errs = %v", errs)
		}
	}
}

func TestConcurrentCallersShareAFailedRenewalAndAllKeepTheValidToken(t *testing.T) {
	var exchanges int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if atomic.AddInt32(&exchanges, 1) == 1 {
			_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
			return
		}
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	source, token := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})
	_, _ = token()
	*clock = clock.Add(54 * time.Minute)

	var wg sync.WaitGroup
	results := make([]string, 8)
	errs := make([]error, 8)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i], errs[i] = token()
		}(i)
	}
	wg.Wait()
	awaitRenewal(t, source)

	if got := atomic.LoadInt32(&exchanges); got != 2 {
		t.Fatalf("exchanges = %d", got)
	}
	for i := range results {
		if results[i] != "whop-token-1" || errs[i] != nil {
			t.Fatalf("results = %v, errs = %v", results, errs)
		}
	}
}

func TestCallersPastTheRenewalPointGetTheValidTokenWithoutWaitingForASlowRenewal(t *testing.T) {
	server, exchanges := scriptedServer(t, granted("whop-token-1", 3600), granted("whop-token-2", 3600))
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	renewalStarted := make(chan struct{})
	release := make(chan struct{})
	var identityCalls int32
	identity := func() (string, error) {
		if atomic.AddInt32(&identityCalls, 1) == 2 {
			close(renewalStarted)
			<-release
		}
		return "id-token", nil
	}
	source, token := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: identity, BaseURL: server.URL})
	releaseRenewal := sync.OnceFunc(func() { close(release) })
	t.Cleanup(releaseRenewal)
	if first, err := token(); err != nil || first != "whop-token-1" {
		t.Fatalf("first = %q, err = %v", first, err)
	}
	*clock = clock.Add(54 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if got, err := source.token(ctx); err != nil || got != "whop-token-1" {
		t.Fatalf("at renewal: token = %q, err = %v", got, err)
	}
	<-renewalStarted
	for i := 0; i < 5; i++ {
		if got, err := source.token(ctx); err != nil || got != "whop-token-1" {
			t.Fatalf("while renewing %d: token = %q, err = %v", i, got, err)
		}
	}

	releaseRenewal()
	awaitRenewal(t, source)
	if renewed, err := token(); err != nil || renewed != "whop-token-2" || atomic.LoadInt32(exchanges) != 2 || atomic.LoadInt32(&identityCalls) != 2 {
		t.Fatalf("renewed = %q, err = %v, exchanges = %d, identity calls = %d", renewed, err, atomic.LoadInt32(exchanges), atomic.LoadInt32(&identityCalls))
	}
}

func TestARenewalFailsAtItsDeadlineAndTheNextOneStartsAfterTheBackoff(t *testing.T) {
	never := make(chan struct{})
	cases := []struct {
		name     string
		identity func(calls *int32) func() (string, error)
		server   func(t *testing.T) *httptest.Server
	}{
		{
			name: "a hung identity token source",
			identity: func(calls *int32) func() (string, error) {
				return func() (string, error) {
					if atomic.AddInt32(calls, 1) == 1 {
						<-never
					}
					return "id-token", nil
				}
			},
			server: func(t *testing.T) *httptest.Server {
				server, _ := scriptedServer(t, granted("whop-token-1", 3600))
				return server
			},
		},
		{
			name: "a hung exchange",
			identity: func(calls *int32) func() (string, error) {
				return func() (string, error) {
					atomic.AddInt32(calls, 1)
					return "id-token", nil
				}
			},
			server: func(t *testing.T) *httptest.Server {
				return stalledServer(t, func(w http.ResponseWriter, _ *http.Request) {
					_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
				})
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			shortenRenewalTimeout(t)
			clock := freezeClock(t)
			pinRandom(t, 0.5)
			var identityCalls int32
			source, _ := sourceAndToken(t, Options{RoleID: roleID, IdentityToken: tc.identity(&identityCalls), BaseURL: tc.server(t).URL})
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := source.token(ctx)
			var oidcErr *Error
			if !errors.As(err, &oidcErr) || oidcErr.Message != "The token renewal timed out after 0.1 seconds." || !oidcErr.RetryAt.Equal(clock.Add(2*time.Second)) {
				t.Fatalf("err = %#v", err)
			}
			*clock = clock.Add(2 * time.Second)
			if token, err := source.token(ctx); err != nil || token != "whop-token-1" || atomic.LoadInt32(&identityCalls) != 2 {
				t.Fatalf("token = %q, err = %v, identity calls = %d", token, err, atomic.LoadInt32(&identityCalls))
			}
		})
	}
}

func TestCallersWithNoValidTokenWaitForTheRenewalAndShareItsOutcome(t *testing.T) {
	var exchanges int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := atomic.AddInt32(&exchanges, 1)
		if n > 1 {
			time.Sleep(200 * time.Millisecond)
		}
		_, _ = fmt.Fprintf(w, `{"access_token":"whop-token-%d","expires_in":40}`, n)
	}))
	t.Cleanup(server.Close)
	clock := freezeClock(t)
	pinRandom(t, 0.5)
	token := tokenFunc(t, Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})
	_, _ = token()
	*clock = clock.Add(20 * time.Second)

	var wg sync.WaitGroup
	results := make([]string, 8)
	errs := make([]error, 8)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i], errs[i] = token()
		}(i)
	}
	wg.Wait()

	if got := atomic.LoadInt32(&exchanges); got != 2 {
		t.Fatalf("exchanges = %d", got)
	}
	for i := range results {
		if results[i] != "whop-token-2" || errs[i] != nil {
			t.Fatalf("results = %v, errs = %v", results, errs)
		}
	}
}

func TestACancelledRequestStopsWaitingWhileTheExchangeFinishesForTheNext(t *testing.T) {
	var exchanges int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&exchanges, 1)
		_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
	}))
	t.Cleanup(server.Close)
	release := make(chan struct{})
	identity := func() (string, error) {
		<-release
		return "id-token", nil
	}
	httpClient, err := NewHTTPClient(Options{RoleID: roleID, IdentityToken: identity, BaseURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := httpClient.source.token(cancelled); !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v", err)
	}
	close(release)
	if token, err := httpClient.source.token(context.Background()); err != nil || token != "whop-token-1" {
		t.Fatalf("token = %q, err = %v", token, err)
	}
	if got := atomic.LoadInt32(&exchanges); got != 1 {
		t.Fatalf("exchanges = %d", got)
	}
}

func TestNewHTTPClientRequiresARoleAndAnIdentitySource(t *testing.T) {
	var oidcErr *Error
	if _, err := NewHTTPClient(Options{IdentityToken: staticIdentity("id-token")}); !errors.As(err, &oidcErr) || oidcErr.Message != "Pass the role to assume as RoleID." {
		t.Fatalf("err = %v", err)
	}
	if _, err := NewHTTPClient(Options{RoleID: roleID}); !errors.As(err, &oidcErr) || oidcErr.Message != "Pass a source of identity tokens as IdentityToken." {
		t.Fatalf("err = %v", err)
	}
}

func TestGitHubActionsRequestsTheWhopAudience(t *testing.T) {
	var gotAudience, gotVersion, gotAuthorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAudience = r.URL.Query().Get("audience")
		gotVersion = r.URL.Query().Get("api-version")
		gotAuthorization = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"value":"github-id-token"}`))
	}))
	t.Cleanup(server.Close)
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", server.URL+"/idtoken?api-version=2.0")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "request-token")

	token, err := GitHubActions(IdentityOptions{})()
	if err != nil || token != "github-id-token" {
		t.Fatalf("token = %q, err = %v", token, err)
	}
	if gotAudience != DefaultAudience || gotVersion != "2.0" || gotAuthorization != "Bearer request-token" {
		t.Fatalf("audience = %q, api-version = %q, authorization = %q", gotAudience, gotVersion, gotAuthorization)
	}
}

func TestGitHubActionsExplainsTheMissingPermission(t *testing.T) {
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")

	_, err := GitHubActions(IdentityOptions{})()
	if err == nil || err.Error() != "GitHub Actions did not provide an identity token. Give the job permissions: id-token: write." {
		t.Fatalf("err = %v", err)
	}
}

func TestTheIdentityTokenSourcesTimeOutWithACustomHTTPClient(t *testing.T) {
	shortenHTTPTimeout(t)
	server := stalledServer(t, http.NotFound)
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", server.URL+"/idtoken")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "request-token")
	stalledMetadataServer := &http.Client{Transport: roundTripper(func(r *http.Request) (*http.Response, error) {
		<-r.Context().Done()
		return nil, r.Context().Err()
	})}

	if _, err := GitHubActions(IdentityOptions{HTTPClient: &http.Client{}})(); err == nil || err.Error() != "The GitHub Actions identity token request timed out after 0.1 seconds." {
		t.Fatalf("GitHubActions err = %v", err)
	}
	if _, err := GoogleCloud(IdentityOptions{HTTPClient: stalledMetadataServer})(); err == nil || err.Error() != "The identity token request to the Google Cloud metadata server timed out after 0.1 seconds." {
		t.Fatalf("GoogleCloud err = %v", err)
	}
}

func TestGoogleCloudAsksTheMetadataServer(t *testing.T) {
	var gotURL, gotFlavor string
	client := &http.Client{Transport: roundTripper(func(r *http.Request) (*http.Response, error) {
		gotURL = r.URL.String()
		gotFlavor = r.Header.Get("Metadata-Flavor")
		recorder := httptest.NewRecorder()
		_, _ = recorder.WriteString("google-id-token\n")
		return recorder.Result(), nil
	})}

	token, err := GoogleCloud(IdentityOptions{Audience: "https://example.com", HTTPClient: client})()
	if err != nil || token != "google-id-token" {
		t.Fatalf("token = %q, err = %v", token, err)
	}
	if gotURL != "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?audience=https%3A%2F%2Fexample.com" || gotFlavor != "Google" {
		t.Fatalf("url = %q, flavor = %q", gotURL, gotFlavor)
	}
}

type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
