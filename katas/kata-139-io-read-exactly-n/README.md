# Kata 139 — Read Exactly N

**Focus:** io.ReadFull semantics

## Your task
Implement:

```go
func ReadExactlyN(r io.Reader, n int) ([]byte, error)
```

### Learning goal
- What you are building: exact-length read behavior with short-read error handling.
- Why this matters in real projects: protocol framing and binary processing often require precise byte counts.
- How this grows your Go skills: you practice length contracts and read failure handling.
- Definition of done (plain English): behavior matches all expectations below.

## Rules / Expectations
- n < 0 => non-nil error
- n == 0 => empty slice and nil error
- return exactly n bytes or error
- short read => non-nil error

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
