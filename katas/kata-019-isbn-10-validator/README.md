# Kata 019 — ISBN-10 Validator

**Focus:** Checksums

## Your task
Implement:

```go
func IsValidISBN10(s string) bool
```

### Learning goal
- What you are building: Build `func IsValidISBN10(s string) bool` as a reliable contract. Focus: Checksums.
- Why this matters in real projects: This is how you keep features fast when input size grows.
- How this grows your Go skills: You practice invariants and complexity reasoning, then prove both with tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: allow hyphens/spaces; x allowed as last digit meaning 10; and must be 10 digits after stripping.

### Tips
- State the invariant first, then code.
- Test shape edges early: empty, one item, duplicates.
- Check complexity after correctness.

## Rules / Expectations
- allow hyphens/spaces
- X allowed as last digit meaning 10
- must be 10 digits after stripping

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- Core lesson: hold invariants first, then optimize safely.
- After this kata, you should be able to defend algorithm choice, complexity, and corner-case behavior.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
