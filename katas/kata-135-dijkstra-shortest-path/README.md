# Kata 135 — Dijkstra Shortest Path

**Focus:** graphs, algorithms

## Your task
Implement:

```go
func ShortestPath(edges map[string]map[string]int, start, goal string) (int, []string, error)
```

### Learning goal
- What you are building: func ShortestPath(edges map[string]map[string]int, start, goal string) (int, []string, error) as a reliable contract. Focus: graphs, algorithms.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): returns the minimum total cost and a path [start...goal] of nodes where every consecutive pair is an edge and the cost sums to the minimum; error when goal is unreachable.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- minimum total cost
- valid path (consecutive edges)
- unreachable => error
- start == goal => cost 0, path [start]

## Prior reading
- [Dijkstra algorithm (Wikipedia)](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
- [Go container/heap package](https://pkg.go.dev/container/heap)

## What this kata is about (and why it matters)
- Core lesson: shortest-path correctness is checkable: verify the cost AND that the returned path is real and sums to it.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
