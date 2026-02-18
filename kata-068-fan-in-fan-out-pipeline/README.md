# Kata 068 — Fan-in/Fan-out Pipeline

**Focus:** Channels, composition

## Your task
Implement:

```go
func Pipeline[T any](ctx context.Context, in <-chan T, stages ...func(context.Context, <-chan T) <-chan T) <-chan T
```

### Learning goal
- Expected work: Implement `Fan-in/Fan-out Pipeline` with clear coordination so concurrent work finishes without races, deadlocks, or goroutine leaks.
- Why: `Fan-in/Fan-out Pipeline` teaches safe coordination patterns that prevent costly production race conditions.
- When correct: `Fan-in/Fan-out Pipeline` should satisfy the required behavior, including: `compose stages in order`; `cancel cleanly`; and `close outputs`.

## Rules / Expectations
- compose stages in order
- cancel cleanly
- close outputs

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains channels, composition by implementing `Fan-in/Fan-out Pipeline` under explicit constraints.
- It is important because `Fan-in/Fan-out Pipeline` builds the synchronization discipline needed for safe parallel systems.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
