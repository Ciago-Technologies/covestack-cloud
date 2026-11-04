# CoveStack V5 — Go Ownership and CLI Learner Handover

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. A versus E boundary

| Area | A / Anuj | E |
| --- | --- | --- |
| Agent/client/signing/reconcile/runtime providers | Owns and explains implementation | Does not modify/implement |
| Shared pkg/kanclient target | Reviews/designs with B+C and consumer E; not moved yet | Uses approved API boundary; any shared change needs A+B+C review |
| CLI/cobra/JSON/wait/errors | Reviews Go and runtime/API separation | Owns thin authenticated user client |
| Packaging | Reviews toolchain/CI/security | Owns approved CLI OS builds and later real OS-run evidence |
| Optional provider | Reviews Go/boundary with B/C/D | Later thin Terraform/OpenTofu provider only after APIs/CLI |

All shared/API/module moves are Planned. No current CLI/provider module is established. Ordinary CLI user auth differs from agent service signing; do not accidentally reuse privileged credentials to make a demo work.

## 2. E's four learning stages — not dated delivery promises

| Stage | Learn and deliver | Review gate |
| --- | --- | --- |
| W1 | Go modules/structs/errors, explainable cobra hello after module layout review; target OS build exercise | A reviews imports; linux/amd64, darwin/arm64, windows/amd64 build evidence when produced |
| W2 | net/http, JSON, context, timeout/cancel/errors against clearly labelled mock | A+C negative tests, bounded HTTP, no redirects/credential logs |
| W3 | Practice owned compute create/list/get/delete, idempotency, finite wait and JSON against mock; network/Kubernetes list/get is mock-only practice | No fake provisioned resources or VM driver claims; approved API shapes only |
| W4 | Table tests/httptest and real user-server integration when available | A code review, C actual auth/owner/idempotency assertions; packaging is not proof runtime runs on every OS |

The learning stages do not select future CLI command names or wire kinds. Do not copy historical Droplet/container specs into VM requests. Real service order follows available owned APIs and failure tests, not a hardcoded old academic chain.

## 3. CLI rules

API-only, no Docker/Incus/k3d/DB import or privileged socket. Config file stores endpoint/profile/key identifier only; no local resource DB. Secrets from approved environment/keychain/custody, not command arguments/logs. Mutations use logical idempotency; retries use reviewed API semantics. Create returns pending; finite wait/cancel shows actual convergence/error. List/get stable JSON and scoped redacted table view. Secret export separately authorized; no kubeconfig in ordinary output.

## 4. A's review checklist

- Explainable code and correct module/user-versus-agent auth boundary.
- Context/timeout/cancel and strict HTTP/JSON/error handling; no credential redirects/log leaks.
- Idempotency and observed wait semantics match real API; no pretend ready from process/row presence.
- Table/mock tests labelled; real auth/foreign ID and negative behavior shown separately.
- No runtime driver or DB dependency in user CLI; exact packaging artifacts and later OS-run evidence.
- No arbitrary package/signer relocation or speculative provider before actual lifecycle API.

## 5. Current first task and next handshake

A uses existing signed no-execution loop with B+C real assigned server credentials. E agrees module/client plan with A and can learn mocks without bypassing a missing server. B/D/C review API/DTO/security; C owns acceptance. [Team checkpoint](../TEAM-FIRST-CHECKPOINT.md), [Anuj tasks](CoveStack-Anuj-Next-Steps.md), [automation](CoveStack-Terraform-and-Ansible.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-go-handover) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
