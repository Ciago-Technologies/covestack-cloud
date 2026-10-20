# CoveStack V5 — Delivery Workflow to 25 January 2027

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **20 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Schedule authority

**Submission is Monday, 25 January 2027.** Planning date: 14 October 2026. The deliverable plan is CORE-5 only: Droplets, required Kubernetes, Managed PostgreSQL, Object storage and VPC/IPAM. Security, Recovery, CLI, optional Terraform, one stack demo and Ansible Part 1 are cross-cutting items, not extra services.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 2. CORE-5 phase — deliverable plan

| Period                 | Work                                                                                                                                                                                                                  | Acceptance evidence                                                                                |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| 14 Oct–31 Oct 2026     | B: server main and migrations from schema; C: auth lookup; A: signed client/main, host exclusion and early capacity spike; B+D/A/C claim/network protocol review; plan signer move; schema review B+C+A                                                                     | Auth-smoke against real server; clean-DB migration                                                 |
| 1 Nov–20 Nov 2026      | A: container driver, labels, reconcile/inspection, host exclusion; B/C default-VPC prerequisite; E: CLI skeleton/droplet; early one-day K8s host spike at first host preparation                                                                                        | Real container create/delete through CLI → API                                                     |
| 21 Nov–5 Dec 2026      | C: tests 1–6 in CI and Java VPC allocator; E: vpc CLI                                                                                                                                                                 | Tests green; ten parallel IP allocations in one VPC; ten accounts create VPCs without CIDR overlap |
| 6 Dec–20 Dec 2026      | A: k3d driver; B: port reservations/quota; scoped kubeconfig; C: tests 7–10; E: k8s CLI                                                                                                                               | Crash recovery; tenant denial; clean deletion; third cluster rejected                              |
| 21 Dec 2026–5 Jan 2027 | PostgreSQL and MinIO process on generic container driver; Java bucket/credential reconcile; restore; db/bucket CLI; D: /v2 subset; Terraform Droplet/K8s only if ahead; SSH key injection and reviewed SSH base image | Delete data, restore known rows/checksum; cross-account bucket key denied                          |
| 6 Jan–10 Jan 2027      | One Postgres + MinIO + app Droplet stack; stack CLI; Ansible Part 1                                                                                                                                                   | One-command stack with safe retry; playbook second run changed=0                                   |
| 11 Jan–25 Jan 2027     | FREEZE; teammate clean-VM install; viva rehearsal; backup video; honest relabel for any service without Gates passed                                                                                                  | Clean install and evidence review                                                                  |

These are target checkpoints, not current capability claims. Dates do not waive gates.

## 3. Stage-specific ownership

| Owner | CORE-5 priority                                                                                     | Review dependency                            |
| ----- | --------------------------------------------------------------------------------------------------- | -------------------------------------------- |
| A     | Agent, two drivers, network helper, reconcile, signing, pkg/kanclient, CI, Ansible Part 1 and ITISS | B SQL; C tests/security; D schema/API; E CLI |
| B     | Java core, all migrations/lease/fencing/idempotency SQL, quota and stack records                    | Schema review B+C+A                          |
| C     | Auth/security; all acceptance tests from day 1; Java IPAM allocator                                 | Agent/API integration with A/B/D             |
| D     | Java /v2 endpoints and DTOs; schema review; Terraform test fixtures                                 | No Go; B core semantics                      |
| E     | Thin cobra CLI, JSON and packaging; later thin Terraform layer                                      | A reviews every Go change                    |

E has weekly Go-learning reviews: W1 basics/cobra/three-OS builds; W2 HTTP/JSON/context/timeouts with mock API; W3 Droplet flows and mock-only VPC/K8s practice; W4 table tests/httptest and real-server CLI. Follow the [handover](CoveStack-Go-Handover.md); learning weeks are not new calendar promises.

## 4. Dependency order

CLI-first order: **droplet → vpc → k8s → db → bucket → stack**. Start the next production CLI service only after the previous service's failure test passes in CI. W3 K8s practice is mock-only.

Auth, durable intent and signed agent scope precede real reconciliation. VPC is server-owned; compute network attachment needs coordinated contract review B+A. Agreed target: one agent process per physical host; separate service credential and authenticated polling context per assigned account. B owns server assignment validation; C credential provisioning/rotation/revocation; A isolated clients and shared bounded execution. Provisioning/storage details still need B+C+A review; no new wire fields. Do not bypass that decision with caller-selected account parameters.

CLI, Terraform provider and Ansible inventory use only the authenticated API. CLI has no local DB and no Docker/k3d dependency; config is server URL/key ID/profile, secrets from env/keychain. Mutations send an idempotency key; create returns pending, --wait polls; get shows desired/converged generation, lease and observed state; list/get support -o json. No UI work during CORE-5.

### 4.1 Runtime prerequisites and early HLD spike

CLI delivery order does not postpone runtime dependencies. The first real Droplet needs the default-VPC claim/execution path; freeze that reviewed protocol before workload-driver dispatch. Shared MinIO bootstrap is one operator-owned process/volume, separate from tenant bucket children. Claim/report/renewal schemas and default-network execution are early B+D+A/C work, not late patches.

Run the one-day AlmaLinux/k3d capacity spike at first host preparation, before committing to the combined deployment shape. It measures all platform dependencies and bounded workload combinations, not only cluster idle RAM. Revalidate in November as integration changes; the early result does not claim a completed driver. The 2-cluster cap remains an upper limit; admission may reject earlier under measured CPU/memory/disk constraints.

The previous last-resort cuts of the installer/stack are superseded: both remain required under current scope, and tests 1–15 are never cut. If required work still cannot meet the deadline, report unmet gates and seek an explicit revised SRS; no mentor approval or scope change is assumed here.

