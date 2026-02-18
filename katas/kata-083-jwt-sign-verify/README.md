# Kata 083 — JWT Sign/Verify

**Focus:** crypto, encoding

## Your task
Implement:

```go
func SignJWT(claims map[string]any, secret []byte, ttl time.Duration) (string, error)
```

### Learning goal
- What you are building: Build `func SignJWT(claims map[string]any, secret []byte, ttl time.Duration) (string, error)` as a reliable contract. Focus: crypto, encoding.
- Why this matters in real projects: Security paths have no soft failures. Exact behavior is non-negotiable.
- How this grows your Go skills: You practice fail-closed validation and byte-accurate handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: use HS256 for sign/verify; include an `exp` claim; and also provide `VerifyJWT(token, secret)`.

### Tips
- Fail closed on malformed data.
- Never ignore crypto or encoding errors.
- Add tamper and expiry negative tests.

## Rules / Expectations
- HS256
- include exp
- also implement VerifyJWT(token, secret)

## Prior reading
- [RFC 7519 (JWT)](https://www.rfc-editor.org/rfc/rfc7519)
- [Go encoding/base64 package](https://pkg.go.dev/encoding/base64)
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
