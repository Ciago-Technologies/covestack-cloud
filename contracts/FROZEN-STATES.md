# CoveStack Cloud — Frozen Lifecycle Contract

| Contract | Rule |
|---|---|
| Version | **1.1.0** |
| Source | Edit `lifecycle-rules.json`, then run `python3 contracts/generate_lifecycle.py`. |
| Review | Anuj/B/C review required; CI checks generated files. |
| Implementation | Defined vocabulary does not mean drivers/controllers are operational. |

## Resource lifecycle vocabulary

| Kind | Service | Owner | Desired states | Terminal state |
|---|---|---|---|---|
| `compute.instance` | Droplets / compute instances | agent | `active`, `off`, `archive` | `archive` |
| `db.instance` | Managed PostgreSQL | agent | `online`, `stopped`, `deleted` | `deleted` |
| `k8s.cluster` | Kubernetes / bounded k3d cluster | agent | `active`, `deleted` | `deleted` |
| `compute.group` | Autoscaling groups | server | `active`, `deleted` | `deleted` |
| `compute.snapshot` | Compute snapshots | agent | `available`, `deleted` | `deleted` |
| `network.load-balancer` | HTTP load balancers | agent | `active`, `deleted` | `deleted` |
| `storage.bucket` | Spaces-compatible buckets | server | `available`, `deleted` | `deleted` |

### Observed states — one state per row

| Kind | Allowed observed state |
|---|---|
| `compute.instance` | `new` |
| `compute.instance` | `provisioning` |
| `compute.instance` | `starting` |
| `compute.instance` | `active` |
| `compute.instance` | `stopping` |
| `compute.instance` | `off` |
| `compute.instance` | `deleting` |
| `compute.instance` | `archive` |
| `compute.instance` | `error` |
| `db.instance` | `creating` |
| `db.instance` | `starting` |
| `db.instance` | `online` |
| `db.instance` | `stopping` |
| `db.instance` | `stopped` |
| `db.instance` | `deleting` |
| `db.instance` | `deleted` |
| `db.instance` | `error` |
| `k8s.cluster` | `provisioning` |
| `k8s.cluster` | `running` |
| `k8s.cluster` | `degraded` |
| `k8s.cluster` | `deleting` |
| `k8s.cluster` | `deleted` |
| `k8s.cluster` | `error` |
| `compute.group` | `new` |
| `compute.group` | `active` |
| `compute.group` | `scaling` |
| `compute.group` | `degraded` |
| `compute.group` | `deleting` |
| `compute.group` | `deleted` |
| `compute.group` | `error` |
| `compute.snapshot` | `creating` |
| `compute.snapshot` | `available` |
| `compute.snapshot` | `deleting` |
| `compute.snapshot` | `deleted` |
| `compute.snapshot` | `error` |
| `network.load-balancer` | `new` |
| `network.load-balancer` | `provisioning` |
| `network.load-balancer` | `active` |
| `network.load-balancer` | `degraded` |
| `network.load-balancer` | `deleting` |
| `network.load-balancer` | `deleted` |
| `network.load-balancer` | `error` |
| `storage.bucket` | `creating` |
| `storage.bucket` | `available` |
| `storage.bucket` | `deleting` |
| `storage.bucket` | `deleted` |
| `storage.bucket` | `error` |

## Convergence, deletion and failure

| Item | Details |
|---|---|
| 1 | Desired state is intent; observed state is reality. Progress/degraded/error is not success. |
| 2 | Set `deleting=true` exactly for deletion intent: `archive` for instances, `deleted` otherwise. |
| 3 | Set observedGeneration only to the current target generation after the matching stable state is reached. Progress/errors/health omit it or use null. |
| 4 | Stale generations return `applied=false`; future/invalid generations are rejected. Actual timestamp, account and agent lease/fencing checks belong in B/C services. |
| 5 | Each observed report includes kind and observedAt; kind must match the stored owned resource. Groups and buckets are server-owned and cannot be agent reports/work. |
| 6 | Every heartbeat advertises supportedKinds; the server filters work accordingly. Reject unknown kinds; never mark unsupported work successful. |
| 7 | Work, known IDs and cleanup remain account/agent scoped. Outages or incomplete known-ID responses never authorize deletion. |
| 8 | No terminal-state resurrection. Desired deletion intent cannot be reversed by retries; GC waits for teardown acknowledgement. Failure does not acknowledge the target generation. |
| 9 | Logical transitions below may have unreported intermediate states. Validation permits reachable states, including same-state reports, without traversing a terminal state. Observations alone do not authorize side effects. |
| 10 | Cancellation applies only to queued operations. A failed operation attempt is terminal; a retry is a new linked attempt. |

