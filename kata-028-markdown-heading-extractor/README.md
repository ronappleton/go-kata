# Kata 028 — Markdown Heading Extractor

**Focus:** Line scanning

## Your task
Implement:

```go
func ExtractHeadings(md string) []string
```

### Learning goal
- Expected work: Implement `Markdown Heading Extractor` idiomatically with strong control-flow clarity and edge-case correctness.
- Why: `Markdown Heading Extractor` teaches core implementation habits that compound across all later katas.
- When correct: `Markdown Heading Extractor` should satisfy the required behavior, including: `lines starting with #+space`; `return heading text`; and `preserve order`.

## Rules / Expectations
- lines starting with #+space
- return heading text
- preserve order

## Prior reading
- [Go strings package](https://pkg.go.dev/strings)
- [Go unicode package](https://pkg.go.dev/unicode)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata trains line scanning by implementing `Markdown Heading Extractor` under explicit constraints.
- It is important because `Markdown Heading Extractor` strengthens the baseline coding discipline every other kata depends on.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
