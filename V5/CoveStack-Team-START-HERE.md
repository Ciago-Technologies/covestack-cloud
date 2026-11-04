# CoveStack V5 — Team Onboarding and Shared Working Rules

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Read first

Start [root Starter.md](../Starter.md) for clone/baseline/branch/check/PR instructions. Then [TEAM-FIRST-CHECKPOINT.md](../TEAM-FIRST-CHECKPOINT.md) for each owner, shared integration acceptance and GitHub card template. [V5 index](README.md) is the current detailed architecture/team map, not a history-only folder.

| Order | Document | What to learn |
| --- | --- | --- |
| 1 | [Product SRS](CoveStack-SRS-Scope.md) and [catalogue](CoveStack-Service-Catalog.md) | Core/future product scope independent of academic limits |
| 2 | [Main design](DocsV5-DO_Team.md) and [resource behavior](CoveStack-Resource-Services.md) | Real VM versus App Container; Java/Go/provider authority |
| 3 | [Roles](Team_Roles.md) and [workflow](CoveStack-Project-Workflow.md) | First issue/owner/dependency and no invented dates |
| 4 | [Setup](CoveStack-Complete-Working-Setup.md) and [status](CoveStack-Implementation-Status.md) | Current code versus missing runnable foundation/runtime |
| 5 | [API](CoveStack-API-and-Compatibility.md) and [schema](CoveStack-Database-Schema.md) | Actual 1.1.0 versus proposed VM/claims/owned DB model |
| 6 | [Gates](CoveStack-Service-Gates.md), [acceptance](CoveStack-AntiWrapper-Acceptance.md), [security](CoveStack-Security-and-ITISS.md) | Required real failure/isolation/recovery evidence |

## 2. First shared work

A already has signed config/client/main/no-execution polling. B+C must supply reviewed real Java foundation/credentials/trusted assignment; D reviews actual discovery/fixtures. A+C prove signed []/negative auth/cancel/outage without resource execution. E can agree thin module/client layout and learn labelled mocks. Do not recreate existing agent code or turn a fake endpoint/default credential into accepted integration.

## 3. Shared product rules

VM Instance is actual guest-kernel compute target; App Container separate. RustFS is preferred Planned provider, operator process/volumes separate from Java tenant bucket/policy/credential reconcile. CLI/UI/providers/inventory use only authenticated API; backend sockets and platform DB remain private. B owns SQL, C acceptance, D no Go, E CLI not agent; A review all Go. Host capacity is measured, not inferred from 8 GB/two-cluster academic history. Product scope has no academic service/template/network/node ceiling.

## 4. Branch and review workflow

Confirm Anuj's reviewed baseline, clone whole repo and branch for one task; never push into a subdirectory or directly to main. Each issue has owner, concrete assertion, dependency/blocked-by and exact evidence scope. Board is view of issues, not another state registry. A PR/CI record establishes what ran; file presence and board Done do not establish service support. No remote issue or branch state is claimed by this guide.

## 5. Documentation and secrets

Share repository links on the agreed branch/revision, not stale ZIP copies. V5 remains detailed current guidance; canonical product/domain specs resolve shared scope/evidence. The previous V5 snapshot is explicitly historical. Keep root checkpoint.md as AI handoff and issues as live task status. Never share real .env/deploy secrets/signing/private/admin keys, secret kubeconfigs/state or raw logs. Apache-2.0 applies to CoveStack-owned material, not upstream projects.

## 6. First task per owner

| Person | Immediate slice |
| --- | --- |
| A | Explain existing loop; B+C handshake and real signed no-work integration; retain no-execution guard |
| B | Startup/foundation/migration prerequisite and real authenticated server slice |
| C | Credential/assignment/custody/negative auth and acceptance from day one |
| D | Existing endpoint/DTO/fixture inventory and future VM/claim compatibility review |
| E | Agree module/user-client boundary with A; small explainable mock CLI learning slice |

Read exact details in [team roles](Team_Roles.md) and root team checkpoint. Resource execution, new schema/driver/module creation and host provisioning require separately reviewed tasks; this documentation update does not authorize them.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-team-start-here) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
