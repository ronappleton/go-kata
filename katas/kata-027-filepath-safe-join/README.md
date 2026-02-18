# Kata 027 — Safe Join

**Focus:** filepath.Join, filepath.Clean, traversal guard checks

## Your task
Implement:

```go
func SafeJoin(base string, parts ...string) (string, error)
```

### Learning goal
- What you are building: a path join helper that blocks directory traversal escapes.
- Why this matters in real projects: file handling becomes a security issue when path boundaries are weak.
- How this grows your Go skills: you practice defensive path normalization with `path/filepath`.
- Definition of done (plain English): returned paths stay inside base, or an error is returned.

### Tips
- Normalize both base and candidate paths.
- Verify prefix boundaries carefully.
- Test traversal attempts such as `..` segments.

## Rules / Expectations
- empty base after clean => non-nil error
- clean and join path parts with base
- if resolved path escapes base => non-nil error
- if resolved path is inside base => return cleaned path

## Prior reading
- [path/filepath package](https://pkg.go.dev/path/filepath)
- [os package](https://pkg.go.dev/os)
- [Go package documentation index](https://pkg.go.dev/std)

## What this kata is about (and why it matters)
- Core lesson: path composition must enforce trust boundaries.
- After this kata, you should be able to review file-path code for traversal risks.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
