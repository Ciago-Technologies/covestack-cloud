# CoveStack V5 — Documentation Index and Reading Guide

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Start here

**Submission: Monday, 25 January 2027.** Product spelling is **CoveStack**; existing server/agent modules remain **kan-server** and **kan-agent**. V5 replaces the older planning timetable, but does not change running code or the current wire-contract version.

This directory contains the reviewed CORE-5 plan, acceptance/schema references and a separate deferred roadmap. Planned services and illustrative commands are explicitly separated from implementation evidence.

## 2. Requested file replacements

| Previous guide/topic | V5 replacement |
|---|---|
| DocsV4-DO_Team / technical design | [DocsV5-DO_Team](DocsV5-DO_Team.md) |
| Folder structure explanation | [Folder map](Folder_Structure_Explanation.md) |
| Complete working setup | [Setup](CoveStack-Complete-Working-Setup.md) |
| Project workflow | [Workflow](CoveStack-Project-Workflow.md) |
| Team Roles | [Ownership](Team_Roles.md) |
| tools | [Tools](tools.md) |

## 3. Complete document index

| File | Purpose |
|---|---|
| [CoveStack-SRS-Scope.md](CoveStack-SRS-Scope.md) | CORE-5 scope and five deliverables; examiner reading |
| [CoveStack-Service-Gates.md](CoveStack-Service-Gates.md) | Promotion rule and evidence register |
| [CoveStack-AntiWrapper-Acceptance.md](CoveStack-AntiWrapper-Acceptance.md) | Tests 1–15, CORE-5 viva; separate extended appendix |
| [CoveStack-Database-Schema.md](CoveStack-Database-Schema.md) | Reference schema, migration order and open decisions; not executable |
| [CoveStack-Go-Handover.md](CoveStack-Go-Handover.md) | A/E boundary and learner plan |
| [CoveStack-Extended-Roadmap.md](CoveStack-Extended-Roadmap.md) | Extended track and Optional future work; preserved design |
| [DocsV5-DO_Team.md](DocsV5-DO_Team.md) | Technical Design and Team Architecture |
| [CoveStack-Service-Catalog.md](CoveStack-Service-Catalog.md) | Services, Capabilities and Release Boundaries |
| [Folder_Structure_Explanation.md](Folder_Structure_Explanation.md) | Folder and File Responsibilities |
| [CoveStack-Complete-Working-Setup.md](CoveStack-Complete-Working-Setup.md) | Setup, Deployment and Verification |
| [CoveStack-Project-Workflow.md](CoveStack-Project-Workflow.md) | Delivery Workflow to 25 January 2027 |
| [Team_Roles.md](Team_Roles.md) | Ownership, Reviews and ITISS Responsibilities |
| [tools.md](tools.md) | Tools, Versions and Why They Are Used |
| [CoveStack-Stack-Templates.md](CoveStack-Stack-Templates.md) | One Stack Integration Template |
| [CoveStack-VPC-IPAM-and-Networking.md](CoveStack-VPC-IPAM-and-Networking.md) | VPC, IPAM and Connectivity Policy |
| [CoveStack-Publication-Domains-and-Tunneling.md](CoveStack-Publication-Domains-and-Tunneling.md) | Application Publication, Domains and Tunnels |
| [CoveStack-Instance-Metadata.md](CoveStack-Instance-Metadata.md) | Instance Metadata and Identity Boundaries |
| [CoveStack-Terraform-and-Ansible.md](CoveStack-Terraform-and-Ansible.md) | Terraform, Ansible and Controlled Automation |
| [CoveStack-UI-CLI-and-Monitoring.md](CoveStack-UI-CLI-and-Monitoring.md) | React UI, CLI Parity and Monitoring |
| [CoveStack-Security-and-ITISS.md](CoveStack-Security-and-ITISS.md) | Security, Infrastructure and ITISS Evidence |
| [CoveStack-Operations-and-Recovery.md](CoveStack-Operations-and-Recovery.md) | Operations, Failure Handling and Data Recovery |
| [CoveStack-Testing-and-Acceptance.md](CoveStack-Testing-and-Acceptance.md) | Test Matrix and Gate Evidence |
| [CoveStack-Implementation-Status.md](CoveStack-Implementation-Status.md) | Baseline Status, Missing Work and Contract Changes |
| [CoveStack-Anuj-Next-Steps.md](CoveStack-Anuj-Next-Steps.md) | Anuj Implementation and ITISS Work Order |
| [CoveStack-Team-START-HERE.md](CoveStack-Team-START-HERE.md) | Team Onboarding and Shared Working Rules |
| [CoveStack-Current-Setup-Fixes.md](CoveStack-Current-Setup-Fixes.md) | Baseline Audit and Extension Checklist |
| [CoveStack-PDF-and-Formatting.md](CoveStack-PDF-and-Formatting.md) | Markdown and PDF Export |
| [CoveStack-References-and-Decision-Log.md](CoveStack-References-and-Decision-Log.md) | References, Assumptions and Decisions |
| [CoveStack-Resource-Services.md](CoveStack-Resource-Services.md) | Resource Specifications and Service Behavior |
| [CoveStack-API-and-Compatibility.md](CoveStack-API-and-Compatibility.md) | API Boundaries and Compatibility Matrix |

