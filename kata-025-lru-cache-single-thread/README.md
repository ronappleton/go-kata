# Kata 025 — LRU Cache (single-thread)

**Focus:** Structs, maps, list

## Your task
Implement:

```go
func NewLRU(capacity int) (*LRU, error)
```

### Learning goal
- Expected work: Implement `LRU Cache (single-thread)` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `LRU Cache (single-thread)` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- capacity>0
- Get/Put methods
- evict least recently used

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains structs, maps, list by implementing `LRU Cache (single-thread)` under explicit constraints.
- It is important because `LRU Cache (single-thread)` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
