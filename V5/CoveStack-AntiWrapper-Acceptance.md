# CoveStack V5 — Anti-Wrapper Acceptance and CORE-5 Viva

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Scope and wrapper definition

> **Wrapper definition:** A CLI/API handler that directly runs docker run, k3d cluster create or writes proxy configuration without persisted desired state fails this acceptance design. Execution belongs behind durable intent and reconciliation, not in user-interface handlers.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

C owns ALL acceptance tests from day 1. The owner column also identifies implementation/pairing owners. All tests below are targets; no passing results are asserted.

## 2. CORE-5 acceptance — SRS tests

| Test                    | Exact assertion                                                                                                                                          | Owner                                       | Runs in CI      | Schema constraint relied on                                       |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- | --------------- | ----------------------------------------------------------------- |
| 1 — Container crash     | Kill -9 agent after container create, before acknowledgment; restart: exactly one owned container and target convergence.                                | C owns tests; A/C pairing                   | Required target | Resource ID/name uniqueness; generation and fenced acknowledgment |
| 2 — Fencing             | A claims; lease expires; B claims; A acknowledges: reject A without changing convergence.                                                                | C owns tests; B/C pairing                   | Required target | Lease owner/epoch match; expiry decision B                        |
| 3 — Tenant              | A GET/DELETE on B resource returns 404; no side effect or existence leak.                                                                                | C owns tests; C pairing                     | Required target | Account predicates and composite ownership FKs                    |
| 4 — Idempotency         | Same key ×10 parallel creates yields one resource/ID; different body with same key returns 422.                                                          | C owns tests; B/C pairing                   | Required target | Account/key PK, fingerprint, one transaction                      |
| 5 — Stale generation    | Report generation 3 after desired generation 4: ignored; acknowledgment and current observed state unchanged.                                            | C owns tests; B/C pairing                   | Required target | Acknowledgment equality; separate guarded progress statement      |
| 6 — Server down         | Stop server: agent creates and deletes nothing while disconnected.                                                                                       | C owns tests; A/C pairing                   | Required target | DB authoritative intent; no cleanup from failed poll              |
| 7 — Kubernetes crash    | Kill -9 agent mid cluster create; restart: exactly one deterministic cluster and target convergence.                                                     | C owns tests; A/C pairing                   | Required target | Cluster identity, reserved port and proposed labels               |
| 8 — Kubernetes tenant   | B fetches A kubeconfig: 404; kubeconfig absent from list/get/observations.                                                                               | C owns tests; C/D pairing                   | Required target | Account FK/predicate; kubeconfig reference only                   |
| 9 — Kubernetes cleanup  | Delete cluster: no owned container, volume or network with kan.resource=<id> remains.                                                                    | C owns tests; A/C pairing                   | Required target | Proposed label contract and ownership inventory                   |
| 10 — Kubernetes quota   | Third cluster on host: documented quota error; existing clusters unchanged.                                                                              | C owns tests; B/C pairing                   | Required target | Host cap transaction lock and derived count                       |
| 11 — PostgreSQL restore | Back up known rows, delete test data, restore: rows and canonical dataset checksum match; B cannot reach A DB.                                           | C owns tests; A/B/C pairing                 | Required target | Owned volume/backup/credentials and private routing               |
| 12 — MinIO tenant       | A access key on B bucket: denied; API requests for B resource by A return 404.                                                                           | C owns tests; C/B pairing                   | Required target | Scoped bucket credentials and account predicates                  |
| 13 — MinIO rebuild      | Restart MinIO; Java server reconciles bucket and credential policy from DB rows, without duplicate or cross-account binding.                             | C owns tests; B/C pairing                   | Required target | storage.bucket is server-owned; ciphertext custody TBD B+C        |
| 14 — VPC allocation     | 10 parallel IP allocations in one VPC: no duplicates; 10 accounts create VPCs concurrently: no CIDR overlap; A cannot see B network; no user CIDR input. | C owns tests; C/B; A network checks pairing | Required target | Subnet exclusion, live-IP uniqueness, one VPC/account             |
| 15 — Stack              | One CLI command provisions PostgreSQL + MinIO + app Droplet; partial failure recorded; retry resumes without duplicate children.                         | C owns tests; B/C/E; A execution pairing    | Required target | Stack/role uniqueness and idempotent child records                |

Test 14 does not permit custom user CIDRs: overlap is exercised against allocator/database proposals. Tests 11/12 include protocol/network denial alongside the account-scoped API 404 rule. K8s quota error name/status is **TBD with B+D**; do not invent an error wire string. Proposed `kan.account`, `kan.resource`, `kan.agent` labels need coordinated review; they are not present compatibility.

## 3. Reference concurrency design

**Reference design, B finalizes; not executable, not migration-ready.** Resources propose generation, acked_generation, lease_owner, lease_epoch, lease_expires_at and deleted_at. Exact schemas/SQL are in [Database Schema](CoveStack-Database-Schema.md).

