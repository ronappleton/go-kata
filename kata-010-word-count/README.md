# Kata 010 — Word Count

**Focus:** Strings, fields, maps

## Your task
Implement:

```go
func WordCount(s string) map[string]int
```

## Rules / Expectations
- split on whitespace
- lowercase words
- empty => empty map (not nil)

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)

## What this kata is about (and why it matters)
- This kata is about implementing Word Count with constraints that make you practice strings, fields, maps.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
