# Kata 068 — Fan-in/Fan-out Pipeline

**Focus:** Channels, composition

## Your task
Implement:

```go
func Pipeline[T any](ctx context.Context, in <-chan T, stages ...func(context.Context, <-chan T) <-chan T) <-chan T
```

### Learning goal
- What you are practicing: Build `Fan-in/Fan-out Pipeline` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `compose stages in order`; `cancel cleanly`; and `close outputs`.

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
- This kata is focused practice in Channels, composition through `Fan-in/Fan-out Pipeline`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
