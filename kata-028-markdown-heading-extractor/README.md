# Kata 028 — Markdown Heading Extractor

**Focus:** Line scanning

## Your task
Implement:

```go
func ExtractHeadings(md string) []string
```

### Learning goal
- What you are practicing: Build `Markdown Heading Extractor` with clean control flow and reliable edge-case behavior.
- Why it matters: You will use this in everyday Go code where small correctness habits prevent larger defects later.
- How this grows your Go skills: This builds the base layer of Go fluency that every advanced kata depends on.
- When correct: When your solution is correct, it should satisfy: `lines starting with #+space`; `return heading text`; and `preserve order`.

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
- This kata is focused practice in Line scanning through `Markdown Heading Extractor`.
- You will use this in everyday Go code where small correctness habits prevent larger defects later.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
