# Kata 096 — Benchmarking Kata

**Focus:** testing/benchmark

## Your task
Implement:

```go
func Concat(parts []string) string
```

### Learning goal
- What you are practicing: Build `Benchmarking Kata` as a reusable, testable abstraction with explicit contracts.
- Why it matters: You will use this when designing packages that need to stay maintainable as features grow.
- How this grows your Go skills: This builds API design judgment and composable package structure in Go.
- When correct: When your solution is correct, it should satisfy: `write 2 benchmarks: naive vs builder`; `compare allocations`; and `document results in README`.

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
- This kata is focused practice in testing/benchmark through `Benchmarking Kata`.
- You will use this when designing packages that need to stay maintainable as features grow.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
