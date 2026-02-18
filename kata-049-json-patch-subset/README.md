# Kata 049 — JSON Patch (subset)

**Focus:** encoding/json

## Your task
Implement:

```go
func ApplyPatch(doc []byte, patch []byte) ([]byte, error)
```

### Learning goal
- What you are practicing: Build `JSON Patch (subset)` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `support add/replace/remove at top-level keys only`; and `invalid => error`.

## Rules / Expectations
- support add/replace/remove at top-level keys only
- invalid => error

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in encoding/json through `JSON Patch (subset)`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
