# Kata 086 — File Integrity Checker

**Focus:** io, crypto/sha256

## Your task
Implement:

```go
func SHA256File(path string) (string, error)
```

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
- This kata is about implementing File Integrity Checker with constraints that make you practice io, crypto/sha256.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
