# Kata 022 — JSON Pretty Print

**Focus:** encoding/json

## Your task
Implement:

```go
func PrettyJSON(input []byte) ([]byte, error)
```

### Learning goal
- What you are practicing: Build `JSON Pretty Print` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `invalid json => error`; `indent 2 spaces`; and `stable formatting`.

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
- This kata is focused practice in encoding/json through `JSON Pretty Print`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
