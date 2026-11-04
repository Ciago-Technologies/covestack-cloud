# CoveStack V5 — Product Requirements and Scope

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Status, authority and terminology

This active V5 requirements view follows the [product SRS](../docs/product/CoveStack-Product-SRS.md) and canonical catalogue. Product scope is independent of historical academic obligations/dates. All capabilities below are target requirements; no service is made available by this document. V5 is not a wire version.

## 2. Core requirements

| ID | Product capability | Required outcome |
| --- | --- | --- |
| 1 | Compute Instances (VMs) | Image, vCPU/RAM/disk/NIC, secure initialization; create/inspect/start/stop/restart/resize/snapshot/backup/restore/delete |
| 2 | App Containers | OCI image, ports, env/secret refs, CPU/RAM, volumes/network, health/logs, start/stop/restart/update/delete |
| 3 | Kubernetes Clusters | Create/delete/inspect/lifecycle, nodes/limits, reserved API endpoint, protected kubeconfig, health/network and standard-tool workload deployment |
| 4 | VM Images and Templates | Approved OS images/custom-image intake, versioned VM templates, validated cloud-init or equivalent, provenance |
| 5 | VPC / Private Networks | Owned network create/inspect/membership/attachment/dependency-safe delete |
| 6 | Subnets | Operator-reviewed ranges, allocate subnet from permitted pool, reservations and no overlap |
| 7 | IPAM | Allocate/reserve/inspect/release private addresses; interrupted-allocation recovery |
| 8 | Network Interfaces | Owned interface create/inspect and validated attach/detach; address tracking |
| 9 | Firewalls / Security Groups | Default deny, ingress/egress protocol/port/source/destination rules, associations and audit |
| 10 | Block Storage / Volumes | Create/inspect/attach/detach, supported resize, ownership/usage and safe delete |
| 11 | Object Storage | S3 bytes/buckets, scoped keys/policies, quota, persistent objects, backup/restore; capability-declared versioning/lifecycle |
| 12 | Managed PostgreSQL | Private database/user/credential lifecycle, real readiness, persistent storage, start/stop/health/backup/restore/delete/rotate |
| 13 | Application Environments | Owned resources and bounded dependencies, provision order/readiness/status/progress, safe retry/recovery/cleanup |
| 14 | IAM and RBAC | Users, organization/account membership, roles/permissions, default-deny resource/project authorization |
| 15 | API Keys and Service Accounts | Scoped programmatic identity, issuance/revocation/rotation/expiration policy and trusted agent assignment |
| 16 | Secrets Management | Sensitive credential references, scoped retrieval/redaction/audit/rotation through evaluated mature backend |
| 17 | Audit Logs and Events | API/admin/agent/resource/security/credential events, correlation, append-only app access |
| 18 | Resource Quotas | CPU/RAM/disk/VM/container/cluster/storage/network per-account limits |
| 19 | Monitoring and Resource Health | Host/VM/container/Kubernetes/agent/scheduler CPU/RAM/disk/health and observation freshness |
| 20 | Basic Logging | Platform/agent/provisioning and authorized workload logs, sanitized errors and retention |
| 21 | Snapshots and Backup/Restore Foundations | Owned VM/volume recovery points and metadata, bounded schedules/retention, verify/restore/history |
| 22 | CLI and API | Unified scoped lifecycle API and thin CLI, JSON, idempotency, wait/cancel/progress and explicit errors |


## 3. Cross-cutting requirements

| Concern | Required behavior |
| --- | --- |
| Ownership | Account/project/member/action authorization before lookup/mutation; foreign resource 404 without existence leak |
| Desired state | PostgreSQL authority, stable identities/generations/operations, transactional idempotency/reservations and safe reconciliation |
| Runtime trust | True VM guest kernel; explicit shared-kernel container class; trusted agent only at privileged sockets |
| Capacity | Fresh host capability/health, physical reserves plus account quotas, safe refusal, no assumed density/HA |
| Secrets | Hash user tokens, protect signing/DB/S3 material and key recovery, narrow delivery and rotation/revocation; no logs/spec leaks |
| Recovery | Known actual guest/data/SQL/object restoration, protected independent copies and retained-data policy |
| Automation | API-only CLI/provider/inventory, idempotent mutations, JSON inspection, truthful finite waits |
| Audit/telemetry | Sanitized scoped immutable events, actual health/freshness and meaningful incident diagnostics |

## 4. Deliverable evidence by domain

| Domain | Acceptance reference | Actual evidence required |
| --- | --- | --- |
| VM / App Container | Tests 1–6 plus named VM/container readiness and privilege cases | Guest-kernel VM distinct from container; owned inspect/start/stop/recovery and no cross-account access |
| Kubernetes | Tests 7–10 | Actual cluster/workload, scoped kubeconfig, measured quota/admission and owned cleanup |
| PostgreSQL | Test 11 | Real SQL, private/scoped credentials, durable rows and verified restore |
| RustFS | Tests 12–13 | Known object bytes, denied foreign key, restart persistence and Java bucket/policy recovery; independent verified restore |
| Networks/volumes/security/capacity | Test 14 plus named traffic/disk/admission cases | No allocation races, real forbidden packets/attachments denied and no overspend |
| Environments | Test 15 | Usable dependencies/app, partial retry, no duplicate child and shared-retained data preserved |
| Platform interfaces/operations | Service-gate and named CLI/audit/custody/install checks | Actual authenticated API/JSON, freshness, clean install, upgrade/rollback and data/key recovery |

IDs 1–15/E1–E8 are retained traceability labels, not an old service-count ceiling or proof previous container tests validate a VM. Product-domain additions remain necessary. [Acceptance definitions](CoveStack-AntiWrapper-Acceptance.md).

## 5. Future scope and constraints

Future candidates/dependencies appear in [catalogue](CoveStack-Service-Catalog.md#3-future-candidates) and [extended dependencies](CoveStack-Extended-Roadmap.md). No arbitrary workflow runner, new Labs product, full cloud parity, unsupported node density or automatic public exposure is implied. Profiles/capability differences are documented and rejected explicitly when unsupported. [Known constraints](CoveStack-Complete-Working-Setup.md#10-known-constraints).

## 6. Compatibility and implementation boundaries

The actual contract is 1.1.0. VM/App Container types, capability/host assignment, full claims/renew/reports/network/capacity messages require coordinated B+D+A/C design, fixtures and generated outputs. API state strings are taken from the current authoritative catalogue; missing values are TBD with B/D, not invented here. Java owns DB transactions; Go agents and clients do not access the platform DB directly.

## 7. Promotion and security

G1: PostgreSQL desired state and generation are authoritative. G2: crash-safe inspection/reconciliation rebuilds state/config without duplicates. G3: account authorization denies foreign API resources (404 without existence leak) and actual traffic/storage boundaries are tested. G4: meaningful failure acceptance runs in CI on an authorized capable runner. **Acceptance Gates Passed** requires this evidence; **Supported** additionally needs a declared subset, qualified profile and maintenance policy.

The product remains Experimental; no qualified production profile. Apache-2.0 is adopted for CoveStack-owned material; dependency/image notices and copied-reference rights remain separate. [Threat/security](CoveStack-Security-and-ITISS.md), [gates](CoveStack-Service-Gates.md), [status](CoveStack-Implementation-Status.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-srs-scope) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
