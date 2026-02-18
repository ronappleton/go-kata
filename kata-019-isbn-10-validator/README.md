# Kata 019 — ISBN-10 Validator

**Focus:** Checksums

## Your task
Implement:

```go
func IsValidISBN10(s string) bool
```

### Learning goal
- What you are practicing: Build `ISBN-10 Validator` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `allow hyphens/spaces`; `X allowed as last digit meaning 10`; and `must be 10 digits after stripping`.

## Rules / Expectations
- allow hyphens/spaces
- X allowed as last digit meaning 10
- must be 10 digits after stripping

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Checksums through `ISBN-10 Validator`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
