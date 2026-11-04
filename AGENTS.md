# CoveStack task rules

- Product authority: docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md; proposed product core does not replace academic CORE-5. No Labs/learning-management functionality or new release plan. Object-storage provider remains under evaluation.

- During this task: documentation only; edit only the approved batch.
- Never create .sql or migration files; reference DDL belongs inside Markdown.
- Do not modify code, contracts/, deploy/, scripts, HTML or .gitignore.
- Preserve all pre-existing changes and compare content fingerprints.
- Never claim implementation from documentation alone. Record code-confirmed, tested, partial, planned/proposed/blocked/deferred evidence separately; service Supported still requires G1–G4.
- Preserve academic CORE-5 categories/tests: compute, required Kubernetes, Managed PostgreSQL, Object storage, VPC/IPAM. New target compute is a true VM Instance; separate App Container requires explicit academic mapping review, never a silent waiver.
- SRS and deliverables mention only the five services plus approved cross-cutting items.
- Extended track is separate, not in SRS/deliverables; no numeric service-count claim.
- Optional future work is separate from the named extended track.
- Service Gate Rule: Supported requires Acceptance Gates Passed after G1–G4 plus a qualified profile and maintenance policy; historical V5 Gates passed maps to Acceptance Gates Passed.
- G1: DB desired state and generation are authoritative.
- G2: crash-safe reconcile; rebuild state/config from DB; no duplicates.
- G3: another account receives 404; protocol/network denial adds evidence.
- G4: service failure test runs in CI.
- Active service status: Proposed | Planned | In Progress | Partially Implemented | Implemented | Runtime Verified | Acceptance Gates Passed | Supported | Deferred | Blocked. Historical labels remain historical; require evidence per docs/platform/readiness.md.
- CORE-5 tests 1–15; extended tests E1–E8; C owns all acceptance tests.
- A: all agent/client/driver/signing/reconcile Go, CI, Part 1 and ITISS.
- E: cobra CLI, JSON, packaging; later thin Terraform provider; no agent Go.
- B: Java core, all SQL, leases/fencing/idempotency, quotas and stack records.
- C: auth/security, all acceptance tests and Java IPAM allocator.
- D: Java /v2/DTO/schema and Terraform fixtures; writes no Go.
- Any Go change: A review; pkg/kanclient and k3d: A+B+C; schema: B+C+A.
- CLI never imports or calls Docker or k3d; authenticated API only.
- CLI holds config only; mutations use idempotency; list/get offer JSON.
- AI-generated code is reference only, A types and explains agent Go.
- VM-first target: evaluate Incus/QEMU/KVM; distinct trusted App Container and k3d backends; explicit network/storage capabilities. Old two-driver ceiling is superseded, not implemented. No privileged runtime sockets/host mounts for tenants.
- Kubernetes required: max 2 clusters per host, 1 server + max 2 agents each.
- Contract-first: lifecycle-rules.json 1.1.0; missing strings TBD with B.
- Signing credentials use service kind; API tokens are hash-only.
- Agreed: one agent process/host; separate credential/client/poll context per account; assignment/custody details B+C+A.
- Acknowledgment requires exact generation, owner and epoch; progress never acks.
- Planning began 14 October 2026. Old V5 deadline 25 January 2027 conflicts with latest user-stated 15 December; confirm revised year/scope, never invent mentor approval or release dates.
- Do not delete requirements; preserve moved text with a relocation record.
- No commit, push, network or installs; stop after each batch for review.

## Mandatory task handoff — checkpoint.md

- Root checkpoint.md is the single canonical AI-agent task handoff. Read it and AGENTS.md before each task; verify current files/git status rather than trusting stale notes. Do not create competing checkpoint/handoff files.
- After every task, before the final response, update checkpoint.md in place, including documentation-only tasks, reviews and investigations. For long tasks, also checkpoint meaningful milestones before stopping, handing off or approaching context limits. This maintenance is authorized within ordinary task scope unless the latest user instruction explicitly prohibits edits or checkpoint changes; then report that the handoff was not updated.
- Record the task/objective and authorization, approach and important decisions, files changed or inspected, actual result, tests/checks with commands/results and evidence scope, unresolved risks/dependencies, remaining work and exact next action with file paths. Include branch/revision and commit/push status when checked; never invent results or remote state.
- State progress as concrete verified steps versus pending steps. If reporting a percentage, define its checklist/denominator; never guess project completion or confuse task progress with service Gates passed. Label historical evidence, mocks, unverified claims and blocked work explicitly.
- Keep the current summary and next-step instructions coherent: replace stale statements, preserve still-relevant requirements/decisions and retain a concise dated task record. Do not append contradictory snapshots or duplicate the full conversation. Include enough context for another agent to continue without chat history.
- Never store secrets, tokens, private configuration contents or sensitive logs in checkpoint.md. Preserve unrelated user changes. The checkpoint does not override contracts, source/CI evidence, user instructions or GitHub's live issue status, and does not authorize the next implementation task or a commit/push.

## Active documentation authority

- Start docs/platform/README.md, resource-model.md and implementation-status.md for the new VM-first target. V5 body is superseded design/academic history; current 1.1.0 contracts/source remain unchanged and cannot execute invented VM/claim fields.
- Incus is preferred evaluation only; 8 GB is a lab constraint, never a production profile or density guarantee. No VM-to-container fallback. Single-host hardening is not HA.
- Floci is optional external emulation integration, not a mandatory core, fork decision or wholesale Java-to-Go rewrite. Java control plane stays; product-owned new Go components need reviewed boundaries/ownership.
- No license/support/production readiness inferred: no root license grant or supported profile exists. Preserve historical evidence and actual environment restrictions; runtime acceptance requires a separately authorized capable host.
