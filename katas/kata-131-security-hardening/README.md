# Kata 131 — Security Hardening

**Focus:** input validation, SQL injection, XSS, secrets management

## Your task
Implement input sanitization and email validation.

### Learning goal
- What you are practicing: defensive coding against common security vulnerabilities.
- Why this matters: security bugs are the most expensive bugs. Input validation is your first line of defense.
- How this grows your Go skills: you learn to think like an attacker and validate everything.

### Tips
- Strip or encode HTML/script tags.
- Validate email format with regex or string checks.
- Never trust user input.

## Rules / Expectations
- Script tags => removed
- SQL injection patterns => removed
- Valid email => true
- Invalid email => false

## What this kata is about (and why it matters)
- Core lesson: validate and sanitize all user input. Always.
- After this kata, you'll think about security in every function that handles input.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
