// Package lifecycle validates the shared contract. Controllers/drivers must also
// enforce ownership, readiness, report ordering and agent lease/fencing.
package lifecycle

import (
	"encoding/json"
	"fmt"
)

type resourceRule struct {
	Owner          string              `json:"owner"`
	InitialState   string              `json:"initialState"`
	DesiredStates  []string            `json:"desiredStates"`
	ObservedStates []string            `json:"observedStates"`
	TerminalStates []string            `json:"terminalStates"`
	Acknowledges   map[string]string   `json:"acknowledges"`
	Transitions    map[string][]string `json:"transitions"`
}
type entityRule struct {
	States      []string            `json:"states"`
	Transitions map[string][]string `json:"transitions"`
}
type catalog struct {
	Resources map[string]resourceRule `json:"resources"`
	Entities  map[string]entityRule   `json:"entities"`
}

var rules = func() catalog {
	var value catalog
	if err := json.Unmarshal([]byte(catalogJSON), &value); err != nil {
		panic(err)
	}
	return value
}()

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

// ValidateResource rejects unsupported kinds, cross-kind intent, contradictory
// deletion flags and server-only kinds in agent work. SupportedKinds must be
// checked separately against the agent's installed drivers.
func ValidateResource(kind, desired string, deleting, agentWork bool) error {
	rule, ok := rules.Resources[kind]
	if !ok {
		return fmt.Errorf("unknown resource kind: %s", kind)
	}
	if agentWork && rule.Owner != "agent" {
		return fmt.Errorf("server-managed kind cannot be agent work: %s", kind)
	}
	if !contains(rule.DesiredStates, desired) {
		return fmt.Errorf("invalid desired state for %s: %s", kind, desired)
	}
	terminalIntent := contains(rule.TerminalStates, rule.Acknowledges[desired])
	if deleting != terminalIntent {
		return fmt.Errorf("deleting contradicts desired state")
	}
	return nil
}

// CanTransition permits missed intermediate observations, but terminal states
// have no outgoing paths. Same-state reports are legal for known states only.
func CanTransition(kind, from, to string) bool {
	rule, ok := rules.Resources[kind]
	if !ok || !contains(rule.ObservedStates, from) || !contains(rule.ObservedStates, to) {
		return false
	}
	return reachable(rule.Transitions, from, to)
}
func reachable(edges map[string][]string, from, to string) bool {
	if from == to {
		return true
	}
	seen := map[string]bool{from: true}
	queue := []string{from}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, next := range edges[current] {
			if next == to {
				return true
			}
			if !seen[next] {
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}
	return false
}

// ValidateObservation returns applied=false for stale-generation reports.
// Nil generation means progress/health and never acknowledges desired work.
// Readiness and observedAt ordering are not established by a state string.
func ValidateObservation(kind, desired, previous, state string, target int64, observed *int64) (bool, error) {
	return validateObservation(kind, desired, previous, state, target, observed, "agent")
}

// ValidateServerObservation covers server-owned groups/buckets in scoped transactions.
func ValidateServerObservation(kind, desired, previous, state string, target int64, observed *int64) (bool, error) {
	return validateObservation(kind, desired, previous, state, target, observed, "server")
}
func validateObservation(kind, desired, previous, state string, target int64, observed *int64, owner string) (applied bool, err error) {
	rule, ok := rules.Resources[kind]
	if !ok || rule.Owner != owner {
		return false, fmt.Errorf("invalid observation owner/kind")
	}
	if !contains(rule.DesiredStates, desired) || !contains(rule.ObservedStates, state) {
		return false, fmt.Errorf("invalid lifecycle state")
	}
	if target < 1 {
		return false, fmt.Errorf("invalid target generation")
	}
	if observed != nil {
		if *observed < 1 || *observed > target {
			return false, fmt.Errorf("invalid or future observed generation")
		}
		if *observed < target {
			return false, nil
		}
		if state != rule.Acknowledges[desired] {
			return false, fmt.Errorf("generation acknowledged before convergence")
		}
	}
	if (state == "deleting" || contains(rule.TerminalStates, state)) && !contains(rule.TerminalStates, rule.Acknowledges[desired]) {
		return false, fmt.Errorf("deletion observation without deletion intent")
	}
	if previous == "" {
		previous = rule.InitialState
	}
	if !CanTransition(kind, previous, state) {
		return false, fmt.Errorf("invalid lifecycle transition")
	}
	return true, nil
}

// Entity transitions are direct: operation/notification updates are server-owned
// transactions, rather than sampled agent observations.
func CanEntityTransition(entity, from, to string) bool {
	rule, ok := rules.Entities[entity]
	if !ok || !contains(rule.States, from) || !contains(rule.States, to) {
		return false
	}
	return from == to || contains(rule.Transitions[from], to)
}
