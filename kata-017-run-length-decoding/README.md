# Kata 017 — Run-Length Decoding

**Focus:** Parsing, errors

## Your task
Implement:

```go
func RLDecode(s string) (string, error)
```

### Learning goal
- Expected work: Implement `Run-Length Decoding` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Run-Length Decoding` teaches core implementation habits that compound across all later katas.

## Rules / Expectations
- invalid format => error
- count can be multiple digits
- roundtrip with encode for valid inputs

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains parsing, errors by implementing `Run-Length Decoding` under explicit constraints.
- It is important because `Run-Length Decoding` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
