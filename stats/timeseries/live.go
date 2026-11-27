package timeseries

import (
	context "context"

	live "github.com/whopio/whopsdk-go/v2/live"
	option "github.com/whopio/whopsdk-go/v2/option"
	stats "github.com/whopio/whopsdk-go/v2/stats"
)

// Watch streams Retrieve's answer now and each time it changes, until the stream
// ends; see package live.
func (c *Client) Watch(
	ctx context.Context,
	request *stats.RetrieveTimeSeriesRequest,
	opts ...live.WatchOption,
) *live.Stream[*stats.RetrieveTimeSeriesResponse] {
	return live.Watch(ctx, c.options.HTTPClient, func(connection context.Context, requestOptions ...option.RequestOption) (*stats.RetrieveTimeSeriesResponse, error) {
		return c.Retrieve(connection, request, requestOptions...)
	}, opts...)
}

// WaitFor returns the first of Retrieve's answers that until accepts, the current
// one included; see package live.
func (c *Client) WaitFor(
	ctx context.Context,
	request *stats.RetrieveTimeSeriesRequest,
	until func(*stats.RetrieveTimeSeriesResponse) bool,
	opts ...live.WaitForOption,
) (*stats.RetrieveTimeSeriesResponse, error) {
	return live.WaitFor(ctx, c.options.HTTPClient, func(connection context.Context, requestOptions ...option.RequestOption) (*stats.RetrieveTimeSeriesResponse, error) {
		return c.Retrieve(connection, request, requestOptions...)
	}, until, opts...)
}
