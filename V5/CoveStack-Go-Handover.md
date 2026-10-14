# CoveStack V5 — Go Ownership and Learner Handover

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Boundary

| Owner | May change | Must not change |
|---|---|---|
| A — Anuj | kan-agent, container/k3d drivers, network helper, reconcile, signing, pkg/kanclient, CI | Do not hand agent responsibility to E silently |
| E — other Go learner | cmd/covestack thin cobra client, JSON, cross-build packaging; later thin provider | Agent, signing, drivers and reconcile |
| B/C/D | Java/SQL/policy/schema/API roles | D writes no Go |

**AI-generated code is reference only: A types and can explain agent Go.** Any Go change needs A review; pkg/kanclient needs A+B+C; k3d driver A+B+C; schema B+C+A. The module layout and signer move are planned decisions, not changes in this task.

## 2. E's four-week learning plan

Weeks are learning stages, not invented calendar dates. Production CLI service order remains droplet → vpc → k8s → db → bucket → stack, with each previous failure test in CI before the next service starts.

| Week | Deliverable | Review gate |
|---|---|---|
| W1 | Basics/modules/structs/errors; cobra hello; linux/amd64, darwin/arm64, windows/amd64 builds | A: explain types/errors/module boundary; binaries build |
| W2 | net/http, JSON, context/timeouts against mock API | A: timeout/cancel/non-2xx/JSON tests; no runtime dependencies |
| W3 | Droplet create/list/get/delete, --wait, JSON and auto idempotency; mock-only vpc/k8s list/get practice | A: API-only boundaries; mock practice does not start production K8s work |
| W4 | Table tests, httptest and CLI against real server | A+C: ownership/idempotency/generation results; actual OS run evidence |

## 3. CLI rules

| Rule | Requirement |
|---|---|
| Platform imports | pkg/kanclient only; cobra and standard-library use allowed |
| Backend | Authenticated API only; no Docker/k3d/DB access |
| Config | Server URL, key ID, profile only; no local database |
| Secrets | Env/OS keychain; not arguments/logs |
| Transport | Finite timeouts and no credential redirects |
| Mutations | Every request carries an idempotency key |
| Create | Returns pending immediately; --wait polls convergence |
| Get | Desired/converged generation, lease, observed state |
| Automation | Every list/get supports -o json |

## 4. A's review checklist

| Check | Evidence |
|---|---|
| Understand code | E explains each changed flow and error path |
| Dependency boundary | Go dependency scan contains no Docker/k3d |
| API policy | C's account/idempotency tests cover real server |
| Wait/result | Pending, cancellation and convergence are distinct |
| Output | Stable JSON and no credential output |
| Packaging | Three target builds; two actual OS demonstrations |

Ask the mentor in writing whether ITISS students may own agent/driver Go. Confirmation is pending; if prohibited, name an agent developer and record handover. No mentor decision is asserted here.

[Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Gates](CoveStack-Service-Gates.md) · [Schema](CoveStack-Database-Schema.md) · [SRS](CoveStack-SRS-Scope.md) · [Roadmap](CoveStack-Extended-Roadmap.md).
