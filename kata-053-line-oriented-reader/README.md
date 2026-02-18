# Kata 053 — Line-Oriented Reader

**Focus:** bufio

## Your task
Implement:

```go
func ReadLines(r io.Reader, maxLine int) ([]string, error)
```

### Learning goal
- Expected work: Implement `Line-Oriented Reader` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `Line-Oriented Reader` teaches strict boundary handling so malformed input cannot silently corrupt results.

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
- This kata trains bufio by implementing `Line-Oriented Reader` under explicit constraints.
- It is important because `Line-Oriented Reader` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
