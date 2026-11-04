# CoveStack — current canonical AI-agent handoff

Updated: **4 November 2026**. Branch inspected: **feat/agent-signed-client**, revision **84975bb**. No Git branch/history changes, commit, push or remote synchronization were performed. Read AGENTS.md, this current section and actual Git status before starting.

## Current task, authority and result

Latest user master task: repository-wide documentation refinement for local-first, account/environment-owned real infrastructure. **Markdown only**; no application/config/SQL/migration/API-contract/test/build/CI changes, installs or infrastructure actions. Read-only source inspection and official research authorized. Pre-existing uncommitted changes (including .gitignore) are preserved and separately fingerprinted. User previously confirmed deleting three redundant PDFs; do not restore them.

Canonical product authority: [Product Architecture and Service Catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md). It defines **22 proposed core capabilities and 31 deferred future candidates**, not implemented services, a release plan or an academic amendment. VM Instance is actual guest-kernel compute target; separate App Container is shared-kernel/trusted. Incus/QEMU/KVM is preferred evaluation, no driver exists. Floci is optional external testing integration, not mandatory foundation or rewrite target. Account/environment ownership, bounded dependency/readiness orchestration, capacity scheduling, security and verified recovery are hypotheses to prove. No Labs/assignments/learning-management product and no new release plan.

The object-storage provider has **no selected new default**: evaluate Garage, SeaweedFS and Ceph RGW; legacy MinIO needs current artifact/maintenance/rights review. Secrets should use a reviewed mature provider, not a homemade vault. No root license grant is present; open-source is intended direction. Multi-host is future scope, single-host is not HA, 8 GB is a historical development/lab constraint, not production sizing. No supported production profile.

## Actual implementation and evidence

Go signed client/config/signer/lifecycle helpers, cancellable single-account no-execution poll loop and entry point exist. Returned work stops with ErrWorkAvailable. Timeout/selected HTTP retry exists; general network-failure and real-server behavior are not proven. Java auth/replay/interfaces/helpers and extended V5/V6 SQL exist, but foundational migration chain, runnable scoped APIs, concrete credentials/assignment, VM/container/network/storage execution, full CLI and environment/scheduler behavior remain missing or unverified. Existing tests/reports are historical evidence, not fresh runtime verification. No application/infrastructure tests ran for this documentation task.

Current contract 1.1.0 stays unchanged; conceptual VM/App Container/project/capacity/network fields are not approved wire kinds/fields/states. Claims use exact generation/owner/epoch target fence, progress never acknowledges; stale runtime tasks require outcome inspection. One agent per host with account credential/client contexts is retained pending host-assignment/custody detail review.

## Roles and scope review

A/Anuj owns agent/client/driver/signing/reconcile Go, CI, installer/ITISS and explains code. E thin API-only CLI/JSON/packaging/later provider, never agent implementation. B Java core/all SQL/leases/idempotency/quotas/stacks; C security/all acceptance/Java allocator; D Java API/DTO/schema, no Go. Future scheduler ownership needs explicit team review. Academic CORE-5 categories/tests 1–15, fixed stack and approved cross-cutting obligations remain visible; product VM conversion and wider core require team/mentor approval, not silent scope expansion. Latest user deadline 15 December lacks restated year; historical V5 says 25 January 2027. Do not invent resolution.

Active statuses: Proposed, Planned, In Progress, Partially Implemented, Implemented, Runtime Verified, Acceptance Gates Passed, Supported, Deferred, Blocked. G1 intent, G2 crash-safe reconciliation, G3 ownership/actual isolation and G4 automated failure evidence remain mandatory. Supported additionally needs declared scope, qualified profile and maintenance policy. No gate promotion in this task.

## Verification, files and next action

Current refinement modified 57 existing Markdown files and created 6 Markdown files; exact paths/dispositions and current versus earlier audit are in [documentation audit](docs/documentation-audit.md). Static tables/fences/active links/anchors, catalogue IDs, historical V5/archive bodies and protected-file fingerprints checked. 48 existing missing historical links retained; no active link failures. Mermaid structurally inspected, not visually rendered. Ignore rules unchanged; ignored Markdown needs deliberate maintainer handling if later publishing.

Next action is **team review**, not automatically authorized code: read canonical catalogue, [academic transition](docs/platform/academic-transition.md), [decisions](docs/platform/decisions.md), [provider evaluation](docs/platform/provider-evaluation.md), [capacity](docs/platform/capacity.md), [readiness](docs/platform/readiness.md) and [implementation status](docs/platform/implementation-status.md). B+C+A/D settle protocol, custody, host assignment, VM scope and acceptance mapping; confirm deadline and license. Then separately authorize real-server integration and a capable-host Incus spike. Preserve current restrictions: no Docker/Compose/k3d/MinIO runtime, no manifests, Terraform deferred and unauthenticated GH CLI; VM capability not probed.

## Dated task continuity

- Earlier 4 November redesign: established docs/platform, VM-first target and historical relocation notices; preserved original guide/evidence bodies. Its earlier audit is retained beneath the new audit. No implementation or service promotion.
- Current 4 November refinement: created canonical core/future catalogue and five focused domain documents, strengthened environment/network/provider/scheduling/observability/comparison coverage, removed active release assignments by superseding the old plan, aligned navigation/status and preserved historical bodies. No source or infrastructure work.

## Preserved previous checkpoint history

Recorded prior actions/results below are unchanged. Old architecture, dates and next-task guidance do not override this current summary or user authorization.

