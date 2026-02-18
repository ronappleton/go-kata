# Kata 045 — Worker Pool

**Focus:** Concurrency patterns

## Your task
Implement:

```go
func RunPool(ctx context.Context, workers int, jobs <-chan Job) <-chan Result
```

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
- This kata is about implementing Worker Pool with constraints that make you practice concurrency patterns.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
