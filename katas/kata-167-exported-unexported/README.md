# Kata 167 — Exported vs Unexported Identifiers

**Focus:** uppercase/lowercase naming, visibility rules, API design

## Your task

Implement functions that understand Go's visibility rules.

### Learning goal
- What you are practicing: understanding that Go uses naming case for visibility, not keywords.
- Why this matters: API design in Go is entirely based on exported/unexported naming.
- How this grows your Go skills: you'll design clean, well-encapsulated packages.

## Rules / Expectations
- IsExported returns true for uppercase-starting names
- FilterExported returns only exported names
- MaskUnexported replaces unexported names with "***"

## What this kata is about (and why it matters)
- Core lesson: Go's visibility is simple but powerful — uppercase = public, lowercase = private.
- After this kata, you should design APIs with clear public/private boundaries.

## What you must submit for marking
- `kata.go`
