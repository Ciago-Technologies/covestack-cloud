# CoveStack V5 — Dependency-Driven Team Workflow

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Planning authority

This is a working dependency plan, not a dated academic delivery schedule or release-by-release roadmap. Product core/future classification follows the catalogue. Historical dates/counts remain in the linked snapshot, not current requirements. Existing team roles remain; no service or runtime evidence is declared finished by this workflow.

## 2. Current starting checkpoint

The existing Go signed client/main/poll loop runs no backend work. Immediate goal: real authenticated Java no-work polling with protected assigned credentials, negative auth and cancellation/error observations. Server foundation, concrete credential lookup and trusted assignment are missing or unverified. Do not start by recreating loop.go or wiring Docker/Incus commands into a handler.

| Step | Lead / dependencies | Reviewable output | Exit evidence |
| --- | --- | --- | --- |
| Handshake agreement | B+C+A; D compatibility | Real endpoint, no-work shape, key custody/provisioning and trusted host/account assignment | Written nonsecret agreement and missing decisions |
| Server foundation | B+C; D API | Runnable minimal server and reviewed prerequisite migration/auth path | Actual startup and clean/upgrade schema evidence on approved environment |
| Signed integration | A+C; B real server | Existing client/loop against assigned real identity | Valid [], negative auth/replay, cancellation and observed retry/failure |
| Protocol/resource review | B+D+A+C | VM/App Container types, claims/renew/report, capability/network/capacity and migration plan | Reviewed authoritative schemas/fixtures/version strategy before execution |
| Backend spike | A+C; B reservations | Authorized Incus/OS/image/network/storage profile | Real guest kernel/readiness, task recovery, resource measurement and teardown |

## 3. Foundation-to-resource dependency order

| Slice | Must exist first | Work and required exit |
| --- | --- | --- |
| Identity / account/project ownership | Reviewed custody and assignment | Real credentials/membership, fail-closed auth, audit and owner-scoped API |
| Durable operations / admission | Authoritative schema and real DB | Idempotent intent+operation+reservation, claims/fence policy and no parallel overspend |
| Images / disk / networking policy | Trusted host/profile and ownership schema | Approved image provenance, volume/NIC/IP/port allocation, actual default-deny traffic |
| VM lifecycle / App Container | Above plus approved type/driver interface | Inspect/create/readiness/start-stop/cleanup; uncertain outcome and crash no duplicates |
| Kubernetes | Safe container runtime, ports/capacity/secrets/network | Tested k3d/k3s pair, real workload, protected kubeconfig and clean recovery/deletion |
| PostgreSQL / object storage | Persistent network/storage/secrets and trusted executors | Actual SQL/object bytes, scoped credentials, restart and verified restore; Java RustFS reconcile |
| Bounded environment | Owned ready child services and reservation policy | Typed graph/retry/retention; usable app-data workflow and shared-resource preservation |
| Interfaces / operations | Real APIs and evidence | CLI JSON/wait/errors, installer, monitoring/freshness, audit, operator runbooks and profile qualification |

Security, audit, failure tests, quotas and recovery are integrated into each slice rather than deferred to the final demonstration. E may learn/build mocks in parallel, but a mock or planned path is never labelled an available runtime.

## 4. Capacity and runtime spikes

VM host evaluation checks hardware/nested virtualization, supported Linux/kernel/Incus/QEMU/KVM, guest image/firmware provenance, management socket permissions, disk/NIC/firewall, unknown task outcomes and measured overhead. Kubernetes spike checks reviewed OS/cgroup/SELinux/firewall behavior, Docker+k3d+k3s compatibility and actual load. No fixed 8 GB production recommendation, two-cluster product cap or idle-memory estimate qualifies a host.

## 5. Daily engineering workflow

1. Confirm baseline, issue owner, acceptance, affected files and dependency owner; use [Starter.md](../Starter.md) Git/PR steps.
2. Read current source/contract, permanent architecture and the existing-code change plan; do not patch generated files alone.
3. Implement the smallest approved slice; do not add dependencies/runtimes without a consumer and review.
4. Verify unit/contract checks, then real environment/failure evidence on an authorized host; preserve actual negative results.
5. Open a concrete PR with owner/reviewers, migration/API/security impact and limitations; no direct main push.
6. Update the same issue/card, affected docs and canonical AI checkpoint. Board Done is not service Supported.

Weekly learning review for E is an explainable code/test/package outcome, not a fixed calendar promise. A pairs on Go; B/C/D provide mocks/fixtures without secretly changing interfaces.

## 6. Scope pressure and stop rules

Stop speculative future work when it blocks required reviews, security, ownership or recovery. Do not cut failure tests to meet a demonstration date. If a slice lacks data/auth/cleanup evidence, keep it Planned/In Progress/Blocked with an explicit issue rather than marking it Supported. Replan capacity or workload subset, not safety controls. No academic cut-order or automatic service deadline applies to the product.

## 7. Demonstration and evidence

Live highlights can show reviewed real no-work/auth, idempotent requests, owner denial and truthful desired/observed views once APIs exist. VM/RustFS/SQL/cluster/traffic/restore proof may be real CI output or recorded demonstrations from the exact revision. Record backup video and sanitized setup; do not depend on hotspot/router rights. Never use a fictional CLI/server command to conceal absent code. [Acceptance](CoveStack-AntiWrapper-Acceptance.md).

## 8. Operations handover

Before a supported profile: another operator clean-installs exact artifacts, credentials/keys recover, known data restores, stale agent health is truthful, quotas/admission reject safely, upgrades/rollback are rehearsed and license/notices/disclosure responsibilities are known. Single-host hardening is not HA.

## 9. Future development dependencies

Future candidates are listed in [extended dependencies](CoveStack-Extended-Roadmap.md), not promoted by schedule. Multi-host needs partition-safe ownership/reservations plus locality; publication needs explicit route/domain/TLS controls; richer databases/registries/migration/GPU work needs demand, rights and backend conformance. No new release plan is generated.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-project-workflow) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
