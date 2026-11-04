# CoveStack — first team checkpoint and GitHub board guide

Updated: **4 November 2026**. Audience: **A/Anuj, B, C, D and E**.

**Keep this file and Starter.md.** [Starter.md](Starter.md) covers cloning, branching and getting started. This file defines the team's next integration checkpoint: who does what, what blocks progress, which evidence to collect and how to track it on GitHub. It does not claim the checkpoint has passed or impose a product deadline.

## 1. Which record should you follow?

| Record | Authority | What to update |
| --- | --- | --- |
| [Product SRS](docs/product/CoveStack-Product-SRS.md) and [catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) | Product requirements and core/future scope | Reviewed architecture/scope decisions |
| [Architecture](docs/platform/architecture.md), [API boundaries](docs/platform/api-and-cli.md) and actual contracts | Target boundaries versus current wire behavior | Coordinated design/schema/fixture changes, never invented wire strings |
| [Starter.md](Starter.md) and this file | Team onboarding, role boundaries and checkpoint acceptance | Update when the agreed next slice changes; do not maintain daily completion ticks here |
| GitHub issues/project board | Live assignee, status, dependencies, blocker and linked PR | Owner updates the same issue as work progresses |
| PR/CI/runtime evidence | Exact revision, commands and actual results | Attach evidence to the issue; label source/unit/mock/integration/runtime scope |
| [checkpoint.md](checkpoint.md) | Single canonical AI-agent handoff | Agent records task authorization, decisions, changes, verification and next action |

Do not make this team guide a second AI handoff or a second live task database. Link issues to this checkpoint; the board is a view of those issues. No GitHub items were created or inspected live by this documentation update.

## 2. Current architecture and actual starting point

| Area | Current direction / evidence boundary |
| --- | --- |
| Product | Local-first, account-owned real infrastructure; scope is independent of historical academic CORE-5 |
| Compute | Real VM Instance is the primary target; Incus/QEMU/KVM preferred evaluation. App Container is separate and shared-kernel |
| Control plane | Java/Spring Boot with JdbcClient/plain SQL; PostgreSQL owns desired state/operations. No JPA/Hibernate |
| Agent | Go signed client and single-account polling exist; returned work deliberately stops the loop without execution |
| Missing integration | Concrete credential provisioning, trusted account/host assignment, runnable scoped server foundation and real-server evidence remain pending |
| Execution | Claims, VM/container/network/storage drivers, reconciliation and capacity admission are target work, not verified service support |
| Object storage | RustFS preferred planned provider; operator-managed process/volumes, account-owned buckets/keys, Java-side policy and credential reconciliation |
| CLI | E's thin authenticated Go API client is Planned; no direct runtime socket or platform DB access |
| License | [Apache License 2.0](LICENSE), SPDX Apache-2.0, adopted for CoveStack-owned material; third-party terms remain separate |

Existing Go test sources, Java auth components or extended SQL files do not establish a bootable platform, live schema or provisioned service. Read the [implementation ledger](docs/platform/implementation-status.md) before assigning a task. Old academic dates, cluster caps and template counts are historical, not product requirements. No new delivery date or runtime approval is invented here.

## 3. The next shared checkpoint

**Goal: the existing single-account signed agent polls a real authenticated Java endpoint with trusted assignment, handles no work/errors and stops cleanly—without executing resources.**

The immediate checkpoint is a safe management-plane integration, not a VM/RustFS feature demo. A must not recreate the client/loop; B+C must not bypass identity to supply work. D reviews real versus proposed interfaces. E can learn with labelled mocks while the server dependencies are pending.

| Step | Lead / collaborators | Concrete exit evidence | Dependency |
| --- | --- | --- | --- |
| 1. Agree the handshake | B+C+A; D compatibility | Recorded endpoint/no-work behavior, credential kind/provisioning path, trusted host/account assignment and protected delivery; no secret values | Read source and current contract together |
| 2. Make the server slice testable | B+C; D API mapping | Reviewed Java startup/auth and prerequisite migration path on an authorized environment; actual results recorded | Credential/assignment design and migration history review |
| 3. Prove authentication | C+B; A client review | Real signed request accepted for the assigned identity; invalid/revoked/replayed/wrong-assignment cases rejected | Runnable server and protected test identities |
| 4. Integrate existing agent | A+C | Actual no-work polling, cancellation/auth-stop/error observations; returned work never claimed or executed | Steps 1–3 |
| 5. Review next execution design | D+B+A+C | Future claim/renew/report, VM/type/network/capacity gaps and fixture plan recorded | Current integration evidence; no silent contract reinterpretation |
| 6. Connect CLI later | E with A/C; B/D API | Thin API behavior against a real available endpoint with user auth and truthful JSON/error output | Agreed module boundary and real user API; mock work remains separately labelled |

