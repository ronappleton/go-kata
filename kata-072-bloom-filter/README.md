# Kata 072 — Bloom Filter

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata72() error
```

### Learning goal
- Expected work: Implement `Bloom Filter` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Bloom Filter` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Bloom filter overview](https://en.wikipedia.org/wiki/Bloom_filter)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `Bloom Filter` under explicit constraints.
- It is important because `Bloom Filter` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
