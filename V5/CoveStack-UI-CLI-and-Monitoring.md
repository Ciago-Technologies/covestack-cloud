# CoveStack V5 — React UI, CLI Parity and Monitoring

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Single source of state

**React, CLI, Terraform and Ansible use the same authenticated server services.** A UI-created resource appears in CLI results; a CLI operation appears in the UI. There is no separate synchronization database in the browser.

| Interface | Responsibility |
|---|---|
| UI | Forms, resource views, progress and authorized downloads |
| CLI | Scriptable requests, inspect/wait and useful exit codes |
| Backend | Validation, account policy, durable state and actions |
| Agent | Actual infrastructure convergence and observations |

## 2. UI screens

| Screen | Required content |
|---|---|
| Overview | Account resource counts, recent actions and stale/unknown health |
| Create resource | Supported images/versions/sizing/VPC and validation |
| Stack wizard | Template version, inputs, publication and retention choices |
| Stack details | Dependency progress, errors and child-resource links |
| Resource details | Desired/observed state, IDs and connection metadata |
| Networks | CIDR, attachments and allocation/error status |
| Publish application | Owned service/port, approved hostname and route status |
| CLI help | Copyable examples with placeholders |
| Monitoring | Actual metrics/events and time of last observation |

Enable only real supported operations. A form can be designed with mock fixtures during development, but a release must not show mock creation as successful provisioning.

## 3. Forms and secret handling

| Form field | Validation |
|---|---|
| Image/template/version | Approved supported catalog entry |
| CPU/memory/storage | Valid units/ranges and account quota |
| VPC | Owned and reachable on resource host |
| Hostname | Safe syntax, ownership and uniqueness |
| App port | Approved application endpoint |
| Persistent-data action | Explicit retain/delete choice |
| Password/token | Separate secret workflow; not ordinary metadata |

Client-side checks improve feedback; the backend revalidates every input. Do not expose service signing secrets, tunnel provider credentials or kubeconfig content in general listings. Do not add a generic browser shell-command runner.

Choose a reviewed user-auth/session approach. For cookie sessions, evaluate CSRF and cookie policy; Bearer usage requires protecting user tokens. Agent signing credentials are never browser credentials. Do not claim CORS supplies authorization.

## 4. CLI behavior

| Behavior | Requirement |
|---|---|
| Create | Returns ID/action; supports explicit wait timeout |
| Inspect/list | Same account-scoped records as UI |
| JSON output | Stable documented shape for scripts |
| Pending/failure | Truthful status; meaningful nonzero exits |
| Credentials | Restricted local config; no secret output by default |
| Remote doctor | Queries remote status; no local daemon assumption |
| Kubeconfig export | Separate authorization; preserve existing files |

CLI names in V5 examples are proposed until the executable exists. Publish actual commands and help during interface implementation, not fictional startup instructions.

## 5. Monitoring model

| Signal | Meaning | Do not infer |
|---|---|---|
| Agent last-seen | Agent communication freshness | App readiness |
| Container/Pod running | Process/runtime state | Correct app response |
| App probe | Defined endpoint readiness | Entire stack/data safety |
| Failed operation | A lifecycle attempt failed | Resource is automatically deleted |
| Host CPU/memory/disk | Measured host load/capacity | Per-tenant permission to see host metrics |
| Backup age/result | Recorded backup status | Restore success |

Poll at a modest configured interval initially; show timestamps and unknown/stale states. Streaming can be added later. Host-wide metrics are admin-only unless a reviewed scoped aggregation is supplied. Alerts should have threshold, duration, destination and a tested trigger.

## 6. User workflow acceptance

UI create → CLI inspect; CLI stop → UI update; Terraform update → both reflect same desired state; invalid/other-account actions denied in every interface. Test slow provisioning, server errors, double-submit, expiration and secret downloads. Do not duplicate provisioning logic inside React or the CLI.
