# Kata 126 — Benchmarking Deep Dive

**Focus:** benchmark patterns, sub-benchmarks, comparison, regression

## Your task
Implement ConcatStringsBuilder and compare performance with the loop version.

### Learning goal
- What you are practicing: comparing algorithm implementations with benchmarks.
- Why this matters: choosing the right implementation can mean 10x performance differences.
- How this grows your Go skills: you learn sub-benchmarks, benchmark comparison, and performance regression testing.

### Tips
- Use sub-benchmarks (`b.Run`) to compare implementations.
- Run `go test -bench=. -benchmem` to see allocation differences.
- Use `benchstat` to compare benchmark results statistically.

## Rules / Expectations
- ConcatStringsBuilder produces correct output
- Benchmark compares both implementations
- Builder version should be more efficient

## What this kata is about (and why it matters)
- Core lesson: benchmarks prove which implementation is actually faster.
- After this kata, you'll benchmark before choosing algorithms.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test -bench=. -benchmem ./...
```
