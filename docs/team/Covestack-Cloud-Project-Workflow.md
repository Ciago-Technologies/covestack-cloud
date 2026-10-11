# Covestack Cloud — Delivery Workflow to 15 December 2026

| Item | Details |
|---|---|
| 1 | Plan updated **Sunday, 11 October 2026**. Final submission is **Tuesday, 15 December 2026**; half demo is **Saturday, 28 November 2026**. There are 65 calendar days after today. Dates are actual 2026 weekdays, not copied 2025 Friday labels. |
| 2 | Read [design](Covestack-Cloud-Technical-Design.md), [setup fixes](Covestack-Cloud-Current-Setup-Fixes.md) and [team roles](Covestack-Cloud-Team-Roles.md). This plan assumes five available team members, daily short coordination and a running integration VM. It is a schedule with measurable gates, not a guarantee independent of progress. |

## 1. Baseline and rules for finishing

| Item | Details |
|---|---|
| 1 | The security/build baseline is committed as `349a2f4` (Fixed security issues) on published branch `fix/security-baseline`, which tracks `origin/fix/security-baseline`. The lifecycle changes are currently applied but uncommitted. Remote main/PR merge status has not been verified; review and merge through the shared PR. Three signature vectors, negative tests, registered agent/Bearer filters, Valkey replay plumbing, CI, secret templates and auth-smoke exist. The config.go file currently contains only `package config`; it is a placeholder, not a working loader. Real credential/token lookup, application main, migrations/controllers, full agent and live deployment tests remain pending. |
| 2 | Product name: **Covestack Cloud**; intended repository slug: `covestack-cloud`. The actual checkout and GitHub remote are still named `kan-cloud` until the repository owner renames them. Use the real paths in commands. Existing module folders, Go import paths, Java packages, signing headers/algorithm, KAN_* environment names, API paths, labels, database identity and Valkey key prefixes remain compatibility identifiers. A future technical rename needs a separate coordinated migration. |
| 3 | One Java application, one host agent and one integration VM. Avoid microservices, module relocations, a new frontend, new cloud providers and new frameworks. Keep maintained Java 25-compatible dependencies and update the actual contracts with Java/Go together. |
| 4 | Security gates are not optional feature cuts: auth, replay/body limits, ownership, workload/backend separation, persistence and truthful action completion must work before release. Planned feature depth is intentionally small so the required application stack can finish. |
| 5 | A task is done only when its acceptance check passes on the integration VM, CI is green and the relevant guide describes actual behavior. Stubs/fakes are explicitly labeled. Never merge a bypass so a demo turns green. |

## 2. Scope and order

| Priority | Deliverable | Depth |
|---|---|---|
| P0 | Safe auth, corrected contracts, entry points, migrations, repeatable build | Fix existing audit blockers first |
| P1 | Droplets/actions/account/catalogs/tags/SSH/CLI/Mailpit/ASG | Complete tested lifecycle |
| P1 | Managed Postgres + Spaces | Real persistent data, account-scoped credentials |
| P2 | HTTP LB | One rule, healthy members, tested reload |
| P2 | Kubernetes | One tested version/node shape and reachable kubeconfig |
| P2 | Snapshot | DO action + scoped container image; volume exclusions documented |
| Deferred depth | MySQL, complex DB operations, advanced LB/DOKS, multi-provider | Add only after all required paths and security gates pass |

| Item | Details |
|---|---|
| 1 | P2 features are scheduled below, not silently removed. If a stability gate is missed, keep required paths stable and publish exact supported/501 behavior rather than presenting unfinished features as working. |

## 3. First 72 hours: 12–14 October

| Owner | Monday 12 | Tuesday 13 | Wednesday 14 acceptance |
|---|---|---|---|
| Anuj | Pin Go/tools; add Go/Java CI with C; establish VM baseline | Update signed-header rules and Go signer; infra/secrets proposal | Go vectors/negative tests; recorded VM versions |
| B | Spring entry point, config and Flyway Postgres dependency | Accounts/resources/operations migrations and scoped repos | Fresh DB migration + application boots |
| C | Filter registration, bounded body and credential principal | Verify-before-nonce, shared authenticated replay store | Replay mutation/poison/body/key-kind tests pass |
| D | Provider response inventory and contract schemas with B | Account controller, states/examples/HTTP codes | Authenticated account read + contract-example validation |
| E | Local bootstrap and restricted credential files | Health + minimal CLI configure/doctor scaffold | Initial user token works; no public key-issuance bypass |

