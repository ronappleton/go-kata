# Kata 100 — LRU Cache (thread-safe)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata70() error
```

### Learning goal
- What you are building: Build `func Kata70() error` as a reliable contract. Focus: real-world Go integration and testing.
- Why this matters in real projects: Concurrency bugs are expensive. You are learning to prevent them by design.
- How this grows your Go skills: You practice ownership, cancellation, synchronization, and leak-free shutdown.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: match the full README contract; include tests for happy path and edge cases; and keep API and error style idiomatic Go.

### Tips
- Decide ownership first: who starts, stops, and closes.
- Test cancellation and shutdown before throughput.
- Run `go test -race ./...` regularly.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)

## What this kata is about (and why it matters)
- Core lesson: own lifecycle and shutdown before chasing throughput.
- After this kata, you should be able to explain who starts, who stops, and who closes every path.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
