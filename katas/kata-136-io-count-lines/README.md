# Kata 136 — Count Lines

**Focus:** bufio.Scanner, io.Reader

## Your task
Implement:

```go
func CountLines(r io.Reader) (int, error)
```

### Learning goal
- What you are building: line-counting logic with explicit reader error handling.
- Why this matters in real projects: stream boundaries and scanner errors are common in logs and ingestion services.
- How this grows your Go skills: you practice reader-based processing with deterministic output contracts.
- Definition of done (plain English): behavior matches all expectations below.

## Rules / Expectations
- nil reader => non-nil error
- count logical lines from reader
- empty input => 0 and nil error
- scanner error => propagate error

## Prior reading
- [bufio package](https://pkg.go.dev/bufio)
- [io package](https://pkg.go.dev/io)
- [Go package documentation index](https://pkg.go.dev/std)

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
