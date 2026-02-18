# Kata 005 — Count Vowels

**Focus:** Loops, unicode, maps

## Your task
Implement:

```go
func CountVowels(s string) int
```

### Learning goal
- What you are practicing: Build `Count Vowels` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `count a,e,i,o,u (case-insensitive)`; `treat vowels as ASCII set only`; and `empty => 0`.

## Rules / Expectations
- count a,e,i,o,u (case-insensitive)
- treat vowels as ASCII set only
- empty => 0

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata is focused practice in Loops, unicode, maps through `Count Vowels`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
