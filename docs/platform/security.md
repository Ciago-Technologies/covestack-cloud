# Security target and threat model

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Updated: 4 November 2026. **Requirements, not a production certification.** Review owner C with B+A/D. Current source evidence is in [implementation status](implementation-status.md); no security runtime exercise ran in this documentation task.

## Assets, actors and trust boundaries

Assets: account resources/data, VM disks/images/NICs, App Container volumes, DB and S3 credentials, kubeconfigs, signing/decryption material, backup copies, capacity reservations, audit history and release artifacts. Actors: anonymous caller, authenticated tenant, operator, agent, compromised guest/container, malicious image/source/registry and a failed/stale worker.

Boundaries: internet/client → API; user → account membership; API → DB/key store; assigned agent → signed API; privileged agent → Incus/Docker/k3d/storage/network; VM/container → host/other accounts; tenant → selected object-storage data plane; backup/release producer → restore/install consumer. The trusted host and privileged agent can defeat tenant controls if compromised; mitigate and document rather than promise resistance to a hostile root operator.

## Threat/control/evidence register

| Threat | Required control | Evidence before claiming readiness |
| --- | --- | --- |
| Spoofed user or membership change | Authenticate every management API; explicit RBAC/action/account checks before lookup/mutation; bounded sessions and revocation | Unknown/expired/revoked token, role downgrade, foreign ID and membership race tests |
| Spoofed agent/cross-account work | Stored credential-to-host/account assignment; signed request; immutable per-account client context | Wrong assignment/agent/key and revoked context denied; no default key or caller-selected scope |
| Replay/tampering | Preserve current canonical request rules, atomic authenticated nonce claim and bounded clock policy; fail closed on replay-store outage | Signature vectors plus real valid/replay/tamper/outage tests |
| Credential leakage | Hash opaque user tokens; encrypt signing/DB/S3 secrets where recovery required; separate versioned master-key custody; authorized secret delivery | Redaction tests, protected backup/key rotation/revocation/restore; no plaintext in ordinary spec/observed/events |
| Agent/host compromise | Restricted service identity, socket/device ACLs, patched dedicated host, constrained adapter inputs; protect updates and recovery keys | Host hardening inspection, access matrix, update verification and incident drill |
| VM escape/malicious guest | Supported patched hypervisor, reviewed image provenance/firmware, no unnecessary devices or passthrough, default-deny vNIC | Guest-kernel proof, escape exposure review, patch/image inventory and isolation tests; no absolute isolation claim |
| Container escape | Trusted workloads only initially; no daemon socket/privileged mode/root mounts; reviewed user/capability/seccomp/mount policy | Container config inspection and denied host/socket/device access |
| Lateral movement | Account networks/default-deny ingress and egress; separate management route; no direct DB/cache/Incus/Docker exposure | VM-to-VM, VM-container, cross-account and management reachability tests including restart |
| DNS/IPAM collision/spoofing | Locked durable allocations, pool containment, reserved addresses, no competing allocator; review resolver/DHCP/metadata controls | Parallel allocation, overlapping route, stale lease and spoofing/metadata-exposure tests |
| Storage cross-account access | Ownership-constrained disk/volume/backup references; scoped S3 policy; no arbitrary host paths | Foreign disk attach/restore/key reads fail; same friendly names safe |
| Backup corruption/lost key | Integrity/authenticity checks, protected off-host copies, separate key recovery and restore drill | Corrupt/unauthorized backup rejected; known rows/files restored and verified |
| Stale worker/uncertain side effect | Exact generation/owner/epoch/expiry validation; per-host exclusion; settle provider task before retry | Pause worker, kill parent/child, drop provider response, reject stale report and prove no duplicate VM |
| Resource exhaustion | Durable quota and physical-capacity reservations, rate/body limits, bounded queues/timeouts | Concurrent create, full disk/RAM and stale capacity gracefully rejected without partial resources |
| Supply-chain compromise | Pinned verified releases/images, dependency/secret/static scans, SBOM, artifact signatures/provenance | Reproducible release evidence, signature rejection, vulnerability triage and published reporting channel |
| Communication failure | Verified TLS off loopback; no redirects; bounded signed attempts, no fallback identity; deny new mutations when authority uncertain | Server/replay/backend outage and credential loss recover without unsafe cleanup |
| Audit forgery/disclosure | Append-only app privileges, scoped read, redaction, correlation and protected external retention | App cannot alter events; denied access/reassignment logged; administrator authority disclosed |

## Identity and access design

Account/membership and explicit least-privilege role/action matrix remain B+C review; no new wire role values are introduced. Tenants cannot grant operator rights or choose an agent by an unvalidated identifier. User tokens and service credentials have separate issuance, expiry, rotation and revocation semantics. Evaluate short-lived service credentials with secure renewal/recovery; do not pretend current HMAC keys already have automated renewal. Bootstrap must require operator-controlled provisioning, not first-request admin creation or a shared default secret.

Optional policy-engine extraction is a later ADR. [OPA](https://www.openpolicyagent.org/docs) is a policy component to evaluate, not mandatory or installed. Keep authoritative identity and policy versioning consistent; an evaluator outage must not become allow-all. Secret/KMS provider selection remains open.

## Host/runtime controls

Use a supported hardened Linux host with reviewed kernel, virtualization, firewall, disk protection and time synchronization. Patching, maintenance windows, operator access, secure boot/image support and incident responsibility belong in the supported profile. Least privilege reduces attack surface but cannot neutralize the authority of a required privileged runtime API.

[Incus security documentation](https://linuxcontainers.org/incus/docs/main/explanation/security/) describes full control through the local socket, including host attachments. Do not expose that interface as the tenant API. Remote backend administration, if ever required, needs authenticated TLS and explicit authorization. A dedicated per-account backend project may help ownership but does not replace CoveStack authorization or backend hardening.

## Data protection and audit

Encrypt transport to management and data endpoints in the production profile; development loopback exceptions are labelled. At-rest encryption is claimed only for a tested backend/disk/key configuration, not a checkbox. Include backup keys and deletion/retention policy. Never describe DB/S3 credential scoping as VM-grade isolation. A selected object provider may use one listener for administration/data requests; keep administrator keys and console routes operator-only.

Audit who requested what, acting account, resource/operation/revision, decision, provider task and result with sanitized correlation. Unauthenticated security events cannot invent a trusted account; nullable scope versus separate security events remains a DB/access decision. Hash chains alone do not protect against a privileged administrator rewriting history.

## Security release gate and disclosure

Record findings, severity, affected version, exploitability, remediation or approved residual risk with owner. No Stable profile with unresolved critical/high exposure unless a documented security decision explicitly rules the finding inapplicable; establish a severity/exception policy during review, not assumed here. Provide vulnerability-reporting contacts only after a maintained monitored channel is designated; do not invent security@ addresses or an SLA. See [contribution/attribution](license-and-attribution.md) and [OpenSSF criteria](https://www.bestpractices.dev/en). No badge or compliance claim is earned by linking the criteria.
