# Kata 042 — Context Timeout Wrapper

**Focus:** context

## Your task
Implement:

```go
func WithTimeoutDo(ctx context.Context, d time.Duration, fn func(context.Context) error) error
```

## Rules / Expectations
- fn called with derived ctx
- returns ctx error on timeout
- propagates fn error

## Prior reading
- [Go context package](https://pkg.go.dev/context)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Context Timeout Wrapper with constraints that make you practice context.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
