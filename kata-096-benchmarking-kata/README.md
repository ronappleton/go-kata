# Kata 096 — Benchmarking Kata

**Focus:** testing/benchmark

## Your task
Implement:

```go
func Concat(parts []string) string
```

### Learning goal
- What you are building: Build `func Concat(parts []string) string` as a reliable contract. Focus: testing/benchmark.
- Why this matters in real projects: Good structure compounds. Clear contracts reduce future rework.
- How this grows your Go skills: You practice interface-first design and explicit dependencies.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: write 2 benchmarks: naive vs builder; compare allocations; and document results in readme.

### Tips
- Write contract tests before implementation details.
- Keep dependencies explicit, not implicit.
- Prefer small interfaces with one reason to change.

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
- Core lesson: design seams for testability and change.
- After this kata, you should be able to describe component contracts and test at those seams.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
