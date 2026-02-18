# Kata 038 — Mini Template Renderer

**Focus:** Placeholders

## Your task
Implement:

```go
func RenderTemplate(tpl string, vars map[string]string) string
```

### Learning goal
- What you are practicing: Build `Mini Template Renderer` to transform/parse input strictly, with correct handling of malformed data.
- Why it matters: You will use this at system boundaries where external input must be handled safely and predictably.
- How this grows your Go skills: This builds robust boundary validation and deterministic data transformation habits.
- When correct: When your solution is correct, it should satisfy: `placeholders {{key}}`; `missing => empty`; and `no recursion`.

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
- This kata is focused practice in Placeholders through `Mini Template Renderer`.
- You will use this at system boundaries where external input must be handled safely and predictably.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
