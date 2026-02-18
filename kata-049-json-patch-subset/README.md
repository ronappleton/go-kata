# Kata 049 — JSON Patch (subset)

**Focus:** encoding/json

## Your task
Implement:

```go
func ApplyPatch(doc []byte, patch []byte) ([]byte, error)
```

## Rules / Expectations
- support add/replace/remove at top-level keys only
- invalid => error

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing JSON Patch (subset) with constraints that make you practice encoding/json.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
