# CoveStack V5 — Ownership, Reviews and ITISS Responsibilities

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Five-person plan

The existing A–E assignments are preserved. New frontend/provider work is a proposed extension of those assignments; confirm capacity in G0. Limited ITISS application coding does not silently transfer Anuj's existing agent ownership.

| Person | Primary implementation | V5 additions |
|---|---|---|
| A — Anuj | Go agent, drivers, infrastructure, CI | VPC/IPAM integration, ingress, metadata networking, ITISS evidence |
| B | Java core, migrations, resources/actions | Stack orchestration, leases, subnet/domain records, ASG |
| C | Authentication, permissions and security | Domain claims, secrets, metadata/SSRF policy, independent tests |
| D | Provider DTOs/endpoints and compatibility | Terraform provider API mappings; API/schema review |
| E | Management API, CLI, packaging and mail | React SPA and Ansible inventory/workflow examples |

D/E's extra integrations need explicit resourcing. A reviews the Go provider lifecycle and infrastructure; B owns server semantics; C approves authorization. Reassign a named task if workload exceeds capacity rather than giving every feature to A.

## 2. Anuj's current sequence

| Priority | Work | Acceptance |
|---|---|---|
| 1 | Review config and add signed client/polling main | Real auth plus bounded error behavior |
| 2 | Compute reconciliation | Create/stop/start/delete; no restart duplicates |
| 3 | Networking and publication | Collision-free subnets; external HTTPS and denied private routes |
| 4 | Kubernetes and persistent services | Bounded cluster; retained data and private endpoints |
| 5 | Metadata and operations | Self-only data; monitoring and restore |

The local config implementation exists, but review/commit/pinned-toolchain evidence remain separate. Anuj should learn the relevant Go patterns and explain accepted code; the team must not assign him unexplained generated application code and call it ITISS work.

## 3. Java core and contracts — B

| Deliverable | Completion evidence |
|---|---|
| Main/migrations and real configuration | Fresh DB and packaged boot |
| Scoped resource/action transactions | Cross-account and concurrent-create tests |
| Agent work/observations and fencing | Old lease cannot write; stale generation ignored |
| Stack DAG and retries | Partial failure recorded; retry resumes safely |
| VPC/lease/domain records | Atomic uniqueness and ownership |
| ASG controller | Pending count, limits, cooldown and outage pause |
| Persistent-data policy | Delete versus retain choices and restore metadata |

## 4. Authentication and security — C

Implement real credential/token stores, revocation/expiry, permissions, encrypted signing-secret storage and protected audit evidence. Coordinate signed-byte and nonce behavior with A. Never use a public bootstrap endpoint that issues unrestricted account/agent keys.

| Boundary | Required test |
|---|---|
| User versus agent credentials | Cannot swap credential kinds |
| Object/function authorization | Every owned-resource and admin action scoped |
| Replay | Valid duplicate rejected; bad MAC cannot poison nonce |
| Shared storage | A credential cannot operate on B bucket |
| Domains | Unverified/reassigned hostname cannot serve another tenant |
| Metadata | No cross-instance selection or credential leak |
| Automation runner | Job cannot access another account/host secret |

## 5. Provider-facing integration — D

Maintain the exact tested provider-style subset and explicit unsupported responses. Build Terraform incrementally against the common authenticated APIs: one resource lifecycle, refresh/import, update, delete, then stack support. Avoid a second implementation bypassing core policy.

Coordinate Kubernetes versions/shapes, snapshot mapping and client-facing connection metadata with A/B. Standard DigitalOcean tooling is not automatically compatible with a custom management API; show the exact endpoint/client tests that pass.

## 6. Interfaces and configuration workflows — E

| Area | Requirement |
|---|---|
| CLI | Resource/stack/API operations and truthful pending status |
| React | Same records/actions as CLI; forms only for supported behavior |
| Help | Copyable placeholder commands; no embedded credentials |
| Ansible | Owned inventory, scoped connectivity and repeatable example |
| Kubeconfig | Authorized export with no accidental overwrite |
| Packaging | Clean install, remote doctor, artifact checksums |
| Mail | Post-commit outbox; SMTP failure does not roll back resource creation |

## 7. ITISS evidence ownership

Anuj leads the VM/network/privilege plan, DNS/TLS/tunnel deployment, testable isolation, vulnerability triage, monitoring, incident exercises and recovery. C collaborates on security policy; application developers implement the backend changes those controls require.

If the mentor expects ITISS students to avoid extensive app coding, propose a named agent developer and agreed handover. Infrastructure configuration, shell/Ansible automation and API testing remain technical work requiring understanding and reproducibility.

## 8. Required pair reviews

| Change | Reviewers |
|---|---|
| Signing/wire contract | A + B + C |
| Compute/observations/leases | A + B |
| VPC/IPAM and gateway | A + B + C |
| Provider compatibility/Terraform | D + B + A |
| UI/CLI/Ansible | E + B + C; A for host connectivity |
| Domain/metadata | A + C + B |
| Stack retention/rollback | B + A + C + E |

## 9. Definition of task done

A task includes behavior, meaningful tests, updated supported-feature status, sanitized evidence and a reviewable PR. Documentation can describe the target, but only tested implementation and agreed acceptance mark the feature complete. Use [workflow](CoveStack-Project-Workflow.md) for the authoritative dated plan.
