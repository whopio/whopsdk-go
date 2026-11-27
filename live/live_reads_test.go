package live_test

import (
	"testing"

	client "github.com/whopio/whopsdk-go/v2/client"
	timeseries "github.com/whopio/whopsdk-go/v2/stats/timeseries"
)

// Each live read's Watch and WaitFor take exactly its read's parameters and answer its answer type:
// checkLiveRead binds its type parameters to the read's, so a method that differs does not compile.
func TestEveryLiveReadStreamsItsRead(t *testing.T) {
	t.Parallel()
	t.Run("stats.timeSeries.retrieve", func(t *testing.T) {
		t.Parallel()
		checkLiveRead(
			t,
			func(whop *client.Whop) *timeseries.Client { return whop.Stats.TimeSeries },
			(*timeseries.Client).Retrieve,
			(*timeseries.Client).Watch,
			(*timeseries.Client).WaitFor,
		)
	})
}
