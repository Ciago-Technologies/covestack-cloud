# CoveStack — teammate start guide

Updated: **4 November 2026**. For **A/Anuj, B, C, D and E**. Read this after cloning, then work on one small reviewed task at a time. Use [the team checkpoint and board guide](TEAM-FIRST-CHECKPOINT.md) for shared acceptance, dependencies and suggested issue cards.

**CoveStack is Experimental.** This guide explains how to contribute to the project; it is not a full-product installation guide. Architecture describes what we intend to build. Source and actual test evidence establish what currently exists.

## 1. Understand what we are building

CoveStack is a local-first infrastructure control plane: users request account-owned infrastructure on machines they control, and CoveStack should create, observe, isolate and recover it through an API and CLI.

| Part | Current direction |
| --- | --- |
| Compute Instance | A real VM with its own guest kernel; Incus/QEMU/KVM is the preferred backend to evaluate; driver is Planned |
| App Container | Separate shared-kernel resource for trusted applications; never described as a VM |
| Kubernetes | k3d/k3s candidate; host capacity and supported limits must be measured |
| Networks, disks and databases | Owned resources with capacity, access controls, durable intent and verified recovery |
| Object storage | RustFS preferred planned provider; operator manages the service/volumes; Java reconciles account-owned buckets, policies and scoped credentials |
| Control plane | Java/Spring Boot, JdbcClient/plain SQL and PostgreSQL; no JPA/Hibernate |
| Agent | Go; signed outbound communication and future trusted runtime execution |
| CLI | Go/cobra thin authenticated API client; does not call Docker/Incus/k3d or access the platform database |

The [product catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) defines core and future scope. Historical academic CORE-5, dates, template counts and cluster caps do not limit the product. They are not current delivery promises.

## 2. Know what exists before starting

| Present in source | Still pending |
| --- | --- |
| Go config loader, signer and signed HTTP client | Full credential provisioning and trusted account/host assignment |
| Single-account cancellable signed polling and agent entry point | Claims, runtime drivers, resource execution and reconciliation |
| Java auth/replay components and lifecycle helpers | Runnable scoped resource APIs and foundational DB migration chain |
| Unit-test sources and dated test reports | Real VM/RustFS provisioning, isolation, restore and supported deployment evidence |

**The agent deliberately stops when work arrives.** It does not claim, execute, acknowledge or delete the returned work. Do not rebuild the existing loop, call it a runtime driver, or treat a mock/unit test as successful resource provisioning.

Use the [implementation ledger](docs/platform/implementation-status.md) for evidence. Use [updated V5 guidance](V5/README.md) with the canonical product docs; do not copy runnable-service assumptions from archived V5 bodies or old diagrams.

## 3. Read these files in order

| Order | Read | What you should understand |
| --- | --- | --- |
| 1 | [README](README.md) | Product, actual status and limits |
| 2 | [Product SRS](docs/product/CoveStack-Product-SRS.md) and [catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) | Current requirements and core/future classification |
| 3 | [Architecture](docs/platform/architecture.md) and [resource model](docs/platform/resource-model.md) | Java/Go/backend boundaries; VM versus container |
| 4 | [Contribution guide](docs/platform/contributing.md) | Roles, review responsibilities and working process |
| 5 | [Implementation status](docs/platform/implementation-status.md) and [decisions](docs/platform/decisions.md) | Existing code, blockers and missing decisions |
| 6 | [API/CLI boundary](docs/platform/api-and-cli.md) and [contracts README](contracts/README.md) | Actual contract 1.1.0; proposed concepts are not current wire fields |
| 7 | [Security](docs/platform/security.md), [capacity](docs/platform/capacity.md) and [readiness](docs/platform/readiness.md) | Required isolation, admission and failure evidence |
| 8 | Your role section below | Your first discussion, files and reviewable output |

