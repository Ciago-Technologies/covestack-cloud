# Historical pre-VM entry-point snapshot

> Preserved original body below, not active installation, architecture or schedule guidance. Use [current platform docs](../../platform/README.md). Original relative links require the original checkout layout; no runtime result is revised.

<!-- Historical body preserved below. -->

# CoveStack Cloud — Local Infrastructure

| Item | Details                                                                                                                                                                                                                                                                                                      |
| ---- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| 1    | Actual Compose provides Postgres, Valkey and Mailpit. From root use `docker compose --env-file deploy/.env -f deploy/docker-compose.yml config --quiet`. Create local secrets/image references from templates; never commit them. See [setup guide](../V5/CoveStack-Complete-Working-Setup.md). |
