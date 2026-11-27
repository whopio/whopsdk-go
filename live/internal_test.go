package live

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBacksOffFromUpToASecondToItsCapWithUpToAQuarterOfJitter(t *testing.T) {
	t.Parallel()
	var waits []time.Duration
	for attempt := 0; attempt < 8; attempt++ {
		waits = append(waits, backoff(attempt, 30*time.Second, 0))
	}
	second := time.Second
	require.Equal(t, []time.Duration{0, second, 2 * second, 4 * second, 8 * second, 16 * second, 30 * second, 30 * second}, waits)
	require.Equal(t, 250*time.Millisecond, backoff(0, 30*time.Second, 0.25))
	require.Equal(t, second, backoff(0, 30*time.Second, 1))
	require.Equal(t, 500*time.Millisecond, backoff(0, 500*time.Millisecond, 1))
	require.Equal(t, 5*second, backoff(3, 30*time.Second, 1))
	require.Equal(t, 30*second, backoff(1000, 30*time.Second, 0))
}

func TestParsesEventsAcrossLinesSkippingCommentsAndUnknownFields(t *testing.T) {
	t.Parallel()
	var events parser
	var parsed [][2]string
	for _, line := range []string{": keep-alive", "", "id: 7", "event: stream.snapshot", "data: {\"a\":", "data:1}", "retry: 5", "", "data", ""} {
		if name, data, ok := events.feed(line); ok {
			parsed = append(parsed, [2]string{name, data})
		}
	}
	require.Equal(t, [][2]string{{SnapshotEvent, "{\"a\":\n1}"}, {"message", ""}}, parsed)
}

func TestEndsReconnectsOrFailsOnACloseByItsReasonAndStatus(t *testing.T) {
	t.Parallel()
	require.NoError(t, closedBy(`{"reason":"not_live","status":null}`))
	require.NoError(t, closedBy(`{"reason":"not_live","status":204}`))
	require.NoError(t, closedBy(`{"reason":"not_live"}`))

	var closed *StreamClosedError
	require.ErrorAs(t, closedBy(`{"reason":"not_live","status":403}`), &closed)
	require.Equal(t, StreamClosedError{Reason: "not_live", Status: 403}, *closed)
	require.ErrorAs(t, closedBy(`{"reason":"unreadable","status":null}`), &closed)
	require.Equal(t, StreamClosedError{Reason: "unreadable"}, *closed)

	for _, data := range []string{`{"reason":"reconnect","status":null}`, `{"reason":"moved"}`, `[]`, `not json`} {
		require.ErrorIs(t, closedBy(data), errReconnect, data)
	}
}

func TestDecodesASnapshotWithoutItsEventNameAndRefusesOneThatIsNotAnObject(t *testing.T) {
	t.Parallel()
	type answer struct {
		A     int    `json:"a"`
		Event string `json:"event"`
	}
	decoded, err := decode[*answer](`{"a":1,"event":"stream.snapshot"}`)
	require.NoError(t, err)
	require.Equal(t, &answer{A: 1}, decoded)

	for _, data := range []string{`null`, `[]`, `1`} {
		decoded, err := decode[*answer](data)
		require.ErrorContains(t, err, "is not a JSON object", data)
		require.Nil(t, decoded, data)
	}
}

func TestReconnectsOnlyAfterAStatusAReconnectMayAnswerDifferently(t *testing.T) {
	t.Parallel()
	for status, reconnects := range map[int]bool{200: false, 204: false, 304: true, 400: false, 401: false, 404: false, 408: true, 410: false, 422: false, 429: true, 500: true, 502: true, 503: true} {
		require.Equal(t, reconnects, reconnectsAfter(status), status)
	}
}