Steps 1–4 define the immediate integration exit. Steps 5–6 prepare later work; do not block a clear account of what did and did not pass. No step is pre-marked complete by this document.

## 4. A / Anuj — first tasks and limits

| Order | Your work | Relevant files / review | Evidence to attach |
| --- | --- | --- | --- |
| 1 | Explain the existing config → client/signing → poll → cancellation/no-execution path | kan-agent/cmd/kan-agent/main.go; internal/config/; internal/client/; internal/agent/loop.go; sigv1/ | Short explanation and exact source revision; no duplicate implementation |
| 2 | Agree real URL, credential delivery and trusted assignment with B+C | Current [internal API](contracts/internal-api.yaml); [API boundary](docs/platform/api-and-cli.md) | Written non-secret agreement and blockers |
| 3 | Run relevant existing Go checks on the authorized dev environment | kan-agent/; CI commands in Starter.md | Actual command/results; mocks and unit checks clearly labelled |
| 4 | Test the existing loop against B+C's real server | Existing client/loop; C's integration tests | Empty-array behavior, auth rejection, cancellation, observed outage/retry behavior |
| 5 | Review future execution boundary before coding drivers | A+B+C+D, VM/type and claim design | Agreed scope/fixtures and separately approved backend evaluation plan |

Current discovery is `GET /internal/v1/agent/work`; `200` with `[]` means no work. Environment names are `KAN_SERVER_URL`, `KAN_ACCESS_KEY`, `KAN_SECRET_KEY`, optional `KAN_RULES_PATH`. The loader does not automatically read .env. A localhost example is not a verified server URL. Do not paste secrets into an issue, command argument, shared screenshot or this file.

Your later ownership remains agent/shared client/signing/drivers/reconcile, CI, installer and ITISS evidence. You review E's Go but do not take over the CLI. Runtime evaluation requires an authorized capable host; no tenant privileged sockets, host root mounts or arbitrary host execution. Explain the code you contribute.

## 5. B — server foundation and durable state

| Order | Your work | First output / evidence |
| --- | --- | --- |
| 1 | Inspect kan-server source/pom and actual migration prerequisites; do not assume V5/V6 form a clean schema chain | Reviewed startup/migration dependency plan and blockers |
| 2 | Agree credential lookup/auth wiring with C and assignment/discovery needs with A+D | A bounded authenticated server slice, not a fake key or anonymous fallback |
| 3 | Own application/configuration and all SQL/migrations, idempotency, claims/fencing, quota/reservations | Reviewed forward-only migration plan; actual empty-DB/startup evidence when implemented |
| 4 | Make the real signed no-work path testable with C | Reproducible server/test setup, sanitized results and integration dependencies |

Work in kan-server/. B owns SQL; C reviews tenancy/secrets/tests, A agent-facing state and D DTO alignment. Establish applied history before migrations; never edit an applied migration or copy historical container reference DDL directly into a VM schema. Java's RustFS reconciliation is later work; object data still requires independent persistence/restore evidence.

## 6. C — identity, security and acceptance from day one

| Order | Your work | First output / evidence |
| --- | --- | --- |
| 1 | Read existing auth/replay interfaces and current signing rules with B+A | Credential/assignment/custody decision and threat boundary |
| 2 | Design real credential lookup/provisioning, protected storage, rotation/revocation and replay behavior | Valid/invalid/revoked/replayed/wrong-assignment cases; no shared default key |
| 3 | Own acceptance for the real server/client integration | Reproducible tests and exact evidence class/revision; no unit-only runtime claim |
| 4 | Later own Java IPAM and service isolation/recovery acceptance | Actual allocation/traffic/storage denial and verified restoration, not merely DB records |

Work in kan-server auth/test areas and reviewed acceptance specifications. Keep hash-only API tokens separate from encrypted signing credentials; agent service identity is not ordinary CLI user identity. A request's claimed account/agent is not trusted assignment by itself. C owns acceptance; implementers contribute evidence and fixes.

## 7. D — Java API, DTO and contract review

| Order | Your work | First output / evidence |
| --- | --- | --- |
| 1 | Inventory actual Java endpoints/interfaces versus missing server APIs | Reviewed list of existing behavior, gaps and next DTO/fixture slice |
| 2 | Review discovery/no-work/error/ownership semantics with B+C+A | Fixtures aligned with actual contract; no invented fields |
| 3 | Record future VM/App Container, assignment/claims/network/capacity changes | Coordinated schema/fixture/generated-output proposal, not an assumed current wire version |
| 4 | Later map bounded public APIs and Terraform fixtures | Exact supported subset verified against real server behavior |

