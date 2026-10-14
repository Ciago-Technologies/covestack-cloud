// auth-smoke checks the Go signer against a running Java agent API.
package main

import (
	"context"
	"fmt"
	"github.com/ciagotech/kan-agent/internal/config"
	"github.com/ciagotech/kan-agent/sigv1"
	"net/http"
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
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	base, key, secret, rules := cfg.ServerURL, cfg.AccessKey, cfg.SecretKey, cfg.Rules

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
