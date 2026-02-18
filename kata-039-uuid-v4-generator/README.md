# Kata 039 — UUID v4 Generator

**Focus:** crypto/rand

## Your task
Implement:

```go
func NewUUIDv4() (string, error)
```

### Learning goal
- What you are practicing: Build `UUID v4 Generator` with careful validation and exact byte-level behavior for security-sensitive inputs/outputs.
- Why it matters: You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.
- How this grows your Go skills: This builds your discipline for correctness-first coding in high-risk code paths.
- When correct: When your solution is correct, it should satisfy: `RFC4122 version/variant bits`; and `format 8-4-4-4-12 lowercase`.

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
- This kata is focused practice in crypto/rand through `UUID v4 Generator`.
- You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
