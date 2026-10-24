// Package oidc authenticates the Whop client with a cloud provider's OIDC identity token instead
// of an API key.
//
// NewHTTPClient returns an HTTP client for option.WithHTTPClient. Before each request it exchanges
// an identity token for a short-lived Whop access token, reuses that token until a few minutes before
// it expires, and then exchanges a fresh identity token. When a renewal fails, requests keep using the
// current token until 30 seconds before it expires, and the exchange is retried after a backoff. When
// no valid token remains, the request fails with the exchange's error and is not sent, and later
// requests return the same error without another exchange until its RetryAt. The package uses only the
// standard library, never generated client code, so it survives the client being regenerated.
//
//	httpClient, err := oidc.NewHTTPClient(oidc.Options{
//		RoleID:        "aurl_xxxxxxxxxxxxx",
//		IdentityToken: oidc.GitHubActions(oidc.IdentityOptions{}),
//	})
//	if err != nil {
//		return err
//	}
//	whop := client.NewWhop(option.WithHTTPClient(httpClient))
package oidc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DefaultAudience is the audience GitHubActions and GoogleCloud request, and the one to set on your trust.
const DefaultAudience = "https://api.whop.com"

const (
	defaultBaseURL            = "https://api.whop.com/api/v1"
	grantType                 = "urn:ietf:params:oauth:grant-type:token-exchange"
	idTokenType               = "urn:ietf:params:oauth:token-type:id_token"
	renewBeforeExpiryMin      = 2 * time.Minute
	renewBeforeExpiryMax      = 10 * time.Minute
	expiryMargin              = 30 * time.Second
	firstBackoff              = 2 * time.Second
	maxRetryWait              = time.Minute
	googleMetadataIdentityURL = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity"
)

var (
	now         = time.Now
	random      = rand.Float64
	httpTimeout = 10 * time.Second
)

// Options configures NewHTTPClient.
type Options struct {
	// RoleID is the role to assume, prefixed aurl_. A trust linked to it must accept the identity token.
	RoleID string
	// IdentityToken returns an identity token from your cloud provider. It is called again for every
	// exchange, so return a fresh token rather than a stored one.
	IdentityToken func() (string, error)
	// BaseURL is the API's base URL, for example https://sandbox-api.whop.com/api/v1. Defaults to production.
	BaseURL string
	// HTTPClient sends the exchange and the API requests. Defaults to http.DefaultClient, as the Whop
	// client does. The exchange is cut off after 10 seconds whichever client sends it.
	HTTPClient *http.Client
}

// IdentityOptions configures the built-in identity token sources.
type IdentityOptions struct {
	// Audience is the audience to request. It must equal the audience on your trust. Defaults to DefaultAudience.
	Audience string
	// HTTPClient sends the request. Defaults to http.DefaultClient. The request is cut off after 10
	// seconds whichever client sends it.
	HTTPClient *http.Client
}

// Error reports a failed exchange or identity token request. Code is the API's error code when it sent one.
//
// On every error a request returns because the exchange failed, RetryAt is when the next exchange may
// start: the API's Retry-After or RateLimit-Reset, up to a minute, or else a backoff that starts at 2
// seconds and doubles up to a minute. Until then requests fail with this error without an exchange. A
// failure from outside this package, such as a network error or an error from your IdentityToken, is
// wrapped in an Error, and Unwrap returns it.
type Error struct {
	Status  int
	Code    string
	Message string
	RetryAt time.Time
	err     error
}

func (e *Error) Error() string { return e.Message }

func (e *Error) Unwrap() error { return e.err }

// HTTPClient sends the Whop client's requests with an access token from the exchange. Pass it to
// option.WithHTTPClient. A request-level option.WithHTTPClient replaces it, and its authentication,
// for that request.
type HTTPClient struct {
	source *tokenSource
	send   *http.Client
}

// NewHTTPClient returns an HTTP client for option.WithHTTPClient that exchanges your provider's identity
// token for a Whop access token and renews it before it expires. Concurrent requests share one exchange
// and its outcome.
func NewHTTPClient(options Options) (*HTTPClient, error) {
	if options.RoleID == "" {
		return nil, &Error{Message: "Pass the role to assume as RoleID."}
	}
	if options.IdentityToken == nil {
		return nil, &Error{Message: "Pass a source of identity tokens as IdentityToken."}
	}
	return &HTTPClient{source: &tokenSource{options: options}, send: httpClient(options.HTTPClient)}, nil
}

