#!/usr/bin/env python3
"""Generate wire schemas, state reference and language catalogs from one JSON source.
Use --check in CI; no third-party dependencies are required.
"""
import argparse
import copy
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parent
REPO = ROOT.parent

def ref(name):
    return {"$ref": "#/components/schemas/" + name}

def yaml_text(value, indent=0):
    """Emit the JSON subset of YAML with quoted keys/scalars (deterministic)."""
    pad = " " * indent
    if isinstance(value, dict):
        lines = []
        for key, child in value.items():
            prefix = pad + json.dumps(key, ensure_ascii=False) + ":"
            if isinstance(child, (dict, list)) and child:
                lines.extend([prefix, yaml_text(child, indent + 2).rstrip()])
            else:
                lines.append(prefix + " " + json.dumps(child, ensure_ascii=False))
        return "\n".join(lines) + "\n"
    if isinstance(value, list):
        lines = []
        for child in value:
            if isinstance(child, (dict, list)) and child:
                lines.extend([pad + "-", yaml_text(child, indent + 2).rstrip()])
            else:
                lines.append(pad + "- " + json.dumps(child, ensure_ascii=False))
        return "\n".join(lines) + "\n"
    return pad + json.dumps(value, ensure_ascii=False) + "\n"

def validate_catalog(catalog):
    for name, rule in {**catalog["resources"], **catalog["entities"]}.items():
        states = rule.get("observedStates", rule.get("states"))
        assert len(states) == len(set(states)), name
        assert rule["initialState"] in states, name
        assert set(rule["transitions"]) == set(states), name
        for state, targets in rule["transitions"].items():
            assert set(targets) <= set(states), (name, state)
            if state in rule["terminalStates"]:
                assert targets == [], (name, state)
        if "desiredStates" in rule:
            assert set(rule["acknowledges"]) == set(rule["desiredStates"]), name
            assert set(rule["acknowledges"].values()) <= set(states), name
            assert not set(rule["reservedObservedStates"]) & set(states), name

