# Kata 100 — Plugin-free DI Container

**Focus:** Design

## Your task
Implement:

```go
func NewContainer() *Container
```

## Rules / Expectations
- Register(name, ctor)
- Resolve(name)
- singleton scope

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
