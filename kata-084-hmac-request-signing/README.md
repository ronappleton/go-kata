# Kata 084 — HMAC Request Signing

**Focus:** crypto/hmac

## Your task
Implement:

```go
func SignRequest(method, path string, body []byte, secret []byte) string
```

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
- This kata is about implementing HMAC Request Signing with constraints that make you practice crypto/hmac.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
