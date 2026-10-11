package main

import (
	"github.com/ciagotech/kan-agent/sigv1"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestSmokeAgainstHTTPVerifier(t *testing.T) {
	path, err := filepath.Abs("../../../contracts/sigv1-rules.json")
	if err != nil {
		t.Fatal(err)
	}
	rules, err := sigv1.LoadRules(path)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	calls := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		headers := map[string]string{"Host": r.Host}
		for k, v := range r.Header {
			headers[k] = v[0]
		}
		expected, _, _, err := sigv1.Sign("test-secret", r.Method, r.URL.EscapedPath(), r.URL.RawQuery, headers, body, r.Header.Get(rules.DateHeader), rules)
		nonce := r.Header.Get(rules.NonceHeader)
		if err != nil || expected != r.Header.Get(rules.SignatureHeader) || seen[nonce] {
			w.WriteHeader(401)
			return
		}
		seen[nonce] = true
		w.WriteHeader(200)
	})
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		request.Host = request.URL.Host
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
		return recorder.Result(), nil
	})}
	t.Setenv("KAN_SERVER_URL", "http://localhost:4700")
	t.Setenv("KAN_ACCESS_KEY", "test-key")
	t.Setenv("KAN_SECRET_KEY", "test-secret")
	t.Setenv("KAN_RULES_PATH", path)
	if err := runWithClient(client); err != nil {
		t.Fatal(err)
	}
	if calls != 5 {
		t.Fatalf("got %d calls, want 5", calls)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) { return fn(request) }
