package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ciagotech/kan-agent/internal/config"
	"github.com/ciagotech/kan-agent/sigv1"
)

func testClient(t *testing.T, endpoint string) *Client {
	t.Helper()
	rules, err := sigv1.LoadRules("../../../contracts/sigv1-rules.json")
	if err != nil {
		t.Fatal(err)
	}
	c, err := New(config.Config{ServerURL: endpoint, AccessKey: "synthetic-key", SecretKey: "synthetic-secret", Rules: rules})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestSignedNoWorkAndFreshNonce(t *testing.T) {
	var previous string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.RequestURI() != "/internal/v1/agent/work" {
			t.Error("incorrect request")
		}
		headers := map[string]string{"host": r.Host}
		for name, values := range r.Header {
			headers[name] = values[0]
		}
		rules, err := sigv1.LoadRules("../../../contracts/sigv1-rules.json")
		if err != nil {
			t.Error(err)
			w.WriteHeader(500)
			return
		}
		signature, _, _, err := sigv1.Sign("synthetic-secret", r.Method, r.URL.EscapedPath(), r.URL.RawQuery, headers, nil, r.Header.Get(rules.DateHeader), rules)
		if err != nil || signature != r.Header.Get(rules.SignatureHeader) {
			t.Error("invalid signature")
		}
		nonce := r.Header.Get(rules.NonceHeader)
		if nonce == "" || nonce == previous {
			t.Error("nonce not fresh")
		}
		previous = nonce
		fmt.Fprint(w, "[]")
	}))
	defer server.Close()
	c := testClient(t, server.URL)
	for i := 0; i < 2; i++ {
		work, err := c.GetWork(context.Background())
		if err != nil || work == nil || len(work) != 0 {
			t.Fatalf("work=%v err=%v", work, err)
		}
	}
	if strings.Contains(fmt.Sprintf("%+v %#v", c, c), "synthetic-secret") {
		t.Fatal("secret leaked")
	}
}

func TestWorkResponseValidation(t *testing.T) {
	valid := `{"id":"r1","kind":"compute.instance","generation":1,"desiredState":"active","deleting":false,"spec":{},"labels":{}}`
	cases := []struct {
		name, body string
		wantError  bool
	}{
		{"resource", "[" + valid + "]", false},
		{"null", "null", true}, {"object", "{}", true}, {"empty body", "", true},
		{"trailing JSON", "[] []", true}, {"null item", "[null]", true},
		{"missing deleting", "[" + strings.Replace(valid, `"deleting":false,`, "", 1) + "]", true},
		{"invalid generation", "[" + strings.Replace(valid, `"generation":1`, `"generation":0`, 1) + "]", true},
		{"server owned kind", "[" + strings.Replace(valid, "compute.instance", "storage.bucket", 1) + "]", true},
		{"invalid spec", "[" + strings.Replace(valid, `"spec":{}`, `"spec":[]`, 1) + "]", true},
		{"oversized", strings.Repeat(" ", MaxResponseBytes+1), true},
		{"too many", "[" + strings.Repeat(valid+",", 1000) + valid + "]", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, tc.body) }))
			defer server.Close()
			work, err := testClient(t, server.URL).GetWork(context.Background())
			if (err != nil) != tc.wantError {
				t.Fatalf("unexpected result: %v", err)
			}
			if tc.wantError && work != nil {
				t.Fatal("failure returned work")
			}
			if !tc.wantError && (len(work) != 1 || work[0].ID != "r1") {
				t.Fatal("resource not decoded")
			}
		})
	}
}

func TestHTTPFailuresAndRedirect(t *testing.T) {
	destination := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { t.Error("redirect followed") }))
	defer destination.Close()
	for _, status := range []int{204, 302, 401, 403, 429, 503} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Location", destination.URL)
				w.Header().Set("Retry-After", "12")
				w.WriteHeader(status)
				fmt.Fprint(w, "do-not-log-this-secret")
			}))
			defer server.Close()
			_, err := testClient(t, server.URL).GetWork(context.Background())
			var httpErr *HTTPError
			if !errors.As(err, &httpErr) || httpErr.StatusCode != status || httpErr.RetryAfter != "12" {
				t.Fatalf("unexpected error: %v", err)
			}
			if strings.Contains(err.Error(), "do-not-log") {
				t.Fatal("body leaked")
			}
		})
	}
}

func TestCancellation(t *testing.T) {
	started := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { close(started); <-r.Context().Done() }))
	defer server.Close()
	c := testClient(t, server.URL)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() { <-started; cancel() }()
	_, err := c.GetWork(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation: %v", err)
	}
}

func TestOriginValidation(t *testing.T) {
	rules, err := sigv1.LoadRules("../../../contracts/sigv1-rules.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, origin := range []string{"http://example.com", "https://user:secret@example.com", "https://example.com/path", "https://example.com?x=y", "https://example.com#", "http://localhost:0", ""} {
		if _, err := New(config.Config{ServerURL: origin, AccessKey: "key", SecretKey: "secret", Rules: rules}); err == nil {
			t.Errorf("accepted invalid origin %q", origin)
		}
	}
}
