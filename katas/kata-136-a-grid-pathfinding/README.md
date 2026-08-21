# Kata 136 — A* Grid Pathfinding

**Focus:** algorithms, grids

## Your task
Implement:

```go
func FindPath(start, goal [2]int, grid [][]int) ([][2]int, error)
```

### Learning goal
- What you are building: func FindPath(start, goal [2]int, grid [][]int) ([][2]int, error) as a reliable contract. Focus: algorithms, grids.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): returns a path of [row,col] cells from start to goal moving only through passable cells (0), with consecutive steps adjacent (4-directional); error when no path exists.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- only passable cells
- 4-directional adjacency
- start and goal endpoints
- unreachable => error

## Prior reading
- [A* search algorithm (Wikipedia)](https://en.wikipedia.org/wiki/A*_search_algorithm)
- [Amit Patel's A* tutorial](https://www.redblobgames.com/pathfinding/a-star/introduction.html)

## What this kata is about (and why it matters)
- Core lesson: A* is admissible with a Manhattan heuristic; the test checks path validity, not optimality, so any correct path passes.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
