# CoveStack V5 — Detailed Product and Team Documentation

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Start here

**V5 is retained and its bodies are updated for the current product.** It is not a folder to delete or a wrapper notice over old requirements. Product is local-first account-owned infrastructure: real VM Instance primary, App Container separate, Java/Spring/JDBC authority, Go trusted agent and thin CLI, RustFS preferred Planned object provider, capacity/safe reconciliation/security/verified recovery.

[Root Starter](../Starter.md) explains clone/branch/first work; [team checkpoint](../TEAM-FIRST-CHECKPOINT.md) explains shared acceptance and board; [canonical AI checkpoint](../checkpoint.md) preserves task continuity. The [product catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) remain shared scope authorities. V5 supplies detailed aligned working guidance.

## 2. Product status and boundaries

Current Go config/signer/client/no-execution loop/main exist; Java auth/replay/helpers and partial extended SQL exist. Full server/credential/host assignment, claims, real VM/container/Kubernetes/data services and clients still require implementation/evidence. No service becomes Supported here; no qualified production profile. Current actual wire is 1.1.0, not V5 or an invented future version.

Core/future scope follows the catalogue, independent of old academic limits/dates/template/node/network caps. No new Labs product or release-by-release plan is introduced. RustFS Java bucket/policy/key reconcile is separate from operator service/volumes and actual byte restore. Apache-2.0 is adopted for CoveStack-owned material; third-party obligations remain.

## 3. Full current document index

| File | Current responsibility |
| --- | --- |
| [DocsV5-DO_Team.md](DocsV5-DO_Team.md) | HLD/LLD, authority, provider planes, safe lifecycle and decisions |
| [CoveStack-Service-Catalog.md](CoveStack-Service-Catalog.md) | 22 core / 31 future catalogue and ownership boundaries |
| [CoveStack-SRS-Scope.md](CoveStack-SRS-Scope.md) | Independent product requirements, domain deliverables and acceptance |
| [Team_Roles.md](Team_Roles.md) | A–E ownership, pair reviews, first tasks and ITISS |
| [CoveStack-Project-Workflow.md](CoveStack-Project-Workflow.md) | Dependency-driven integration/implementation/review workflow |
| [CoveStack-Complete-Working-Setup.md](CoveStack-Complete-Working-Setup.md) | Actual checks, missing setup, future profile/install and constraints |
| [CoveStack-Current-Setup-Fixes.md](CoveStack-Current-Setup-Fixes.md) | Concrete current blockers and owners before coding |
| [CoveStack-Database-Schema.md](CoveStack-Database-Schema.md) | VM-first proposed relational model, exact current states, concurrency and migration boundaries |
| [CoveStack-Service-Gates.md](CoveStack-Service-Gates.md) | Core/future promotion, domain failure evidence and profile qualification |
| [CoveStack-AntiWrapper-Acceptance.md](CoveStack-AntiWrapper-Acceptance.md) | Exact current 1–15/E1–E8 assertions and truthful viva answers |
| [CoveStack-Testing-and-Acceptance.md](CoveStack-Testing-and-Acceptance.md) | Evidence levels, current checks, real acceptance and CI boundaries |
| [CoveStack-Go-Handover.md](CoveStack-Go-Handover.md) | A/E boundaries, staged CLI learning and review |
| [CoveStack-Anuj-Next-Steps.md](CoveStack-Anuj-Next-Steps.md) | Existing loop integration first, later driver/infra/ITISS dependencies |
| [CoveStack-Team-START-HERE.md](CoveStack-Team-START-HERE.md) | Team entry point, clone/issue/read/review rules |
| [Folder_Structure_Explanation.md](Folder_Structure_Explanation.md) | Actual source tree versus planned modules and generated/secrets policy |
| [CoveStack-API-and-Compatibility.md](CoveStack-API-and-Compatibility.md) | Actual 1.1.0 versus future VM/type/claim/network/capacity protocols |
| [CoveStack-Resource-Services.md](CoveStack-Resource-Services.md) | Owned VM/container/Kubernetes/SQL/RustFS/network/storage lifecycle |
| [CoveStack-VPC-IPAM-and-Networking.md](CoveStack-VPC-IPAM-and-Networking.md) | IPAM/NIC/firewall/traffic/route ownership, actual denial and recovery |
| [CoveStack-Stack-Templates.md](CoveStack-Stack-Templates.md) | Bounded environment dependencies/readiness/retry/retention |
| [CoveStack-Terraform-and-Ansible.md](CoveStack-Terraform-and-Ansible.md) | Thin provider, operator installer, future inventory and automation safety |
| [CoveStack-UI-CLI-and-Monitoring.md](CoveStack-UI-CLI-and-Monitoring.md) | API-only client/UI, core fresh health/logs and advanced observability |
| [CoveStack-Security-and-ITISS.md](CoveStack-Security-and-ITISS.md) | Threat model, actual isolation/custody, host controls and evidence |
| [CoveStack-Instance-Metadata.md](CoveStack-Instance-Metadata.md) | Management inspection versus future guest identity/delivery |
| [CoveStack-Publication-Domains-and-Tunneling.md](CoveStack-Publication-Domains-and-Tunneling.md) | Future opt-in gateway/routes/domain/TLS/actual reachability |
| [CoveStack-Operations-and-Recovery.md](CoveStack-Operations-and-Recovery.md) | Incident, retention, independent backup/verified restore and upgrade |
| [CoveStack-Implementation-Status.md](CoveStack-Implementation-Status.md) | Source evidence, missing code, environment restrictions and next integration |
| [CoveStack-References-and-Decision-Log.md](CoveStack-References-and-Decision-Log.md) | Authority, existing research, open decisions and relocation record |
| [CoveStack-Extended-Roadmap.md](CoveStack-Extended-Roadmap.md) | Future dependency catalogue, not release dates or academic ceiling |
| [CoveStack-PDF-and-Formatting.md](CoveStack-PDF-and-Formatting.md) | Markdown formatting, stale export/tooling limits and sharing |
| [README.md](README.md) | This current V5 reading and authority map |
| [tools.md](tools.md) | Existing pins versus evaluated providers/tooling and rights/profile checks |


