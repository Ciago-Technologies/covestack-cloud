# CoveStack V5 — Anuj Implementation and ITISS Work Order

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **15 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Current starting point

Config/signing/lifecycle source and auth-smoke are inspection evidence only; no service is Gates passed from file presence. The next A work is reviewed client/polling, production agent and generic container reconciliation. Do not reapply old patches or rename compatibility packages. No code changes happen in this documentation task.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 2. Immediate implementation sequence

| Order | A work                                                  | Acceptance/review                                                     |
| ----- | ------------------------------------------------------- | --------------------------------------------------------------------- |
| 1     | Review local config/tests; schema and signed-poll scope | Implement agreed one-process/per-account-client model after assignment/custody review; no cross-account shortcut |
| 2     | Plan module layout and signer move into pkg/kanclient   | A+B+C review; move not performed now                                  |
| 3     | Typed HTTP client and injected-transport tests          | Finite timeout, cancellation, no redirects, bounded body              |
| 4     | Production heartbeat/poll/shutdown main                 | Real server/auth-smoke; advertise only installed reviewed kinds       |
| 5     | Generic container driver and reconcile                  | Tests 1–6 with C; B owns all lease/fence/idempotency SQL              |
| 6     | Docker network helper inside container driver           | C Java allocator; B records; test 14; attachment contract change B+A  |
| 7     | k3d CLI driver after early host spike; revalidate in November                | Tests 7–10; max 2 clusters/host; reserved port and scoped kubeconfig  |
| 8     | PostgreSQL and MinIO process/volume lifecycle           | Tests 11–13; Java reconciles buckets/credentials                      |
| 9     | One stack integration; Ansible Part 1 and evidence      | Test 15; installer second run changed=0                               |

Only two workload drivers: container and k3d. Generic container inputs are image/env/volume/labels/health. Proposed kan.* reserved labels require approval. Unknown states/wire fields remain TBD with B; contract 1.1.0 wins.

A types and can explain agent Go; AI-generated code is reference only. E owns thin CLI, never agent/signing/drivers/reconcile. Follow CLI service order droplet → vpc → k8s → db → bucket → stack, waiting for the previous failure test in CI.

### 2.1 HLD/LLD changes to your first slices

Run the total-host capacity spike at first host preparation, not after committing the combined workload shape. Before mutation, pair with B+D/C on claim/renew/report and separate signed network-control schemas; existing 1.1.0 fields do not define the full protocol. Do not add wire fields alone.

Acquire host OS exclusion, serialize resource mutations, retain child-operation exclusion and record durable attempts with B. A timeout does not prove a backend stopped. Ambiguous old execution blocks conflicting takeover until settled; DB epoch rejection alone is insufficient. Renew only current unexpired claims and stop new mutations when authorization is lost.

Default VPC execution is needed for the first real Droplet. Add periodic inspection of converged rows, using fresh claims and no health acknowledgment. Repair drift with unchanged desired generation; stale inspection cannot overwrite newer intent. For MinIO, execute one operator-owned shared process/volume from reviewed intent/host secret references; Java owns tenant buckets and the stack owns a bucket child, not the shared process. Preserve the fixed stack/test 15 and installer; optional/extended work is cut first.

These are Planned implementation duties. [HLD](DocsV5-DO_Team.md#22-hld--control-execution-and-data-paths) and [LLD](CoveStack-Database-Schema.md#7-concurrency-and-reference-queries) define the precise dependencies and failure evidence.

## 3. Client rules

| Concern         | Requirement                                                        |
| --------------- | ------------------------------------------------------------------ |
| Context/timeout | Caller cancellation, finite HTTP timeout                           |
| Redirects       | Refuse forwarding signed credentials                               |
| Request bytes   | Marshal once; sign/transmit identical bytes                        |
| Response size   | Bounded read; detect overflow                                      |
| Status/JSON     | Non-2xx and malformed data are errors                              |
| Retry           | Fresh signed request; mutations need agreed idempotency            |
| Logging         | No secret structs, canonical real-key payloads or full kubeconfigs |

Coordinate endpoints, retry/idempotency, agent leases and readiness with B. Test using synthetic credentials; a fake verifier is not the deployed auth/store.

## 4. First compute acceptance

Create one owned container-backed Droplet, inspect by approved labels before create, restart after mid-create crash, converge without duplicates and delete only the owned resource. Preserve volumes/unrelated resources. Never describe Droplets as VMs. Server outage or incomplete/empty poll must not trigger cleanup.

Epoch is the fence; convergence ack requires exact generation/owner/epoch and clears lease_owner plus lease_expires_at together, leaving epoch unchanged. Progress is separate and never changes acked_generation. A reviews agent integration; B owns SQL; C owns tests. Lifecycle success mapping uses existing contract strings only.

Agreed target: one agent process per physical host; separate service credential and authenticated polling context per assigned account. B owns server assignment validation; C credential provisioning/rotation/revocation; A isolated clients and shared bounded execution. Provisioning/storage details still need B+C+A review; no new wire fields.

## 5. ITISS parallel track

| CORE-5 phase | Evidence                                                                             |
| ------------ | ------------------------------------------------------------------------------------ |
| Foundation   | AlmaLinux VM/version/resource budget, threat model, real signing/replay tests and CI |
| Compute      | Privilege/mount review, crash/fence/isolation tests with C                           |
| Network      | Pool/allocator decision, private firewall and actual denied paths                    |
| Kubernetes   | One-day host spike: cgroup v2, SELinux, firewalld, memory and tested k3d+k3s flags   |
| Data         | Owned volumes/credentials; measured backup/restore and cross-account denial          |
| Install      | Vault/no_log/checksums/systemd hardening; second run changed=0                       |

Publication/TLS/tunnel, metadata, metrics/alerts/incidents and dynamic inventory remain Extended track evidence in the roadmap. A begins extended Go only after all CORE-5 gates pass. If A is one week behind, all extended work stops.

Ask mentor in writing whether ITISS students may own agent/driver Go; if not, name a developer and record handover. Confirmation remains pending.

## 6. Learning sequence

Study structs/pointers/interfaces/errors, HTTP/JSON/readers, contexts/timers/signals, tests and then concurrency. Inspect one fixture/result before coding. Use Go documentation and the prior learning guide if present; V5 does not replace code-level walkthroughs.

## 7. Dated targets

| Period                          | A checkpoint                                                      |
| ------------------------------- | ----------------------------------------------------------------- |
| 14–31 October 2026              | Client/main/poll/auth-smoke and schema review                     |
| 1–20 November 2026              | Container/reconcile/fencing integration; revalidate early capacity spike   |
| 21 November–5 December 2026     | Tests 1–6 in CI and network helper/allocator integration          |
| 6–20 December 2026              | Required bounded K8s and tests 7–10                               |
| 21 December 2026–5 January 2027 | PostgreSQL/MinIO process lifecycle and restore with C/B           |
| 6–10 January 2027               | One stack and Ansible Part 1                                      |
| 11–25 January 2027              | Freeze, clean-VM test, two-OS CLI evidence, viva and backup video |

No new CORE-5 features after 10 January. Never cut tests 1–15; cut extended work first and report any unmet gate honestly. Final deadline: Monday, 25 January 2027.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