| Item | Details |
|---|---|
| 1 | Pair C with Anuj on signing and B with D on the API/database boundary. Use a temporary dev credential implementation only in isolated automated tests; the integration profile must use the real store. Acceptance date may expose missing work; move it explicitly into the next gate, not into a hidden TODO. |

## 4. Weekly gates

| Gate | Date | Required evidence | Recovery if missed |
|---|---|---|---|
| G1 | Friday 16 Oct | App boots; migrated DB; real registered auth; principal/service key; corrected shared signed-nonce vectors; CI builds Java/Go | Entire team fixes G1 before feature expansion |
| G2 | Friday 23 Oct | Real create → agent → container → observed → action complete; negative tenant/auth tests; fake and real runs distinguished | B/C/D pair on flow; Anuj pauses extra drivers |
| G3 | Friday 30 Oct | Power on/off/reboot/delete; retry failure produces failed operation attempt; key-only SSH; ports rebuilt after agent restart | No advanced API work until lifecycle is reliable |
| G4 | Friday 6 Nov | ASG pending counts/lease/outage handling; tags and SSH key API; CLI commands; mail after commit | Fix slow-pull/outage tests; preserve core count limits |
| G5 | Friday 13 Nov | Persistent managed Postgres; actual returned user/DB; scoped Spaces credentials; object/DB recreation tests | Whole next integration slot devoted to data/connectivity |
| G6 | Friday 20 Nov | HTTP LB works; bounded Kubernetes path; snapshot action baseline; public DNS/TLS + second device | Stop expanding depth; repair supported paths |
| G7 | Friday 27 Nov | Half-demo rehearsal from clean start; auth/internal routing/persistence evidence; timed script | Bug fixes only through half demo |
| Half demo | Saturday 28 Nov | Real stable supported features; known limitations stated | Record feedback and prioritize defects |
| Feature freeze | Tuesday 1 Dec | Planned feature paths pass or are explicitly unsupported; no new scope | Finish defects and integration, not new framework work |
| G8 | Friday 4 Dec | Packaged app/agent/CLI; clean-machine install; security scans and restore | Packaging/security fixes take priority |
| Release candidate | Tuesday 8 Dec | Full supported script on clean VM + second OS; no unresolved release-blocking defects | Fix blockers and rerun affected acceptance |
| Final rehearsal | Friday 11 Dec | Backup/restore, certificate check, install artifacts and evidence complete | Saturday reserved for stabilization |
| Buffer | Sunday 13–Monday 14 Dec | Critical fixes only; final archive/report ready | No new features or cosmetic refactors |
| Submit | Tuesday 15 Dec | Tested version, source, binaries, docs and evidence delivered | Submission, not development day |

## 5. Work allocation by phase
### 12–23 October: unblock and integrate

| Item | Details |
|---|---|
| 1 | Anuj: signed Go client, driver skeleton, scoped labels, pinned CI/tools and infrastructure readiness. |
| 2 | B: real entry point/migrations/repositories, transactional create/operation and agent work/observed APIs. |
| 3 | C: full auth matrix, service principal, AES-GCM/secret handling, authenticated shared replay store and negative tests. |
| 4 | D: account/Droplet DTOs, kind states/spec validation, response/error codes, numeric IDs and SDK test. |
| 5 | E: local bootstrap, CLI config, restricted credentials, minimal public health and authenticated doctor interface. |
| 6 | Gate evidence includes a correctly signed nonce mutation being rejected and a valid resource action actually reaching Docker; old two-vector success alone is insufficient. |

### 24 October–6 November: complete compute and ASG

| Item | Details |
|---|---|
| 1 | Anuj: restart/stop/delete handling, safe port allocator, periodic observations, dirty-event retries, agent lease/fencing and actual SSH test. |
| 2 | B: generation/report-order rules, restart intent, failure/retry semantics, atomic idempotency, ASG reservations/parent retirement/scale-down. |
| 3 | C: ownership reviews on every endpoint, revoked-token behavior, replay restart/store-failure tests and quotas. |
| 4 | D: all core action reads/writes, pagination, tags/SSH keys and error DTOs; godo/doctl lifecycle checks. |
| 5 | E: CLI instances/actions/asg, remote doctor, mail outbox delivery and recovery instructions. |
| 6 | Test an image pull slower than two scheduler ticks and an agent outage. Neither may create unlimited replacements or mark a failed operation completed. |

### 7–13 November: persistent application stack

