# Kata 034 — Temperature Converter

**Focus:** Parsing floats

## Your task
Implement:

```go
func ConvertTemp(input string) (string, error)
```

### Learning goal
- What you are practicing: Build `Temperature Converter` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `supports C/F/K suffix`; `output all 3 units with 2 decimals`; and `invalid => error`.

## Rules / Expectations
- supports C/F/K suffix
- output all 3 units with 2 decimals
- invalid => error

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Parsing floats through `Temperature Converter`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
