# VM Instance and App Container resource model

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **VM compute decision selected by user; backend evaluation and implementation Planned.** No wire schema was changed.

## Terminology and isolation

The primary general-purpose compute resource is **VM Instance** (also Compute Instance): an actual machine with guest OS/kernel, virtual CPU/memory, disk and NIC. Incus is the preferred initial management backend to evaluate; QEMU/KVM supplies the proposed Linux virtualization path. Do not present a Docker container as a VM or select a container when virtualization is unavailable.

| Property | VM Instance | App Container |
| --- | --- | --- |
| Kernel | Separate guest kernel | Host kernel |
| Isolation | Hardware-assisted virtualization target | Namespaces/cgroups and runtime restrictions |
| Startup/resource use | Higher; measure on chosen host | Usually lower; still reserve and measure |
| OS | Validated guest images/architectures; no blanket OS guarantee | Linux-compatible workloads on selected Linux runtime |
| User | Account-owned general-purpose compute | Trusted application/development workloads and internal services |
| Backend | Incus/QEMU/KVM evaluation target | Docker or later reviewed alternative |
| Security classification | Stronger boundary, still exposed to hypervisor/agent/host compromise | Shared-kernel boundary; not arbitrary hostile-tenant compute |

[Incus distinguishes system containers from VMs](https://linuxcontainers.org/incus/docs/main/explanation/containers_and_vms/). [KVM](https://www.linux-kvm.org/page/Main_Page) and [QEMU](https://www.qemu.org/docs/master/system/introduction.html) are execution foundations, not tenancy/HA guarantees. A guest may require an agent for file/exec/bootstrap features; supported guest images and readiness probes must be tested rather than inferred from a hypervisor boot.

## VM lifecycle target

Validate VM request → authenticate/authorize account → check virtualization/capability → capacity/quota admission → persist intent/operation/reservations → assign trusted agent → claim current authorized attempt → inspect ownership → provision through selected API → configure disk/CPU/RAM/vNIC → default-deny network → bounded secure initial access → verify guest/application readiness → report current convergence → inspect drift/recover.

| Action | Expected target semantics | Failure/acceptance requirement |
| --- | --- | --- |
| Create | Reviewed image, reserved resources, deterministic owned instance; no host mount/device supplied by tenant | Interrupt after provider accepts create; reconcile recorded backend task and prove exactly one owned VM |
| Inspect | Redacted inventory/guest state/capacity and observed freshness | Booted is not ready; foreign backend ID cannot be adopted |
| Start | Only owned retained VM with capacity re-admission and network controls | Failure preserves disks and reports actual status |
| Stop | Prefer bounded graceful guest shutdown; force only reviewed explicit policy | Power-off observation required; retain disk/IP per policy |
| Restart | Authorized stop/start or backend restart with operation tracking | No second VM; explain downtime and readiness |
| Resize | Only validated backend/image operations; initial scope may require powered-off guest | No default disk shrink; reserve delta first; unsupported request rejected; old data retained on failure |
| Network/storage attach or detach | Authorize both resource owners, reserve address/disk, check backend compatibility and dependency safety | Actual interface/route/volume state verified; interrupted attachment resumes without cross-account adoption |
| Snapshot | Record owned recovery point and backend consistency limits; application quiescing where required | Inspect snapshot completion and restore known data; snapshot alone is not an independent backup |
| Backup / restore | Scoped encrypted artifact, checksum, retention and target ownership | Restore known file/data values to a safe target; corrupted or foreign artifact rejected |
| Recover / reconcile | Compare durable generation, owned backend task and actual state after interruption | No duplicate VM or unrelated cleanup; unknown outcomes retain reservations |
| Archive | Product retention/export intent to be specified with B+D; never equate it automatically with VM delete or snapshot | Existing archive wire semantics are container-specific; future mapping/retention approval required |
| Delete | Versioned intent, revoke ingress/credentials, remove only owned backend resources under retention policy | Release reservations only after confirmed cleanup; retain immutable identity tombstone |

## Contract and schema compatibility

Current contract 1.1.0 compute.instance examples describe the prior container execution model. Keep all existing kind/state/header names and historical fixtures unchanged. **VM Instance and App Container are product terms, not approved new kind strings.** Current clients must not interpret an old image/spec as a guest VM request.

A future versioned ADR must settle: distinct kind versus explicit backend/spec revision, capability negotiation, identity retention, old desired-state mapping (including archive), VM images/disks/NICs, action availability and migration of existing container records to the separate App Container model. No implicit conversion; running containers remain containers until an explicit validated migration. Existing numeric /v2 compatibility is a bounded adapter, not a right to reuse incompatible specs. Old agents reject or are withheld incompatible work; mixed-version upgrade/rollback must be tested.

B owns schema/migrations; D owns wire/DTO compatibility; A owns driver/client; C reviews privilege/identity/tests. Coordinated source schemas, fixtures and generated outputs precede execution. V5 reference DDL must not be copied into a VM migration without review; this task changes none.

## Backend evaluation exit

On a separate authorized host: verify supported host/kernel/backend pair, hardware/nested virtualization and device permissions; boot a signed/verified guest image; prove its own kernel; exercise network/disk ownership, readiness, resource enforcement, reboot, cleanup, crash recovery and restore. Record actual CPU/RAM/disk/latency and privileged access. No exact versions, VM density or minimum RAM is invented here.

Incus local API socket grants highly privileged instance/host attachment control. Only trusted agent/operator may reach it; never tenants, guests or App Containers. No host root filesystem, unrestricted devices, arbitrary host exec or daemon socket in workload specs. See [threat model](security.md) and [capacity plan](capacity.md).
