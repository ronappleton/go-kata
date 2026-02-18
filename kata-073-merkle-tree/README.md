# Kata 073 — Merkle Tree

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata73() error
```

### Learning goal
- What you are practicing: Build `Merkle Tree` with careful validation and exact byte-level behavior for security-sensitive inputs/outputs.
- Why it matters: You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.
- How this grows your Go skills: This builds your discipline for correctness-first coding in high-risk code paths.
- When correct: When your solution is correct, it should satisfy: `follow README spec`; `write tests`; and `keep it idiomatic`.

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
- This kata is focused practice in real-world Go design and testing through `Merkle Tree`.
- You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
