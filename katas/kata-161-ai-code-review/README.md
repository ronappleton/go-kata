# Kata 161 — AI-Assisted Code Review

**Focus:** using AI for code review, verifying AI suggestions, critical thinking

## Your task
Implement a code review function that demonstrates critical thinking.

### Learning goal
- What you are practicing: using AI as a review partner while maintaining your own judgment.
- Why this matters: AI can catch issues you miss, but it also hallucinates. You need to know when to trust it and when to question it.
- How this grows your Go skills: you develop the habit of reviewing code critically — your own and AI's.

### The Verification Loop
1. **Ask AI to review your code** — it will find real issues and fake ones
2. **Verify each suggestion** — does it actually apply? Is it correct?
3. **Run the tests** — AI suggestions that break tests are wrong
4. **Apply your judgment** — some AI suggestions are technically correct but practically wrong

### Red Flags in AI Code Review
- Suggesting changes that would break existing tests
- Over-engineering simple solutions
- Ignoring the project's existing patterns
- Suggesting deprecated or insecure approaches
- Being confidently wrong about language features

## Rules / Expectations
- Returns structured review with issues and suggestions
- Confidence score reflects actual certainty
- Suggestions are actionable and specific

## What this kata is about (and why it matters)
- Core lesson: AI is a review partner, not a replacement for your judgment.
- After this kata, you'll use AI reviews critically and effectively.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
