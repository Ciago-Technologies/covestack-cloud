# CoveStack V5 — Operations, Failure Handling and Data Recovery

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Operational principle

**Preserve owned resources and persistent data during uncertainty.** A network outage, incomplete list or malformed work response is not deletion intent. Restart recovery inspects durable IDs and reserved labels before making changes.

| Event | Immediate behavior |
|---|---|
| Server unreachable | Pause new work; retain workloads |
| Agent stops | Mark stale; avoid speculative replacement storms |
| Replay store unavailable | Reject affected authenticated requests safely |
| App fails | Report health/error; honor desired state |
| Tunnel fails | Public endpoint unavailable; internal resources remain |
| Proxy config invalid | Keep last valid configuration; report failure |

## 2. Routine operations

| Procedure | Required checks |
|---|---|
| Startup | Secrets, versions, private listeners, schema and dependency health |
| Deployment/update | Reviewed image/config changes; health and compatibility |
| Shutdown | Drain/stop supported work; preserve volumes |
| Key rotation | Scoped replacement, revocation and tested rollback |
| Certificate renewal | Valid certificate plus tested proxy reload |
| Resource deletion | Owned dependency cleanup and explicit retention |

Do not use `docker compose down -v` or blanket prune as a default repair. A workload's desired off state is different from an active workload crashing. Recovery must honor that distinction.

## 3. Backup coverage

| Data | Backup/recovery need |
|---|---|
| Control-plane PostgreSQL | Accounts, resource/action/stack records and ownership |
| Workload PostgreSQL | Database-consistent data backup |
| Object storage | Objects plus policies/ownership references |
| Master secrets | Protected key recovery distinct from DB backup |
| Runtime config | Proxy/domain/network/version records |
| Kubernetes | Defined workload/data recovery; kubeconfig alone is not backup |

A container filesystem snapshot excludes mounted-volume data and is not a database backup or VM memory snapshot. Record precisely what each backup/snapshot contains.

## 4. Restore exercise

| Step | Evidence |
|---|---|
| Seed known data | Expected record/object/checksum |
| Take documented backup | Time, source revision and protected artifact |
| Prepare separate environment | Avoid destroying the current demo |
| Restore dependencies/keys | Correct configuration and ownership |
| Restore data/state | Verified DB/object access |
| Reconcile workload identity | No duplicates or cross-account rebinding |
| Verify app operation | Defined usable result |
| Measure | RPO/RTO and unresolved gaps |

RTO is time to the defined usable service. RPO describes data/time lost relative to the recovery event. A same-host restore does not prove host-loss recovery if all backups are on that host. State the scenario honestly.

## 5. Idempotency and fencing

| Scenario | Rule |
|---|---|
| Request response lost | Refresh durable operation before resubmitting |
| Agent retry | Fresh signed request; mutation semantics agreed |
| Worker duplicate | Per-resource serialization and durable identity |
| Old agent resumes | Lease/epoch check prevents stale authority |
| Partial stack failure | Resume safe steps; do not erase persistent data |
| IP/hostname reuse | Wait until stale route/metadata attachment removed |

## 6. Incident runbook outline

Record symptom, time/commit/environment, affected accounts, observed state and safe containment. Restrict/revoke relevant credentials, preserve sanitized evidence, correct the cause, verify recovery, and write follow-up fixes. A reboot without diagnosis is not an incident report.

## 7. Release evidence

Clean VM install, remote browser app, approved private admin connectivity, restart test, credential revocation, alert trigger and separate restoration all need actual results. Store exact versions/hardware and identify cases not tested. Use the [acceptance matrix](CoveStack-Testing-and-Acceptance.md).
