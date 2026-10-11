// auth-smoke checks the Go signer against a running Java agent API.
package main

import (
	"context"
	"fmt"
	"github.com/ciagotech/kan-agent/sigv1"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run() error {
	client := &http.Client{Timeout: 10 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	return runWithClient(client)
}

func runWithClient(client *http.Client) error {
	base := os.Getenv("KAN_SERVER_URL")
	key, secret := os.Getenv("KAN_ACCESS_KEY"), os.Getenv("KAN_SECRET_KEY")
	if base == "" || key == "" || secret == "" {
		return fmt.Errorf("set KAN_SERVER_URL, KAN_ACCESS_KEY and KAN_SECRET_KEY")
	}
	parsed, err := url.Parse(base)
	if err != nil {
		return err
	}
	if parsed.Scheme != "https" && !(parsed.Scheme == "http" && (parsed.Hostname() == "localhost" || parsed.Hostname() == "127.0.0.1" || parsed.Hostname() == "::1")) {
		return fmt.Errorf("use HTTPS outside localhost")
	}
	if parsed.Path != "" && parsed.Path != "/" || parsed.RawQuery != "" || parsed.Fragment != "" {
		return fmt.Errorf("server URL must be an origin, without path/query/fragment")
	}
	rulesPath := os.Getenv("KAN_RULES_PATH")
	if rulesPath == "" {
		rulesPath = "../contracts/sigv1-rules.json"
	}
	rules, err := sigv1.LoadRules(rulesPath)
	if err != nil {
		return err
	}

	target := strings.TrimRight(base, "/") + "/internal/v1/agent/work"
	fresh := func() (*http.Request, error) {
		return sigv1.NewSignedRequest(context.Background(), "GET", target, key, secret, nil, "", rules)
	}
	check := func(name string, request *http.Request, want int) error {
		response, err := client.Do(request)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		response.Body.Close()
		if response.StatusCode != want {
			return fmt.Errorf("%s: status %d, expected %d", name, response.StatusCode, want)
		}
		fmt.Printf("PASS %s (%d)\n", name, response.StatusCode)
		return nil
	}
	valid, err := fresh()
	if err != nil {
		return err
	}
	replay := valid.Clone(context.Background())
	if err := check("valid signed work request", valid, 200); err != nil {
		return err
	}
	if err := check("identical nonce replay", replay, 401); err != nil {
		return err
	}
	bad, err := fresh()
	if err != nil {
		return err
	}
	original := bad.Clone(context.Background())
	bad.Header.Set(rules.SignatureHeader, strings.Repeat("0", 64))
	if err := check("invalid signature", bad, 401); err != nil {
		return err
	}
	if err := check("valid request after nonce poisoning attempt", original, 200); err != nil {
		return err
	}
	tampered, err := fresh()
	if err != nil {
		return err
	}
	tampered.Header.Set(rules.NonceHeader, tampered.Header.Get(rules.NonceHeader)+"a")
	return check("nonce tampering", tampered, 401)
}
