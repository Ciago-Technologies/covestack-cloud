# CoveStack V5 — Service Gates and Promotion Evidence

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Scope and promotion rule

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

Never write that all services are complete. Extended services are not SRS requirements or deliverables; a Gates passed extended feature may be shown only as a bonus. Preserve the distinction between status vocabulary and runtime contract states.

## 2. CORE-5 — SRS gate register

Every table/column below is proposed reference design, B finalizes. No live schema is asserted.

| Service | G1 DB intent | G2 rebuild from DB | G3 tenant test | G4 failure test | Owner | Status |
|---|---|---|---|---|---|---|
| Droplets | resources row and generation | Inspect proposed labels before create | B gets 404 | 1; related 2–6 | A/B; C tests | Planned |
| Kubernetes | Resource/cluster row; reserved API port | Find half-created deterministic cluster; finish or recreate same identity | B kubeconfig request gets 404 | 7, 9, 10; related 8 | A/B/D; C tests | Planned |
| Managed PostgreSQL | db_instances, owned volume and resource intent | Reattach volume; restore verified backup | API 404 plus forbidden DB path | 11 | A/B; C tests | Planned |
| Object storage | Bucket and scoped credential rows | Java recreates buckets/policy; MinIO process uses container driver | API 404 plus denied S3 key | 12, 13 | B/C; A runtime | Planned |
| VPC/IPAM | vpcs generation/intent; IP allocation rows | Rebuild bridge/leases; network execution helper | A cannot see B network: 404 | 14 | C allocator; B SQL; A execution | Planned |

VPC has no kind in contract 1.1.0. It is a server-owned design, not new agent work. Its state vocabulary and compute attachment message remain **TBD with B**; decision owners B+A. Bucket reconciliation stays server-side under `storage.bucket`.

## 3. Extended track — not in SRS

| Service | G1 DB intent | G2 rebuild from DB | G3 tenant test | G4/CI test | Owner | Status |
|---|---|---|---|---|---|---|
| Auto scaling groups | asg_groups, asg_instances; state/gen fields TBD with B where missing | Desired membership, reservations and scaling policy | Other-account API 404 | E1 | B; A after CORE-5 | Extended track |
| Load balancing | load_balancers, lb_backends, health_check_configs; state/gen fields TBD with B where missing | Backend membership and validated proxy configuration | Other-account API 404 | E2 | D/B; A after CORE-5 | Extended track |
| Application publication | publications; state/gen fields TBD with B where missing | Owned hostname-to-upstream routes | Other-account API 404 | E5 | B/C; A after CORE-5 | Extended track |
| Metadata | metadata_identities; state/gen fields TBD with B where missing | Self-identity bindings and permitted fields | Other-account API 404 | E4 | C/B; A after CORE-5 | Extended track |
| Monitoring | metric_samples, alert_rules, alerts; state/gen fields TBD with B where missing | Observed freshness and alert configuration; measurements are evidence | Other-account API 404 | E7 | C/B; A after CORE-5 | Extended track |
| Snapshots | snapshots; state/gen fields TBD with B where missing | Owned immutable image references and retention | Other-account API 404 | E3 | D/B; A after CORE-5 | Extended track |
| Domains | domains, domain_verifications; state/gen fields TBD with B where missing | Verified claims and current account bindings | Other-account API 404 | E6 | C/B | Extended track |
| React UI | No new tables; state/gen fields TBD with B where missing | Account-scoped list responses through API | Other-account API 404 | Interface gate; relevant tests 3/8 | E | Extended track |
| Ansible inventory (Part 2) | No new tables; state/gen fields TBD with B where missing | Owned inventory JSON through CLI/API | Other-account API 404 | E8 | E; A review | Extended track |

React and inventory are interfaces, not persisted services: their underlying CORE-5 resources meet G1–G4; the interface itself has the access/lifecycle checks below. Monitoring samples are measured evidence, never authoritative desired state. Publication/metadata/domain enums are TBD with B. See [roadmap](CoveStack-Extended-Roadmap.md).

## 4. Interface and cross-cutting gates

| Item | Access boundary | Evidence |
|---|---|---|
| CLI | Authenticated API only; no Docker/k3d/DB | Mutation idempotency, lifecycle/--wait, stable JSON, dependency scan |
| Terraform (optional) | API only; Droplet/Kubernetes resources | Create/read/update/delete/import; no direct runtime access |
| Ansible Part 1 | Operator installer; no platform DB as client inventory | Vault/no_log, checksum, lint; second run changed=0 |
| React (extended) | API only; list page first | Same account-scoped state and JSON; no secret/config authority |
| Ansible inventory (extended) | CLI JSON/authenticated API only | E8; no Docker/k3d import/call |
| One stack demo | Account-owned records; fixed sequence, no DAG engine | Test 15; cross-cutting integration, not sixth service |

## 5. Freeze relabel procedure

Before 11 January, C attaches CI results to each gate; B/A/D/E review their boundaries. Promote only when G1–G4 are evidenced. Otherwise retain Planned/In progress and state the missing gate. No new CORE-5 features after 10 January; no release claim is based on a form, fixture or reference DDL. Do not cut tests 1–15 to meet the deadline.

## 6. Links

[Acceptance](CoveStack-AntiWrapper-Acceptance.md), [schema](CoveStack-Database-Schema.md), [SRS scope](CoveStack-SRS-Scope.md), [Go handover](CoveStack-Go-Handover.md), [main design](DocsV5-DO_Team.md).
