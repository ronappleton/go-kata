# Kata 168 — Makefile Basics

**Focus:** Make targets, variables, dependencies, phony targets, build/test/clean patterns

## Your task

Parse and generate Makefiles for Go projects.

### Learning goal
- What you are practicing: understanding Make syntax, targets, variables, and common Go build patterns.
- Why this matters: every Go project needs a consistent build interface. Makefiles standardize build/test/lint/clean.
- How this grows your Go skills: you'll create reproducible build systems that work on any machine.

## Rules / Expectations
- ParseMakefile extracts target names in order
- ExtractVariables returns variable assignments
- ValidateTargets finds missing required targets
- GenerateMakefile creates a complete Go project Makefile

## What this kata is about (and why it matters)
- Core lesson: Makefiles are the universal build system. Understanding targets, dependencies, and recipes is essential.
- After this kata, you should write Makefiles for any Go project.

## What you must submit for marking
- `kata.go`
