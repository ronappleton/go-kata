# Kata 163 — Variable Scope & Lifetime

**Focus:** local vs package vs global scope, block scoping, short vs long declaration

## Your task

Implement `CountDeclarations()` and understand variable scoping.

### Learning goal
- What you are practicing: understanding Go's variable declaration syntax, scope rules, and lifetime.
- Why this matters: scope bugs are silent killers — a shadowed variable can cause logic errors that compile fine.
- How this grows your Go skills: you'll write safer, more predictable code.

## Rules / Expectations
- CountDeclarations returns the number of distinct variables you declare
- ShadowExample should demonstrate understanding of block scoping

## What this kata is about (and why it matters)
- Core lesson: Go has strict block scoping. := creates block-scoped variables. var works at package level.
- After this kata, you should understand when variables live and die.

## What you must submit for marking
- `kata.go`
