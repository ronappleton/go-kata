# Kata 096 — Benchmarking Kata

**Focus:** testing/benchmark

## Your task
Implement:

```go
func Concat(parts []string) string
```

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
- This kata is about implementing Benchmarking Kata with constraints that make you practice testing/benchmark.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
