# Kata 042 — Context Timeout Wrapper

**Focus:** context

## Your task
Implement:

```go
func WithTimeoutDo(ctx context.Context, d time.Duration, fn func(context.Context) error) error
```

### Learning goal
- What you are practicing: Build `Context Timeout Wrapper` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `fn called with derived ctx`; `returns ctx error on timeout`; and `propagates fn error`.

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
- This kata is focused practice in context through `Context Timeout Wrapper`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
