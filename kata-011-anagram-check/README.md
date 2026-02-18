# Kata 011 — Anagram Check

**Focus:** Rune counts, normalization

## Your task
Implement:

```go
func IsAnagram(a, b string) bool
```

### Learning goal
- What you are practicing: Build `Anagram Check` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `case-insensitive`; `ignore spaces/punctuation`; and `Unicode-safe via rune counts`.

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
- This kata is focused practice in Rune counts, normalization through `Anagram Check`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
