# Kata 083 — JWT Sign/Verify

**Focus:** crypto, encoding

## Your task
Implement:

```go
func SignJWT(claims map[string]any, secret []byte, ttl time.Duration) (string, error)
```

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
- This kata is about implementing JWT Sign/Verify with constraints that make you practice crypto, encoding.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
