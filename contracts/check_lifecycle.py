#!/usr/bin/env python3
"""Offline lifecycle/fixture checks, using only the stdlib.
The schema checker covers the keywords emitted by our generator, not arbitrary
OpenAPI/JSON Schema documents. Runtime validation/authorization remains Go/Java work.
"""
import copy
import datetime
import json
import subprocess
import sys
from pathlib import Path
from generate_lifecycle import ROOT, build_api, validate_catalog

class Invalid(ValueError):
    pass

def validate(schema, value, schemas):
    if "$ref" in schema:
        return validate(schemas[schema["$ref"].rsplit("/", 1)[1]], value, schemas)
    if value is None and schema.get("nullable"):
        return
    for branch in schema.get("allOf", []):
        validate(branch, value, schemas)
    if "oneOf" in schema:
        count = 0
        for branch in schema["oneOf"]:
            try:
                validate(branch, value, schemas); count += 1
            except Invalid:
                pass
        if count != 1:
            raise Invalid("oneOf matched " + str(count) + " variants")
    expected = schema.get("type")
    checks = {"object": isinstance(value, dict), "array": isinstance(value, list),
              "string": isinstance(value, str), "integer": type(value) is int,
              "boolean": type(value) is bool}
    if expected and not checks[expected]:
        raise Invalid("wrong type: " + expected)
    if "enum" in schema and not any(type(value) is type(item) and value == item for item in schema["enum"]):
        raise Invalid("not in enum")
    if type(value) is int:
        if "minimum" in schema and value < schema["minimum"]:
            raise Invalid("below minimum")
        if schema.get("format") == "int64" and not -(2**63) <= value < 2**63:
            raise Invalid("out of int64 range")
    if isinstance(value, str):
        if len(value) < schema.get("minLength", 0) or len(value) > schema.get("maxLength", sys.maxsize):
            raise Invalid("string bounds")
        if schema.get("format") == "date-time":
            try:
                date = datetime.datetime.fromisoformat(value.replace("Z", "+00:00"))
                if date.utcoffset() != datetime.timedelta(0):
                    raise ValueError("not UTC")
            except ValueError as error:
                raise Invalid("invalid UTC observation time") from error
    if isinstance(value, dict):
        if not set(schema.get("required", [])) <= set(value):
            raise Invalid("required field absent")
        for key, child in value.items():
            if key in schema.get("properties", {}):
                validate(schema["properties"][key], child, schemas)
            elif isinstance(schema.get("additionalProperties"), dict):
                validate(schema["additionalProperties"], child, schemas)
            elif schema.get("additionalProperties") is False:
                raise Invalid("unknown field")
    if isinstance(value, list):
        if len(value) > schema.get("maxItems", sys.maxsize):
            raise Invalid("array bounds")
        if schema.get("uniqueItems") and len({json.dumps(v, sort_keys=True) for v in value}) != len(value):
            raise Invalid("duplicate capability")
        for child in value:
            validate(schema.get("items", {}), child, schemas)

def load(name):
    return json.loads((ROOT/"examples"/name).read_text())

def rejects(schema, value, schemas):
    try:
        validate(schema, value, schemas)
    except Invalid:
        return
    raise AssertionError("invalid fixture accepted")

def main():
    subprocess.run([sys.executable, str(ROOT/"generate_lifecycle.py"), "--check"], check=True)
    catalog = json.loads((ROOT/"lifecycle-rules.json").read_text()); validate_catalog(catalog)
    api = build_api(catalog); schemas = api["components"]["schemas"]
    resources = load("server-resources-all-services.json")
    work = load("get-work-response-all-services.json")
    observations = load("observed-reports-all-services.json")
    agent_kinds = {kind for kind, rule in catalog["resources"].items() if rule["owner"] == "agent"}
    assert {r["kind"] for r in resources} == set(catalog["resources"])
    assert {r["kind"] for r in work} == agent_kinds
    assert {r["kind"] for r in observations} == agent_kinds
    for resource in resources:
        validate(schemas["Resource"], resource, schemas)
    for resource in work + load("get-work-response-k8s.json"):
        validate(schemas["AgentResource"], resource, schemas)
    for observation in observations + [load("put-observed-request-k8s.json")]:
        validate(schemas["ObservedReport"], observation, schemas)
    heartbeat_schema = api["paths"]["/internal/v1/agent/heartbeat"]["post"]["requestBody"]["content"]["application/json"]["schema"]
    heartbeat = load("heartbeat-request.json"); validate(heartbeat_schema, heartbeat, schemas)
    server_entities = load("server-operation-and-notification.json")
    for kind, item in server_entities.items():
        validate(schemas[kind.title()+"Status"], item["status"], schemas)
    group = next(r for r in resources if r["kind"] == "compute.group")
    rejects(schemas["AgentResource"], group, schemas)
    wrong = copy.deepcopy(group); wrong["state"] = "running"
    rejects(schemas["Resource"], wrong, schemas)
    bucket = next(r for r in resources if r["kind"] == "storage.bucket")
    rejects(schemas["AgentResource"], bucket, schemas)
    wrong = copy.deepcopy(bucket); wrong["state"] = "active"
    rejects(schemas["Resource"], wrong, schemas)
    wrong = copy.deepcopy(work[0]); wrong["desiredState"] = "online"
    rejects(schemas["AgentResource"], wrong, schemas)
    wrong = copy.deepcopy(work[0]); wrong["desiredState"] = "archive"; wrong["deleting"] = False
    rejects(schemas["AgentResource"], wrong, schemas)
    wrong = copy.deepcopy(observations[2]); wrong["state"] = "active"  # k8s needs running
    rejects(schemas["ObservedReport"], wrong, schemas)
    wrong = copy.deepcopy(observations[0]); del wrong["kind"]
    rejects(schemas["ObservedReport"], wrong, schemas)
    wrong = copy.deepcopy(observations[0]); del wrong["observedAt"]
    rejects(schemas["ObservedReport"], wrong, schemas)
    wrong = copy.deepcopy(heartbeat); wrong["supportedKinds"] = ["compute.group"]
    rejects(heartbeat_schema, wrong, schemas)
    print("Lifecycle coverage: 7 resource kinds, 5 agent kinds, server operations/notifications; positive/negative schema fixtures PASS")

if __name__ == "__main__":
    main()
