package live

import (
	"context"
	"time"
)

// WithClock, WithRandom, WithPause and WithIdleTimeout let the tests in package live_test drive the
// stream's time, jitter, waits and idle watchdog.

func WithClock(now func() time.Time) WatchOption {
	return watchOption(func(s *settings) { s.now = now })
}

func WithRandom(random func() float64) WatchOption {
	return watchOption(func(s *settings) { s.random = random })
}

func WithPause(pause func(ctx context.Context, wait time.Duration) bool) WatchOption {
	return watchOption(func(s *settings) { s.pause = pause })
}

func WithIdleTimeout(idleTimeout time.Duration) WatchOption {
	return watchOption(func(s *settings) { s.idleTimeout = idleTimeout })
}
