# Apache-2.0 adoption and RustFS confirmation — current follow-up

Updated: 4 November 2026. Targeted documentation/licensing update; reused the existing audit and provider analysis. Older reports below preserve the accurate earlier approval-pending history but no longer describe the current license.

Maintainer explicitly approved Apache License 2.0. Root LICENSE is the unmodified official apache.org text (11,358 bytes, 202 lines; SHA256 cfc7749b96f63bd31c3c42b5c471bf756814053e847c10f3eb003417bc523d30). SPDX identifier Apache-2.0 now appears consistently in active project/contribution/SRS documentation. No fabricated holder/author or source header added. No NOTICE invented: preserve any applicable upstream NOTICE attribution in actual distributions, after inspecting the selected artifacts. Existing dependency/external-project licenses remain unchanged.

RustFS remains preferred and Planned. Official S3 matrix and IAM docs rechecked; explicit bucket/object/multipart subset documented, ACL authorization excluded and bucket logging/ownership controls planned. Operator storage, account buckets, Java policy/credential reconcile and replaceable adapters retained. Existing real-byte, tenant/foreign denial, restart and verified restore acceptance requirements remain pending; no runtime tests run.

## Files changed in this follow-up

| File | Disposition |
| --- | --- |
| `AGENTS.md` | Modified |
| `COVESTACK-PROJECT-OVERVIEW.md` | Modified |
| `LICENSE` | Created |
| `README.md` | Modified |
| `checkpoint.md` | Modified |
| `docs/README.md` | Modified |
| `docs/documentation-audit.md` | Modified |
| `docs/platform/README.md` | Modified |
| `docs/platform/contributing.md` | Modified |
| `docs/platform/decisions.md` | Modified |
| `docs/platform/implementation-status.md` | Modified |
| `docs/platform/license-and-attribution.md` | Modified |
| `docs/platform/licensing-decision.md` | Modified |
| `docs/platform/provider-evaluation.md` | Modified |
| `docs/platform/references.md` | Modified |
| `docs/platform/strategy.md` | Modified |
| `docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md` | Modified |
| `docs/product/CoveStack-Product-SRS.md` | Modified |

Scope: 17 existing Markdown files updated; root LICENSE created. Pre-existing changes remain untouched. No source/config/test/schema/migration/contract/dependency/Compose/infrastructure modifications; .gitignore fingerprint unchanged. No installs, commit, push or application execution.

Verification: LICENSE byte equality with official download; active license wording and internal links; Markdown tables/fences/README rendering; historical V5/archive/checkpoint preservation; protected-file fingerprints and git diff whitespace. Existing 48 historical missing local references remain documented. Prior code/runtime reports were not rerun. Remaining concerns: resolved transitive licenses/SBOM, bundled notices/source obligations, image/runtime/firmware rights and copied handbook/reference provenance; contributor governance beyond the license is not invented. RustFS artifact/admin/custody/security/restart/service-wide backup profile still requires implementation and real evidence.

<!-- Earlier audit and correction reports preserved below; historical results. -->

# Targeted corrections — RustFS, independent product scope and licensing

Updated: 4 November 2026. This is a focused follow-up using the existing audit, not a new repository-wide audit. Latest user decisions supersede the earlier approval/scope conditions below; historical results remain preserved. Product core/future classifications are unchanged but independent of academic CORE-5.

## Decisions and evidence

- RustFS preferred candidate, integration Planned. Official S3/IAM/persistence/license material reviewed; compatibility limits and unverified Java admin, scoped credentials, restart and actual service backup/restore are documented in provider-evaluation.md. Java owns bucket/policy/credential reconciliation; operator owns process/volumes; alternatives remain possible.
- New active product SRS adopts the catalogue independently. Historical V5 bodies and previous academic transition text remain intact under current notices. Old academic dates, counts and caps are not product requirements.
- Apache-2.0 recommended over MIT for express patent terms, pending rights-holder approval and root publication. No root LICENSE added, no upstream license altered, no final transitive redistribution clearance claimed.
- README rendered using the already-installed markdown-it parser: one H1, one table/five rows, 22 internal links. No escaped markup/entities. Agent source was read: signed single-account polling/no-execution exists; full provisioning/assignment/claims/runtime execution pending.
- .gitignore had no staged/unstaged diff at task start. Its SHA256 matches both prior task-start snapshots. Commit 51014ba contains the overview ignore rule; that content predates the previous Codex documentation work. No accidental Codex hunk is established, so no ignore rule was reverted. Earlier author intent is not independently known; preserve existing user work rather than assume it was accidental.

## Every file changed in this correction