<!-- Historical checkpoint body starts here. -->

# CoveStack — AI agent handoff checkpoint

Updated: **4 November 2026**. This is a continuity note, not proof of deployed features. Read actual files, current user instructions and git status before acting. Do not assume the snapshot below remains current.

## 1. Read this first

The user is **Anuj, Person A**, learning Go. Explain code in simple language and small steps. He wants to build and understand the agent himself. The single-account polling loop and agent entry point now exist in internal/agent/loop.go and cmd/kan-agent/main.go. The current task creates a standalone CoveStack project overview and maintains this checkpoint; no further implementation or Floci replacement is authorized. Next integration work requires B+C-coordinated real-server evidence, not recreation of the existing loop.

Read AGENTS.md and latest user instructions. Current rules restrict work to approved documentation, forbid code/SQL/contracts/deploy/script/HTML/ignore edits and commit/push/network/installs unless a later user request specifically authorizes an exception. Earlier explicit client-code authorization does not grant unrestricted future edits. Preserve unrelated changes. AI-generated agent Go is reference only: A must type and explain it. No automatic delegation is needed.

## 2. Project identity and target

**A local cloud control plane with an account-scoped, crash-safe resource model. Docker and k3d are replaceable execution backends.**

| Component            | Responsibility                                                                                                                                   |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| kan-server           | Java 25 / Spring Boot; JdbcClient/plain SQL, PostgreSQL 17, Flyway; authoritative intent, tenancy, assignment and transactions; no JPA/Hibernate |
| kan-agent            | Go; signed outbound API requests, future polling/reconciliation and bounded backend execution; never direct platform DB access                   |
| CLI                  | E's thin Go/cobra authenticated API client; no Docker/k3d calls or resource DB                                                                   |
| Shared client layout | pkg/kanclient is planned, not present; module/signer relocation needs A+B+C review                                                               |

CORE-5: container-backed Droplets (not VMs), required Kubernetes through k3d, Managed PostgreSQL with backup/restore/scoped credentials/private network, MinIO-backed object storage, VPC/IPAM (one private network/account, no custom CIDR).

One fixed stack integrates PostgreSQL + tenant bucket in shared MinIO + app Droplet. It must run a usable app, not just create records. No DAG engine or automatic public route. CLI, security/recovery, optional Terraform for Droplet/Kubernetes, and Ansible Part 1 are cross-cutting scope. No UI during CORE-5.

Kubernetes: one host; max 2 clusters, each 1 server plus max 2 agents. Cap is not a measured RAM guarantee. Only two workload drivers: generic container and k3d; network execution is a container-driver helper. Server Java reconciles bucket/credential policy; agent runs the shared MinIO process/volume. Do not turn buckets into agent work.

Extended track: ASG, load balancing, application publication, metadata, monitoring, snapshots, domains, React read-only UI and Ansible inventory Part 2. A starts extended Go only after CORE-5 gates pass. Optional future work remains separate; no numeric extended-service-count claim. Publication/domain SQL alone does not supply a reconciler, DNS verifier, tunnel or internet access.

Planning began 14 October 2026; first workflow checkpoint 31 October 2026. Feature freeze after 10 January 2027; submission Monday, 25 January 2027. Report unmet gates rather than fabricate completion or change scope.

## 3. People and review boundaries

| Role     | Ownership                                                                                             |
| -------- | ----------------------------------------------------------------------------------------------------- |
| A — Anuj | All agent/client/signing/reconcile/driver Go, network helper, CI, Ansible Part 1 and ITISS evidence   |
| B        | Java core; all SQL/migrations, leases/fencing/idempotency, quotas and stack records                   |
| C        | Auth/security; owner of all acceptance tests from day one; Java IPAM allocator                        |
| D        | Java /v2 endpoints, DTO/schema alignment and Terraform fixtures; no Go                                |
| E        | Thin cobra CLI, JSON, packaging and later thin Terraform provider; no agent/signing/drivers/reconcile |

Any Go needs A review; pkg/kanclient and k3d need A+B+C; schema needs B+C+A, with D reviewing DTOs. Do not silently shift ownership. ITISS permission for agent-code ownership remains a mentor confirmation item; do not invent approval.

## 4. Current checkout snapshot

Observed branch: feat/agent-signed-client. HEAD at inspection: 4ef30f7 (Docs: Add Docs). Earlier local history includes 7dfe0a4 (client functionality) and 8dc78dc (checkpoint docs). No remote fetch or PR inspection performed; do not infer current upstream/merge state.

**Pre-existing modified file: config.json. Do not print its contents, stage it or overwrite it as part of unrelated work.** The new checkpoint.md is the only intended change in this handoff task. Recheck git status before every edit/commit.

| Existing agent file                               | Purpose                                                                                           |
| ------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| kan-agent/go.mod / go.sum                         | Module github.com/ciagotech/kan-agent; go 1.22, toolchain go1.27.2; dependencies/checksums        |
| internal/config/config.go and config_test.go      | KAN\_\* environment loading, origin/rule validation and redacted formatting                       |
| sigv1/rules.go                                    | Loads validated shared signing rules                                                              |
| sigv1/sign.go                                     | Canonical request and HMAC signing                                                                |
| sigv1/request.go                                  | Fresh signed request, nonce/date; caller must refuse redirects                                    |
| sigv1/sign_test.go / negative_test.go             | Shared vectors and invalid/tampered signing cases                                                 |
| internal/lifecycle/catalog_generated.go           | Generated contract catalogue; do not edit manually                                                |
| internal/lifecycle/validate.go / validate_test.go | Kind/state/deletion/transition and observation checks                                             |
| internal/client/client.go / client_test.go        | Single signed work GET, typed common fields, bounded decoding, HTTP errors and cancellation tests |
| cmd/auth-smoke/main.go / main_test.go             | Existing authentication diagnostic, not a production agent loop                                   |

