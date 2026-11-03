package agent

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/ciagotech/kan-agent/internal/client"
)

type workResult struct {
	resources []client.Resource
	err       error
}

type fakeWorkClient struct {
	results []workResult
	calls   int
}

func (f *fakeWorkClient) GetWork(context.Context) ([]client.Resource, error) {
	f.calls++
	if len(f.results) == 0 {
		return nil, errors.New("fake client has no result")
	}
	result := f.results[0]
	f.results = f.results[1:]
	return result.resources, result.err
}

func testOptions(waits *[]time.Duration) runOptions {
	return runOptions{
		pollInterval: time.Second,
		backoffMin:   time.Second,
		backoffMax:   30 * time.Second,
		now:          func() time.Time { return time.Unix(0, 0) },
		jitter:       func(delay time.Duration) time.Duration { return delay },
		wait: func(ctx context.Context, delay time.Duration) error {
			if waits != nil {
				*waits = append(*waits, delay)
			}
			return nil
		},
	}
}

func TestRunPollsNoWorkUntilCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	waits := make([]time.Duration, 0, 2)
	fake := &fakeWorkClient{results: []workResult{{}, {}, {}}}
	opts := testOptions(&waits)
	opts.wait = func(ctx context.Context, delay time.Duration) error {
		waits = append(waits, delay)
		if len(waits) == 2 {
			cancel()
		}
		return ctx.Err()
	}

	if err := run(ctx, fake, opts); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if fake.calls != 2 || len(waits) != 2 {
		t.Fatalf("calls=%d waits=%d, want two of each", fake.calls, len(waits))
	}
}

func TestRunCancellationDuringRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan struct{})
	work := workFunc(func(ctx context.Context) ([]client.Resource, error) {
		close(started)
		<-ctx.Done()
		return nil, ctx.Err()
	})
	go func() {
		<-started
		cancel()
	}()

	if err := run(ctx, work, testOptions(nil)); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
}

func TestRunStopsOnAuthenticationFailure(t *testing.T) {
	waits := make([]time.Duration, 0, 1)
	fake := &fakeWorkClient{results: []workResult{{err: &client.HTTPError{StatusCode: http.StatusUnauthorized}}}}
	err := run(context.Background(), fake, testOptions(&waits))
	var httpErr *client.HTTPError
	if err == nil || !errors.As(err, &httpErr) || httpErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(waits) != 0 {
		t.Fatalf("authentication failure was retried: %v", waits)
	}
}

func TestRunRetriesAndResetsBackoff(t *testing.T) {
	waits := make([]time.Duration, 0, 4)
	fake := &fakeWorkClient{results: []workResult{
		{err: &client.HTTPError{StatusCode: http.StatusServiceUnavailable}},
		{},
		{err: &client.HTTPError{StatusCode: http.StatusServiceUnavailable}},
	}}
	opts := testOptions(&waits)
	opts.wait = func(ctx context.Context, delay time.Duration) error {
		waits = append(waits, delay)
		if len(waits) == 3 {
			return context.Canceled
		}
		return nil
	}
	if err := run(context.Background(), fake, opts); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	want := []time.Duration{time.Second, time.Second, time.Second}
	if len(waits) != len(want) {
		t.Fatalf("waits=%v, want %v", waits, want)
	}
	for i := range want {
		if waits[i] != want[i] {
			t.Fatalf("waits=%v, want %v", waits, want)
		}
	}
}

func TestRunHonorsRetryAfter(t *testing.T) {
	waits := make([]time.Duration, 0, 1)
	fake := &fakeWorkClient{results: []workResult{
		{err: &client.HTTPError{StatusCode: http.StatusTooManyRequests, RetryAfter: "7"}},
	}}
	opts := testOptions(&waits)
	opts.wait = func(ctx context.Context, delay time.Duration) error {
		waits = append(waits, delay)
		return context.Canceled
	}
	if err := run(context.Background(), fake, opts); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if len(waits) != 1 || waits[0] != 7*time.Second {
		t.Fatalf("waits=%v, want [7s]", waits)
	}
}

func TestRunDoesNotExecuteWork(t *testing.T) {
	waits := make([]time.Duration, 0, 1)
	fake := &fakeWorkClient{results: workResultWithResource()}
	err := run(context.Background(), fake, testOptions(&waits))
	if !errors.Is(err, ErrWorkAvailable) {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(waits) != 0 {
		t.Fatalf("work was acknowledged or delayed: %v", waits)
	}
}

func TestRunReturnsPermanentError(t *testing.T) {
	waits := make([]time.Duration, 0, 1)
	fake := &fakeWorkClient{results: []workResult{{err: errors.New("malformed response")}}}
	err := run(context.Background(), fake, testOptions(&waits))
	if err == nil || err.Error() != "malformed response" {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(waits) != 0 {
		t.Fatalf("permanent error was retried: %v", waits)
	}
}

type workFunc func(context.Context) ([]client.Resource, error)

func (f workFunc) GetWork(ctx context.Context) ([]client.Resource, error) { return f(ctx) }

func workResultWithResource() []workResult {
	return []workResult{{resources: []client.Resource{{ID: "resource-1"}}}}
}
