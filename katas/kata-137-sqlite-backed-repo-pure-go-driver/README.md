# Kata 137 — Repository Pattern (in-memory store)

**Focus:** design patterns, interfaces

## Your task
Implement:

```go
type User struct {...}; type UserRepo interface {...}; type InMemoryRepo struct{}; func NewInMemoryRepo() *InMemoryRepo
```

### Learning goal
- What you are building: type User struct {...}; type UserRepo interface {...}; type InMemoryRepo struct{}; func NewInMemoryRepo() *InMemoryRepo as a reliable contract. Focus: design patterns, interfaces.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): UserRepo offers Create/Get/Update/Delete/List over User; InMemoryRepo implements it; Get on a missing id returns ErrNotFound; Create assigns an id when empty; Update on missing id errors; Delete is idempotent.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- UserRepo interface with Create/Get/Update/Delete/List
- ErrNotFound sentinel
- InMemoryRepo satisfies the interface
- insertion-order List

## Prior reading
- [Repository pattern (Martin Fowler)](https://martinfowler.com/eaaCatalog/repository.html)
- [Go: Effective Go](https://go.dev/doc/effective_go)

## What this kata is about (and why it matters)
- Core lesson: the repository pattern hides the store behind an interface so behavior can be tested without a database.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
