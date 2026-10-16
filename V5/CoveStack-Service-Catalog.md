# CoveStack V5 — Services, Capabilities and Release Boundaries

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **15 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. CORE-5 resource services — SRS scope

**A local cloud control plane with an account-scoped, crash-safe resource model. Docker and k3d are replaceable execution backends.**

| Service            | Bounded deliverable                                                                                | Limit                                                                | Status  |
| ------------------ | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------- | ------- |
| Compute/Droplets   | Container create/inspect/start/stop/delete; tenancy, generation, leases/fencing and crash recovery | Single host; containers, never VMs                                   | Planned |
| Kubernetes         | Required k3d cluster ownership, quota and separately scoped kubeconfig                             | Max 2 clusters/host; 1 server + max 2 agents; one tested pinned pair | Planned |
| Managed PostgreSQL | Owned volume, backup/restore, scoped credentials and private networking                            | PostgreSQL first; no HA or extra-engine promise                      | Planned |
| Object storage     | S3-compatible MinIO; per-account credentials and bucket scoping                                    | Selected client operations; Java bucket reconcile                    | Planned |
| VPC/IPAM           | One private bridge/account; simple allocation and DB leases                                        | Single host; no custom CIDR; pool TBD                                | Planned |

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

### 1.1 HLD/LLD dependency interpretation

Default private networking precedes workload execution; CLI delivery order is not permission to skip that dependency. Shared MinIO is one operator-owned runtime, with per-account buckets/credentials reconciled in Java. The fixed stack's storage child is a bucket, never ownership of the shared process. Claims require reviewed correlation/renewal schemas, and periodic inspections repair converged drift without changing requested intent. Runtime exclusion is separate from database epoch fencing. See [HLD/LLD](DocsV5-DO_Team.md#22-hld--control-execution-and-data-paths).

## 2. Extended track and optional future work — relocation index

The former original-service and extension categories do not determine SRS membership. Load balancing, ASG and snapshots move out of deliverables; VPC/IPAM moves into CORE-5. The scope limits below are explicit user decisions, not evidence of implementation.

| Catalog item               | Classification / authoritative destination                                                                                                       |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| Load balancing             | Extended track; [roadmap section 4](CoveStack-Extended-Roadmap.md#4-load-balancing)                                                              |
| Auto scaling groups        | Extended track; [roadmap section 3](CoveStack-Extended-Roadmap.md#3-auto-scaling-groups)                                                         |
| Snapshots                  | Extended track; [roadmap section 8](CoveStack-Extended-Roadmap.md#8-snapshots)                                                                   |
| Metadata                   | Extended track; [roadmap section 6](CoveStack-Extended-Roadmap.md#6-metadata)                                                                    |
| Application publication    | Extended track; [roadmap section 5](CoveStack-Extended-Roadmap.md#5-application-publication); simulated public IP/direct IP browsing attach here |
| Domains                    | Extended track; [roadmap section 9](CoveStack-Extended-Roadmap.md#9-domains)                                                                     |
| Monitoring                 | Extended track; [roadmap section 7](CoveStack-Extended-Roadmap.md#7-monitoring)                                                                  |
| React UI                   | Extended track; [roadmap section 10](CoveStack-Extended-Roadmap.md#10-react-ui); read-only list first                                            |
| Ansible inventory (Part 2) | Extended track; [roadmap section 11](CoveStack-Extended-Roadmap.md#11-ansible-inventory-part-2)                                                  |

Cross-cutting CORE-5 items are Security, Recovery, CLI first, optional Terraform for Droplet/K8s only, ONE Postgres + MinIO + app Droplet stack integration and Ansible Part 1 operator install. Broader customization/DAG and parity goals remain deferred in the roadmap; they are not extra deliverables.

## 3. Publication scope — relocated

The earlier default-publication design is an Extended track target, not a CORE-5 feature. Its shared gateway, explicit exposure, approved application port and HTTPS rules are preserved in the [roadmap](CoveStack-Extended-Roadmap.md#5-application-publication). No simulated address routes internet traffic.

## 4. CORE-5 resource information

| Resource          | Safe list/get information                                                                               | Protected retrieval                  |
| ----------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------ |
| Droplet           | Public ID, image, desired/converged generation, observed state, lease and approved private-network info | SSH material by policy               |
| Kubernetes        | Public ID, version/node shape, generations and observed state; no kubeconfig                            | Scoped kubeconfig export only        |
| Database          | Public ID, engine, private endpoint, volume reference and generations                                   | Scoped password/credential reference |
| Bucket            | Public ID, name, endpoint/policy summary and generations                                                | Scoped access credentials            |
| VPC/IPAM          | Owned CIDR, allocations and generations; state TBD with B                                               | No ordinary secret                   |
| Stack integration | Template version, child IDs and progress                                                                | Authorized secret references only    |

These are control-plane views; they do not add a workload metadata service. Internal uuid, CLI prefixed public_id and numeric /v2 provider_id are B+D decisions. No credential or kubeconfig is placed in ordinary spec/observed/list/get.

## 5. Deliverable acceptance

| CORE-5 service     | Failure/acceptance tests |
| ------------------ | ------------------------ |
| Compute/Droplets   | 1–6                      |
| Kubernetes         | 7–10                     |
| Managed PostgreSQL | 11                       |
| Object storage     | 12–13                    |
| VPC/IPAM           | 14                       |

Test 15 is the one integration stack demo. See [SRS scope](CoveStack-SRS-Scope.md), [gate register](CoveStack-Service-Gates.md) and [exact assertions](CoveStack-AntiWrapper-Acceptance.md). Gate evidence, not a scheduled date, determines status.

## 6. Reuse and differentiation

Docker already offers IPAM/network primitives; Compose/Swarm support coordinated workloads; Terraform/Ansible can manage Docker. CoveStack adds its selected account/resource/stack contract and integrated workflow. It is not original merely because forms replace YAML. Evaluate whether the complete workflow helps the intended shared lab.

## 7. Optional future work — relocated

The existing nested Docker, VM-backed Droplet, multi-host VPC, additional DB engine and full-provider-parity designs are **Optional future work** in [roadmap section 12](CoveStack-Extended-Roadmap.md#12-optional-future-work--not-additional-extended-services). No platform Docker socket is exposed to user workloads. This classification preserves the catalog items without expanding the SRS.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
