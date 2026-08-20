# Kata 162 — AI as a Debugging Partner

**Focus:** using AI to understand errors, not just fix them, root cause analysis

## Your task
Implement a structured debugging session that uses AI to understand errors.

### Learning goal
- What you are practicing: using AI to understand WHY code fails, not just HOW to fix it.
- Why this matters: AI can generate a fix in seconds, but if you don't understand the root cause, you'll hit the same bug again. Understanding > Fixing.
- How this grows your Go skills: you learn systematic debugging and root cause analysis.

### The Anti-Copy-Paste Debugging Rule
When you hit an error:
1. **Read the error message yourself first.** Understand what it says.
2. **Check the stack trace.** Where did it happen? What called what?
3. **Form a hypothesis.** "I think X is nil because Y."
4. **THEN ask AI** — but ask "WHY is this happening?" not "HOW do I fix this?"
5. **Verify the AI's explanation** matches your hypothesis
6. **Write a test** that reproduces the bug before fixing it

### Never Do This
- Copy-pasting AI's "fix" without understanding why it works
- Asking AI to "just fix it" without providing error context
- Skipping the hypothesis step — that's where learning happens

## Rules / Expectations
- Debugging session includes hypotheses
- Includes test suggestions
- Conclusion is empty until analysis is complete

## What this kata is about (and why it matters)
- Core lesson: AI helps you understand, not just fix. Understanding is the skill.
- After this kata, you'll debug with AI as a thinking partner, not a crutch.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
