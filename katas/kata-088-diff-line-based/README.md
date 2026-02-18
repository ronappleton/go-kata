# Kata 088 — Diff (line-based)

**Focus:** Algorithms

## Your task
Implement:

```go
func DiffLines(a, b []string) []DiffOp
```

### Learning goal
- What you are building: Build `func DiffLines(a, b []string) []DiffOp` as a reliable contract. Focus: Algorithms.
- Why this matters in real projects: This is how you keep features fast when input size grows.
- How this grows your Go skills: You practice invariants and complexity reasoning, then prove both with tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: lcs-based diff; ops: add/del/keep; and deterministic.

### Tips
- State the invariant first, then code.
- Test shape edges early: empty, one item, duplicates.
- Check complexity after correctness.

## Rules / Expectations
- LCS-based diff
- ops: Add/Del/Keep
- deterministic

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
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
