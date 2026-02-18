# Kata 035 — Log Line Parser

**Focus:** Key=value parsing

## Your task
Implement:

```go
func ParseLogLine(line string) (map[string]string, error)
```

### Learning goal
- What you are practicing: Build `Log Line Parser` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `key=value pairs separated by spaces`; `values may be quoted`; and `duplicate keys last wins`.

## Rules / Expectations
- key=value pairs separated by spaces
- values may be quoted
- duplicate keys last wins

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Key=value parsing through `Log Line Parser`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
