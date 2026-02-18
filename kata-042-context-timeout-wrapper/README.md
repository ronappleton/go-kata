# Kata 042 — Context Timeout Wrapper

**Focus:** context

## Your task
Implement:

```go
func WithTimeoutDo(ctx context.Context, d time.Duration, fn func(context.Context) error) error
```

### Learning goal
- What you are building: Build `func WithTimeoutDo(ctx context.Context, d time.Duration, fn func(context.Context) error) error` as a reliable contract. Focus: context.
- Why this matters in real projects: This is everyday Go work: small rules, clear behavior, zero surprises.
- How this grows your Go skills: You practice zero-value handling, explicit branching, and table-driven tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: fn called with derived ctx; returns ctx error on timeout; and propagates fn error.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

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
- Core lesson: turn plain rules into deterministic Go behavior.
- After this kata, you should be able to write rule-first tests and explain each edge case clearly.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
