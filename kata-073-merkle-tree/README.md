# Kata 073 — Merkle Tree

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata73() error
```

### Learning goal
- Expected work: Implement `Merkle Tree` with strict validation and byte-level correctness for all cryptographic operations.
- Why: `Merkle Tree` teaches security precision, where small mistakes can become real vulnerabilities.
- When correct: `Merkle Tree` should satisfy the required behavior, including: `follow README spec`; `write tests`; and `keep it idiomatic`.

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
- This kata trains integration design, boundary handling, and robust testing by implementing `Merkle Tree` under explicit constraints.
- It is important because `Merkle Tree` reflects production security code, where correctness mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