| File | Disposition | Change |
| --- | --- | --- |
| `AGENTS.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `COVESTACK-PROJECT-OVERVIEW.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `README.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `V5/CoveStack-API-and-Compatibility.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-AntiWrapper-Acceptance.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Anuj-Next-Steps.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Complete-Working-Setup.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Current-Setup-Fixes.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Database-Schema.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Extended-Roadmap.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Go-Handover.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Implementation-Status.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Instance-Metadata.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Operations-and-Recovery.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-PDF-and-Formatting.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Project-Workflow.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Publication-Domains-and-Tunneling.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-References-and-Decision-Log.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Resource-Services.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-SRS-Scope.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Security-and-ITISS.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Service-Catalog.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Service-Gates.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Stack-Templates.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Team-START-HERE.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Terraform-and-Ansible.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-Testing-and-Acceptance.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-UI-CLI-and-Monitoring.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/CoveStack-VPC-IPAM-and-Networking.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/DocsV5-DO_Team.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/Folder_Structure_Explanation.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/README.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/Team_Roles.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `V5/tools.md` | Modified | Historical body preserved; active relocation scope/provider/SRS pointer corrected |
| `agent/README.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `checkpoint.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `deploy/README.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/README.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/documentation-audit.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/README.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/academic-transition.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/architecture.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/capacity.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/decisions.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/environments.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/feature-specifications.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/future-services.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/implementation-status.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/installation-and-operations.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/license-and-attribution.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/licensing-decision.md` | Created | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/network-and-storage.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/networking-and-policy.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/provider-evaluation.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/readiness.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/references.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/platform/strategy.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md` | Modified | Current product scope/provider/licensing/status/navigation/handoff correction |
| `docs/product/CoveStack-Product-SRS.md` | Created | Current product scope/provider/licensing/status/navigation/handoff correction |

Result: **57 existing Markdown files modified, 2 new Markdown files created**. The ignored root overview remains ignored by the unchanged existing rule; it will need explicit maintainer publication handling later if desired. No Git mutation is authorized here.

Static verification: Markdown tables/fences and active local links/anchors; unchanged catalogue IDs/classifications; historical-body and protected-file fingerprints; README parse/render; diff whitespace. Existing 48 missing historical references are preserved, not current product link failures. No runtime acceptance or legal clearance is implied. No implementation/config/schema/migration/contract/dependency/test/infrastructure files changed; .gitignore unchanged. No app/infrastructure tests, installs, commits or pushes.

Remaining decisions: exact RustFS artifact and support profile, Java admin adapter/permissions, bootstrap and tenant key custody/rotation, consistent protected backup/restore and independent target, upgrade/security/capacity evidence; rights-holder approval, dependency inventory/notices and published root license. Engineering protocol/schema review remains necessary independently of academic history.

<!-- Earlier documentation audit preserved below; historical task results. -->

# Current documentation catalogue refinement — 4 November 2026

This section records the latest master task. The earlier redesign audit below is retained as historical evidence; its counts are not the totals for this refinement. Scope is Markdown only, with read-only source/config/contract inspection and official research. No application or infrastructure tests, installs, Git mutations or environment provisioning.

## Current inventory, disposition and evidence

The starting checkout contained **227 inventoried files and 97 Markdown documents**, including historical source documents and the previous redesign. All active documentation and relevant source/contract/config surfaces were reviewed; historical records were inspected for authority/disposition rather than treated as current implementation. The previous archive inventory is retained below. No MDX/reStructuredText/AsciiDoc source was found. Fingerprints before this refinement are held in /tmp/covestack-catalogue-before.json for this session only.

| Document / group | Current purpose | Problems | Required changes / disposition |
| --- | --- | --- | --- |
| Root README and overview | Discovery and project explanation | Old positioning, MinIO default, active release-plan link | Point to canonical product catalogue; factual Experimental status and no default provider |
| docs/platform active specifications | Architecture and domain evidence | Partial catalogue, missing policy/environment detail and comparison dimensions | Align target core/future scope, capacity, providers, security, domain acceptance and evidence |
| docs/product canonical catalogue | Product authority | Missing before refinement | Created: 22 proposed core capabilities, 31 deferred candidates, shared lifecycle/security; no new microservice count commitment |
| V5 historical set | Academic strategy/requirements/evidence | Conflicts with VM-first product | Preserve entire historical body; update relocation pointer only; academic changes need approval |
| Old release roadmap | Previous strategy | User now excludes release planning | Preserved verbatim beneath explicit supersession notice; no active release assignments |
| Previous checkpoints, archived guides, generated exports | Historical evidence | Stale architecture and broken historical local links | Preserve; active links fixed, 48 existing historical missing targets explicitly retained |
| AGENTS.md and checkpoint.md | Task boundaries / canonical continuity | Stale active scope/status direction | Update current authority and exact next action; preserve earlier checkpoint history |

## Current modified / created Markdown files

| File | Disposition | Change |
| --- | --- | --- |
| `AGENTS.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `COVESTACK-PROJECT-OVERVIEW.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `README.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `V5/CoveStack-API-and-Compatibility.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-AntiWrapper-Acceptance.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Anuj-Next-Steps.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Complete-Working-Setup.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Current-Setup-Fixes.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Database-Schema.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Extended-Roadmap.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Go-Handover.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Implementation-Status.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Instance-Metadata.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Operations-and-Recovery.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-PDF-and-Formatting.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Project-Workflow.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Publication-Domains-and-Tunneling.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-References-and-Decision-Log.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Resource-Services.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-SRS-Scope.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Security-and-ITISS.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Service-Catalog.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Service-Gates.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Stack-Templates.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Team-START-HERE.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Terraform-and-Ansible.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-Testing-and-Acceptance.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-UI-CLI-and-Monitoring.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/CoveStack-VPC-IPAM-and-Networking.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/DocsV5-DO_Team.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/Folder_Structure_Explanation.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/README.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/Team_Roles.md` | Modified | Historical body unchanged; current catalogue pointer |
| `V5/tools.md` | Modified | Historical body unchanged; current catalogue pointer |
| `checkpoint.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/README.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/documentation-audit.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/README.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/academic-transition.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/agent-and-reconciliation.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/api-and-cli.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/architecture.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/capacity.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/competitive-analysis.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/contributing.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/decisions.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/environments.md` | Created | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/feature-specifications.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/future-services.md` | Created | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/implementation-status.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/installation-and-operations.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/license-and-attribution.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/network-and-storage.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/networking-and-policy.md` | Created | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/observability.md` | Created | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/provider-evaluation.md` | Created | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/readiness.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/references.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/resource-model.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/roadmap.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/security.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/platform/strategy.md` | Modified | Current domain, evidence, navigation or handoff alignment |
| `docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md` | Created | Canonical product identity and complete proposed catalogue |

**Current refinement: 57 existing Markdown files modified; 6 new Markdown files created.** Repository Git diff additionally includes changes already present before this refinement; do not attribute those to this task. Some legacy Markdown is ignored locally, so it will not appear in ordinary Git status; ignore rules were not edited.

## Verification and unresolved approvals

Static validation: 602 Markdown tables checked, no uneven columns/unclosed fences, no broken active local links/heading anchors. The 48 missing targets are existing historical references. Catalogue IDs 1–22 and 23–53 are checked against the requested core/future lists. Mermaid source is structurally reviewed; no diagram renderer or visual export was run. V5 historical bodies, previous archived originals, contract/code/config/tests/SQL and pre-existing .gitignore fingerprint are preserved. No service gate was promoted.

Open decisions: B+C+A/D approve versioned VM/container/assignment/network/capacity protocols; mentor/team reconcile VM compute and broader product core with academic CORE-5 and deadline; B+C/A select object-storage/secret custody and tested artifacts; maintainers decide license grant and supported profile. Scheduler ownership is not silently assigned. Incus is the preferred evaluation candidate, not an implemented/qualified backend. Multi-host, advanced services and optional Floci remain future/integration scope.

No requirements were deleted: prior roadmap is now historical; legacy V5 bodies and old audit are retained. No Labs/learning-management feature or new release plan was introduced. Maturity gates describe evidence, not a dated release roadmap. No commit/push.

---

<!-- Previous redesign audit retained below; historical results. -->

# Repository documentation audit and redesign report

Audit date: 4 November 2026. Scope: documentation only. No application/infrastructure tests or package installation. This inventory was established before rewriting the overview; every text artifact was read for structure/content and relevant implementation/architecture was inspected in depth. PDF text was extracted to temporary files; archived Markdown members were read without changing the ZIPs. Binary renders were inventoried/inspected as artifacts, not certified for visual rendering.

## 1. Baseline inventory and method

- 80 loose documentation artifacts: 69 Markdown, 3 PDF, 1 HTML, 2 SVG and 5 generated test-report TXT files.
- 34 archived text documents inside 3 existing ZIP containers, read-only; 114 document entries inspected counting archive members separately, plus ZIP container inventory.
- No MDX, reStructuredText or AsciiDoc files were found. Historical test-report output was not rerun.
- Read source: Go config/signer/client/poll loop and test surfaces; Java auth/replay/contracts, pom/go.mod, SQL file inventory, Compose and CI definitions. Source/build/config were read only.
- Planning policy: deep repository/compatibility/security analysis, no delegation; verification by static document/link/diff/fingerprint invariants, not runtime claims.

## 2. Per-document baseline inventory

Problems identify this snapshot, not a claim that historical authors were describing current code. Disposition distinguishes active replacement from preservation. No original requirement or historical test outcome was deleted.

| Document | Current purpose | Problems | Required changes | Disposition |
| --- | --- | --- | --- | --- |
| `Starter.md` | CoveStack — Teammate Start Guide | Old strategy, container model, dates or moved links | Supersession notice and current references; repair relocation links where possible | Historical body retained; link-only repair in two files |
| `COVESTACK-PROJECT-OVERVIEW.md` | CoveStack — Project, Use Cases, Workflow and Infrastructure | Stale positioning/current-state or dependency-only setup unclear | Current VM-first/evidence-accurate entry point; original body copied to explicit archive | Updated active; old body preserved |
| `checkpoint.md` | CoveStack — AI agent handoff checkpoint | Competing historical strategy and changing next-task snapshots | Current authoritative summary, retain all prior entries verbatim | Updated active prefix, historical body preserved |
| `HANDBOOK.md` | LocalCloud handbook | Foreign LocalCloud source, not CoveStack; 27 missing sibling links | Reference/attribution notice; retain unresolved external historical names | Prefix only; external body preserved |
| `Kan_Cloud_Database_Schema_Simple_Guide.pdf` | Kan_Cloud_Database_Schema_Simple_Guide.pdf | Historical rendered model; may lag source and new VM target | Label through indexes/audit, no regeneration or binary rewrite | Unchanged historical export |
| `AGENTS.md` | CoveStack task rules | Old container/two-driver/status/deadline rules conflict with latest task | Reconcile latest target and evidence/maturity authority | Updated active rules |
| `TEAM-FIRST-CHECKPOINT.md` | CoveStack — First Team Checkpoint and GitHub Board Guide | Old strategy, container model, dates or moved links | Supersession notice and current references; repair relocation links where possible | Historical body retained; link-only repair in two files |
| `README.md` | CoveStack Cloud | Stale positioning/current-state or dependency-only setup unclear | Current VM-first/evidence-accurate entry point; original body copied to explicit archive | Updated active; old body preserved |
| `Starter.pdf` | Starter.pdf | Historical rendered model; may lag source and new VM target | Label through indexes/audit, no regeneration or binary rewrite | Unchanged historical export |
| `V5/CoveStack-Implementation-Status.md` | CoveStack V5 — Baseline Status, Missing Work and Contract Changes | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Project-Workflow.md` | CoveStack V5 — Delivery Workflow to 25 January 2027 | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Operations-and-Recovery.md` | CoveStack V5 — Operations, Failure Handling and Data Recovery | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-API-and-Compatibility.md` | CoveStack V5 — API Boundaries and Compatibility Matrix | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Database-Schema.md` | CoveStack V5 — Database Schema Reference | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-UI-CLI-and-Monitoring.md` | CoveStack V5 — React UI, CLI Parity and Monitoring | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Team-START-HERE.md` | CoveStack V5 — Team Onboarding and Shared Working Rules | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Terraform-and-Ansible.md` | CoveStack V5 — Terraform, Ansible and Controlled Automation | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/tools.md` | CoveStack V5 — Tools, Versions and Why They Are Used | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Current-Setup-Fixes.md` | CoveStack V5 — Baseline Audit and Extension Checklist | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Publication-Domains-and-Tunneling.md` | CoveStack V5 — Application Publication, Domains and Tunnels | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Testing-and-Acceptance.md` | CoveStack V5 — Test Matrix and Gate Evidence | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-AntiWrapper-Acceptance.md` | CoveStack V5 — State Reconciliation Acceptance and CORE-5 Viva | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-V5-Print.html` | V5/CoveStack-V5-Print.html | Historical rendered model; may lag source and new VM target | Label through indexes/audit, no regeneration or binary rewrite | Unchanged historical export |
| `V5/CoveStack-VPC-IPAM-and-Networking.md` | CoveStack V5 — VPC, IPAM and Connectivity Policy | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Anuj-Next-Steps.md` | CoveStack V5 — Anuj Implementation and ITISS Work Order | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Extended-Roadmap.md` | CoveStack V5 — Extended Development Roadmap | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/Folder_Structure_Explanation.md` | CoveStack V5 — Folder and File Responsibilities | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/DocsV5-DO_Team.md` | CoveStack V5 — Technical Design and Team Architecture | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-References-and-Decision-Log.md` | CoveStack V5 — References, Assumptions and Decisions | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Complete-Working-Setup.md` | CoveStack V5 — Setup, Deployment and Verification | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Service-Catalog.md` | CoveStack V5 — Services, Capabilities and Release Boundaries | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Stack-Templates.md` | CoveStack V5 — One Stack Integration Template | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-PDF-and-Formatting.md` | CoveStack V5 — Markdown and PDF Export | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Resource-Services.md` | CoveStack V5 — Resource Specifications and Service Behavior | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Service-Gates.md` | CoveStack V5 — Service Gates and Promotion Evidence | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Go-Handover.md` | CoveStack V5 — Go Ownership and Learner Handover | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-SRS-Scope.md` | CoveStack V5 — CORE-5 SRS Scope for Integration | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/Team_Roles.md` | CoveStack V5 — Ownership, Reviews and ITISS Responsibilities | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Instance-Metadata.md` | CoveStack V5 — Instance Metadata and Identity Boundaries | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/README.md` | CoveStack V5 — Documentation Index and Reading Guide | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `V5/CoveStack-Security-and-ITISS.md` | CoveStack V5 — Security, Infrastructure and ITISS Evidence | Previous container-compute/two-driver/sizing/schedule authority; some stale source status | Per-topic supersession pointer; preserve full original reference/DDL/tests | Superseded design body retained; active platform replacement |
| `contracts/FROZEN-STATES.md` | CoveStack Cloud — Frozen Lifecycle Contract | Current generated/signing reference describes legacy compatibility, not VM support | Keep exact state/signing reference; route through updated contract README | Unchanged authoritative current-wire documentation |
| `contracts/CANONICAL-REQUEST.md` | CoveStack Cloud — Canonical Request and Signing Guide | Current generated/signing reference describes legacy compatibility, not VM support | Keep exact state/signing reference; route through updated contract README | Unchanged authoritative current-wire documentation |
| `contracts/README.md` | CoveStack Cloud — contracts/ | Incorrect agent-kind count and no VM evolution boundary | Correct prose count, add unchanged 1.1.0 compatibility notice | Documentation only; generated contracts unchanged |
| `deploy/README.md` | CoveStack Cloud — Local Infrastructure | Stale positioning/current-state or dependency-only setup unclear | Current VM-first/evidence-accurate entry point; original body copied to explicit archive | Updated active; old body preserved |
| `agent/README.md` | CoveStack Cloud — Agent placeholder | Stale positioning/current-state or dependency-only setup unclear | Current VM-first/evidence-accurate entry point; original body copied to explicit archive | Updated active; old body preserved |
| `docs/CoveStack-VPC-IPAM-Ingress-Metadata-Proposal.md` | CoveStack: VPC, IPAM, internet access, metadata and dashboard proposal | Old strategy, container model, dates or moved links | Supersession notice and current references; repair relocation links where possible | Historical body retained; link-only repair in two files |
| `docs/development-environment.md` | CoveStack Development Environment | Environment version/check reports can be mistaken for supported product profile | Scope note only; preserve actual results/restrictions | Historical evidence body unchanged |
| `docs/development-toolchain-status.md` | CoveStack Development Toolchain Status | Environment version/check reports can be mistaken for supported product profile | Scope note only; preserve actual results/restrictions | Historical evidence body unchanged |
| `docs/CoveStack-Docker-Difference-and-ITISS-Defense.md` | CoveStack: why build it above Docker, and where is the ITISS contribution? | Old strategy, container model, dates or moved links | Supersession notice and current references; repair relocation links where possible | Historical body retained; link-only repair in two files |
| `docs/Anuj_Working.pdf` | docs/Anuj_Working.pdf | Historical rendered model; may lag source and new VM target | Label through indexes/audit, no regeneration or binary rewrite | Unchanged historical export |
| `docs/Anuj_Working.md` | Anuj Working — CoveStack Agent, Infrastructure and ITISS | Old strategy, container model, dates or moved links | Supersession notice and current references; repair relocation links where possible | Historical body retained; link-only repair in two files |
| `docs/README.md` | CoveStack Cloud — Documentation | Stale positioning/current-state or dependency-only setup unclear | Current VM-first/evidence-accurate entry point; original body copied to explicit archive | Updated active; old body preserved |
| `Workflow_Diagrams/CoveStack-Workflow-and-Gates.svg` | Workflow_Diagrams/CoveStack-Workflow-and-Gates.svg | Historical rendered model; may lag source and new VM target | Label through indexes/audit, no regeneration or binary rewrite | Unchanged historical export |
| `Workflow_Diagrams/CoveStack-Service-Gate-Matrix.svg` | Workflow_Diagrams/CoveStack-Service-Gate-Matrix.svg | Historical rendered model; may lag source and new VM target | Label through indexes/audit, no regeneration or binary rewrite | Unchanged historical export |
| `Workflow_Diagrams/CoveStack-Workflow-Diagrams.md` | CoveStack — Workflow and Gate Diagrams | Old strategy, container model, dates or moved links | Supersession notice and current references; repair relocation links where possible | Historical body retained; link-only repair in two files |
| `control-plane/README.md` | CoveStack Cloud — Control-plane placeholder | Stale positioning/current-state or dependency-only setup unclear | Current VM-first/evidence-accurate entry point; original body copied to explicit archive | Updated active; old body preserved |
| `kan-server/target/surefire-reports/com.ciagotech.kan.server.contract.LifecycleValidatorTest.txt` | kan-server/target/surefire-reports/com.ciagotech.kan.server.contract.LifecycleValidat | Generated prior test evidence, not fresh runtime validation | Record provenance/scope only | Unchanged generated evidence |
| `kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.ValkeyNonceStoreTest.txt` | kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.ValkeyNonceStoreTest | Generated prior test evidence, not fresh runtime validation | Record provenance/scope only | Unchanged generated evidence |
| `kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.AgentAuthSecurityTest.txt` | kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.AgentAuthSecurityTes | Generated prior test evidence, not fresh runtime validation | Record provenance/scope only | Unchanged generated evidence |
| `kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.BearerAuthFilterTest.txt` | kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.BearerAuthFilterTest | Generated prior test evidence, not fresh runtime validation | Record provenance/scope only | Unchanged generated evidence |
| `kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.KanSigV1Test.txt` | kan-server/target/surefire-reports/com.ciagotech.kan.server.auth.KanSigV1Test.txt | Generated prior test evidence, not fresh runtime validation | Record provenance/scope only | Unchanged generated evidence |
| `docs/go-learning/03-running-and-debugging.md` | Run, inspect and debug the Go code | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/sigv1--negative_test.go.md` | Signing security regression tests | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/internal--lifecycle--validate_test.go.md` | Shared lifecycle vector tests | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/sigv1--sign.go.md` | Canonicalization and HMAC signing | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/01-repository-and-runtime.md` | Folder map and how the current code works | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/internal--config--config_test.go.md` | Configuration tests | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/sigv1--request.go.md` | Building signed HTTP requests | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/internal--lifecycle--catalog_generated.go.md` | Embedded generated lifecycle catalog | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/internal--lifecycle--validate.go.md` | Lifecycle validation | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/00-go-basics.md` | Go basics through your CoveStack code | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/cmd--auth-smoke--main.go.md` | Authentication smoke executable | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/sigv1--sign_test.go.md` | Shared signing vector tests | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/02-build-the-next-parts-yourself.md` | Anuj's next steps: learn, practice, implement | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/README.md` | Anuj's Go learning guide — CoveStack Cloud | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/internal--config--config.go.md` | Configuration loading | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/sigv1--rules.go.md` | Shared signing rules | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/go-learning/cmd--auth-smoke--main_test.go.md` | In-memory smoke verifier test | Dated line-by-line source and old next steps may appear current | Mark snapshot; link source/status/current architecture | Prefix only; tutorial body preserved |
| `docs/archive/README.md` | Historical CoveStack documentation | Old V5 authority and mixed historical artifacts | Current index pointer and pre-redesign snapshot scope | Prefix only; existing archive body preserved |

## 3. Archived members inspected and preserved

All ZIP bytes are unchanged. Internal links and original names are historical layout, not live documentation routing.

| Archive | Member | Disposition |
| --- | --- | --- |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Anuj-Next-Steps.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Complete-Working-Setup.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Current-Setup-Fixes.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Documentation-Cleanup.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Project-Workflow.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Renaming.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-START-HERE.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Team-Roles.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/team.zip` | `team/Covestack-Cloud-Technical-Design.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `README.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `contracts/FROZEN-STATES.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `contracts/CANONICAL-REQUEST.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `contracts/README.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `deploy/README.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `agent/README.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/README.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/ANUJ-PERSON-C-HANDOFF.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `control-plane/README.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Anuj-Next-Steps.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Complete-Working-Setup.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Technical-Design.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Repository-Layout.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Team-Roles.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Audit-2026-10-11.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Code-Changes.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-START-HERE.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Lifecycle-Reference.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Project-Workflow.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Current-Setup-Fixes.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Documentation-Change-Log.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Tools.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Branding-Changes.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/Documentation-Before-Tables-2026-10-11.zip` | `docs/team/Covestack-Cloud-Lifecycle-Changes.md` | Read-only historical body; no ZIP rewrite |
| `docs/archive/ANUJ-PERSON-C-HANDOFF.zip` | `ANUJ-PERSON-C-HANDOFF.md` | Read-only historical body; no ZIP rewrite |

## 4. Contradictions and resolution

| Finding | Resolution | Remaining human/implementation decision |
| --- | --- | --- |
| Root README says client/agent pending | Report actual signed client and no-execution loop source | Real server credentials/assignment and live integration |
| Container Droplet versus mandatory VM | Separate VM Instance/App Container, preserve old current-wire specimens | Versioned type/spec/IDs, guest backend and migration with B+D/A/C |
| Fixed two drivers | Explicit VM/container/k3d and network/storage capability boundaries | Interface/driver implementation/conformance |
| Docker-only networking versus VM NICs | Separate runtime attach/ACL/IPAM mapping and no double authority | Actual Incus/container traffic and ownership tests |
| 8 GB/two clusters interpreted as production | Lab constraint/ceiling only, measured capacity admission | Host spike and supported profile qualification |
| January plan versus user 15 December | Retain historical dates; no current calendar guarantee | Confirm year/academic scope/mentor decision |
| Product roadmaps defer safety | Mandatory authorization/capacity/idempotency/crash baseline in 0.1 | Later releases deepen evidence; no unsafe initial provisioning |
| Source library/build equals supported service | Distinct feature evidence, service gates and maturity/profile qualification | Actual integration/runtime/restore and security acceptance |
| Flyway declared but no complete foundation | Two extended SQL files exist; prerequisites missing | Applied-history/migration plan and VM schema, no SQL edits here |
| Go/Java future strategy | Keep Java; selective product-owned Go, optional external Floci | Migration ADR/ownership, no upstream rewrite |
| Open-source marketing without license | Explicit release-rights blocker | Maintainer license/contributor/third-party review |
| Foreign handbook and duplicate guides | Mark external/superseded; canonical docs/platform and checkpoint | Verify handbook permission before redistribution |

## 5. New active documentation and coverage

| File | Purpose |
| --- | --- |
| `docs/platform/README.md` | CoveStack platform documentation |
| `docs/platform/academic-transition.md` | Academic obligations and VM architecture transition |
| `docs/platform/agent-and-reconciliation.md` | Agent execution, reconciliation and recovery |
| `docs/platform/api-and-cli.md` | API, CLI and compatibility boundaries |
| `docs/platform/architecture.md` | Production-target architecture |
| `docs/platform/capacity.md` | Scheduling, capacity and benchmark plan |
| `docs/platform/competitive-analysis.md` | Competitive analysis and differentiation hypotheses |
| `docs/platform/contributing.md` | Development and contribution guide |
| `docs/platform/decisions.md` | Architecture decision and implementation gap register |
| `docs/platform/feature-specifications.md` | Feature specifications and implementation boundaries |
| `docs/platform/implementation-status.md` | Implementation and evidence ledger |
| `docs/platform/installation-and-operations.md` | Installation, deployment, operations and troubleshooting |
| `docs/platform/license-and-attribution.md` | License, attribution and release-rights review |
| `docs/platform/network-and-storage.md` | Tenant networking, storage and backup |
| `docs/platform/readiness.md` | Maturity, acceptance and production readiness |
| `docs/platform/references.md` | Research references and evidence limits |
| `docs/platform/resource-model.md` | VM Instance and App Container resource model |
| `docs/platform/roadmap.md` | Product release roadmap and dependency gates |
| `docs/platform/security.md` | Security target and threat model |
| `docs/platform/strategy.md` | Product strategy and user validation |

The 20 active platform documents cover positioning/users, fair competition, HLD/control/data planes, VM/container lifecycle, identity/threat model, network/storage/backup, agent/claims/fencing, measured scheduling, install/operations/troubleshooting, API/CLI, feature specifications, maturity/acceptance, roadmap, contribution, rights/attribution, academic transition, evidence, decisions and researched references. Root README/overview and nested indexes route to them. No additional public service or unofficial API enum is implemented.

## 6. Preservation, relocation and exceptions

- V5's 31 bodies remain intact below notices; no reference DDL was changed. Old service/test/requirement numbers remain historical obligations pending VM mapping.
- Original bodies of six rewritten entry points are saved under docs/archive/pre-vm-redesign-2026-11-04, with an index and historical notices. No prior requirement is silently removed.
- 38 bad links in docs/Anuj_Working.md and Workflow_Diagrams/CoveStack-Workflow-Diagrams.md were corrected only for relocation; wording and prior results remain.
- HTML/SVG/ZIP/generated report bytes were preserved. During this task the user intentionally removed Starter.pdf, docs/Anuj_Working.pdf and Kan_Cloud_Database_Schema_Simple_Guide.pdf, confirming this in chat because Markdown versions suffice. The agent did not remove/restore those files. They remain in the baseline inventory as inspected historical artifacts; exports were not regenerated.
- Environment/checkpoint historical results were preserved; static source inspection is not a fresh test.
- Baseline missing targets: 65 loose Markdown references. Remaining missing links are explicitly historical external handbook names or archive snapshot links reflecting original root layout, not active instructions. Detailed exceptions follow verification; no missing local implementation file was fabricated.
- Pre-existing .gitignore and checkpoint changes were present. Compare task-start hashes, not just git diff, to determine this task's scope. No ignore rule changed; some local/ignored documentation may not appear in git status.

## 7. Verification and counts

Verification is static documentation validation only: table widths/fences, active file/anchor links, semantic authority, historical-body hashes, all-file fingerprints and git diff. Application/unit/infrastructure tests and renderer scripts were deliberately not run. Source/config/API/SQL/build/CI/lock files must retain task-start hashes. Counts and exact historical link exceptions are recorded after the final check below.

### Final documentation validation record

| Check | Observed result / scope |
| --- | --- |
| Existing document bodies modified | 67 existing Markdown files updated, including local/ignored learning/overview files; prefixes preserve historical bodies where specified |
| New documentation files | 28 Markdown files: 20 platform docs, 7 archive copies/index, this audit |
| Baseline inspected | 80 loose artifacts plus 34 archived documents; ZIP containers also inventoried |
| Active local links/anchors | No missing active file targets or heading anchors; all historical links also scanned |
| Historical missing links | 48 retained historical references: foreign handbook sibling names and pre-redesign snapshots in their old layout; table below |
| Relocation repairs | 38 links fixed in two historical guides; no requirement wording removed |
| V5 body preservation | All 31 original bodies remain intact below notices; reference DDL unchanged |
| Other historical preservation | 57 original docs receive notices; two have link-only relocation repairs; environment reports and prior checkpoint body unchanged |
| Files lost outside edits | Three PDFs intentionally removed by the user during work, confirmed in chat; not agent edits |
| Scope | All files written by this task are Markdown; protected implementation/config/contracts/SQL/CI/build/locks retain starting hashes; pre-existing .gitignore change untouched |
| Git/infra/runtime | No branch/commit/push/merge/install/provisioning or application/acceptance tests; Git status/diff inspection only |
| Documentation checks | Structural/link/semantic/fingerprint checks pass; final git diff whitespace check performed separately |

“Inspected” does not imply competitor source audit, fresh tests or visual PDF/HTML certification. Semantic review treats text explicitly under a historical marker as dated history, not active target instructions. Generated signing/state docs remain current-contract authority and are not rewritten into the VM target.

### Historical link exceptions

These are retained source/history references, not missing active features. No placeholder implementation file is created merely to satisfy a link checker.

| Historical document | Original target not present in that layout |
| --- | --- |
| `HANDBOOK.md` | `SERVICE_IMPLEMENTATION_ROADMAP.md` |
| `HANDBOOK.md` | `FEATURES.md` |
| `HANDBOOK.md` | `DASHBOARD.md` |
| `HANDBOOK.md` | `NODEJS_LOCAL_CI.md` |
| `HANDBOOK.md` | `FEATURES.md` |
| `HANDBOOK.md` | `EC2_TEMPLATES.md` |
| `HANDBOOK.md` | `EC2_INFRASTRUCTURE.md#ssh-into-an-os-like-guest` |
| `HANDBOOK.md` | `EC2_NETWORKING.md` |
| `HANDBOOK.md` | `EC2_RELEASE.md` |
| `HANDBOOK.md` | `EC2_INFRASTRUCTURE.md` |
| `HANDBOOK.md` | `KUBERNETES.md` |
| `HANDBOOK.md` | `EC2_RELEASE.md` |
| `HANDBOOK.md` | `SERVICE_IMPLEMENTATION_ROADMAP.md` |
| `HANDBOOK.md` | `CLI.md` |
| `HANDBOOK.md` | `HTTP_API.md` |
| `HANDBOOK.md` | `NODEJS_LOCAL_CI.md` |
| `HANDBOOK.md` | `CI_CD.md` |
| `HANDBOOK.md` | `DASHBOARD.md` |
| `HANDBOOK.md` | `KUBERNETES.md` |
| `HANDBOOK.md` | `EC2_RELEASE.md` |
| `HANDBOOK.md` | `EC2_INFRASTRUCTURE.md` |
| `HANDBOOK.md` | `EC2_NETWORKING.md` |
| `HANDBOOK.md` | `EC2_TEMPLATES.md` |
| `HANDBOOK.md` | `ENGINE_RE.md` |
| `HANDBOOK.md` | `FEATURES.md` |
| `HANDBOOK.md` | `SERVICE_IMPLEMENTATION_ROADMAP.md` |
| `HANDBOOK.md` | `EVIDENCE.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/CoveStack-SRS-Scope.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/DocsV5-DO_Team.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/CoveStack-Database-Schema.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/CoveStack-AntiWrapper-Acceptance.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/CoveStack-Stack-Templates.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/CoveStack-Complete-Working-Setup.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/COVESTACK-PROJECT-OVERVIEW.md` | `V5/CoveStack-Anuj-Next-Steps.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/docs--README.md` | `../V5/CoveStack-Team-START-HERE.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/docs--README.md` | `../V5/CoveStack-Anuj-Next-Steps.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/docs--README.md` | `../V5/CoveStack-Security-and-ITISS.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/docs--README.md` | `../V5/CoveStack-Go-Handover.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/docs--README.md` | `../V5/CoveStack-Anuj-Next-Steps.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/docs--README.md` | `archive/README.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `V5/CoveStack-Team-START-HERE.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `V5/README.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `docs/archive/README.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `V5/CoveStack-Anuj-Next-Steps.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `V5/CoveStack-Security-and-ITISS.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `contracts/lifecycle-rules.json` |
| `docs/archive/pre-vm-redesign-2026-11-04/README-original.md` | `contracts/FROZEN-STATES.md` |
| `docs/archive/pre-vm-redesign-2026-11-04/deploy--README.md` | `../V5/CoveStack-Complete-Working-Setup.md` |

