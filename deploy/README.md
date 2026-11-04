# Local infrastructure definitions — dependencies only

The existing [Compose file](docker-compose.yml) defines PostgreSQL, authenticated Valkey and Mailpit. It does not install/start Incus, VM/App Container execution, MinIO, Java application or the complete Go control loop environment. No configuration was changed by the VM documentation redesign.

[Installation and operations](../docs/platform/installation-and-operations.md) separates current definitions from future installer requirements; [capacity](../docs/platform/capacity.md) defines measured host profiles. Exact tested image references and secrets come from reviewed local templates, not invented docs values. Current sandbox prohibits daemon/runtime provisioning; respect [environment evidence/restrictions](../docs/development-toolchain-status.md).

Preserve initialized PostgreSQL/Valkey volumes; changing a secret file is not rotating initialized credentials. Never use down -v or blanket pruning as a repair step. The prior command reference is preserved in [the historical snapshot](../docs/archive/pre-vm-redesign-2026-11-04/deploy--README.md); it is not evidence of a supported product deployment.
