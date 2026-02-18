# Kata 020 — Roman Numerals

**Focus:** Greedy mapping

## Your task
Implement:

```go
func ToRoman(n int) (string, error)
```

### Learning goal
- What you are practicing: Build `Roman Numerals` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `1..3999 only`; `invalid => error`; and `use standard subtractive notation`.

## Rules / Expectations
- 1..3999 only
- invalid => error
- use standard subtractive notation

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Greedy mapping through `Roman Numerals`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
