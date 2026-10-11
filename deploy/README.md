# Covestack Cloud — Local Infrastructure

| Item | Details |
|---|---|
| 1 | Actual Compose provides Postgres, Valkey and Mailpit. From root use `docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet`. Create local secrets/image references from templates; never commit them. See [setup guide](../docs/team/Covestack-Cloud-Complete-Working-Setup.md). |