// Do sends the request with the current access token, exchanging for one first when there is none or it
// is about to expire. When the exchange fails and no valid token remains, Do returns its error and does
// not send the request.
func (c *HTTPClient) Do(request *http.Request) (*http.Response, error) {
	token, err := c.source.token(request.Context())
	if err != nil {
		return nil, err
	}
	authenticated := request.Clone(request.Context())
	authenticated.Header.Set("Authorization", "Bearer "+token)
	return c.send.Do(authenticated)
}

type tokenSource struct {
	options             Options
	mu                  sync.Mutex
	current             string
	renewAt             time.Time
	expiresAt           time.Time
	failure             *Error
	consecutiveFailures int
	inFlight            *exchangeCall
}

type grant struct {
	token    string
	lifetime time.Duration
}

type exchangeCall struct {
	done  chan struct{}
	token string
	err   error
}

func (s *tokenSource) token(ctx context.Context) (string, error) {
	s.mu.Lock()
	if s.current != "" && now().Before(s.renewAt) {
		token := s.current
		s.mu.Unlock()
		return token, nil
	}
	if s.failure != nil && now().Before(s.failure.RetryAt) {
		token, err := s.currentOrFailure()
		s.mu.Unlock()
		return token, err
	}
	call := s.inFlight
	if call == nil {
		call = &exchangeCall{done: make(chan struct{})}
		s.inFlight = call
		go s.renew(call)
	}
	s.mu.Unlock()

	select {
	case <-call.done:
		return call.token, call.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (s *tokenSource) renew(call *exchangeCall) {
	granted, err := s.exchange()
	s.mu.Lock()
	if err == nil {
		s.current, s.failure, s.consecutiveFailures = granted.token, nil, 0
		s.expiresAt = now().Add(granted.lifetime)
		s.renewAt = s.expiresAt.Add(-renewBeforeExpiry(granted.lifetime))
		call.token = granted.token
	} else {
		s.consecutiveFailures++
		s.failure = failed(err, s.consecutiveFailures)
		call.token, call.err = s.currentOrFailure()
	}
	s.inFlight = nil
	s.mu.Unlock()
	close(call.done)
}

func (s *tokenSource) currentOrFailure() (string, error) {
	if s.current != "" && now().Before(s.expiresAt.Add(-expiryMargin)) {
		return s.current, nil
	}
	return "", s.failure
}

func renewBeforeExpiry(lifetime time.Duration) time.Duration {
	return min(between(renewBeforeExpiryMin, renewBeforeExpiryMax), lifetime/2)
}

func failed(err error, consecutiveFailures int) *Error {
	failure := &Error{Message: err.Error(), err: err}
	if oidcErr, ok := err.(*Error); ok {
		copied := *oidcErr
		failure = &copied
	}
	if !failure.RetryAt.After(now()) {
		failure.RetryAt = now().Add(backoff(consecutiveFailures))
	}
	return failure
}

func backoff(consecutiveFailures int) time.Duration {
	wait := firstBackoff
	for i := 1; i < consecutiveFailures && wait < maxRetryWait; i++ {
		wait *= 2
	}
	wait = min(wait, maxRetryWait)
	return between(wait*9/10, wait*11/10)
}

func (s *tokenSource) exchange() (grant, error) {
	subjectToken, err := s.options.IdentityToken()
	if err != nil {
		return grant{}, err
	}

	body, err := json.Marshal(map[string]string{
		"grant_type":         grantType,
		"subject_token_type": idTokenType,
		"subject_token":      subjectToken,
		"role_id":            s.options.RoleID,
	})
	if err != nil {
		return grant{}, err
	}

	baseURL := s.options.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return withTimeout("The token exchange", func(ctx context.Context) (grant, error) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/access_tokens/exchange", bytes.NewReader(body))
		if err != nil {
			return grant{}, err
		}
		request.Header.Set("Content-Type", "application/json")

		response, err := httpClient(s.options.HTTPClient).Do(request)
		if err != nil {
			return grant{}, err
		}
		defer func() { _ = response.Body.Close() }()

		var payload struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int64  `json:"expires_in"`
			Error       *struct {
				Message string `json:"message"`
				Code    string `json:"code"`
			} `json:"error"`
		}
		decodeErr := json.NewDecoder(response.Body).Decode(&payload)

		if response.StatusCode >= 400 {
			message := fmt.Sprintf("The token exchange failed with %d.", response.StatusCode)
			code := ""
			if payload.Error != nil {
				if payload.Error.Message != "" {
					message = payload.Error.Message
				}
				code = payload.Error.Code
			}
			return grant{}, &Error{Status: response.StatusCode, Code: code, Message: message, RetryAt: retryHint(response.Header)}
		}
		if decodeErr != nil || payload.AccessToken == "" || payload.ExpiresIn <= 0 {
			return grant{}, &Error{Status: response.StatusCode, Message: "The token exchange returned no access token."}
		}
		return grant{token: payload.AccessToken, lifetime: time.Duration(payload.ExpiresIn) * time.Second}, nil
	})
}

