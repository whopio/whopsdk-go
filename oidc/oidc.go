// Package oidc authenticates the Whop client with a cloud provider's OIDC identity token instead
// of an API key.
//
// Token returns a function for option.WithTokenFunc. The client calls it before each request; it
// exchanges an identity token for a short-lived Whop access token, reuses that token until shortly
// before it expires, and then exchanges a fresh identity token. The package uses only the standard
// library, never generated client code, so it survives the client being regenerated.
//
//	whop := client.NewWhop(option.WithTokenFunc(oidc.Token(oidc.Options{
//		RoleID:        "aurl_xxxxxxxxxxxxx",
//		IdentityToken: oidc.GitHubActions(oidc.IdentityOptions{}),
//	})))
package oidc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	renewBeforeExpiry         = time.Minute
	httpTimeout               = 10 * time.Second
	googleMetadataIdentityURL = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity"
)

var now = time.Now

// Options configures Token.
type Options struct {
	// RoleID is the role to assume, prefixed aurl_. A trust linked to it must accept the identity token.
	RoleID string
	// IdentityToken returns an identity token from your cloud provider. It is called again for every
	// exchange, so return a fresh token rather than a stored one.
	IdentityToken func() (string, error)
	// BaseURL is the API's base URL, for example https://sandbox-api.whop.com/api/v1. Defaults to production.
	BaseURL string
	// HTTPClient sends the exchange. Defaults to a client with a 10 second timeout.
	HTTPClient *http.Client
}

// IdentityOptions configures the built-in identity token sources.
type IdentityOptions struct {
	// Audience is the audience to request. It must equal the audience on your trust. Defaults to DefaultAudience.
	Audience string
	// HTTPClient sends the request. Defaults to a client with a 10 second timeout.
	HTTPClient *http.Client
}

// Error reports a failed exchange or identity token request. Code is the API's error code when it sent one.
type Error struct {
	Status  int
	Code    string
	Message string
}

func (e *Error) Error() string { return e.Message }

// Token returns a function for option.WithTokenFunc that exchanges your provider's identity token for a
// Whop access token and renews it before it expires. Concurrent requests share one exchange.
func Token(options Options) func() (string, error) {
	source := &tokenSource{options: options}
	return source.token
}

type tokenSource struct {
	options Options
	mu      sync.Mutex
	current string
	renewAt time.Time
}

func (s *tokenSource) token() (string, error) {
	if s.options.RoleID == "" {
		return "", &Error{Message: "Pass the role to assume as RoleID."}
	}
	if s.options.IdentityToken == nil {
		return "", &Error{Message: "Pass a source of identity tokens as IdentityToken."}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.current != "" && now().Before(s.renewAt) {
		return s.current, nil
	}

	token, expiresIn, err := s.exchange()
	if err != nil {
		return "", err
	}
	s.current = token
	s.renewAt = now().Add(time.Duration(expiresIn)*time.Second - renewBeforeExpiry)
	return token, nil
}

func (s *tokenSource) exchange() (string, int64, error) {
	subjectToken, err := s.options.IdentityToken()
	if err != nil {
		return "", 0, err
	}

	body, err := json.Marshal(map[string]string{
		"grant_type":         grantType,
		"subject_token_type": idTokenType,
		"subject_token":      subjectToken,
		"role_id":            s.options.RoleID,
	})
	if err != nil {
		return "", 0, err
	}

	baseURL := s.options.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	request, err := http.NewRequest(http.MethodPost, strings.TrimRight(baseURL, "/")+"/access_tokens/exchange", bytes.NewReader(body))
	if err != nil {
		return "", 0, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient(s.options.HTTPClient).Do(request)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()

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
		return "", 0, &Error{Status: response.StatusCode, Code: code, Message: message}
	}
	if decodeErr != nil || payload.AccessToken == "" || payload.ExpiresIn <= 0 {
		return "", 0, &Error{Status: response.StatusCode, Message: "The token exchange returned no access token."}
	}
	return payload.AccessToken, payload.ExpiresIn, nil
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

		request, err := http.NewRequest(http.MethodGet, parsed.String(), nil)
		if err != nil {
			return "", err
		}
		request.Header.Set("Authorization", "Bearer "+requestToken)

		response, err := httpClient(options.HTTPClient).Do(request)
		if err != nil {
			return "", err
		}
		defer response.Body.Close()
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
	}
}

// GoogleCloud reads the identity token of the service account the workload runs as from the metadata
// server, on Cloud Run, Cloud Functions, Compute Engine, or GKE with Workload Identity.
func GoogleCloud(options IdentityOptions) func() (string, error) {
	return func() (string, error) {
		request, err := http.NewRequest(http.MethodGet, googleMetadataIdentityURL+"?audience="+url.QueryEscape(audience(options)), nil)
		if err != nil {
			return "", err
		}
		request.Header.Set("Metadata-Flavor", "Google")

		response, err := httpClient(options.HTTPClient).Do(request)
		if err != nil {
			return "", err
		}
		defer response.Body.Close()
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
	return &http.Client{Timeout: httpTimeout}
}
