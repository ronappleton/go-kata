# Kata 166 — init() Function & Package Lifecycle

**Focus:** init() execution order, package-level initialization, side effects

## Your task

Understand and implement initialization order tracking.

### Learning goal
- What you are practicing: understanding Go's initialization sequence — variables first, then init(), then main().
- Why this matters: init() bugs are hard to debug because the execution order isn't obvious.
- How this grows your Go skills: you'll understand why some Go code "just works" and some doesn't.

## Rules / Expectations
- Initialize returns the correct order: package_vars, init_func, main
- GetInitOrder and ResetInitOrder work correctly

## What this kata is about (and why it matters)
- Core lesson: Go has a strict initialization order. Understanding it prevents subtle bugs.
- After this kata, you should predict initialization behavior in any Go program.

## What you must submit for marking
- `kata.go`
