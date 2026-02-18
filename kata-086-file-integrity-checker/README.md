# Kata 086 — File Integrity Checker

**Focus:** io, crypto/sha256

## Your task
Implement:

```go
func SHA256File(path string) (string, error)
```

### Learning goal
- Expected work: Implement `File Integrity Checker` with strict validation and byte-level correctness for all cryptographic operations.
- Why: `File Integrity Checker` teaches security precision, where small mistakes can become real vulnerabilities.
- When correct: `File Integrity Checker` should satisfy the required behavior, including: `stream file`; `hex digest`; and `error on read`.

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
- This kata trains io, crypto/sha256 by implementing `File Integrity Checker` under explicit constraints.
- It is important because `File Integrity Checker` reflects production security code, where correctness mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
