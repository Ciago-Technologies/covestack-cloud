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

Source files remain unchanged until Batch 3; this appendix is the relocation destination. It preserves design intent, not old claims of availability. CORE-5 topics appearing as prerequisites here remain governed by the main design and gate register.

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

#### 1. Default supported capability

**Publication support is built in for eligible Droplets and Kubernetes applications. Exposure is explicit.** One shared tunnel/gateway can carry many hostname routes; each published app receives a validated owned mapping. This is not one new tunnel per container.

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

This version preserves one-host execution as the initial supported deployment. Nested Docker, VM-backed instances and multi-host networking are future/optional work; a container-backed Droplet is not a general-purpose cloud VM.


## 14. Gate evidence and reading links

[Gates](CoveStack-Service-Gates.md) · [E1–E8 acceptance](CoveStack-AntiWrapper-Acceptance.md#appendix-a-extended-track-tests--not-srs-acceptance) · [Layer 2 schema](CoveStack-Database-Schema.md) · [CORE-5 SRS](CoveStack-SRS-Scope.md) · [Go handover](CoveStack-Go-Handover.md).
