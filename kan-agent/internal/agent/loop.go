// Package agent polls the server without executing resource work yet.
package agent

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ciagotech/kan-agent/internal/client"
)

const (
	PollInterval = 5 * time.Second
	BackoffMin   = 1 * time.Second
	BackoffMax   = 30 * time.Second
)

// ErrWorkAvailable stops safely when the server returns work that this slice
// cannot execute. It never acknowledges or cleans up that work.
var ErrWorkAvailable = errors.New("work received but execution is not implemented")

// WorkClient is the read-only server boundary used by the polling loop.
type WorkClient interface {
	GetWork(context.Context) ([]client.Resource, error)
}

type runOptions struct {
	pollInterval time.Duration
	backoffMin   time.Duration
	backoffMax   time.Duration
	now          func() time.Time
	jitter       func(time.Duration) time.Duration
	wait         func(context.Context, time.Duration) error
}

// Run polls one account until the context is cancelled or a non-retryable
// error occurs. It does not execute, acknowledge, or delete resources.
func Run(ctx context.Context, work WorkClient) error {
	return run(ctx, work, runOptions{
		pollInterval: PollInterval,
		backoffMin:   BackoffMin,
		backoffMax:   BackoffMax,
		now:          time.Now,
		jitter:       boundedJitter,
		wait:         wait,
	})
}

func run(ctx context.Context, work WorkClient, opts runOptions) error {
	if work == nil {
		return errors.New("work client is required")
	}
	if opts.now == nil {
		opts.now = time.Now
	}
	if opts.jitter == nil {
		opts.jitter = func(d time.Duration) time.Duration { return d }
	}
	if opts.wait == nil {
		opts.wait = wait
	}

	backoff := opts.backoffMin
	for {
		resources, err := work.GetWork(ctx)
		if err == nil {
			if len(resources) > 0 {
				return fmt.Errorf("%w: %d resource(s)", ErrWorkAvailable, len(resources))
			}
			backoff = opts.backoffMin
			if err := opts.wait(ctx, opts.jitter(opts.pollInterval)); err != nil {
				return cleanWaitError(err)
			}
			continue
		}

		if errors.Is(err, context.Canceled) {
			return nil
		}
		if errors.Is(err, context.DeadlineExceeded) && ctx.Err() != nil {
			return ctx.Err()
		}

		var httpErr *client.HTTPError
		retryable := false
		if errors.As(err, &httpErr) {
			switch {
			case httpErr.StatusCode == http.StatusUnauthorized || httpErr.StatusCode == http.StatusForbidden:
				return fmt.Errorf("agent authentication rejected: %w", err)
			case !temporaryStatus(httpErr.StatusCode):
				return err
			default:
				retryable = true
			}
		}
		if !retryable && !temporaryError(err) {
			return err
		}

		delay := backoff
		if httpErr != nil && httpErr.StatusCode == http.StatusTooManyRequests {
			if retryDelay, ok := parseRetryAfter(httpErr.RetryAfter, opts.now()); ok {
				delay = boundedDelay(retryDelay, opts.backoffMax)
			} else {
				delay = boundedDelay(opts.jitter(delay), opts.backoffMax)
			}
		} else {
			delay = boundedDelay(opts.jitter(delay), opts.backoffMax)
		}
		if err := opts.wait(ctx, delay); err != nil {
			return cleanWaitError(err)
		}
		backoff = nextBackoff(backoff, opts.backoffMax)
	}
}

func temporaryStatus(status int) bool {
	switch status {
	case http.StatusRequestTimeout, http.StatusTooEarly, http.StatusTooManyRequests,
		http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	default:
		return false
	}
}

func temporaryError(err error) bool {
	var networkErr net.Error
	// net.Error.Temporary is deprecated and has no well-defined replacement;
	// timeout is the portable network condition that warrants retrying here.
	return errors.As(err, &networkErr) && networkErr.Timeout()
}

func parseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	if seconds, err := strconv.ParseInt(value, 10, 64); err == nil && seconds >= 0 {
		if seconds > int64(time.Duration(1<<63-1)/time.Second) {
			return time.Duration(1<<63 - 1), true
		}
		return time.Duration(seconds) * time.Second, true
	}
	when, err := http.ParseTime(value)
	if err != nil {
		return 0, false
	}
	if delay := when.Sub(now); delay > 0 {
		return delay, true
	}
	return 0, true
}

func nextBackoff(current, maximum time.Duration) time.Duration {
	if current <= 0 || current >= maximum {
		return maximum
	}
	if current > maximum/2 {
		return maximum
	}
	return current * 2
}

func boundedDelay(delay, maximum time.Duration) time.Duration {
	if delay < 0 {
		return 0
	}
	if maximum > 0 && delay > maximum {
		return maximum
	}
	return delay
}

func boundedJitter(delay time.Duration) time.Duration {
	if delay <= 0 {
		return 0
	}
	spread := delay / 10
	if spread == 0 {
		return delay
	}
	return delay - spread + time.Duration(rand.Int63n(int64(spread*2)+1))
}

func wait(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return nil
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func cleanWaitError(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}