| Item | Details |
|---|---|
| 1 | Anuj: managed Postgres volume/readiness/ports and MinIO persistence; client-reachable network mapping. |
| 2 | B: DB service/credentials, account-scoped Spaces metadata and explicit deletion/backup behavior. |
| 3 | C: scoped S3 policies and cross-account read/write denial; workload/backend network verification by IP. |
| 4 | D: DB DTO/API/status and golden files; authenticated SDK checks. |
| 5 | E: DB/Spaces CLI, masked lists, scoped upload/download and readiness output. |
| 6 | Create data, recreate containers without deleting volumes and verify data. Test restore from backup. User DB containers never join the control-plane backend. Do not implement MySQL until the required Postgres path passes. |

### 14–20 November: bounded LB, Kubernetes and snapshot

| Item | Details |
|---|---|
| 1 | Anuj: valid HTTP LB config directory/reload, one k3d driver/version/network, scoped snapshot images. |
| 2 | B: healthy membership sync, cluster/snapshot services, truthful observations/actions. |
| 3 | C: validate shell/tool arguments, ownership, kubeconfig/credential protection and k3d privilege boundary. |
| 4 | D: LB/DOKS DTOs and snapshot through the Droplet action endpoint; schema/SDK compatibility for supported subset. |
| 5 | E: CLI lb/k8s/snapshot paths; separate kubeconfig export; demo-friendly supported-feature listing. |
| 6 | Use an HTTP-ready workload, not an SSH-only image, for LB responses. Pin Kubernetes versions and keep bounded node counts. Do not claim container snapshots include volumes or full VM memory. |

### 21–28 November: rehearsals and half demo

| Item | Details |
|---|---|
| 1 | All owners repair integration failures before new work. Two rehearsals: one on the main VM and one from a clean installation/second device. Keep a timed script and collect actual request/response outputs without secrets. Test cert expiry, routes and port mappings. |
| 2 | Half demo order: boot/doctor → account → create/key SSH → actions/reboot → ASG heal → Postgres table → scoped object upload → mail → supported LB/k8s/snapshot paths → replay/body/tenant evidence. Allow 25–35 minutes according to stable implemented scope. |

### 29 November–8 December: freeze and release candidate

| Item | Details |
|---|---|
| 1 | Finish half-demo defects and planned supported features by Dec 1. Then package app, foreground agent and CLI; verify downloads/checksums/install directions and CI. Foreground agent startup is the required release path; systemd is optional polish after it passes, not a blocker for five OS service managers. |
| 2 | Write threat model, incident runbook and pentest evidence based on actual implementation. Test master-key/data restoration and no secret leaks. No release-blocking auth, ownership, persistence or lifecycle defects may remain. |

### 9–15 December: stabilize and submit

| Item | Details |
|---|---|
| 1 | Dec 9–11: clean-machine/second-device final run and documentation review. Dec 12: reserve repair slot. Dec 13–14: critical fixes and submission assembly only. Tag the tested commit after required checks pass; don't tag an untested tree because the calendar says release day. |

## 6. Daily execution and escalation

| Item | Details |
|---|---|
| 1 | Standup: yesterday's accepted result, today's acceptance and blocker. Each owner commits a small reviewable slice; every behavior/security change has meaningful checks. Contract PRs require Anuj and B plus C for auth changes. No direct push to main. |
| 2 | A blocker lasting 90 minutes triggers pairing with the relevant owner. After one working day, reassign help and report which gate is affected. Friday demo runs from documented commands; documentation does not substitute for actual evidence. |
| 3 | Track tasks as: Not started / In progress / Implemented awaiting validation / Accepted / Blocked with owner and next action. Do not mark Accepted based on generated code alone. Keep one defect list in the repository and map it to the 38 audit IDs. |

## 7. Submission checklist

| Item | Details |
|---|---|
| 1 | [ ] Known tested commit and pinned version matrix; reproducible build/CI outputs. |
| 2 | [ ] No runtime auth bypass; negative tests and account/agent ownership evidence. |
| 3 | [ ] Complete core lifecycle and ASG outage/slow-create tests. |
| 4 | [ ] Persistent DB/object data, backup/restore and protected master key. |
| 5 | [ ] Correct public/private routing and second-device DNS/TLS/client endpoints. |
| 6 | [ ] LB/k8s/snapshot supported subsets demonstrated or limitations explicitly declared. |
| 7 | [ ] App/agent/CLI artifacts with checksums, install/setup/recovery instructions. |
| 8 | [ ] Threat model/pentest report reflect measured controls, not full CIS/parity claims. |
| 9 | [ ] Team can explain every major boundary and failure behavior without relying on AI output. |
| 10 | [ ] Final archive/report assembled by Dec 14; submit Dec 15. |
