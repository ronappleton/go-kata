# Kata 088 — Diff (line-based)

**Focus:** Algorithms

## Your task
Implement:

```go
func DiffLines(a, b []string) []DiffOp
```

### Learning goal
- Expected work: Implement `Diff (line-based)` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Diff (line-based)` teaches complexity-aware correctness that impacts throughput and latency at scale.

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
- This kata trains algorithms by implementing `Diff (line-based)` under explicit constraints.
- It is important because `Diff (line-based)` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
