# Kata 039 — UUID v4 Generator

**Focus:** crypto/rand

## Your task
Implement:

```go
func NewUUIDv4() (string, error)
```

### Learning goal
- What you are building: Build `func NewUUIDv4() (string, error)` as a reliable contract. Focus: crypto/rand.
- Why this matters in real projects: Security paths have no soft failures. Exact behavior is non-negotiable.
- How this grows your Go skills: You practice fail-closed validation and byte-accurate handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: rfc4122 version/variant bits; and format 8-4-4-4-12 lowercase.

### Tips
- Fail closed on malformed data.
- Never ignore crypto or encoding errors.
- Add tamper and expiry negative tests.

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
- Core lesson: validate early, fail closed, and be exact.
- After this kata, you should be able to justify validation rules and add negative tests for tampering.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
