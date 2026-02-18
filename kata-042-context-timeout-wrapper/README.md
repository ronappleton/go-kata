# Kata 042 — Context Timeout Wrapper

**Focus:** context

## Your task
Implement:

```go
func WithTimeoutDo(ctx context.Context, d time.Duration, fn func(context.Context) error) error
```

### Learning goal
- Expected work: Implement `Context Timeout Wrapper` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Context Timeout Wrapper` teaches core implementation habits that compound across all later katas.

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
- This kata trains context by implementing `Context Timeout Wrapper` under explicit constraints.
- It is important because `Context Timeout Wrapper` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
