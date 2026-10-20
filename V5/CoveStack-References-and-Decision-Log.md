# CoveStack V5 — References, Assumptions and Decisions

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **15 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Document authority

| Material                        | Authority                                               |
| ------------------------------- | ------------------------------------------------------- |
| V5 docs                         | New target architecture/schedule and agreed scope       |
| Current runtime contracts       | Authoritative wire/generator source; lifecycle 1.1.0    |
| Actual commits/tests/deployment | G1–G4 evidence for Gates passed                         |
| Old team guides/archives        | Historical design and findings                          |
| Mentor rubric                   | External acceptance requirement; obtain current version |

No official C-DAC marks allocation or no-coding rule is assumed. Planning date is 14 October 2026; deadline is Monday, 25 January 2027.

## 2. Primary source links

| Reference                                                                                                           | Purpose                             |
| ------------------------------------------------------------------------------------------------------------------- | ----------------------------------- |
| [Docker bridge](https://docs.docker.com/engine/network/drivers/bridge/)                                             | Local network primitives and limits |
| [Docker IPAM options](https://docs.docker.com/reference/cli/docker/network/create/)                                 | Existing address mechanisms         |
| [Docker security](https://docs.docker.com/engine/security/)                                                         | Daemon/workload authority           |
| [k3d exposure](https://k3d.io/stable/usage/exposing_services/)                                                      | Ingress integration                 |
| [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework)                                   | Provider development                |
| [Ansible Docker](https://docs.ansible.com/projects/ansible/latest/collections/community/docker/index.html)          | Existing automation capabilities    |
| [OWASP object permissions](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/) | Authorization test rationale        |

Links are carried from earlier documents; no network access or source revalidation occurs in this documentation task. Extended references/candidates are preserved in the roadmap. Exact tool/provider compatibility must be retested against release versions. Do not claim Docker lacks IPAM or Terraform/Ansible integrations.

## 3. Decisions before implementation

| Decision                                                    | Owners / evidence required                                             |
| ----------------------------------------------------------- | ---------------------------------------------------------------------- |
| Go module layout and signer move                            | A+B+C; shared client imports and CLI boundary                          |
| CLI/agent boundary                                          | A/E with B+C; dependency check and real auth                           |
| Credential-account/agent assignment and multi-account polls | Agreed one process/host, separate account credentials/clients; B+C+A finalize provisioning/custody/assignment; no new wire fields       |
| k3d+k3s tested pair and flags                               | A+B+C; one-day AlmaLinux host spike                                    |
| K8s port range/release and kubeconfig delivery              | B+A+C; reserve before create; cleanup; scoped export                   |
| Subnet pool and VPC/compute attachment                      | B+A; C allocator; route overlap and coordinated contract review        |
| MinIO scoped credential model/admin custody                 | B+C; private server path, rotation and negative tests                  |
| Backup storage format/retention                             | B+A; measured rows/checksum restore                                    |
| Internal/public/provider IDs                                | B+D; uuid/public_id/provider_id mappings                               |
| Audit/RLS/key custody and missing enums                     | B+C/D as applicable; nullable/security-events tradeoff; contract-first |
| One stack inputs/app contract                               | B+E+A/C; fixed sequence, JSON shapes and test 15                       |
| ITISS coding ownership                                      | Ask mentor in writing; if disallowed, named agent developer/handover   |

No versions, CIDRs, quota-error wire names or confirmations are chosen here. Extended proxy/health/cooldown/storage/identity/domain/tunnel decisions live only in the roadmap; old G0 schedule wording is superseded by the workflow.

### 3.1 HLD/LLD review decisions — target resolution

| Decision | Target selected | Remaining review |
|---|---|---|
| Runtime safety | OS host exclusion plus per-resource serialization; uncertain actions block takeover | A+C prototype; no Docker epoch-fencing claim |
| Lease semantics | Current owner/epoch/generation and unexpired DB-clock lease for renew/report/ack | B+D/A/C schema and fixtures; values TBD |
| Drift/health | next_attempt_at schedules retries and periodic inspections; new claim after convergence | B+A query/scheduling and C tests |
| Network transport | Separate signed network-control exchange, server-owned vpcs, default network prerequisite | Routes/fields/states TBD B+D/A/C |
| Shared MinIO | Operator-owned durable compute.instance runtime; protected host secret references; stack tenant bucket | B+C/A bootstrap/custody/private route; approved compute mapping |
| Required cuts | Preserve stack/test 15 and installer; optional provider and extended work cut first | Explicit SRS revision required for any other reduction |
| Capacity | Spike at first host preparation, before driver capacity commitment | Actual workload matrix; revalidate integration in November |

Prior contradictory cut wording is superseded, not a removed service requirement. Historical roadmap text and the gate spreadsheet preserve earlier targets/slogans; they do not establish runtime capability or override this review. No spreadsheet/code/contract edit occurs here.

## 4. Claim and evidence policy

Do not claim provider parity, VM isolation for Droplets, fake-IP routing, tamper-proof audit or an operational service from documentation. Unknown kinds/states are TBD with B; new messages require coordinated schema/fixtures/generated outputs. Reference DDL is not executable or migration-ready; no schema exists yet.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
