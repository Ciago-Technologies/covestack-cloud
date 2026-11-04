# CoveStack V5 — Implementation Evidence and Missing Work

| Document control | Value |
| --- | --- |
| Updated | 4 November 2026 |
| Authority | Detailed current product specification; aligned to the [canonical catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) and [product SRS](../docs/product/CoveStack-Product-SRS.md) |
| Evidence | Target design, not a completion certificate; source and actual test results determine implementation status |
| Compatibility | V5 is a documentation label; current wire contract remains 1.1.0; future fields/kinds/versions require coordinated review |

## 1. Evidence policy

Snapshot: 4 November 2026. Source inspected read-only; no application/compiler/unit/runtime/infrastructure acceptance tests ran for this V5 alignment. Historical reports keep dates/scope; table categories describe actual component evidence, not service support. Product is Experimental and no supported production profile is established.

## 2. Source-confirmed status register

| Feature | Current category | Source evidence | Evidence still required |
| --- | --- | --- | --- |
| Agent config/signing | Implemented — confirmed in code | config.Load, sigv1 rules/request/sign and test sources | Deployment credential/key/clock/TLS/replay integration |
| Signed work client | Implemented — confirmed in code | internal/client.New/GetWork, status/error/response bounds | Real Java assigned-account response and secret-handling review |
| Single-account poll/startup | Implemented — confirmed in code | internal/agent.Run and cmd/kan-agent/main.go; safely stops on returned work | Actual Java outage/recovery/auth-stop/signal evidence |
| Lifecycle validation | Implemented — confirmed in code | Generated 1.1.0 catalogue and Go/Java validation helpers | VM/container schema compatibility; real service readiness enforcement |
| Java auth/replay components | Partially Implemented | Filters/interfaces, ValkeyNonceStore and auth configuration | Concrete persistent KeyService/BearerTokenService, host assignment, runnable app/controllers |
| Java control plane / DB foundation | Blocked | pom declares Java 25/Spring Boot 3.5.16/JDBC/Flyway; Boot repackage skip true | Entry point/configuration, scoped APIs, prerequisite migrations and empty DB acceptance |
| SQL files | Partially Implemented | V5 publications, V6 domains exist | Not a working DB chain; accounts/resources prerequisites absent; no applied-history inspection |
| VM Instance / Incus | Planned | New design decision only; no driver in source | Host spike, schemas, secure API driver and actual guest-kernel/failure acceptance |
| App Container | Planned | Separate resource model only | Explicit resource/spec, restricted driver and trusted-workload acceptance |
| Kubernetes | Planned | Contract/example kinds only | Real provider, capacity, scoped export, cleanup/workload evidence |
| PostgreSQL / object-storage services | Planned | Target docs; platform dependency Compose is not a tenant service | Drivers, ownership/policy, data path/restore and secret custody |
| VPC/IPAM | Planned | Prior reference design, no current VPC contract kind | VM/container network protocol/allocator/enforcement and collisions/traffic tests |
| Capacity scheduler / reservations | Planned | Architecture/benchmark plan | Transactional admission, fresh host reports, exhaustion tests |
| CLI / Terraform | Planned | No current CLI/provider module; Terraform deferred in environment | Reviewed module/API/auth boundary and lifecycle/JSON tests |
| Backups/audit/installer/monitoring | Planned | Documentation and acceptance intent | Persisted implementation, real restore/privileges/clean install/freshness |
| Optional Floci | Proposed | Product integration proposal only | Pinned external artifact/license/capability model, network/auth/cleanup tests |
| Customer advantage | Proposed | User hypotheses and competitor research | Interviews, pilots and comparison of actual recurring workflows |

## 3. What the agent actually does

cmd/kan-agent loads environment config and creates signed bounded HTTP client/loop under signal cancellation. GET work returning [] is no work. Any returned resource work stops with ErrWorkAvailable: no claim, execution, acknowledgement, deletion or reconciliation. Retry covers selected HTTP statuses/timeouts, not all transport failures; real Java outage/assignment/credential and multi-account independence need evidence. Current principal fields do not establish trusted host assignment.

## 4. Product and backend decisions

VM Instance must be a true guest kernel; App Container separate trusted shared-kernel type. Incus/QEMU/KVM preferred evaluation; no driver exists. RustFS preferred Planned operator-managed service with Java account-owned bucket/key/policy reconcile, not MinIO selection or an implemented adapter. SeaweedFS/alternatives remain possible. Current wire 1.1.0 unchanged; future VM/type/claim/capacity/network schema is reviewed before execution.

Core catalogue includes more than runtimes: IAM/secrets/quotas/observability/logging/audit/recovery/interfaces and typed environments. Historical service counts/caps/dates do not limit product. Component tests do not demonstrate every foundation.

## 5. Existing reports and environmental limits

Earlier Go race/vet/lint, Java tests and contract/dependency reports remain [dated toolchain evidence](../docs/development-toolchain-status.md), not fresh service acceptance. Manifest Go toolchain/CI pin was 1.27.2 while a historical environment reported 1.27.3; resolve availability/pin review separately, no silent upgrade. Java 25/Spring/JDBC/Flyway are actual source inputs. No live DB/applied migrations, runtime/socket/Incus/RustFS capability, remote CI/GitHub task or production host profile was established.

Preserve current restricted sandbox policy: no Docker/Compose/k3d/S3 acceptance or new manifests, Terraform deferred and GitHub CLI unauthenticated in the recorded environment. VM host not probed/authorized here; documentation does not lift restrictions.

## 6. First integration and status updates

B+C real Java foundation/credential provisioning/trusted assignment, D API/fixtures and A existing signed no-execution loop are the first shared checkpoint. E reviews thin client/module and learns with labelled mocks. Record issue/revision, commands/results, mock versus real scope, blocker/owner and next action. No guessed completion percentage or blanket all-services-complete claim. [Team checkpoint](../TEAM-FIRST-CHECKPOINT.md), [gates](CoveStack-Service-Gates.md), [fixes](CoveStack-Current-Setup-Fixes.md).

## 7. License and distribution

Root LICENSE contains maintainer-approved official Apache License 2.0, SPDX Apache-2.0, for CoveStack-owned code/docs. Upstream licenses/notices, resolved dependency/image distribution and copied-reference provenance remain separate; no invented copyright holder or implied security certification. No runtime support follows from licensing.

## Historical traceability

The [previous body of this file](../docs/archive/V5-before-full-product-alignment-2026-11-04.md#v5-covestack-implementation-status) is preserved unchanged as historical reference, including earlier requirements/DDL/acceptance text. It is not current VM/RustFS setup guidance. Use this active body, the linked canonical domain specifications and [actual implementation evidence](CoveStack-Implementation-Status.md) for new work. No source, migrations, contracts, tests, configuration or runtime state is changed by this document.
