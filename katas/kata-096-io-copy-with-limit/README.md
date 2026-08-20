# Kata 137 — Copy With Limit

**Focus:** io.CopyN, boundary validation

## Your task
Implement:

```go
func CopyWithLimit(dst io.Writer, src io.Reader, limit int64) (int64, error)
```

### Learning goal
- What you are building: bounded copy behavior for safe stream transfers.
- Why this matters in real projects: unbounded copies can cause resource exhaustion and unpredictable runtime behavior.
- How this grows your Go skills: you practice transfer limits and explicit error paths.
- Definition of done (plain English): behavior matches all expectations below.

## Rules / Expectations
- limit <= 0 => return 0 and nil error
- copy at most limit bytes
- return bytes copied and any copy error
- nil src/dst => non-nil error

## Prior reading
- [io package](https://pkg.go.dev/io)
- [Go package documentation index](https://pkg.go.dev/std)

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
