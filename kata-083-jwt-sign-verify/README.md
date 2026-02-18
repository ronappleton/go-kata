# Kata 083 — JWT Sign/Verify

**Focus:** crypto, encoding

## Your task
Implement:

```go
func SignJWT(claims map[string]any, secret []byte, ttl time.Duration) (string, error)
```

### Learning goal
- Expected work: Implement `JWT Sign/Verify` with strict validation and byte-level correctness for all cryptographic operations.
- Why: `JWT Sign/Verify` teaches security precision, where small mistakes can become real vulnerabilities.
- When correct: `JWT Sign/Verify` should satisfy the required behavior, including: `HS256`; `include exp`; and `also implement VerifyJWT(token, secret)`.

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
- This kata trains crypto, encoding by implementing `JWT Sign/Verify` under explicit constraints.
- It is important because `JWT Sign/Verify` reflects production security code, where correctness mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
