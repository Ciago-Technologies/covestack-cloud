# Maturity, acceptance and production readiness

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **Current maturity: Experimental. No supported production deployment profile.** Historical unit/lint reports are not live infrastructure acceptance. This document defines future gates; no application or infrastructure tests ran for this documentation redesign.

## Status has three independent dimensions

| Dimension | Values | Interpretation |
| --- | --- | --- |
| Service/evidence | Proposed; Planned; In Progress; Partially Implemented; Implemented; Runtime Verified; Acceptance Gates Passed; Supported; Deferred; Blocked | Record source and domain-specific runtime/automated evidence separately; none advances from documentation alone |
| Historical mapping | V5 In progress → In Progress; Gates passed → Acceptance Gates Passed; Extended track → Deferred/Proposed by current classification | Historical bodies are unchanged; Supported additionally needs declared subset, qualified profile and maintenance policy |
| Product maturity | Experimental; Alpha; Beta; Release Candidate; Stable 1.0; Supported production deployment profile | Maturity and particular host-profile qualification are not automatic from a feature status |

G1: DB desired state/generation authoritative. G2: crash-safe inspection/rebuild without duplicate backend objects. G3: account/RBAC denial (foreign API resources 404) plus actual network/storage credential isolation. G4: meaningful failure acceptance runs in CI on an authorized capable runner. New VM acceptance needs runtime evidence; existing container mock/vector tests cannot pass it.

## Measurable maturity exits

| Stage | Required exit evidence | Current disposition |
| --- | --- | --- |
| Experimental | Truthful source inventory, contract/architecture decisions and documented limits; no support claims | Current source stage; new VM model not implemented |
| Alpha | Versioned VM/container schemas and compatible agents; real guest-kernel VM and distinct trusted App Container; authenticated ownership/capacity; idempotent create/cleanup and interrupted-create proof; reproducible isolated install | Blocked on server/foundation/driver/protocol/host evidence |
| Beta | Multi-account real network/disk/S3 denial, claim/renew/fence/uncertain-task recovery, drift, tested workload and control DB restore, operation/audit visibility, measured capacity, real CLI workflows | Planned; destructive testing only on dedicated test data |
| Release Candidate | Frozen release contracts/profile, clean install by another operator, full matrix acceptance, upgrade/rollback rehearsals, signed artifacts/SBOM/provenance, vulnerability review and operational docs | Planned |
| Stable 1.0 | RC evidence reviewed and published, known limitations and support policy approved, no release-blocking security/data-loss defect, licensed distribution/contribution/disclosure channels | Planned; no calendar promise |
| Supported production deployment profile | Exact OS/kernel/backend/storage/network/hardware/workload profile plus security, endurance, restore and upgrade evidence and maintenance ownership | None qualified; Stable alone cannot qualify arbitrary hosts |

## Production evidence checklist

| Requirement | Acceptance that must actually run | Owner |
| --- | --- | --- |
| Clean install | Another operator installs exact verified artifacts on empty supported host without hidden steps | A+B; C evidence |
| VM lifecycle | Create/inspect/start/stop/restart/delete prove guest kernel, identity, readiness and safe retention; resize/archive only declared supported cases | A; C tests; B/D contract |
| Distinct containers | Explicit shared-kernel type; restricted image/privilege/device/socket policy | A+C |
| Identity/isolation | Missing/revoked/wrong-account/role credentials fail; current membership checked | C+B/D |
| Real networks/storage | Cross-account packets/disk attach/backup read/S3 key denied; management route blocked | C+A/B |
| Durable intent/idempotency | Parallel same-key requests one result; mismatched fingerprint rejection; intent/operation/reservations atomic | B; C tests |
| Crash recovery | Interrupt before/after provider acceptance/report; same one resource; unrelated resources unchanged | A+B; C tests |
| Fencing/unknown outcome | Expired/stale owner/epoch/generation rejected; paused child/provider task blocks conflicting takeover | B+A; C tests |
| Secure communication/secrets | Real signing/replay/TLS/custody/rotation/revocation and redaction; fail closed on verifier/replay outage | C+A/B |
| Capacity | Parallel admission cannot overspend; stale host/full disk/RAM produces explicit safe refusal | B+A; C tests |
| Recovery | Known rows/object bytes/VM file hashes restored; corrupt/unauthorized archive rejected; DB and key recovery rehearsed | A+B+C |
| Observability/audit | Fresh/stale state truthful, attempts correlated, denied/admin actions captured, app cannot rewrite audit | B+C/A; E CLI |
| Upgrade/rollback | Old/new agent contract negotiation; schema expand/contract; recovery from a failed upgrade with data consistency | B+D+A/C |
| Release chain | Static/dependency/secret scans, SBOM, signature/provenance verification, vulnerability intake | A+C; maintainers approve |
| CLI/user workflow | Independent user provisions and inspects usable VM/app/dependencies, cancels/waits and cleans up | E+A/B/D; C tests |

All thresholds/profile dimensions, security exceptions and result artifacts need reviewer approval before running qualification; no invented SLA or uptime percentage. Unit-only tests are labelled as such. Benchmarks do not establish resistance to all malicious workloads.

## Academic test continuity

Preserve V5 tests 1–15 and E1–E8 as historical obligations. New VM/container tests supplement rather than silently replace them. Test 1 previously asserted exactly one container; it must be mapped to exactly one VM plus guest readiness for the new compute target, with separate container coverage. Mapping requires team/mentor acceptance and coordinated contracts; see [academic transition](academic-transition.md). Do not invent approved SRS requirement IDs.

## Failure domain and support

Single-host production means a hardened measured installation with backup/recovery; it still loses service when that host fails. HA is separate future work: DB replication/quorum, independent control-plane services, leader/lease safety, host failover, storage/network redundancy and partition tests. Replicating a Java process or using a clustered backend does not pass that gate.

[Implementation evidence](implementation-status.md) records current gaps. Production-oriented means aiming at these requirements, not declaring production-ready by editing documentation.
