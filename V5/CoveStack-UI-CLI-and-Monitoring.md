# CoveStack V5 — CLI, Web Interface and Freshness-Aware Monitoring

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. One authoritative state and clear maturity

CLI/API is the first target interface; an eventual React SPA uses the same authenticated APIs. There is no current public CLI/dashboard implementation established by source. UI is not authorization or a second database. Basic monitoring/health/logging/audit are core foundations; rich dashboards/tracing/prediction/alert rules are later extensions, not reasons to skip current diagnostics.

## 2. CLI experience

Go/cobra thin client; endpoint/profile/key identifier config only, secrets protected separately. Mutations carry idempotency key. Create returns pending with durable ID; finite wait polls actual desired-versus-converged generation and fails truthfully on timeout/failure. All list/get offer stable JSON with scoped IDs, observation age and operation details. No runtime socket, platform DB or agent signing key for ordinary user access. Actual output schema/command names require API/module agreement; do not publish working commands before implementation.

## 3. UI screens and forms — future interface

| Screen | Target information / action | Security boundary |
| --- | --- | --- |
| Read-only list first | Scoped resources, operation and last observed freshness | Account/project authorization remains server-side |
| Resource detail | Desired/converged state, owned dependencies, host/profile limits, sanitized events/logs | No ordinary secret/kubeconfig/admin endpoint disclosure |
| Create form | Validated VM/container/cluster/network/database/storage options returned by capabilities | No unsupported backend knobs, host paths, arbitrary command or fabricated size choices |
| Environment | Typed child readiness/failure and safe retry/retention | Shared resource reference is not permission to delete its owner |
| Recovery/keys | Explicit authorized audited action and required warning about target/data effects | Secret export separate, no persistence in URLs, ordinary logs or browser storage |

Start UI only after real authenticated list/get and truthful health exist. Form completeness does not establish resource support. Public application exposure is a separate future opt-in, never an automatic effect of a create form.

## 4. Metrics and health model

| Signal | Authority / semantics |
| --- | --- |
| Desired revision | DB intent, not evidence the VM/service is healthy |
| Convergence revision | Accepted successful readiness for a particular desired revision, not perpetual health |
| Current health | Timestamped/ordered probe/provider observation with freshness budget |
| Host capacity | Fresh compatible host observations plus durable reservations, never usage alone |
| Logs/events | Scoped sanitized correlation of user/account/resource/operation/attempt/provider task |
| Alerts | Policy-driven later delivery with durable event; no alerts already running is claimed |

Agent loss makes health stale/unknown rather than healthy; observation payload labels are not invented wire enums. Reject stale/future/out-of-order reports per approved protocol. Redact credentials, query strings, private configs and backend responses; retention/tenant logging access and volume limits require review.

## 5. Core versus advanced observability

Basic host CPU/RAM/disk pressure, agent reachability, resource/service readiness, errors, operation events and useful logs belong with the lifecycle being built. Advanced dashboards, historical utilization, prediction, distributed tracing, richer alerting/notifications and external integrations need measured demand/storage/retention and secure provider choices. No collector/metrics database/alert tool is selected just by drawing a chart.

## 6. Acceptance

Same account sees same real resource/operation across CLI/API/future UI; wrong account denied. JSON parsable and redacted; mutating retries one ID; create/wait never turns server failure into success. Stop agent and verify actual freshness changes (E7), retained stale report cannot make it fresh, and useful event/log errors contain no secrets. UI/server independently validate input and permission. No interface bypasses G1–G4/profile gates.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-ui-cli-and-monitoring) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
