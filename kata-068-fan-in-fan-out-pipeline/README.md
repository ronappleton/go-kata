# Kata 068 — Fan-in/Fan-out Pipeline

**Focus:** Channels, composition

## Your task
Implement:

```go
func Pipeline[T any](ctx context.Context, in <-chan T, stages ...func(context.Context, <-chan T) <-chan T) <-chan T
```

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
- This kata is about implementing Fan-in/Fan-out Pipeline with constraints that make you practice channels, composition.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