## 4. Reading order and teammate sharing

Share the whole reviewed V5 folder and exact source branch/revision. Ask everyone to read these in order:

| Order | Main file | Why everyone needs it |
|---|---|---|
| 1 | [START-HERE](CoveStack-Team-START-HERE.md) | Clone/branch rules and first team task |
| 2 | [SRS scope](CoveStack-SRS-Scope.md) | What the team must deliver |
| 3 | [Main design](DocsV5-DO_Team.md) | Architecture, owned logic and boundaries |
| 4 | [Team roles](Team_Roles.md) | Exact A–E tasks and required reviews |
| 5 | [Workflow](CoveStack-Project-Workflow.md) | Dates, dependencies, freeze and cut rules |
| 6 | [Setup](CoveStack-Complete-Working-Setup.md) | Real baseline checks, target installation and limits |
| 7 | [Service gates](CoveStack-Service-Gates.md) | When a service may be promoted |
| 8 | [Acceptance](CoveStack-AntiWrapper-Acceptance.md) | Tests 1–15 and honest CORE-5 viva answers |
| 9 | [Implementation status](CoveStack-Implementation-Status.md) | Planned work versus source/CI evidence |

| Owner | Additional required reading |
|---|---|
| A — Anuj | Go Handover; Anuj Next Steps; Resource Services; Networking; Security; Operations; Schema agent columns |
| B | Database Schema; API; Resource Services; Stack Templates; Current Setup Fixes |
| C | Database Schema; Security; Acceptance/Testing; Networking; credential and audit decisions |
| D | API/Compatibility; Database Schema; Resource Services; Terraform fixtures |
| E | Go Handover; API; Folder Structure; tools; Terraform/Ansible bounded scope |
| Everyone later | Extended Roadmap before starting any deferred task |

Runtime contracts are authoritative: read contracts/README.md, contracts/internal-api.yaml, contracts/lifecycle-rules.json and signing rules/vectors for assigned integration work. Clone the whole repo; share repository links rather than stale separate source copies. The Java team works in kan-server; A agent work in kan-agent; E planned CLI in cmd/covestack. These are paths within one repository, not push destinations.

## 5. Agreed offering — CORE-5 only

| Service | Bound | Status |
|---|---|---|
| Compute/Droplets | Container-backed, single host; not VMs | Planned |
| Kubernetes | Required; max 2 clusters/host; 1 server + max 2 agents | Planned |
| Managed PostgreSQL | Backup/restore, scoped credentials and private networking | Planned |
| Object storage | S3-compatible MinIO, per-account credentials and bucket scoping | Planned |
| VPC/IPAM | One private network/account; simple allocator; no custom CIDR | Planned |

Security, Recovery, CLI first, optional Terraform Droplet/K8s, one Postgres + MinIO + app Droplet stack and Ansible Part 1 are cross-cutting. The separate [roadmap](CoveStack-Extended-Roadmap.md) is not SRS scope or deliverables. No UI work during CORE-5.

**Service Gate Rule:** A service is called **Supported** only after all four gates pass and its status is **Gates passed**: **G1** a DB row holds desired state and generation; **G2** reconciliation survives a crash without duplicates and rebuilds state/configuration from DB rows; **G3** another account receives **404**; **G4** one service failure test runs in CI. The DB is the source of truth, never a container, cluster, bucket, proxy or configuration file.

Status vocabulary: **Planned | In progress | Gates passed | Extended track**. CORE-5 defaults to **Planned**; the extended track defaults to **Extended track**. Do not promote a status from documentation or a mock result.

## 6. Print, Git and sharing

All comparisons use real Markdown tables. Wider gate/schema tables and Mermaid ERD require suitable rendering; see [PDF limitations](CoveStack-PDF-and-Formatting.md). The existing HTML/print scripts are unchanged and do not represent these edits. No current PDF is generated.

V5 is now tracked in checkpoint **40291e6** (docs: batch 1.1); current batches remain uncommitted. Earlier local/ignored statements describe the pre-checkpoint state. No ignore rule, root README, source file, script or HTML is changed by Batch 3. GitHub links show only the revision actually pushed; do not describe local edits as published. Share only reviewed docs/source, never deployment secrets, private .env, state, kubeconfigs or backups containing credentials.
