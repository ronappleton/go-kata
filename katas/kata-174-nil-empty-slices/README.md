# Kata 174 — Nil vs Empty Slices & Maps

**Focus:** nil vs empty, initialization patterns, when to use which

## Your task

Handle nil and empty collections correctly.

### Learning goal
- What you are practicing: understanding nil vs empty slices/maps, safe operations, and API-friendly patterns.
- Why this matters: nil vs empty bugs cause JSON null issues, panics on nil maps, and confusing behavior.
- How this grows your Go skills: you'll write predictable code that handles zero values gracefully.

## Rules / Expectations
- ClassifySlice distinguishes nil, empty, populated
- SafeMapGet handles nil maps
- EnsureEmpty returns non-nil empty slices
- NormalizeMap produces non-nil maps

## What this kata is about (and why it matters)
- Core lesson: nil slices are valid but behave differently from empty slices. Maps must be initialized before writes.
- After this kata, you should handle nil/empty collections without fear.

## What you must submit for marking
- `kata.go`