## compute.instance

| Item | Details |
|---|---|
| 1 | active only after container and required readiness checks succeed; off after intentional stop; archive after owned workload teardown. |
| 2 | Success mapping: `active` → `active`; `off` → `off`; `archive` → `archive`. |

| From | Allowed next logical states |
|---|---|
| `new` | `provisioning`, `deleting`, `error` |
| `provisioning` | `active`, `off`, `deleting`, `error` |
| `starting` | `active`, `deleting`, `error` |
| `active` | `stopping`, `deleting`, `error` |
| `stopping` | `off`, `deleting`, `error` |
| `off` | `starting`, `deleting`, `error` |
| `deleting` | `archive`, `error` |
| `archive` | Terminal; no outgoing transition |
| `error` | `provisioning`, `starting`, `stopping`, `deleting` |

## db.instance

| Item | Details |
|---|---|
| 1 | online requires successful database connection/readiness; preserve data volumes through stop/recreate; deletion follows explicit retention policy. |
| 2 | Success mapping: `online` → `online`; `stopped` → `stopped`; `deleted` → `deleted`. |

| From | Allowed next logical states |
|---|---|
| `creating` | `online`, `stopped`, `deleting`, `error` |
| `starting` | `online`, `deleting`, `error` |
| `online` | `stopping`, `deleting`, `error` |
| `stopping` | `stopped`, `deleting`, `error` |
| `stopped` | `starting`, `deleting`, `error` |
| `deleting` | `deleted`, `error` |
| `deleted` | Terminal; no outgoing transition |
| `error` | `creating`, `starting`, `stopping`, `deleting` |

| Item | Details |
|---|---|
| 1 | Reserved future states (not accepted in this contract): `resizing`, `migrating`, `forking`. Add schemas/operations/tests in a reviewed revision before enabling them. |

## k8s.cluster

| Item | Details |
|---|---|
| 1 | running requires reachable Kubernetes API and required nodes ready; degraded is not successful convergence. |
| 2 | Success mapping: `active` → `running`; `deleted` → `deleted`. |

| From | Allowed next logical states |
|---|---|
| `provisioning` | `running`, `degraded`, `deleting`, `error` |
| `running` | `degraded`, `deleting`, `error` |
| `degraded` | `running`, `deleting`, `error` |
| `deleting` | `deleted`, `error` |
| `deleted` | Terminal; no outgoing transition |
| `error` | `provisioning`, `deleting` |

## compute.group

| Item | Details |
|---|---|
| 1 | active requires desired capacity of healthy children; scaling includes pending reservations; deleted only after scoped child policy completes; never agent work. |
| 2 | Success mapping: `active` → `active`; `deleted` → `deleted`. |

| From | Allowed next logical states |
|---|---|
| `new` | `scaling`, `active`, `deleting`, `error` |
| `active` | `scaling`, `degraded`, `deleting`, `error` |
| `scaling` | `active`, `degraded`, `deleting`, `error` |
| `degraded` | `scaling`, `active`, `deleting`, `error` |
| `deleting` | `deleted`, `error` |
| `deleted` | Terminal; no outgoing transition |
| `error` | `scaling`, `deleting` |

## compute.snapshot

| Item | Details |
|---|---|
| 1 | available requires immutable account-scoped image persisted and inspected; container snapshots exclude mounted volumes and VM memory. |
| 2 | Success mapping: `available` → `available`; `deleted` → `deleted`. |

| From | Allowed next logical states |
|---|---|
| `creating` | `available`, `deleting`, `error` |
| `available` | `deleting`, `error` |
| `deleting` | `deleted`, `error` |
| `deleted` | Terminal; no outgoing transition |
| `error` | `creating`, `deleting` |

