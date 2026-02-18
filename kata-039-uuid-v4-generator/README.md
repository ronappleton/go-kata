# Kata 039 — UUID v4 Generator

**Focus:** crypto/rand

## Your task
Implement:

```go
func NewUUIDv4() (string, error)
```

### Learning goal
- Expected work: Implement `UUID v4 Generator` with strict validation and byte-level correctness for all cryptographic operations.
- Why: `UUID v4 Generator` teaches security precision, where small mistakes can become real vulnerabilities.

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
- This kata trains crypto/rand by implementing `UUID v4 Generator` under explicit constraints.
- It is important because `UUID v4 Generator` reflects production security code, where correctness mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
