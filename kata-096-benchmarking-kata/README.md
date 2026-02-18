# Kata 096 — Benchmarking Kata

**Focus:** testing/benchmark

## Your task
Implement:

```go
func Concat(parts []string) string
```

### Learning goal
- Expected work: Implement `Benchmarking Kata` as a composable abstraction with explicit contracts and testable behavior.
- Why: `Benchmarking Kata` teaches maintainable design through explicit contracts and low coupling.
- When correct: `Benchmarking Kata` should satisfy the required behavior, including: `write 2 benchmarks: naive vs builder`; `compare allocations`; and `document results in README`.

## Rules / Expectations
- write 2 benchmarks: naive vs builder
- compare allocations
- document results in README

## Prior reading
- [Go testing package (benchmarks)](https://pkg.go.dev/testing)
- [Go blog: Profiling Go Programs](https://go.dev/blog/pprof)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains testing/benchmark by implementing `Benchmarking Kata` under explicit constraints.
- It is important because `Benchmarking Kata` develops design choices that improve extensibility and testability over time.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
