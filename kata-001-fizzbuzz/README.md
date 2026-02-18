# Kata 001 — FizzBuzz

**Focus:** Basics: loops, conditionals, slices, strconv

## Your task
Implement:

```go
func FizzBuzz(n int) []string
```

### Learning goal
- Expected work: Implement `FizzBuzz` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `FizzBuzz` teaches core implementation habits that compound across all later katas.

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
- This kata trains basics: loops, conditionals, slices, strconv by implementing `FizzBuzz` under explicit constraints.
- It is important because `FizzBuzz` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
