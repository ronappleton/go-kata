# Kata 005 — Count Vowels

**Focus:** Loops, unicode, maps

## Your task
Implement:

```go
func CountVowels(s string) int
```

## Rules / Expectations
- count a,e,i,o,u (case-insensitive)
- treat vowels as ASCII set only
- empty => 0

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata is about implementing Count Vowels with constraints that make you practice loops, unicode, maps.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
