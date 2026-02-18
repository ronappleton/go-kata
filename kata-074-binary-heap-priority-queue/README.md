# Kata 074 — Binary Heap Priority Queue

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata74() error
```

### Learning goal
- Expected work: Implement `Binary Heap Priority Queue` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Binary Heap Priority Queue` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go container/heap package](https://pkg.go.dev/container/heap)
- [Go sort package](https://pkg.go.dev/sort)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `Binary Heap Priority Queue` under explicit constraints.
- It is important because `Binary Heap Priority Queue` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
