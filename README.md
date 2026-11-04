# CoveStack

**A local-first infrastructure control plane for account-owned environments on machines you control.**

CoveStack aims to combine real infrastructure provisioning, capacity-aware scheduling, crash-safe reconciliation, built-in security and verified recovery through one API and CLI. **CoveStack-owned code and documentation are licensed under Apache License 2.0 (SPDX: `Apache-2.0`); see [LICENSE](LICENSE).**

Primary compute is a **VM Instance with its own guest kernel**. **App Containers** are distinct shared-kernel resources for trusted workloads. Incus/QEMU/KVM is the preferred VM backend to evaluate; no VM driver exists yet.

## Status: Experimental

| Present in source | Still pending |
| --- | --- |
| Go configuration, HMAC signing and bounded signed HTTP client | Full credential provisioning and trusted account/host assignment |
| Cancellable single-account signed polling with deliberate no-execution behavior | Claims, runtime drivers and crash-safe resource execution |
| Java authentication/replay components and lifecycle validators | Runnable scoped resource APIs and foundational migration chain |
| Unit-test sources and dated reports | Real isolation, verified recovery and a supported deployment profile |

The agent stops safely when work arrives. It does not claim, execute, acknowledge or delete that work. Source and mock tests do not prove live-server or infrastructure readiness. See the [implementation ledger](docs/platform/implementation-status.md).

## Product scope

The [canonical architecture and service catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) defines proposed core and future infrastructure capabilities independently of historical academic CORE-5. Core includes compute, networking, storage, PostgreSQL, environments and security/operations foundations. Future candidates remain dependency-bound targets, not available services or a release promise.

Potential users are small teams, DevOps engineers, organizations and self-hosters. Demand and comparative advantage require interviews and pilots. The hypothesis is a coherent account/environment lifecycle and recovery experience—not simply more services. Read the [competitive analysis](docs/platform/competitive-analysis.md).

## Architecture

Java/Spring Boot control plane → authoritative PostgreSQL → authenticated outbound Go agent → future VM, container, Kubernetes and networking/storage providers. Valkey is the replay store; Mailpit is development email capture. Users never receive privileged sockets or direct platform database access.

**RustFS is the preferred object-storage candidate; integration is Planned.** The operator manages the storage service and persistent volumes; Java reconciles account-owned buckets, policies and scoped credentials. The provider adapter retains SeaweedFS and other alternatives. [RustFS compatibility gaps and required tests](docs/platform/provider-evaluation.md#rustfs-compatibility-and-acceptance).

The current wire contract remains 1.1.0. VM/type, trusted assignment, claim and capacity evolution requires coordinated schema/fixture review. Floci is an optional external testing integration. Historical 8 GB sizing is not a production recommendation; single-host deployment is not HA.

## Documentation

- [Product SRS](docs/product/CoveStack-Product-SRS.md) and [service catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md)
- [Architecture and HLD/LLD](docs/platform/architecture.md), [VM/container model](docs/platform/resource-model.md) and [environments](docs/platform/environments.md)
- [Security](docs/platform/security.md), [capacity](docs/platform/capacity.md) and [acceptance/readiness](docs/platform/readiness.md)
- [Installation and operations](docs/platform/installation-and-operations.md), [API/CLI](docs/platform/api-and-cli.md) and [documentation index](docs/platform/README.md)
- [Future dependencies](docs/platform/future-services.md) and [historical academic reference](docs/platform/academic-transition.md)

No full-product quick-start is offered yet: existing Compose starts dependencies, not the complete platform or VM service.

## Contributing and licensing

Read the [contribution guide](docs/platform/contributing.md), [open decisions](docs/platform/decisions.md) and [checkpoint](checkpoint.md). Java remains the control plane; Go components grow selectively without a wholesale rewrite.

**License: Apache License 2.0 — SPDX `Apache-2.0`.** The maintainer approved adoption; the complete text is in [LICENSE](LICENSE). Existing third-party material, dependencies and external projects retain their own licenses and notices. See the [licensing decision](docs/platform/licensing-decision.md) and [attribution review](docs/platform/license-and-attribution.md) for outstanding distribution checks. License adoption does not establish production support.

Historical V5 guides and evidence retain their original statements under supersession notices. They do not constrain product scope or establish current functionality.
