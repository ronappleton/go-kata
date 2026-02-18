# Kata 035 — Log Line Parser

**Focus:** Key=value parsing

## Your task
Implement:

```go
func ParseLogLine(line string) (map[string]string, error)
```

## Rules / Expectations
- key=value pairs separated by spaces
- values may be quoted
- duplicate keys last wins

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Log Line Parser with constraints that make you practice key=value parsing.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
