# Kata 172 — Composition Over Inheritance

**Focus:** Struct embedding, small interfaces, accept interfaces return structs

## Your task

Build composable types using embedding and small interfaces.

### Learning goal
- What you are practicing: Go's composition model — struct embedding for reuse, small interfaces for flexibility.
- Why this matters: Go has no class inheritance. Composition is the Go way to build complex types from simple ones.
- How this grows your Go skills: you'll design flexible, testable systems using Go's type system.

## Rules / Expectations
- BaseItem provides common String() and HasTag()
- Task and TodoList embed BaseItem
- DescribeItem accepts any Stringer (interface flexibility)
- FilterItems works with any TagChecker

## What this kata is about (and why it matters)
- Core lesson: Go uses composition, not inheritance. Small interfaces + struct embedding = clean, testable design.
- After this kata, you should design composable Go types naturally.

## What you must submit for marking
- `kata.go`
