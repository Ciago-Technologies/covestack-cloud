# CoveStack V5 — Stack Templates and Complete Environment Workflow

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. The central user experience

**A stack is an owned environment containing references to individually managed resources.** One request coordinates dependencies, configuration, readiness and publication. Users still inspect each resource's metadata; the stack is not a replacement for its constituent services.

| Template | Resources | Gate |
|---|---|---|
| Single web app | Droplet, VPC, optional HTTPS route | First vertical slice |
| Web/data/storage | Droplet app, PostgreSQL, bucket, VPC, publication | First full environment |
| Kubernetes web/data | Cluster/workload, PostgreSQL, bucket, network, publication | After cluster and full-stack dependencies |
| Customizable environment | Validated supported components and dependencies | After fixed templates work |

Selecting a template is not enough to create an arbitrary application configuration. The app must accept the template's documented environment variables or secret references. Unsupported application assumptions produce an explicit validation error.

## 2. End-to-end example

| Step | User input/action | CoveStack work |
|---|---|---|
| Select | Web/data/storage template | Load immutable template version |
| Configure | Image, app port, sizing, region and network needs | Validate supported choices and quota |
| Publish choice | Enable HTTPS with approved hostname | Reserve route; do not expose private services |
| Submit | UI, CLI or provider request | Store stack/action atomically |
| Provision | User observes progress | Create network and persistent services first |
| Connect | No manual password copying required | Issue scoped secrets and pass references |
| Deploy | Supported app configuration | Start workload and check readiness |
| Publish | Route usable | Return actual public URL |
| Inspect | CLI/UI resource view | Same stack/child IDs and status |
| Remove | Retain/delete data choice | Reverse dependencies with safe policy |

Example proposed CLI, **not an existing executable command**:

```bash
covestack stack create shop --template web-postgres-storage
covestack stack inspect shop
```

## 3. Template contract

| Field | Required meaning |
|---|---|
| Template ID/version | Immutable selected behavior; no silent mutable updates |
| Input schema | Types, ranges, approved images and supported flags |
| Dependencies | Acyclic ordering and completion prerequisites |
| Outputs | Resource references and reachable endpoints |
| Secret wiring | Protected references, not plain JSON in listings |
| Readiness | Actual criteria per resource/application |
| Retention | Explicit persistent-data policy |
| Update policy | Supported edits and replacement consequences |

A conceptual template may declare dependencies in YAML, but the schema below is **proposed CoveStack format**, not Docker Compose or an already implemented parser:

```yaml
apiVersion: covestack.example/v1
kind: EnvironmentTemplate
metadata:
  name: web-postgres-storage
spec:
  version: "1"
  components:
    - name: network
      type: vpc
    - name: database
      type: postgres
      dependsOn: [network]
    - name: storage
      type: s3-bucket
      dependsOn: [network]
    - name: app
      type: droplet-app
      dependsOn: [database, storage]
  publication:
    enabledByUser: true
  retention:
    persistentData: retain
```

Validate cycles, unknown components, unbounded sizes and account ownership before provisioning. User customization operates within the schema; it is not raw shell/proxy/privileged-container configuration.

## 4. Stack state and failure policy

Proposed stack states are separate from the existing resource contract and require review before implementation.

| Stack state | Meaning |
|---|---|
| Validating/queued | Request accepted for checks/scheduling |
| Provisioning | Child operations running in dependency order |
| Ready | Defined required outputs/readiness satisfied |
| Degraded/failed | Required component unavailable or provisioning failed |
| Updating/deleting | Explicit action in progress |
| Deleted | Required cleanup complete; retention references recorded |

Never overwrite a child's resource state to make the stack look ready. Show which components are pending, failed or retained. Store last error, operation attempt and timestamps without sensitive payloads.

| Failure | Required response |
|---|---|
| DB fails before app | Stop dependent step; expose cause and bounded retry |
| Publication fails after app ready | App remains private; report publication failure |
| Agent restarts | Resume from durable references; inspect before create |
| Duplicate submit | Same idempotent result or conflict, not second stack |
| Partial delete | Preserve retryable cleanup records; no hidden leak |
| Persistent-data retention | Keep protected data ownership and restoration info |

Compensating cleanup is not unconditional rollback. An app failure must not erase a database that contains user data. Default release policy retains persistent data unless the authorized user explicitly requests deletion.

## 5. Single ownership model

UI and CLI read the same server records. Terraform either manages an entire stack or separate child resources: do not let both models concurrently own the same child's desired specification. Ansible configures approved targets; the stack/provider owns resource provisioning. Record external-management ownership and supported adoption/import behavior.

## 6. Acceptance

Create through UI and inspect through CLI; repeat the same request without duplication; fail a dependency and resume; restart agent/server during provisioning; deny another account; delete with retained data; confirm actual endpoints and secret scoping. Compose already offers multi-service definitions; the evidence must establish CoveStack's controlled self-service and policies.
