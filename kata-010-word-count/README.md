# Kata 010 — Word Count

**Focus:** Strings, fields, maps

## Your task
Implement:

```go
func WordCount(s string) map[string]int
```

### Learning goal
- What you are practicing: Build `Word Count` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `split on whitespace`; `lowercase words`; and `empty => empty map (not nil)`.

## Rules / Expectations
- split on whitespace
- lowercase words
- empty => empty map (not nil)

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata is focused practice in Strings, fields, maps through `Word Count`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
