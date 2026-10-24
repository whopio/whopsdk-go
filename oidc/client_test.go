package oidc_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	whopsdk "github.com/whopio/whopsdk-go/v2"
	client "github.com/whopio/whopsdk-go/v2/client"
	"github.com/whopio/whopsdk-go/v2/oidc"
	option "github.com/whopio/whopsdk-go/v2/option"
)

func newWhop(t *testing.T, serverURL string) *client.Whop {
	t.Helper()
	httpClient, err := oidc.NewHTTPClient(oidc.Options{
		RoleID:        "aurl_xxxxxxxxxxxxx",
		IdentityToken: func() (string, error) { return "id-token", nil },
		BaseURL:       serverURL,
	})
	if err != nil {
		t.Fatal(err)
	}
	return client.NewWhop(option.WithBaseURL(serverURL), option.WithHTTPClient(httpClient))
}

func TestTheClientAuthenticatesWithTheExchangedToken(t *testing.T) {
	var gotAuthorization atomic.Value
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/access_tokens/exchange" {
			_, _ = w.Write([]byte(`{"access_token":"whop-token-1","expires_in":3600}`))
			return
		}
		gotAuthorization.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"user_xxxxxxxxxxxxx"}`))
	}))
	t.Cleanup(server.Close)

	_, err := newWhop(t, server.URL).Users.Retrieve(context.Background(), &whopsdk.RetrieveUsersRequest{ID: "user_xxxxxxxxxxxxx"})

	if err != nil || gotAuthorization.Load() != "Bearer whop-token-1" {
		t.Fatalf("Authorization = %v, err = %v", gotAuthorization.Load(), err)
	}
}

func TestTheClientReturnsTheExchangeErrorWithoutSendingTheRequest(t *testing.T) {
	var apiRequests int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/access_tokens/exchange" {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"error":{"type":"forbidden","message":"This network is not allowed.","code":"oidc_exchange_ip_not_allowed"}}`))
			return
		}
		atomic.AddInt32(&apiRequests, 1)
		w.WriteHeader(http.StatusUnauthorized)
	}))
	t.Cleanup(server.Close)

	_, err := newWhop(t, server.URL).Users.Retrieve(context.Background(), &whopsdk.RetrieveUsersRequest{ID: "user_xxxxxxxxxxxxx"})

	var oidcErr *oidc.Error
	if !errors.As(err, &oidcErr) || oidcErr.Status != http.StatusForbidden || oidcErr.Code != "oidc_exchange_ip_not_allowed" {
		t.Fatalf("err = %#v", err)
	}
	if got := atomic.LoadInt32(&apiRequests); got != 0 {
		t.Fatalf("unauthenticated API requests = %d", got)
	}
}
