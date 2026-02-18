# Kata 008 — Filter Even

**Focus:** Slices, order, immutability

## Your task
Implement:

```go
func FilterEven(nums []int) []int
```

### Learning goal
- What you are practicing: Build `Filter Even` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `preserve order`; `nil => empty slice`; and `do not mutate input`.

## Rules / Expectations
- preserve order
- nil => empty slice
- do not mutate input

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is focused practice in Slices, order, immutability through `Filter Even`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
