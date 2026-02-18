# Kata 011 — Anagram Check

**Focus:** Rune counts, normalization

## Your task
Implement:

```go
func IsAnagram(a, b string) bool
```

### Learning goal
- Expected work: Implement `Anagram Check` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Anagram Check` teaches core implementation habits that compound across all later katas.

## Rules / Expectations
- case-insensitive
- ignore spaces/punctuation
- Unicode-safe via rune counts

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains rune counts, normalization by implementing `Anagram Check` under explicit constraints.
- It is important because `Anagram Check` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
