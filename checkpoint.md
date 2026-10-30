# CoveStack — AI agent handoff checkpoint

Updated: **30 October 2026**. This is a continuity note, not proof of deployed features. Read actual files, current user instructions and git status before acting. Do not assume the snapshot below remains current.

## 1. Read this first

The user is **Anuj, Person A**, learning Go. Explain code in simple language and small steps. He wants to build and understand the agent himself. His latest coding direction is the **single-account agent startup and polling loop**. That code has not been written in the inspected tree. The current handoff-maintenance task authorizes AGENTS.md and checkpoint.md only; it is not permission to implement the next coding task automatically.

Read AGENTS.md and latest user instructions. Current rules restrict work to approved documentation, forbid code/SQL/contracts/deploy/script/HTML/ignore edits and commit/push/network/installs unless a later user request specifically authorizes an exception. Earlier explicit client-code authorization does not grant unrestricted future edits. Preserve unrelated changes. AI-generated agent Go is reference only: A must type and explain it. No automatic delegation is needed.

## 2. Project identity and target

**A local cloud control plane with an account-scoped, crash-safe resource model. Docker and k3d are replaceable execution backends.**

| Component | Responsibility |
| --- | --- |
| kan-server | Java 25 / Spring Boot; JdbcClient/plain SQL, PostgreSQL 17, Flyway; authoritative intent, tenancy, assignment and transactions; no JPA/Hibernate |
| kan-agent | Go; signed outbound API requests, future polling/reconciliation and bounded backend execution; never direct platform DB access |
| CLI | E's thin Go/cobra authenticated API client; no Docker/k3d calls or resource DB |
| Shared client layout | pkg/kanclient is planned, not present; module/signer relocation needs A+B+C review |

CORE-5: container-backed Droplets (not VMs), required Kubernetes through k3d, Managed PostgreSQL with backup/restore/scoped credentials/private network, MinIO-backed object storage, VPC/IPAM (one private network/account, no custom CIDR).

One fixed stack integrates PostgreSQL + tenant bucket in shared MinIO + app Droplet. It must run a usable app, not just create records. No DAG engine or automatic public route. CLI, security/recovery, optional Terraform for Droplet/Kubernetes, and Ansible Part 1 are cross-cutting scope. No UI during CORE-5.

Kubernetes: one host; max 2 clusters, each 1 server plus max 2 agents. Cap is not a measured RAM guarantee. Only two workload drivers: generic container and k3d; network execution is a container-driver helper. Server Java reconciles bucket/credential policy; agent runs the shared MinIO process/volume. Do not turn buckets into agent work.

Extended track: ASG, load balancing, application publication, metadata, monitoring, snapshots, domains, React read-only UI and Ansible inventory Part 2. A starts extended Go only after CORE-5 gates pass. Optional future work remains separate; no numeric extended-service-count claim. Publication/domain SQL alone does not supply a reconciler, DNS verifier, tunnel or internet access.

Planning began 14 October 2026; first workflow checkpoint 31 October 2026. Feature freeze after 10 January 2027; submission Monday, 25 January 2027. Report unmet gates rather than fabricate completion or change scope.

## 3. People and review boundaries

| Role | Ownership |
| --- | --- |
| A — Anuj | All agent/client/signing/reconcile/driver Go, network helper, CI, Ansible Part 1 and ITISS evidence |
| B | Java core; all SQL/migrations, leases/fencing/idempotency, quotas and stack records |
| C | Auth/security; owner of all acceptance tests from day one; Java IPAM allocator |
| D | Java /v2 endpoints, DTO/schema alignment and Terraform fixtures; no Go |
| E | Thin cobra CLI, JSON, packaging and later thin Terraform provider; no agent/signing/drivers/reconcile |

Any Go needs A review; pkg/kanclient and k3d need A+B+C; schema needs B+C+A, with D reviewing DTOs. Do not silently shift ownership. ITISS permission for agent-code ownership remains a mentor confirmation item; do not invent approval.

## 4. Current checkout snapshot

Observed branch: feat/agent-signed-client. HEAD at inspection: 4ef30f7 (Docs: Add Docs). Earlier local history includes 7dfe0a4 (client functionality) and 8dc78dc (checkpoint docs). No remote fetch or PR inspection performed; do not infer current upstream/merge state.

**Pre-existing modified file: config.json. Do not print its contents, stage it or overwrite it as part of unrelated work.** The new checkpoint.md is the only intended change in this handoff task. Recheck git status before every edit/commit.

