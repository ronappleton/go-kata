# Kata 032 — Top N Words

**Focus:** Sorting, ties

## Your task
Implement:

```go
func TopNWords(s string, n int) []string
```

### Learning goal
- What you are building: Build `func TopNWords(s string, n int) []string` as a reliable contract. Focus: Sorting, ties.
- Why this matters in real projects: This is how you keep features fast when input size grows.
- How this grows your Go skills: You practice invariants and complexity reasoning, then prove both with tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: lowercase, split whitespace; ties sorted alphabetically; and if `n<=0`, return `empty slice`.

### Tips
- State the invariant first, then code.
- Test shape edges early: empty, one item, duplicates.
- Check complexity after correctness.

## Rules / Expectations
- lowercase, split whitespace
- ties sorted alphabetically
- n<=0 => empty slice

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
