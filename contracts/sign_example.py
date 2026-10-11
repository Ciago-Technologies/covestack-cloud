#!/usr/bin/env python3
"""Rules-driven reference signer and shared-vector verifier (stdlib only)."""
import hashlib
import hmac
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parent
RULES = json.loads((ROOT / "sigv1-rules.json").read_text())

def build_canonical_request(method, path, query, headers, body):
    lower = {}
    for key, value in headers.items():
        name = key.lower()
        if name in lower:
            raise ValueError("duplicate header: " + name)
        lower[name] = value.strip(" \t")
    names = set(RULES["alwaysSignedHeaders"])
    if body and RULES["signBodyContentTypeHeader"]:
        names.add("content-type")
    names = sorted(names)
    for name in names:
        value = lower.get(name, "")
        if not value or "\r" in value or "\n" in value:
            raise ValueError("missing or invalid signed header: " + name)
    canonical_headers = "".join(f"{name}:{lower[name]}\n" for name in names)
    raw_body = body.encode("utf-8") if isinstance(body, str) else body
    return "\n".join([method, path, query, canonical_headers, ";".join(names),
                       hashlib.sha256(raw_body).hexdigest()]), ";".join(names)

def sign(secret_key, method, path, query, headers, body, date_str):
    dates = [v.strip(" \t") for k, v in headers.items() if k.lower() == RULES["dateHeader"].lower()]
    if dates != [date_str] or not secret_key:
        raise ValueError("date mismatch or empty signing secret")
    canonical, names = build_canonical_request(method, path, query, headers, body)
    sts = "\n".join([RULES["algorithm"], date_str, hashlib.sha256(canonical.encode("utf-8")).hexdigest()])
    signature = hmac.new(secret_key.encode("utf-8"), sts.encode("utf-8"), hashlib.sha256).hexdigest()
    return canonical, sts, signature, names

if __name__ == "__main__":
    for vector in json.loads((ROOT / "sigv1-vectors.json").read_text())["vectors"]:
        _, _, actual, _ = sign(vector["secretKey"], vector["method"], vector["path"], vector["query"],
                               vector["headers"], vector["body"], vector["headers"][RULES["dateHeader"]])
        if actual != vector["expectedSignature"]:
            raise SystemExit("Signature mismatch: " + vector["name"])
        print(vector["name"] + ": PASS")
