# Kata 078 — SQLite-backed Repo (pure Go driver)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata78() error
```

### Learning goal
- Expected work: Implement `SQLite-backed Repo (pure Go driver)` with clear repository boundaries and predictable transaction-aware persistence behavior.
- Why: `SQLite-backed Repo (pure Go driver)` teaches persistence invariants that protect data integrity during failures.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go database/sql package](https://pkg.go.dev/database/sql)
- [modernc.org/sqlite driver docs](https://pkg.go.dev/modernc.org/sqlite)
- [SQLite transaction docs](https://www.sqlite.org/lang_transaction.html)

## What this kata is about (and why it matters)
- This kata trains integration design, boundary handling, and robust testing by implementing `SQLite-backed Repo (pure Go driver)` under explicit constraints.
- It is important because `SQLite-backed Repo (pure Go driver)` reinforces data-consistency and transaction thinking used in real persistence layers.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
