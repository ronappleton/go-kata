# Kata 084 — HMAC Request Signing

**Focus:** crypto/hmac

## Your task
Implement:

```go
func SignRequest(method, path string, body []byte, secret []byte) string
```

### Learning goal
- What you are practicing: Build `HMAC Request Signing` with careful validation and exact byte-level behavior for security-sensitive inputs/outputs.
- Why it matters: You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.
- How this grows your Go skills: This builds your discipline for correctness-first coding in high-risk code paths.
- When correct: When your solution is correct, it should satisfy: `canonicalize input`; `hex signature`; and `write VerifyRequest helper`.

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
- This kata is focused practice in crypto/hmac through `HMAC Request Signing`.
- You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
