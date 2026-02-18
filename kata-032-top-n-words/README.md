# Kata 032 — Top N Words

**Focus:** Sorting, ties

## Your task
Implement:

```go
func TopNWords(s string, n int) []string
```

### Learning goal
- What you are practicing: Build `Top N Words` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `lowercase, split whitespace`; `ties sorted alphabetically`; and `n<=0 => empty slice`.

## Rules / Expectations
- lowercase, split whitespace
- ties sorted alphabetically
- n<=0 => empty slice

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Sorting, ties through `Top N Words`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