### Team review required

| Document / decision | Reviewers and reason |
| --- | --- |
| resource-model.md / architecture.md | A+B+C/D: VM versus App Container, Incus privileges, provider/claim/network/spec boundaries and ownership |
| academic-transition.md / roadmap.md | Team and mentor: preserve CORE-5, map VM tests/stack, confirm 15 December versus January schedule; no approval inferred |
| security.md / network-and-storage.md | C+B+A/D: credential custody, default-deny traffic, disk/object/backup ownership and residual risks |
| capacity.md | A+B+C: actual host measurements, admission/reservations, overcommit, scheduler boundary and profile limits |
| api-and-cli.md / feature-specifications.md | B+D+A/C/E: exact future schemas/DTOs, mixed versions, migration and user CLI behavior |
| readiness.md / installation-and-operations.md | A+B+C/D/E: measurable thresholds, upgrade/rollback, supported matrix and authorized runtime runners |
| strategy.md / competitive-analysis.md | Maintainers and pilot users: demand, fair comparisons and actual additional value |
| license-and-attribution.md | Maintainers/rights holders: root license, contributor and third-party rights, disclosure/contact policy |

No document is a completion certificate. Current blockers remain runnable Java/identity/foundation, coordinated future protocols, actual VM/container/network/storage implementations, safe capacity/claims, runtime failure/restore evidence, a license grant and confirmed academic mapping. No SQL/state/DTO/backend code was implemented to make the documentation appear complete.

### Git visibility limitation

23 modified pre-existing local/ignored documents (the root overview, foreign handbook, old diagram/proposal/Anuj guides and 18 local learning guides) do not appear in ordinary Git status. Their contents were inspected/updated locally as requested; ignore policy was not changed. All 28 newly created Markdown documents are visible/unignored. Root README/docs index route to unignored docs/platform rather than require the ignored root overview. Future publication of those original local files requires an explicit maintainer decision, not an unsolicited force-add or ignore edit.

The required feature specification list contains 14 features; all 14 have explicit scope, problem/story/outcome, API/data/runtime/security/failure/test/status/dependency/release fields. Documentation verification uses that list, not an invented feature count.
