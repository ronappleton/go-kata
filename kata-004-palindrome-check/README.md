# Kata 004 — Palindrome Check

**Focus:** Strings, normalization basics

## Your task
Implement:

```go
func IsPalindrome(s string) bool
```

### Learning goal
- Expected work: Implement `Palindrome Check` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Palindrome Check` teaches core implementation habits that compound across all later katas.
- When correct: `Palindrome Check` should satisfy the required behavior, including: `case-insensitive`; `ignore non-letters/digits`; and `Unicode letters treated as letters (use unicode package)`.

## Rules / Expectations
- case-insensitive
- ignore non-letters/digits
- Unicode letters treated as letters (use unicode package)

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains strings, normalization basics by implementing `Palindrome Check` under explicit constraints.
- It is important because `Palindrome Check` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
