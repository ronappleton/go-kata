# Kata 035 — Log Line Parser

**Focus:** Key=value parsing

## Your task
Implement:

```go
func ParseLogLine(line string) (map[string]string, error)
```

### Learning goal
- What you are building: Build `func ParseLogLine(line string) (map[string]string, error)` as a reliable contract. Focus: Key=value parsing.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: key=value pairs separated by spaces; values may be quoted; and duplicate keys last wins.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

## Rules / Expectations
- key=value pairs separated by spaces
- values may be quoted
- duplicate keys last wins

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

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
