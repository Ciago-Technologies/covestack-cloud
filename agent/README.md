# Agent directory placeholder

The actual Go module is [kan-agent](../kan-agent), not this placeholder. Existing config, signing/client and no-execution polling entry point are described in [implementation status](../docs/platform/implementation-status.md).

New target: real VM Instance through evaluated Incus backend; distinct App Container and k3d providers; explicit network/storage capabilities. No driver is established. See [agent/reconciliation](../docs/platform/agent-and-reconciliation.md), [API/configuration](../docs/platform/api-and-cli.md) and [contribution guide](../docs/platform/contributing.md). No direct DB access or privileged socket access for tenants. Existing wire and module layout are unchanged.

A full runnable product quick-start is unavailable pending real server/credentials and runtime gates. Previous environment commands are preserved in the [pre-redesign snapshot](../docs/archive/pre-vm-redesign-2026-11-04/agent--README.md), not a current deployment guarantee.
