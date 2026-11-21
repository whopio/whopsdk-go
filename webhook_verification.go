package whopsdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// This file is hand-written and preserved by .fernignore. It restores the
// verification half of the `client.webhooks.unwrap` the Stainless-generated SDK
// shipped: Fern generates from OpenAPI paths and `unwrap` was never a path, so
// no Fern SDK has it. Ruby, Python and TypeScript each regained it as a
// standalone helper; this is Go's.
//
// It is a package-level function rather than a method on a generated client so
// that nothing generated has to be patched. It depends only on the standard
// library — crypto/hmac and crypto/sha256 are all the scheme needs — so it
// survives the client being replaced, and it adds nothing to go.mod.
//
// It deliberately does NOT coerce the parsed body into a typed event model, as
// the Stainless SDK did. Fern generates none: WebhookEvent is the enum of event
// *names* an endpoint subscribes to, not a payload type, so there is nothing to
// coerce into. Callers who want a struct should json.Unmarshal the same body
// once this has returned without error.

// WebhookSignatureTolerance is how far the webhook-timestamp header may sit
// from the current clock, in either direction, before UnwrapWebhook rejects the
// delivery. It bounds how long a captured request stays replayable.
const WebhookSignatureTolerance = 5 * time.Minute

// The signature scheme Whop's backend implements in
// WebhooksManager::SignWebhook: HMAC-SHA256 over "<id>.<timestamp>.<body>",
// base64-encoded, sent as "v1,<signature>". The header may in principle carry a
// space-separated list of versioned entries, so unknown versions are skipped
// rather than treated as a mismatch.
const (
	webhookIDHeader        = "webhook-id"
	webhookSignatureHeader = "webhook-signature"
	webhookTimestampHeader = "webhook-timestamp"
	webhookSignatureScheme = "v1"
)

// Errors UnwrapWebhook returns, wrapped with the specifics of the delivery.
// Match them with errors.Is; the wrapped text is for logs, not for branching.
var (
	// ErrWebhookSecretMissing means no signing secret was supplied, so nothing
	// could be checked. It is a caller mistake, not a rejected delivery.
	ErrWebhookSecretMissing = errors.New("whopsdk: cannot verify a webhook without the endpoint's signing secret")

	// ErrWebhookHeaderMissing means one of webhook-id, webhook-timestamp or
	// webhook-signature was absent or empty.
	ErrWebhookHeaderMissing = errors.New("whopsdk: webhook signature header missing")

	// ErrWebhookTimestamp means webhook-timestamp was unparseable or outside
	// WebhookSignatureTolerance. It is checked before the signature, so a
	// replayed delivery is refused for the timestamp rather than looking valid.
	ErrWebhookTimestamp = errors.New("whopsdk: webhook-timestamp is not within the tolerance window")

	// ErrWebhookSignature means no v1 entry in webhook-signature matched the
	// body. Either the body is not what was signed, or the secret is not the
	// one it was signed with.
	ErrWebhookSignature = errors.New("whopsdk: webhook signature does not match")

	// ErrWebhookBody means the signature verified but the body is not a JSON
	// object. The delivery is authentic; it is not something this can parse.
	ErrWebhookBody = errors.New("whopsdk: webhook body is not a JSON object")
)

// UnwrapWebhook verifies a webhook delivery against its signature headers and
// returns the parsed body.
//
//	payload, err := whopsdk.UnwrapWebhook(body, r.Header, os.Getenv("WHOP_WEBHOOK_SECRET"))
//	if err != nil {
//	    http.Error(w, "invalid signature", http.StatusBadRequest)
//	    return
//	}
//
// body must be the raw, unmodified request bytes. The signature covers exactly
// what was sent, so re-serialising the JSON first — or reading through anything
// that rewrites whitespace — fails verification even for a genuine delivery.
// Read it with io.ReadAll(r.Body) before any decoding.
//
// headers is the request's headers; only webhook-id, webhook-timestamp and
// webhook-signature are read. Whop sends them lowercase, and http.Header
// canonicalises what it stores, so the lookup here is case-insensitive in both
// directions and a hand-built map with lowercase keys works too.
//
// secret is the endpoint's signing secret exactly as Whop issued it — the whole
// ws_-prefixed string. Pass it verbatim: the backend HMACs with the secret's
// literal bytes, prefix included, so stripping the prefix or base64-decoding it
// first derives a different key and every genuine delivery fails. (That decode
// is what the Standard Webhooks libraries do to their key, which is why the
// other SDKs' helpers have to base64-encode the secret before handing it over.
// Computing the HMAC here directly means there is nothing to cancel out.)
//
// The returned map is the decoded body. On any error it is nil and nothing
// about the delivery should be trusted.
func UnwrapWebhook(body []byte, headers http.Header, secret string) (map[string]any, error) {
	if secret == "" {
		return nil, ErrWebhookSecretMissing
	}

	id, err := webhookHeaderValue(headers, webhookIDHeader)
	if err != nil {
		return nil, err
	}
	timestamp, err := webhookHeaderValue(headers, webhookTimestampHeader)
	if err != nil {
		return nil, err
	}
	signatures, err := webhookHeaderValue(headers, webhookSignatureHeader)
	if err != nil {
		return nil, err
	}

	if err := checkWebhookTimestamp(timestamp); err != nil {
		return nil, err
	}
	if err := checkWebhookSignature(secret, id, timestamp, body, signatures); err != nil {
		return nil, err
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrWebhookBody, err)
	}
	return payload, nil
}

func webhookHeaderValue(headers http.Header, name string) (string, error) {
	if value := headers.Get(name); value != "" {
		return value, nil
	}
	for key, values := range headers {
		if !strings.EqualFold(key, name) {
			continue
		}
		for _, value := range values {
			if value != "" {
				return value, nil
			}
		}
	}
	return "", fmt.Errorf("%w: %s", ErrWebhookHeaderMissing, name)
}

func checkWebhookTimestamp(timestamp string) error {
	seconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: %s is not a unix timestamp", ErrWebhookTimestamp, timestamp)
	}
	drift := time.Since(time.Unix(seconds, 0))
	if drift < 0 {
		drift = -drift
	}
	if drift > WebhookSignatureTolerance {
		return fmt.Errorf(
			"%w: %s is %s away from now, tolerance is %s",
			ErrWebhookTimestamp, timestamp, drift.Round(time.Second), WebhookSignatureTolerance,
		)
	}
	return nil
}

func checkWebhookSignature(secret, id, timestamp string, body []byte, signatures string) error {
	expected := signWebhookPayload(secret, id, timestamp, body)
	versioned := 0
	for _, entry := range strings.Fields(signatures) {
		scheme, encoded, found := strings.Cut(entry, ",")
		if !found || scheme != webhookSignatureScheme {
			continue
		}
		versioned++
		candidate, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			continue
		}
		if hmac.Equal(candidate, expected) {
			return nil
		}
	}
	return fmt.Errorf("%w: %d %s entries checked", ErrWebhookSignature, versioned, webhookSignatureScheme)
}

func signWebhookPayload(secret, id, timestamp string, body []byte) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(id))
	mac.Write([]byte("."))
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write(body)
	return mac.Sum(nil)
}
