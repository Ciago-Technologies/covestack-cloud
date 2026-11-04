# CoveStack V5 — Instance Metadata and Workload Identity

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Two different metadata meanings

| Metadata | Scope and authority |
| --- | --- |
| Management resource description | Owned API/CLI/UI view: resource IDs, desired/observed generation, readiness/freshness and dependencies; core resource inspection |
| Guest self-metadata endpoint | Future controlled workload-accessible service for its own identity/network; not current IMDS support |

Do not confuse an API details screen with a guest route. Current source establishes neither a VM metadata driver nor a link-local endpoint. VM compute remains core; richer workload metadata delivery is a separate future capability.

## 2. Target bounded payload

Read-only self identity, approved network/address/interface information and declared image/placement fields where needed. Never expose signing/admin/master keys, other accounts, database passwords, kubeconfig, root S3 credentials or arbitrary control-plane views. Exact schema/version/path/values require B+D+A/C agreement; no AWS compatibility or wire string is invented.

## 3. Delivery options and routing

Evaluate guest agent, approved network proxy or a routed link-local service per runtime/profile. Merely binding an HTTP server to 169.254.169.254 does not install guest routes, identify callers or prevent spoofing. A VM can have its own NIC/kernel; App Container/k3d NAT/proxy paths have different identity visibility. Do not infer identity from source IP after shared NAT.

## 4. Reliable workload identity

Bind account/resource/host/interface/task to current durable assignment; protect against stale IP reuse, spoofed packet, reassigned resource and shared proxy identity. Review guest-issued short-lived tokens/binding if needed; no token protocol is selected here. Fail closed when caller identity cannot be established. Disable before teardown and invalidate previous-generation identity on reassignment.

## 5. Ownership and implementation dependencies

B Java metadata/owned records; C policy/acceptance; A approved VM/network delivery; D DTO/wire review. Depends on actual VM/container identity/network ownership and management isolation, not on a generic HTTP listener. Backend/provider metadata is not authority over account policy. Preserve account references and secret redaction through restore/migration.

## 6. Acceptance and operations

E4: workload A cannot read B's metadata, including shared NAT/spoof/reused-IP cases; no cloud/agent/DB/S3 secret exposed. Test denied management routes, stale generation/reassignment, endpoint outage/cancellation and restoration with exact own identity. Keep request/audit limits without sensitive payloads. No current metadata service is Supported; future service requires domain gates/profile evidence.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-instance-metadata) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
