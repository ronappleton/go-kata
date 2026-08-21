# Kata 111 — LRU Cache (thread-safe)

**Focus:** Concurrency, structs

## Your task
Implement:

```go
func NewLRUCache(capacity int) (*LRUCache, error)
```

### Learning goal
- What you are building: func NewLRUCache(capacity int) (*LRUCache, error) as a reliable contract. Focus: Concurrency, structs.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): capacity>0; thread-safe Get/Put; evict least recently used; Len() reflects size.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- thread-safe under concurrent access
- evict LRU on overflow
- Get returns ok=false on miss
- error when capacity < 1

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)

## What this kata is about (and why it matters)
- Core lesson: own the lock discipline: every public method takes the mutex and releases it exactly once.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
