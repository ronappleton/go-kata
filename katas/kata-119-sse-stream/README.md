# Kata 119 — SSE Stream

**Focus:** net/http, streaming

## Your task
Implement:

```go
func SSEHandler(events <-chan string) http.Handler
```

### Learning goal
- What you are building: func SSEHandler(events <-chan string) http.Handler as a reliable contract. Focus: net/http, streaming.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): writes each event as a "data: <event>\n\n" block with Content-Type text/event-stream, flushes per event, and ends the stream when the channel closes.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- text/event-stream content type
- data: prefix per event
- flush per event
- ends on channel close

## Prior reading
- [SSE spec (WHATWG)](https://html.spec.whatwg.org/multipage/server-sent-events.html)
- [Go net/http package](https://pkg.go.dev/net/http)

## What this kata is about (and why it matters)
- Core lesson: streaming contracts: set headers before writing, flush after each payload, and terminate cleanly.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
