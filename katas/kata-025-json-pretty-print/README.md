# Kata 025 — Pretty JSON

**Focus:** encoding/json validation and formatting

## Your task
Implement:

```go
func PrettyJSON(input []byte) ([]byte, error)
```

### Learning goal
- What you are building: strict JSON validation plus standardized pretty output.
- Why this matters in real projects: logs, debug tools, and fixtures rely on stable JSON formatting.
- How this grows your Go skills: you practice safe JSON transformations with `encoding/json`.
- Definition of done (plain English): valid JSON is reformatted predictably and invalid JSON fails clearly.

### Tips
- Validate before formatting.
- Assert semantic equivalence by unmarshalling in tests.
- Test malformed inputs explicitly.

## Rules / Expectations
- invalid JSON input => non-nil error
- valid JSON => pretty-print with 2-space indentation
- output ends with a trailing newline
- semantic JSON content remains unchanged

## Prior reading
- [encoding/json package](https://pkg.go.dev/encoding/json)
- [bytes package](https://pkg.go.dev/bytes)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: formatting output should never silently change meaning.
- After this kata, you should be able to implement safe JSON normalization utilities.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
