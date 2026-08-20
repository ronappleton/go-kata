# Databases and I/O

**Kata range:** `136-140`

## What you are building
You are learning persistence and input/output boundary handling with explicit contracts and resource safety.

## Why this matters in real Go work
Database and I/O code drives most production failures: leaks, partial writes, inconsistent reads, and bad error handling.

## What "correct" looks like by the end
- You handle query/stream errors clearly and early.
- You manage resources (`rows.Close`, file handles, buffers) intentionally.
- You can explain transaction and I/O behavior under failure conditions.

## How to practice this category
- Test success, partial failure, and malformed input cases.
- Keep boundary logic thin and explicit.
- Validate output contracts at the edges, not only in core logic.
