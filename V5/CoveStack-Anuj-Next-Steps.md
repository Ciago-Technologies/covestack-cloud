# CoveStack V5 — Anuj Implementation and ITISS Work Order

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Current starting point

The config loader is implemented locally; signing and lifecycle helpers exist. The typed HTTP client, production agent executable and first real compute driver remain next work. Do not reapply old patches or rename compatibility packages.

## 2. Immediate implementation sequence

| Order | Code/work | Why |
|---|---|---|
| 1 | Review config tests and usage | Validated inputs before network calls |
| 2 | `internal/client` typed API/data | Bounded signed communication |
| 3 | Client tests with injected transport | Progress independent of Java readiness |
| 4 | `cmd/kan-agent` heartbeat/poll | Read-only real integration first |
| 5 | Compute driver/reconcile | One owned resource, no duplicates |
| 6 | Network/IPAM integration | Correct owned allocation and routing |
| 7 | Kubernetes driver | Bounded cluster and kubeconfig |
| 8 | Ingress and persistence work | Actual app publication and data |

These proposed paths do not exist merely because they appear here. Build one behavior per reviewed change.

## 3. Client rules

| Concern | Requirement |
|---|---|
| Context/timeout | Caller cancellation, finite HTTP timeout |
| Redirects | Refuse forwarding signed credentials |
| Request bytes | Marshal once; sign/transmit identical bytes |
| Response size | Bounded read; detect overflow |
| Status/JSON | Non-2xx and malformed data are errors |
| Retry | Fresh signed request; mutations need agreed idempotency |
| Logging | No secret structs, canonical real-key payloads or full kubeconfigs |

Coordinate endpoints, retry/idempotency, agent leases and readiness with B. Test using synthetic credentials; a fake verifier is not the deployed auth/store.

## 4. First compute acceptance

Create one owned workload, inspect it, stop/start by requested intent, restart the agent without duplication, then explicitly delete that resource. Preserve existing volumes and unrelated resources. On failure report error without acknowledging convergence. Backend/empty-poll failure must not trigger cleanup.

## 5. ITISS parallel track

| Phase | Deliverable |
|---|---|
| Foundation | VM/version baseline, CI and threat model |
| Compute | Mount/privilege review and ownership evidence |
| Network | Subnet pool, firewall and actual denied paths |
| Publication | Real TLS/tunnel route and external-device test |
| Data | Separate backup/restore exercise |
| Extended platform | Metadata boundary, automation review, alerts/incidents |

The role remains agent plus infrastructure unless the team/mentor explicitly changes ownership. Ask for a named counterpart if broad frontend/provider work creates overload; do not silently abandon assigned integration.

## 6. Learning sequence

Study structs/pointers/interfaces/errors, HTTP/JSON/readers, contexts/timers/signals, tests and then concurrency. Inspect one fixture/result before coding. Use Go documentation and the prior learning guide if present; V5 does not replace code-level walkthroughs.

## 7. Dated targets

G1 on 23 October needs real auth/client/compute; G2 on 6 November adds reliable lifecycle/networking; G3 on 20 November needs Kubernetes/public ingress. Stack/data by 4 December; automation/UI by 18 December; expanded features and evidence by 8 January. Final submission is 25 January 2027. Report dependency failures and revise gates openly.
