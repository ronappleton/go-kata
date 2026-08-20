# Kata 091 — testify Assertions

**Focus:** github.com/stretchr/testify, assert/require, test helpers

## Your task
Rewrite tests using testify-style assertions.

### Learning goal
- What you are practicing: using the testify library for cleaner, more expressive tests.
- Why this matters: testify is the most popular Go testing library. Its assert/require packages reduce boilerplate.
- How this grows your Go skills: you learn assert.Equal, require.NoError, and testify suites.

### Tips
- `assert.Equal(t, want, got)` replaces manual if-checks.
- `require.NoError(t, err)` stops the test immediately on failure.
- Use `assert.Contains`, `assert.Len`, `assert.True` for specific checks.

## Rules / Expectations
- Multiply(2,3) => 6
- Multiply(0,5) => 0
- Tests use assertion patterns

## What this kata is about (and why it matters)
- Core lesson: good assertions make tests readable and failures obvious.
- After this kata, your tests will be cleaner and more maintainable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
