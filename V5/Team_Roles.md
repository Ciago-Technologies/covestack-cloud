# CoveStack V5 — Ownership, Reviews and ITISS Responsibilities

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **15 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Five-person plan

| Person         | CORE-5 duties                                                                                                            | Extended duties                                                       |
| -------------- | ------------------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------- |
| A — Anuj       | ALL agent Go, container/k3d drivers, Docker network helper, reconcile, signing, pkg/kanclient, CI, Ansible Part 1, ITISS | Proxy-config/health/snapshot helpers only after all CORE-5 gates pass |
| B              | Java core, migrations, ALL lease/fencing/idempotency SQL, quota records and stack record                                 | ASG controller and extended records                                   |
| C              | Auth/security; owner of ALL acceptance tests from day 1; Java IPAM allocator                                             | Domain verification and metadata policy; E1–E8 tests                  |
| D              | Java /v2 endpoints including Kubernetes cluster subset, DTO mappings, API/schema review, Terraform fixtures; no Go       | LB/snapshot API mappings                                              |
| E — Go learner | cmd/covestack thin cobra API client, -o json, cross-compile packaging; later thin Terraform provider                     | React read-only list page; Ansible inventory Part 2                   |

**Risk: A is the bottleneck** for agent, two drivers, reconcile/fencing integration, infrastructure and ITISS. Mitigation: B owns all SQL, C owns tests/allocator, E owns CLI, D owns API mappings. Extended work must never starve CORE-5 reviews. No responsibility moves silently.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 2. Anuj's current sequence and Go ownership

A writes and can explain all agent Go. AI-generated code is reference only: A types it and explains its error paths. E touches only cmd/covestack and the later thin Terraform layer; never agent, signing, drivers or reconcile. Any pkg/kanclient change needs A review and A+B+C pairing.

| Priority | Work                                                                     | Acceptance                                                    |
| -------- | ------------------------------------------------------------------------ | ------------------------------------------------------------- |
| 1        | Review config; signed client/poll loop; decide module layout/signer move | Real auth-smoke; authenticated account/agent binding reviewed |
| 2        | Container driver and reconcile                                           | Tests 1–6, crash/fencing/idempotency                          |
| 3        | VPC/IPAM integration                                                     | C allocator; A network helper; test 14                        |
| 4        | Required k3d driver                                                      | Tests 7–10; host cap 2 clusters                               |
| 5        | PostgreSQL and MinIO process lifecycle                                   | Tests 11–13; Java bucket reconciliation                       |
| 6        | One stack and installer evidence                                         | Test 15; second Ansible run changed=0                         |

Source presence is not Gates passed. Follow the dated [workflow](CoveStack-Project-Workflow.md).

## 3. Java core and contracts — B

| Duty                                     | Acceptance evidence                                                                      |
| ---------------------------------------- | ---------------------------------------------------------------------------------------- |
| Main/configuration and Flyway migrations | Empty PostgreSQL 17 migration and packaged boot                                          |
| Scoped resource/operation transactions   | Account predicates/composite FKs; tests 3–4                                              |
| Lease/fencing/idempotency SQL            | Epoch fence, equality generation ack, release both lease fields; progress never acks     |
| Port/quota reservations                  | Reserve before cluster create; release after cleanup; transactional two-cluster host cap |
| Stack record                             | One fixed-sequence integration; retry no duplicates; no DAG engine                       |
| Persistent-data records                  | Owned volume/backup references and measured restore                                      |

No live schema exists yet. Schema migration stages are proposal B finalizes; C tests, A agent-column review, D DTO review.

## 4. Authentication and security — C

C owns every acceptance test from day 1, including tests 1–15 and E1–E8. Implement hash-only api_tokens and encrypted service-kind signing_credentials; custody is open B+C. Keep expiry/revocation, replay and account enforcement distinct.

