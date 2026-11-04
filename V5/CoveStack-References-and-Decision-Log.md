# CoveStack V5 — Authority, References and Open Decisions

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Document authority

| Material | Authority |
| --- | --- |
| Latest explicit user decisions | Independent VM-first infrastructure product, Planned RustFS, Apache-2.0 adoption and retained useful V5 detail |
| Canonical product/domain docs | Shared scope/architecture/security/capacity/requirements/evidence |
| V5 active bodies | Detailed working roles, setup, API/schema/concurrency/service/acceptance and topic guidance aligned to canonical docs |
| Actual contracts/source/PR/CI/runtime artifacts | Existing wire behavior and actual implementation/test scope; documentation cannot override them |
| Historical snapshots/exports/old plans | Past requirements, alternative designs and results only; not current product limitations |
| GitHub issues / root checkpoint | Live task status / single AI continuity, respectively |

No mentor confirmation, market advantage, production profile, deadline or new wire version is inferred from an architecture document. Cite the exact revision and verified evidence type when claiming behavior.

## 2. Primary technical and licensing references

| Source | Purpose / evidence boundary |
| --- | --- |
| [Incus docs](https://linuxcontainers.org/incus/docs/main/) | VM/container/projects/storage/network/security; not an implemented CoveStack backend |
| [KVM](https://www.linux-kvm.org/page/Main_Page) / [QEMU](https://www.qemu.org/docs/master/system/introduction.html) | Virtualization foundations; actual hardware/profile measured separately |
| [RustFS compatibility](https://docs.rustfs.com/en/reference/s3-compatibility) / [IAM](https://docs.rustfs.com/en/security-compliance/iam) | Tested subset, credential/policy model and explicit unsupported/planned behavior; not CoveStack runtime acceptance |
| [Provider evaluation](../docs/platform/provider-evaluation.md) | Private administration/custody, alternatives and real restart/restore requirements |
| [Apache-2.0 official text](https://www.apache.org/licenses/LICENSE-2.0.txt) / [root LICENSE](../LICENSE) | Exact maintainer-approved project license; does not relicense upstream projects |
| [Competitive official-source review](../docs/platform/competitive-analysis.md) | Fair existing capabilities/uncertainties; no false missing-feature claim |
| [Canonical research references](../docs/platform/references.md) | Additional reviewed official sources and failed-lookup limits; latest docs are not support pins |

No fresh network research, dependency install or runtime verification was performed for this alignment; existing verified source references are reused. Revalidate exact selected artifact/security/version documentation before qualification.

## 3. Current decisions and unresolved encoding

| Decision | Disposition | Owner / next evidence |
| --- | --- | --- |
| Independent core/future product scope | User-selected; no academic ceiling | Keep canonical 22/31 classification; no separate microservice count promise |
| True VM primary, distinct App Container | User-selected; Planned execution | B+D+A/C reviewed spec/migration; A real guest/provider spike |
| Incus over QEMU/KVM | Preferred evaluation | Approved host/image/privilege/capacity/task/restore tests |
| Java control plane / Go agents/clients | Retained; no wholesale rewrite | New extraction/scheduler ownership/transaction boundary reviewed separately |
| RustFS object backend | Preferred Planned provider | B+C/A artifact/admin API/policy/key/volume/private management and verified byte restore |
| API 1.1.0 | Actual unchanged wire | B+D+A/C future type/claim/assignment/network/capacity and fixtures/generated consumers |
| One agent host, isolated account contexts | Target model, concrete assignment pending | B+C+A secure enrollment/dispatch/credential loading and failure isolation |
| Claim/ack/unknown outcome | Exact generation/owner/epoch and separate progress target | B+C+A expiry/renew policy, runtime task exclusion and real stale-worker tests |
| Schema / IDs / capacity / ranges / retention | Reference design, not migration-ready | B finalizes with C tenancy and A runtime/D DTO; actual DB history before migration |
| CLI/shared module/provider | Planned; E thin user client | A+B+C+E module/auth review, D fixtures; API-only lifecycle first |
| Apache-2.0 | Maintainer approved; official root LICENSE present | Exact transitive/image/notices/copied-reference provenance still reviewed |
| Production profile / HA / demand | Unqualified / future / hypothesis | Clean install, security/endurance/restore/upgrade and pilot comparison evidence |

## 4. Evidence and change policy

Current Go loop safely stops on work and does not provision. Pure lifecycle validators do not enforce DB ownership/leases/readiness. Existing extended SQL is not a complete migration chain. HMAC is not a tenant authorization or host assignment grant. Runtime operations require intent, physical admission, exact current authority and actual backend inspection/readiness.

Any future message/schema change updates authoritative sources, fixtures, generated outputs and consumers together; no silent API strings or supported option claims. Current roles remain A agent Go, B all SQL/Java core, C security/acceptance/Java allocator, D Java API no Go, E CLI/packaging/later provider. [Existing-file plan](../EXISTING-CODE-CHANGE-PLAN.md), [gate evidence](CoveStack-Service-Gates.md).

## 5. Historical relocation record

All prior V5 bodies—including original 32-table/extended DDL, 38-issue audit text, old calendar/caps, previous extension/service assumptions and relocated design alternatives—are preserved unchanged in the single snapshot linked below. Active bodies have actual updated content, not transition notices over old instructions. No unique historical requirement/test result is deleted; obsolete applicability is explained and current requirements are cross-linked.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-references-and-decision-log) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
