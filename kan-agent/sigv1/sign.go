// Package sigv1 signs outgoing requests from kan-agent to kan-server.
//
// The agent only SIGNS. It never verifies anything, because it has no
// open port to receive requests on — it only calls out to kan-server.
// Verification lives entirely on the Java side (see KanSigV1AuthFilter).
package sigv1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"
)

func sha256Hex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// BuildCanonicalRequest returns the canonical request string and the
// signed-headers list, per CANONICAL-REQUEST.md.
func BuildCanonicalRequest(method, path, query string, headers map[string]string,
	body []byte, rules *Rules) (canonicalRequest string, signedHeaders string, err error) {

	signSet := map[string]bool{}
	for _, h := range rules.AlwaysSignedHeaders {
		signSet[strings.ToLower(h)] = true
	}
	if len(body) > 0 && rules.SignBodyContentTypeHeader {
		signSet["content-type"] = true
	}

	lower := map[string]string{}
	for k, v := range headers {
		name := strings.ToLower(k)
		if _, exists := lower[name]; exists {
			return "", "", fmt.Errorf("duplicate header: %s", name)
		}
		lower[name] = strings.Trim(v, " \t")
	}

	var names []string
	for name := range signSet {
		value, ok := lower[name]
		if !ok || value == "" || strings.ContainsAny(value, "\r\n") {
			return "", "", fmt.Errorf("missing or invalid signed header: %s", name)
		}
		names = append(names, name)
	}
	sort.Strings(names)

	var sb strings.Builder
	for _, n := range names {
		sb.WriteString(n)
		sb.WriteString(":")
		sb.WriteString(lower[n])
		sb.WriteString("\n")
	}
	signedHeaders = strings.Join(names, ";")
	hashedPayload := sha256Hex(body)

	canonicalRequest = strings.Join([]string{
		method, path, query, sb.String(), signedHeaders, hashedPayload,
	}, "\n")
	return canonicalRequest, signedHeaders, nil
}

func stringToSign(rules *Rules, dateStr, canonicalRequest string) string {
	return strings.Join([]string{
		rules.Algorithm, dateStr, sha256Hex([]byte(canonicalRequest)),
	}, "\n")
}

// Sign returns the hex signature to send in the X-Kan-Signature header.
func Sign(secretKey, method, path, query string, headers map[string]string,
	body []byte, dateStr string, rules *Rules) (signature, canonicalRequest, signedHeaders string, err error) {

	cr, sh, err := BuildCanonicalRequest(method, path, query, headers, body, rules)
	if err != nil {
		return "", "", "", err
	}
	lowerDate := ""
	for name, value := range headers {
		if strings.EqualFold(name, rules.DateHeader) {
			lowerDate = strings.Trim(value, " \t")
		}
	}
	if lowerDate != dateStr || secretKey == "" {
		return "", "", "", fmt.Errorf("date mismatch or empty signing secret")
	}
	sts := stringToSign(rules, dateStr, cr)

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(sts))
	sig := hex.EncodeToString(mac.Sum(nil))

	return sig, cr, sh, nil
}

// NowKanDate formats the current time the way kan-server expects it,
// for the X-Kan-Date header.
func NowKanDate() string {
	return time.Now().UTC().Format("20060102T150405Z")
}