Paths in the table are relative to kan-agent except the module row. The earlier snapshot lacked a loop/entry point; inspection on 4 November 2026 found internal/agent/loop.go, loop_test.go and cmd/kan-agent/main.go. No root pkg/kanclient or cmd/covestack module is established by this snapshot. No functioning execution driver or service gate is established.

### Client behavior already present in source

- New(config.Config) creates a reusable 10-second HTTP client, rejects redirects, validates origin and copies signing context.
- GetWork(context.Context) returns ([]client.Resource, error); performs one GET, no polling/retries/claim/ack/driver calls.
- HTTP 200 with [] returns a non-nil empty slice. Null, malformed/oversized JSON and invalid common/lifecycle fields return errors, not fake empty success.
- Local response bound: 4 MiB; max 1000 work items from current contract. Local bound is not a wire guarantee; review with B+C before release.
- HTTPError exposes StatusCode and RetryAfter without response body. Future loop decides retry policy.
- Resource includes ID, Kind, Generation, DesiredState, State, Deleting, opaque Spec and Labels. Drivers must later validate spec and installed capabilities; labels are not authorization.
- Client has redacted String/GoString. Reuse existing signer, never invent headers.

## 5. Verification: evidence versus pending work

On 28 October 2026 the user reported successful go test ./... for auth-smoke, client, config, lifecycle and sigv1, followed by go vet ./... with no diagnostic shown. Earlier implementation-session tests and vet also passed after local-loopback permission was granted. These are historical observations, not a fresh run on today's tree. No tests were rerun to create this handoff note.

**Real Java-server integration remains pending in the conversation.** No observed startup URL, issued service secret, trusted assignment or live polling/shutdown evidence was supplied. Never claim Gates passed from source/tests alone.

Run appropriate checks when code changes:

```bash
cd kan-agent
go test ./...
go vet ./...
```

No network/install authorization is implied. Use available local toolchain/dependencies; ask through proper tool permissions if sandbox restrictions block required local listener tests. Unit/local test-server success does not replace real Java integration.

## 6. Agreed first communication milestone

