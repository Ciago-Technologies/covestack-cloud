# CoveStack V5 — Operations, Failure Handling and Data Recovery

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Operational principle

**Preserve owned resources and persistent data during uncertainty.** A network outage, incomplete list or malformed work response is not deletion intent. Restart recovery inspects durable IDs and approved proposed-label contract before making changes. kan.* is not existing compatibility.

| Event | Immediate behavior |
|---|---|
| Server unreachable | Pause new work; retain workloads |
| Agent stops | Create stays pending; bounded waits; no claimed convergence |
| Replay store unavailable | Reject affected authenticated requests safely |
| App fails | Report health/error; honor desired state |
| Extended operation failures | Prior route/config requirements relocated to roadmap; not CORE-5 release gates |

## 2. Routine operations

| Procedure | Required checks |
|---|---|
| Startup | Secrets, versions, private listeners, schema and dependency health |
| Deployment/update | Reviewed image/config changes; health and compatibility |
| Shutdown | Drain/stop reviewed work; preserve volumes |
| Key rotation | Scoped replacement, revocation and tested rollback |
| Extended edge procedures | Prior renewal/reload policy preserved in roadmap |
| Resource deletion | Owned dependency cleanup and explicit retention |

Do not use `docker compose down -v` or blanket prune as a default repair. A workload's desired off state is different from an active workload crashing. Recovery must honor that distinction.

## 3. Backup coverage

| Data | Backup/recovery need |
|---|---|
| Control-plane PostgreSQL | Accounts, resource/action/stack records and ownership |
| Workload PostgreSQL | Database-consistent data backup |
| Object storage | Objects plus policies/ownership references |
| Master secrets | Protected key recovery distinct from DB backup |
| Runtime config | Account/private-network/cluster/volume records and tested version decisions |
| Kubernetes | Defined workload/data recovery; kubeconfig alone is not backup |

Filesystem image-capture limits are preserved in the roadmap. CORE-5 PostgreSQL backup must be database-consistent, with verified known rows/checksum. Control-plane DB backup and protected key recovery are separate necessities; record exactly what artifacts contain.

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
| Old agent resumes | Exact owner/epoch/generation ack predicate; old epoch rejected |
| Accepted convergence | Clear lease_owner and lease_expires_at together; keep epoch unchanged; progress never changes acked_generation |
| Partial stack failure | Resume safe steps; do not erase persistent data |
| IP/port reuse | Release only after owned attachment/cluster cleanup; live partial uniqueness remains |

## 6. Incident runbook outline

Record symptom, time/commit/environment, affected accounts, observed state and safe containment. Restrict/revoke relevant credentials, preserve sanitized evidence, correct the cause, verify recovery, and write follow-up fixes. A reboot without diagnosis is not an incident report.

## 7. CORE-5 release evidence

Clean-VM install, tests 1–15, private administration, credential revocation and measured PostgreSQL restoration need actual results. Live demo highlights tests 1, 3, 4 and get view; remaining tests are CI/video. Keep backup recording for the phone hotspot; CLI cross-builds three targets and runs on two actual OSes. Public-edge/alert scenarios are preserved as extended targets in the roadmap.

Only two drivers: generic container and k3d. Java reconciles bucket/policy after MinIO restart; failed/partial work inventory never authorizes deletion. Multi-account agent polls remain B+C+A decision. Retention defaults are TBD with B+A.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
