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

This directory provides the six replacement guides requested and the additional specifications needed for the expanded platform. Planned services and illustrative commands are explicitly separated from implementation evidence.

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
| [DocsV5-DO_Team.md](DocsV5-DO_Team.md) | Technical Design and Team Architecture |
| [CoveStack-Service-Catalog.md](CoveStack-Service-Catalog.md) | Services, Capabilities and Release Boundaries |
| [Folder_Structure_Explanation.md](Folder_Structure_Explanation.md) | Folder and File Responsibilities |
| [CoveStack-Complete-Working-Setup.md](CoveStack-Complete-Working-Setup.md) | Setup, Deployment and Verification |
| [CoveStack-Project-Workflow.md](CoveStack-Project-Workflow.md) | Delivery Workflow to 25 January 2027 |
| [Team_Roles.md](Team_Roles.md) | Ownership, Reviews and ITISS Responsibilities |
| [tools.md](tools.md) | Tools, Versions and Why They Are Used |
| [CoveStack-Stack-Templates.md](CoveStack-Stack-Templates.md) | Stack Templates and Complete Environment Workflow |
| [CoveStack-VPC-IPAM-and-Networking.md](CoveStack-VPC-IPAM-and-Networking.md) | VPC, IPAM and Connectivity Policy |
| [CoveStack-Publication-Domains-and-Tunneling.md](CoveStack-Publication-Domains-and-Tunneling.md) | Application Publication, Domains and Tunnels |
| [CoveStack-Instance-Metadata.md](CoveStack-Instance-Metadata.md) | Instance Metadata and Identity Boundaries |
| [CoveStack-Terraform-and-Ansible.md](CoveStack-Terraform-and-Ansible.md) | Terraform, Ansible and Controlled Automation |
| [CoveStack-UI-CLI-and-Monitoring.md](CoveStack-UI-CLI-and-Monitoring.md) | React UI, CLI Parity and Monitoring |
| [CoveStack-Security-and-ITISS.md](CoveStack-Security-and-ITISS.md) | Security, Infrastructure and ITISS Evidence |
| [CoveStack-Operations-and-Recovery.md](CoveStack-Operations-and-Recovery.md) | Operations, Failure Handling and Data Recovery |
| [CoveStack-Testing-and-Acceptance.md](CoveStack-Testing-and-Acceptance.md) | Test Matrix and Definition of Done |
| [CoveStack-Implementation-Status.md](CoveStack-Implementation-Status.md) | Baseline Status, Missing Work and Contract Changes |
| [CoveStack-Anuj-Next-Steps.md](CoveStack-Anuj-Next-Steps.md) | Anuj Implementation and ITISS Work Order |
| [CoveStack-Team-START-HERE.md](CoveStack-Team-START-HERE.md) | Team Onboarding and Shared Working Rules |
| [CoveStack-Current-Setup-Fixes.md](CoveStack-Current-Setup-Fixes.md) | Baseline Audit and Extension Checklist |
| [CoveStack-PDF-and-Formatting.md](CoveStack-PDF-and-Formatting.md) | Markdown and PDF Export |
| [CoveStack-References-and-Decision-Log.md](CoveStack-References-and-Decision-Log.md) | References, Assumptions and Decisions |
| [CoveStack-Resource-Services.md](CoveStack-Resource-Services.md) | Resource Specifications and Service Behavior |
| [CoveStack-API-and-Compatibility.md](CoveStack-API-and-Compatibility.md) | API Boundaries and Compatibility Matrix |

## 4. Reading order

Read design, service catalog, roles, workflow, implementation status and setup. Then read stack templates, networking, publication, metadata and integrations for your assigned feature. Security, operations and acceptance define proof requirements; the PDF guide covers export.

## 5. Agreed offering

| Original services | Extensions/platform capability |
|---|---|
| Droplets and Kubernetes | Owned VPC/IPAM and self-metadata |
| PostgreSQL and S3-compatible storage | Built-in optional app publication and domains |
| HTTP LB and auto scaling groups | Complete stack templates and customization |
| Filesystem snapshots | Terraform/Ansible, shared React/CLI and monitoring |

First implement Droplets/Kubernetes; a full web/data/storage template needs its actual database/storage dependencies. Nested Docker/VM-backed instances remain optional later scope. Fake public IPs do not provide internet connectivity. Users visit a real published HTTPS URL.

## 6. Print and sharing

Run `python3 V5/render_v5.py` for a printable combined HTML file; follow [PDF formatting](CoveStack-PDF-and-Formatting.md). All comparison matrices are real Markdown pipe tables with spacing and headers, not code blocks. Each guide has a consistent control table, headings and highlighted decisions.

The repository currently ignores `V5`. This bundle is local and uncommitted; generation does not modify that policy, active code or older documents.
