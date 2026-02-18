# Kata 068 — Mini Template Renderer

**Focus:** Placeholders

## Your task
Implement:

```go
func RenderTemplate(tpl string, vars map[string]string) string
```

### Learning goal
- What you are building: Build `func RenderTemplate(tpl string, vars map[string]string) string` as a reliable contract. Focus: Placeholders.
- Why this matters in real projects: Parsers are trust boundaries. Loose parsing creates downstream bugs.
- How this grows your Go skills: You practice strict input contracts and explicit error paths.
- Definition of done (plain English): A reviewer should be able to confirm this behavior in tests: placeholders {{key}}; when `missing`, return `empty`; and no recursion.

### Tips
- Reject ambiguous input early.
- Make malformed-input tests first-class.
- Do not silently coerce data.

## Rules / Expectations
- placeholders {{key}}
- missing => empty
- no recursion

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- Core lesson: parse strictly and fail loudly on bad input.
- After this kata, you should be able to state input contracts and return precise parse errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
