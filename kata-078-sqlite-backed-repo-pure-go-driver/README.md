# Kata 078 — SQLite-backed Repo (pure Go driver)

**Focus:** Advanced practice

## Your task
Implement:

```go
func Kata78() error
```

### Learning goal
- What you are practicing: Build `SQLite-backed Repo (pure Go driver)` with solid persistence boundaries and transaction-aware behavior.
- Why it matters: You will use this in repository/service layers where data correctness and failure handling matter.
- How this grows your Go skills: This builds stronger transactional thinking and cleaner separation of persistence concerns.
- When correct: When your solution is correct, it should satisfy: `follow README spec`; `write tests`; and `keep it idiomatic`.

## Rules / Expectations
- follow README spec
- write tests
- keep it idiomatic

## Prior reading
- [Go database/sql package](https://pkg.go.dev/database/sql)
- [modernc.org/sqlite driver docs](https://pkg.go.dev/modernc.org/sqlite)
- [SQLite transaction docs](https://www.sqlite.org/lang_transaction.html)

## What this kata is about (and why it matters)
- This kata is focused practice in real-world Go design and testing through `SQLite-backed Repo (pure Go driver)`.
- You will use this in repository/service layers where data correctness and failure handling matter.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
