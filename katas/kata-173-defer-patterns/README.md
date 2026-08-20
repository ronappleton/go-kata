# Kata 173 — Defer Patterns & Cleanup

**Focus:** defer, resource cleanup, panic recovery, named returns

## Your task

Use defer for cleanup and error recovery.

### Learning goal
- What you are practicing: defer for guaranteed cleanup, recover() for panic safety, and understanding defer timing.
- Why this matters: resource leaks are silent killers. defer ensures cleanup happens. recover prevents cascading panics.
- How this grows your Go skills: you'll write robust code that cleans up after itself.

## Rules / Expectations
- CaptureOutput pairs open/close with defer
- SafeDivide uses recover() to catch divide-by-zero panics
- Pipeline ensures completed/failed is always recorded
- DeferredMultiply demonstrates argument evaluation timing

## What this kata is about (and why it matters)
- Core lesson: defer is Go's RAII equivalent. Use it for file/lock/connection cleanup. Use recover for defensive boundaries.
- After this kata, you should use defer confidently and idiomatically.

## What you must submit for marking
- `kata.go`
