# Kata 073 — Merkle Tree

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata73() error
```

### Learning goal
- What you are building: Build `func Kata73() error` as a reliable contract. Focus: real-world Go integration and testing.
- Why this matters in real projects: Security paths have no soft failures. Exact behavior is non-negotiable.
- How this grows your Go skills: You practice fail-closed validation and byte-accurate handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: match the full README contract; include tests for happy path and edge cases; and keep API and error style idiomatic Go.

### Tips
- Fail closed on malformed data.
- Never ignore crypto or encoding errors.
- Add tamper and expiry negative tests.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
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
