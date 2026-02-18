# Kata 038 — Mini Template Renderer

**Focus:** Placeholders

## Your task
Implement:

```go
func RenderTemplate(tpl string, vars map[string]string) string
```

## Rules / Expectations
- placeholders {{key}}
- missing => empty
- no recursion

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
