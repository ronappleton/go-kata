# Kata 005 — Go Vet & Lint

**Focus:** go vet, golangci-lint, static analysis, code quality

## Your task
Implement:

```go
func CheckCode(code string) []string
```

### Learning goal
- What you are practicing: using static analysis tools to catch bugs before they reach production.
- Why this matters: `go vet` and linters catch common mistakes that the compiler misses. Every Go team uses them.
- How this grows your Go skills: you develop the habit of checking code quality automatically.

### Tips
- Run `go vet ./...` on any Go project to see it in action.
- Try installing golangci-lint for more thorough checks.

## Rules / Expectations
- clean code => empty slice

## What this kata is about (and why it matters)
- Core lesson: automated quality checks are free bugs prevention.
- After this kata, you should always run `go vet` before committing.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
