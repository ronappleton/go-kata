# Kata 018 — Caesar Cipher

**Focus:** ASCII letters shifting

## Your task
Implement:

```go
func Caesar(s string, shift int) string
```

### Learning goal
- Expected work: Implement `Caesar Cipher` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Caesar Cipher` teaches core implementation habits that compound across all later katas.
- When correct: `Caesar Cipher` should satisfy the required behavior, including: `only shift A-Z and a-z`; `preserve case`; and `non-letters unchanged`.

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
- This kata trains ascii letters shifting by implementing `Caesar Cipher` under explicit constraints.
- It is important because `Caesar Cipher` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
