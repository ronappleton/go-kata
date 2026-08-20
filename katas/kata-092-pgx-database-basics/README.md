# Kata 092 — pgx Database Basics

**Focus:** database/sql, pgx, connection, query, scan

## Your task
Implement a UserStore interface with a mock backing store.

### Learning goal
- What you are practicing: database access patterns with Go interfaces.
- Why this matters: every production Go app talks to a database. Understanding the interface pattern enables testing without a real DB.
- How this grows your Go skills: you learn the repository pattern and how to mock database access.

### Tips
- Define an interface for your store.
- Test with a mock implementation.
- Use context.Context for cancellation support.

## Rules / Expectations
- Create returns an ID
- GetByID returns the user or error
- Tests use mock store (no real DB)

## What this kata is about (and why it matters)
- Core lesson: mock your database boundary for fast, reliable tests.
- After this kata, you can design data access layers that are testable.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
