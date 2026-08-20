# Kata 164 — Go Application Structure

**Focus:** cmd/, internal/, pkg/, directory layout, package naming conventions

## Your task

Implement functions that understand Go project structure.

### Learning goal
- What you are practicing: understanding how Go projects are organized, import paths, and package privacy.
- Why this matters: correct structure makes your code maintainable and importable.
- How this grows your Go skills: you'll architect real Go applications properly.

## Rules / Expectations
- PackagePath returns the correct cmd/ import path
- PackageName extracts the last component
- IsInternal checks for /internal/ in the path

## What this kata is about (and why it matters)
- Core lesson: Go has strong conventions for project layout. cmd/, internal/, and pkg/ are standard.
- After this kata, you should structure any Go project correctly.

## What you must submit for marking
- `kata.go`
