# Kata 018 — Caesar Cipher

**Focus:** ASCII letters shifting

## Your task
Implement:

```go
func Caesar(s string, shift int) string
```

### Learning goal
- What you are practicing: Build `Caesar Cipher` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `only shift A-Z and a-z`; `preserve case`; and `non-letters unchanged`.

## Rules / Expectations
- only shift A-Z and a-z
- preserve case
- non-letters unchanged

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in ASCII letters shifting through `Caesar Cipher`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
