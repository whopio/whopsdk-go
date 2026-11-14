package whopsdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Every fixture below is signed the way WebhooksManager::SignWebhook signs, not
// the way UnwrapWebhook verifies: HMAC-SHA256 over the single string
// "<id>.<timestamp>.<body>", base64-encoded, sent as "v1,<signature>". Building
// it with the code under test would only prove that function agrees with
// itself, which is exactly the failure mode that broke the other SDKs' helpers.
//
// The secret is a real-shaped one: "ws_" plus 64 hex characters, 67 bytes, used
// as the raw HMAC key with the prefix included.
const (
	webhookTestSecret      = "ws_1f0c3a5e7b9d2f4a6c8e0b1d3f5a7c9e1b3d5f7a9c1e3b5d7f9a1c3e5b7d9f1a"
	webhookTestWrongSecret = "ws_9a1c3e5b7d9f1a3c5e7b9d1f3a5c7e9b1d3f5a7c9e1b3d5f7a9c1e3b5d7f9a1c"
	webhookTestID          = "msg_2Qb7XtLm9Kd4Rp1Ns6Vw8Zc0"
	webhookTestBody        = `{"id":"msg_2Qb7XtLm9Kd4Rp1Ns6Vw8Zc0","type":"product.created","api_version":"2026-08-01","data":{"id":"prod_A1b2C3d4","title":"Detailing Masterclass"}}`
)

func webhookTestSignature(t *testing.T, secret, id, timestamp, body string) string {
	t.Helper()
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(id + "." + timestamp + "." + body))
	require.NoError(t, err)
	return "v1," + base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func webhookTestTimestamp(offset time.Duration) string {
	return strconv.FormatInt(time.Now().Add(offset).Unix(), 10)
}

// Whop sends the three headers lowercase. http.Header{} built as a literal
// keeps that casing verbatim, which is what arrives on the wire; the canonical
// spelling is exercised separately in TestUnwrapWebhookHeaderCasing.
func webhookTestHeaders(t *testing.T, secret, id, timestamp, body string) http.Header {
	t.Helper()
	return http.Header{
		"webhook-id":        {id},
		"webhook-timestamp": {timestamp},
		"webhook-signature": {webhookTestSignature(t, secret, id, timestamp, body)},
		"content-type":      {"application/json"},
	}
}

func webhookTestDelivery(t *testing.T) (string, http.Header) {
	t.Helper()
	timestamp := webhookTestTimestamp(0)
	return webhookTestBody, webhookTestHeaders(t, webhookTestSecret, webhookTestID, timestamp, webhookTestBody)
}

func TestUnwrapWebhookAcceptsABackendSignedDelivery(t *testing.T) {
	body, headers := webhookTestDelivery(t)

	payload, err := UnwrapWebhook([]byte(body), headers, webhookTestSecret)

	require.NoError(t, err)
	assert.Equal(t, webhookTestID, payload["id"])
	assert.Equal(t, "product.created", payload["type"])
	data, ok := payload["data"].(map[string]any)
	require.True(t, ok, "expected the envelope's data object")
	assert.Equal(t, "prod_A1b2C3d4", data["id"])
}

// The signature covers the exact bytes sent. json.Marshal of the decoded body
// is semantically identical and byte-different, and it must not verify —
// otherwise the helper is checking something other than what was signed.
func TestUnwrapWebhookRejectsAReserializedBody(t *testing.T) {
	body, headers := webhookTestDelivery(t)
	var decoded map[string]any
	require.NoError(t, json.Unmarshal([]byte(body), &decoded))
	reserialized, err := json.Marshal(decoded)
	require.NoError(t, err)
	require.NotEqual(t, body, string(reserialized), "the fixture must not survive a round-trip unchanged")

	payload, err := UnwrapWebhook(reserialized, headers, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookSignature)
}