For storage work also read [RustFS requirements](docs/platform/provider-evaluation.md#rustfs-compatibility-and-acceptance); for combined applications read [environments](docs/platform/environments.md). AI assistants must read [AGENTS.md](AGENTS.md) and maintain [checkpoint.md](checkpoint.md).

## 4. Prepare your development tools

| Who | Tools / source of version truth | Safe first check |
| --- | --- | --- |
| Everyone | Git, editor and repository access | `git --version` |
| A and E | Go; inspect kan-agent/go.mod and CI | `go version` |
| B, C and D | JDK and Maven; inspect kan-server/pom.xml and CI | `java -version`, `javac -version`, `mvn -version` |
| Contract contributors | Python 3 and checker prerequisites | `python3 --version` |
| Backend integration owners | Separately approved capable host and reviewed runtime/profile | Agree requirements before installing or creating resources |

The inspected repository declares Go toolchain **1.27.2** and Java **25**. Recheck your checkout; do not silently change pins to match your machine. A missing or unavailable toolchain is a blocker to report. No Maven wrapper is assumed.

You do not need to create a VM, Docker container, cluster or RustFS service just to read the code or start a bounded task. Follow [installation boundaries](docs/platform/installation-and-operations.md); Compose contains dependencies, not the whole application. No supported production host profile exists yet.

## 5. Clone and create your feature branch

Shell examples use Bash/Git Bash. They are instructions for teammates; this documentation update did not run them against GitHub.

### 5.1 Clone once

```bash
git clone https://github.com/Ciago-Technologies/covestack-cloud.git
cd covestack-cloud
git status -sb
git fetch origin
git branch -r
```

If you already cloned it, start with `git status -sb`. Preserve your work; do not clone over it or use reset/clean to get past changes.

### 5.2 Confirm the starting branch with Anuj

Use `main` only when Anuj confirms the needed baseline changes are merged. Otherwise use the exact reviewed baseline branch he provides. Do not assume another person's feature branch is the agreed baseline.

```bash
COVESTACK_BASELINE_BRANCH=main
```

Replace `main` before continuing if a different baseline was agreed. This is a shell variable, not an application setting.

### 5.3 Create one branch for your task

These are suggested names, not branches that already exist:

| Person | Suggested branch | First task |
| --- | --- | --- |
| A | feat/a-agent-server-integration | Existing signed loop against real assigned server credentials |
| B | feat/b-server-foundation | Bootable Java foundation and reviewed migration plan |
| C | feat/c-auth-acceptance | Real credential/authentication path and acceptance foundation |
| D | feat/d-api-contract-review | Actual endpoint/DTO inventory and compatibility review |
| E | feat/e-cli-foundation | Reviewed CLI module/client boundary and mock API learning slice |

Example for B; substitute your own branch name:

```bash
git switch -c feat/b-server-foundation "origin/$COVESTACK_BASELINE_BRANCH"
git status -sb
```

Create only your own branch. If it exists already, switch to it rather than recreate it. Everyone clones the whole repository and pushes a branch for the repository; nobody pushes “into kan-server”.

## 6. Agree a small first issue

Before coding, record the following in one GitHub issue:

| Field | What to write |
| --- | --- |
| Outcome | One concrete behavior to deliver |
| Owner and reviewer | Who changes it and who reviews it |
| Files / boundary | Which existing module is affected; which interface needs agreement |
| Dependency | Another person's server, credential, fixture or schema decision you need |
| Acceptance | The exact assertion that demonstrates the behavior |
| Evidence | Command/result and whether it is unit, mock, real-server or runtime evidence |
| Next action | One clear step or blocker; no guessed completion percentage |

Use the existing project board consistently:

| Column | Meaning |
| --- | --- |
| Backlog | Idea or task whose scope/dependencies are not ready |
| Todo | One owner, clear acceptance and dependencies ready |
| In progress | Owner is actively working or verifying |
| In Review | Reviewable PR opened with actual evidence |
| Done | Reviewed/merged task with acceptance evidence; this does not mean a whole service is Supported |

GitHub issues own live task status. Documentation owns architecture; checkpoint.md owns AI-agent continuity. Link them rather than maintaining conflicting task copies. If an issue is blocked, name the dependency and owner instead of inventing an API or credential.

## 7. A / Anuj — your next work

**Start with integration of the existing signed polling loop, not another loop or VM driver.**

| Step | What you do | Files / people | Reviewable result |
| --- | --- | --- | --- |
| 1 | Walk through how configuration creates the client, how signing works and how polling stops | kan-agent/cmd/kan-agent/main.go; internal/config/, internal/client/, internal/agent/loop.go; sigv1/ | You can explain the current path and its no-execution boundary |
| 2 | Agree real server URL, protected test credentials and trusted account/agent assignment | B+C; [API boundary](docs/platform/api-and-cli.md) | Written integration agreement; no secret values in issue/docs |
| 3 | Check existing unit/static baseline and attach actual results | kan-agent/ and current CI | Pass or precise blocker; no runtime claim |
| 4 | With B+C's server ready, test signed no-work polling, invalid/revoked credentials, cancellation and errors | Existing client/loop and C's integration tests | Real-server evidence; record any retry gaps instead of assuming they work |
| 5 | Review future work-claim, VM/type and network contracts before execution code | A+B+C+D; contracts/ read first | Coordinated design/fixture plan; no invented current fields |

Actual discovery is `GET /internal/v1/agent/work`; a successful empty array means no work. `KAN_SERVER_URL`, `KAN_ACCESS_KEY`, `KAN_SECRET_KEY` and optional `KAN_RULES_PATH` are current config names. The loader does not automatically read .env. B+C must supply real assigned credentials securely; a localhost URL shown in a contract is not proof a listener exists.

Later you own provider/runtime drivers, inspection/reconciliation, capacity reporting, signing, CI, installer and ITISS evidence. Incus evaluation requires a separately approved capable host. Preserve the agent's privileged boundary; no tenant runtime sockets, arbitrary host commands or direct DB access. Explain the Go you contribute. A shared pkg/kanclient layout is still a reviewed future task; do not move the signer independently.

## 8. B — Java core and database foundation

1. Inspect kan-server/pom.xml, existing Java packages and migration files. Current source does not establish a bootable product server or foundational schema; existing extended migrations reference prerequisite tables.
2. Agree the next minimal server/auth/assignment slice with C, D and A. Review current V5 VM-first schema requirements; old full DDL is archived and not executable/migration-ready.
3. Own Java application startup/configuration, the forward-only migration plan and all lease/fencing/idempotency/quota SQL. Never edit an already applied migration; establish actual history first.
4. Make the reviewed foundation testable from an empty approved database, with real authentication and owner-scoped queries. C owns acceptance; A needs a real endpoint/assignment path for signed polling.
5. Record startup/migration/auth evidence and unresolved protocol decisions in your PR.

**First output:** a reviewed foundation/migration dependency plan, followed by the smallest bootable authenticated slice. Locations: kan-server/src/main/java/ and kan-server/src/main/resources/db/migration/. B owns SQL; use JdbcClient/plain SQL, not JPA.

## 9. C — security, credentials and acceptance

1. Read kan-server/src/main/java/com/ciagotech/kan/server/auth/ and contracts/ signing/lifecycle rules.
2. Agree credential creation/lookup, revocation, secret custody, replay behavior and trusted assignment with B+A. A caller-provided account/agent identifier is not trusted assignment.
3. Keep hash-only API tokens separate from encrypted signing credentials. No shared default key, plaintext secret logs or unauthenticated shortcut.
4. Establish acceptance tests from the first slice: missing/bad/replayed signatures, wrong assignment/account and revocation. Distinguish mocks from real Java/Valkey evidence.
5. Later own Java IPAM allocation and all domain-specific acceptance, including actual packet/storage denial and restore verification.

**First output:** reviewed credential/assignment behavior and reproducible authentication tests that A can exercise against B's server. RustFS-scoped policies, credential rotation and cross-account denial are future acceptance work, not already verified.

## 10. D — Java API and compatibility

1. Read contracts/internal-api.yaml, lifecycle-rules.json and existing Java auth/contract helpers before defining DTOs.
2. Inventory existing endpoints versus interfaces that still need implementation. Coordinate public account/resource API boundaries with B+C and internal assignment/claim needs with A.
3. Document the VM Instance/App Container compatibility gap. Current contract 1.1.0 is unchanged and must not silently reinterpret an old container specification as a VM request.
4. Agree request/response/error fixtures and DTO ownership before implementing the reviewed endpoint slice. No full /v2/provider parity promise.
5. Later help with Terraform compatibility fixtures and API mappings once actual resources exist.

**First output:** reviewed endpoint/DTO/fixture plan for the next server slice. Location: kan-server/ and reviewed contract sources. **D writes Java, not Go.** Contract changes need coordinated schemas, fixtures and generated outputs.

## 11. E — Go CLI learner

1. Learn the current API/client boundary and review the proposed CLI/module layout with A before adding modules or moving shared signing code.
2. Start with a small cobra command and three-OS packaging exercise after that layout is agreed. Future targets are linux/amd64, darwin/arm64 and windows/amd64.
3. Practice net/http, JSON, context cancellation, finite timeouts and HTTP error handling against a clearly labelled mock API. Mock responses are learning/test fixtures, not provisioned resources.
4. Add reviewed lifecycle commands only as the real APIs become available; use idempotency for mutations and stable JSON for list/get. Keep credentials out of command arguments/logs and hold no local resource database.
5. With A/C review, move the same thin client to the real server and verify actual user authorization. A later Terraform provider stays a thin API layer; it does not become a second control plane.

**First output:** an agreed module/client plan and small CLI/mock test slice. cmd/covestack/ and pkg/kanclient/ are planned locations, not claims that those modules exist. E does not implement agent, signing, drivers or reconcile; no direct Docker/Incus/k3d/DB calls.

## 12. First shared integration checkpoint

| Order | Owner | Outcome before moving on |
| --- | --- | --- |
| 1 | B+C+A; D compatibility | Agree real credential provisioning, trusted assignment and the current discovery/no-work response |
| 2 | B+C | Bring up the reviewed server/auth/foundation on an authorized development setup; record actual evidence |
| 3 | A+C | Existing signed client polls real no-work safely; test auth rejection, cancellation and observed failure behavior |
| 4 | B+D+C; A review | Agree future mutation/claim/type/schema/fixture changes before runtime execution |
| 5 | A+B+C+D | Implement a bounded owned resource lifecycle with capacity/security and meaningful failure tests under separate task approval |
| 6 | E with A/C | Expose the real authenticated workflow through the CLI; retain clear mock versus runtime evidence |

RustFS integration remains Planned: operator manages the process/persistent volumes; Java reconciles bucket/policy/credential intent. Test actual object bytes, foreign-account denial, key revocation, restart persistence and verified backup/restore. Recreating an empty bucket does not restore lost data. See [provider acceptance](docs/platform/provider-evaluation.md#rustfs-compatibility-and-acceptance).

## 13. Verify your work without overstating it

Run the checks appropriate to your changed module on your authorized development environment. These match current repository CI; this documentation update did not execute them.

From kan-agent/:

```bash
go test -race ./...
go vet ./...
```

From kan-server/:

```bash
mvn --batch-mode --no-transfer-progress verify
```

For relevant contract changes, from the repository root with checker prerequisites available:

```bash
python3 contracts/sign_example.py
python3 contracts/check_lifecycle.py
```

Report missing tools/network/dependencies as blockers. A successful Java build does not prove the server boots; a passing signer test does not prove trusted assignment; mock create does not prove a VM boots. Do not invent startup commands for a missing application or bypass authentication to get green output.

A service becomes **Acceptance Gates Passed** only with persisted authoritative intent (G1), crash-safe reconciliation (G2), ownership and actual isolation (G3), and meaningful automated failure acceptance (G4). **Supported** additionally requires a qualified profile and maintenance policy. Follow [readiness](docs/platform/readiness.md); no status promotion from documentation alone.

## 14. Commit, push and raise a PR

Review your own diff first. Replace the staged path with the exact files for your task; do not blindly stage teammates' changes or secrets.

```bash
git status -sb
git diff --check
git diff --stat
git add path/to/your/changed/files
git diff --cached --check
git diff --cached --stat
git commit -m "Describe the concrete change"
```

Then push **your feature branch**; example for B:

```bash
git push -u origin feat/b-server-foundation
```

Open a PR into the agreed baseline, usually main once integration prerequisites are merged. Include behavior, evidence, dependencies, limitations and contract/schema impact. Do not push directly to main or force-push over another person's work.

| Change | Review |
| --- | --- |
| Any Go | A |
| Shared client/signing or VM/backend boundary | A+B+C; D for wire/DTO impact |
| SQL/schema | B+C+A; D for DTO alignment |
| Authentication/security/acceptance | C plus affected implementers |
| Java API/DTO | B/C/D as affected; A for agent-facing changes |

## 15. Daily follow-up

- Check your GitHub issue and the latest agreed baseline before starting.
- Keep local work committed/stashed safely before merging updates; never discard changes to force a clean tree.
- Fetch and merge the agreed origin baseline into your feature branch after confirming the branch name; do not silently switch everyone to a new baseline.
- Update the issue with behavior delivered, actual test results, blocker/owner and the next small step.
- Update affected current documentation; AI agents also maintain checkpoint.md. Preserve dated evidence instead of rewriting past results.
- Ask for review when your slice is concrete; do not wait for an entire service to be finished.

Never commit real .env files, deploy/secrets/, signing/master/private keys, credential-bearing kubeconfigs, Terraform state or secret logs. Never give tenants Docker/Incus sockets or host root access. Keep kan-server, kan-agent, KAN1 and existing contract names until a coordinated compatibility change is approved.

CoveStack-owned code/documentation uses [Apache License 2.0](LICENSE), SPDX **Apache-2.0**. Third-party notices and rights remain separate. The previous starter is preserved as [historical reference](docs/archive/pre-vm-redesign-2026-11-04/Starter-original.md); follow this current guide for new work.
