# CoveStack — First Team Checkpoint and GitHub Board Guide

| Document control | Value |
| --- | --- |
| Updated | 27 October 2026 |
| Scope | First implementation checkpoint; target work, not acceptance evidence |
| Team | A: Anuj; B: Java core/SQL; C: security/tests; D: Java API/DTO; E: CLI |
| Existing first checkpoint | 31 October 2026; report missing evidence honestly rather than marking tasks finished |
| Submission deadline | Monday, 25 January 2027 |
| Service status | Planned until actual evidence justifies a change under G1–G4 |

## 1. How to maintain this and GitHub

Use two connected records with different responsibilities, rather than two competing sources of truth:

| Record | Authoritative for | How to maintain it |
| --- | --- | --- |
| This file and linked V5/contract documents | Checkpoint scope, ownership, dependencies and exit criteria | Change through a reviewed PR; do not tick live progress here |
| GitHub issues and project board | Assignee, current task status, blockers, PR and test evidence | Update the same issue as work progresses; link this file/section |
| PR/CI results | Exact code revision reviewed and tests actually executed | Attach links to the issue; mocks do not establish real runtime results |

The board is a view of issues, not a separate task list. Create one issue per bounded task and add that issue to the project. Do not duplicate an issue as a draft card. Use repository links from the agreed branch, then a stable merged revision when available. Never include credentials or secret configuration in either record.

A task in the board's Done column is not a service at Gates passed. Service status remains **Planned | In progress | Gates passed | Extended track**. Supported requires G1 DB-authoritative intent/generation, G2 crash-safe reconciliation/rebuild without duplicates, G3 other-account 404 plus applicable protocol/network denial, and G4 a failure test running in CI. C owns tests 1–15 and E1–E8. This first communication checkpoint passes no resource service gate by itself.

## 2. Everyone: start after cloning

Use the whole repository; do not clone just kan-server or kan-agent. These are instructions to run deliberately, not commands already executed by this document.

```bash
git clone https://github.com/Ciago-Technologies/covestack-cloud.git
cd covestack-cloud
git status -sb
git switch main
git pull --ff-only origin main
git switch -c feat/your-specific-task
```

If the checkout already has uncommitted work, preserve it before switching branches. If your task already has a branch, switch to it rather than creating another. Anuj's existing task branch is feat/agent-signed-client. Branch from the reviewed baseline; coordinate shared changes rather than copying unfinished files between teammates.

| Read | Purpose |
| --- | --- |
| [Team onboarding](V5/CoveStack-Team-START-HERE.md) | Shared rules and first tasks |
| [SRS scope](V5/CoveStack-SRS-Scope.md) | Exactly CORE-5 and approved cross-cutting scope |
| [Roles](V5/Team_Roles.md) | Code ownership and reviewers |
| [Workflow](V5/CoveStack-Project-Workflow.md) | Dates and dependency order |
| [Setup](V5/CoveStack-Complete-Working-Setup.md) | Existing commands, prerequisites and warnings |
| [Agent API](contracts/internal-api.yaml) and [lifecycle rules](contracts/lifecycle-rules.json) | Actual contract; version 1.1.0, not guessed messages |

Open your assigned issue before coding. Write its expected result and dependencies. Run the existing relevant checks and record the baseline, including failures. Each person creates private local configuration from reviewed templates; never copy another person's real .env, keys or deployment secrets.

## 3. Person A — Anuj: signed communication first

**First task:** understand the existing client and integrate one-account signed communication. client.go/client_test.go exist; source presence does not prove current review, real-server connectivity or a production loop. Review the code and explain it before extending it. AI-generated agent code is reference only; A types and can explain its behavior.

| Order | Your action | Evidence / dependency |
| --- | --- | --- |
| 1 | Read config/config.go, sigv1/request.go, internal/client/client.go and its tests; study auth-smoke | Explain configuration, signing, HTTP errors, decoding and cancellation |
| 2 | Review client tests and run relevant Go checks from kan-agent | Record actual output/revision; distinguish local test server from Java server |
| 3 | Pair with B+C on URL, one account, service credential and trusted host assignment | B supplies real server; C securely provisions key; no caller-selected account |
| 4 | Use client.GetWork for one real signed GET | HTTP 200 with [] when no work; no dispatch/acknowledgment |
| 5 | Add agent startup and polling integration after reviewing entry-point/file placement | One account, 10-second requests, 5-second polling base, cancellable waits |
| 6 | Add temporary-failure backoff, authentication stop and shutdown tests | 1-to-30-second backoff cap with bounded jitter; valid 429 guidance; 401/403 stops account loop; Ctrl+C interrupts request/wait |
| 7 | Pair with C against the real Java server | Outage/restart recovery, invalid key, no work and no secret logs; attach results |

