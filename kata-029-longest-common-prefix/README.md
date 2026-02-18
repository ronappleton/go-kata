# Kata 029 — Longest Common Prefix

**Focus:** Strings

## Your task
Implement:

```go
func LongestCommonPrefix(items []string) string
```

### Learning goal
- What you are practicing: Build `Longest Common Prefix` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `empty => empty`; `single => itself`; and `case-sensitive`.

## Rules / Expectations
- empty => empty
- single => itself
- case-sensitive

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Strings through `Longest Common Prefix`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
