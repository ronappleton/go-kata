# Kata 012 — Min/Max

**Focus:** Errors, edge cases

## Your task
Implement:

```go
func MinMax(nums []int) (min int, max int, err error)
```

### Learning goal
- Expected work: Implement `Min/Max` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Min/Max` teaches core implementation habits that compound across all later katas.
- When correct: `Min/Max` should satisfy the required behavior, including: `empty or nil => error`; `single element => min=max`; and `handles negatives`.

## Rules / Expectations
- empty or nil => error
- single element => min=max
- handles negatives

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains errors, edge cases by implementing `Min/Max` under explicit constraints.
- It is important because `Min/Max` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
