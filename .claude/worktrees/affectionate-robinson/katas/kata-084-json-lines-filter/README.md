# Kata 084 — JSON Lines Filter

**Focus:** Stream processing

## Your task
Implement:

```go
func FilterJSONL(r io.Reader, w io.Writer, predicate func(map[string]any) bool) error
```

### Learning goal
- What you are building: Build `func FilterJSONL(r io.Reader, w io.Writer, predicate func(map[string]any) bool) error` as a reliable contract. Focus: Stream processing.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: each line json object; write only matches; and skip blank lines.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

## Rules / Expectations
- each line JSON object
- write only matches
- skip blank lines

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)

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
