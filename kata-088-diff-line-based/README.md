# Kata 088 — Diff (line-based)

**Focus:** Algorithms

## Your task
Implement:

```go
func DiffLines(a, b []string) []DiffOp
```

### Learning goal
- What you are practicing: Build `Diff (line-based)` by holding algorithm invariants and edge cases, not just passing easy examples.
- Why it matters: You will use this when performance matters and correctness must hold at larger input sizes.
- How this grows your Go skills: This builds complexity awareness and confidence in proving behavior with tests.
- When correct: When your solution is correct, it should satisfy: `LCS-based diff`; `ops: Add/Del/Keep`; and `deterministic`.

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
- This kata is focused practice in Algorithms through `Diff (line-based)`.
- You will use this when performance matters and correctness must hold at larger input sizes.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
