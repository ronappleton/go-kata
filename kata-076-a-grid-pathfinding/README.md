# Kata 076 — A* Grid Pathfinding

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata76() error
```

### Learning goal
- Expected work: Implement `A* Grid Pathfinding` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `A* Grid Pathfinding` teaches complexity-aware correctness that impacts throughput and latency at scale.

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
- This kata trains integration design, boundary handling, and robust testing by implementing `A* Grid Pathfinding` under explicit constraints.
- It is important because `A* Grid Pathfinding` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
