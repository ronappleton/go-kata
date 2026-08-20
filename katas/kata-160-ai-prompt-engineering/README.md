# Kata 160 — AI Prompt Engineering for Developers

**Focus:** effective prompting, context setting, constraint specification

## Your task
Implement a prompt builder that constructs effective AI prompts.

### Learning goal
- What you are practicing: writing clear, structured prompts that get useful AI responses.
- Why this matters: AI is only as good as your prompt. Vague prompts get vague answers. Specific prompts get specific, testable code.
- How this grows your Go skills: you learn to articulate requirements precisely — a skill that helps even when you're NOT using AI.

### The Developer's AI Contract
1. **Never copy-paste AI output without understanding it.** Read every line. If you can't explain it, you can't maintain it.
2. **Use AI for boilerplate, not for thinking.** AI is great at "write a struct with these fields." It's bad at "design the architecture."
3. **Always verify with tests.** AI-generated code that passes tests is trustworthy. AI-generated code without tests is a liability.
4. **Treat AI like a junior developer.** Review its code. Ask questions. Don't blindly merge.

### Tips
- Start with "I want to..." not "Can you..."
- Include the language, error handling style, and test expectations
- Show existing code patterns you want to follow
- Ask for explanations, not just code

## Rules / Expectations
- prompt includes the goal
- prompt includes constraints when provided
- prompt is clear and actionable

## What this kata is about (and why it matters)
- Core lesson: AI is a tool, not a replacement for thinking.
- After this kata, you'll prompt AI effectively and verify its output.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
