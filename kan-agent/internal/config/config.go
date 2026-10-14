// Package config loads agent settings from the process environment.
// It does not source .env files or print credential values.
package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ciagotech/kan-agent/sigv1"
)

// Config contains validated connection settings and shared signing rules.
// Use String for logging; do not log individual credential fields.
type Config struct {
	ServerURL string
	AccessKey string
	SecretKey string
	RulesPath string
	Rules     *sigv1.Rules
}

// String deliberately excludes credentials and user-provided paths.
func (Config) String() string     { return "agent configuration (credentials redacted)" }
func (c Config) GoString() string { return c.String() }

// Load reads KAN_* variables. The default rules path assumes execution from
// kan-agent; set KAN_RULES_PATH explicitly when running from another directory.
func Load() (Config, error) { return load(os.Getenv) }

func load(getenv func(string) string) (Config, error) {
	var c Config
	base := getenv("KAN_SERVER_URL")
	key, secret := getenv("KAN_ACCESS_KEY"), getenv("KAN_SECRET_KEY")
	for _, field := range []struct{ name, value string }{
		{"KAN_SERVER_URL", base}, {"KAN_ACCESS_KEY", key}, {"KAN_SECRET_KEY", secret},
	} {
		if strings.TrimSpace(field.value) == "" {
			return c, fmt.Errorf("%s is required", field.name)
		}
	}
	// Credentials are opaque: reject whitespace rather than silently changing them.
	if strings.ContainsAny(key, " \t\r\n") || strings.ContainsAny(secret, " \t\r\n") {
		return c, fmt.Errorf("credentials must not contain whitespace")
	}
	u, err := url.Parse(base)
	if err != nil {
		return c, fmt.Errorf("KAN_SERVER_URL must be a valid HTTP(S) origin")
	}
	if u.User != nil || u.Opaque != "" || u.Hostname() == "" || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" || strings.Contains(base, "#") || (u.Path != "" && u.Path != "/") || u.RawPath != "" {
		return c, fmt.Errorf("KAN_SERVER_URL must be an origin without credentials, path, query or fragment")
	}
	if strings.HasSuffix(u.Host, ":") {
		return c, fmt.Errorf("KAN_SERVER_URL has an invalid port")
	}
	if port := u.Port(); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil || n < 1 || n > 65535 {
			return c, fmt.Errorf("KAN_SERVER_URL has an invalid port")
		}
	}
	host := strings.ToLower(u.Hostname())
	ip := net.ParseIP(host)
	local := host == "localhost" || (ip != nil && ip.IsLoopback())
	if u.Scheme != "https" && !(u.Scheme == "http" && local) {
		return c, fmt.Errorf("KAN_SERVER_URL requires HTTPS outside localhost")
	}
	path := getenv("KAN_RULES_PATH")
	if path == "" {
		path = "../contracts/sigv1-rules.json"
	}
	rules, err := sigv1.LoadRules(path)
	if err != nil {
		return c, fmt.Errorf("KAN_RULES_PATH must reference a readable, valid signing rules file")
	}
	c = Config{ServerURL: strings.TrimSuffix(u.String(), "/"), AccessKey: key, SecretKey: secret, RulesPath: path, Rules: rules}
	return c, nil
}
