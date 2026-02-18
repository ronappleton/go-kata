# Kata 003 — Reverse String

**Focus:** Strings, runes vs bytes

## Your task
Implement:

```go
func Reverse(s string) string
```

### Learning goal
- Expected work: Implement `Reverse String` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Reverse String` teaches core implementation habits that compound across all later katas.
- When correct: `Reverse String` should satisfy the required behavior, including: `must handle Unicode correctly (runes)`; `empty => empty`; and `palindrome unchanged`.

## Rules / Expectations
- must handle Unicode correctly (runes)
- empty => empty
- palindrome unchanged

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains strings, runes vs bytes by implementing `Reverse String` under explicit constraints.
- It is important because `Reverse String` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
