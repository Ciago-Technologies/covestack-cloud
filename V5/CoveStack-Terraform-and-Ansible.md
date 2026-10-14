# CoveStack V5 — Terraform, Ansible and Controlled Automation

| Document control | Value |
|---|---|
| Product / documentation | **CoveStack / V5** |
| Updated | **14 October 2026** |
| Submission deadline | **Monday, 25 January 2027** |
| Compatibility names | Server: `kan-server`; agent: `kan-agent` |
| Status | Target design and implementation guidance; not a completion certificate |

## 1. Two separate provisioning interfaces

| Mode | User experience | Implementation |
|---|---|---|
| Direct Terraform | User writes configuration and runs Terraform | CoveStack provider calls authenticated API |
| Platform template | User selects approved stack/variables | Stack engine; optionally a controlled IaC runner |
| Ansible configuration | User configures owned targets | Scoped inventory/modules/connection |

**Uploading `.tf` or playbook files does not establish support.** Terraform needs a provider, state and resource lifecycle. Ansible needs a working target/connection method and configuration semantics. They already manage Docker through existing integrations; CoveStack adds its account/resource/stack contract.

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

Use the official [Terraform framework](https://developer.hashicorp.com/terraform/plugin/framework) and acceptance tests. Start with one Droplet lifecycle, then VPC/publication and whole-stack support. Pin the provider version after testing.

Illustrative future HCL, **not a published provider schema**:

```hcl
provider "covestack" {
  endpoint = var.covestack_endpoint
}

resource "covestack_stack" "shop" {
  name             = "shop"
  template_id      = "web-postgres-storage"
  template_version = "1"
  publish          = true
}
```

Supply tokens through an approved secret mechanism; do not embed them in examples. The supported resource schema must be implemented before this example can run.

## 3. Terraform and controller ownership

| Situation | Rule |
|---|---|
| Stack resource manages children | Do not separately manage same children in another Terraform resource |
| UI modifies Terraform-owned spec | Warn/reject unsupported conflicting writes or document drift policy |
| Controller observes health | Does not overwrite provider's desired configuration |
| ASG-managed members | Do not independently count/manage each generated member |
| Terraform timeout | Refresh durable server result before blindly resubmitting |

A file upload/run interface is a job system, not a provider substitute. Use reviewed templates first. The safest initial implementation lets users run Terraform themselves against APIs; it avoids executing arbitrary user code on the platform host.

## 4. Ansible scope

| Target | Supported approach |
|---|---|
| Control-plane VM | Operator-reviewed roles with appropriate privileges |
| Container-backed Droplet | Approved image/connectivity; module/API or restricted SSH path |
| Kubernetes workload | Authorized kubeconfig and relevant Kubernetes modules |
| CoveStack records | Authenticated API modules/inventory |

An arbitrary minimal image may lack SSH/Python requirements for ordinary Ansible modules. Validate the chosen image/transport rather than promising every playbook will work. Do not give a user Ansible job the host Docker socket to gain convenient container access.

Start with one playbook that installs/configures a supported application, reports failure clearly, and makes no unintended changes on a second run. Return owned targets and authorized connection details from inventory; secrets are protected references and sensitive tasks suppress output.

## 5. Platform-run job model, if implemented

| Control | Requirement |
|---|---|
| Execution isolation | Job cannot modify platform host or other account state |
| Inputs | Reviewed templates/schema; no arbitrary executable plugins |
| Identity | Separate scoped job credential |
| State | Per-account locking and protected storage |
| Secrets | Restricted injection, redacted logs, lifecycle cleanup |
| Limits | Time, CPU/memory, process/output/network bounds |
| Outcome | Exit result, durable operation and failed-step detail |
| Cancellation | Defined partial-result/cleanup behavior |

A provider or playbook can execute programs and reach networks. Treat unrestricted uploads as later scope requiring a separate threat model, not moderate configuration work. Never run all account jobs with one global administrator credential.

## 6. State and sensitive data

Terraform state can contain sensitive values even when CLI output is marked sensitive. Protect access, backups and locking; minimize server-returned secret values. Define when credential rotation changes state and who can retrieve it. Ansible facts/logs can leak passwords too.

## 7. Integration acceptance

Create with Terraform, inspect via CLI/UI, refresh/import, change a supported field and delete under retention policy. Deny another account and test interrupted apply/retry. Run the Ansible example twice, test unreachable/unauthorized targets, and verify logs contain no credentials. Record exact provider/collection versions and real endpoints.
