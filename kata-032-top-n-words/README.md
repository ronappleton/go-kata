# Kata 032 — Top N Words

**Focus:** Sorting, ties

## Your task
Implement:

```go
func TopNWords(s string, n int) []string
```

### Learning goal
- Expected work: Implement `Top N Words` by preserving core invariants and handling edge cases, not just happy-path output.
- Why: `Top N Words` teaches complexity-aware correctness that impacts throughput and latency at scale.

## Rules / Expectations
- lowercase, split whitespace
- ties sorted alphabetically
- n<=0 => empty slice

## Prior reading
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains sorting, ties by implementing `Top N Words` under explicit constraints.
- It is important because `Top N Words` makes performance/correctness tradeoffs explicit and measurable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
