// Package client sends signed requests, without polling or executing workloads.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ciagotech/kan-agent/internal/config"
	"github.com/ciagotech/kan-agent/internal/lifecycle"
	"github.com/ciagotech/kan-agent/sigv1"
)

const (
	RequestTimeout = 10 * time.Second
	// MaxResponseBytes is a local safety bound, not a wire-contract limit.
	MaxResponseBytes = 4 << 20
	maxWorkItems     = 1000
)

// Resource contains common contract fields. Drivers must separately validate
// kind-specific Spec and installed capabilities. Labels are not authorization.
type Resource struct {
	ID           string            `json:"id"`
	Kind         string            `json:"kind"`
	Generation   int64             `json:"generation"`
	DesiredState string            `json:"desiredState"`
	State        string            `json:"state,omitempty"`
	Deleting     bool              `json:"deleting"`
	Spec         json.RawMessage   `json:"spec"`
	Labels       map[string]string `json:"labels"`
}

// HTTPError retains status and retry guidance, never the possibly sensitive body.
// The future polling loop decides whether and when a request may be retried.
type HTTPError struct {
	StatusCode int
	RetryAfter string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("work request returned HTTP %d", e.StatusCode)
}

// Client holds one account's private signing context and reusable HTTP client.
type Client struct {
	workURL   string
	accessKey string
	secretKey string
	rules     sigv1.Rules
	http      *http.Client
}

func (*Client) String() string     { return "signed API client (credentials redacted)" }
func (c *Client) GoString() string { return c.String() }

// New normally receives config.Load's validated settings. It performs no I/O.
func New(cfg config.Config) (*Client, error) {
	u, err := url.Parse(cfg.ServerURL)
	if err != nil || u.Hostname() == "" || u.User != nil || u.Opaque != "" ||
		(u.Path != "" && u.Path != "/") || u.RawPath != "" || u.RawQuery != "" ||
		u.ForceQuery || strings.Contains(cfg.ServerURL, "#") || strings.HasSuffix(u.Host, ":") {
		return nil, errors.New("server URL must be an HTTP(S) origin")
	}
	if port := u.Port(); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil || n < 1 || n > 65535 {
			return nil, errors.New("invalid server port")
		}
	}
	ip := net.ParseIP(u.Hostname())
	local := strings.EqualFold(u.Hostname(), "localhost") || (ip != nil && ip.IsLoopback())
	if u.Scheme != "https" && !(u.Scheme == "http" && local) {
		return nil, errors.New("server requires HTTPS outside localhost")
	}
	if cfg.AccessKey == "" || cfg.SecretKey == "" ||
		strings.ContainsAny(cfg.AccessKey+cfg.SecretKey, " \t\r\n") || cfg.Rules == nil {
		return nil, errors.New("validated credentials and signing rules are required")
	}
	rules := *cfg.Rules
	rules.AlwaysSignedHeaders = append([]string(nil), cfg.Rules.AlwaysSignedHeaders...)
	return &Client{
		workURL:   strings.TrimRight(u.String(), "/") + "/internal/v1/agent/work",
		accessKey: cfg.AccessKey, secretKey: cfg.SecretKey, rules: rules,
		http: &http.Client{Timeout: RequestTimeout, CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		}},
	}, nil
}

// GetWork performs one read-only signed GET. A non-nil empty slice represents
// 200 []; all failures return nil plus an error. No retries, claims or acks occur.
func (c *Client) GetWork(ctx context.Context) ([]Resource, error) {
	req, err := sigv1.NewSignedRequest(ctx, http.MethodGet, c.workURL,
		c.accessKey, c.secretKey, nil, "", &c.rules)
	if err != nil {
		return nil, errors.New("could not construct signed work request")
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("work request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{StatusCode: resp.StatusCode, RetryAfter: resp.Header.Get("Retry-After")}
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, MaxResponseBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read work response: %w", err)
	}
	if len(body) > MaxResponseBytes {
		return nil, errors.New("work response exceeds local size limit")
	}
	var items []json.RawMessage
	if err := json.Unmarshal(body, &items); err != nil || items == nil || len(items) > maxWorkItems {
		return nil, errors.New("work response must be a JSON array of at most 1000 resources")
	}
	resources := make([]Resource, 0, len(items))
	for i, item := range items {
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(item, &fields); err != nil || fields == nil {
			return nil, fmt.Errorf("work item %d must be an object", i)
		}
		for _, name := range []string{"id", "kind", "generation", "desiredState", "deleting", "spec", "labels"} {
			value, ok := fields[name]
			if !ok || bytes.Equal(bytes.TrimSpace(value), []byte("null")) {
				return nil, fmt.Errorf("work item %d is missing required fields", i)
			}
		}
		var resource Resource
		if err := json.Unmarshal(item, &resource); err != nil || resource.ID == "" ||
			len([]rune(resource.ID)) > 128 || resource.Generation < 1 || resource.Labels == nil ||
			!bytes.HasPrefix(bytes.TrimSpace(resource.Spec), []byte("{")) {
			return nil, fmt.Errorf("work item %d has invalid common fields", i)
		}
		if lifecycle.ValidateResource(resource.Kind, resource.DesiredState, resource.Deleting, true) != nil ||
			(resource.State != "" && !lifecycle.CanTransition(resource.Kind, resource.State, resource.State)) {
			return nil, fmt.Errorf("work item %d violates lifecycle rules", i)
		}
		resources = append(resources, resource)
	}
	return resources, nil
}
