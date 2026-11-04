# Historical release planning reference - superseded

> The numbered release plan below was produced by the previous documentation task. It is not an active commitment and is not extended or regenerated here. Current product scope is the [canonical catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md); future work is a [dependency-only index](future-services.md), with no release sequence or Labs scope. Academic obligations/deadline remain subject to explicit review.

<!-- Historical roadmap body preserved below. -->

# Product release roadmap and dependency gates

Updated: 4 November 2026. **All releases are targets, not shipped capability claims or date commitments.** Academic scope/deadline is separate in [academic transition](academic-transition.md). Keep Java control plane and existing Go agent; no Floci rewrite.

## Safety is not postponed

The proposed release grouping is adjusted to avoid an unsafe 0.1: basic authenticated ownership, atomic intent/idempotency, capacity reservation, bounded claims and crash-safe create/cleanup are prerequisites for any actual provisioning, not features withheld until 0.2/0.3. Later releases deepen those guarantees and evidence. Unfinished portions remain Experimental; a release tag never substitutes for gates.

| Release target | User outcome and scope | Dependencies | Exit acceptance | Main risk |
| --- | --- | --- | --- | --- |
| 0.1 reliable single-host foundation | Real VM Instance via selected backend, separate trusted App Container, one owned app/PG/bucket stack, basic CLI, real readiness and safe cleanup | Incus evaluation; versioned VM/container/claim/network contracts; server boot and foundational migrations; credential custody; capacity/isolation baseline | Independent clean experimental install; actual guest kernel and app data path; basic account denial; same-key/interrupt create yields one resource; cleanup respects retention | VM host requirements and wide integration scope; narrow declared subset rather than claim all targets passed |
| 0.2 durable control plane | Drift scheduling, richer operations, uncertain-action settlement, extended generation/fencing tests and measured VM/data/control DB restore | 0.1 safety path and persisted provider-task identity | Paused worker/lost response/restart/concurrent mutation tests; restored known data and key recovery | Assuming DB epoch fences backend actions; snapshot mistaken for backup |
| 0.3 isolation/security qualification | Full reviewed RBAC, rotated/revoked credentials, default-deny VM/container boundaries, storage credentials, quota/audit evidence | 0.1 mandatory baseline plus threat-model decisions and security runner | Two-account API/packet/disk/backup denial; secret rotation, management-route blocking, audit privileges and security review | Host/agent compromise; settings may not enforce packets; no hostile-tenant claim without evidence |
| 0.4 infrastructure expansion | Required academic Kubernetes mapped to product lifecycle, richer network/storage capability adapters, optional Floci environment, scheduling refinements | Core runtime gates and measured new capacity; compatibility ADR | Real cluster workload/cleanup; adapter conformance; emulator opt-in cleanup and scope; concurrent capacity tests | Over-expanding service list, network conflicts and backend capability drift |
| 0.5 developer experience | Better CLI, eventual UI, reviewed templates, useful logs/monitoring, recovery interface and installer assistance | Stable API, secret delivery and operation evidence | User workflow survives refresh/retry; truthful freshness, accessible error handling, no UI-only success | Interface work conceals unfinished control/data paths |
| 1.0 stable open-source target | Reproducible installer, profile matrix, operations, upgrades/rollback, signed releases, SBOM, contribution/disclosure and license | RC and readiness evidence reviewed; distribution rights decided | Every declared feature/profile meets readiness matrix; known gaps and security policy published | A tag or README implies support beyond tested profiles |

## Execution sequence

1. Resolve VM compatibility, ownership/network/storage model, Incus privileged boundary and academic scope impact.
2. Finish real Java/Go signed no-execution integration using current contract; do not manufacture claims that 1.1.0 cannot represent.
3. Run authorized VM/backend host spike on suitable separate hardware; document negative findings and capacity, then freeze reviewed interfaces.
4. Implement minimal owned VM and distinct App Container lifecycle with mandatory safety baseline; test one resource before multi-account concurrency.
5. Integrate PG/MinIO, fixed app stack, network ownership, recoverable reservations and CLI.
6. Expand recovery/isolation evidence, then additional providers and DX under their gates.

Basic scheduler/reservation exists before first VM; a later Go component is a reviewed extraction, not permission for two independent schedulers. Kubernetes remains an academic obligation even if product packaging lists expansion in 0.4. Reordering a product milestone does not waive academic delivery. No instruction here guarantees completion by 15 December or the older January plan.

## Release review and cuts

For each release publish actual implemented subset, unit/integration/runtime evidence, supported host matrix, omitted capabilities and migration/rollback limits. If late, reduce a proposed product release subset explicitly; do not cut security/data-loss protections or quietly remove signed-off academic requirements. Academic revisions require mentor approval, which is not recorded. No new people, dates or completion percentages are invented.

Go growth applies only to product-owned components. Optional Floci remains an external emulator with its own version/license and integration tests. Cloud API parity is not the product roadmap.