// One byte changed inside a JSON string value: still valid JSON, so a rejection
// means the signature was checked rather than the parser tripping.
func TestUnwrapWebhookRejectsATamperedBody(t *testing.T) {
	body, headers := webhookTestDelivery(t)
	tampered := strings.Replace(body, "prod_A1b2C3d4", "prod_A1b2C3d5", 1)
	require.NotEqual(t, body, tampered)
	require.NoError(t, json.Unmarshal([]byte(tampered), &map[string]any{}))

	payload, err := UnwrapWebhook([]byte(tampered), headers, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookSignature)
}

func TestUnwrapWebhookRejectsTheWrongSecret(t *testing.T) {
	body, headers := webhookTestDelivery(t)

	payload, err := UnwrapWebhook([]byte(body), headers, webhookTestWrongSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookSignature)
}

// The backend HMACs with the secret's literal bytes, prefix included. The two
// key derivations that would silently replace that — base64-encoding the secret
// first, which is what the other SDKs' helpers must do to cancel out the
// Standard Webhooks libraries' decode, and stripping the ws_ prefix the way
// those libraries strip whsec_ — each produce a signature this must refuse.
func TestUnwrapWebhookUsesTheSecretsLiteralBytesAsTheKey(t *testing.T) {
	timestamp := webhookTestTimestamp(0)

	for name, key := range map[string]string{
		"base64-encoded secret": base64.StdEncoding.EncodeToString([]byte(webhookTestSecret)),
		"secret without ws_":    strings.TrimPrefix(webhookTestSecret, "ws_"),
	} {
		t.Run(name, func(t *testing.T) {
			headers := webhookTestHeaders(t, key, webhookTestID, timestamp, webhookTestBody)

			payload, err := UnwrapWebhook([]byte(webhookTestBody), headers, webhookTestSecret)

			assert.Nil(t, payload)
			assert.ErrorIs(t, err, ErrWebhookSignature)
		})
	}
}

func TestUnwrapWebhookRejectsAMissingSecret(t *testing.T) {
	body, headers := webhookTestDelivery(t)

	payload, err := UnwrapWebhook([]byte(body), headers, "")

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookSecretMissing)
}

func TestUnwrapWebhookRejectsAMissingHeader(t *testing.T) {
	for _, name := range []string{"webhook-id", "webhook-timestamp", "webhook-signature"} {
		t.Run(name, func(t *testing.T) {
			body, headers := webhookTestDelivery(t)
			delete(headers, name)

			payload, err := UnwrapWebhook([]byte(body), headers, webhookTestSecret)

			assert.Nil(t, payload)
			assert.ErrorIs(t, err, ErrWebhookHeaderMissing)
			assert.Contains(t, err.Error(), name)
		})
	}
}

func TestUnwrapWebhookRejectsAnEmptyHeader(t *testing.T) {
	body, headers := webhookTestDelivery(t)
	headers["webhook-id"] = []string{""}

	payload, err := UnwrapWebhook([]byte(body), headers, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookHeaderMissing)
}

// Back-dating the header also invalidates the signature, so the tolerance check
// has to run first for the rejection to name the timestamp. A caller that logs
// "bad signature" for an hour-old replay is debugging the wrong thing.
func TestUnwrapWebhookRejectsATimestampOutsideTolerance(t *testing.T) {
	for name, offset := range map[string]time.Duration{
		"stale":  -time.Hour,
		"future": time.Hour,
	} {
		t.Run(name, func(t *testing.T) {
			timestamp := webhookTestTimestamp(offset)
			headers := webhookTestHeaders(t, webhookTestSecret, webhookTestID, timestamp, webhookTestBody)

			payload, err := UnwrapWebhook([]byte(webhookTestBody), headers, webhookTestSecret)

			assert.Nil(t, payload)
			assert.ErrorIs(t, err, ErrWebhookTimestamp)
			assert.Contains(t, err.Error(), "timestamp")
		})
	}
}

func TestUnwrapWebhookAcceptsATimestampInsideTolerance(t *testing.T) {
	timestamp := webhookTestTimestamp(-4 * time.Minute)
	headers := webhookTestHeaders(t, webhookTestSecret, webhookTestID, timestamp, webhookTestBody)

	payload, err := UnwrapWebhook([]byte(webhookTestBody), headers, webhookTestSecret)

	require.NoError(t, err)
	assert.Equal(t, webhookTestID, payload["id"])
}

