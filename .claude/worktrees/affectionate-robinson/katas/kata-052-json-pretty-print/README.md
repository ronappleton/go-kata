# Kata 052 — JSON Pretty Print

**Focus:** encoding/json

## Your task
Implement:

```go
func PrettyJSON(input []byte) ([]byte, error)
```

### Learning goal
- What you are building: Build `func PrettyJSON(input []byte) ([]byte, error)` as a reliable contract. Focus: encoding/json.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: when `invalid json`, return `error`; indent 2 spaces; and stable formatting.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

## Rules / Expectations
- invalid json => error
- indent 2 spaces
- stable formatting

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
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
