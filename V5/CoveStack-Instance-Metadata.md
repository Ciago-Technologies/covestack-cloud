# Architecture transition notice — 4 November 2026

> **Superseded design reference.** Active guidance: [platform README](../docs/platform/README.md). Old container-Droplet architecture, two-driver ceiling, fixed sizing and delivery claims below are historical. The active platform docs define VM compute, a distinct App Container and evidence-based releases.
>
> The original V5 body is retained below without changing its requirements, DDL, test numbers or historical evidence. It must not be read as the current VM architecture or supported setup. [Academic transition](../docs/platform/academic-transition.md) records unresolved SRS/deadline decisions; [contract 1.1.0](../contracts/README.md) remains unchanged. No VM backend or production profile is implemented by this notice.

Current product scope: [canonical architecture and service catalogue](../docs/product/CoveStack-Product-Architecture-and-Service-Catalogue.md); wider proposed core is not an approved replacement for academic CORE-5.

<!-- Historical body preserved below; see supersession notice above. -->

# CoveStack V5 — Instance Metadata and Identity Boundaries

| Document control        | Value                                                                   |
| ----------------------- | ----------------------------------------------------------------------- |
| Product / documentation | **CoveStack / V5**                                                      |
| Updated                 | **14 October 2026**                                                     |
| Submission deadline     | **Monday, 25 January 2027**                                             |
| Compatibility names     | Server: `kan-server`; agent: `kan-agent`                                |
| Status                  | Target design and implementation guidance; not a completion certificate |

## 1. Two meanings of metadata

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#6-metadata). CORE-5 has priority; no capability is claimed here.

## 2. Provider contract

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#6-metadata). CORE-5 has priority; no capability is claimed here.

## 3. Bounded first response

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#6-metadata). CORE-5 has priority; no capability is claimed here.

## 4. Delivery stages

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#6-metadata). CORE-5 has priority; no capability is claimed here.

## 5. Caller identity

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#6-metadata). CORE-5 has priority; no capability is claimed here.

## 6. Tests and operational policy

**Extended track — not in the SRS or deliverables.** Design details and acceptance are relocated to the [extended roadmap](CoveStack-Extended-Roadmap.md#6-metadata). CORE-5 has priority; no capability is claimed here.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
