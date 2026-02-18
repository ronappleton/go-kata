# Kata 083 — JWT Sign/Verify

**Focus:** crypto, encoding

## Your task
Implement:

```go
func SignJWT(claims map[string]any, secret []byte, ttl time.Duration) (string, error)
```

### Learning goal
- What you are practicing: Build `JWT Sign/Verify` with careful validation and exact byte-level behavior for security-sensitive inputs/outputs.
- Why it matters: You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.
- How this grows your Go skills: This builds your discipline for correctness-first coding in high-risk code paths.
- When correct: When your solution is correct, it should satisfy: `HS256`; `include exp`; and `also implement VerifyJWT(token, secret)`.

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
- This kata is focused practice in crypto, encoding through `JWT Sign/Verify`.
- You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
