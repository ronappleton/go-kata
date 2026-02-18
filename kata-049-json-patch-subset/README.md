# Kata 049 — JSON Patch (subset)

**Focus:** encoding/json

## Your task
Implement:

```go
func ApplyPatch(doc []byte, patch []byte) ([]byte, error)
```

### Learning goal
- Expected work: Implement `JSON Patch (subset)` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `JSON Patch (subset)` teaches strict boundary handling so malformed input cannot silently corrupt results.
- When correct: `JSON Patch (subset)` should satisfy the required behavior, including: `support add/replace/remove at top-level keys only`; and `invalid => error`.

## Rules / Expectations
- support add/replace/remove at top-level keys only
- invalid => error

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains encoding/json by implementing `JSON Patch (subset)` under explicit constraints.
- It is important because `JSON Patch (subset)` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
