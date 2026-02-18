# Kata 035 — Log Line Parser

**Focus:** Key=value parsing

## Your task
Implement:

```go
func ParseLogLine(line string) (map[string]string, error)
```

### Learning goal
- Expected work: Implement `Log Line Parser` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `Log Line Parser` teaches strict boundary handling so malformed input cannot silently corrupt results.

## Rules / Expectations
- key=value pairs separated by spaces
- values may be quoted
- duplicate keys last wins

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains key=value parsing by implementing `Log Line Parser` under explicit constraints.
- It is important because `Log Line Parser` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
