# CoveStack V5 — Current Integration Blockers and Ordered Fixes

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Starting evidence

Existing signed client/config/signer/lifecycle helpers, cancellable no-execution loop and main are retained. Java auth/replay/contract components and V5/V6 extended SQL are partial foundation evidence, not a running owned infrastructure API. VM/App Container/k3d/RustFS/provider/CLI/scheduler runtime acceptance remains absent. This register describes pending work, not fixes already applied.

## 2. Ordered blocker register

| Priority / owner | Gap to resolve | Existing surfaces to inspect | Required output |
| --- | --- | --- | --- |
| First — B+C+A/D | Real server origin, protected credential provisioning and trusted host/account assignment | Java KeyService/BearerTokenService/auth config; Go config/client; current internal API | Written agreement plus actual runnable authenticated discovery/no-work path |
| First — B | Runnable Java entry point and migration prerequisites | pom repackage skip, source packages, V5/V6 FK dependencies | Reviewed startup/foundation/clean migration strategy; never assume live schema |
| First — C+B | Concrete token/key store, custody/revoke/replay and membership | Existing service interfaces/filters/ValkeyNonceStore | Real lookup/assignment/negative tests without default key |
| Next — A+C | Real loop integration and transport-error policy | Existing loop/client/tests | []/auth/cancel/outage evidence; preserve no-execution stop and bounded fresh-signature retry |
| Before execution — B+D+A+C | VM/App Container/capability/claim/renew/report/network/capacity schema | Contract base/catalogue/generator/examples, Java/Go validators | Versioned reviewed design and coordinated fixtures/outputs; no fictional wire version |
| Before VM — A+B+C | Authorized Incus/profile/image/disk/network/capacity spike | No existing VM driver; current host not qualified | Real guest kernel/readiness, privilege and unknown-task recovery evidence |
| Before data services — B+C+A | Java RustFS admin adapter, private root custody, volumes and consistent restore | No object adapter/process deployment established | Actual object/key/policy/restart/isolated restore, not config-only recreation |
| Before CLI production — E+A/B/C/D | Module/shared-client/user API boundary | No current CLI/provider module | Explainable API-only client, JSON/wait/idempotency and real user auth |

## 3. File-level work plan and ownership

Use [EXISTING-CODE-CHANGE-PLAN.md](../EXISTING-CODE-CHANGE-PLAN.md) while it exists for inspected exact current files/actions. It deliberately does not invent missing driver/controller/repository/migration filenames. Permanent requirements are this V5 set plus canonical product/domain specs; migrate residual checklist decisions to issues/docs before deleting the temporary plan. B owns all SQL, C all acceptance, A agent Go, D Java API, E CLI.

## 4. What must not be “fixed” by bypass

No anonymous/hardcoded credential, in-memory replay fallback, arbitrary account selector, direct CLI DB/backend, container-as-VM, unverified guest image, unsafe public route, deleting DB volumes, skipping failure tests or editing generated catalogue alone. Do not rewrite shared signer/modules just to hide unresolved auth/consumer design. A green build is not real server or runtime acceptance.

## 5. Next integration and verification

Immediate shared exit is real signed no-work/auth/cancel with existing loop; returned work must still stop safely. Current Go retries timeout-class transport/selected HTTP only, not universal network failures; observe and agree before change. CI commands and exact evidence are in [setup](CoveStack-Complete-Working-Setup.md) and [tests](CoveStack-Testing-and-Acceptance.md). No implementation/test/infrastructure work ran in this docs task.

## 6. Preserve data and current decisions

VM-first product is independent of old academic dates/counts/caps. RustFS Planned, alternatives behind provider boundary. Apache-2.0 adopted; third-party notices still reviewed. Existing .gitignore, secrets, source/contracts/deploy and untracked user release work remain untouched. No supported profile, migration history or backend runtime is inferred.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-current-setup-fixes) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
