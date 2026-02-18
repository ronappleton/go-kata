# Kata 010 — Word Count

**Focus:** Strings, fields, maps

## Your task
Implement:

```go
func WordCount(s string) map[string]int
```

### Learning goal
- Expected work: Implement `Word Count` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Word Count` teaches core implementation habits that compound across all later katas.
- When correct: `Word Count` should satisfy the required behavior, including: `split on whitespace`; `lowercase words`; and `empty => empty map (not nil)`.

## Rules / Expectations
- split on whitespace
- lowercase words
- empty => empty map (not nil)

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata trains strings, fields, maps by implementing `Word Count` under explicit constraints.
- It is important because `Word Count` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
