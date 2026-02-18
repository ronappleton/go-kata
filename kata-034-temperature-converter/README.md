# Kata 034 — Temperature Converter

**Focus:** Parsing floats

## Your task
Implement:

```go
func ConvertTemp(input string) (string, error)
```

### Learning goal
- Expected work: Implement `Temperature Converter` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Temperature Converter` teaches core implementation habits that compound across all later katas.
- When correct: `Temperature Converter` should satisfy the required behavior, including: `supports C/F/K suffix`; `output all 3 units with 2 decimals`; and `invalid => error`.

## Rules / Expectations
- supports C/F/K suffix
- output all 3 units with 2 decimals
- invalid => error

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains parsing floats by implementing `Temperature Converter` under explicit constraints.
- It is important because `Temperature Converter` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
