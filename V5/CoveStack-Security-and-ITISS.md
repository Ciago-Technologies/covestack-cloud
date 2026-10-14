# CoveStack V5 — Security, Infrastructure and ITISS Evidence

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. ITISS contribution

A (Anuj) owns all agent/client/driver/reconcile Go, VM/infrastructure hardening, CI, Ansible Part 1 and ITISS evidence. C owns all acceptance tests from day 1 and reviews security; B owns SQL; D Java API mappings; E thin CLI. Extended infrastructure evidence is relocated to the roadmap. **This is a project work plan, not an official C-DAC grading rubric.** Confirm the institute's expectations with the mentor.

| Domain | Deliverable | Acceptance |
|---|---|---|
| Linux/VM | Users, permissions, version baseline and services | Clean install by teammate |
| Network | Server-owned VPC allocator and private bridge policy | Test 14 plus allowed/denied paths |
| Identity | Roles and two-account action matrix | Every resource action scoped |
| Hardening | Privileges, capabilities, mounts and image review | Actual configuration inspection |
| Communication | Private signed agent path and protected API | Real auth-smoke; no private credential leakage |
| Vulnerability management | Scans, triage and remediation | Dated retest and residual risks |
| Evidence | Sanitized CI/logs, versions and failure timeline | Tests 1–15; no mock promotion |
| Recovery | Backups, restore and incident procedures | Known data recovered separately |

## 2. Threat model

| Threat | Required control/test |
|---|---|
| A operates on B's resource | Object/action authorization in backend |
| User key calls agent API | Credential-kind separation |
| Signed replay | Authenticate before atomic nonce claim |
| Invalid MAC consumes nonce | Valid preserved request still succeeds |
| Agent takeover/host exposure | Trusted high-authority boundary; no workload daemon socket |
| Network crossing | Explicit Docker/Kubernetes routing tests |
| Outage causes deletion | Failed/partial inventory never authorizes cleanup |

Threats include compromised workload/agent and accidental operator mistakes. Simulate them only in the team's authorized lab. Object-level authorization is not replaced by random IDs, a login screen or firewall rules. [OWASP object authorization](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/).

## 3. Signing and user identity

| Existing compatibility rule | Requirement |
|---|---|
| `KAN1-HMAC-SHA256` | Preserve current shared algorithm |
| Date/key/nonce/Host | Required authenticated input |
| Body/path/raw query | Exact bytes across client/proxy/server |
| 1 MiB signed request body | Enforce actual bytes, including chunked input |
| Replay store | Atomic authenticated claim; fail closed on unavailable store |
| User tokens | api_tokens hash-only lookup, expiry/revocation and permissions |
| Service secret storage | signing_credentials ciphertext/key version; service kind; custody B+C |

Signing establishes authenticity/integrity, not encryption or resource authorization. TLS provides transport confidentiality; backend policy determines account access. A real store/bootstrap and running-server tests remain necessary. Each poll derives account scope from its verified service principal; trusted credential-to-agent assignment and multi-account polling remain open B+C+A. Never treat a label/caller-selected account as authorization.

## 4. Host and workload authority

A normal privileged Docker daemon is powerful. A service user with socket access may effectively control the host. Removing root from the agent process does not remove the authority of that socket. [Docker security](https://docs.docker.com/engine/security/).

| Boundary | Policy |
|---|---|
| User workload | No host socket, broad mounts or unnecessary privilege |
| Platform agent | Trusted, protected identity and restricted exposure |
| Lab host | Dedicated VM boundary from personal machine |
| Tenant containers | Not equivalent to independent VM isolation |
| k3d nodes | Special privileges reviewed; never generic tenant privilege grant |
| Rootless option | Evaluate feature compatibility and actual limitations |

Rootless compatibility and stronger workload isolation remain separate evaluated options. Former nested/VM-backed designs are Optional future work in the roadmap; CORE-5 Droplets are containers.

## 5. Secrets and audit

| Material | Rule |
|---|---|
| Agent/signing keys | Never ordinary resource responses or client logs |
| User DB/S3 credentials | Scoped issuance; separate authorized retrieval |
| Kubeconfig | Protected export and log redaction |
| Terraform state | Sensitive storage/access/backups/locking |
| Audit records | Actor, action, resource, result, request/time; protected writes |
| Evidence | Sanitized actual results; no secret dumps |

MinIO administration is server-side; root/admin keys are distinct from tenant credentials. Custody, rotation and server-to-MinIO private path remain B+C decisions. The schema shows nullable audit_log.account_id as a proposed option for unauthenticated events; separate security events versus nullable audit remains B+C decision, with operator-only access and no tenant leak.

Append-only DB permissions reduce accidental alteration but do not justify a tamper-proof audit claim. Record the trust model and administrator authority.

## 6. Vulnerability and incident work

Record scanner/tool/input versions, findings, severity, exploitability context, remediation and verified exceptions. Scanner success is not proof of isolation. Revoke a test credential, investigate sanitized logs and restore normal operation in a controlled exercise. Alert/metrics and public-edge scenarios remain preserved in the roadmap.

## 7. ITISS package for mentor review

Deliver threat model, diagrams, rule inventory, two-account matrix, actual reachability checks, configuration inspection, scan triage, outage timeline and restore evidence. Ask mentor in writing whether ITISS students may own Go agent/driver code; if not, name an agent developer and record handover. Confirmation is pending. AI-generated code is reference only; A types and can explain agent Go. E never takes agent/signing/drivers/reconcile ownership silently.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
