# Kata 001 — FizzBuzz

**Focus:** Basics: loops, conditionals, slices, strconv

## Your task
Implement:

```go
func FizzBuzz(n int) []string
```

### Learning goal
- What you are practicing: Build `FizzBuzz` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `n<=0 => empty slice (not nil)`; `multiples of 3 => Fizz`; `multiples of 5 => Buzz`; `both => FizzBuzz`; and `otherwise number string`.

## Rules / Expectations
- n<=0 => empty slice (not nil)
- multiples of 3 => Fizz
- multiples of 5 => Buzz
- both => FizzBuzz
- otherwise number string

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Basics: loops, conditionals, slices, strconv through `FizzBuzz`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
