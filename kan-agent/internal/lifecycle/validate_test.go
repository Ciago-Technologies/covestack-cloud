package lifecycle

import (
	"encoding/json"
	"os"
	"testing"
)

type vector struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Kind      string `json:"kind"`
	Desired   string `json:"desiredState"`
	Deleting  bool   `json:"deleting"`
	AgentWork bool   `json:"agentWork"`
	Previous  string `json:"previousState"`
	State     string `json:"state"`
	Target    int64  `json:"targetGeneration"`
	Observed  *int64 `json:"observedGeneration"`
	Entity    string `json:"entity"`
	Result    string `json:"result"`
}

func TestSharedLifecycleVectors(t *testing.T) {
	data, err := os.ReadFile("../../../contracts/lifecycle-test-vectors.json")
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Vectors []vector `json:"vectors"`
	}
	if err = json.Unmarshal(data, &file); err != nil {
		t.Fatal(err)
	}
	for _, v := range file.Vectors {
		t.Run(v.Name, func(t *testing.T) {
			result := "invalid"
			switch v.Type {
			case "resource":
				if ValidateResource(v.Kind, v.Desired, v.Deleting, v.AgentWork) == nil {
					result = "valid"
				}
			case "observation":
				applied, err := ValidateObservation(v.Kind, v.Desired, v.Previous, v.State, v.Target, v.Observed)
				if err == nil {
					result = "ignored"
					if applied {
						result = "applied"
					}
				}
			case "server-observation":
				applied, err := ValidateServerObservation(v.Kind, v.Desired, v.Previous, v.State, v.Target, v.Observed)
				if err == nil {
					result = "ignored"
					if applied {
						result = "applied"
					}
				}
			case "transition":
				result = "denied"
				if CanTransition(v.Kind, v.Previous, v.State) {
					result = "allowed"
				}
			case "entity":
				result = "denied"
				if CanEntityTransition(v.Entity, v.Previous, v.State) {
					result = "allowed"
				}
			default:
				t.Fatal("unknown vector type")
			}
			if result != v.Result {
				t.Fatalf("got %s want %s", result, v.Result)
			}
		})
	}
}
