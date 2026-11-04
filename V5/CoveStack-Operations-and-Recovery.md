# CoveStack V5 — Operations, Incident Handling and Verified Recovery

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Operational authority

DB intent is authority; provider objects/tasks and observed health are evidence to reconcile, not instructions to adopt/delete unrelated infrastructure. Service shutdown is not resource deletion. Operator must understand trust/capacity/retention and the actual supported subset/profile before executing runbooks. No full platform or production profile is established today.

## 2. Routine operations and failure handling

| Operation/failure | Required behavior | Evidence to retain |
| --- | --- | --- |
| Control server down | No newly unauthorized mutation/claim; do not turn discovery failure into empty inventory | Actual outage/recovery and stale-status behavior |
| Replay store down | Internal auth fails closed; no in-memory fallback | Real Valkey negative/outage tests, not mocked API calls alone |
| Agent restart | Resolve current assignment/tasks and inspect owned objects before resume | No duplicate VM/cluster, retained data and stale-worker rejection |
| Backend/task timeout | Keep intent/reservations, inspect accepted task/outcome, block conflict | Parent/child death/lost reply and actual safe takeover |
| Disk/RAM pressure | Stop unsafe admission/updates; preserve existing owned workloads | Useful capacity/refusal and no reservation overspend |
| Drift | Fresh authorized inspection repairs current intent or reports policy-limited failure | Same desired revision where spec unchanged; no false health/ack |
| Maintenance/drain | Stop new admission, settle tasks, protect keys/data and record current assignments | Recoverable restart/re-entry; no unrelated cleanup |
| Delete/retention | Authorized intent, reverse owned dependencies, preserve shared/retained objects | Actual teardown before live IP/port/capacity release |

DB epoch/generation rejects stale reports but does not cancel a running backend action automatically. No bulk name-prefix scan or partial known-IDs result authorizes GC; complete scoped inventory, exact ownership and reviewed grace are required.

## 3. Backup coverage

| Scope | Required protected recovery material | Actual restore assertion |
| --- | --- | --- |
| Control-plane DB | Consistent schema/state, identity/assignments, operations/tasks/reservations, policy and audit references | Restore isolated DB, verify owners/claims and reconcile provider inventory before mutation |
| Master/credential custody | Reviewed protected key versions, revocation and bootstrap recovery | Authorized recovery/rotation; old revoked identity stays denied |
| VM/volume | Backend-consistent guest disks/config/identity and required key/image provenance | Boot safe restored owned target, compare known file/application hash |
| PostgreSQL workload | Consistent DB/role/config data and secret references | Restore expected rows/checksum, private access and foreign denial |
| RustFS object service | Actual bytes/required versions/metadata plus needed policy/credential/key/config recovery | Isolated target, known bytes checksum, scoped key/deny and persistence after restart |
| Templates/environments | Definition/input/secret refs, child/sharing/retention bindings | No duplicate child or deletion of shared infrastructure after resume |

A snapshot may not be an independent/app-consistent backup. A successful file write is not restore proof. Copying an unknown active object-store directory or assuming cross-provider encrypted formats are portable is not an approved backup method. Select consistent RustFS service recovery with B+C/A; a tool using RustFS as a backup destination is not evidence of backing up RustFS itself.

## 4. Restore exercise

1. Identify exact owner/source/recovery artifact/version/consistency/key requirements; authorize safe target and capacity.
2. Verify integrity/authenticity/access before touching data; reject corrupt/foreign/incompatible source.
3. Restore into isolated network/identity target with protected keys; no overwriting unrelated production resources.
4. Reconcile DB intent against actual provider tasks/objects before enabling new mutations.
5. Verify guest/SQL/object known content, correct owned policy and foreign-access denial; record result/time/reviewer and failures.
6. Recheck retained reservations, DNS/routes/credentials if applicable and safe cleanup of only disposable test data.

Recovery objectives/retention/timeouts require measurements and policy; no invented RPO/RTO or guarantee. Keep independent protected/off-host copies where host-loss recovery is required; same-host backup alone is not HA or disaster recovery.

## 5. Upgrade, rollback and installation

Use exact signed/checksummed artifacts and compatible schema/agent/backend matrix. Never edit applied migrations or assume binary rollback works with an incompatible schema. Expand/contract plus restore-forward/isolated DB recovery need rehearsals. Driver/API changes preserve current work ownership; unknown capabilities/types fail before execution. Second operator clean/repeat install must be reproducible. Current Java foundation/VM/RustFS setup remains pending.

## 6. Incident runbook

| Step | Required action |
| --- | --- |
| Identify | Record sanitized revision/profile, affected account/resource/operation/host and actual freshness |
| Contain | Revoke exposed credential, isolate approved management/runtime path, stop unsafe admission; preserve evidence |
| Diagnose | Inspect authoritative assignment/claim/provider task without treating missing telemetry as deleted resource |
| Recover | Use reviewed safe retry/restore; no fallback credentials, forced volume deletion or destructive broad cleanup |
| Verify | Real service/data/traffic and denial checks, stale-worker safety and protected key handling |
| Learn | Record root cause/fix/evidence and remaining limits; update issue/runbook, not past test history |

Privileged host/agent compromise can defeat local trust boundaries; no tamper-proof audit claim. Security contact/response policy must be maintained rather than invented. Never use `down -v` or reset the DB to conceal an unresolved incident.

## 7. Evidence and owners

A host/driver/installer/CI, B durable state/recovery, C security and all acceptance, D API semantics, E client diagnostics. Tests 1–15/E1–E8 and named VM/image/packet/capacity/key/DB/S3 recovery cases follow [acceptance](CoveStack-AntiWrapper-Acceptance.md). No old calendar/service-count constraint defines support; [gates](CoveStack-Service-Gates.md) and profile evidence do.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-operations-and-recovery) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