Relevant locations: kan-agent/internal/config/, kan-agent/sigv1/, kan-agent/internal/client/, kan-agent/cmd/auth-smoke/. Agent production entry point/poll files and shared pkg/kanclient module layout still require review; do not claim they exist or move the signer blindly.

The client currently has a local response-size bound; review it with B+C before release. GetWork is discovery, not permission to execute. Start with supportedKinds empty; heartbeat requires truthful version inputs, including dockerVersion. No invented Docker version. Empty/failed polls never authorize cleanup.

**Exit:** you can explain the code; unit/client checks pass; one assigned credential communicates with the real server; no work is normal; outage recovery and shutdown are demonstrated. C records integration evidence. **Next:** reviewed claim/network schemas, then one Droplet lifecycle with crash recovery and fencing; do not jump to multi-account concurrency or extended publication.

Full details: [Anuj first milestone, section 2.3](V5/CoveStack-Anuj-Next-Steps.md#23-first-milestone-decision--single-account-communication-planned).

## 4. Person B — Java application and database foundation

**First task:** provide a bootable server and safe reviewed foundation, not every table at once.

| Order | Your action | Exit evidence |
| --- | --- | --- |
| 1 | Read schema, inspect current Flyway files and determine any applied history | Written prerequisite/order review with C+A and D for DTOs |
| 2 | Resolve the missing foundation before existing V5/V6 extended migrations can run | Empty PostgreSQL 17 migration test; no rewriting an applied migration |
| 3 | Implement the reviewed initial identity, credential, resource, operation, idempotency/event/audit foundation | Tenant constraints and transaction tests; staged scope from schema section 9.1 |
| 4 | Add/verify actual application entry point, configuration and packaging | Real Java process starts using reviewed configuration; no imaginary startup command |
| 5 | Pair with C on credential lookup and trusted account/agent assignment | Persisted assignment is checked; no global/default credential |
| 6 | Expose existing read-only work behavior with D/C | Authenticated assigned account, capability filtering, 200 [] for no work |

**Exit:** teammate can follow the verified startup procedure; foundation migrates cleanly; real authentication integrates and the empty work path behaves correctly. **Next:** scoped resource transactions and B-owned lease/fencing/idempotency SQL after coordinated schemas. Agent/client never writes the DB directly.

Read [database schema](V5/CoveStack-Database-Schema.md), especially migration inventory and open decisions. Existing publication/domain SQL does not establish a functioning CORE-5 database. Share settings names and startup instructions, never DB passwords.

## 5. Person C — credentials, security and acceptance from day one

**First task:** make authentication real and own the tests that prove the first integration.

| Order | Your action | Exit evidence |
| --- | --- | --- |
| 1 | Review auth filters, KeyService, nonce store and B's credential schema | Agree ciphertext key custody, service eligibility and account/agent assignment |
| 2 | Implement real active service-credential lookup with B | Unknown/revoked keys rejected; no fallback development key |
| 3 | Provision one dedicated service credential securely | Restricted issuance, secret delivered once by protected channel, no Git/chat/log leakage |
| 4 | Verify signature/replay behavior and store outages | Valid request accepted; replay/tampering denied; unavailable replay protection fails closed |
| 5 | Own real-server communication tests with A+B/D | Empty work, wrong credentials, account scope, outage/restart and sanitized errors |
| 6 | Prepare CORE-5 acceptance tests incrementally | Tests 1–15 remain C-owned; no resource gate promoted by the communication test |

**Exit:** real stored credential authenticates; invalid/revoked/replayed requests are denied; assignment and secrets are reviewed; tests run reproducibly. **Next:** tests 1–6 with B+A and Java IPAM allocator/test 14 in agreed dependency order. Do not wait until the end to add acceptance tests.

## 6. Person D — exact Java API/DTO contract

**First task:** align the real Java HTTP boundary with the existing contract and identify missing execution messages.

| Order | Your action | Exit evidence |
| --- | --- | --- |
| 1 | Read internal-api.yaml and generated/base-source instructions; pair with B+C | Agreed existing route, auth context and 200 array response |
| 2 | Implement/review existing work DTO/controller mapping with B | Scoped query, capability filter, server-owned kinds excluded; empty result is [] |
| 3 | Add or pair on Java HTTP serialization/negative fixtures | Missing/invalid requests fail; no accidental secret fields; C owns acceptance suite |
| 4 | Review UUID, CLI public ID and numeric /v2 ID mappings | Written mapping decisions with B; no invented wire IDs |
| 5 | Specify claim/renew/report gaps with B+A/C before execution | Coordinated schema/base source, examples and generated outputs in a later reviewed code change |

**Exit:** Java wire response matches existing contract; open message gaps are explicit; A can consume the real response without guessed fields. **Next:** bounded /v2 Droplet/Kubernetes mappings and fixtures. D writes no Go and does not take B's SQL ownership.

## 7. Person E — CLI foundation and learning boundary

**First task:** build a thin cobra CLI foundation, not an agent or frontend.

| Order | Your action | Exit evidence |
| --- | --- | --- |
| 1 | Read Go handover; agree module/import layout with A | No conflicting root module or premature signer move |
| 2 | Learn structs/errors/modules; create agreed cobra entry point/help | Explain command flow; no fake resource success |
| 3 | Practice HTTP/JSON/context against a mock API | Timeout, cancellation, non-2xx and malformed response checks |
| 4 | Establish packaging for linux/amd64, darwin/arm64, windows/amd64 | Cross-build artifacts; build success is not actual OS-run evidence |
| 5 | Agree JSON/error/idempotency/wait behavior with A+B/D/C | API-only path; no Docker/k3d/DB dependency; separate user auth from service key |

**Exit:** bounded CLI foundation reviewed by A; mock tests and builds recorded; no agent ownership. **Next:** real droplet create/list/get/delete once corresponding API is available; production order droplet → vpc → k8s → db → bucket → stack, previous failure test passing in CI first. W3 Kubernetes practice remains mock-only. Terraform and read-only UI are not first tasks.

Read [Go handover](V5/CoveStack-Go-Handover.md). CLI must not reuse the agent's service key as ordinary user authorization; C+B define the correct API-token flow.

## 8. Shared first-checkpoint exit and permission to advance

| Check | Owners | Required evidence before marking this checkpoint accepted |
| --- | --- | --- |
| Real server | B | Reviewed startup and clean-DB migration results |
| Real identity | C+B | Dedicated service credential, eligibility/replay checks and trusted assignment |
| Exact API | D+B/C | Read-only scoped work response; 200 [] with no advertised execution kinds |
| Agent communication | A with C | Valid request, outage/recovery, invalid-auth stop and prompt shutdown |
| CLI foundation | E with A | Explainable cobra skeleton, transport practice and cross-build evidence |
| Next execution contract | B+D+A/C | Missing claim/renew/report/network encoding identified; approved schemas required before backend mutation |

Parallel work is allowed: A can test transport and E can learn against mocks while B/C/D implement the real server. Neither mock path passes live integration. If blocked, update the issue with the dependency, owner and next action; do not pretend the dependency is finished.

At handoff attach PR, commit, executed checks, runtime prerequisites and known gaps. Service/core acceptance remains governed by [Service Gates](V5/CoveStack-Service-Gates.md) and [tests 1–15](V5/CoveStack-AntiWrapper-Acceptance.md). Do not move to extended work to bypass an unmet core checkpoint.

## 9. Configure the board shown in the screenshot

Recommended visual order: **Backlog → Todo → In progress → In Review → Done**. Your screenshot places Done before In Review; reorder if helpful. These are task states, not lifecycle or service-state strings.

| Board column | What belongs here | Move forward when |
| --- | --- | --- |
| Backlog | Later tasks and ideas not ready for execution | Dependencies and acceptance are clear enough to schedule |
| Todo | Small ready tasks with one owner, acceptance and known dependencies | Assignee actually starts |
| In progress | Task actively being implemented or verified | Reviewable PR exists with relevant test evidence |
| In Review | PR awaiting review/checks or requested revisions | Review requirements and checks pass; PR merged into agreed baseline |
| Done | Accepted task with merged PR and required evidence | No automatic service-gate promotion |

Keep the screenshot's In progress limit of five for the team: normally one active issue per person. For Anuj, one active implementation task at a time. A blocked task keeps its honest status plus a blocked label/dependency; moving it to Done is not a blocker workaround. Partial code can be a draft PR and remain In progress.

Suggested issue/project fields: Assignee, Status, Priority, Area, Checkpoint, linked PR and blocked-by issue. Start with the current checkpoint; use the existing workflow for later dates. Do not invent teammate usernames, issue numbers or completed dates. Add labels by component/owner role if useful; role A/B/C/D/E does not replace an actual assignee.

### 9.1 Anuj's first issue cards

These are issue-title suggestions, not issues already created. Choose status from actual evidence; this document does not assert current board status.

| Suggested title | Initial placement guidance | Acceptance / dependency |
| --- | --- | --- |
| A: Review and explain the signed work client | In progress if you are doing this now; otherwise Todo | Explain New/GetWork and failure paths; run checks; review changed files |
| A: Verify signed client against Java server | Todo if B/C server/key ready; otherwise Backlog with dependencies | Valid signed GET gives 200 []; no secrets; C records result |
| A: Add one-account agent startup and poll loop | Todo after client review and agreed entry-point layout | No execution; 10-second timeout; normal wait and no-work handling |
| A: Add polling backoff and safe shutdown | Backlog until polling slice is ready | Bounded retry/jitter, 429 policy, auth stop, cancellation tests |
| A+C: Demonstrate real-server outage and recovery | Backlog until loop and B/C environment ready; A assigned, C reviewer/test owner | Stop/restart server, invalid credential and Ctrl+C scenarios recorded |
| A+B+C: Review shared-client module layout | Todo when reviewers available; A assigned | Decide layout; no unreviewed move/import changes |
| A: Implement container reconciliation | Backlog | Requires execution/claim/network schema and C's tests; test 1 no duplicate |
| A: Integrate VPC, Kubernetes and persistence | Backlog as planning epic; split before execution | Existing staged CORE-5 dependencies; not one giant active task |

Do not mark client implementation Done just because client.go exists or an earlier chat reported tests. Locate the relevant PR/revision, rerun appropriate checks when needed and attach evidence. If already merged and accepted, record that actual PR and move the issue to Done; the separate real-server milestone can still be pending.

### 9.2 First cards for the other owners

| Owner | Initial issue title | Dependency / exit |
| --- | --- | --- |
| B | B: Bootstrap Java server and reviewed foundational migrations | Inspect V5/V6 prerequisites; clean PostgreSQL 17 migration and boot |
| C | C: Provision service credentials and verify replay protection | B identity/store integration; valid and invalid real-server cases |
| D | D: Align scoped work endpoint DTO and HTTP contract | B query and C identity; 200 [] plus negative HTTP cases |
| E | E: Build reviewed cobra CLI foundation | A module review; help, tests and three-target builds |

Split these into reviewable issues when they span several PRs. A task can be ready for design/test work while blocked on later integration; describe the boundary instead of hiding it.

### 9.3 Copy this into each issue body

```markdown
## Outcome
One observable result for this task.

## Scope and owner
Assignee: actual teammate username
Role/component:
Checkpoint: first integration
Reference: TEAM-FIRST-CHECKPOINT.md, relevant section

## Acceptance
- [ ] Expected behavior and error cases described.
- [ ] Relevant unit/integration checks executed and linked.
- [ ] Real runtime evidence attached where promised; mocks labelled clearly.
- [ ] No secrets in source, output or screenshots.
- [ ] PR reviewed by the boundary owners and merged.

## Dependencies
Blocked by: actual issue links, or none
Reviewers:

## Evidence
PR / commit:
Commands and actual results:
Known gaps / next action:
```

No GitHub changes are performed by this document. Create the issues manually or authorize a separate board-update task. Keep live checkboxes/status in those issues; change this file only when the agreed checkpoint or acceptance changes.
