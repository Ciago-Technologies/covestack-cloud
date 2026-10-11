package sigv1

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const MaxBodyBytes = 1024 * 1024

// NewSignedRequest signs immutable body bytes, the escaped path and raw query.
// Construct a fresh request on each retry; never reuse nonce/signature headers.
// The caller must use a client that refuses redirects to avoid credential leakage.
func NewSignedRequest(ctx context.Context, method, rawURL, accessKey, secret string,
	body []byte, contentType string, rules *Rules) (*http.Request, error) {
	if len(body) > MaxBodyBytes {
		return nil, fmt.Errorf("body exceeds 1 MiB")
	}
	if strings.ContainsAny(accessKey, "\r\n") {
		return nil, fmt.Errorf("invalid access key")
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if (parsed.Scheme != "https" && parsed.Scheme != "http") || parsed.Host == "" || parsed.User != nil || parsed.Fragment != "" {
		return nil, fmt.Errorf("expected HTTP(S) URL without userinfo/fragment")
	}
	// Detach bytes from the caller: changing them after signing must not change the wire body.
	payload := bytes.Clone(body)
	request, err := http.NewRequestWithContext(ctx, method, rawURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	nonceBytes := make([]byte, 32)
	if _, err := rand.Read(nonceBytes); err != nil {
		return nil, err
	}
	date := NowKanDate()
	request.Header.Set(rules.AccessKeyHeader, accessKey)
	request.Header.Set(rules.DateHeader, date)
	request.Header.Set(rules.NonceHeader, hex.EncodeToString(nonceBytes))
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	headers := map[string]string{"host": request.URL.Host}
	for name, values := range request.Header {
		headers[name] = values[0]
	}
	path := request.URL.EscapedPath()
	if path == "" {
		path = "/"
	}
	signature, _, _, err := Sign(secret, request.Method, path, request.URL.RawQuery, headers, payload, date, rules)
	if err != nil {
		return nil, err
	}
	request.Header.Set(rules.SignatureHeader, signature)
	return request, nil
}