## network.load-balancer

| Item | Details |
|---|---|
| 1 | active requires valid configuration, successful reload, reachable listener and healthy required members; empty/unhealthy targets are degraded. |
| 2 | Success mapping: `active` → `active`; `deleted` → `deleted`. |

| From | Allowed next logical states |
|---|---|
| `new` | `provisioning`, `deleting`, `error` |
| `provisioning` | `active`, `degraded`, `deleting`, `error` |
| `active` | `degraded`, `deleting`, `error` |
| `degraded` | `active`, `deleting`, `error` |
| `deleting` | `deleted`, `error` |
| `deleted` | Terminal; no outgoing transition |
| `error` | `provisioning`, `deleting` |

## storage.bucket

| Item | Details |
|---|---|
| 1 | available requires physical MinIO bucket and account-scoped access policy; delete empty buckets by default, require explicit authorization to purge objects; server-managed through trusted MinIO administration, with root credentials kept server-side. |
| 2 | Success mapping: `available` → `available`; `deleted` → `deleted`. |

| From | Allowed next logical states |
|---|---|
| `creating` | `available`, `deleting`, `error` |
| `available` | `deleting`, `error` |
| `deleting` | `deleted`, `error` |
| `deleted` | Terminal; no outgoing transition |
| `error` | `creating`, `deleting` |

## Server-owned operation and notification lifecycles
### operation

| Item | Details |
|---|---|
| 1 | cancelled only before execution; failures are terminal for that attempt; retry creates a linked new attempt; complete only on confirmed target convergence. |

| From | Allowed next |
|---|---|
| `queued` | `running`, `failed`, `cancelled` |
| `running` | `completed`, `failed` |
| `completed` | Terminal |
| `failed` | Terminal |
| `cancelled` | Terminal |

### notification

| Item | Details |
|---|---|
| 1 | post-commit outbox only; bounded retry policy; SMTP failure never rolls back a committed resource operation; permanent failures remain failed. |

| From | Allowed next |
|---|---|
| `queued` | `delivering`, `failed` |
| `delivering` | `delivered`, `failed` |
| `failed` | `queued` |
| `delivered` | Terminal |

## Services without agent resource lifecycles

| Service | Owner | Policy |
|---|---|---|
| Accounts / membership | server | Account activation/suspension and memberships belong to authorization policy, not agent resource work. |
| Access keys / Bearer tokens | server | Active/revoked/expired eligibility, key kind and permissions; no resource lifecycle or agent work. |
| SSH keys | server | Account-scoped metadata; validated public keys; referenced only by owned instances. |
| Tags | server | Account-scoped metadata and associations; reserved ownership labels are not editable user tags. |
| Regions / sizes / images catalogs | server | Read-only provider metadata; snapshots have their own resource lifecycle. |
| Spaces objects | storage | S3 put/get/list/delete behavior within an owned bucket; no asynchronous resource-state enum. |
| CLI / doctor / health | client/server | API clients/status endpoints, not agent resource kinds; detailed doctor requires authentication. |

## Provider mapping and storage boundaries

| Item | Details |
|---|---|
| 1 | These are internal states. D maps them to provider DTOs/status vocabulary; internal strings must not be copied blindly into public responses. Operation statuses are separate from resource states. E/C define administrative credential/account policy. MinIO/Postgres/Valkey/Mailpit daemon readiness is infrastructure health, not a new agent resource kind. Managed databases are account resources; control-plane Postgres is infrastructure. Storage buckets/objects are not compute snapshots. |

## Example: Kubernetes convergence

| Item | Details |
|---|---|
| 1 | Java creates desiredState=active, state=provisioning, generation=1. Agent reports kind=k8s.cluster, state=running, observedAt and observedGeneration=1 only after readiness. Delete sets desiredState=deleted and deleting=true, increments generation; deleting progress omits observedGeneration, then deleted acknowledges that generation. |

## Acceptance

| Item | Details |
|---|---|
| 1 | Run `python3 contracts/check_lifecycle.py`, Go lifecycle tests, Java lifecycle tests and real scoped controller/driver integration. All planned resource kinds have examples; groups/buckets appear only in server-resource examples. Do not claim those examples prove a service is operational. |
