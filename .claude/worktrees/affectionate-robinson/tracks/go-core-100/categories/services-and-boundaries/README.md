# Services and Boundaries

**Kata range:** `071-090`

## What you are building
You are practicing HTTP, network, and service-boundary behavior with clear contracts and defensive error handling.

## Why this matters in real Go work
External systems fail, timeout, and return unexpected payloads. Senior Go code treats these as normal operating conditions.

## What "correct" looks like by the end
- You set clear request/response behavior for success and failure.
- You respect context cancellation and timeout boundaries.
- You surface failures with enough detail for diagnosis.

## How to practice this category
- Treat boundary code as contract code, not glue code.
- Verify retries/backoff/cancellation with focused tests.
- Keep side effects isolated behind small interfaces.
