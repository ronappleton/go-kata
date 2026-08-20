# Kata 170 — Error Handling Idioms

**Focus:** if err != nil, sentinel errors, error wrapping, custom types, Is/As

## Your task

Implement idiomatic error handling patterns.

### Learning goal
- What you are practicing: the Go error handling contract — always check, wrap with context, use sentinels and types appropriately.
- Why this matters: error handling is THE defining Go pattern. Bad error handling is the #1 source of bugs.
- How this grows your Go skills: you'll write robust code that fails gracefully and is debuggable.

## Rules / Expectations
- ValidateEmail returns ErrInvalidEmail sentinel
- ValidateAge returns *ValidationError with context
- ClassifyError uses errors.Is and errors.As
- WrapWithContext wraps with fmt.Errorf %w

## What this kata is about (and why it matters)
- Core lesson: Go errors are values. Handle them explicitly, wrap them with context, and use types/sentinels for callers.
- After this kata, you should handle errors idiomatically in any Go program.

## What you must submit for marking
- `kata.go`
