package oidc

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
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

func staticIdentity(token string) func() (string, error) {
	return func() (string, error) { return token, nil }
}

func TestTokenExchangesTheIdentityTokenForAnAccessTokenAssumingTheRole(t *testing.T) {
	server, bodies := exchangeServer(t, "whop-token-1")

	token, err := Token(Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL + "/api/v1/"})()
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

func TestTokenRenewsAMinuteBeforeExpiryWithAFreshIdentityToken(t *testing.T) {
	server, bodies := exchangeServer(t, "whop-token-1", "whop-token-2")
	clock := time.Date(2026, 10, 23, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return clock }
	t.Cleanup(func() { now = time.Now })
	identities := []string{"id-token-1", "id-token-2"}
	identity := func() (string, error) {
		next := identities[0]
		identities = identities[1:]
		return next, nil
	}
	token := Token(Options{RoleID: roleID, IdentityToken: identity, BaseURL: server.URL + "/api/v1"})

	first, _ := token()
	clock = clock.Add(58 * time.Minute)
	reused, _ := token()
	if first != "whop-token-1" || reused != "whop-token-1" || len(*bodies) != 1 {
		t.Fatalf("first = %q, reused = %q, exchanges = %d", first, reused, len(*bodies))
	}

	clock = clock.Add(2 * time.Minute)
	renewed, _ := token()
	if renewed != "whop-token-2" || (*bodies)[1]["subject_token"] != "id-token-2" {
		t.Fatalf("renewed = %q, subject = %q", renewed, (*bodies)[1]["subject_token"])
	}
}

func TestTokenSharesOneExchangeBetweenConcurrentCallers(t *testing.T) {
	server, bodies := exchangeServer(t, "whop-token-1", "unexpected")
	token := Token(Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL + "/api/v1"})

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

func TestTokenSurfacesTheAPIErrorAndRetriesOnTheNextCall(t *testing.T) {
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
	token := Token(Options{RoleID: roleID, IdentityToken: staticIdentity("id-token"), BaseURL: server.URL})

	_, err := token()
	var oidcErr *Error
	if !errors.As(err, &oidcErr) || oidcErr.Status != 400 || oidcErr.Code != "oidc_exchange_no_match" || oidcErr.Message != "No trust linked to this role accepts this token." {
		t.Fatalf("err = %#v", err)
	}
	if retried, err := token(); err != nil || retried != "whop-token-1" {
		t.Fatalf("retried = %q, err = %v", retried, err)
	}
}

func TestTokenRequiresARoleAndAnIdentitySource(t *testing.T) {
	if _, err := Token(Options{IdentityToken: staticIdentity("id-token")})(); err == nil {
		t.Fatal("expected an error without RoleID")
	}
	if _, err := Token(Options{RoleID: roleID})(); err == nil {
		t.Fatal("expected an error without IdentityToken")
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
