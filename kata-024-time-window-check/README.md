# Kata 024 — Time Window Check

**Focus:** time package

## Your task
Implement:

```go
func WithinWindow(t, start, end time.Time) bool
```

## Rules / Expectations
- inclusive start, exclusive end
- if start after end => false

## Prior reading
- [Go time package](https://pkg.go.dev/time)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Time Window Check with constraints that make you practice time package.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
