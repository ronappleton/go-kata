# Kata 045 — Worker Pool

**Focus:** Concurrency patterns

## Your task
Implement:

```go
func RunPool(ctx context.Context, workers int, jobs <-chan Job) <-chan Result
```

### Learning goal
- Expected work: Implement `Worker Pool` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Worker Pool` teaches safe coordination patterns that prevent costly production race conditions.
- When correct: `Worker Pool` should satisfy the required behavior, including: `cancels on ctx`; `close output when done`; and `no leaks`.

## Rules / Expectations
- cancels on ctx
- close output when done
- no leaks

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains concurrency patterns by implementing `Worker Pool` under explicit constraints.
- It is important because `Worker Pool` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
