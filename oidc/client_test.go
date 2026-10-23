package oidc_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	whopsdk "github.com/whopio/whopsdk-go/v2"
	client "github.com/whopio/whopsdk-go/v2/client"
	"github.com/whopio/whopsdk-go/v2/oidc"
	option "github.com/whopio/whopsdk-go/v2/option"
)

func TestTheClientAuthenticatesWithTheExchangedToken(t *testing.T) {
	var gotAuthorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/access_tokens/exchange" {
			_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
			return
		}
		gotAuthorization = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"user_xxxxxxxxxxxxx"}`))
	}))
	t.Cleanup(server.Close)

	whop := client.NewWhop(
		option.WithBaseURL(server.URL),
		option.WithTokenFunc(oidc.Token(oidc.Options{
			RoleID:        "aurl_xxxxxxxxxxxxx",
			IdentityToken: func() (string, error) { return "id-token", nil },
			BaseURL:       server.URL,
		})),
	)

	_, _ = whop.Users.Retrieve(context.Background(), &whopsdk.RetrieveUsersRequest{ID: "user_xxxxxxxxxxxxx"})

	if gotAuthorization != "Bearer whop-token-1" {
		t.Fatalf("Authorization = %q", gotAuthorization)
	}
}
