# Kata 007 — Fibonacci

**Focus:** Iteration, slices

## Your task
Implement:

```go
func Fibonacci(n int) ([]int, error)
```

### Learning goal
- What you are practicing: Build `Fibonacci` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `n<0 => error`; `n=0 => empty slice`; `n=1 => [0]`; and `sequence starts 0,1,...`.

## Rules / Expectations
- n<0 => error
- n=0 => empty slice
- n=1 => [0]
- sequence starts 0,1,...

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Iteration, slices through `Fibonacci`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