func TestUnwrapWebhookRejectsANonNumericTimestamp(t *testing.T) {
	body, headers := webhookTestDelivery(t)
	headers["webhook-timestamp"] = []string{"not-a-timestamp"}

	payload, err := UnwrapWebhook([]byte(body), headers, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookTimestamp)
}

// The header is a space-separated list. Unknown versions are skipped rather
// than failing the delivery, and the valid v1 entry is found wherever it sits.
func TestUnwrapWebhookAcceptsAMultiEntrySignatureHeader(t *testing.T) {
	timestamp := webhookTestTimestamp(0)
	valid := webhookTestSignature(t, webhookTestSecret, webhookTestID, timestamp, webhookTestBody)
	stale := webhookTestSignature(t, webhookTestWrongSecret, webhookTestID, timestamp, webhookTestBody)

	for name, header := range map[string]string{
		"valid last":          stale + " " + valid,
		"valid between":       stale + " " + valid + " " + stale,
		"after an unknown v2": "v2,cGxhY2Vob2xkZXI= " + valid,
		"after a malformed":   "v1,!!!not-base64!!! " + valid,
	} {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{
				"webhook-id":        {webhookTestID},
				"webhook-timestamp": {timestamp},
				"webhook-signature": {header},
			}

			payload, err := UnwrapWebhook([]byte(webhookTestBody), headers, webhookTestSecret)

			require.NoError(t, err)
			assert.Equal(t, webhookTestID, payload["id"])
		})
	}
}

func TestUnwrapWebhookRejectsASignatureHeaderWithNoUsableEntry(t *testing.T) {
	timestamp := webhookTestTimestamp(0)
	headers := http.Header{
		"webhook-id":        {webhookTestID},
		"webhook-timestamp": {timestamp},
		"webhook-signature": {"v2,cGxhY2Vob2xkZXI="},
	}

	payload, err := UnwrapWebhook([]byte(webhookTestBody), headers, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookSignature)
}

// Whop sends the headers lowercase, net/http canonicalises what it stores, and
// a hand-built map keeps whatever the caller typed. All three must work.
func TestUnwrapWebhookHeaderCasing(t *testing.T) {
	timestamp := webhookTestTimestamp(0)
	signature := webhookTestSignature(t, webhookTestSecret, webhookTestID, timestamp, webhookTestBody)

	canonical := http.Header{}
	canonical.Set("Webhook-Id", webhookTestID)
	canonical.Set("Webhook-Timestamp", timestamp)
	canonical.Set("Webhook-Signature", signature)

	mixed := http.Header{
		"WEBHOOK-ID":        {webhookTestID},
		"Webhook-TimeStamp": {timestamp},
		"webhook-SIGNATURE": {signature},
	}

	for name, headers := range map[string]http.Header{"canonical": canonical, "mixed": mixed} {
		t.Run(name, func(t *testing.T) {
			payload, err := UnwrapWebhook([]byte(webhookTestBody), headers, webhookTestSecret)

			require.NoError(t, err)
			assert.Equal(t, webhookTestID, payload["id"])
		})
	}
}

func TestUnwrapWebhookRejectsNilHeaders(t *testing.T) {
	payload, err := UnwrapWebhook([]byte(webhookTestBody), nil, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookHeaderMissing)
}

// A body that verifies but is not a JSON object is authentic and unparseable —
// a distinct outcome from a signature mismatch, so callers can tell them apart.
func TestUnwrapWebhookSeparatesAnUnparseableBodyFromABadSignature(t *testing.T) {
	body := "not json"
	timestamp := webhookTestTimestamp(0)
	headers := webhookTestHeaders(t, webhookTestSecret, webhookTestID, timestamp, body)

	payload, err := UnwrapWebhook([]byte(body), headers, webhookTestSecret)

	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrWebhookBody)
	assert.NotErrorIs(t, err, ErrWebhookSignature)
}
