# Kata 116 — JSON Schema Validator (subset)

**Focus:** encoding/json, validation

## Your task
Implement:

```go
func ValidateJSON(doc, schema []byte) error
```

### Learning goal
- What you are building: func ValidateJSON(doc, schema []byte) error as a reliable contract. Focus: encoding/json, validation.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): validates type, required properties, and property types; returns a descriptive error for the first violation; nil when valid.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- support type/properties/required keywords
- validate nested objects
- error on first violation
- invalid JSON input => non-nil error

## Prior reading
- [JSON Schema core spec](https://json-schema.org/draft/2020-12/json-schema-core.html)
- [Go encoding/json package](https://pkg.go.dev/encoding/json)

## What this kata is about (and why it matters)
- Core lesson: validation is fail-closed: unknown types and malformed input must error, never pass silently.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
