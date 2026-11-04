# Tenant networking, storage and backup

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. Status: **Planned**; existing Compose defines platform dependencies only. No network/storage isolation or restore is certified.

## Networking ownership

One account-owned private network remains the initial product/lab constraint; no custom CIDR in the initial subset. VM vNICs use the selected Incus network integration; App Container and k3d adapters need explicit attach/reachability behavior. Reusing a Docker-only helper cannot automatically attach Incus guests. Backend network mapping, subnet pools and allowed cross-runtime paths are B+C+A decisions, not new wire kinds.

| Requirement | Ownership and implementation target | Evidence |
| --- | --- | --- |
| IPAM | PostgreSQL reservations and server allocator or reviewed coordinated backend delegation; one allocator authority | Parallel allocations/account subnet races; no collision with host/VPN/runtime networks |
| Management separation | API/DB/Valkey/Incus/Docker/agent administration denied to tenants | Actual guest/container packets and API authorization tests |
| Default deny | Account-isolated VM-to-VM and cross-account ingress/egress with explicit permitted app/DB/S3 paths | Allow/revoke changes actual traffic, including restart and spoof attempts |
| DNS | Scoped naming, resolver policy and prevention of cross-account disclosure/rebinding/metadata reachability | Name/ownership changes and forbidden destination tests |
| Attachment | Bind account/resource/host/NIC/IP and current network generation | Cross-account attach denied; network not deleted with live/reserved attachments |
| Ports | Reserve management/app ports before use; avoid unapproved management publication | Parallel collision, cleanup/reuse and control-port denial |
| Publication | Explicit opt-in and registered verified route, TLS and hostname ownership | Extended scope only; simulated address is not internet routing |

Do not claim a bridge is a complete VPC/firewall. Incus ACL/network facilities are backend mechanisms to evaluate; [official networking/security docs](https://linuxcontainers.org/incus/docs/main/explanation/security/) do not establish CoveStack policy. Backend IPAM, host firewall and control-plane reservations must not fight each other. No router-level DHCP/DNS rights are assumed for an institute hotspot.

## Storage ownership and secrets

VM disks and application volumes are distinct account-owned resources even when backed by one pool. Bind volume/backend handle to host and account; reject arbitrary host paths/devices, foreign attach and out-of-policy resizing. Reservation includes capacity/snapshot/backup overhead. Expose encryption only for a tested storage/key profile; no unsupported encryption claim.

PostgreSQL tenant instance is separate from the control-plane DB, with private endpoint, persistent storage, scoped credentials and verified read/write/restore. A trusted App Container may host it in the initial reviewed subset; potentially untrusted DB/workload isolation may require a VM-backed profile. No automatic assumption that all tenants may run privileged DB images.

Object storage provider selection remains open; see [provider evaluation](provider-evaluation.md). The selected provider's operator-owned service/storage is separate from tenant bucket/credential records. Java reconciles buckets/policies, not the agent work queue. Protected administrator custody, rotation, version/API/license and private server path remain open. Tenant credentials cannot administer or delete another account's data. Rebuilding bucket rows/policy does not restore lost bytes.

## Backup and verified recovery

| Backup scope | Must include | Restore proof |
| --- | --- | --- |
| Control plane | Consistent PostgreSQL backup, schema version, ownership/reservation/task history and protected key recovery | Restore in isolated environment; validate assignments, leases and no accidental backend duplication |
| VM | Supported guest/disk consistency method, ownership, image/backend compatibility and retention | Boot restored VM under new reviewed identity; known file/application hashes match |
| PostgreSQL workload | Consistent data backup and needed roles/config/secret references | Restore known canonical rows/checksum; backup-file checksum alone insufficient |
| Object data | Selected provider object bytes/version/config as required by tested mode, policy and credential recovery | Exact object byte checksums and cross-account denial after restoration |
| Agent/operator | Protected identity/bootstrap references, approved artifact versions and recovery procedure | Reissue/revoke safely; no secrets in ordinary metadata backup |

Snapshot is not automatically an off-host backup or application-consistent capture. Record quiesce policy, checksum/authenticity, provenance, size, source generation, verification time and retention; exact durable schema still needs B review. Protect off-host copies and keys against same-host loss. Define RPO/RTO goals after measurements, never invent recovery guarantees. Test corrupt, unauthorized and incompatible restores; interrupted restore must preserve recovery options.

Deletion is durable intent, followed by confirmed teardown and explicit retain/purge policy. Release IP/port/storage capacity only after known cleanup; retain tombstones so public IDs are never reassigned accidentally. GC requires complete scoped inventory, a reviewed grace period and operator evidence; partial known-ids/timeout is not deletion authority. Secure erasure depends on storage/encryption backend and must not be promised from file unlink alone.
