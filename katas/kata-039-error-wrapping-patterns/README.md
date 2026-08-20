# Kata 039 — Error Wrapping Patterns

**Focus:** error wrapping, fmt.Errorf, %w, errors.Is, errors.As

## Your task
Implement:

```go
func FindUser(id int) (string, error)
```

### Learning goal
- What you are practicing: wrapping errors with context while preserving the ability to check error types.
- Why this matters: in real systems, you need to add context to errors ("failed to fetch user 123") while still being able to check the root cause.
- How this grows your Go skills: you learn the `%w` verb, `errors.Is`, and `errors.As` — the modern Go error handling toolkit.

### Tips
- Use `fmt.Errorf("context: %w", err)` to wrap errors.
- Use `errors.Is(err, ErrNotFound)` to check error identity.
- Use `errors.As(err, &target)` to extract error types.

## Rules / Expectations
- existing user => returns name, nil error
- missing user => returns ErrNotFound

## What this kata is about (and why it matters)
- Core lesson: wrap errors with context, check with errors.Is/As.
- After this kata, you should never lose error context in your code.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
