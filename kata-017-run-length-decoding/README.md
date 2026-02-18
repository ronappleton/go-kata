# Kata 017 — Run-Length Decoding

**Focus:** Parsing, errors

## Your task
Implement:

```go
func RLDecode(s string) (string, error)
```

### Learning goal
- What you are practicing: Build `Run-Length Decoding` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `invalid format => error`; `count can be multiple digits`; and `roundtrip with encode for valid inputs`.

## Rules / Expectations
- invalid format => error
- count can be multiple digits
- roundtrip with encode for valid inputs

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Parsing, errors through `Run-Length Decoding`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
