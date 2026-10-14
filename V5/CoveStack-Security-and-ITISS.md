# CoveStack V5 — Security, Infrastructure and ITISS Evidence

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. ITISS contribution

Anuj owns reproducible systems work: VM deployment, networking, hardening, DNS/TLS/tunnel, independent tests, monitoring and recovery evidence. Application developers implement code-level controls with him. **This is a project work plan, not an official C-DAC grading rubric.** Confirm the institute's expectations with the mentor.

| Domain | Deliverable | Acceptance |
|---|---|---|
| Linux/VM | Users, permissions, version baseline and services | Clean install by teammate |
| Network | Routes, VPC/address pool, firewall and proxy policy | Allowed/denied paths measured |
| Identity | Roles and two-account action matrix | Every resource action scoped |
| Hardening | Privileges, capabilities, mounts and image review | Actual configuration inspection |
| Communication | TLS, private agent route and renewal | External access/private denial |
| Vulnerability management | Scans, triage and remediation | Dated retest and residual risks |
| Observability | Metrics, logs and alert triggers | Detection of controlled failure |
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
| Proxy SSRF/config injection | Upstreams resolved from owned records |
| Domain reassignment | Verified claim and safe route cleanup |
| Metadata identity spoofing | Per-instance binding; no secret payload |
| Untrusted automation | Restricted identity/execution/state |
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
| User tokens | Protected hash lookup, expiry/revocation and permissions |
| Service secret storage | Encrypted with protected master key/recovery policy |

Signing establishes authenticity/integrity, not encryption or resource authorization. TLS provides transport confidentiality; backend policy determines account access. A real store/bootstrap and running-server tests remain necessary.

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

Do not weaken protections to support nested Docker in MVP. Rootless/VM-backed modes are evaluated deployment options, not automatic security guarantees.

## 5. Secrets and audit

| Material | Rule |
|---|---|
| Agent/tunnel/provider keys | Never public UI/general metadata |
| User DB/S3 credentials | Scoped issuance; separate authorized retrieval |
| Kubeconfig | Protected export and log redaction |
| Terraform state | Sensitive storage/access/backups/locking |
| Audit records | Actor, action, resource, result, request/time; protected writes |
| Evidence | Sanitized actual results; no secret dumps |

Append-only DB permissions reduce accidental alteration but do not justify a tamper-proof audit claim. Record the trust model and administrator authority.

## 6. Vulnerability and incident work

Record scanner/tool/input versions, findings, severity, exploitability context, remediation and verified exceptions. Scanner success is not proof of isolation. Revoke a test credential, trigger an alert, investigate logs and restore normal operation in a controlled incident exercise.

## 7. ITISS package for mentor review

Deliver threat model, diagrams, rule inventory, two-account matrix, actual reachability checks, configuration inspection, scan triage, outage timeline and restore evidence. If coding assignments are changed, record the named owner/handover; limited ITISS application coding cannot leave the production agent unowned.
