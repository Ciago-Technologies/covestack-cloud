# Control-plane directory placeholder

The actual Java module is [kan-server](../kan-server), not this placeholder. Keep Java/Spring Boot and JdbcClient/plain SQL; PostgreSQL is resource authority. Auth/replay components exist but complete server boot/resource APIs/foundational migrations remain pending.

Read [target architecture](../docs/platform/architecture.md), [current evidence](../docs/platform/implementation-status.md) and [decision gaps](../docs/platform/decisions.md). Future Go component extraction requires an explicit ADR, not a wholesale language rewrite. No backend host shell execution or direct tenant runtime socket exposure.
