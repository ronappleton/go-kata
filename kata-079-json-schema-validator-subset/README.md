# Kata 079 — JSON Schema Validator (subset)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata79() error
```

### Learning goal
- Expected work: Implement `JSON Schema Validator (subset)` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `JSON Schema Validator (subset)` teaches strict boundary handling so malformed input cannot silently corrupt results.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `JSON Schema Validator (subset)` under explicit constraints.
- It is important because `JSON Schema Validator (subset)` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
