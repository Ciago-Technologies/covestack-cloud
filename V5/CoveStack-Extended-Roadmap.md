# CoveStack V5 — Extended Development Roadmap

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |
| Track status | Extended track |

## 1. Extended development track — NOT part of SRS or deliverables

Status: **Extended track**. Entry: a feature starts only after its CORE-5 dependency has Gates passed. B/C/D/E may start server-side extended design/API/controller work from **1 December 2026** only when their CORE-5 duties are not blocked. CORE-5 reviews and failure tests always take priority. A begins extended Go execution/reconciliation only after **all CORE-5 gates pass**. If A is one week behind any checkpoint, stop all extended work.

No service here is **Supported** until status **Gates passed** and the same G1–G4 rule is evidenced: DB desired state/generation; crash-safe rebuild without duplicate; other-account API 404; CI failure test. Interfaces reuse gated backend resources and add their interface tests rather than inventing tables.

Extended features are excluded from SRS/deliverables. They may be shown only as **bonus** after Gates passed. The track continues **after 25 January 2027** as post-submission development. Cut the entire extended track first if CORE-5 is late.

## 2. Order and execution model

ASG → Load balancing → Application publication → Metadata → Monitoring → Snapshots → Domains → read-only list page → Ansible Part 2. ASG/LB reuse the Droplet controller and are the cheapest next candidates after CORE-5, not guaranteed effort estimates.

Still two workload drivers: **generic container** and **k3d**. Networking is a container-driver helper. Extended proxy configuration, health checking and snapshot work are A-owned reconcilers/helpers using those backends, not automatically additional driver types.

## 3. Auto scaling groups

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Approved compute membership and scaling policy |
| Platform-value note | Execution backend only (container driver executes instances). Platform value: B controller decides policy |
| DB tables | asg_groups, asg_instances; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Reservations, membership and min/max policy |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | B controller; C tests; A execution after CORE-5 |
| Gate test | E1; C owns acceptance |
| CORE-5 dependency | Droplets |
| Decisions | Cooldown defaults and selected metric: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 4. Load balancing

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Owned HTTP backend distribution and health |
| Platform-value note | Execution backend only (HTTP proxy, tool TBD). Platform value: owned health/membership/config recovery |
| DB tables | load_balancers, lb_backends, health_check_configs; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Owned membership/config; observed health remeasured |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | D API mappings; B records; C tests; A after CORE-5 |
| Gate test | E2; C owns acceptance |
| CORE-5 dependency | Droplets and VPC/IPAM |
| Decisions | Proxy tool and health method: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 5. Application publication

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Explicitly requested website exposure through shared gateway |
| Platform-value note | Execution backend only (tunnel/reverse proxy, tool TBD). Platform value: authorized route lifecycle |
| DB tables | publications; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Registered hostname, resource and approved port |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | B records; C policy; A after CORE-5 |
| Gate test | E5; C owns acceptance |
| CORE-5 dependency | Droplets or Kubernetes and VPC/IPAM |
| Decisions | Tunnel provider, proxy topology and TLS method: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

Simulated public IP and direct-IP browsing attach here, not additional services. A simulated value does not route internet traffic. Real browser access needs a real gateway/hostname and DNS/TLS; managed application publication remains explicit, never automatic exposure of private management ports.

## 6. Metadata

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Bound instance self-information |
| Platform-value note | Execution backend only (read-only metadata listener, delivery TBD). Platform value: identity and secret exclusion |
| DB tables | metadata_identities; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Self-identity and permitted policy fields |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | C policy; B records; A network delivery after CORE-5 |
| Gate test | E4; C owns acceptance |
| CORE-5 dependency | Droplets |
| Decisions | Delivery, route/namespace and provider subset: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 7. Monitoring

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Observed health, metrics, events and alerts |
| Platform-value note | Execution backend only (measurement collector, choice TBD). Platform value: freshness and scoped alerts |
| DB tables | metric_samples, alert_rules, alerts; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Alert rules; samples remain evidence, not desired-state authority |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | C policy/tests; B records; A collection after CORE-5 |
| Gate test | E7; C owns acceptance |
| CORE-5 dependency | CORE-5 resource/agent APIs |
| Decisions | Metrics, freshness thresholds and alert destinations: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 8. Snapshots

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Owned instance-filesystem capture/restore |
| Platform-value note | Execution backend only (container driver image operations). Platform value: owned immutable references |
| DB tables | snapshots; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Source/image ownership and artifact references |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | D mappings; B records; C tests; A after CORE-5 |
| Gate test | E3; C owns acceptance |
| CORE-5 dependency | Droplets |
| Decisions | Snapshot format/storage and retention: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 9. Domains

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Platform subdomains and verified custom hostnames |
| Platform-value note | Execution backend only (DNS/TLS provider, choice TBD). Platform value: claim ownership and reassignment safety |
| DB tables | domains, domain_verifications; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Current claim and verification generation |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | C verification; B records; A routes after CORE-5 |
| Gate test | E6; C owns acceptance |
| CORE-5 dependency | Application publication |
| Decisions | Verification method and certificate/DNS flow: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 10. React UI

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Read-only list page first |
| Platform-value note | Execution backend only (browser API client). Platform value: shared scoped state, never runtime authority |
| DB tables | No new tables; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Fetch authenticated list/get JSON; no browser resource registry |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | E; C tests; A reviews any Go change |
| Gate test | Interface gate + applicable tenant tests; C owns acceptance |
| CORE-5 dependency | CORE-5 APIs |
| Decisions | Frontend/auth transport choice: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

## 11. Ansible inventory (Part 2)

