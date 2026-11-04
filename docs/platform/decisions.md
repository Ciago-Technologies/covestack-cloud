# Architecture decision and implementation gap register

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. Decision labels distinguish user-selected direction from approval of detailed schemas/backend versions or production acceptance.

| Decision | Disposition | Required review / evidence |
| --- | --- | --- |
| Local-first owned infrastructure, not emulator replacement | Selected product direction | Validate target users/pilots; no adoption guarantee |
| Primary VM Instance, distinct App Container | Selected by user; implementation Planned | B+D/A/C versioned resource/spec/IDs, old-container compatibility and engineering migration review |
| Incus preferred VM management backend | Preferred evaluation, not supported | A+C/B host spike, privileges, task recovery, resource enforcement and actual readiness |
| QEMU/KVM under VM backend | Evaluation stack, no direct shell driver chosen | Supported Linux/virtualization/profile and signed guest-image path |
| Java remains control plane | Retained | No language rewrite; current server boot/persistence still missing |
| New product-owned Go components | Long-term direction | Scheduler/service extraction ownership, atomic state and upgrade/rollback before code |
| Backend interfaces beyond two drivers | Selected target | Versioned capabilities and conformance; no untrusted dynamic plugin execution |
| Current wire 1.1.0 unchanged | Compatibility boundary | Future VM/container, claim/renew/report/capacity/network schema/fixtures/generated output review |
| Database and migration chain | Blocked foundation | B checks applied history; accounts/resources prerequisites, VM-owned records/reservations; never edit applied migrations |
| Capacity profile | No production sizing approved | A/B/C measurements; 8 GB lab only; conservative initial overcommit policy and rejection behavior |
| Identity and keys | Partially Implemented | C+B actual token/credential store, account-host assignment, TTL/rotation/revocation/master-key custody |
| Network enforcement and IPAM | Planned | C allocator, A Incus/container adapters, B transactions; default-deny packet tests; avoid double allocation |
| Storage/backup | Planned | Owner/disk/restore model, encryption/backend/key and retention policy; real restore |
| Object-storage provider/custody and rights | RustFS preferred; integration Planned | B+C/A exact artifact/admin API, private management, scoped policy/key rotation, persistent volume and verified restore; alternatives remain behind adapter |
| Optional Floci | Proposed integration | Pinned external artifact, endpoint isolation, image/license, state/cleanup and no production AWS credentials |
| Product scope independent of academic CORE-5 | Selected by latest user instruction | Core/future catalogue governs functionality; old SRS/tests remain historical, not product limits or approval gates |
| Historical academic dates | Reference only | Preserve old date records; no academic date is a product deadline and no new product schedule is invented |
| Product license/security disclosure | Apache-2.0 approved by maintainer; root LICENSE added | Third-party dependency/NOTICE and copied-material rights review, contribution governance, maintainer/contact and response process remain; see licensing-decision.md |
| Stable profile/HA | None / future | Readiness gates, profile matrix; separate HA architecture and partition/host-loss evidence |

## What documentation cannot establish

No live migration history, exact production kernel/backend pair, VM density, actual network isolation, real credential provisioning, successful VM/data restore, end-user demand, mentor approval or remote CI/PR state was established. The latest maintainer instruction approves Apache-2.0; the official root LICENSE now records adoption. Historical environment test results remain historical. Code changes must be separately authorized and owned; no schema/manifest/build/CI/source changes were made here.

Resolve only decisions blocking the next authorized slice; avoid inventing implementation detail to make tables look finished. [Feature specs](feature-specifications.md) and [readiness](readiness.md) define the evidence each decision unlocks.
