# Kata 054 — JSON Lines Filter

**Focus:** Stream processing

## Your task
Implement:

```go
func FilterJSONL(r io.Reader, w io.Writer, predicate func(map[string]any) bool) error
```

### Learning goal
- Expected work: Implement `JSON Lines Filter` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `JSON Lines Filter` teaches strict boundary handling so malformed input cannot silently corrupt results.

## Rules / Expectations
- each line JSON object
- write only matches
- skip blank lines

## Prior reading
- [Go encoding/json package](https://pkg.go.dev/encoding/json)
- [Go io package](https://pkg.go.dev/io)
- [Go bufio package](https://pkg.go.dev/bufio)

## What this kata is about (and why it matters)
- This kata trains stream processing by implementing `JSON Lines Filter` under explicit constraints.
- It is important because `JSON Lines Filter` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
