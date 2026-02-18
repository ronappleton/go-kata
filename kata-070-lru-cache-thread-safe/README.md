# Kata 070 — LRU Cache (thread-safe)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata70() error
```

### Learning goal
- What you are practicing: Build `LRU Cache (thread-safe)` with safe coordination so concurrent work finishes cleanly under load.
- Why it matters: You will use this any time work runs in parallel and must shut down cleanly without races or leaks.
- How this grows your Go skills: This builds mental models for goroutines, channels, cancellation, and synchronization.
- When correct: When your solution is correct, it should satisfy: `follow README spec`; `write tests`; and `keep it idiomatic`.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go memory model](https://go.dev/ref/mem)
- [Go sync package](https://pkg.go.dev/sync)

## What this kata is about (and why it matters)
- This kata is focused practice in real-world Go design and testing through `LRU Cache (thread-safe)`.
- You will use this any time work runs in parallel and must shut down cleanly without races or leaks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