| Existing agent file | Purpose |
| --- | --- |
| kan-agent/go.mod / go.sum | Module github.com/ciagotech/kan-agent; go 1.22, toolchain go1.27.2; dependencies/checksums |
| internal/config/config.go and config_test.go | KAN_* environment loading, origin/rule validation and redacted formatting |
| sigv1/rules.go | Loads validated shared signing rules |
| sigv1/sign.go | Canonical request and HMAC signing |
| sigv1/request.go | Fresh signed request, nonce/date; caller must refuse redirects |
| sigv1/sign_test.go / negative_test.go | Shared vectors and invalid/tampered signing cases |
| internal/lifecycle/catalog_generated.go | Generated contract catalogue; do not edit manually |
| internal/lifecycle/validate.go / validate_test.go | Kind/state/deletion/transition and observation checks |
| internal/client/client.go / client_test.go | Single signed work GET, typed common fields, bounded decoding, HTTP errors and cancellation tests |
| cmd/auth-smoke/main.go / main_test.go | Existing authentication diagnostic, not a production agent loop |

Paths in the table are relative to kan-agent except the module row. No internal/agent loop or cmd/kan-agent entry point was found. No root pkg/kanclient or cmd/covestack module is established by this snapshot. No functioning execution driver or service gate is established.

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

| Item | Existing rule or selected starting behavior |
| --- | --- |
| Development URL | http://localhost:4700 in contract; same-host example, not proof of running server; remote origin needs verified HTTPS |
| Configuration | KAN_SERVER_URL, KAN_ACCESS_KEY, KAN_SECRET_KEY, KAN_RULES_PATH; loader does not source .env automatically |
| Signing rules path | Default ../contracts/sigv1-rules.json assumes running from kan-agent; set explicit path otherwise |
| Credential | Dedicated service key, one account/host assignment; C provisions securely, B persists and validates trusted assignment |
| Secret custody | api_tokens hash-only; signing_credentials encrypted plus key version; master-key delivery/rotation B+C |
| Scope | Account derived from verified credential; one process/host with separate immutable credential/client context per account later |
| Read-only work | GET /internal/v1/agent/work; no-work is HTTP 200 [] |
| Capabilities | Empty supportedKinds until reviewed tested driver; unexpected nonempty work must not execute or acknowledge |
| Heartbeat | Requires agentVersion, dockerVersion, supportedKinds; agree truthful version source, never fabricate values |
| Timing | 10-second HTTP timeout; 5-second normal polling base with bounded jitter; temporary backoff 1 second doubling to 30-second cap |
| Errors | Respect valid 429 Retry-After with cancellable waiting; invalid value uses backoff; 401/403 stops this account loop; other malformed/nonretryable responses are explicit errors |
| Request freshness | Every attempt freshly signed with new nonce/date; no redirect/credential forwarding |
| Shutdown | Cancel active HTTP request, timer/backoff and future polling on Ctrl+C/termination |

Timing values are selected initial operational defaults, not new environment variables or protocol constants. Exact jitter bounds, retry classification, Retry-After parsing/wait limits and local response bound still need reviewed code/tests. Do not hide permanent protocol errors in endless retries. No secret values in logs, issue evidence or screenshots.

B+C+A must finalize credential-to-agent lookup/provisioning; existing KeyService exposes account/key/kind but does not itself prove host assignment. Use mocks for isolated development if needed; keep real-server exit criteria pending.

## 7. Next coding task — where to start

The user asked where to write the startup/poll loop. Suggested files below **do not exist in the inspected tree**. They are proposed locations inside the existing agent module, not a new module or shared-client move.

```text
kan-agent/
  internal/agent/
    loop.go
    loop_test.go
  cmd/kan-agent/
    main.go
```

When the user authorizes coding, review current files/instructions first. Implement a small no-execution loop before drivers:

1. internal/agent/loop.go: a WorkClient interface whose GetWork(context.Context) returns []client.Resource and error. Existing *client.Client satisfies it; inject a fake client for tests.
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

The only inspected SQL migrations are V5__create_publications.sql and V6__create_domains.sql. Foundational accounts/resources migrations were not found. They are required dependencies; no clean-database readiness is proven. B must inspect applied history before choosing migration versions; never rewrite an applied migration or renumber blindly. This checkpoint does not authorize SQL edits.

[V5 schema](V5/CoveStack-Database-Schema.md) is reference design with unresolved placeholders, not executable migration-ready SQL. It distinguishes existing extended SQL from stronger target constraints and missing workflow details. PostgreSQL 16/text ULIDs/merged access_keys/in-progress/errored from the simple guide do not override the agreed model and contract.

## 9. Roadmap after the communication milestone

