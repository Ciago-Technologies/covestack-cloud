# CoveStack V5 — CORE-5 SRS Scope for Integration

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Status and purpose

This scope file is to be merged into the future SRS; no existing SRS/requirement-ID scheme was found. Do not invent requirement IDs. All five services are **Planned**. V5 is documentation version; lifecycle contract remains 1.1.0.

**A local cloud control plane with an account-scoped, crash-safe resource model. Docker and k3d are replaceable execution backends.**

> A handler that directly executes runtime create commands without persisted desired state fails this SRS acceptance design. Resource state is authoritative in the DB, not execution objects or client configuration.

## 2. Exactly five services

| Service | Scope | Status |
|---|---|---|
| Compute/Droplets | Container-backed, single host; never a VM | Planned |
| Kubernetes | Required k3d; max 2 clusters/host; 1 server + max 2 agents; one pinned k3s version | Planned |
| Managed PostgreSQL | Backup/restore, scoped credentials, private network | Planned |
| Object storage | S3-compatible MinIO, per-account credentials, scoped buckets | Planned |
| VPC/IPAM | One private network/account; simple allocator; no custom CIDR | Planned |

## 3. Cross-cutting items — not additional services

| Item | Bound |
|---|---|
| Security | Authenticated API, account predicates, ciphertext/hash custody |
| Recovery | DB desired state/generation; crash tests and measured restore |
| CLI first | Go/cobra API client; no UI work during CORE-5 |
| Terraform optional | Droplet and Kubernetes resources only |
| One integration stack | PostgreSQL + MinIO + app Droplet; fixed sequence, no DAG engine |
| Ansible Part 1 | Install playbook and ITISS evidence |

## 4. Five deliverables and one integration demo

| Deliverable | G1 DB intent | G2 crash/rebuild | G3 isolation | G4 failure evidence |
|---|---|---|---|---|
| Compute/Droplets | Resource generation | Inspect before create | Other-account 404 | Tests 1–6 |
| Kubernetes | Cluster/port reservation | Recover half-create | Kubeconfig 404 | Tests 7–10 |
| Managed PostgreSQL | Owned resource/volume | Reattach/restore | API 404/private DB denial | Test 11 |
| Object storage | Bucket/credential rows | Java reconciles after restart | API 404/S3 denial | Tests 12–13 |
| VPC/IPAM | Network intent/IP leases | Rebuild helper configuration | Network 404 | Test 14 |

Stack demo: test **15** proves one-command provisioning, recorded partial failure and retry without duplicates. It is integration of the five-service deliverable, not a sixth service.

## 5. Gates and CLI boundary

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

The CLI talks only to the authenticated API through pkg/kanclient. No Docker/k3d calls, direct DB, or local resource database. Config contains server URL/key ID/profile; secrets come from env/keychain. Every mutation uses an idempotency key. Create returns pending immediately; --wait polls convergence. Get shows desired/converged generation, lease and observed state; every list/get supports stable -o json.

## 6. Assumptions and constraints

Single 8 GB AlmaLinux host; two-cluster cap and node limits; a phone hotspot without institute LAN/router administration; one Go developer and one Go learner. Cluster memory estimate is unverified until the host spike. Use CI/video and two-OS CLI evidence; no reliance on a live hotspot alone. See [Known constraints](CoveStack-Complete-Working-Setup.md#10-known-constraints) (section added in Batch 2).

Limits communicated to mentor on <date: fill in>. Mentor confirmation: pending. No cluster version, subnet pool, key custody, missing enum or quota-error wire string is invented; open decisions remain with the named owners.

## 7. Out of scope for this SRS

Further services are in the development roadmap (CoveStack-Extended-Roadmap.md) and are not requirements or deliverables of this SRS.

## 8. Acceptance references

[Tests 1–15](CoveStack-AntiWrapper-Acceptance.md#2-core-5-acceptance--srs-tests) · [Gate register](CoveStack-Service-Gates.md) · [Schema reference](CoveStack-Database-Schema.md) · [Go ownership](CoveStack-Go-Handover.md).
