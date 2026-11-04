# CoveStack platform documentation

Updated: 4 November 2026. **Experimental; no supported deployment profile.** Documentation specifies targets, not runtime acceptance.

CoveStack targets **an open-source, local-first infrastructure control plane for small teams that want account-owned environments, intelligent resource scheduling, reliable provisioning, built-in security and recoverable infrastructure across their own machines**. Open-source distribution rights and multi-host support are not yet established. Primary compute is an actual **VM Instance**; an **App Container** is a separate shared-kernel resource. Incus is the preferred VM backend to evaluate, using QEMU/KVM on a qualifying Linux host. No such driver exists yet.

Users should request infrastructure, observe lifecycle and recover interrupted operations without manually administering every runtime object. Demand and comparative advantage remain hypotheses. This is not an AWS, Floci, LocalStack, Coolify, OpenStack or CloudStack replacement.

## Reading paths

| Reader | Start | Continue |
| --- | --- | --- |
| New user | [Product strategy](strategy.md) | [Actual implementation](implementation-status.md), [installation limits](installation-and-operations.md) |
| Architect | [Target architecture](architecture.md) | [VM/container model](resource-model.md), [agent/reconcile](agent-and-reconciliation.md), [capacity](capacity.md) |
| Security/operator | [Threat model](security.md) | [Network/storage](network-and-storage.md), [operations](installation-and-operations.md), [readiness](readiness.md) |
| Implementer | [API/CLI boundaries](api-and-cli.md) | [Feature specifications](feature-specifications.md), [decision register](decisions.md), [contribution guide](contributing.md) |
| Team/mentor | [Academic transition](academic-transition.md) | [Core/future dependencies](future-services.md), [acceptance](readiness.md) |
| Product reviewer | [Competitive analysis](competitive-analysis.md) | [Validation plan](strategy.md), [license/attribution](license-and-attribution.md) |
| Documentation maintainer | [Repository audit](../documentation-audit.md) | [References](references.md), [canonical AI checkpoint](../../checkpoint.md) |

## Authority and compatibility

1. Latest explicit user decisions select the VM-based target; no mentor scope waiver or implementation acceptance is inferred.
2. [Contract 1.1.0](../../contracts/README.md) remains the actual wire contract. It has no approved VM/App Container separation or full claim/renew protocol. Product terminology is not a new wire kind.
3. The [canonical product architecture and service catalogue](../product/CoveStack-Product-Architecture-and-Service-Catalogue.md) owns identity and core/future classification; this directory supplies domain specifications. [V5](../../V5/README.md) preserves the previous academic design, schema reference and tests; its container-Droplet/two-driver/8 GB production implications are superseded, not evidence to reproduce.
4. [Academic obligations](academic-transition.md) remain recorded. A revised SRS/backend mapping needs B+C+A/D and mentor review before implementation changes; do not quietly drop Kubernetes or tests 1–15.
5. Historical logs, test reports and learning snapshots retain their dates. [Status](implementation-status.md) separates source evidence, historical checks and absent runtime proof.

No current application-environment engine, VM installer, runnable product API, market validation, license grant or production certification is established by these documents. The repository lacks a root license file; distribution policy needs maintainer approval.
