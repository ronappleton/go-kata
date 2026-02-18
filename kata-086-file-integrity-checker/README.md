# Kata 086 — File Integrity Checker

**Focus:** io, crypto/sha256

## Your task
Implement:

```go
func SHA256File(path string) (string, error)
```

### Learning goal
- What you are building: Build `func SHA256File(path string) (string, error)` as a reliable contract. Focus: io, crypto/sha256.
- Why this matters in real projects: Security paths have no soft failures. Exact behavior is non-negotiable.
- How this grows your Go skills: You practice fail-closed validation and byte-accurate handling.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: stream file; hex digest; and error on read.

### Tips
- Fail closed on malformed data.
- Never ignore crypto or encoding errors.
- Add tamper and expiry negative tests.

## Rules / Expectations
- stream file
- hex digest
- error on read

## Prior reading
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)
- [Go os package](https://pkg.go.dev/os)
- [Go path/filepath package](https://pkg.go.dev/path/filepath)
- [Go crypto/sha256 package](https://pkg.go.dev/crypto/sha256)

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
