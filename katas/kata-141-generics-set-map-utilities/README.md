# Kata 141 — Generics: Set Utilities

**Focus:** generics

## Your task
Implement:

```go
type Set[T comparable] struct{}; func NewSet[T comparable]() *Set[T]; func (s *Set[T]) Add(v T); func (s *Set[T]) Remove(v T); func (s *Set[T]) Contains(v T) bool; func (s *Set[T]) Len() int; func Union[T comparable](a, b *Set[T]) *Set[T]; func Intersection[T comparable](a, b *Set[T]) *Set[T]
```

### Learning goal
- What you are building: type Set[T comparable] struct{}; func NewSet[T comparable]() *Set[T]; func (s *Set[T]) Add(v T); func (s *Set[T]) Remove(v T); func (s *Set[T]) Contains(v T) bool; func (s *Set[T]) Len() int; func Union[T comparable](a, b *Set[T]) *Set[T]; func Intersection[T comparable](a, b *Set[T]) *Set[T] as a reliable contract. Focus: generics.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): Set holds unique values of any comparable type; Add/Remove/Contains/Len work; Union and Intersection combine sets.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- works for int and string
- no duplicates
- Union and Intersection
- Len reflects unique count

## Prior reading
- [Go generics tutorial](https://go.dev/doc/tutorial/generics)
- [Go type parameters proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)

## What this kata is about (and why it matters)
- Core lesson: generics let one implementation serve many element types; the tests prove it for two.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
