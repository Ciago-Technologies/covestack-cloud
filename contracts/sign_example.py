#!/usr/bin/env python3
"""
Reference implementation of the proposed kan sigv1 canonical request + signature.
No dependencies beyond the Python standard library. Run: python3 sign_example.py

This is a PROPOSAL to match against, not a copy of an already-agreed spec —
see CANONICAL-REQUEST.md for why. Use this to sanity-check your Java/Go
implementations produce the same output for the same input.
"""
import hashlib
import hmac


def sha256_hex(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def build_canonical_request(method, path, query, headers, body):
    """headers: dict of header name -> value, ONLY the ones being signed."""
    lowered = {k.lower(): v.strip() for k, v in headers.items()}
    sorted_names = sorted(lowered.keys())
    canonical_headers = "".join(f"{name}:{lowered[name]}\n" for name in sorted_names)
    signed_headers = ";".join(sorted_names)
    hashed_payload = sha256_hex(body.encode())
    canonical_request = "\n".join([
        method,
        path,
        query,
        canonical_headers,
        signed_headers,
        hashed_payload,
    ])
    return canonical_request, signed_headers


def sign(secret_key: str, method: str, path: str, query: str,
         headers: dict, body: str, date_str: str):
    canonical_request, signed_headers = build_canonical_request(
        method, path, query, headers, body
    )
    hashed_cr = sha256_hex(canonical_request.encode())
    string_to_sign = "\n".join(["KAN1-HMAC-SHA256", date_str, hashed_cr])
    signature = hmac.new(
        secret_key.encode(), string_to_sign.encode(), hashlib.sha256
    ).hexdigest()
    return canonical_request, string_to_sign, signature, signed_headers


def run_example(title, secret_key, method, path, query, headers, body, date_str):
    cr, sts, sig, sh = sign(secret_key, method, path, query, headers, body, date_str)
    print(f"=== {title} ===")
    print("-- Canonical Request --")
    print(cr)
    print("-- String To Sign --")
    print(sts)
    print("-- Signed Headers --", sh)
    print("-- Signature --", sig)
    print()


if __name__ == "__main__":
    # Vector 1: matches contracts/sigv1-vectors.json -> "get-work-basic"
    run_example(
        "GET /work",
        secret_key="test-secret-do-not-use-in-prod",
        method="GET",
        path="/internal/v1/agent/work",
        query="",
        headers={"Host": "localhost:4700", "X-Kan-Date": "20260926T120000Z"},
        body="",
        date_str="20260926T120000Z",
    )

    # Vector 2: matches contracts/sigv1-vectors.json -> "put-observed-k8s-cluster"
    body2 = (
        '{"observedGeneration":2,"state":"active",'
        '"detail":{"clusterId":"k3d-web-cluster","nodeCount":3},"error":null}'
    )
    run_example(
        "PUT /observed (k8s.cluster)",
        secret_key="test-secret-do-not-use-in-prod",
        method="PUT",
        path="/internal/v1/agent/resources/01J8ZK9XCLUSTER1/observed",
        query="",
        headers={
            "Host": "localhost:4700",
            "X-Kan-Date": "20260926T120005Z",
            "Content-Type": "application/json",
        },
        body=body2,
        date_str="20260926T120005Z",
    )
