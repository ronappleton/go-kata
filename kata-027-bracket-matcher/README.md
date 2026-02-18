# Kata 027 — Bracket Matcher

**Focus:** Stacks, runes

## Your task
Implement:

```go
func IsBalanced(s string) bool
```

### Learning goal
- What you are practicing: Build `Bracket Matcher` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `supports (), {}, []`; `ignores other chars`; and `empty => true`.

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
- This kata is focused practice in Stacks, runes through `Bracket Matcher`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
