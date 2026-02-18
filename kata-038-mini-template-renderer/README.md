# Kata 038 — Mini Template Renderer

**Focus:** Placeholders

## Your task
Implement:

```go
func RenderTemplate(tpl string, vars map[string]string) string
```

### Learning goal
- Expected work: Implement `Mini Template Renderer` as strict input transformation logic that accepts valid data and rejects malformed input.
- Why: `Mini Template Renderer` teaches strict boundary handling so malformed input cannot silently corrupt results.

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
- This kata trains placeholders by implementing `Mini Template Renderer` under explicit constraints.
- It is important because `Mini Template Renderer` builds strict parser habits that prevent downstream data and logic errors.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