Full reference: [Anuj checklist, section 2.3](V5/CoveStack-Anuj-Next-Steps.md#23-first-milestone-decision--single-account-communication-planned).

| Item               | Existing rule or selected starting behavior                                                                                                                                     |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Development URL    | http://localhost:4700 in contract; same-host example, not proof of running server; remote origin needs verified HTTPS                                                           |
| Configuration      | KAN_SERVER_URL, KAN_ACCESS_KEY, KAN_SECRET_KEY, KAN_RULES_PATH; loader does not source .env automatically                                                                       |
| Signing rules path | Default ../contracts/sigv1-rules.json assumes running from kan-agent; set explicit path otherwise                                                                               |
| Credential         | Dedicated service key, one account/host assignment; C provisions securely, B persists and validates trusted assignment                                                          |
| Secret custody     | api_tokens hash-only; signing_credentials encrypted plus key version; master-key delivery/rotation B+C                                                                          |
| Scope              | Account derived from verified credential; one process/host with separate immutable credential/client context per account later                                                  |
| Read-only work     | GET /internal/v1/agent/work; no-work is HTTP 200 []                                                                                                                             |
| Capabilities       | Empty supportedKinds until reviewed tested driver; unexpected nonempty work must not execute or acknowledge                                                                     |
| Heartbeat          | Requires agentVersion, dockerVersion, supportedKinds; agree truthful version source, never fabricate values                                                                     |
| Timing             | 10-second HTTP timeout; 5-second normal polling base with bounded jitter; temporary backoff 1 second doubling to 30-second cap                                                  |
| Errors             | Respect valid 429 Retry-After with cancellable waiting; invalid value uses backoff; 401/403 stops this account loop; other malformed/nonretryable responses are explicit errors |
| Request freshness  | Every attempt freshly signed with new nonce/date; no redirect/credential forwarding                                                                                             |
| Shutdown           | Cancel active HTTP request, timer/backoff and future polling on Ctrl+C/termination                                                                                              |

Timing values are selected initial operational defaults, not new environment variables or protocol constants. Exact jitter bounds, retry classification, Retry-After parsing/wait limits and local response bound still need reviewed code/tests. Do not hide permanent protocol errors in endless retries. No secret values in logs, issue evidence or screenshots.

B+C+A must finalize credential-to-agent lookup/provisioning; existing KeyService exposes account/key/kind but does not itself prove host assignment. Use mocks for isolated development if needed; keep real-server exit criteria pending.

## 7. Historical loop implementation plan — now present; verify integration

The following was the original plan. On 4 November 2026, these files exist and the loop/entry point were read. Do not recreate them. Review remaining integration and recorded verification before another authorized code change.

```text
kan-agent/
  internal/agent/
    loop.go
    loop_test.go
  cmd/kan-agent/
    main.go
```

When the user authorizes coding, review current files/instructions first. Implement a small no-execution loop before drivers:

1. internal/agent/loop.go: a WorkClient interface whose GetWork(context.Context) returns []client.Resource and error. Existing \*client.Client satisfies it; inject a fake client for tests.
2. Build the loop gradually: call once, handle [], wait on timer or context cancellation, repeat. No plain time.Sleep that prevents prompt shutdown.
3. Add bounded temporary-failure backoff/jitter, valid 429 retry guidance, clear permanent errors and 401/403 stop. No automatic resource execution from a nonempty response; fail safely without ack/cleanup.
4. internal/agent/loop_test.go: deterministic fake-client/timing tests for no work, cancellation during request/wait/backoff, authentication stop, retry reset, rate limiting and unexpected work. Avoid long real sleeps.
5. cmd/kan-agent/main.go: config.Load, client.New, signal.NotifyContext for interrupt/termination, then agent.Run using the context. Keep entry point small; choose exact Run signature during implementation, it is not an existing API.
6. Define clean cancellation versus error exit; sanitize logs. Do not add heartbeat until required inputs are truthful and scope is reviewed.
7. Run tests/vet; pair with C+B for a real 200 [] request, server stop/restart recovery, invalid credentials and prompt Ctrl+C. Attach actual evidence.

No container execution, multi-account concurrency or new claim wire fields in this slice. Do not take E's CLI work. Explain the code to Anuj in beginner-friendly steps.

## 8. Protocol and schema blockers before execution

Lifecycle contract is 1.1.0. New messages require coordinated base/schema, fixtures and generated outputs with B+D+A/C; documentation V5 or a future-version note is not an implemented wire release. GET /work stays read-only. Do not invent claim routes/fields or mutate a lease on GET.

Before backend mutation agree signed idempotent claim, renewal, progress/convergence and network-control exchanges. Acknowledgment validates account/assignment, exact generation, owner, epoch and live expiry; clears owner and expiry together, retains epoch. Progress never updates acked_generation. DB fences reports, not Docker/k3d directly: later OS host exclusion/per-resource serialization and settlement of ambiguous backend commands are also required.

The only inspected SQL migrations are V5**create_publications.sql and V6**create_domains.sql. Foundational accounts/resources migrations were not found. They are required dependencies; no clean-database readiness is proven. B must inspect applied history before choosing migration versions; never rewrite an applied migration or renumber blindly. This checkpoint does not authorize SQL edits.

[V5 schema](V5/CoveStack-Database-Schema.md) is reference design with unresolved placeholders, not executable migration-ready SQL. It distinguishes existing extended SQL from stronger target constraints and missing workflow details. PostgreSQL 16/text ULIDs/merged access_keys/in-progress/errored from the simple guide do not override the agreed model and contract.

## 9. Roadmap after the communication milestone

Reviewed claim/default-network protocol → one container Droplet with crash/fencing tests → second isolated account context → VPC/IPAM execution integration → bounded k3d lifecycle → PostgreSQL/MinIO persistence/recovery → fixed stack, installer and ITISS evidence. Default network execution is an early dependency for the first real workload; CLI delivery order does not postpone it.

Production CLI order: droplet → vpc → k8s → db → bucket → stack; prior failure test passes in CI before next production service starts. C owns tests; no changes to test numbering. Do not add features merely to match LocalCloud's roadmap; focus on usable account-isolated environments and explainable recovery.

A service is Supported only at **Gates passed** after G1–G4: durable DB intent/generation; crash-safe no-duplicate reconcile/rebuild; cross-account 404 with relevant protocol/network denial; failure test in CI. Status vocabulary: Planned | In progress | Gates passed | Extended track. Resource/library source presence is not service acceptance. No CORE-5 service is promoted by this checkpoint.

## 10. Team board and references

[TEAM-FIRST-CHECKPOINT.md](TEAM-FIRST-CHECKPOINT.md) covers each person's first tasks and issue templates. Documentation owns agreed scope/exit criteria; GitHub issues/project own live task status/blockers/evidence. Board order: Backlog → Todo → In progress → In Review → Done. Normally one active issue/person. Task Done means reviewed/merged with its evidence, not service Gates passed. Do not update GitHub or send messages without authorization.

| Reference                                                   | Read for                                                         |
| ----------------------------------------------------------- | ---------------------------------------------------------------- |
| [Team onboarding](V5/CoveStack-Team-START-HERE.md)          | First team workflow and baseline                                 |
| [Roles](V5/Team_Roles.md)                                   | Ownership/review boundaries                                      |
| [Anuj next steps](V5/CoveStack-Anuj-Next-Steps.md)          | Immediate milestone and defaults                                 |
| [Setup](V5/CoveStack-Complete-Working-Setup.md)             | Local configuration and actual versus target commands            |
| [Decision log](V5/CoveStack-References-and-Decision-Log.md) | Selected defaults versus open items                              |
| [SRS](V5/CoveStack-SRS-Scope.md)                            | Exactly CORE-5 and cross-cutting scope                           |
| [Acceptance](V5/CoveStack-AntiWrapper-Acceptance.md)        | Tests 1–15, extended E1–E8                                       |
| [Gates](V5/CoveStack-Service-Gates.md)                      | Service promotion evidence                                       |
| [Technical design](V5/DocsV5-DO_Team.md)                    | HLD/LLD and execution safety                                     |
| [Contract](contracts/internal-api.yaml)                     | Current generated wire API; follow source/generator instructions |
| [Lifecycle](contracts/lifecycle-rules.json)                 | Exact state/kind strings                                         |

Older docs/team and simple guides can be historical; do not silently override V5/contracts. The root HANDBOOK.md describes LocalCloud, a comparison/reference project, not proof of CoveStack capabilities. Existing print HTML may lag Markdown; this handoff does not authorize regeneration. The previous website-refactor request targeted a different repository; no website change was made here.

## 11. Handoff maintenance

Before continuing: check git status, applicable instructions, actual code and task authorization. Update this note when a milestone/decision genuinely changes; distinguish inspected source, historically reported tests and freshly executed runtime evidence. Never copy credentials or config.json contents into the note. On finishing, report changed files, checks actually run, remaining blockers and whether anything was committed/pushed. This handoff creation performs no commit, push, network access, install or implementation change.

## 12. Latest task record — 30 October 2026

**Task:** require a single maintained AI-agent handoff after every task. **Progress:** 2 of 2 documentation steps addressed: AGENTS.md rule added; this checkpoint synchronized. No implementation milestone or service gate advanced.

**Approach/changes:** appended mandatory read/update rules to AGENTS.md, selecting this root checkpoint.md as the sole canonical agent handoff. Rules cover objective, approach, changed/inspected files, actual checks, evidence scope, blockers, concrete progress, remaining work and next file/action. Explicit no-edit requests take precedence; GitHub live issue status and contract/source authority remain separate. No competing handoff file was created.

**Verification scope:** documentation checks passed for required rule text, balanced Markdown fences, local file links and before/after fingerprints; only AGENTS.md and checkpoint.md changed during this task. No Go/Java or live-server test is required for this instruction-only change. Do not infer runtime acceptance from documentation validation.

**Preserved work:** config.json was already modified and remains untouched. checkpoint.md was already untracked at task start; no staging/commit/push was authorized or performed. Branch observed at task start: feat/agent-signed-client; no remote lookup performed.

**Remaining/next:** single-account polling loop and startup are still pending. On the user's coding authorization, reread sections 6–8, inspect current files, then start proposed kan-agent/internal/agent/loop.go and loop_test.go before proposed kan-agent/cmd/kan-agent/main.go. B+C real server/credential/assignment and coordinated execution schemas remain dependencies. Keep Anuj's beginner-friendly learning and review requirements.

## 13. Latest task record — 3 November 2026

**Task:** inspect `TEAM-FIRST-CHECKPOINT.md` and explain the next build target for Person A. **Result:** no implementation milestone or service gate advanced.

**Inspected:** `TEAM-FIRST-CHECKPOINT.md` sections 3, 8 and 9.1, plus this checkpoint. The immediate target is the single-account, read-only agent startup and polling slice: proposed `kan-agent/internal/agent/loop.go`, `loop_test.go`, then `kan-agent/cmd/kan-agent/main.go`. The existing signed client is the dependency; drivers, claim/ack, cleanup and multi-account execution remain out of scope for this slice.

**Checks/evidence:** inspected documentation and ran `git status --short --branch`; no Go/Java tests were run because this was an explanation-only task. Existing working-tree changes were preserved; no commit, push, network access or install was performed.

**Next action:** after explicit coding authorization, inspect the current agent files and implement the loop incrementally with cancellation-safe waits, bounded retry/backoff, authentication-stop handling and fake-client tests before adding the small signal-aware command entry point. Real Java-server, credential and assignment evidence remains pending with B+C.

## 14. Latest task record — 3 November 2026

**Task:** explain Person A's responsibilities from `checkpoint.md`. **Result:** no implementation milestone or service gate advanced.

**Summary given:** Anuj owns the Go agent/client/signing/reconcile/driver work. The immediate build is the single-account, read-only polling slice: `kan-agent/internal/agent/loop.go`, deterministic `loop_test.go`, and then the signal-aware `kan-agent/cmd/kan-agent/main.go`. The loop must poll the existing signed client, wait cancellably, apply bounded temporary-error backoff and 429 handling, stop on 401/403, and never execute or acknowledge work.

**Checks/evidence:** reread `checkpoint.md` and checked the branch/status header with `git status --short --branch`; no Go/Java tests or live-server checks were run. Existing working-tree changes were preserved; no commit, push, network access or install was performed.

**Next action:** after explicit coding authorization, recheck current source and instructions, then implement and test the loop before the command entry point. Real Java-server credential/assignment, outage recovery and shutdown evidence remain pending with B+C.

## 15. Latest task record — 3 November 2026

**Task:** inspect V5 for required skills/MCP and prepare a requested installation. **Result:** no skill or MCP installation performed because the request supplied the literal placeholders `SKILL_NAME` and `MCP_SERVER_NAME` rather than package names.

**Findings:** V5 defines a tested toolchain—Go, Java/Maven/Spring, Python, Docker/Compose, k3d/k3s, kubectl, PostgreSQL, Valkey, MinIO/S3, cobra, optional Terraform, Ansible, ansible-lint, ansible-vault, golangci-lint, Git/GitHub and curl/API testing. It does not select an installable skill or MCP package. Relevant native workflow skills are already available for setup, feature work, build, testing, debugging, review, handoff, minimal coding and skill discovery; MCP use is supportive rather than a V5 product dependency.

**Checks/evidence:** read the V5 onboarding, setup, Go handover, tools, Terraform/Ansible, UI/CLI and PDF documents; checked the branch/status header. No network, package install, global configuration backup/edit, privileged command or project file change was performed.

**Next action:** obtain the exact skill and MCP server names (or approved package/repository identifiers), then verify official sources, dependencies, environment/API-key requirements and existing installations before any global change. Never overwrite existing skills or MCP configuration.

## 16. Latest task record — 3 November 2026

**Task:** check for and install PostgreSQL and Valkey MCP servers. **Result:** trusted packages were verified and their harmless help commands ran successfully; active MCP registration remains pending.

**Sources and checks:** PostgreSQL uses Microsoft's `@microsoft/postgres-mcp` from `microsoft/postgres-mcp`, requiring Node 22+ and supporting `npx -y @microsoft/postgres-mcp run`; its profile must be explicitly read-only for autonomous use. Valkey uses AWS Labs' `awslabs.valkey-mcp-server` from `awslabs/mcp`, requiring Python 3.10+/uv and supporting `uvx awslabs.valkey-mcp-server@latest --readonly`; no API key is needed for basic local Valkey access, while optional embeddings may require AWS/OpenAI/Ollama configuration.

**Actual result:** Node 22.23.3, uv 0.13.0 and Python 3.11.2 are available. `npx --yes @microsoft/postgres-mcp --help` and `uvx awslabs.valkey-mcp-server@latest --help` completed successfully and cached dependencies outside the repository. Registry checks still report `postgres` and `valkey` as not registered. The requested `/home/anujtiwari/.config/omnirush/opencode.json` is absent in this runtime, so no backup or guessed global configuration was created.

**Next action:** provide or expose the actual OpenCode global configuration path/runtime, then back it up and add only read-only PostgreSQL and Valkey entries with explicit local endpoints and no secrets. Recheck registry tools after reload; do not configure database credentials or write/admin capabilities without separate authorization.

## 17. Latest task record — 3 November 2026

**Task:** implement Person A's single-account agent polling slice after rereading root instructions. **Authorization:** the user explicitly requested coding and selected Ponytail mode; the earlier documentation-only restriction was superseded for this coding batch.

**Changed:** added `kan-agent/internal/agent/loop.go`, `kan-agent/internal/agent/loop_test.go` and `kan-agent/cmd/kan-agent/main.go`. The loop reuses the existing signed `GetWork` client, polls one account, waits with cancellable timers, applies bounded jitter/backoff, honors bounded `429 Retry-After`, stops on `401/403` and other permanent errors, and returns safely without execution/acknowledgment when work is present. The command loads existing `KAN_*` configuration and handles interrupt/SIGTERM. No drivers, heartbeat, claim protocol, database access or CLI work was added.

**Verification:** `go test ./...` and `gofmt` were attempted from `kan-agent` but both failed immediately with `command not found` because the container has no Go toolchain. A Python static check passed for final newlines and zero trailing-whitespace lines in all three files. No live Java-server integration was run. Existing working-tree changes were preserved; no commit or push was performed.

**Remaining risk/dependency:** Go formatting, compilation, unit tests and vet remain unverified until a Go toolchain is available. Real server credential/assignment, outage recovery and shutdown evidence remain pending with B+C. Next action: run `gofmt`, `go test ./...` and `go vet ./...` in an environment with the approved Go toolchain, then review the loop with A/B/C boundaries before any execution driver.

## 18. Latest task record — 3 November 2026

**Task:** reread root instructions and synchronize the canonical handoff after the Person A polling-loop task. **Result:** `checkpoint.md` was updated in place; no code or contract changes were made in this maintenance step.

**Evidence scope:** reread `AGENTS.md`, reread the current checkpoint tail, and checked `git status --short --branch`. The working tree remains on `feat/agent-signed-client`; the three new agent files are untracked and existing unrelated modifications remain preserved. No tests, commit, push, network access, install or privileged command was performed for this checkpoint-only step.

**Current verified state:** this record's missing-Go-toolchain statement was superseded by the development-environment setup recorded below. No service gate has advanced and no live Java-server evidence exists.

**Next action:** use the current toolchain evidence below, then address the remaining lint findings and Docker-daemon dependency before any driver or execution work. Keep B+C dependencies for credential assignment and real-server recovery evidence explicit.

## 19. Latest task record — 3 November 2026

**Task:** audit and set up the Covestack development/testing toolchain inside the actual OmniRush sandbox. **Authorization:** the user explicitly requested installation, local service validation, project checks and documentation. Root `AGENTS.md` was reread; documentation-only repository rules were preserved and no application code, contracts, deployment files or ignore files were changed by this setup task.

**Environment:** Debian GNU/Linux 12, x86_64, Linux 7.2.11-200.fc44.x86_64; working directory `/run/media/anujtiwari/DevMount/Project-CDAC/covestack-cloud`; outbound HTTPS passed; overlay had 163–164 GB free at checks. Package-manager wrappers and sandbox paths were recorded in `docs/development-environment.md` and `docs/development-toolchain-status.md`.

**Installed and verified:** Go 1.27.3/gofmt, Temurin JDK 25.0.5, Maven 3.8.7, Docker CLI 20.10.24, Docker Compose 5.7.0, kubectl 1.37.2, k3d 5.9.0, PostgreSQL 17.11/client tools, Valkey 9.1.3/server/CLI, MinIO `mc`, Ansible 7.7.0/core 2.14.18, ansible-lint 6.13.1, ansible-vault, golangci-lint 2.15.0, Git 2.39.5, GitHub CLI 2.23.0, curl, pip and venv. Terraform was intentionally deferred as optional until E begins provider work; no separate host k3s binary was needed because k3d reports its bundled k3s version.

**Checks and local services:** Go race tests passed; `gofmt -l` reported no differences; `go vet ./...` produced no diagnostic; Maven test and verify passed with 15 Java tests and zero failures/errors; contract scripts and `V5/validate_v5.py` passed; Python compileall passed using an external bytecode cache; Compose configuration parsing passed. A disposable PostgreSQL probe passed on sandbox port 5433 and Valkey `PING` plus `SET/GET/DEL` passed on 127.0.0.1:6379. No production endpoint or credential was used.

**Blocked or incomplete:** Docker CLI and Compose are installed but `/var/run/docker.sock` is unavailable, so Compose startup/shutdown, k3d cluster validation and container-backed MinIO/S3 tests are blocked. No MinIO endpoint is running. No Kubernetes manifest with an `apiVersion:` field was found; the YAML inventory contains the OpenAPI contract and Docker Compose configuration. `gh auth status` reports no login. At the time of this historical record, `golangci-lint run` reported eight findings; those findings were later fixed in record 21. No host daemon, socket permission or paid cloud resource was changed.

**Files changed:** added `docs/development-environment.md` and `docs/development-toolchain-status.md`; existing code/config changes and untracked agent files were preserved. Branch remains `feat/agent-signed-client`, observed HEAD `caffee3`; no commit or push was performed.

**Next action:** follow the permanent environment disposition recorded below for excluded infrastructure. Go lint remediation is tracked separately in record 21 and is not an accepted limitation.

## 20. Latest task record — 3 November 2026

**Task:** record the user's permanent disposition for the remaining setup findings. **Result:** documentation only; no application, lint configuration, ignore file, daemon/socket, endpoint, Terraform installation or GitHub credential change was made.

**Accepted permanent limitations:** leave Docker daemon/socket unavailable; leave Compose runtime, k3d clusters, MinIO containers and container-backed acceptance unavailable; keep the repository without Kubernetes manifests; do not configure MinIO/S3; keep Terraform and `terraform-plugin-framework` deferred; keep GitHub CLI unauthenticated. These are accepted environment decisions, not passing evidence or service-gate results. The eight Go lint findings are not part of these restrictions and were fixed in record 21.

**Files updated:** `docs/development-environment.md`, `docs/development-toolchain-status.md` and this canonical `checkpoint.md`. Existing code/config changes were preserved. Branch remains `feat/agent-signed-client`, HEAD remains `caffee3`; no commit or push was performed.

**Checks:** reread `AGENTS.md`, reread the current checkpoint/status documentation, and checked the working-tree status. No tests or installation commands were rerun because this was a documentation-only disposition update. No secrets or credentials were recorded.

**Next action:** treat the excluded infrastructure decisions as the terminal setup baseline and continue only with work that does not require those capabilities. Go lint status is governed by record 21.

## 21. Latest task record — 4 November 2026

**Task:** correct the prior environment disposition and remediate all eight Go lint findings. **Authorization:** the user explicitly requested code-level fixes, verification, and documentation updates, while reaffirming the permanent infrastructure restrictions. `AGENTS.md` was reread before changes.

**Original findings and corrections:**

1. `cmd/auth-smoke/main.go:42:22`, `errcheck`: checked `response.Body.Close` and return a contextual close error.
2. `internal/client/client.go:113:23`, `errcheck`: checked the deferred response-body close using named results and surface a close failure when no earlier error exists.
3. `internal/client/client_test.go:54:13`, `errcheck`: checked the `fmt.Fprint` result and report writer errors through `testing.T`.
4. `internal/client/client_test.go:87:107`, `errcheck`: expanded the test handler and checked its `fmt.Fprint` result.
5. `internal/client/client_test.go:112:15`, `errcheck`: checked the response write and omitted a body for HTTP 204, which does not permit one.
6. `internal/agent/loop.go:136:65`, `staticcheck SA1019`: removed deprecated `net.Error.Temporary`; retry classification now uses the supported portable timeout signal.
7. `internal/client/client.go:82:28`, `staticcheck QF1001`: simplified the equivalent HTTPS/local-HTTP condition using De Morgan's law.
8. `internal/config/config.go:68:28`, `staticcheck QF1001`: applied the same equivalent condition simplification.

**Verification:** from `kan-agent`, `gofmt -w internal/agent/loop.go internal/agent/loop_test.go cmd/kan-agent/main.go` passed; the final rerun of `go test -race ./...` passed for all packages; `go vet ./...` passed with no diagnostic; and `golangci-lint run` passed with `0 issues`. `gofmt -l .` still lists 13 pre-existing tracked files with CRLF line endings. Those unrelated line-ending changes were preserved rather than rewritten.

**Files changed:** `kan-agent/cmd/auth-smoke/main.go`, `kan-agent/internal/client/client.go`, `kan-agent/internal/client/client_test.go`, `kan-agent/internal/config/config.go`, `kan-agent/internal/agent/loop.go`, `docs/development-environment.md`, `docs/development-toolchain-status.md`, and this `checkpoint.md`. Documentation headers were refreshed to 4 November 2026. No API contract, security behavior, lint configuration, infrastructure restriction, credential or unrelated working-tree change was altered.

**Remaining issues:** only the permanently excluded environment capabilities remain unavailable. The Go lint gate is clean. No commit or push was performed; branch remains `feat/agent-signed-client`.

**Next action:** in the next authorized agent batch, review the five corrected Go files against the preserved working-tree changes, then collect B/C-coordinated live Java-server evidence for `200 []`, outage recovery, invalid-auth stop, and Ctrl+C cancellation using non-production credentials. Do not attempt Docker, Compose runtime, k3d/k3s, Kubernetes, MinIO/S3, Terraform, GitHub authentication or container-backed acceptance.

## Floci comparison task — 4 November 2026

**Objective:** explain Floci versus CoveStack using current primary documentation and local source. **Approach:** inspected floci-io/floci README at https://github.com/floci-io/floci and local loop/main, SRS, AGENTS.md and latest checkpoint records. This was a documentation-level comparison, not a full source audit or runtime benchmark of Floci. No Floci installation or workload was run.

**Findings:** Floci documents AWS-shaped APIs for local development/testing/CI, existing SDK/CLI/IaC integration, configurable persistence, real Docker compute/databases and k3s-backed EKS. Do not claim these mechanisms are unique to CoveStack, or assume Floci lacks identity/isolation/recovery without a dedicated audit. CoveStack targets the smaller CORE-5 account-scoped relational control plane, split Java/Go execution, explainable recovery and usable fixed-stack lab. These are target guarantees until evidenced, not proven superiority.

**Current source correction:** loop.go and cmd/kan-agent/main.go exist; the loop polls without execution, handles errors/waits and stops on unexecutable work. Updated stale opening/section 7 statements accordingly. Previously recorded race/vet/lint passes remain historical; no tests or live server checks were rerun for this comparison. No service gate advanced.

**Changes/checks:** checkpoint.md only. Read-only git status showed feat/agent-signed-client with a clean working tree before this update; no remote sync checked. Local links/fences and file fingerprints checked during task. No code, SQL, contracts or other documentation changed; no commit/push/install.

**Next:** when explicitly authorized, collect B+C real-server evidence for 200 [], outage/recovery, invalid-auth stop and Ctrl+C against existing loop. Respect environment restrictions in the latest environment documentation; no Docker/container-backed acceptance or new product scope is authorized by the comparison.

## Floci ITISS clarification — 4 November 2026

**Task/approach:** explain IT infrastructure/security learning relevance using the Floci upstream README (https://github.com/floci-io/floci); no installation or runtime audit. Its documented IAM/STS, secrets/KMS, CloudTrail/CloudWatch and actual compute/DB runtimes provide potential lab exercises, not proof of production-equivalent controls or mentor acceptance. API resource configuration is not proof of packet filtering or isolation; verify actual permitted/denied behavior for a chosen exercise.

**Result/progress:** research answer prepared; no implementation or gate changed. Relevant learning can include authorization tests, protected secrets, audit investigation, service operation and recovery. CoveStack's ITISS evidence must demonstrate its own implemented controls; security relevance is not exclusive to either project. No claim of an official C-DAC marks allocation or approval.

**Changes/verification:** only checkpoint.md updated for handoff; inspected git status showed existing checkpoint changes on feat/agent-signed-client, preserved. No code/tests/runtime executed, no commit/push/install. Next implementation action remains authorized B+C real-server integration of the existing agent loop, under current environment restrictions.

## ITISS mechanism comparison clarification — 4 November 2026

**Request/result:** compare how IT infrastructure/security is performed in Floci versus CoveStack, rather than generic security-topic lists. Read upstream docs/services/iam.md and ec2.md and local V5 security rules/Java auth-filter references. Floci documents optional IAM enforcement (default permissive, test-key bypass) and opt-in network-namespace/nftables security-group filtering; legacy disabled-enforcement publication does not enforce source CIDRs. Do not generalize every Floci network record into actual routing or assume default security enforcement.

**CoveStack distinction:** target mandatory verified credential/account/assignment, tenant predicates, HMAC+nonce replay checks, private network helper, scoped secrets, persisted audit and crash-safe reconcile/restore evidence. Signing/filter/client/loop source exists but full identity provisioning, network/runtime/restore and live integration remain pending. Neither stronger security nor full runtime enforcement was benchmarked/proven by this research.

**Evidence/changes:** documentation-level primary-source comparison only; no runtime/tests/installation/commit/push. Only this existing checkpoint updated, preserving prior handoff entries. Current implementation next step unchanged: authorized B+C real-server tests of the existing agent loop. No new implementation authorized.

## Standalone project overview task — 4 November 2026

**Request/authorization:** one root Markdown document explaining CoveStack, use cases, workflow and infrastructure; ordinary checkpoint maintenance. **Approach:** inspect actual agent entry/loop/client paths, Java source inventory, deploy Compose and V5 scope/HLD/stack. Write COVESTACK-PROJECT-OVERVIEW.md as a readable target/current-state overview with architecture diagram, infrastructure bindings, control/data paths, CORE-5 and extended boundaries, recovery/security, roles and evidence requirements.

**Result:** overview created, no implementation changed. It explicitly records latest user-stated deadline 15 December versus old V5 25 January 2027; year/revised scope not confirmed and no replacement schedule invented. No Floci fork or lab-assessment pivot was approved; they remain discussions. Runtime capabilities/gates are not inferred from source.

**Files/checks:** overview and checkpoint.md only; Markdown table widths, fences, local file links and file fingerprints passed. No Go/Java test rerun or live backend validation for this docs task. Pre-existing checkpoint changes preserved; inspected branch feat/agent-signed-client. No commit/push/install or infrastructure operations.

**Remaining/next:** user should review the overview and decide any revised submission scope with mentor before changing CORE-5. Technical integration next step remains B+C real-server credentials/assignment and no-execution agent tests, respecting environment limits; no new coding task is authorized by this document.
