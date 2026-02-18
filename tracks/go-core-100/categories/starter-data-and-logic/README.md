# Starter: Data and Logic

**Kata range:** `011-020`

## What you are building
You are learning how to transform slices, maps, and structs into predictable outputs.

## Why this matters in real Go work
Business logic is mostly filtering, searching, mapping, and combining data. Strong habits here make API and service code easier to maintain.

## What "correct" looks like by the end
- You preserve required ordering and return shape.
- You avoid mutating caller input unless the contract says to.
- You can explain why each branch exists and what rule it satisfies.

## How to practice this category
- Write table-driven tests for normal and edge cases.
- Check behavior for empty collections and missing values.
- Prefer readable loops and clear naming over compressed logic.
