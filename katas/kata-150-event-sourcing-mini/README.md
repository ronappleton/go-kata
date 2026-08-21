# Kata 150 — Event Sourcing Mini

**Focus:** event-driven, design

## Your task
Implement:

```go
type Event struct { Type string; Data map[string]any }; type Store struct{}; func NewStore() *Store; func (s *Store) Append(e Event) int; func (s *Store) Events() []Event; func ApplyAccount(events []Event) (int, error)
```

### Learning goal
- What you are building: type Event struct { Type string; Data map[string]any }; type Store struct{}; func NewStore() *Store; func (s *Store) Append(e Event) int; func (s *Store) Events() []Event; func ApplyAccount(events []Event) (int, error) as a reliable contract. Focus: event-driven, design.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): Store appends events in order and returns monotonically increasing sequence numbers; ApplyAccount replays deposit/withdraw events into a balance and errors on unknown types or negative balance.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- ordered event log
- sequence numbers
- replay to a balance
- invalid events error

## Prior reading
- [Event sourcing (Martin Fowler)](https://martinfowler.com/eaaDev/EventSourcing.html)
- [Event sourcing (Wikipedia)](https://en.wikipedia.org/wiki/Event_sourcing)

## What this kata is about (and why it matters)
- Core lesson: event sourcing treats the log as the source of truth; state is always derived by replay.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
