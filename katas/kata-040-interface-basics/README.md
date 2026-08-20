# Kata 040 — Interface Basics

**Focus:** interfaces, implicit satisfaction, accept interfaces return structs

## Your task
Implement the Greeter interface for both FormalGreeter and CasualGreeter.

### Learning goal
- What you are practicing: defining and implementing interfaces in Go.
- Why this matters: interfaces are how Go achieves polymorphism and testability. "Accept interfaces, return structs" is a core Go idiom.
- How this grows your Go skills: you learn that Go interfaces are satisfied implicitly — no `implements` keyword needed.

### Tips
- In Go, a type satisfies an interface just by implementing its methods.
- Keep interfaces small — 1-3 methods is ideal.

## Rules / Expectations
- FormalGreeter => formal greeting
- CasualGreeter => casual greeting
- Both implement Greeter

## What this kata is about (and why it matters)
- Core lesson: interfaces enable polymorphism without inheritance.
- After this kata, you should design your code around small interfaces.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
