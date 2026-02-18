# Kata 084 — HMAC Request Signing

**Focus:** crypto/hmac

## Your task
Implement:

```go
func SignRequest(method, path string, body []byte, secret []byte) string
```

### Learning goal
- Expected work: Implement `HMAC Request Signing` with strict validation and byte-level correctness for all cryptographic operations.
- Why: `HMAC Request Signing` teaches security precision, where small mistakes can become real vulnerabilities.
- When correct: `HMAC Request Signing` should satisfy the required behavior, including: `canonicalize input`; `hex signature`; and `write VerifyRequest helper`.

## Rules / Expectations
- canonicalize input
- hex signature
- write VerifyRequest helper

## Prior reading
- [Go crypto/hmac package](https://pkg.go.dev/crypto/hmac)
- [Go crypto/sha256 package](https://pkg.go.dev/crypto/sha256)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains crypto/hmac by implementing `HMAC Request Signing` under explicit constraints.
- It is important because `HMAC Request Signing` reflects production security code, where correctness mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
