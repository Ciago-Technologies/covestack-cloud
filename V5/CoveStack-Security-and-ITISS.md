# CoveStack V5 — Security Architecture and ITISS Evidence

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Security and ITISS contribution

ITISS evidence is actual host/virtualization hardening, identity/replay/custody, private traffic enforcement, capacity admission, audited operations, incident handling and verified data recovery. A Go agent is a trusted privileged component, not a reason to replace infrastructure evidence with code volume. No course/mentor approval is assumed; clarify course-specific ownership expectations separately without constraining product architecture.

## 2. Trust boundaries and threat model

| Threat | Target control | Required real/negative evidence |
| --- | --- | --- |
| User impersonation or removed membership | Authenticate every management API; account/project/action authorization | Missing/expired/revoked credentials, role downgrade and owner race |
| Forged agent or wrong assignment | Stored credential/host/account mapping, signed outbound requests, narrow dispatch | Wrong host/account/key denied; labels/caller selector cannot grant ownership |
| Replay/tamper | Shared KAN1 canonicalization, MAC before nonce reservation, atomic Valkey TTL and safe clock policy | Signature vectors plus real replay/nonce-store outage and no poisoning |
| Host/agent compromise | Dedicated hardened supported host, restricted sockets/devices, patch/update/secret policy | Access matrix/privilege inspection, verified updates and incident recovery |
| VM escape/malicious image | Patched reviewed hypervisor/images/firmware, narrow devices/initialization | Actual guest boundary/provenance and denied host access; no absolute isolation guarantee |
| Container escape | Trusted shared-kernel class, no privileged mode/socket/root mounts/broad devices | Runtime config inspection and real management/host denial |
| Lateral movement/spoofing | Owned NIC/network/default-deny ingress/egress, management separation and guarded metadata | Actual forbidden packets, cross-runtime/restart/spoof cases |
| Storage cross-account access | Owned volume/backup refs and scoped bucket policy/keys | Foreign disk attach/read/restore/object/admin access denied |
| Credential leakage | Token hashes, protected recoverable keys/secret refs, least-privilege delivery/rotation | Redaction/ciphertext, revoke/rotate/restore and no logs/spec/query leaks |
| Stale worker / uncertain side effect | Exact owner/epoch/generation guards and runtime exclusion/task inspection | Paused worker, lost response, parent/child death and no duplicate resource |
| Denial of service | Physical/account reservations, bounded body/queue/timeout/rate policy | Parallel overspend, disk/RAM full, stale-host refusal and existing data retained |
| Backup corruption/lost key | Protected independent copy, digest/authenticity, key recovery and rehearsal | Corrupt/foreign restore rejected; known data and policy actually recover |
| Supply-chain compromise | Verified pinned artifacts/images, static/dependency/secret scans, SBOM/signatures/provenance | Tampered artifact rejected; exact distribution/notices/security evidence |
| Audit forgery/disclosure | Append-only app roles, scoped sanitized event/log retention and external policy | App cannot rewrite audit/events; privileged admin risk remains explicit |

## 3. Identity, signing and secrets

Internal service credentials and public user tokens have separate trust/lifecycle. Current HMAC signer/filter/rules exist; concrete issuance/store/host assignment/revocation integration is pending. Hash opaque API tokens; protect recoverable signing/DB/S3 keys with reviewed key version/custody. No default development key or caller-selected account authority. Signing is authenticity/integrity, not TLS, user RBAC, lease or encrypted data.

Root/admin RustFS credentials stay operator/private Java administration; tenant keys only receive scoped policy. Actual revoke/rotate/restart/restore denial is mandatory. Mature secret/KMS provider selection and bootstrap/recovery are open; do not invent an unsafe homemade vault or pretend current keys already auto-rotate.

## 4. Claims, fences and host execution

Current 1.1.0 lacks the full reviewed advanced claim protocol. Future claim/renew/convergence enforce trusted assignment, exact generation, owner and epoch, with explicitly reviewed validity/expiry semantics. Progress does not acknowledge. DB fence is not a hypervisor fence: serialize host/resource/child actions and settle actual async backend task before takeover. Lease expiry never proves a subprocess/task stopped.

Only trusted agent/operator touches Incus/Docker APIs; tenants get no socket, host filesystem, arbitrary command, broad device or administration credential. Least-privilege service flags reduce exposure but cannot erase the authority of a privileged runtime. Review remote backend TLS/auth separately if needed.

## 5. Network/data policy

Default deny and explicit necessary ingress/egress; private addressing alone is not isolation. Verify VM-to-VM, VM/container/cluster, management/replay/control DB and provider admin routes. Guard NIC/IP spoofing, DNS rebinding/metadata paths and address reuse; no competing IP allocator. At-rest encryption is only claimed for a tested backend/key profile; backups and keys require protected independent recovery and retention/deletion policy.

## 6. Logging, vulnerability and incident work

Record actor/account/project/resource/operation/attempt/provider outcome with sanitized correlation. Unauthenticated events have no trusted account; nullable scope versus separate security events is B+C design. Hash chains alone are not tamper-proof against administrators. Preserve denied/admin/credential events; no secrets in payload/logs.

Security issue triage needs severity, impact/version, owner, fix/retest and residual-risk disposition. Maintained disclosure contact/SLA/governance remain decisions; no invented email or certification. Apache-2.0 license does not certify security or relicense dependencies. [Threat model](../docs/platform/security.md), [recovery](CoveStack-Operations-and-Recovery.md).

## 7. ITISS evidence package

| Evidence | What to demonstrate | Owner / scope |
| --- | --- | --- |
| Host/runtime | Authorized compatible host, hypervisor/image trust, patch/privilege/socket inventory | A; C review; actual profile |
| Network | Owned allocation, default deny and real allowed/forbidden packet traces | A+C/B; sanitized disposable tests |
| Identity | Real signing/replay/revoke/assignment/auth-store outage | C+B+A; not unit mocks alone |
| Capacity | Fresh host budget, quota/concurrency refusal and unknown-task reservations | B+A; C tests |
| Recovery | Independent protected DB/key/VM/SQL/object restore with known content | A+B+C; actual safe target |
| Operations | Clean install, repeatability, upgrade/rollback and incident explanation | A+B+C; no fictional runnable setup |

Keep test revision/setup/results/limitations and explain failures honestly. Do not add a Labs/assignment/instructor-dashboard product or new deadline. A should understand and explain contributed agent Go; E remains CLI learner, D Java only. [Roles](Team_Roles.md), [team checkpoint](../TEAM-FIRST-CHECKPOINT.md).

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-security-and-itiss) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
