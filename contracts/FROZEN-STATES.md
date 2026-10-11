# Frozen state names

Do not change spelling or casing of these. If Java writes "Terminated"
and Go checks for "terminated", things break silently. These strings
are locked the day the contract is frozen.

| kind             | desiredState values           | state values (internal, written to DB)                                           |
|------------------|-------------------------------|----------------------------------------------------------------------------------|
| compute.instance | active, off, archive          | new, active, off, archive, error                                                 |
| db.instance      | online, stopped, deleted      | creating, online, stopped, resizing, migrating, forking, deleted, error                   |
| k8s.cluster      | active, deleted               | provisioning, running, degraded, error, deleting, deleted                        |
| compute.group    | (Java-only, agent ignores)    | active, deleting                                                                 |

## Example: k8s.cluster lifecycle

1. User asks for a cluster -> Java inserts row: `desiredState: active`, `state: provisioning`, `generation: 1`
2. Agent polls `/work`, sees the row, starts k3d cluster
3. Agent PUTs `/observed`: `observedGeneration: 1`, `state: running`
4. Java's `describe-clusters` now shows `running`
5. User deletes -> Java sets `deleting: true`, bumps `generation`
6. Agent tears down k3d cluster, PUTs `/observed`: `state: deleted`
