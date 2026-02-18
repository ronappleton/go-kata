# Kata 053 — Line-Oriented Reader

**Focus:** bufio

## Your task
Implement:

```go
func ReadLines(r io.Reader, maxLine int) ([]string, error)
```

### Learning goal
- What you are practicing: Build `Line-Oriented Reader` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `error if any line exceeds maxLine`; `keep line content only`; and `handles last line`.

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
- This kata is focused practice in bufio through `Line-Oriented Reader`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
