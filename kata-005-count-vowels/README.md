# Kata 005 — Count Vowels

**Focus:** Loops, unicode, maps

## Your task
Implement:

```go
func CountVowels(s string) int
```

### Learning goal
- Expected work: Implement `Count Vowels` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Count Vowels` teaches core implementation habits that compound across all later katas.
- When correct: `Count Vowels` should satisfy the required behavior, including: `count a,e,i,o,u (case-insensitive)`; `treat vowels as ASCII set only`; and `empty => 0`.

## Rules / Expectations
- count a,e,i,o,u (case-insensitive)
- treat vowels as ASCII set only
- empty => 0

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata trains loops, unicode, maps by implementing `Count Vowels` under explicit constraints.
- It is important because `Count Vowels` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
