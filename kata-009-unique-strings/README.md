# Kata 009 — Unique Strings

**Focus:** Maps as sets, order preservation

## Your task
Implement:

```go
func Unique(items []string) []string
```

### Learning goal
- What you are practicing: Build `Unique Strings` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `preserve first occurrence order`; `nil => empty slice`; and `case-sensitive`.

## Rules / Expectations
- preserve first occurrence order
- nil => empty slice
- case-sensitive

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata is focused practice in Maps as sets, order preservation through `Unique Strings`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