func retryHint(header http.Header) time.Time {
	var wait time.Duration
	if retryAfter := header.Get("Retry-After"); retryAfter != "" {
		if seconds, err := strconv.ParseInt(retryAfter, 10, 64); err == nil {
			wait = time.Duration(min(seconds, int64(maxRetryWait/time.Second))) * time.Second
		} else if date, err := http.ParseTime(retryAfter); err == nil {
			wait = date.Sub(now())
		}
	}
	if wait <= 0 {
		if reset, err := strconv.ParseInt(header.Get("RateLimit-Reset"), 10, 64); err == nil {
			wait = time.Unix(reset, 0).Sub(now())
		}
	}
	return now().Add(min(wait, maxRetryWait))
}

// GitHubActions reads the job's identity token from GitHub Actions. The job needs permissions: id-token: write.
func GitHubActions(options IdentityOptions) func() (string, error) {
	return func() (string, error) {
		requestURL := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
		requestToken := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
		if requestURL == "" || requestToken == "" {
			return "", &Error{Message: "GitHub Actions did not provide an identity token. Give the job permissions: id-token: write."}
		}

		parsed, err := url.Parse(requestURL)
		if err != nil {
			return "", err
		}
		query := parsed.Query()
		query.Set("audience", audience(options))
		parsed.RawQuery = query.Encode()

		return withTimeout("The GitHub Actions identity token request", func(ctx context.Context) (string, error) {
			request, err := http.NewRequestWithContext(ctx, http.MethodGet, parsed.String(), nil)
			if err != nil {
				return "", err
			}
			request.Header.Set("Authorization", "Bearer "+requestToken)

			response, err := httpClient(options.HTTPClient).Do(request)
			if err != nil {
				return "", err
			}
			defer func() { _ = response.Body.Close() }()
			if response.StatusCode >= 400 {
				return "", &Error{Status: response.StatusCode, Message: fmt.Sprintf("GitHub Actions refused the identity token request with %d.", response.StatusCode)}
			}

			var payload struct {
				Value string `json:"value"`
			}
			if err := json.NewDecoder(response.Body).Decode(&payload); err != nil || payload.Value == "" {
				return "", &Error{Status: response.StatusCode, Message: "GitHub Actions returned no identity token."}
			}
			return payload.Value, nil
		})
	}
}

// GoogleCloud reads the identity token of the service account the workload runs as from the metadata
// server, on Cloud Run, Cloud Functions, Compute Engine, or GKE with Workload Identity.
func GoogleCloud(options IdentityOptions) func() (string, error) {
	return func() (string, error) {
		return withTimeout("The identity token request to the Google Cloud metadata server", func(ctx context.Context) (string, error) {
			request, err := http.NewRequestWithContext(ctx, http.MethodGet, googleMetadataIdentityURL+"?audience="+url.QueryEscape(audience(options)), nil)
			if err != nil {
				return "", err
			}
			request.Header.Set("Metadata-Flavor", "Google")

			response, err := httpClient(options.HTTPClient).Do(request)
			if err != nil {
				return "", err
			}
			defer func() { _ = response.Body.Close() }()
			if response.StatusCode >= 400 {
				return "", &Error{Status: response.StatusCode, Message: fmt.Sprintf("The Google Cloud metadata server refused the identity token request with %d.", response.StatusCode)}
			}

			body, err := io.ReadAll(response.Body)
			if err != nil {
				return "", err
			}
			token := strings.TrimSpace(string(body))
			if token == "" {
				return "", &Error{Status: response.StatusCode, Message: "The Google Cloud metadata server returned no identity token."}
			}
			return token, nil
		})
	}
}

func audience(options IdentityOptions) string {
	if options.Audience != "" {
		return options.Audience
	}
	return DefaultAudience
}

func httpClient(client *http.Client) *http.Client {
	if client != nil {
		return client
	}
	return http.DefaultClient
}

func withTimeout[T any](request string, send func(context.Context) (T, error)) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), httpTimeout)
	defer cancel()
	result, err := send(ctx)
	if err != nil && ctx.Err() != nil {
		return result, &Error{Message: fmt.Sprintf("%s timed out after %g seconds.", request, httpTimeout.Seconds())}
	}
	return result, err
}

func between(low, high time.Duration) time.Duration {
	return low + time.Duration(random()*float64(high-low))
}
