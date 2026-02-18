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

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
