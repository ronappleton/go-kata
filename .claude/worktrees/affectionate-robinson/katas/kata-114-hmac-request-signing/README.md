# Kata 114 — HMAC Request Signing

**Focus:** crypto/hmac

## Your task
Implement:

```go
func SignRequest(method, path string, body []byte, secret []byte) string
```

### Learning goal
- What you are building: Build `func SignRequest(method, path string, body []byte, secret []byte) string` as a reliable contract. Focus: crypto/hmac.
- Why this matters in real projects: Security paths have no soft failures. Exact behavior is non-negotiable.
- How this grows your Go skills: You practice fail-closed validation and byte-accurate handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: canonicalize input; hex signature; and write verifyrequest helper.

### Tips
- Fail closed on malformed data.
- Never ignore crypto or encoding errors.
- Add tamper and expiry negative tests.

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
- Core lesson: validate early, fail closed, and be exact.
- After this kata, you should be able to justify validation rules and add negative tests for tampering.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
