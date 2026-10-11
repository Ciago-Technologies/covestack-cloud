package sigv1

import (
	"encoding/json"
	"fmt"
	"os"
)

// Rules is loaded from contracts/sigv1-rules.json so Go and Java never
// disagree on header names or timing windows without both files changing.
type Rules struct {
	Algorithm                 string   `json:"algorithm"`
	DateHeader                string   `json:"dateHeader"`
	AccessKeyHeader           string   `json:"accessKeyHeader"`
	SignatureHeader           string   `json:"signatureHeader"`
	NonceHeader               string   `json:"nonceHeader"`
	AlwaysSignedHeaders       []string `json:"alwaysSignedHeaders"`
	SignBodyContentTypeHeader bool     `json:"signBodyContentTypeHeader"`
	MaxClockSkewSeconds       int      `json:"maxClockSkewSeconds"`
	NonceReplayWindowSeconds  int      `json:"nonceReplayWindowSeconds"`
	DateFormat                string   `json:"dateFormat"`
}

func LoadRules(path string) (*Rules, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var r Rules
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	required := map[string]bool{"host": false, "x-kan-date": false, "x-kan-access-key": false, "x-kan-nonce": false}
	for _, name := range r.AlwaysSignedHeaders {
		if _, ok := required[name]; ok {
			required[name] = true
		}
	}
	for name, present := range required {
		if !present {
			return nil, fmt.Errorf("unsafe rules: missing %s", name)
		}
	}
	if r.Algorithm != "KAN1-HMAC-SHA256" || r.DateHeader != "X-Kan-Date" || r.AccessKeyHeader != "X-Kan-Access-Key" || r.NonceHeader != "X-Kan-Nonce" || r.SignatureHeader != "X-Kan-Signature" || !r.SignBodyContentTypeHeader || r.MaxClockSkewSeconds < 1 || r.NonceReplayWindowSeconds < 2*r.MaxClockSkewSeconds {
		return nil, fmt.Errorf("unsafe signing rules")
	}
	return &r, nil
}
