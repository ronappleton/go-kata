# Kata 125 — pprof Profiling

**Focus:** runtime/pprof, profiling, CPU/memory analysis, benchmarks

## Your task
Implement FindDuplicates and write a benchmark for it.

### Learning goal
- What you are practicing: writing benchmarks and using profiling to find performance bottlenecks.
- Why this matters: "make it work, then make it fast" — profiling tells you where to optimize.
- How this grows your Go skills: you learn to use `go test -bench`, `go tool pprof`, and benchmark-driven optimization.

### Tips
- Write benchmarks with `func BenchmarkXxx(b *testing.B)`.
- Use `b.ResetTimer()` after setup.
- Run `go test -bench=. -benchmem` for memory allocation info.
- Use `go tool pprof` to analyze CPU/memory profiles.

## Rules / Expectations
- FindDuplicates returns duplicate values
- Benchmark runs without errors
- Code is efficient (O(n) expected)

## What this kata is about (and why it matters)
- Core lesson: measure before optimizing. Profiling finds real bottlenecks.
- After this kata, you'll benchmark before claiming code is "fast enough."

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
go test -bench=. -benchmem ./...
```
