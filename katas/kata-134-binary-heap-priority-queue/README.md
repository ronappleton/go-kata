# Kata 134 — Binary Heap Priority Queue

**Focus:** data structures

## Your task
Implement:

```go
func NewPriorityQueue(minHeap bool) *PriorityQueue; func (q *PriorityQueue) Push(priority int, value string); func (q *PriorityQueue) Pop() (priority int, value string, ok bool); func (q *PriorityQueue) Len() int; func (q *PriorityQueue) Peek() (priority int, value string, ok bool)
```

### Learning goal
- What you are building: func NewPriorityQueue(minHeap bool) *PriorityQueue; func (q *PriorityQueue) Push(priority int, value string); func (q *PriorityQueue) Pop() (priority int, value string, ok bool); func (q *PriorityQueue) Len() int; func (q *PriorityQueue) Peek() (priority int, value string, ok bool) as a reliable contract. Focus: data structures.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): min-heap pops lowest priority first, max-heap pops highest; Pop/Peek return ok=false when empty; Len reflects size.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- min and max ordering
- ok=false on empty Pop/Peek
- Len tracks size
- stable for duplicate priorities (any order)

## Prior reading
- [Binary heap (Wikipedia)](https://en.wikipedia.org/wiki/Binary_heap)
- [Go container/heap package](https://pkg.go.dev/container/heap)

## What this kata is about (and why it matters)
- Core lesson: heaps guarantee the extreme element in O(log n); the invariant to test is pop order, not internal layout.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