| Field | Extended-track target |
|---|---|
| Status | Extended track |
| Purpose | Owned targets from CLI/API |
| Platform-value note | Execution backend only (Ansible inventory consumer). Platform value: authenticated scoped target selection |
| DB tables | No new tables; [Layer 2 schema](CoveStack-Database-Schema.md#appendix-a-extended-track-schema--not-in-v1-migrations) |
| Rebuilt from DB | Fetch owned inventory through authenticated API |
| Tenant rule | Other-account API returns 404; protocol-specific denial also tested |
| Owner | E inventory; A SSH/infra review; C tests |
| Gate test | E8; C owns acceptance |
| CORE-5 dependency | Droplets and API SSH prerequisites |
| Decisions | Approved image/SSH/Python/connectivity: TBD decision record |

Missing lifecycle/state/message strings are **TBD with B**. New types require coordinated schemas, fixtures and generated outputs; contract 1.1.0 is unchanged.

Inventory script prints JSON derived from `covestack droplet list -o json`; **NEVER docker ps**. It talks only to the authenticated API through the CLI/client. E8: Bob inventory contains none of Alice Droplets. Prerequisites: an approved sshd + Python base image and owned SSH-key injection through the API. This is distinct from the CORE-5 operator installation playbook; it creates no DB tables.

## 12. Optional future work — not additional extended services

| Catalog item | Classification |
|---|---|
| Nested Docker | Optional future work; privilege/storage isolation review |
| VM-backed Droplets | Optional future work; CORE-5 Droplets remain containers |
| Multi-host VPC | Optional future work; routing/distributed allocation |
| More DB engines | Optional future work; engine-specific restore/policy |
| Full provider parity | Optional future work; no compatibility claim |
| Broader template customization/DAG | Preserved deferred design; not the single CORE-5 stack |
| Unrestricted IaC jobs | Preserved deferred design; separate execution threat model |

Do not mount the platform Docker socket into a user workload. All tool choices remain decision records; existing references/examples below are historical candidates, not selections.

## 13. Preserved design text and relocation register

The following text is preserved from existing V5 topic guides so requirements are not lost. **Every requirement here is an Extended track or Optional future work target, not a current capability or CORE-5 deliverable.** Old phase names/examples are not approved wire/schema strings; missing vocabulary is TBD with B. Old scheduling/ownership assumptions are superseded by sections 1–11.

Batch 3 replaces source-topic detail with CORE-5 guidance or relocation pointers; this appendix preserves the deferred detail. It preserves design intent, not old claims of availability. CORE-5 topics appearing as prerequisites here remain governed by the main design and gate register.

### Source: CoveStack-Instance-Metadata.md

**Relocation status: Extended track; historical design, not availability.**

#### 1. Two meanings of metadata

| Meaning | Reader | Example |
|---|---|---|
| Control-plane resource information | Authorized UI/CLI/API user | ID, state, network and endpoint |
| Instance metadata service | Workload querying itself | Own ID, hostname, private IP and tags |

These are different services. The current catalog of accounts/tags/regions does not mean an IMDS implementation exists.

#### 2. Provider contract

| Contract | Endpoint model | V5 handling |
|---|---|---|
| DigitalOcean | `/metadata/v1/` at link-local address | Document/test selected Droplet subset |
| AWS EC2 | `/latest/meta-data/` and IMDSv2 token flow | Separate future compatibility module |
| CoveStack custom fields | Explicit namespace | Real publication URL and simulator attributes |

Do not call generic JSON automatically AWS-compatible; many leaf metadata responses are text. DigitalOcean and AWS use different paths and semantics. [Droplet metadata](https://docs.digitalocean.com/products/droplets/how-to/access-metadata/), [AWS IMDSv2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html).

#### 3. Bounded first response

| Field | Policy |
|---|---|
| Instance ID/hostname | Bound to this workload |
| Private address/VPC | Actual inspected assignment |
| Region/tags | Reviewed self-information subset |
| Public app URL | Separate custom field, if published |
| Simulated address | Clearly marked; not a working URL |
| User data | Excluded initially or separately reviewed for sensitivity |
| Signing/provider/database secrets | Never included |

#### 4. Delivery stages

| Stage | Design | Acceptance |
|---|---|---|
| First | Per-instance read-only sidecar; explicit internal endpoint | Only its bound instance data |
| Compatibility | Namespace-local `169.254.169.254` listener | Correct routes/binding and provider response tests |
| Shared gateway alternative | Enforced identity/anti-spoofing | Cannot select another instance |

An HTTP listener does not automatically assign a virtual IP or intercept packets. Link-local support needs namespace/address/route setup. A controlled helper can configure the interface; the user workload should not receive host network administration privileges.

A per-instance sidecar sharing the instance network namespace can serve data bound to that instance. Assigning a link-local /32 inside separate namespaces avoids a single ambiguous host listener, but must be prototyped against actual Docker mode and host routes. Check that no genuine host/cloud metadata endpoint leaks into the workload.

#### 5. Caller identity

| Untrusted input | Why insufficient |
|---|---|
| Requested instance/account ID | User can change it |
| `X-Forwarded-For` | Can be spoofed without a trusted proxy boundary |
| Source IP alone | NAT/spoofing/reuse can break identity |
| Token alone | Must be bound, scoped and validated under the chosen model |

The simplest MVP has no caller-selected target: the sidecar has only its assigned instance data. A token may mitigate some SSRF patterns but does not make an already authorized compromised workload unable to read its own metadata. Document that boundary.

#### 6. Tests and operational policy

Deny public ingress to metadata, cross-instance lookup, forged identifiers, stale data after delete/reuse and access to secret fields. Bound request rates/response sizes. Metadata can use last approved local state during a control-plane outage, with freshness explicitly represented. Delete/recreate must not deliver the old instance's account data to a new owner.

No link-local compatibility claim is accepted until a real workload query succeeds with the exact expected response and negative identity tests pass.

### Source: CoveStack-Publication-Domains-and-Tunneling.md

**Relocation status: Extended track; historical design, not availability.**

#### 1. Target publication capability

**Extended-track target: publication for eligible Droplets and Kubernetes applications. Exposure is explicit; no current capability is asserted.** One shared tunnel/gateway can carry many hostname routes; each published app receives a validated owned mapping. This is not one new tunnel per container.

```text
Internet browser
   -> real HTTPS hostname / public edge
   -> shared tunnel or VPS gateway
   -> trusted local reverse proxy
      -> Droplet application port
      -> Kubernetes ingress -> Service -> ready Pods
```

| Layer | Responsibility |
|---|---|
| Public DNS/TLS | Real hostname and trusted certificate |
| Tunnel | Reach the private lab without inbound router forwarding |
| Reverse proxy | Dispatch approved hostnames to approved upstreams |
| Application | Listen on expected interface/port and satisfy health criteria |
| Control plane | Account permission, domain claim and route lifecycle |

#### 2. Request and route lifecycle

| Stage | Required behavior |
|---|---|
| Select resource | Must belong to the caller |
| Select service/port | Allowed app endpoint, not arbitrary upstream URL |
| Reserve hostname | Unique account/resource binding |
| Configure | Generate validated config safely; atomic apply/reload |
| Verify | Test endpoint and record actual readiness |
| Return | Public URL separate from private/simulated addresses |
| Stop/failure | Truthful unavailable status |
| Delete/unpublish | Remove routes before address/resource reuse |

The user should not paste Docker IPs into proxy configuration. Resolve upstreams from trusted resource state. Reject configuration injection, arbitrary private URLs and attempts to publish the control-plane DB, replay store, metadata or daemon.

#### 3. Proxy topology choices

| Topology | Consequence |
|---|---|
| Proxy on host | Can use explicitly loopback-published workload ports |
| Proxy in container | Its localhost is not host localhost; requires approved connectivity |
| Proxy joined to tenant bridges | Increases cross-network reach; review compromise boundary |
| Kubernetes path | Stable Service/Ingress with mapped cluster ingress port |

Record the chosen method and measure isolation. Kubernetes Pod IPs can change. Cluster admin connectivity is separate from website ingress; do not publish admin APIs just to serve an app.

Cloudflare is one candidate for public hostname routing; a VPS is another. Do not require a specific vendor unless the team's chosen deployment does. [Tunnel routing](https://developers.cloudflare.com/tunnel/concepts/routing/), [k3d service exposure](https://k3d.io/stable/usage/exposing_services/).

#### 4. Domain types

| Type | Initial handling |
|---|---|
| Team-owned subdomain | Automatically assigned within reviewed zone |
| Customer subdomain | Ownership challenge, DNS target and certificate validation |
| Apex domain | Provider-specific DNS/certificate method |
| Local-only name | Lab DNS; not public internet access |
| Fake public IP | Not a substitute for any domain route |

One certificate wildcard covers one label depth. Validate the actual `app-id.apps.domain` certificate rather than assuming `*.domain` covers it. Domain registrar and DNS provider can differ.

Ordinary tunnel setup for an owned zone is not automatic arbitrary-domain support. Use a supported custom-hostname product or a tested VPS certificate flow. Validate hostname ownership separately from certificate issuance. [Custom hostname validation](https://developers.cloudflare.com/cloudflare-for-platforms/cloudflare-for-saas/domain-support/hostname-validation/).

#### 5. Security and operations

| Condition | Rule |
|---|---|
| Unknown Host/SNI | Reject/404; never default to another tenant |
| Unverified customer domain | No serving tenant traffic |
| Stale route after delete | Remove before reassignment |
| Provider token | Restricted operator secret; no browser exposure |
| Forwarded headers | Trust only known proxies; no spoofed identity |
| Internal signed API | Private route; preserve original signed authority |
| Tunnel outage | Show endpoint unavailable; resources/data remain |
| TLS expiry | Monitor and test renewal/reload |

#### 6. First acceptance demo

Publish a Droplet app and a Kubernetes app under separate real HTTPS hostnames. Access both from mobile data or another unrelated network without VPN. Deny an unknown hostname and another account's publish request. Stop the app, restart the tunnel, delete the resource and verify route removal. Only real HTTP/HTTPS app access is promised; SSH/database/UDP publication requires separate protocol design.

### Source: CoveStack-UI-CLI-and-Monitoring.md

**Relocation status: Extended track; historical design, not availability.**

#### 1. Single source of state

**React, CLI, Terraform and Ansible use the same authenticated server services.** A UI-created resource appears in CLI results; a CLI operation appears in the UI. There is no separate synchronization database in the browser.

| Interface | Responsibility |
|---|---|
| UI | Forms, resource views, progress and authorized downloads |
| CLI | Scriptable requests, inspect/wait and useful exit codes |
| Backend | Validation, account policy, durable state and actions |
| Agent | Actual infrastructure convergence and observations |

#### 2. UI screens

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

#### 3. Forms and secret handling

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

#### 4. CLI behavior

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

#### 5. Monitoring model

| Signal | Meaning | Do not infer |
|---|---|---|
| Agent last-seen | Agent communication freshness | App readiness |
| Container/Pod running | Process/runtime state | Correct app response |
| App probe | Defined endpoint readiness | Entire stack/data safety |
| Failed operation | A lifecycle attempt failed | Resource is automatically deleted |
| Host CPU/memory/disk | Measured host load/capacity | Per-tenant permission to see host metrics |
| Backup age/result | Recorded backup status | Restore success |

Poll at a modest configured interval initially; show timestamps and unknown/stale states. Streaming can be added later. Host-wide metrics are admin-only unless a reviewed scoped aggregation is supplied. Alerts should have threshold, duration, destination and a tested trigger.

#### 6. User workflow acceptance

UI create → CLI inspect; CLI stop → UI update; Terraform update → both reflect same desired state; invalid/other-account actions denied in every interface. Test slow provisioning, server errors, double-submit, expiration and secret downloads. Do not duplicate provisioning logic inside React or the CLI.

### Source: Resource Services, former sections 6–8

#### 6. HTTP load balancer

| Input | Validation |
|---|---|
| Backend resources | Owned, allowed app endpoints |
| Listener/rule | Initial bounded HTTP routing behavior |
| Health check | Approved path/protocol and timeouts |
| Public endpoint | Optional registered gateway publication |

CoveStack owns membership/health decisions; the proxy executes traffic forwarding. Generate valid complete include structures, test configuration before atomically reloading, and keep the last good config on failure. Reconcile changed IP/port assignments and removed members. Define empty-pool/unhealthy behavior rather than routing to a default unrelated app.

Acceptance: actual requests reach healthy members, failed/removed backend stops receiving new traffic, invalid reload preserves service, and cross-account/private-service targets are rejected.

#### 7. Auto scaling groups

| Policy | Requirement |
|---|---|
| Minimum/desired/maximum | Ordered bounded values under account quota |
| Member template | Approved image/size/VPC and parent ownership |
| Scale signal | Actual selected metric or explicit desired-size request |
| Cooldown | Bounded rate of decisions and stable window |
| Pending capacity | Count reserved/provisioning members |
| Outage | Pause speculative replacement when agent information is stale |

The controller—not Docker's mere ability to run replicas—decides scaling. Serialize group changes; avoid excess launches during slow image pulls. Retire failed attempts and scale down under defined ownership/data policy. Do not delete persistent data as a hidden consequence of membership replacement.

Acceptance: manual/metric-driven supported change, min/max/quota enforcement, pending counts during slow provision, cooldown, lease outage and owned member cleanup. Record the tested metrics; do not present random chart values as scaling inputs.

#### 8. Snapshots

| Snapshot contains | Snapshot does not contain |
|---|---|
| Selected container filesystem/image state | Mounted database/object volume data |
| Owned immutable image reference | VM memory/kernel state |
| Source/resource/account mapping | Automatically safe reusable credentials |

Create through the documented mapped action, record immutable identity, and restrict image access. Restored instances need their own identity/SSH host-key/credential policy; copying secrets across instances is not harmless. Snapshot retention and garbage collection must respect references and owners.

Acceptance: owned snapshot action completes honestly, restored instance is usable, wrong-account image use is denied, mounted data exclusion is explicit, and unrelated images are not pruned.


### Source: broader template design, Optional future work

The fixed CORE-5 integration has no DAG engine; the broader requirements below are deferred.

#### 3. Template contract

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

#### 4. Stack state and failure policy

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


### Source: inventory/configuration and platform-run automation design

Part 2/inventory is Extended track; unrestricted platform-run jobs are Optional future work, not the install playbook.

#### 4. Ansible scope

| Target | Supported approach |
|---|---|
| Control-plane VM | Operator-reviewed roles with appropriate privileges |
| Container-backed Droplet | Approved image/connectivity; module/API or restricted SSH path |
| Kubernetes workload | Authorized kubeconfig and relevant Kubernetes modules |
| CoveStack records | Authenticated API modules/inventory |

An arbitrary minimal image may lack SSH/Python requirements for ordinary Ansible modules. Validate the chosen image/transport rather than promising every playbook will work. Do not give a user Ansible job the host Docker socket to gain convenient container access.

Start with one playbook that installs/configures a supported application, reports failure clearly, and makes no unintended changes on a second run. Return owned targets and authorized connection details from inventory; secrets are protected references and sensitive tasks suppress output.

#### 5. Platform-run job model, if implemented

| Control | Requirement |
|---|---|
| Execution isolation | Job cannot modify platform host or other account state |
| Inputs | Reviewed templates/schema; no arbitrary executable plugins |
| Identity | Separate scoped job credential |
| State | Per-account locking and protected storage |
| Secrets | Restricted injection, redacted logs, lifecycle cleanup |
| Limits | Time, CPU/memory, process/output/network bounds |
| Outcome | Exit result, durable operation and failed-step detail |
| Cancellation | Defined partial-result/cleanup behavior |

A provider or playbook can execute programs and reach networks. Treat unrestricted uploads as later scope requiring a separate threat model, not moderate configuration work. Never run all account jobs with one global administrator credential.


### Source: main design, former extension/decision sections

CORE-5 network/client decisions are retained in the revised main document; remaining entries move here.

#### 7. Extension boundaries

| Extension | Architecture requirement |
|---|---|
| VPC/IPAM | Durable lease policy plus owned bridge lifecycle |
| Publication | Registered hostname/resource/port mapping; explicit exposure selection |
| Metadata | Trusted workload identity; no cloud/agent secrets |
| Terraform | Provider API client with state refresh/import/lifecycle tests |
| Ansible | Inventory/connection method and approved configuration actions |
| React | Same backend state; no separate resource inventory |
| Monitoring | Actual observations, timestamp and unknown/stale display |

#### 8. Decision records required before code

Record subnet pool, IP allocator authority, proxy location/network namespace, tunnel provider, domain/certificate method, Kubernetes shape, metadata delivery method, retention defaults, and Terraform ownership of stack children. Each decision includes tradeoffs, tested assumptions and acceptance criteria.

This version preserves one-host execution as the initial target deployment. Nested Docker, VM-backed instances and multi-host networking are future/optional work; a container-backed Droplet is not a general-purpose cloud VM.


## 14. Gate evidence and reading links

[Gates](CoveStack-Service-Gates.md) · [E1–E8 acceptance](CoveStack-AntiWrapper-Acceptance.md#appendix-a-extended-track-tests--not-srs-acceptance) · [Layer 2 schema](CoveStack-Database-Schema.md) · [CORE-5 SRS](CoveStack-SRS-Scope.md) · [Go handover](CoveStack-Go-Handover.md).


### Relocated from CoveStack-Service-Catalog.md — Historical scope; extended targets / optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. Extensions agreed for V5

| Request | Feasible bounded version | Important limit |
|---|---|---|
| VPC | Owned private bridge per VPC on one host | Local network simulation |
| IPAM | Non-overlapping subnet leases; recorded actual IPs | One authoritative allocation policy |
| Internet website | Public hostname through tunnel and reverse proxy | Host, tunnel and app must run |
| Simulated public IP | Optional visibly simulated API/UI value | Not internet-routable |
| Direct IP browsing | Real gateway or routed private access | Fake address is insufficient |
| Custom domain | Verified owned name, DNS/TLS and route | Provider-specific setup required |
| Instance metadata | Read-only self-metadata | Identity/routing must be enforced |
| React SPA | Scoped resources, CLI help and real health | Backend still enforces policy |

| Additional capability | Required delivery |
|---|---|
| Stack templates | One-request dependency orchestration and resource references |
| Terraform | CoveStack provider for supported API resources |
| Ansible | Inventory/connectivity plus supported configuration workflows |
| Monitoring | Actual metrics/events and alert evidence |
| Recovery | Safe restart, bounded retry and verified restoration |
| UI/CLI parity | Same backend records, policy and operation results |


### Relocated from CoveStack-Service-Catalog.md — Historical scope; extended targets / optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Publication is standard support, not automatic exposure

Every eligible Droplet/Kubernetes application can use the built-in publication workflow. A shared gateway is installed by the administrator. Users explicitly request publication for an approved service/port; no database, metadata or private management service becomes public merely because its resource exists.

The initial public endpoint is an HTTPS hostname. Arbitrary customer domains are a later gated extension. Kubernetes requests reach a Service/Ingress through the configured cluster ingress path, not a hardcoded Pod IP.


### Relocated from CoveStack-Service-Catalog.md — Historical scope; extended targets / optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 5. Demonstrated release tiers

| Tier | Required outcome |
|---|---|
| Foundation | Real authenticated backend/agent and first compute flow |
| Initial services | Droplet and Kubernetes lifecycle with UI/CLI consistency |
| Network/publication | Owned VPC, address policy and external HTTPS app |
| Full template | App, PostgreSQL, S3 bucket, network and publication |
| Automation/operations | Provider, playbook, monitoring and restore evidence |
| Expanded services | LB, ASG and snapshots with explicit tested limits |

All agreed services remain targets. If a gate fails, report the affected feature unsupported/pending and revise workload with the mentor; do not publish a success-shaped placeholder to preserve a timetable.


### Relocated from CoveStack-Service-Catalog.md — Historical scope; extended targets / optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 7. Optional future features

| Future feature | Reason to defer |
|---|---|
| Nested Docker | Privilege/storage/network complexity |
| VM-backed Droplets | New virtualization driver and operational model |
| Multi-host VPC | Routing, distributed allocation and fencing |
| Additional database engines | Engine-specific persistence/readiness tests |
| Full provider parity | Scope and semantics substantially exceed MVP |

Never mount the platform Docker socket into user workloads as a shortcut to nested Docker.


### Relocated from CoveStack-VPC-IPAM-and-Networking.md — Historical alternatives; no pool decision

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 1. Three separate networking concerns

| Concern | Purpose |
|---|---|
| Private east-west | Permitted app/database/workload communication |
| Outbound egress | Workloads reaching approved outside destinations |
| Public ingress | Explicit published applications reached by visitors |

A VPC is not automatically an internet website. The initial implementation maps an owned VPC onto a user-defined bridge on one host. Shared-kernel containers remain a limited isolation model.


### Relocated from CoveStack-VPC-IPAM-and-Networking.md — Historical alternatives; no pool decision

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. Address allocation

Example pools are illustrative and must be checked against actual routes. Choose a reviewed private pool, such as `10.80.0.0/16`, only when it does not overlap host/VM/LAN/VPN or existing Docker networks. Centrally allocate bounded subnets (for example /24) in the initial routing domain.

| Allocation | Authority | Storage |
|---|---|---|
| VPC subnet | CoveStack transactional allocator | Account/host/CIDR lease |
| Container IP — MVP | Docker IPAM | Inspect and record actual assignment |
| Explicit stable IP — later | One coordinated allocator | Reserved IP plus Docker acceptance |
| Gateway/system addresses | Deployment policy | Reserved before allocation |
| Public hostname | CoveStack domain/route registry | Unique verified mapping |

Do not run two independent workload-IP allocators. Docker already provides IPAM; using it under CoveStack's owned lease policy is legitimate integration. [Docker network options](https://docs.docker.com/reference/cli/docker/network/create/).


### Relocated from CoveStack-VPC-IPAM-and-Networking.md — Historical alternatives; no pool decision

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Lease lifecycle

| Step | Rule |
|---|---|
| Reserve | Atomic uniqueness/locking, not check-then-insert race |
| Provision | Create only the intended owned bridge |
| Record | Save runtime network ID and actual addressing |
| Attach | Validate resource/VPC account and host |
| Recover | Inspect/reconcile pending leases after failure |
| Release | Confirm attachments, ingress and metadata cleanup |
| Reuse | Only after owned objects have actually gone |

Track allocation exhaustion and explicit conflicting-CIDR errors. A partially created bridge/DB lease needs a reconciled failure state; deleting unrelated networks is never a valid recovery shortcut.


### Relocated from CoveStack-VPC-IPAM-and-Networking.md — Historical alternatives; no pool decision

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 4. Isolation and firewall policy

| Path | Initial policy |
|---|---|
| Same VPC approved app → DB | Allowed under selected service policy |
| Different account/VPC | Denied unless explicitly modeled |
| Workload → control-plane DB/cache | Denied |
| Workload → Docker management API/socket | Denied |
| Visitor → published app | Approved gateway only |
| Visitor → metadata/private admin | Denied |

Bridge defaults and published host ports can create alternate paths. Test backend IPs, gateways, published ports and both directions. DNS lookup failure alone is not proof of isolation. Inspect the actual Docker iptables/nftables behavior; do not assume ordinary host INPUT rules govern forwarded container packets.

Kubernetes adds its own service/Pod CIDRs and network policy mechanisms. Reserve those ranges and verify cluster-to-VPC routes. A cluster network does not automatically become an isolated VPC attachment; the driver must define its boundary.


### Relocated from CoveStack-VPC-IPAM-and-Networking.md — Historical alternatives; no pool decision

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 5. Address presentation

| Value | User meaning |
|---|---|
| Private IP | Reachable only through permitted private routes |
| Simulated public address | Explicit emulation value; not real internet reachability |
| Public HTTPS URL | Actually routed visitor endpoint |
| API endpoint/kubeconfig | Separately authorized administrative connectivity |

Documentation-only `203.0.113.42` is a fictitious example, not an address to allocate to internet services. DNS cannot make a fake address reachable. Use a real tunnel/gateway hostname for external visitors.


### Relocated from CoveStack-VPC-IPAM-and-Networking.md — Historical alternatives; no pool decision

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 6. Future constraints

| Request | Additional requirement |
|---|---|
| Overlapping tenant ranges | Stronger namespace/routing isolation |
| Multi-host VPC | Overlay or explicitly managed routing and distributed leases |
| Per-resource public IP | Real address supply and routing; not API fields alone |
| Private remote admin | VPN/client routes and authentication |
| Hostile tenant isolation | Evaluate VM/sandbox boundaries, not bridge names |


### Relocated from CoveStack-API-and-Compatibility.md — Historical extended models / optional broader ownership

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Proposed management models

| Model | Required semantics |
|---|---|
| Resource/action | Stable IDs, pending/completed/failed outcome |
| VPC/attachment | Account/host/CIDR ownership and uniqueness |
| Stack/template | Immutable template version, DAG and child IDs |
| Publication/domain | Verified hostname, approved upstream and lifecycle |
| Metrics/doctor | Scope, timestamps and unknown/stale status |

Endpoint paths for these extensions must be agreed before coding; illustrative names are not an implemented API. Document pagination, timeouts, errors and idempotency for each operation.


### Relocated from CoveStack-API-and-Compatibility.md — Historical extended models / optional broader ownership

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 5. Cross-interface conflicts

Store a durable operation once, then let UI/CLI wait or inspect it. Protect concurrent writes with optimistic/version checks and appropriate idempotency. Define external IaC ownership; do not let UI and Terraform silently fight over the same desired fields.


### Relocated from CoveStack-Stack-Templates.md — Broader templates / customization: Optional future work; publication/UI portions Extended track

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 1. The central user experience

**A stack is an owned environment containing references to individually managed resources.** One request coordinates dependencies, configuration, readiness and publication. Users still inspect each resource's metadata; the stack is not a replacement for its constituent services.

| Template | Resources | Gate |
|---|---|---|
| Single web app | Droplet, VPC, optional HTTPS route | First vertical slice |
| Web/data/storage | Droplet app, PostgreSQL, bucket, VPC, publication | First full environment |
| Kubernetes web/data | Cluster/workload, PostgreSQL, bucket, network, publication | After cluster and full-stack dependencies |
| Customizable environment | Validated supported components and dependencies | After fixed templates work |

Selecting a template is not enough to create an arbitrary application configuration. The app must accept the template's documented environment variables or secret references. Unsupported application assumptions produce an explicit validation error.


### Relocated from CoveStack-Stack-Templates.md — Broader templates / customization: Optional future work; publication/UI portions Extended track

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. End-to-end example

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


### Relocated from CoveStack-Stack-Templates.md — Broader templates / customization: Optional future work; publication/UI portions Extended track

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 5. Single ownership model

UI and CLI read the same server records. Terraform either manages an entire stack or separate child resources: do not let both models concurrently own the same child's desired specification. Ansible configures approved targets; the stack/provider owns resource provisioning. Record external-management ownership and supported adoption/import behavior.


### Relocated from CoveStack-Stack-Templates.md — Broader templates / customization: Optional future work; publication/UI portions Extended track

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 6. Acceptance

Create through UI and inspect through CLI; repeat the same request without duplication; fail a dependency and resume; restart agent/server during provisioning; deny another account; delete with retained data; confirm actual endpoints and secret scoping. Compose already offers multi-service definitions; the evidence must establish CoveStack's controlled self-service and policies.


### Relocated from CoveStack-Terraform-and-Ansible.md — Historical broader automation; Part 2 Extended track, arbitrary runners Optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 1. Two separate provisioning interfaces

| Mode | User experience | Implementation |
|---|---|---|
| Direct Terraform | User writes configuration and runs Terraform | CoveStack provider calls authenticated API |
| Platform template | User selects approved stack/variables | Stack engine; optionally a controlled IaC runner |
| Ansible configuration | User configures owned targets | Scoped inventory/modules/connection |

**Uploading `.tf` or playbook files does not establish support.** Terraform needs a provider, state and resource lifecycle. Ansible needs a working target/connection method and configuration semantics. They already manage Docker through existing integrations; CoveStack adds its account/resource/stack contract.


### Relocated from CoveStack-Terraform-and-Ansible.md — Historical broader automation; Part 2 Extended track, arbitrary runners Optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. Provider slice

| First resource capability | Acceptance |
|---|---|
| Configure endpoint/token | Scoped account identity and TLS |
| Create | Stable ID plus operation wait/cancel timeout |
| Read/refresh | Actual server state and absent-resource handling |
| Update | Valid supported changes; accurate replacement decisions |
| Delete | Idempotent cleanup and retention behavior |
| Import | Adopt an authorized existing resource |
| Drift | Explain difference; no silent duplication |

Use the official [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework) and acceptance tests. Start with one Droplet lifecycle, then VPC/publication and whole-stack support. Pin the provider version after testing.

Illustrative future HCL, **not a published provider schema**:

```hcl
provider "covestack" {
  endpoint = var.covestack_endpoint
}

resource "covestack_stack" "shop" {
  name             = "shop"
  template_id      = "web-postgres-storage"
  template_version = "1"
  publish          = true
}
```

Supply tokens through an approved secret mechanism; do not embed them in examples. The supported resource schema must be implemented before this example can run.


### Relocated from CoveStack-Terraform-and-Ansible.md — Historical broader automation; Part 2 Extended track, arbitrary runners Optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Terraform and controller ownership

| Situation | Rule |
|---|---|
| Stack resource manages children | Do not separately manage same children in another Terraform resource |
| UI modifies Terraform-owned spec | Warn/reject unsupported conflicting writes or document drift policy |
| Controller observes health | Does not overwrite provider's desired configuration |
| ASG-managed members | Do not independently count/manage each generated member |
| Terraform timeout | Refresh durable server result before blindly resubmitting |

A file upload/run interface is a job system, not a provider substitute. Use reviewed templates first. The safest initial implementation lets users run Terraform themselves against APIs; it avoids executing arbitrary user code on the platform host.


### Relocated from CoveStack-Terraform-and-Ansible.md — Historical broader automation; Part 2 Extended track, arbitrary runners Optional future work

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 7. Integration acceptance

Create with Terraform, inspect via CLI/UI, refresh/import, change a supported field and delete under retention policy. Deny another account and test interrupted apply/retry. Run the Ansible example twice, test unreachable/unauthorized targets, and verify logs contain no credentials. Record exact provider/collection versions and real endpoints.


### Relocated from CoveStack-Testing-and-Acceptance.md — Historical broad test matrix; extended comparison scenarios deferred

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Service acceptance

| Area | Required cases |
|---|---|
| Authentication | Active/invalid/revoked/expired credentials; wrong credential kind |
| Signing | Tampered nonce/key/body/path/query; replay and nonce-poisoning |
| Droplet | Create/read/start/stop/delete; intentional stop vs crash |
| Recovery | Restart mid-work; inspect IDs; no duplicates |
| Observations | Stale/future generation, old lease and readiness acknowledgment |
| Kubernetes | Real ready cluster, kubeconfig and deployed test app |
| PostgreSQL | Connection readiness, private access and retained/restored data |
| S3 | Put/get/list/delete on scoped bucket; cross-account deny |
| VPC/IPAM | Concurrent/overlap/exhaustion allocation and denied attachment |
| HTTP LB | Healthy member routing, removal, empty pool and reload failure |
| ASG | Pending reservations, min/max/cooldown and outage pause |
| Snapshot | Owned image restore; volume data explicitly excluded |


### Relocated from CoveStack-Testing-and-Acceptance.md — Historical broad test matrix; extended comparison scenarios deferred

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 4. Extension acceptance

| Area | Required cases |
|---|---|
| Public apps | Droplet and Kubernetes site from external mobile network |
| Domain policy | Unknown/unverified/stale hostname and wrong account |
| Metadata | Self-only identity; no secrets/public ingress; IP reuse cleanup |
| Stack | Dependencies/readiness, retry, partial failure and data retention |
| Terraform | Create/read/update/delete/import, drift and interrupted apply |
| Ansible | Supported configuration twice; denied/unreachable target |
| UI/CLI | Create in either, inspect/mutate in other; identical policy |
| Monitoring | Real timestamps, stale/unknown states and triggered alert |
| Restore | Separate environment with known records/objects verified |


### Relocated from CoveStack-Testing-and-Acceptance.md — Historical broad test matrix; extended comparison scenarios deferred

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 6. Stack acceptance scenario

| Order | Expected result |
|---|---|
| Submit approved template | One durable stack/action |
| Network allocation | Owned non-overlapping bridge |
| DB/bucket readiness | Private DB and scoped storage |
| App deploy/config | References correct service endpoints |
| Publication | Trusted real HTTPS URL |
| CLI inspect UI-created stack | Same child IDs and outcomes |
| Dependency failure/retry | Truthful state and no duplicate children |
| Delete with retain | Data remains protected/recoverable |

A ready stack cannot be backed by success-shaped stubs. Test a failure at each dependency boundary; retain user data even if the app step fails.


### Relocated from CoveStack-Testing-and-Acceptance.md — Historical broad test matrix; extended comparison scenarios deferred

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 7. Fair comparison and mentor evidence

Use the same workload/hardware and comparable security assumptions for Docker/Compose versus CoveStack. Measure workflow effort, account controls, publication setup, restart behavior, overhead and restoration. Credit Docker/Compose for mechanisms already supplied; show the added cloud-style account/template policy.


### Relocated from CoveStack-Testing-and-Acceptance.md — Historical broad test matrix; extended comparison scenarios deferred

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 8. Release checklist

| Release check | Required evidence |
|---|---|
| Supported-feature matrix | Exact endpoints/actions/limits |
| Tests | Required CI and recorded live cases |
| Security | Cross-account denial and private-boundary tests |
| Artifacts | Pinned inputs/checksums and clean install |
| Docs | No fictional commands presented as runnable |
| Data | Retention and restore proof |
| Interfaces | Same API/state and non-sensitive listings |
| Rehearsal | Second-device app and complete stack demonstration |

All unresolved tests are recorded as pending/failed with an owner; “not tested” is not “passed.”


### Relocated from CoveStack-Security-and-ITISS.md — Historical extended threat/evidence rows; CORE-5 rows governed by current guide

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 1. ITISS contribution

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


### Relocated from CoveStack-Security-and-ITISS.md — Historical extended threat/evidence rows; CORE-5 rows governed by current guide

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. Threat model

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


### Relocated from CoveStack-Security-and-ITISS.md — Historical extended threat/evidence rows; CORE-5 rows governed by current guide

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 4. Host and workload authority

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


### Relocated from CoveStack-Security-and-ITISS.md — Historical extended threat/evidence rows; CORE-5 rows governed by current guide

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 5. Secrets and audit

| Material | Rule |
|---|---|
| Agent/tunnel/provider keys | Never public UI/general metadata |
| User DB/S3 credentials | Scoped issuance; separate authorized retrieval |
| Kubeconfig | Protected export and log redaction |
| Terraform state | Sensitive storage/access/backups/locking |
| Audit records | Actor, action, resource, result, request/time; protected writes |
| Evidence | Sanitized actual results; no secret dumps |

Append-only DB permissions reduce accidental alteration but do not justify a tamper-proof audit claim. Record the trust model and administrator authority.


### Relocated from CoveStack-Security-and-ITISS.md — Historical extended threat/evidence rows; CORE-5 rows governed by current guide

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 6. Vulnerability and incident work

Record scanner/tool/input versions, findings, severity, exploitability context, remediation and verified exceptions. Scanner success is not proof of isolation. Revoke a test credential, trigger an alert, investigate logs and restore normal operation in a controlled incident exercise.


### Relocated from CoveStack-Operations-and-Recovery.md — Historical extended operation/recovery scenarios

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 1. Operational principle

**Preserve owned resources and persistent data during uncertainty.** A network outage, incomplete list or malformed work response is not deletion intent. Restart recovery inspects durable IDs and reserved labels before making changes.

| Event | Immediate behavior |
|---|---|
| Server unreachable | Pause new work; retain workloads |
| Agent stops | Mark stale; avoid speculative replacement storms |
| Replay store unavailable | Reject affected authenticated requests safely |
| App fails | Report health/error; honor desired state |
| Tunnel fails | Public endpoint unavailable; internal resources remain |
| Proxy config invalid | Keep last valid configuration; report failure |


### Relocated from CoveStack-Operations-and-Recovery.md — Historical extended operation/recovery scenarios

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. Routine operations

| Procedure | Required checks |
|---|---|
| Startup | Secrets, versions, private listeners, schema and dependency health |
| Deployment/update | Reviewed image/config changes; health and compatibility |
| Shutdown | Drain/stop supported work; preserve volumes |
| Key rotation | Scoped replacement, revocation and tested rollback |
| Certificate renewal | Valid certificate plus tested proxy reload |
| Resource deletion | Owned dependency cleanup and explicit retention |

Do not use `docker compose down -v` or blanket prune as a default repair. A workload's desired off state is different from an active workload crashing. Recovery must honor that distinction.


### Relocated from CoveStack-Operations-and-Recovery.md — Historical extended operation/recovery scenarios

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Backup coverage

| Data | Backup/recovery need |
|---|---|
| Control-plane PostgreSQL | Accounts, resource/action/stack records and ownership |
| Workload PostgreSQL | Database-consistent data backup |
| Object storage | Objects plus policies/ownership references |
| Master secrets | Protected key recovery distinct from DB backup |
| Runtime config | Proxy/domain/network/version records |
| Kubernetes | Defined workload/data recovery; kubeconfig alone is not backup |

A container filesystem snapshot excludes mounted-volume data and is not a database backup or VM memory snapshot. Record precisely what each backup/snapshot contains.


### Relocated from CoveStack-Operations-and-Recovery.md — Historical extended operation/recovery scenarios

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 5. Idempotency and fencing

| Scenario | Rule |
|---|---|
| Request response lost | Refresh durable operation before resubmitting |
| Agent retry | Fresh signed request; mutation semantics agreed |
| Worker duplicate | Per-resource serialization and durable identity |
| Old agent resumes | Lease/epoch check prevents stale authority |
| Partial stack failure | Resume safe steps; do not erase persistent data |
| IP/hostname reuse | Wait until stale route/metadata attachment removed |


### Relocated from CoveStack-Operations-and-Recovery.md — Historical extended operation/recovery scenarios

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 7. Release evidence

Clean VM install, remote browser app, approved private admin connectivity, restart test, credential revocation, alert trigger and separate restoration all need actual results. Store exact versions/hardware and identify cases not tested. Use the [acceptance matrix](CoveStack-Testing-and-Acceptance.md).


### Relocated from CoveStack-Current-Setup-Fixes.md — Historical broad audit extension targets

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. V5 extension findings

| Extension | Required review |
|---|---|
| VPC/IPAM | Overlap/exhaustion/atomic allocation and safe release |
| Gateway/domains | Owned upstream, claim uniqueness and stale route cleanup |
| Metadata | Trusted identity, secret exclusion and namespace route |
| Stack engine | Dependency DAG, idempotency, failure and retention |
| Terraform/Ansible | State, connectivity, ownership and execution isolation |
| React/monitoring | API parity, permissions and truthful freshness |

Use the implementation-status and acceptance documents to assign current owners/results. Do not treat historical proposed fix wording as an instruction to undo an already corrected baseline.


### Relocated from CoveStack-References-and-Decision-Log.md — Historical external links/candidates; no offline verification or selected tools

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 2. Primary source links

| Reference | Purpose |
|---|---|
| [Docker bridge](https://docs.docker.com/engine/network/drivers/bridge/) | Local network primitives and limits |
| [Docker IPAM options](https://docs.docker.com/reference/cli/docker/network/create/) | Existing address mechanisms |
| [Docker security](https://docs.docker.com/engine/security/) | Daemon/workload authority |
| [Cloudflare routing](https://developers.cloudflare.com/tunnel/concepts/routing/) | Candidate public hostname transport |
| [Domain validation](https://developers.cloudflare.com/cloudflare-for-platforms/cloudflare-for-saas/domain-support/hostname-validation/) | Ownership/certificate distinction |
| [DigitalOcean metadata](https://docs.digitalocean.com/products/droplets/how-to/access-metadata/) | Selected provider endpoint semantics |
| [AWS metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html) | Separate IMDSv2 behavior |
| [k3d exposure](https://k3d.io/stable/usage/exposing_services/) | Ingress integration |
| [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework) | Provider development |
| [Ansible Docker](https://docs.ansible.com/projects/ansible/latest/collections/community/docker/index.html) | Existing automation capabilities |
| [OWASP object permissions](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/) | Authorization test rationale |

These mechanisms were researched during the preceding design work; k3d, Terraform and Ansible references were checked again while preparing V5. Exact tool/provider compatibility must be retested against release versions. Do not claim Docker lacks IPAM or Terraform/Ansible integrations.


### Relocated from CoveStack-References-and-Decision-Log.md — Historical external links/candidates; no offline verification or selected tools

**Historical target text; not a capability claim.** Scope/owners/dates in sections 1–12 override this text. Old sample names, states, schemas, tool candidates and estimates are not approved wire strings, version pins or decisions; missing values remain TBD with B.

#### 3. Decisions to record in G0

| Decision | Required information |
|---|---|
| First stack | Supported app interface and dependencies |
| Network pool | Host/VM/VPN overlap evidence |
| Proxy/tunnel | Namespace placement, DNS/TLS and provider terms |
| Kubernetes | Tested shape/version, ingress and admin access |
| Metadata | Explicit endpoint first versus link-local milestone |
| Resource ownership | UI/CLI/IaC conflict and adoption policy |
| Persistent-data default | Retain/delete/restore semantics |
| Team capacity | Named frontend/provider/Ansible owners |


### Batch 3 relocation and conflict resolution

| Prior design assumption | Approved replacement |
|---|---|
| Seven original resource services as release scope | CORE-5 only; LB/ASG/filesystem capture move to Extended track; VPC/IPAM is CORE-5 |
| Broad UI/templates/public-route/provider parity before submission | CLI first, one fixed stack, optional Droplet/K8s provider; wider work deferred |
| Arbitrary/illustrative pool and multiple networks | Pool TBD B+A; one private network/account; no custom CIDR |
| Early DNS/router/LAN and external-device demo dependency | Single AlmaLinux 8 GB VM on phone hotspot; no router DNS/DHCP; CI/video backup |
| Agent bucket reconcile or separate network driver | Java bucket/credential reconcile; MinIO process on generic container driver; network helper, two drivers |
| Unbounded cluster shape/version and assumed memory | Required cap 2; 1 server + max 2 agents; tested pair; memory estimate unverified |
| Broad Go/frontend ownership | A all agent Go; E thin CLI/later provider; D Java only; C acceptance/allocator |

These changes are explicit scope/ownership decisions, not deletion of requirements or evidence of completion. Historical sections above preserve earlier requirements, sample values and candidate tools; only sections 1–12 and approved CORE-5 documents govern new implementation. No unsupported sample name/state/version becomes an approved contract string.
