# Kata 075 — Dijkstra Shortest Path

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata75() error
```

### Learning goal
- Expected work: Implement `Dijkstra Shortest Path` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Dijkstra Shortest Path` teaches complexity-aware correctness that impacts throughput and latency at scale.
- When correct: `Dijkstra Shortest Path` should satisfy the required behavior, including: `follow README spec`; `write tests`; and `keep it idiomatic`.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `Dijkstra Shortest Path` under explicit constraints.
- It is important because `Dijkstra Shortest Path` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
