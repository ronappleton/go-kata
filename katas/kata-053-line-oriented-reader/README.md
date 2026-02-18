# Kata 053 — Line-Oriented Reader

**Focus:** bufio

## Your task
Implement:

```go
func ReadLines(r io.Reader, maxLine int) ([]string, error)
```

### Learning goal
- What you are building: Build `func ReadLines(r io.Reader, maxLine int) ([]string, error)` as a reliable contract. Focus: bufio.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: error if any line exceeds maxline; accept only `keep line content`; and handles last line.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

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
- Core lesson: parse strictly and fail loudly on bad input.
- After this kata, you should be able to state input contracts and return precise parse errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
