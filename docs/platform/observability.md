# Monitoring, health, logging, events and quotas

Updated: 4 November 2026. **Core observability foundations Planned**; advanced alerting/dashboards/tracing remain future capabilities. [Catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md) owns classification; these capabilities need not each be a new service process.

| Capability | Required observations / behavior | Security and failure acceptance |
| --- | --- | --- |
| Host metrics | CPU/RAM/pressure, disk capacity/IO, runtime availability and collection freshness | Verified host identity; stale capacity blocks admission; a binary version is not healthy runtime |
| Workload metrics | VM, container and Kubernetes utilization, actual readiness and recovery status | Account/project-scoped reads; guest probes cannot expose host secrets or grant host administration |
| Agent/scheduler health | Last validated observation, assignment, queue/reservation/attempt state and current authority | Stop agent/backend: show stale/unknown, not cached healthy; no observation increments desired revision |
| Resource health | Desired versus observed configuration, generation correlation, sanitized errors and provider readiness | Progress/health never advances acked_generation or claims success from process startup |
| Basic logging | Platform, agent, provisioning and authorized workload logs with operation/request correlation | Redact tokens/env/kubeconfig; bounded storage/body/retention; foreign logs denied |
| Audit/events | API/admin/agent/resource/security/credential actions and explicit decision/result | Append-only app privileges, failed-auth scope handled, privileged-operator threat documented |
| Quotas | Account CPU/RAM/disk/VM/container/cluster/storage/network limits plus physical reservations | Concurrent requests cannot overspend; pending/deleting/uncertain work retained; usage not duplicated as conflicting authority |

Metric collection and backend exporters must be evaluated for supported versions/privileges/footprint; no telemetry stack or numeric thresholds selected. Alerts/notifications for capacity/backup/agent/security and historical predictions/tracing are future extensions, not a new core dashboard or SaaS dependency. Basic diagnostics remain necessary before those extensions.

Observed timestamps use reviewed source/clock/sequence semantics. Report collection time separately from last successful health check and desired/converged revision. An empty work list is not a host-health signal. Network failure is not a reason to delete an observed resource or guess readiness.

Define retention/backpressure and disk-exhaustion behavior; critical ownership/idempotency/security state cannot disappear because a logging system is unavailable. Audit failure policy versus fail-closed sensitive administration needs explicit C+B review; no silent dropped security events. All log/export access obeys current membership/resource policy, not a UI filter alone.

Acceptance: multi-account metric/log denial, secret redaction, stale/offline reporting, bounded log storage and actual host/workload observations, event-operation correlation, quota races and operating diagnosis of a failed create/restore. C owns tests; A provides agent collection, B state/coordination, D API, E CLI. No collector, database, alert manager or runtime check was installed/executed here.
