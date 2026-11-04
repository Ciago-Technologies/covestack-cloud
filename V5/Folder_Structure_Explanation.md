# CoveStack V5 — Folder and File Responsibilities

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Names and present-versus-planned rule

Product name is CoveStack. Existing kan-server, kan-agent, Go module/Java package/KAN1/header identifiers remain compatibility names until coordinated change. VM Instance and App Container are product types, not permission to rename current wire kinds. A planned path is not a file that already exists.

## 2. Existing repository surfaces

| Path | Actual responsibility / evidence | Owner |
| --- | --- | --- |
| kan-agent/cmd/kan-agent/ | Existing config/client/cancellable polling entry point; returned work stops safely | A |
| kan-agent/cmd/auth-smoke/ | Signing diagnostic and local verifier test; real Java proof separate | A; C acceptance |
| kan-agent/internal/config/ | Environment/rules validation and credential-safe formatting | A |
| kan-agent/internal/client/ | Bounded signed GET work client and HTTP/schema tests | A; B+C/D review |
| kan-agent/internal/agent/ | No-execution polling/backoff/auth-stop/cancel and tests | A |
| kan-agent/sigv1/ | Existing request/rules/signing/vector tests | A+C |
| kan-agent/internal/lifecycle/ | Pure shared validator and generated catalogue/tests | A+B/C/D |
| kan-server/src/main/java/.../auth/ | Existing auth/replay interfaces/filters/helpers, not full app server | C+B |
| kan-server/src/main/java/.../contract/ | Java lifecycle helper and generated data | B+D/C |
| kan-server/src/test/java/ | Existing Java unit/mock/vector tests | C; B/A/D inputs |
| kan-server/src/main/resources/db/migration/ | Existing publication/domain SQL only; foundation prerequisites unresolved | B SQL; C+A/D review |
| contracts/ | Authoritative base/lifecycle/signing, examples, generator/checkers and generated references | D+B+C; A consumer |
| deploy/ | Dependency Compose/examples for platform DB/Valkey/Mailpit, not complete product | A infra; B/C |
| .github/workflows/ | Current Go/Java/contract checks and secret scan | A; C acceptance |
| docs/product/ and docs/platform/ | Canonical product/domain/evidence specifications | Team ownership per topic |
| V5/ | Current detailed implementation/team documentation aligned with canonical specs | Team review; docs not runtime proof |

## 3. Planned layout — no creation/move in this task

| Target area | Purpose / decision before code |
| --- | --- |
| Shared Go API/client module | Reuse protected HTTP/signing where appropriate without confusing user and agent auth; module/import review A+B+C+E |
| User CLI/provider | Thin API-only Go consumer and later optional provider; no runtime drivers/DB |
| VM/App Container/k3d providers | Trusted execution interfaces, capabilities, owned inspection/task/readiness/conformance; A implementation |
| Java app/controllers/repositories | Real authenticated resource/assignment/operations and SQL foundation under B/C/D |
| Host/capacity/network/storage model | Assignment/reservations/attachments/traffic enforcement and recovery; no fixed two-driver ceiling |
| Installer and future inventory | Operator installation versus account-scoped API inventory; actual service/artifact first |

Do not create guessed filenames/modules/migration numbers from this table. [Existing-code plan](../EXISTING-CODE-CHANGE-PLAN.md) deliberately maps actual present files; missing components get separately reviewed interfaces/paths.

## 4. Generated versus handwritten

contracts/lifecycle-rules.json and internal-api.base.json drive generated internal-api.yaml, FROZEN-STATES.md, Go catalog_generated.go and Java LifecycleData.java. Change authoritative inputs/fixtures/consumers together when approved; do not hand-edit generated outputs. Signing rules/vectors remain shared. Compiled target/classes/jars and __pycache__ are build artifacts, not edit targets or fresh runtime proof.

## 5. Secrets and data placement

Real .env/deploy secrets/signing/master/private keys, credential kubeconfigs, Terraform state and raw sensitive logs never enter source/issues/ordinary docs. Protect provider/custody data and independent backups; reviewed runtime storage is not a random host directory supplied by tenant. Do not change ignore rules or delete volumes to hide setup errors. Keep versioned nonsecret examples accurate without fake executable credentials.

## 6. Human and AI documentation flow

Starter.md: clone/branch/start. TEAM-FIRST-CHECKPOINT.md: owned shared tasks/board/acceptance. checkpoint.md: single AI continuity file. GitHub issues: live status. V5/catalogue/domain docs: current requirements; historical snapshot: old text only. The [V5 index](README.md) and [root documentation flow](../DOCUMENTATION-IMPLEMENTATION-FLOW.md) link the reading path. No obsolete-file deletion is implied by this map.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-folder-structure-explanation) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
