package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func settings(t *testing.T) map[string]string {
	t.Helper()
	path, err := filepath.Abs("../../../contracts/sigv1-rules.json")
	if err != nil {
		t.Fatal(err)
	}
	return map[string]string{"KAN_SERVER_URL": "https://agent.example.test", "KAN_ACCESS_KEY": "private-access-key", "KAN_SECRET_KEY": "private-signing-secret", "KAN_RULES_PATH": path}
}

func TestServerOriginValidation(t *testing.T) {
	for _, tc := range []struct {
		url   string
		valid bool
	}{
		{"https://agent.example.test", true}, {"https://agent.example.test:8443/", true},
		{"http://localhost:4700", true}, {"http://127.0.0.1:4700", true}, {"http://[::1]:4700", true},
		{"http://agent.example.test", false}, {"http://localhost.evil.test", false},
		{"https:///missing-host", false}, {"https://user:password@agent.example.test", false},
		{"https://agent.example.test/api", false}, {"https://agent.example.test?", false},
		{"https://agent.example.test?token=private-signing-secret", false},
		{"https://agent.example.test#", false}, {"https://agent.example.test#fragment", false},
		{"https://agent.example.test/%2f", false}, {"ftp://localhost", false},
		{"https://agent.example.test:0", false}, {"https://agent.example.test:65536", false},
		{"https://agent.example.test:", false}, {"https://agent.example.test:bad", false},
	} {
		t.Run(tc.url, func(t *testing.T) {
			env := settings(t)
			env["KAN_SERVER_URL"] = tc.url
			cfg, err := load(func(k string) string { return env[k] })
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v, error=%v", tc.valid, err)
			}
			if err != nil && strings.Contains(err.Error(), "private-signing-secret") {
				t.Fatal("error leaked secret")
			}
			if tc.valid && (cfg.Rules == nil || strings.HasSuffix(cfg.ServerURL, "/")) {
				t.Fatal("configuration not ready for requests")
			}
		})
	}
}

func TestRequiredSettingsAndRules(t *testing.T) {
	for _, name := range []string{"KAN_SERVER_URL", "KAN_ACCESS_KEY", "KAN_SECRET_KEY"} {
		t.Run(name, func(t *testing.T) {
			env := settings(t)
			env[name] = " \t"
			cfg, err := load(func(k string) string { return env[k] })
			if err == nil || !strings.Contains(err.Error(), name) {
				t.Fatalf("expected missing %s", name)
			}
			if cfg.SecretKey != "" {
				t.Fatal("returned partial credentials on failure")
			}
		})
	}
	env := settings(t)
	env["KAN_RULES_PATH"] = filepath.Join(t.TempDir(), "private-signing-secret")
	_, err := load(func(k string) string { return env[k] })
	if err == nil || strings.Contains(err.Error(), "private-signing-secret") {
		t.Fatal("rules error missing or leaked path")
	}
	env = settings(t)
	invalid := filepath.Join(t.TempDir(), "rules.json")
	if err := os.WriteFile(invalid, []byte(`{"algorithm":"unsafe"}`), 0600); err != nil {
		t.Fatal(err)
	}
	env["KAN_RULES_PATH"] = invalid
	if _, err := load(func(k string) string { return env[k] }); err == nil {
		t.Fatal("accepted unsafe rules")
	}
	env = settings(t)
	env["KAN_ACCESS_KEY"] = "key\r\nInjected: header"
	if _, err := load(func(k string) string { return env[k] }); err == nil {
		t.Fatal("accepted header injection")
	}
}

func TestLoadAndRedactedFormatting(t *testing.T) {
	env := settings(t)
	for k, v := range env {
		t.Setenv(k, v)
	}
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.AccessKey != env["KAN_ACCESS_KEY"] || cfg.SecretKey != env["KAN_SECRET_KEY"] {
		t.Fatal("credentials changed")
	}
	for _, format := range []string{"%v", "%+v", "%#v"} {
		s := fmt.Sprintf(format, cfg)
		if strings.Contains(s, cfg.AccessKey) || strings.Contains(s, cfg.SecretKey) {
			t.Fatal("format leaked credentials")
		}
	}
}
