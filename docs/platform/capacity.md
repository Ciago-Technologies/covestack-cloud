# Scheduling, capacity and benchmark plan

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. Status: **Planned**. No supported hardware/VM density is established. Owner B for current durable admission, A host discovery/runtime, C concurrency tests; future Go scheduler needs a reviewed boundary.

## Deployment profiles

| Profile | Intended use | Admission and evidence |
| --- | --- | --- |
| Developer laptop | Source/unit work and isolated experiments; host backend may be absent | No assumption VM/KVM exists; never claim hardware/runtime support from a binary version |
| Single-host lab | One authorized Linux virtualization host | Previous 8 GB AlmaLinux figure is a development/lab constraint only; measure admitted workload combinations |
| Supported self-hosted production | Hardened single host with approved OS/kernel/Incus/QEMU/KVM/storage/network configuration | None approved; requires measured reservations, endurance, restore, security and upgrade gates |
| Future multi-host production | Placement, independent host failures and recoverable distributed state | Deferred; requires host identity, topology, network/storage/claim coordination and split-brain design |

The historical 2-cluster cap and 1-server/max-2-agent shape describe an old academic profile only; they are not product quotas or supported density. Set tested per-profile limits using measured capacity; reject unsafe requests. No VM count, minimum RAM or vCPU-to-core ratio is invented.

## Discovery and budget

Agent reports authenticated host fingerprint, backend capability/version pair, CPU topology, virtualization support/access, memory/pressure, disk free/allocated/IO headroom, network pools/ports and observation time. These are proposed report responsibilities, not current wire fields. Control plane validates assignment and freshness; self-reported capacity is not an excuse to trust unassigned work.

For each allocatable dimension: **admissible capacity = measured safe host budget − platform reserve − committed allocations − outstanding reservations − reviewed safety margin**. Reject when requested incremental capacity exceeds any dimension. Units, freshness thresholds and margins must be benchmarked and documented per profile.

| Dimension | Admission rule |
| --- | --- |
| Memory | Reserve guest allocation plus measured hypervisor/device overhead, container/Kubernetes overhead and platform peak; no implicit memory overcommit initially |
| CPU | Reserve against measured safe budget; initial conservative non-overcommit target, later explicit workload-tested policy |
| Disk | Reserve VM/volume capacity plus images, snapshots/backups and working space; thin provisioning does not make capacity free |
| IO | Measure latency/throughput under mixed DB/VM/backup load; refuse or queue if profile guarantees would be violated |
| Network | Unique owned subnet/IP/port reservations and allowed host-route constraints; avoid Docker/Incus/host/VPN overlap |
| Kubernetes | Budget nodes and control-plane/network overhead rather than multiplying an unverified idle-memory estimate |
| Quotas | Account limits plus canonical physical-host capacity across all account contexts; user account aliases cannot evade admission |

## Scheduling and reservation lifecycle

Authenticate → validate capabilities/trust class → lock canonical host/account/resource/allocation rows in reviewed order → reserve all dimensions atomically with intent/idempotency → return operation → assign ready dependencies → claim and execute. No backend command in the DB transaction. Two simultaneous requests cannot both spend the same headroom. Resource creation requires dependency readiness, not merely dependency rows.

Observed usage informs pressure; desired reservations remain authority for admission. Backfill actual inventory after restart without duplicating accounting. Agent reassignment, resize and restart use re-admission/delta reservations. Keep reservations during ambiguous provider tasks and deletion; release only after confirmed owned cleanup, never lease expiry alone. Downscale/delete behavior does not silently discard retained volumes.

Use bounded queues and fair per-account scheduling. Define reviewed priority classes and starvation prevention before exposing knobs; no tenant-controlled operator priority. Under pressure stop admitting new work, report useful refusal/queued status and preserve existing owned workloads; no arbitrary tenant eviction or fabricated error code. B+D/C decide error/retry contract. Stale host evidence blocks new admission; operator maintenance/draining has explicit policy.

## Benchmark and host-profile qualification

| Phase | Measurement / test | Evidence retained |
| --- | --- | --- |
| Baseline | Hardware, firmware, nested-virtualization status, OS/kernel/backend pair, storage/network setup | Reproducible inventory and image digests; no secret values |
| Platform reserve | API/DB/Valkey/agent and selected object-storage provider idle/peak, backup/upgrade activity | CPU/RAM/disk/IO time series and proposed reserve |
| VM lifecycle | Cold/warm image create, boot/guest readiness, start/stop/restart/resize/delete | Latency, overhead and resource enforcement; no density inferred from one boot |
| Mixed workload | Guest app + SQL + S3 + optional Kubernetes/containers under increasing load | Response/error/pressure/disk behavior, maximum safe measured combination |
| Concurrency | Parallel create/retry/cancel/deletion and stale reports | No overspend, duplicated reservation or unrelated cleanup |
| Failure | Agent/provider/DB outage, restart, storage full, lost response | Existing data preserved, conservative admission and recoverable ownership |
| Endurance/maintenance | Sustained representative load, backup/restore, upgrade and rollback | Stated duration, actual results, limits, risks and support-profile recommendation |

No benchmarks ran here. A reviewed report must identify profile hardware, versions, workload, duration and pass thresholds before declaring support. Store negative results too. A production profile remains Blocked until [readiness gates](readiness.md) pass.

## Proposed multi-host placement boundary

Multi-host scheduling is Deferred (catalogue item 49); single-host admission is the prerequisite, not an implementation of fleet scheduling. Registration binds the trusted agent to a canonical operator-approved host identity and its authorized accounts. Heartbeats authenticate capability/health observations; stale or draining hosts are not new-placement candidates. Define enrollment, revocation and reassignment in a coordinated contract change before accepting reports.

| Step | Policy responsibility | Failure rule |
| --- | --- | --- |
| Filter | Required backend/image/architecture, network/storage locality, trust class, healthy host and requested placement constraints | No eligible host produces a clear refusal or bounded queued operation, never silent fallback to containers |
| Budget | Account quota and physical CPU/RAM/disk/network reservations across all account contexts | Lock/reserve atomically; concurrent admissions cannot overspend |
| Rank | Reviewed deterministic priority and fit policy with account fairness; record explanation | No claim of AI optimization; operators review scoring and starvation policy before exposure |
| Assign | Persist host assignment/reservation and operation before dispatch; agent validates assignment | Lost replies reconcile the same owned resource rather than creating elsewhere |
| Recover | Rebuild reservations from durable allocations plus scoped provider inventory | Unknown old-host execution blocks relocation until fenced/verified; lease expiry alone cannot power off or migrate a VM |
| Rebalance | Future explicit supported migration only, respecting volumes, networks and downtime | Never promise live migration, automatic HA or eviction of unrelated workloads |

Incus already has cluster placement and resource management. CoveStack owns account/environment dependency, quota, trust and allocation policy; the adapter translates a reviewed placement into backend constraints and reads back actual placement. Do not run two competing allocators or assume Incus clustering provides CoveStack database/agent HA. Test a failed reservation, simultaneous placement, stale host, backend relocation and partial network/storage availability before supporting multiple hosts. Sources: [Incus clustering](https://linuxcontainers.org/incus/docs/main/clustering/) and [projects](https://linuxcontainers.org/incus/docs/main/projects/).
