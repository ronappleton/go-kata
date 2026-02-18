# Kata 054 — JSON Lines Filter

**Focus:** Stream processing

## Your task
Implement:

```go
func FilterJSONL(r io.Reader, w io.Writer, predicate func(map[string]any) bool) error
```

### Learning goal
- What you are practicing: Build `JSON Lines Filter` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `each line JSON object`; `write only matches`; and `skip blank lines`.

## Rules / Expectations
- each line JSON object
- write only matches
- skip blank lines

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)

## What this kata is about (and why it matters)
- This kata is focused practice in Stream processing through `JSON Lines Filter`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
