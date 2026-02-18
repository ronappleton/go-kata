# Kata 078 — SQLite-backed Repo (pure Go driver)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata78() error
```

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go database/sql package](https://pkg.go.dev/database/sql)
- [modernc.org/sqlite driver docs](https://pkg.go.dev/modernc.org/sqlite)
- [SQLite transaction docs](https://www.sqlite.org/lang_transaction.html)

## What this kata is about (and why it matters)
- This kata is about implementing SQLite-backed Repo (pure Go driver) with constraints that make you practice system design, integration boundaries, and robust testing.
- It is important because these same tradeoffs appear in production: correctness at boundaries, predictable behavior under edge cases, and maintainable tests.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
