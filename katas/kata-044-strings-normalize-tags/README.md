# Kata 021 — Normalize Tags

**Focus:** strings.Split, strings.TrimSpace, strings.ToLower

## Your task
Implement:

```go
func NormalizeTags(input string) []string
```

### Learning goal
- What you are building: a deterministic normalization function for comma-delimited tags.
- Why this matters in real projects: request filtering, search metadata, and config labels all need consistent normalized values.
- How this grows your Go skills: you practice reliable text parsing with the `strings` package.
- Definition of done (plain English): your function produces stable normalized tags for valid and messy input.

### Tips
- Start by proving split/trim behavior with tests.
- Keep behavior deterministic for empty and whitespace-only inputs.
- Prefer readable loops over compact one-liners.

## Rules / Expectations
- split input by comma (`,`)
- trim spaces around each tag
- drop empty tags after trim
- convert each kept tag to lowercase
- preserve original left-to-right order

## Prior reading
- [strings package](https://pkg.go.dev/strings)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go: Effective Go](https://go.dev/doc/effective_go)

## What this kata is about (and why it matters)
- Core lesson: standardize user input into one predictable internal form.
- After this kata, you should be able to explain each transform step and why order stability matters.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
