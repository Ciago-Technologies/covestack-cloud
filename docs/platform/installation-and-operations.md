# Installation, deployment, operations and troubleshooting

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **Developer source setup only; no supported VM/product installer or production deployment profile exists.** These are requirements for a future verified installer plus factual current boundaries, not instructions that a full platform already runs.

## Current source layout and checks

kan-server contains Java libraries without an application entry point in the inspected tree; Maven repackage remains skipped. kan-agent has the signed client and read-only no-execution loop. The dependency Compose file defines PostgreSQL, Valkey and Mailpit, not object-storage provider, Incus, tenant VMs or the Java application. See [implementation ledger](implementation-status.md).

When separately authorized, source maintainers can run existing Go/Java/contract checks using the repository's manifest toolchain; this documentation task ran none. Do not download a different toolchain to conceal a pin mismatch. Never copy a future CLI/VM command from a design diagram into a claim of current usage.

## Future supported installation sequence

| Step | Installer/operator requirement | Exit evidence |
| --- | --- | --- |
| Preflight | Detect approved OS/kernel/backend versions, virtualization/device access, firmware, nested constraints, free capacity, routes and conflicts | Refuse unsupported configuration; no silent container or TCG fallback |
| Bootstrap identity | Establish protected operator identity, DB/key custody, replay store and initial account/agent assignments | No anonymous first-request admin, shared default credential or secret logs |
| Install artifacts | Verify signed checksummed binaries/images and approved backend packages | Signature/provenance/SBOM verified; exact versions recorded |
| State service | Run forward-only approved migration chain with separate migrator/app/readonly roles | Empty and upgrade database tests; V5/V6 prerequisite conflict resolved |
| Execution backend | Configure Incus/API access only for trusted agent/operator; storage/network ownership and capacity | Guests/tenants cannot reach daemon/socket or management routes |
| Start platform | Run documented Java/agent service entry points under reviewed least privilege, TLS and restart policy | Real authenticated API, replay denial and health/freshness, not process-only success |
| Provision smoke | Create one owned VM and trusted App Container; exercise fixed data stack where in release subset | Actual readiness, isolation, interrupted creation, cleanup and restore |
| Repeat | A second operator reproduces clean install and a second installer run | No unexpected changes or hidden credentials; operation transcript sanitized |

Ansible Part 1 remains the target installer path with reviewed roles, Vault/no_log, artifact verification, lint and idempotent second run. Existing service hardening suggestions such as NoNewPrivileges/ProtectSystem need testing with privileged runtime/storage access; do not claim a unit setting removes daemon authority. No playbook/systemd file is created here.

## Operational contracts

| Operation | Required runbook |
| --- | --- |
| Routine shutdown | Preserve DB/volumes/guests according to reviewed policy; stopping management is not workload deletion |
| Maintenance/drain | Stop new admission, settle claims/tasks, back up state/keys and record capacity; never kill unknown tenant objects |
| Upgrade | Verify compatibility matrix and backups; expand/contract migrations; mixed-version agents cannot receive unsupported work |
| Rollback | State which binary rollback is safe with current schema; restore-forward or isolated backup recovery if incompatible; no automatic SQL downgrade |
| Monitoring | Track API/DB/replay/backend reachability, lease/attempt backlog, capacity, certificate expiry and backup freshness; stale agent is unknown/stale, not healthy |
| Audit | Correlate user/account/resource/operation/request without secret payloads; protect writes and retention |
| Restore | Use isolated target and known data checks; reconcile restored intent with existing backend inventory before enabling mutation |
| Incident | Revoke identity, contain management exposure, preserve redacted timeline, resolve uncertain tasks and verify recovery |

Off-host backups/key recovery protect against a single-host loss; they do not give HA. Retention/RPO/RTO, supported versions, patch cadence, monitoring thresholds and escalation contacts require reviewed operating policy, not invented defaults.

## Troubleshooting without destructive fixes

| Symptom | First evidence to inspect | Unsafe shortcut to avoid |
| --- | --- | --- |
| Agent 401/403 | Credential eligibility, account/host assignment, signing bytes/date/nonce and trusted clock | Default key, disable auth or print secret |
| Work request [] | Assigned account, capability filter and actual DB intent; no execution kind advertised | Infer deletion from empty poll |
| VM pending | Capability/preflight, capacity reservation, network readiness and recorded provider task | Create again after timeout without inspection |
| Backend unavailable | Authorized socket/service status and current provider operation | Give tenant daemon socket or broad host privileges |
| Endpoint not ready | Guest/service readiness and intended packet path | Report convergence from provider task acceptance |
| Restore failed | Backup identity/version/integrity/key and source retention | Overwrite the only valid copy or delete volumes |
| Disk/memory full | Reservations, platform reserve, observed pressure and orphan-review evidence | Blanket prune, release ambiguous reservations or evict arbitrary tenant |

Existing Compose bindings are loopback 5432 PostgreSQL, host 6390 Valkey to 6379, 1025/8025 Mailpit. Defined binds are not live reachability proof. Initialized PostgreSQL credentials require deliberate rotation, not deleting volumes or changing an env value alone. Never use down -v or blanket pruning as a repair recipe.
