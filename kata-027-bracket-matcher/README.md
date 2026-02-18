# Kata 027 — Bracket Matcher

**Focus:** Stacks, runes

## Your task
Implement:

```go
func IsBalanced(s string) bool
```

### Learning goal
- Expected work: Implement `Bracket Matcher` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Bracket Matcher` teaches core implementation habits that compound across all later katas.
- When correct: `Bracket Matcher` should satisfy the required behavior, including: `supports (), {}, []`; `ignores other chars`; and `empty => true`.

## Rules / Expectations
- supports (), {}, []
- ignores other chars
- empty => true

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains stacks, runes by implementing `Bracket Matcher` under explicit constraints.
- It is important because `Bracket Matcher` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
