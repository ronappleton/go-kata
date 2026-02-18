# Kata 054 — JSON Lines Filter

**Focus:** Stream processing

## Your task
Implement:

```go
func FilterJSONL(r io.Reader, w io.Writer, predicate func(map[string]any) bool) error
```

## Rules / Expectations
- each line JSON object
- write only matches
- skip blank lines

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)

## What this kata is about (and why it matters)
- This kata is about implementing JSON Lines Filter with constraints that make you practice stream processing.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
