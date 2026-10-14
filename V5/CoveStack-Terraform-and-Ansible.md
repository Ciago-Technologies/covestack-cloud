# CoveStack V5 — Terraform, Ansible and Controlled Automation

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Bounded CORE-5 automation

| Item | Scope | Owner/status |
|---|---|---|
| Terraform optional | Thin API provider; Droplet + Kubernetes cluster resources only | E after CLI; A Go review; Planned |
| Ansible Part 1 | Operator-reviewed host installation playbook | A with B/C service/security review; Planned |
| Inventory/configuration expansion | Relocated to [roadmap](CoveStack-Extended-Roadmap.md#11-ansible-inventory-part-2) | Extended track |

Uploading .tf/playbooks does not establish integration. Terraform already integrates with execution tools; CoveStack adds account-scoped durable lifecycle. No platform-run arbitrary job system is a CORE-5 requirement.

## 2. Provider slice

| First resource capability | Acceptance |
|---|---|
| Configure endpoint/token | Scoped account identity and TLS |
| Create | Stable ID plus operation wait/cancel timeout |
| Read/refresh | Actual server state and absent-resource handling |
| Update | Valid supported changes; accurate replacement decisions |
| Delete | Idempotent cleanup and retention behavior |
| Import | Adopt an authorized existing resource |
| Drift | Explain difference; no silent duplication |

Use the official [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework) and acceptance tests. CORE-5 optional provider covers Droplet and Kubernetes only, after CLI; broader resources are deferred in the roadmap. Pin the provider version after testing.

No HCL resource names/version strings are approved here. Previous stack/publication examples are preserved as unapproved future design in the roadmap. Supply tokens through protected configuration, never embedded examples.

## 3. Terraform and controller ownership

Provider calls only the authenticated API using the reviewed client; never Docker/k3d/DB. Every mutation has agreed idempotency. Read/refresh after timeout before resubmission; reconcile desired fields without replacing observed health with requested state. Import/adoption and conflicting-write policy need B+D/E review.

E owns the thin provider, D Java mappings/test fixtures, B lifecycle, C tests/auth, A every Go review. CORE-5 does not expose stack/VPC/public-route Terraform resources; broader ownership requirements remain in the roadmap.

## 4. Ansible Part 1 — operator installer

| Planned layout/control | Evidence |
|---|---|
| ansible/site.yml; roles/docker, infra, server, agent | Actual reviewed installation on single AlmaLinux VM |
| systemd NoNewPrivileges, ProtectSystem=strict | Explicit writable paths; no silent weakening |
| ansible-vault and no_log | Secret custody B+C; no sensitive log output |
| Checksum-verified binaries | Actual versions/checksums; mismatch fails |
| ansible-lint | Reviewed playbook passes |
| Repeat run | Second run changed=0 |

The installer intentionally configures operator host dependencies; the API-only rule applies to CLI/provider/inventory, not the install role. No YAML/systemd/code files are created in this docs task. Part 2 inventory and SSH/Python/API-key-injection prerequisites live in the roadmap.

## 5. Deferred runner/inventory designs — relocated

All previous identity, isolated execution, secret/state, time/resource limits, cancellation and output requirements remain in the [roadmap](CoveStack-Extended-Roadmap.md). Inventory Part 2 derives JSON from authenticated CLI list output, never docker ps. Arbitrary uploaded runners are Optional future work, not a provider substitute.

## 6. State and sensitive data

Terraform state can contain sensitive values even when CLI output is marked sensitive. Protect access, backups and locking; minimize server-returned secret values. Define when credential rotation changes state and who can retrieve it. Ansible facts/logs can leak passwords too.

## 7. Integration acceptance

Optional provider: Droplet/K8s create/read/update/delete/import, denied other-account action and interrupted apply/retry through authenticated API. Part 1: clean-VM installer, lint, checksum verification and second run changed=0; record actual versions and private endpoints. No UI parity or dynamic inventory is required in CORE-5.

[Gates](CoveStack-Service-Gates.md) · [Acceptance](CoveStack-AntiWrapper-Acceptance.md) · [Schema](CoveStack-Database-Schema.md) · [Go handover](CoveStack-Go-Handover.md) · [Roadmap](CoveStack-Extended-Roadmap.md) · [SRS scope](CoveStack-SRS-Scope.md).
