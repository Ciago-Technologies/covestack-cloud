package sigv1

import (
	"context"
	"io"
	"strings"
	"testing"
)

func testHeaders() map[string]string {
	return map[string]string{"Host": "localhost:4700", "X-Kan-Date": "20260926T120000Z",
		"X-Kan-Access-Key": "KANKAGENT0000000001", "X-Kan-Nonce": "negative-test-nonce-0001"}
}
func TestRequiredHeaders(t *testing.T) {
	rules, err := LoadRules(rulesPath)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Host", "X-Kan-Date", "X-Kan-Access-Key", "X-Kan-Nonce"} {
		t.Run(name, func(t *testing.T) {
			headers := testHeaders()
			delete(headers, name)
			if _, _, _, err := Sign("secret", "GET", "/", "", headers, nil, "20260926T120000Z", rules); err == nil {
				t.Fatal("missing signed header accepted")
			}
		})
	}
	headers := testHeaders()
	headers["host"] = "other"
	if _, _, _, err := Sign("secret", "GET", "/", "", headers, nil, "20260926T120000Z", rules); err == nil {
		t.Fatal("duplicate accepted")
	}
	headers = testHeaders()
	if _, _, _, err := Sign("secret", "PUT", "/", "", headers, []byte("{}"), "20260926T120000Z", rules); err == nil {
		t.Fatal("missing content-type accepted")
	}
}
func TestNonceAndKeyAreSigned(t *testing.T) {
	rules, err := LoadRules(rulesPath)
	if err != nil {
		t.Fatal(err)
	}
	original, _, _, err := Sign("secret", "GET", "/", "", testHeaders(), nil, "20260926T120000Z", rules)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"X-Kan-Nonce", "X-Kan-Access-Key"} {
		headers := testHeaders()
		headers[name] += "changed"
		changed, _, _, err := Sign("secret", "GET", "/", "", headers, nil, "20260926T120000Z", rules)
		if err != nil {
			t.Fatal(err)
		}
		if changed == original {
			t.Fatalf("%s not signed", name)
		}
	}
}
func TestNewSignedRequest(t *testing.T) {
	rules, err := LoadRules(rulesPath)
	if err != nil {
		t.Fatal(err)
	}
	payload := []byte("{}")
	ctx := context.Background()
	request, err := NewSignedRequest(ctx, "PUT", "http://localhost:4700/internal/v1/agent/a%2Fb?x=2&x=1", "key", "secret", payload, "application/json", rules)
	if err != nil {
		t.Fatal(err)
	}
	payload[0] = 'x'
	actual, err := io.ReadAll(request.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != "{}" {
		t.Fatal("body changed after signing")
	}
	headers := map[string]string{"Host": request.URL.Host}
	for k, v := range request.Header {
		headers[k] = v[0]
	}
	sig, _, _, err := Sign("secret", request.Method, request.URL.EscapedPath(), request.URL.RawQuery, headers, actual, request.Header.Get(rules.DateHeader), rules)
	if err != nil || sig != request.Header.Get(rules.SignatureHeader) {
		t.Fatal("request does not match signature", err)
	}
	next, err := NewSignedRequest(ctx, "GET", "http://localhost:4700/", "key", "secret", nil, "", rules)
	if err != nil {
		t.Fatal(err)
	}
	if next.Header.Get(rules.NonceHeader) == request.Header.Get(rules.NonceHeader) {
		t.Fatal("nonce reused")
	}
	if _, err := NewSignedRequest(ctx, "POST", "http://localhost:4700/", "key", "secret", []byte(strings.Repeat("x", MaxBodyBytes+1)), "application/json", rules); err == nil {
		t.Fatal("oversized body accepted")
	}
}