Work in kan-server/ and reviewed authoritative contract sources. Current contract remains 1.1.0; do not silently reinterpret container specs as guest VMs or publish a fictional next wire version. D writes Java, **no Go**. New execution requests require coordinated A+B+C+D review before implementation.

## 8. E — CLI learner and user-facing client

| Order | Your work | First output / evidence |
| --- | --- | --- |
| 1 | Agree CLI/shared-client module boundary with A before adding modules or moving signer code | Reviewed imports/layout; cmd/covestack and pkg/kanclient are planned locations |
| 2 | Learn Go/errors/structs and a bounded cobra command/build exercise | Explainable CLI slice and actual three-OS build evidence when produced |
| 3 | Practice JSON, HTTP/context/timeouts/errors against a labelled mock | Table/httptest evidence; mocks never claim infrastructure creation |
| 4 | Connect to the real user API when available | Actual authorization, finite waits, JSON output and useful error evidence |

CLI calls the authenticated API only; no Docker/Incus/k3d or direct DB operations. Mutations use idempotency keys; config is endpoint/profile/key identifier, secrets are protected separately, and no local resource DB exists. E owns CLI/packaging and later a thin provider, never agent/signing/drivers/reconcile. Terraform and advanced integrations are not shortcuts around missing resource APIs.

## 9. Acceptance checklist for the immediate checkpoint

These are required assertions to implement/run, **not checked-off results**.

| Assertion | Owner | Evidence required |
| --- | --- | --- |
| Clean foundation | B; C verifies | Real server startup and reviewed DB prerequisites; build-only evidence labelled separately |
| Protected identity | C+B; A review | Trusted account/host/credential binding; test secrets delivered outside ordinary logs/docs |
| Signed discovery | A+C+B | Assigned signed request returns actual `200` and `[]`; current endpoint/contract matched |
| Negative auth | C with A/B | Invalid/revoked/replayed/wrong-assignment cases fail; record exact actual response without inventing error codes |
| No execution | A+C | Returned resource work stops safely; no claim/ack/backend create/delete is performed |
| Cancellation and failure | A+C | Ctrl+C/cancellation exits cleanly; record actual timeout/server-down/auth-stop behavior and retry gaps |
| Reproducibility | All relevant owners | Revision, sanitized setup, commands/results and reviewer; another teammate can follow it |

Do not enable or test untrusted assignment by accepting arbitrary work for the sake of a demo. Contract change requirements remain explicit blockers, not omitted fields. A failure is useful evidence when its scope and next action are recorded.

## 10. Set up and use the GitHub board

Create **one issue per bounded task**, add that issue to the project and assign the real teammate. Suggested fields: Status, Assignee, Area, Priority, Blocked by and linked PR. These are instructions, not claims that fields/cards already exist remotely.

| Board status | Meaning | Move it when |
| --- | --- | --- |
| Backlog | Needs scope, owner or dependency resolution | Outcome/acceptance and dependencies are understood |
| Todo | Ready task with one owner and a concrete test | Owner starts the task |
| In progress | Active coding/review preparation/verification | A reviewable PR and evidence exist |
| In Review | PR open; reviewers can assess it | Required review/checks pass and PR is merged |
| Done | Accepted task with merged PR and evidence | Keep evidence linked; no automatic service-status promotion |

A blocked task names the dependency and its owner. Use the project's available blocked field/label or record it in the issue body; do not hide an unavailable Java server behind a fake integration result. Do not run two parallel task lists with conflicting statuses.

### Suggested first cards

These are **suggested titles and initial disposition rules**, not created issues or verified board statuses.

| Owner | Issue title | Dependencies / initial disposition |
| --- | --- | --- |
| A | Integrate existing signed poll loop with real Java no-work endpoint | Backlog/blocked until B+C server, credentials and assignment are ready; source review can proceed |
| A | Review VM/claim driver boundary and unknown-outcome safety | Design discussion with B+C+D; no driver execution before protocol/host approval |
| B | Establish bootable authenticated Java foundation and migration prerequisites | Review actual schema history/dependencies first |
| C | Define credential provisioning, trusted assignment and negative auth acceptance | Pair B+A; run real acceptance once server slice exists |
| D | Review discovery fixtures and versioned VM/claim compatibility gaps | Pair B+C+A; preserve current 1.1.0 semantics |
| E | Agree thin CLI/module layout and build mock HTTP learning slice | A review; mocks may proceed once layout is agreed, real CLI waits for user API |