| Flow                 | Required invariant                                                                                                                                                        |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Claim                | Verified credential account/agent context, not caller-selected account; multi-account polling TBD B+C+A; SKIP LOCKED, bump epoch, reference 30-second lease               |
| Acknowledge          | id, lease_owner, lease_epoch match AND generation = :gen; zero rows rejected; accepted ack clears lease_owner and lease_expires_at together, leaves lease_epoch unchanged |
| Progress             | Separate guarded statement; never touches acked_generation                                                                                                                |
| Expiry               | Epoch is the fence; expiry predicate decision remains with B                                                                                                              |
| Idempotent create    | Account/key reservation, resource, operation/event and response in ONE transaction                                                                                        |
| Fingerprint mismatch | Same key and different normalized request body: 422                                                                                                                       |
| Retry                | Refresh/resume durable operation, not duplicate runtime create                                                                                                            |

## 4. CORE-5 live demo and recorded evidence

Live highlights: tests 1, 3, 4 and the CLI get view showing desired versus converged generation, lease and observed state. Remaining tests are CI output/recorded video. Record a backup video; do not rely on hotspot access alone. CLI cross-build/run evidence uses two actual OSes, not compilation alone.

## 5. CORE-5 viva questions

| Question                                 | Honest answer — target, not current claim                                                                                       |
| ---------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| 1. Crash mid-create?                     | DB intent survives. Ensure inspects deterministic owned identity before another create; test 1 must prove one container.        |
| 2. Two agents?                           | Lease epoch changes on reclaim. An old owner/epoch acknowledgment is rejected; test 2 proves it.                                |
| 3. Why not just Docker?                  | Docker is execution. Our target is account-scoped DB intent, fencing and crash-safe lifecycle; claim evidence only after gates. |
| 4. Why not just k3d?                     | k3d executes a cluster request. Our target adds tenancy, port reservation, quotas, reconciliation and scoped kubeconfig.        |
| 5. Why not containers for data services? | Execution is reused. Our target adds volume ownership, measured restore and per-account storage credentials.                    |
| 6. Is a Droplet a VM?                    | No. CORE-5 compute is container-backed on one host. It has no independent-kernel VM boundary.                                   |
| 7. A reads B kubeconfig?                 | API must return 404; kubeconfig stays out of list/get/observations. Test 8 supplies evidence.                                   |
| 8. Server down?                          | Agent performs no new create/delete. No poll result means no authoritative deletion intent.                                     |
| 9. Why only two clusters?                | The 8 GB single-host budget is constrained. The per-cluster estimate is unverified until the host spike.                        |
| 10. Why no direct CLI DB access?         | Authenticated API centralizes policy/idempotency. CLI stores configuration only, not resource state.                            |
| 11. Which services are Gates passed?     | None are promoted by this document. Inspect the current gate register and C's CI evidence.                                      |
| 12. What is not implemented?             | This is target design, not a certificate. Missing gates and contract decisions are recorded explicitly.                         |

## Appendix A. Extended track tests — not SRS acceptance

Extended features may be shown only as bonus after Gates passed. These tests do not enlarge the SRS.

| Test                   | Exact assertion                                                                                            | Owner                               | Runs in CI      | Schema constraint relied on                        |
| ---------------------- | ---------------------------------------------------------------------------------------------------------- | ----------------------------------- | --------------- | -------------------------------------------------- |
| E1 — ASG               | Delete an active desired member: replacement respects cooldown/min/max; server down: no scaling.           | C owns tests; B/C/A pairing         | Required target | Parent ownership and DB policy                     |
| E2 — Load balancing    | Kill backend: health removes it; delete proxy config: regenerate from DB rows.                             | C owns tests; D/B/C/A pairing       | Required target | Owned backend FKs; observations are derived        |
| E3 — Snapshot          | Snapshot, change file, restore new Droplet: original hash; B cannot read snapshot.                         | C owns tests; D/A/C pairing         | Required target | Account resource/image binding                     |
| E4 — Metadata          | A cannot read B metadata; no agent/cloud secrets returned.                                                 | C owns tests; C/A/B pairing         | Required target | Bound self-identity                                |
| E5 — Publication       | Unknown hostname: 404; resource deletion removes route; public route cannot reach private management port. | C owns tests; A/B/C pairing         | Required target | Unique owned hostname and approved resource        |
| E6 — Domains           | Unverified domain serves nothing; reassignment never serves old tenant.                                    | C owns tests; C/B/A pairing         | Required target | Claim/verification revision binding                |
| E7 — Monitoring        | Stop agent: display stale/unknown, never healthy; alert fires and is recorded.                             | C owns tests; C/B/A pairing         | Required target | Freshness thresholds TBD with B; recorded evidence |
| E8 — Ansible inventory | Bob inventory contains none of Alice Droplets; inventory uses CLI JSON/API, never docker ps.               | C owns tests; E/C; A review pairing | Required target | API account predicate                              |

React has an API-only read-only list/interface gate plus relevant account tests; it creates no tables. E1–E8 are the eight named extended acceptance tests, not nine separate service numbers.

### Batch 2 checklist

Batch 2 added the exact heading **## 10. Known constraints**. The [SRS constraint target](CoveStack-Complete-Working-Setup.md#10-known-constraints) now resolves; this checklist records the cross-link correction, not application acceptance.

[Service gates](CoveStack-Service-Gates.md) · [Schema](CoveStack-Database-Schema.md) · [SRS](CoveStack-SRS-Scope.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [Go handover](CoveStack-Go-Handover.md).
