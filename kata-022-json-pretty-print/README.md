# Kata 022 — JSON Pretty Print

**Focus:** encoding/json

## Your task
Implement:

```go
func PrettyJSON(input []byte) ([]byte, error)
```

### Learning goal
- Expected work: Implement `JSON Pretty Print` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `JSON Pretty Print` teaches strict boundary handling so malformed input cannot silently corrupt results.

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
- This kata trains encoding/json by implementing `JSON Pretty Print` under explicit constraints.
- It is important because `JSON Pretty Print` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
