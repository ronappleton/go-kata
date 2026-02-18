# Kata 140 — Run In Transaction

**Focus:** transaction control flow

## Your task
Implement:

```go
func RunInTransaction(begin func() error, commit func() error, rollback func() error, run func() error) error
```

### Learning goal
- What you are building: deterministic transaction lifecycle handling.
- Why this matters in real projects: transaction mis-handling causes data inconsistency and difficult recovery scenarios.
- How this grows your Go skills: you practice explicit begin/run/commit/rollback decision flow.
- Definition of done (plain English): behavior matches all expectations below.

## Rules / Expectations
- begin error => return immediately
- run error => rollback then return run error
- commit error => return commit error
- successful path => nil

## Prior reading
- [database/sql package](https://pkg.go.dev/database/sql)
- [errors package](https://pkg.go.dev/errors)
- [Go package documentation index](https://pkg.go.dev/std)

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
