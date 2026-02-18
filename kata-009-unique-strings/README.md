# Kata 009 — Unique Strings

**Focus:** Maps as sets, order preservation

## Your task
Implement:

```go
func Unique(items []string) []string
```

### Learning goal
- Expected work: Implement `Unique Strings` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Unique Strings` teaches core implementation habits that compound across all later katas.
- When correct: `Unique Strings` should satisfy the required behavior, including: `preserve first occurrence order`; `nil => empty slice`; and `case-sensitive`.

## Rules / Expectations
- preserve first occurrence order
- nil => empty slice
- case-sensitive

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata trains maps as sets, order preservation by implementing `Unique Strings` under explicit constraints.
- It is important because `Unique Strings` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