Do not create a new “implement poll loop” issue as if the existing code were absent. If earlier client/loop work has a merged PR and accepted evidence, link that record; the separate real-server integration can remain pending.

### Issue body template

```markdown
## Outcome
One observable behavior.

## Owner and scope
Assignee: actual teammate
Affected files/modules:
Reviewers:
Reference: TEAM-FIRST-CHECKPOINT.md, relevant section

## Dependencies
Blocked by: actual issue link/owner, or none
Missing contract/schema decision:

## Acceptance
- [ ] Exact behavior asserted, including relevant negative/failure case
- [ ] Actual results and evidence class attached
- [ ] Relevant reviewers agree; secrets redacted

## Evidence
Revision:
Commands and actual results:
Evidence class: source review / unit / mock / real-server / runtime
PR:

## Next action
One concrete step or blocker with owner.
```

## 11. Move from this checkpoint to resource execution safely

The next coding slice follows evidence and dependencies, not an old calendar. Existing discovery does not establish a lease/claim protocol. Before execution, review trusted assignment, exact-generation/owner/epoch fencing, progress versus acknowledgement, unknown provider task outcomes, type/version compatibility and capacity/ownership reservations. These are target requirements; add coordinated authoritative schemas/fixtures/generated outputs before claiming new wire support.

| Later boundary | Acceptance to establish |
| --- | --- |
| VM Instances | Actual guest kernel/readiness, owned disk/NIC, safe lifecycle, interrupted create exactly one VM and owned cleanup |
| App Containers | Explicit shared-kernel trust class, restricted privileges/volumes, health and safe retry |
| Kubernetes | Actual cluster/workload, authorized kubeconfig, measured capacity and recovery/cleanup; no inherited academic density promise |
| Networking/IPAM | Parallel reservations have no collisions; allowed traffic passes and forbidden cross-account traffic actually fails |
| PostgreSQL | Real SQL readiness, scoped credentials, persistence, backup and known-row restore |
| RustFS | Real object bytes/checksums, scoped policy/key revocation and cross-account denial, restart persistence, protected backup and verified restoration |
| Environments | Owned children/readiness, safe retry after partial failure and preservation of shared/retained resources |

RustFS is **Planned**, not Implemented. Operator owns service/volumes; Java owns bucket/policy/credential reconcile behind a provider adapter. SeaweedFS/alternatives remain possible. Unsupported ACL authorization and planned bucket logging/ownership controls remain explicit; [provider requirements](docs/platform/provider-evaluation.md#rustfs-compatibility-and-acceptance) define the tested subset and gaps. Recreating bucket records does not restore lost data.

G1 durable authoritative intent; G2 crash-safe reconcile/no duplicates; G3 account and actual isolation; G4 meaningful automated failure acceptance. A service reaches **Acceptance Gates Passed** only with evidence. **Supported** also needs a qualified deployment profile and maintenance policy. No product-wide completion claim or supported production profile is established by this checkpoint. [Readiness](docs/platform/readiness.md) owns the full promotion criteria.

## 12. Daily follow-up and review

1. Read your current issue, baseline and blockers; use [Starter.md](Starter.md) for branch/update/PR commands.
2. Work only in your reviewed slice. If a missing schema/assignment/runtime behavior blocks you, record it with its owner.
3. Run relevant existing checks on the authorized environment and attach actual output scope. This documentation update did not run application/infrastructure tests.
4. Open a small PR; A reviews Go, A+B+C shared-client/VM boundaries, B+C+A schema, D DTO and C security/acceptance as affected. Extended ownership is reviewed, never silently assigned.
5. Update the same issue/card with PR, result and next action. Do not mark a mock/runtime gap Done by changing prose.
6. Update current docs when behavior changes. AI agents also update root checkpoint.md; retain past test evidence as dated history.

Keep kan-server, kan-agent and existing wire/module names until a coordinated change is approved. Never put .env files, signing/master/private keys, credential-bearing kubeconfigs or secret logs into issues/commits. Do not grant tenants host administration. Source/contract behavior, not an architecture diagram, determines actual current functionality.

CoveStack-owned material uses [Apache License 2.0](LICENSE), SPDX **Apache-2.0**; upstream terms remain separate. The [previous first-team checkpoint](docs/archive/pre-vm-redesign-2026-11-04/TEAM-FIRST-CHECKPOINT-original.md) is preserved as historical planning. Use this active file for the next team discussion and Starter.md for onboarding.
