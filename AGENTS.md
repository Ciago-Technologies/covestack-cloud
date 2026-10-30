# CoveStack task rules

- During this task: documentation only; edit only the approved batch.
- Never create .sql or migration files; reference DDL belongs inside Markdown.
- Do not modify code, contracts/, deploy/, scripts, HTML or .gitignore.
- Preserve all pre-existing changes and compare content fingerprints.
- Never claim a feature implemented; describe target design and evidence.
- CORE-5: Droplets, required Kubernetes, Managed PostgreSQL, Object storage, VPC/IPAM.
- SRS and deliverables mention only the five services plus approved cross-cutting items.
- Extended track is separate, not in SRS/deliverables; no numeric service-count claim.
- Optional future work is separate from the named extended track.
- Service Gate Rule: Supported only when status is Gates passed after G1–G4.
- G1: DB desired state and generation are authoritative.
- G2: crash-safe reconcile; rebuild state/config from DB; no duplicates.
- G3: another account receives 404; protocol/network denial adds evidence.
- G4: service failure test runs in CI.
- Status: Planned | In progress | Gates passed | Extended track.
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
- Two drivers: generic container and k3d; networking is a container-driver helper.
- Kubernetes required: max 2 clusters per host, 1 server + max 2 agents each.
- Contract-first: lifecycle-rules.json 1.1.0; missing strings TBD with B.
- Signing credentials use service kind; API tokens are hash-only.
- Agreed: one agent process/host; separate credential/client/poll context per account; assignment/custody details B+C+A.
- Acknowledgment requires exact generation, owner and epoch; progress never acks.
- Deadline Monday 25 January 2027; planning date 14 October 2026.
- Do not delete requirements; preserve moved text with a relocation record.
- No commit, push, network or installs; stop after each batch for review.

## Mandatory task handoff — checkpoint.md

- Root checkpoint.md is the single canonical AI-agent task handoff. Read it and AGENTS.md before each task; verify current files/git status rather than trusting stale notes. Do not create competing checkpoint/handoff files.
- After every task, before the final response, update checkpoint.md in place, including documentation-only tasks, reviews and investigations. For long tasks, also checkpoint meaningful milestones before stopping, handing off or approaching context limits. This maintenance is authorized within ordinary task scope unless the latest user instruction explicitly prohibits edits or checkpoint changes; then report that the handoff was not updated.
- Record the task/objective and authorization, approach and important decisions, files changed or inspected, actual result, tests/checks with commands/results and evidence scope, unresolved risks/dependencies, remaining work and exact next action with file paths. Include branch/revision and commit/push status when checked; never invent results or remote state.
- State progress as concrete verified steps versus pending steps. If reporting a percentage, define its checklist/denominator; never guess project completion or confuse task progress with service Gates passed. Label historical evidence, mocks, unverified claims and blocked work explicitly.
- Keep the current summary and next-step instructions coherent: replace stale statements, preserve still-relevant requirements/decisions and retain a concise dated task record. Do not append contradictory snapshots or duplicate the full conversation. Include enough context for another agent to continue without chat history.
- Never store secrets, tokens, private configuration contents or sensitive logs in checkpoint.md. Preserve unrelated user changes. The checkpoint does not override contracts, source/CI evidence, user instructions or GitHub's live issue status, and does not authorize the next implementation task or a commit/push.
