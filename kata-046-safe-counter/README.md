# Kata 046 — Safe Counter

**Focus:** sync/atomic or mutex

## Your task
Implement:

```go
type Counter struct
```

### Learning goal
- What you are building: Build `Safe Counter` as a reliable contract. Focus: sync/atomic or mutex.
- Why this matters in real projects: Concurrency bugs are expensive. You are learning to prevent them by design.
- How this grows your Go skills: You practice ownership, cancellation, synchronization, and leak-free shutdown.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: inc/add/value; and race-free.

### Tips
- Decide ownership first: who starts, stops, and closes.
- Test cancellation and shutdown before throughput.
- Run `go test -race ./...` regularly.

## Rules / Expectations
- Inc/Add/Value
- race-free

## Prior reading
- [Go sync/atomic package](https://pkg.go.dev/sync/atomic)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

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
