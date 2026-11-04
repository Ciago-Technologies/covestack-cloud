# Implementation and evidence ledger

Product identity and core/future scope: [canonical architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

Snapshot: 4 November 2026. Source inspected; **no application, compiler, unit, infrastructure or acceptance tests executed in this documentation task**. Historical reports retain their original scope. Product maturity Experimental; supported production profiles: none.

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

No feature here is promoted to Runtime Verified, Acceptance Gates Passed or Supported. Full service support requires G1–G4 plus declared profile qualification. Reading a test function is not a fresh test run.

## Historical evidence retained

[Toolchain status](../development-toolchain-status.md) records earlier Go race/vet/lint, Java 15-test builds, contract checks and dependency availability in an OmniRush sandbox. [Environment record](../development-environment.md) records its specific versions. Do not silently replace those dates/results with today's claims. They are environment evidence, not proof the resource API/VM runtime exists.

Current manifest Go toolchain/CI pin is 1.27.2 while that environment reported 1.27.3; align/review later without changing manifests here. Java source remains JDBC, not JPA. No root license file found. No live credentials, DB schema, upstream sync, GitHub PR acceptance or production runner was inspected.

## Environment boundaries

Preserve user-directed restrictions: no Docker daemon/socket provision, no Compose runtime/k3d clusters/container-backed acceptance in this sandbox, no Kubernetes manifest creation, no configured MinIO/S3 endpoint, Terraform deferred and GitHub CLI unauthenticated. VM/KVM/Incus runtime was not probed or authorized either. New architecture requires a separately authorized capable host; documentation does not lift these restrictions.

Next meaningful evidence: B+C real server/assignment/credential provisioning for existing signed no-execution loop; D/B/A/C VM/claim/network compatibility ADR; authorized backend spike elsewhere. Check current worktree and [checkpoint](../../checkpoint.md), not old next-task notes.

Current loop retry evidence is bounded: timeout-class transient errors and selected HTTP responses are handled; not every network error is retried. Connection-refusal/general transport recovery and multi-account independence still require review and real-server evidence. Returned resource work stops with ErrWorkAvailable; no driver execution occurs.
