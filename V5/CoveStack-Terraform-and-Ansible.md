# CoveStack V5 — Terraform, OpenTofu, Ansible and Controlled Automation

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. API-only automation boundaries

| Tool | Target responsibility | Current evidence / limit |
| --- | --- | --- |
| Go CLI | Authenticated user request, JSON/wait/errors and idempotency | Planned; real resource APIs absent or incomplete |
| Terraform/OpenTofu provider | Thin adapter over approved resource APIs | Later optional work; no current provider implementation |
| Ansible installer | Repeatable operator-owned platform installation | Planned; actual entry points/artifacts/runtime profile must exist |
| Ansible dynamic inventory | Account-scoped inventory from API/CLI JSON | Future integration; no docker ps or direct DB |
| Floci | Optional AWS-shaped application testing inside/alongside an environment | External optional emulator, not the platform foundation or rewrite target |

Terraform files are not interpreted automatically by CoveStack. A provider implements the API mappings and supported lifecycle; Ansible targets reviewed APIs/guest access. No automatic compatibility with every external provider or arbitrary uploaded configuration is assumed.

## 2. Initial provider subset and state

E's later provider starts with real approved compute and Kubernetes resources, after CLI/client and VM compatibility are reviewable. That is an initial engineering slice, not a product service-count ceiling. Define plan/apply/read/refresh/update/delete/import behavior per supported API operation, immutable provider IDs, polling/error/timeout and unknown-outcome semantics. Do not claim refresh/import available until implemented and tested with real resources.

Provider state is a client cache; PostgreSQL remains authority. Use logical idempotency for retried create, detect drift through owned read, never directly call Incus/Docker/k3d or invent a second scheduler. Avoid secrets/kubeconfig in ordinary state; document sensitive/encrypted backend/storage and export policy for unavoidable values. D supplies API/fixture alignment, A Go review, C acceptance, B lifecycle semantics.

## 3. Module and authentication review

Shared pkg/kanclient, CLI/provider layout and HTTP/signing reuse are Planned, not code moves already made. User CLI/provider uses user auth; agent signing credentials are not a substitute. Agree module/toolchain/build boundary before adding dependencies or imports. API/key/profile configuration contains no local resource DB. Source provider/framework/license versions are pinned after actual tests.

## 4. Ansible operator installer

Target reviewed roles for approved backend/runtime, dependencies, Java server, agent and secrets; no playbook is created here. Use protected Vault/no_log or approved secret backend, verified checksummed/signed artifacts, narrow privileges and scoped configuration. Proposed systemd hardening (NoNewPrivileges/ProtectSystem) must be tested with required privileged socket/storage access; it is not a guarantee of harmless agent authority.

Acceptance: another operator installs exact approved profile cleanly, validates actual API/auth/runtime/isolation/restore, and second run has no unexpected changes. Require lint/idempotence/upgrade/rollback evidence. A hosts installer/CI/ITISS; B server/SQL, C security/tests. Do not install into the restricted current sandbox or imply Compose starts the full product.

## 5. Inventory and guest configuration

Future inventory reads authenticated owned resources via `covestack ... list -o json` once the CLI/API exists; command shape is target, not runnable current code. Script outputs Ansible inventory JSON and never uses docker ps/Incus socket/database as tenant authority. Bob's inventory contains none of Alice's instances (E8). Guest configuration requires approved SSH key injection, secure connectivity and an appropriate guest image/tooling (e.g. Python where the playbook requires it). These do not permit arbitrary host commands.

## 6. Arbitrary execution and future scope

No platform-run arbitrary Terraform/Ansible job engine, embedded CI/CD service or tenant host-shell upload is included. If later proposed, it needs its own threat model, restricted identity/artifact/network/secret boundary, quotas, cancellation/audit and immutable inputs. Integrate existing tools before inventing an executor. Rich templates/registries/deployment automation remain classified by the catalogue.

## 7. Acceptance and handover

Prove real create/read/idempotent retry/import/refresh/delete for the exact advertised provider subset; foreign account denied, secret-safe state/errors, no duplicate uncertain outcome and no backend library import. Installer proves clean/repeat install/restore; inventory proves E8 and stale/error behavior. [Go handover](CoveStack-Go-Handover.md), [existing-code plan](../EXISTING-CODE-CHANGE-PLAN.md), [API](CoveStack-API-and-Compatibility.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-terraform-and-ansible) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
