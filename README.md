# CoveStack

**An open-source, local-first infrastructure control plane for small teams that want account-owned environments, intelligent resource scheduling, reliable provisioning, built-in security and recoverable infrastructure across their own machines.**

This is the product target, not a statement of shipped support or a license grant. Read the [canonical product architecture and service catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md).

CoveStack aims to turn a resource request into actual infrastructure, visible lifecycle and safe recovery—without making each user administer the underlying runtimes. Primary compute is a **VM Instance with its own guest kernel**. **App Containers** are a separate, shared-kernel resource for trusted workloads. Incus is the preferred VM backend to evaluate; no VM driver is implemented yet.

## Status: Experimental

| Present in source                                                                                                   | Not yet available as a verified product                                                                            |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------ |
| Go configuration, HMAC request signing, bounded signed client and cancellable single-account no-execution poll loop | Actual VM/App Container provisioning, runtime reconciliation and capacity admission                                |
| Java auth/replay components and lifecycle validators                                                                | Complete runnable account-owned resource API, concrete credential provisioning and foundational DB migration chain |
| Unit-test sources and dated environment reports                                                                     | Supported install/profile, real isolation, verified workload restore and production readiness                      |

The loop deliberately stops if work arrives; it cannot execute resources. Tests/builds do not prove infrastructure readiness. No application/infrastructure checks ran for this documentation update. See [implementation evidence](docs/platform/implementation-status.md).

## Who it is for

Potential users are small teams, organizations, DevOps engineers and self-hosters who need owned VM/application/data environments, predictable cleanup and recoverable operations. This is a **market hypothesis**, not evidence of customers or adoption. [Use cases and validation](docs/platform/strategy.md).

## Why another infrastructure tool?

Floci/LocalStack primarily serve cloud-API development/testing and can run real selected workloads. Coolify/Dokku/Portainer already offer useful app/container management. Incus, CloudStack and OpenStack already manage actual infrastructure and tenancy. CoveStack's hypothesis is a smaller integrated account-owned VM/app/data/recovery workflow—not uniquely inventing those features or replacing those tools. Floci can be an optional testing integration, never a mandatory core or a Java rewrite target. [Fair comparison and evidence to prove value](docs/platform/competitive-analysis.md).

## Architecture and limits

Java/Spring Boot control plane → authoritative PostgreSQL intent/operations → authenticated outbound Go agent → future VM/container/Kubernetes and network/storage providers. Valkey supports replay protection; the object-storage provider is under evaluation (Garage, SeaweedFS and Ceph RGW; legacy MinIO requires maintenance/licensing review); Mailpit is development email capture. API/CLI users never access privileged sockets or the platform DB.

The current wire contract remains 1.1.0; VM/container and claim/capacity/network evolution needs coordinated schemas and tests. Previous 8 GB sizing is a lab constraint, not a production requirement. Single-host operation is not HA. No production host profile or root open-source license grant has been established.

## Start here

- [Current documentation index](docs/platform/README.md)
- [VM/container decision](docs/platform/resource-model.md) · [architecture](docs/platform/architecture.md)
- [Security](docs/platform/security.md) · [capacity](docs/platform/capacity.md) · [readiness](docs/platform/readiness.md)
- [Installation limits and operations](docs/platform/installation-and-operations.md) · [API/CLI boundary](docs/platform/api-and-cli.md)
- [Core/future catalogue](docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md) · [Future dependencies](docs/platform/future-services.md) · [academic scope/deadline transition](docs/platform/academic-transition.md)

No full product quick-start is published because it would imply functionality not yet present. Existing Compose is dependency configuration only; do not assume it starts CoveStack or a VM.

## Contribute

Choose one bounded issue, read [contribution guidance](docs/platform/contributing.md), [decision gaps](docs/platform/decisions.md), [documentation audit](docs/documentation-audit.md) and [checkpoint.md](checkpoint.md). Help validate use cases, VM backend capabilities, compatible protocol design, security and recovery tests. Java remains the control plane; Go grows in product-owned components selectively, not through an unsolicited rewrite.

**Licensing:** no root LICENSE was found; maintainers must decide and publish a real grant before promising open-source redistribution. Third-party artifacts have their own obligations. See [license/attribution review](docs/platform/license-and-attribution.md). No feature badges, availability guarantees or security certification are claimed.

V5, old guides and exports are explicitly historical design/evidence. Their academic requirements remain recorded; new VM mapping and deadline reconciliation require team/mentor review, not silent removal.
