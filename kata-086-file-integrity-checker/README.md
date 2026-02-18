# Kata 086 — File Integrity Checker

**Focus:** io, crypto/sha256

## Your task
Implement:

```go
func SHA256File(path string) (string, error)
```

### Learning goal
- What you are practicing: Build `File Integrity Checker` with careful validation and exact byte-level behavior for security-sensitive inputs/outputs.
- Why it matters: You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.
- How this grows your Go skills: This builds your discipline for correctness-first coding in high-risk code paths.
- When correct: When your solution is correct, it should satisfy: `stream file`; `hex digest`; and `error on read`.

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
- This kata is focused practice in io, crypto/sha256 through `File Integrity Checker`.
- You will use this in authentication, signing, and integrity checks where small mistakes can become incidents.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
