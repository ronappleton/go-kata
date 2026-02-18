# Kata 029 — Increment Concurrently

**Focus:** sync.WaitGroup, sync.Mutex, goroutine coordination

## Your task
Implement:

```go
func IncrementConcurrently(total, workers int) int
```

### Learning goal
- What you are building: a concurrency-safe shared counter increment routine.
- Why this matters in real projects: race conditions often hide in seemingly simple counter/state updates.
- How this grows your Go skills: you practice coordination and safe shared-state mutation.
- Definition of done (plain English): final counter is exact and deterministic across repeated runs.

### Tips
- Make worker partitioning explicit.
- Keep critical sections minimal.
- Test edge inputs and repeated invocations.

## Rules / Expectations
- total <= 0 => return 0
- workers <= 0 => treat as 1 worker
- increment shared counter exactly total times across workers
- return exact final counter value

## Prior reading
- [sync package](https://pkg.go.dev/sync)
- [sync/atomic package](https://pkg.go.dev/sync/atomic)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: concurrent correctness is about ownership and synchronization, not speed first.
- After this kata, you should be able to justify your synchronization choice in code review.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
