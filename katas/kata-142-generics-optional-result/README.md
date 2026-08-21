# Kata 142 — Generics: Optional/Result

**Focus:** generics

## Your task
Implement:

```go
type Optional[T any] struct{}; func Some[T any](v T) Optional[T]; func None[T any]() Optional[T]; func (o Optional[T]) IsPresent() bool; func (o Optional[T]) Get() (T, bool); func (o Optional[T]) OrElse(def T) T; type Result[T any] struct{}; func Ok[T any](v T) Result[T]; func Err[T any](err error) Result[T]; func (r Result[T]) IsOk() bool; func (r Result[T]) Value() (T, error)
```

### Learning goal
- What you are building: type Optional[T any] struct{}; func Some[T any](v T) Optional[T]; func None[T any]() Optional[T]; func (o Optional[T]) IsPresent() bool; func (o Optional[T]) Get() (T, bool); func (o Optional[T]) OrElse(def T) T; type Result[T any] struct{}; func Ok[T any](v T) Result[T]; func Err[T any](err error) Result[T]; func (r Result[T]) IsOk() bool; func (r Result[T]) Value() (T, error) as a reliable contract. Focus: generics.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): Optional models a present-or-absent value; Result models success-or-error without panics; both are generic over T.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- Some/None semantics
- OrElse default
- Ok/Err semantics
- generic over int and string

## Prior reading
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)
- [Option pattern (Wikipedia)](https://en.wikipedia.org/wiki/Option_type)

## What this kata is about (and why it matters)
- Core lesson: Optional/Result make absence and failure explicit in the type system instead of nil/panic conventions.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
