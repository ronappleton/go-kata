# Kata 025 — LRU Cache (single-thread)

**Focus:** Structs, maps, list

## Your task
Implement:

```go
func NewLRU(capacity int) (*LRU, error)
```

### Learning goal
- What you are building: Build `func NewLRU(capacity int) (*LRU, error)` as a reliable contract. Focus: Structs, maps, list.
- Why this matters in real projects: This is how you keep features fast when input size grows.
- How this grows your Go skills: You practice invariants and complexity reasoning, then prove both with tests.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: capacity>0; get/put methods; and evict least recently used.

### Tips
- State the invariant first, then code.
- Test shape edges early: empty, one item, duplicates.
- Check complexity after correctness.

## Rules / Expectations
- capacity>0
- Get/Put methods
- evict least recently used

## Prior reading
- [Go map types in language spec](https://go.dev/ref/spec#Map_types)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- Core lesson: hold invariants first, then optimize safely.
- After this kata, you should be able to defend algorithm choice, complexity, and corner-case behavior.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
