package sigv1

import (
	"encoding/json"
	"os"
	"testing"
)

// Adjust these two paths if your repo layout puts kan-agent/ somewhere
// other than next to contracts/.
const (
	rulesPath   = "../../contracts/sigv1-rules.json"
	vectorsPath = "../../contracts/sigv1-vectors.json"
)

type vector struct {
	Name              string            `json:"name"`
	Method            string            `json:"method"`
	Path              string            `json:"path"`
	Query             string            `json:"query"`
	Headers           map[string]string `json:"headers"`
	Body              string            `json:"body"`
	SecretKey         string            `json:"secretKey"`
	ExpectedSignature string            `json:"expectedSignature"`
}

type vectorFile struct {
	Vectors []vector `json:"vectors"`
}

func TestVectorsMatch(t *testing.T) {
	rules, err := LoadRules(rulesPath)
	if err != nil {
		t.Fatalf("load rules: %v (check rulesPath)", err)
	}

	data, err := os.ReadFile(vectorsPath)
	if err != nil {
		t.Fatalf("read vectors: %v (check vectorsPath)", err)
	}

	var vf vectorFile
	if err := json.Unmarshal(data, &vf); err != nil {
		t.Fatalf("parse vectors: %v", err)
	}

	for _, v := range vf.Vectors {
		dateStr := v.Headers["X-Kan-Date"]
		got, _, _ := Sign(v.SecretKey, v.Method, v.Path, v.Query, v.Headers, []byte(v.Body), dateStr, rules)
		if got != v.ExpectedSignature {
			t.Errorf("%s: got %s want %s", v.Name, got, v.ExpectedSignature)
		}
	}
}