def build_api(catalog):
    api = json.loads((ROOT / "internal-api.base.json").read_text())
    api["info"]["version"] = catalog["version"]
    api["info"]["description"] = (
        "Covestack Cloud pull-only signed agent API. Generated from lifecycle-rules.json "
        "and internal-api.base.json. Defined kinds do not imply implemented drivers. "
        "Agents advertise supportedKinds; work must be filtered to that capability set. "
        "Server-owned groups/buckets/operations/metadata never become agent work. Coordinated "
        "Anuj/B/C review required for contract changes.")
    schemas = api["components"]["schemas"]
    for old in ["Resource", "ObservedReport"]:
        schemas.pop(old, None)
    resources = catalog["resources"]
    agent_kinds = [k for k, rule in resources.items() if rule["owner"] == "agent"]
    schemas["ResourceKind"] = {"type": "string", "enum": list(resources)}
    schemas["AgentResourceKind"] = {"type": "string", "enum": agent_kinds}
    schemas["ResourceBase"] = {
        "type": "object", "required": ["id", "kind", "generation", "desiredState", "deleting", "spec", "labels"],
        "properties": {
            "id": {"type": "string", "minLength": 1, "maxLength": 128},
            "kind": ref("ResourceKind"),
            "generation": {"type": "integer", "format": "int64", "minimum": 1},
            "desiredState": {"type": "string"},
            "state": {"type": "string", "description": "Optional last server-stored observed state; kind-specific enum enforced when present."},
            "deleting": {"type": "boolean", "description": "True exactly for the kind's terminal deletion intent: archive for compute.instance; deleted otherwise."},
            "spec": {"type": "object", "description": "Kind-specific spec; service DTO validation still required before provisioning. No driver is implied by this schema."},
            "labels": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Server-generated ownership labels; user tags cannot overwrite reserved labels."}}}
    schemas["ObservationBase"] = {
        "type": "object", "required": ["kind", "observedAt", "state"],
        "properties": {
            "kind": ref("AgentResourceKind"),
            "observedGeneration": {"type": "integer", "format": "int64", "minimum": 1, "nullable": True,
                "description": "Only current generation after successful convergence. Omit/null for progress, errors or health; stale generations are ignored and future generations rejected."},
            "observedAt": {"type": "string", "format": "date-time", "description": "UTC observation time; server enforces generation, timestamp and agent ownership/lease ordering."},
            "state": {"type": "string"},
            "detail": {"type": "object"},
            "error": {"type": "string", "nullable": True}}}
    all_names, agent_names, observations, mapping, observed_mapping = [], [], [], {}, {}
    for kind, rule in resources.items():
        name = "".join(part.title() for part in kind.replace("-", ".").split("."))
        resource_name = name + "Resource"
        terminal_intent = next(d for d, s in rule["acknowledges"].items() if s in rule["terminalStates"])
        variants = [{"type": "object", "properties": {"desiredState": {"enum": [d]}, "deleting": {"enum": [d == terminal_intent]}}} for d in rule["desiredStates"]]
        schemas[resource_name] = {"allOf": [ref("ResourceBase"), {"type": "object", "properties": {
            "kind": {"type": "string", "enum": [kind]}, "desiredState": {"type": "string", "enum": rule["desiredStates"]}, "state": {"type": "string", "enum": rule["observedStates"]}}}, {"oneOf": variants}]}
        all_names.append(resource_name); mapping[kind] = "#/components/schemas/" + resource_name
        if rule["owner"] == "agent":
            agent_names.append(resource_name)
            observation_name = name + "Observation"
            schemas[observation_name] = {"allOf": [ref("ObservationBase"), {"type": "object", "properties": {
                "kind": {"type": "string", "enum": [kind]}, "state": {"type": "string", "enum": rule["observedStates"]}}}]}
            observations.append(observation_name); observed_mapping[kind] = "#/components/schemas/" + observation_name
    schemas["Resource"] = {"oneOf": [ref(n) for n in all_names], "discriminator": {"propertyName": "kind", "mapping": mapping}}
    schemas["AgentResource"] = {"oneOf": [ref(n) for n in agent_names], "discriminator": {"propertyName": "kind", "mapping": {k: mapping[k] for k in agent_kinds}}}
    schemas["ObservedReport"] = {"oneOf": [ref(n) for n in observations], "discriminator": {"propertyName": "kind", "mapping": observed_mapping}}
    for name, rule in catalog["entities"].items():
        schemas[name.title() + "Status"] = {"type": "string", "enum": rule["states"], "description": rule["notes"] + "; server-owned, not agent resource work"}
    heartbeat = api["paths"]["/internal/v1/agent/heartbeat"]["post"]
    hb_schema = heartbeat["requestBody"]["content"]["application/json"]["schema"]
    hb_schema["required"] = ["agentVersion", "dockerVersion", "supportedKinds"]
    hb_schema["properties"]["supportedKinds"] = {"type": "array", "uniqueItems": True, "maxItems": len(agent_kinds), "items": ref("AgentResourceKind"), "description": "Advertise only installed, tested drivers. Empty means no resource execution capability; never advertise planned drivers."}
    work = api["paths"]["/internal/v1/agent/work"]["get"]
    work["description"] = "Only agent-owned resources assigned to this authenticated agent/account and advertised supportedKinds. Exclude server-managed groups/buckets and already-acknowledged terminal resources."
    work["responses"]["200"]["content"]["application/json"]["schema"] = {"type": "array", "items": ref("AgentResource"), "maxItems": 1000}
    known = api["paths"]["/internal/v1/agent/known-ids"]["get"]
    known["description"] = "Complete account/agent-scoped IDs for agent-owned supported kinds only. Groups and server-managed buckets are excluded. Partial/time-out results never authorize cleanup; do not silently truncate."
    observed = api["paths"]["/internal/v1/agent/resources/{id}/observed"]["put"]
    observed["description"] = "kind must equal the owned stored resource kind. Enforce lifecycle and successful generation acknowledgement; return applied=false for stale reports. LifecycleValidator alone does not enforce account/lease/timestamp ordering."
    response = observed["responses"]["200"]["content"]["application/json"]["schema"]
    response["required"] = ["ok", "applied"]
    response["properties"] = {"ok": {"type": "boolean"}, "applied": {"type": "boolean"}}
    api["paths"]["/internal/v1/agent/events"]["post"]["responses"]["200"]["description"] = "Event accepted for audit; events do not directly mutate resource state."
    return api

def build_markdown(catalog):
    lines = ["# CloveStack Cloud — Frozen Lifecycle Contract", "", "| Contract | Rule |", "|---|---|", "| Version | **" + catalog["version"] + "** |", "| Source | Edit `lifecycle-rules.json`, then run `python3 contracts/generate_lifecycle.py`. |", "| Review | Anuj/B/C review required; CI checks generated files. |", "| Implementation | Defined vocabulary does not mean drivers/controllers are operational. |", "", "## Resource lifecycle vocabulary", "", "| Kind | Service | Owner | Desired states | Terminal state |", "|---|---|---|---|---|"]
    for kind, rule in catalog["resources"].items():
        lines.append(f'| `{kind}` | {rule["service"]} | {rule["owner"]} | ' + ", ".join('`'+state+'`' for state in rule['desiredStates']) + ' | ' + ", ".join('`'+state+'`' for state in rule['terminalStates']) + ' |')
    lines += ["", "### Observed states — one state per row", "", "| Kind | Allowed observed state |", "|---|---|"]
    for kind, rule in catalog["resources"].items():
        for state in rule["observedStates"]:
            lines.append(f'| `{kind}` | `{state}` |')
    lines += ["", "## Convergence, deletion and failure", "", "- Desired state is intent; observed state is reality. Progress/degraded/error is not success.", "- Set `deleting=true` exactly for deletion intent: `archive` for instances, `deleted` otherwise.", "- Set observedGeneration only to the current target generation after the matching stable state is reached. Progress/errors/health omit it or use null.", "- Stale generations return `applied=false`; future/invalid generations are rejected. Actual timestamp, account and agent lease/fencing checks belong in B/C services.", "- Each observed report includes kind and observedAt; kind must match the stored owned resource. Groups and buckets are server-owned and cannot be agent reports/work.", "- Every heartbeat advertises supportedKinds; the server filters work accordingly. Reject unknown kinds; never mark unsupported work successful.", "- Work, known IDs and cleanup remain account/agent scoped. Outages or incomplete known-ID responses never authorize deletion.", "- No terminal-state resurrection. Desired deletion intent cannot be reversed by retries; GC waits for teardown acknowledgement. Failure does not acknowledge the target generation.", "- Logical transitions below may have unreported intermediate states. Validation permits reachable states, including same-state reports, without traversing a terminal state. Observations alone do not authorize side effects.", "- Cancellation applies only to queued operations. A failed operation attempt is terminal; a retry is a new linked attempt.", ""]
    for kind, rule in catalog["resources"].items():
        lines += ["## " + kind, "", rule['readiness'] + ".", "", "Success mapping: " + "; ".join(f'`{d}` → `{s}`' for d,s in rule['acknowledges'].items()) + ".", "", "| From | Allowed next logical states |", "|---|---|"]
        for state, targets in rule["transitions"].items():
            lines.append('| `'+state+'` | '+(', '.join('`'+s+'`' for s in targets) if targets else 'Terminal; no outgoing transition')+' |')
        if rule["reservedObservedStates"]:
            lines += ["", "Reserved future states (not accepted in this contract): " + ", ".join('`'+s+'`' for s in rule['reservedObservedStates']) + ". Add schemas/operations/tests in a reviewed revision before enabling them."]
        lines.append("")
    lines += ["## Server-owned operation and notification lifecycles", ""]
    for entity,rule in catalog['entities'].items():
        lines += ['### '+entity,'',rule['notes']+'.','','| From | Allowed next |','|---|---|']
        for state,targets in rule['transitions'].items():lines.append('| `'+state+'` | '+(', '.join('`'+s+'`' for s in targets) if targets else 'Terminal')+' |')
        lines.append('')
    lines += ['## Services without agent resource lifecycles','','| Service | Owner | Policy |','|---|---|---|']
    for service in catalog['metadataServices']:lines.append(f'| {service["service"]} | {service["owner"]} | {service["policy"]} |')
    lines += ['', '## Provider mapping and storage boundaries', '', 'These are internal states. D maps them to provider DTOs/status vocabulary; internal', 'strings must not be copied blindly into public responses. Operation statuses are', 'separate from resource states. E/C define administrative credential/account policy.', 'MinIO/Postgres/Valkey/Mailpit daemon readiness is infrastructure health, not a new', 'agent resource kind. Managed databases are account resources; control-plane Postgres', 'is infrastructure. Storage buckets/objects are not compute snapshots.', '', '## Example: Kubernetes convergence', '', 'Java creates desiredState=active, state=provisioning, generation=1. Agent reports', 'kind=k8s.cluster, state=running, observedAt and observedGeneration=1 only after', 'readiness. Delete sets desiredState=deleted and deleting=true, increments generation;', 'deleting progress omits observedGeneration, then deleted acknowledges that generation.', '', '## Acceptance', '', 'Run `python3 contracts/check_lifecycle.py`, Go lifecycle tests, Java lifecycle tests', 'and real scoped controller/driver integration. All planned resource kinds have examples;', 'groups/buckets appear only in server-resource examples. Do not claim those examples', 'prove a service is operational.', '']
    return '\n'.join(lines)

def tabular_markdown(text):
    """Keep headings, existing tables and code; make prose/list items scannable rows."""
    lines = text.splitlines(); result = []; pending = []; fenced = False
    def flush():
        if not pending: return
        result.extend(["", "| Item | Details |", "|---|---|"])
        for index, block in enumerate(pending, 1):
            result.append("| " + str(index) + " | " + block.replace("|", "&#124;") + " |")
        result.append(""); pending.clear()
    paragraph = []
    def end_paragraph():
        if paragraph: pending.append(" ".join(paragraph)); paragraph.clear()
    for line in lines:
        if line.lstrip().startswith("```") or line.lstrip().startswith("~~~"):
            end_paragraph(); flush(); fenced = not fenced; result.append(line); continue
        if fenced: result.append(line); continue
        if line.startswith("#") or line.lstrip().startswith("|") or line.strip() in ("---", "***"):
            end_paragraph(); flush(); result.append(line); continue
        if not line.strip(): end_paragraph(); continue
        if line.lstrip().startswith(("- ", "* ")) or __import__("re").match(r"^\s*\d+\. ", line):
            end_paragraph(); pending.append(__import__("re").sub(r"^\s*(?:[-*]|\d+\.)\s+", "", line)); continue
        paragraph.append(line.strip())
    end_paragraph(); flush()
    spaced = []
    for line in result:
        table = line.lstrip().startswith("|")
        previous_table = bool(spaced) and spaced[-1].lstrip().startswith("|")
        if table and spaced and spaced[-1].strip() and not previous_table: spaced.append("")
        if previous_table and line.strip() and not table: spaced.append("")
        spaced.append(line)
    return "\n".join(spaced).strip() + "\n"

def outputs(catalog):
    raw = json.dumps(catalog, indent=2, ensure_ascii=False)
    go = '// Code generated by contracts/generate_lifecycle.py; DO NOT EDIT.\npackage lifecycle\n\nconst catalogJSON = `' + raw + '`\n'
    java = 'package com.ciagotech.kan.server.contract;\n\n/** Generated by contracts/generate_lifecycle.py; DO NOT EDIT. */\nfinal class LifecycleData {\n    private LifecycleData() {}\n    static final String JSON = """\n' + raw + '\n""";\n}\n'
    return {
        ROOT/'FROZEN-STATES.md': tabular_markdown(build_markdown(catalog)),
        ROOT/'internal-api.yaml': '# Generated by generate_lifecycle.py; edit lifecycle-rules.json/internal-api.base.json.\n'+yaml_text(build_api(catalog)),
        REPO/'kan-agent/internal/lifecycle/catalog_generated.go': go,
        REPO/'kan-server/src/main/java/com/ciagotech/kan/server/contract/LifecycleData.java': java}

def main():
    parser = argparse.ArgumentParser(); parser.add_argument('--check', action='store_true'); args = parser.parse_args()
    catalog = json.loads((ROOT/'lifecycle-rules.json').read_text()); validate_catalog(catalog)
    stale=[]
    for path, content in outputs(catalog).items():
        if args.check:
            if not path.exists() or path.read_text()!=content: stale.append(str(path.relative_to(REPO)))
        else:
            path.parent.mkdir(parents=True, exist_ok=True);path.write_text(content)
    if stale: raise SystemExit('Generated files stale: '+', '.join(stale))
    print('Lifecycle generated files '+('match' if args.check else 'updated'))

if __name__=='__main__': main()