## 4. Implementation reading order

1. Product catalogue/SRS → main design → roles/workflow → source status and current fixes.
2. Team onboarding → actual setup/known constraints → first real signed no-execution server integration.
3. API/compatibility → VM-first schema/concurrency → resources/network/environment design.
4. Gates/acceptance/tests → security/ITISS/recovery → provider/interface/capacity decisions.
5. Future dependencies and advanced topics only after their prerequisites are reviewable; no plan is implementation evidence.

A–E first responsibilities stay in roles/root team guide. A integrates the existing loop with B+C actual credentials/server before runtime work; E uses labelled mock learning and an agreed API/module boundary. Do not recreate source that already exists or use historical fake setup commands.

## 5. History and preservation

One [historical V5 snapshot](../docs/archive/V5-before-full-product-alignment-2026-11-04.md) preserves all 31 previous bodies unchanged, including full old DDL, old audit, timelines and acceptance variants. Every active V5 file links its exact snapshot section. Use those only for traceability; do not treat old container/MinIO/cap/count/version proposals as current implementation. Existing source/contracts/generated code/config are unchanged by this doc alignment.

## 6. Sharing, export and review

Share this index, Starter, team roles/workflow, design/SRS/catalogue, setup/status, API/schema and acceptance/security links from the agreed branch/revision. Ordinary GitHub issues hold live task status; checkpoint.md is AI continuity. No secret config, Terraform state, signing/admin keys or credential kubeconfig in shared material.

Old HTML/PDF/print validator has obsolete calendar/audit assumptions and remains unchanged; no fresh bundle is claimed. Read [formatting/export limits](CoveStack-PDF-and-Formatting.md) before sharing exports. Runtime/profile promotion requires actual gates and operational evidence, not correct Markdown.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-readme) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
