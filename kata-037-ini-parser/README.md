# Kata 037 — INI Parser

**Focus:** Sections, maps

## Your task
Implement:

```go
func ParseINI(s string) (map[string]map[string]string, error)
```

## Rules / Expectations
- sections [name]
- key=value
- comments start with ; or #

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
