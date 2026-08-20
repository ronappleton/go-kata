# Kata 034 — Palindrome Check

**Focus:** Strings, normalization basics

## Your task
Implement:

```go
func IsPalindrome(s string) bool
```

### Learning goal
- What you are building: Build `func IsPalindrome(s string) bool` as a reliable contract. Focus: Strings, normalization basics.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: compare case-insensitively; ignore non-letters/digits; and unicode letters treated as letters (use unicode package).

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

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
- Core lesson: turn plain rules into deterministic Go behavior.
- After this kata, you should be able to write rule-first tests and explain each edge case clearly.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