Reviewed claim/default-network protocol → one container Droplet with crash/fencing tests → second isolated account context → VPC/IPAM execution integration → bounded k3d lifecycle → PostgreSQL/MinIO persistence/recovery → fixed stack, installer and ITISS evidence. Default network execution is an early dependency for the first real workload; CLI delivery order does not postpone it.

Production CLI order: droplet → vpc → k8s → db → bucket → stack; prior failure test passes in CI before next production service starts. C owns tests; no changes to test numbering. Do not add features merely to match LocalCloud's roadmap; focus on usable account-isolated environments and explainable recovery.

A service is Supported only at **Gates passed** after G1–G4: durable DB intent/generation; crash-safe no-duplicate reconcile/rebuild; cross-account 404 with relevant protocol/network denial; failure test in CI. Status vocabulary: Planned | In progress | Gates passed | Extended track. Resource/library source presence is not service acceptance. No CORE-5 service is promoted by this checkpoint.

## 10. Team board and references

[TEAM-FIRST-CHECKPOINT.md](TEAM-FIRST-CHECKPOINT.md) covers each person's first tasks and issue templates. Documentation owns agreed scope/exit criteria; GitHub issues/project own live task status/blockers/evidence. Board order: Backlog → Todo → In progress → In Review → Done. Normally one active issue/person. Task Done means reviewed/merged with its evidence, not service Gates passed. Do not update GitHub or send messages without authorization.

| Reference | Read for |
| --- | --- |
| [Team onboarding](V5/CoveStack-Team-START-HERE.md) | First team workflow and baseline |
| [Roles](V5/Team_Roles.md) | Ownership/review boundaries |
| [Anuj next steps](V5/CoveStack-Anuj-Next-Steps.md) | Immediate milestone and defaults |
| [Setup](V5/CoveStack-Complete-Working-Setup.md) | Local configuration and actual versus target commands |
| [Decision log](V5/CoveStack-References-and-Decision-Log.md) | Selected defaults versus open items |
| [SRS](V5/CoveStack-SRS-Scope.md) | Exactly CORE-5 and cross-cutting scope |
| [Acceptance](V5/CoveStack-AntiWrapper-Acceptance.md) | Tests 1–15, extended E1–E8 |
| [Gates](V5/CoveStack-Service-Gates.md) | Service promotion evidence |
| [Technical design](V5/DocsV5-DO_Team.md) | HLD/LLD and execution safety |
| [Contract](contracts/internal-api.yaml) | Current generated wire API; follow source/generator instructions |
| [Lifecycle](contracts/lifecycle-rules.json) | Exact state/kind strings |

Older docs/team and simple guides can be historical; do not silently override V5/contracts. The root HANDBOOK.md describes LocalCloud, a comparison/reference project, not proof of CoveStack capabilities. Existing print HTML may lag Markdown; this handoff does not authorize regeneration. The previous website-refactor request targeted a different repository; no website change was made here.

## 11. Handoff maintenance

Before continuing: check git status, applicable instructions, actual code and task authorization. Update this note when a milestone/decision genuinely changes; distinguish inspected source, historically reported tests and freshly executed runtime evidence. Never copy credentials or config.json contents into the note. On finishing, report changed files, checks actually run, remaining blockers and whether anything was committed/pushed. This handoff creation performs no commit, push, network access, install or implementation change.

## 12. Latest task record — 30 October 2026

**Task:** require a single maintained AI-agent handoff after every task. **Progress:** 2 of 2 documentation steps addressed: AGENTS.md rule added; this checkpoint synchronized. No implementation milestone or service gate advanced.

**Approach/changes:** appended mandatory read/update rules to AGENTS.md, selecting this root checkpoint.md as the sole canonical agent handoff. Rules cover objective, approach, changed/inspected files, actual checks, evidence scope, blockers, concrete progress, remaining work and next file/action. Explicit no-edit requests take precedence; GitHub live issue status and contract/source authority remain separate. No competing handoff file was created.

**Verification scope:** documentation checks passed for required rule text, balanced Markdown fences, local file links and before/after fingerprints; only AGENTS.md and checkpoint.md changed during this task. No Go/Java or live-server test is required for this instruction-only change. Do not infer runtime acceptance from documentation validation.

**Preserved work:** config.json was already modified and remains untouched. checkpoint.md was already untracked at task start; no staging/commit/push was authorized or performed. Branch observed at task start: feat/agent-signed-client; no remote lookup performed.

**Remaining/next:** single-account polling loop and startup are still pending. On the user's coding authorization, reread sections 6–8, inspect current files, then start proposed kan-agent/internal/agent/loop.go and loop_test.go before proposed kan-agent/cmd/kan-agent/main.go. B+C real server/credential/assignment and coordinated execution schemas remain dependencies. Keep Anuj's beginner-friendly learning and review requirements.
