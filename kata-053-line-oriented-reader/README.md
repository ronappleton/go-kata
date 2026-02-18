# Kata 053 — Line-Oriented Reader

**Focus:** bufio

## Your task
Implement:

```go
func ReadLines(r io.Reader, maxLine int) ([]string, error)
```

## Rules / Expectations
- error if any line exceeds maxLine
- keep line content only
- handles last line

## Prior reading
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Line-Oriented Reader with constraints that make you practice bufio.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
