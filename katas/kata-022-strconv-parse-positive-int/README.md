# Kata 022 — Parse Positive Int

**Focus:** strconv.Atoi, strings.TrimSpace, input validation

## Your task
Implement:

```go
func ParsePositiveInt(input string) (int, error)
```

### Learning goal
- What you are building: a safe parser that rejects invalid or non-positive numeric input.
- Why this matters in real projects: API and CLI inputs are untrusted; parse boundaries must fail clearly.
- How this grows your Go skills: you build strict parse-and-validate habits with `strconv`.
- Definition of done (plain English): callers can trust the return value is a valid positive integer when error is nil.

### Tips
- Treat parsing and semantic validation as separate checks.
- Use table-driven tests for invalid strings and boundary numbers.
- Keep error cases explicit.

## Rules / Expectations
- trim leading/trailing spaces before parsing
- invalid numeric text => non-nil error
- parsed value <= 0 => non-nil error
- valid positive value => return value and nil error

## Prior reading
- [strconv package](https://pkg.go.dev/strconv)
- [errors package](https://pkg.go.dev/errors)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: parsing is not complete until domain validation is also enforced.
- After this kata, you should be able to defend your parse boundary contract in code review.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
