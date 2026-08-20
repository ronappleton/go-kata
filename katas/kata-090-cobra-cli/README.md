# Kata 090 — Cobra CLI

**Focus:** github.com/spf13/cobra, CLI commands, flags, help text

## Your task
Implement CLI argument parsing.

### Learning goal
- What you are practicing: parsing command-line arguments into structured config.
- Why this matters: most Go tools are CLIs. Understanding flag parsing and command structure is essential.
- How this grows your Go skills: you learn the patterns used by every popular Go CLI tool.

### Tips
- Handle both `--flag=value` and `--flag value` formats.
- Always provide helpful error messages for invalid arguments.

## Rules / Expectations
- Parse command name from first argument
- Parse --name= value
- Parse --verbose boolean flag

## What this kata is about (and why it matters)
- Core lesson: good CLIs have clear argument parsing and helpful errors.
- After this kata, you can build production-quality CLI tools.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
