# Kata 050 — Config Loader

**Focus:** env + file

## Your task
Implement:

```go
func LoadConfig(path string, envPrefix string) (map[string]string, error)
```

## Rules / Expectations
- load key=value lines
- override with env
- trim spaces

## Prior reading
- [Go os package](https://pkg.go.dev/os)
- [Go path/filepath package](https://pkg.go.dev/path/filepath)
- [Go: Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation index](https://pkg.go.dev/std)
- [Go language specification](https://go.dev/ref/spec)

## What this kata is about (and why it matters)
- This kata is about implementing Config Loader with constraints that make you practice env + file.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
