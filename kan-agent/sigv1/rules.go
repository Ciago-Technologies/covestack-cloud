package sigv1

import (
	"encoding/json"
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
	return &r, nil
}