## 5. Daily engineering workflow

| Step           | Expected artifact                                      |
| -------------- | ------------------------------------------------------ |
| Plan one slice | Behavior and acceptance case written first             |
| Branch         | Agreed reviewed baseline; no direct main push          |
| Implement      | Small change respecting ownership/contracts            |
| Verify         | Appropriate tests and actual integration evidence      |
| Review         | Boundary owners inspect behavior and errors            |
| Merge          | Required checks and approved PR                        |
| Record         | Feature matrix, evidence and known limitations updated |

Standup: last accepted result, today's acceptance target and current dependency. A blocker lasting a working day triggers pairing or reassignment, not silent schedule drift. Do not force-push teammates' branches or bypass policy to meet a gate.

### 5.1 Experience checkpoints — Planned, within existing dates

| Existing checkpoint | Required user experience | Owner and review |
| --- | --- | --- |
| First Droplet / tests 1–6 | Useful app response, two-account same-name denial, crash/restart and explainable get output | A/B/D/E; C owns tests |
| VPC / test 14 | Allocation races plus actual intended connectivity and cross-account network denial | C allocator/tests, A execution, B records |
| Kubernetes / tests 7–10 | Scoped kubeconfig runs a bounded workload and returns its result | A/B/D/E; C tests |
| PostgreSQL / MinIO / tests 11–13 | Actual row and exact-byte workflows; restore known rows and retain objects across restart | A/B; C tests |
| Fixed stack / test 15 | App uses both dependencies; partial-failure retry preserves identities | B/A/E; C tests |
| Freeze / installer | Teammate follows clean-host walkthrough; records versions, access path, failure diagnosis and restore | A installer/ITISS; C reviews evidence |

For every slice, document inputs/outputs, capability limits, unsupported behavior, readiness, retry, cleanup and retained data before implementation. Test through the authenticated API/CLI and real runtime when promised; record unit, integration and runtime evidence separately. This adds no services, schedule extension or test IDs. See [experience assertions](CoveStack-AntiWrapper-Acceptance.md#33-user-experience-assertions--planned-existing-tests).

## 6. Scope and gate failure policy

| Situation                          | Response                                                                                                                        |
| ---------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| Extended work competes with CORE-5 | Cut the entire extended track first                                                                                             |
| Still late after that cut | Cut optional Terraform K8s resource, then the rest of optional Terraform; preserve required CORE-5 services, fixed stack/test 15, Ansible Part 1 and security/recovery |
| CORE-5 service or tests late       | Never cut a CORE-5 service or tests 1–15; relabel honestly and report unresolved target                                         |
| A one week behind any checkpoint   | Stop all extended work                                                                                                          |
| After 10 January 2027              | No new CORE-5 features; repair, verification and evidence only                                                                  |
| Backend/auth unavailable           | Isolated/mock practice can continue; live acceptance remains pending                                                            |
| Cluster work too large             | Keep required bounded shape/version; never make Kubernetes optional                                                             |

No form, enum or screenshot passes a gate. No mentor acceptance is assumed.

## 7. Final demonstration — CORE-5 only

Five-minute CLI live highlights: test 1 agent crash/restart without duplicate, test 3 other-account 404, test 4 same-key same-ID, then get desired/converged generation, lease and observed state. Show remaining tests 2 and 5–15 as CI results or recorded video. Include required Kubernetes, PostgreSQL restore, MinIO denial/rebuild, VPC allocation and the stack integration evidence. Demonstrate CLI on two actual OSes using cross-built artifacts; keep a backup video and avoid depending on the hotspot alone.

Use [Setup](CoveStack-Complete-Working-Setup.md#9-clean-machine-release-acceptance) for the full evidence script.

## 8. Submission package

| Material        | Requirement                                                 |
| --------------- | ----------------------------------------------------------- |
| Source          | Reviewed tested revision with history                       |
| Binaries/images | Exact versions/digests and checksums                        |
| Documentation   | V5 plus actual runtime command updates                      |
| Test evidence   | Commit/date/environment/results and unresolved issues       |
| Security        | Threat model, policy matrix, remediation and residual risks |
| Recovery        | Repeated restore evidence with measured scenario            |
| Demonstration   | Script and backup recording; no secrets visible             |

## 9. Extended development track — not part of SRS or deliverables

Entry: a service starts only after its CORE-5 dependency has Gates passed. From **1 December 2026**, B/C/D/E may start server-side schema/API/controller work only when CORE-5 duties are not blocked. A starts extended Go work only after all CORE-5 gates pass. CORE-5 reviews always take priority; if A is one week behind, stop all extended work.

Order: **ASG → Load balancing → Application publication → Metadata → Monitoring → Snapshots → Domains → read-only list page → Ansible Part 2**. Each requires its corresponding E1–E8 acceptance evidence or interface gate; UI shares backend tests. After **25 January 2027**, this track continues as post-submission development. Bonus demos/viva mentions require Gates passed.

Requirements retained from the old schedule: external HTTPS ingress, UI/CLI stack parity, domain verification, metadata identity, LB/ASG/snapshot lifecycle, monitoring alerts, Terraform refresh/import and approved Ansible connectivity. Broader stack DAG/customization and nested Docker belong to Optional future work. The previous October/December extended milestones are superseded by dependency gates, not deleted requirements.

Historical scope responses retained: if custom domains are unavailable, report the gap rather than claim arbitrary-domain support; incomplete dependencies must not produce success-shaped stacks; nested Docker must not weaken privilege controls. Future UI create → CLI inspect, external-device publication, Terraform changes, inventory/configuration and separate restore demonstrations remain extended targets. No selected proxy/tunnel tool or date is asserted.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
