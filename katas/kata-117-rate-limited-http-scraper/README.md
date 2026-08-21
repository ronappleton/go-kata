# Kata 117 — Rate-limited HTTP Scraper

**Focus:** net/http, concurrency, time

## Your task
Implement:

```go
func Scrape(ctx context.Context, urls []string, ratePerSec int, client *http.Client) (map[string]string, error)
```

### Learning goal
- What you are building: func Scrape(ctx context.Context, urls []string, ratePerSec int, client *http.Client) (map[string]string, error) as a reliable contract. Focus: net/http, concurrency, time.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): fetches every URL at most ratePerSec requests/sec; returns url->body for 2xx responses; first non-2xx or network error is returned with the partial map.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- rate limit across all requests
- ctx cancellation stops the loop
- 2xx => body in map
- first error returned with partial map

## Prior reading
- [Go net/http package](https://pkg.go.dev/net/http)
- [Go context package](https://pkg.go.dev/context)

## What this kata is about (and why it matters)
- Core lesson: throttling is a global property: the whole scrape shares one budget, not per-URL.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
