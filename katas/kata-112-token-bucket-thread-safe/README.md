# Kata 112 — Token Bucket (thread-safe)

**Focus:** Concurrency, time

## Your task
Implement:

```go
func NewTokenBucket(ratePerSec int, burst int) (*TokenBucket, error)
```

### Learning goal
- What you are building: func NewTokenBucket(ratePerSec int, burst int) (*TokenBucket, error) as a reliable contract. Focus: Concurrency, time.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): ratePerSec>=1, burst>=1; Allow() bool under concurrent use; refills at ratePerSec; never exceeds burst.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- thread-safe under concurrent Allow
- refill at ratePerSec tokens/sec up to burst
- Allow returns false when empty
- error when ratePerSec < 1 or burst < 1

## Prior reading
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)
- [Go time package](https://pkg.go.dev/time)

## What this kata is about (and why it matters)
- Core lesson: make Allow() a single locked decision so burst semantics hold under contention.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
