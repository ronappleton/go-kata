# Kata 024 — Build Search URL

**Focus:** net/url parsing, query encoding

## Your task
Implement:

```go
func BuildSearchURL(base string, params map[string]string) (string, error)
```

### Learning goal
- What you are building: a safe query-URL builder that preserves existing parameters.
- Why this matters in real projects: malformed URL generation causes subtle integration bugs.
- How this grows your Go skills: you use `net/url` instead of manual string concatenation.
- Definition of done (plain English): output URL is valid, encoded, and deterministic.

### Tips
- Parse first, mutate query values second, encode last.
- Include tests for base URLs that already contain query params.
- Avoid manual escaping.

## Rules / Expectations
- invalid base URL => non-nil error
- preserve existing query parameters from base URL
- add only params with non-empty values
- return encoded URL string
- deterministic query encoding order by key

## Prior reading
- [net/url package](https://pkg.go.dev/net/url)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: URL handling should be structured, not stringly typed.
- After this kata, you should be able to build query URLs safely in API/client code.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
