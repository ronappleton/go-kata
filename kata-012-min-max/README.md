# Kata 012 — Min/Max

**Focus:** Errors, edge cases

## Your task
Implement:

```go
func MinMax(nums []int) (min int, max int, err error)
```

### Learning goal
- What you are practicing: Build `Min/Max` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `empty or nil => error`; `single element => min=max`; and `handles negatives`.

## Rules / Expectations
- empty or nil => error
- single element => min=max
- handles negatives

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Errors, edge cases through `Min/Max`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
