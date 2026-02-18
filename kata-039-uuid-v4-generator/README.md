# Kata 039 — UUID v4 Generator

**Focus:** crypto/rand

## Your task
Implement:

```go
func NewUUIDv4() (string, error)
```

## Rules / Expectations
- RFC4122 version/variant bits
- format 8-4-4-4-12 lowercase

## Prior reading
- [RFC 4122 (UUID)](https://www.rfc-editor.org/rfc/rfc4122)
- [Go crypto/rand package](https://pkg.go.dev/crypto/rand)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing UUID v4 Generator with constraints that make you practice crypto/rand.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