| Boundary                       | Required evidence                                                    |
| ------------------------------ | -------------------------------------------------------------------- |
| User versus service credential | Kinds cannot be substituted; no public unrestricted key bootstrap    |
| Object/action access           | Other-account API 404; no existence leak                             |
| Replay/signing                 | Duplicate rejected; invalid MAC cannot poison nonce                  |
| Shared object storage          | A key denied on B bucket; Java recreates buckets after MinIO restart |
| IPAM                           | Ten allocations in one VPC; ten account VPCs, no overlap             |
| Audit                          | Unauthenticated event option and operator-only access reviewed B+C   |
| MinIO administration           | Custody, rotation and server-to-MinIO private path decided B+C       |

Agreed target: one agent process per physical host; separate service credential and authenticated polling context per assigned account. B owns server assignment validation; C credential provisioning/rotation/revocation; A isolated clients and shared bounded execution. Provisioning/storage details still need B+C+A review; no new wire fields. An ownership label is not authentication. Extended domain/metadata/automation policies remain in the roadmap.

## 5. Provider-facing integration — D

D writes Java only: /v2 endpoints, including the bounded Kubernetes cluster subset, DTO mappings, API/schema review and Terraform test fixtures. B owns core semantics; E writes the later thin Go provider and A reviews it. Optional Terraform covers Droplet and Kubernetes resources only; API authentication/ownership, refresh/import and explicit unsupported responses need evidence.

B+D finalize internal uuid, prefixed public_id for CLI and numeric provider_id for /v2. Do not assert full provider compatibility. Extended LB/snapshot mappings wait for dependency gates.

## 6. CLI and packaging — E

| Area         | Requirement                                                           |
| ------------ | --------------------------------------------------------------------- |
| CLI          | Thin cobra API client through pkg/kanclient only; never Docker/k3d/DB |
| State/config | No local DB; server URL/key ID/profile only; secrets env/keychain     |
| Mutations    | Idempotency key; create pending; --wait polls with finite timeout     |
| Get/list     | Desired/converged generation, lease/observed state; stable -o json    |
| Packaging    | linux/amd64, darwin/arm64, windows/amd64; two actual OS demos         |
| Learning     | Weekly W1–W4 review; W3 K8s mock-only                                 |
| Terraform    | Thin provider layer after CLI; Droplet/K8s only                       |

No UI work during CORE-5. Read-only React list page and Ansible inventory are Extended track; mail/outbox remains deferred until a test requires it. Preserve authorized kubeconfig export without overwriting existing files.

## 7. ITISS evidence ownership

Anuj leads VM/resource budget, privileges, firewall/isolation, signing/replay validation, CI failure evidence, installer hardening and backup/restore exercises. C reviews policy and independent tests; B/D implement needed server boundaries. DNS/TLS/tunnel, advanced metrics/incidents and inventory evidence remain Extended track targets.

**Open item:** Ask mentor in writing whether ITISS students may own the Go agent/driver code; if not, name an agent developer and a handover. Confirmation is pending. No informal assumption transfers agent work to E.

## 8. Required pair reviews

| Change                                 | Reviewers                 |
| -------------------------------------- | ------------------------- |
| Any Go change                          | A                         |
| pkg/kanclient or signing/wire contract | A+B+C                     |
| k3d driver                             | A+B+C                     |
| Schema change                          | B+C+A; D DTO alignment    |
| IPAM/network helper                    | C+B+A                     |
| Provider compatibility                 | D+B; A reviews Go; C auth |
| Stack/data retention                   | B+A+C; E CLI              |
| Extended domain/metadata/automation    | C+B+A; E interface        |

Extended tasks never delay CORE-5 reviews.

## 9. Evidence and task review

A reviewable task includes behavior, meaningful tests, sanitized evidence and a PR. A service status changes to Gates passed only with all G1–G4 results. Source files, mocks and screenshots alone are insufficient. Never write that all services are complete.

[Service gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Extended roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
