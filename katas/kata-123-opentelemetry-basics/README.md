# Kata 123 — OpenTelemetry Basics

**Focus:** tracing, spans, context propagation, observability

## Your task
Implement a simplified tracing system.

### Learning goal
- What you are practicing: distributed tracing concepts with spans and context propagation.
- Why this matters: in production systems, you need to trace requests across services. OpenTelemetry is the standard.
- How this grows your Go skills: you learn how context carries trace information through call chains.

### Tips
- Spans have a name, parent, and lifecycle (start/finish).
- Context propagation links parent and child spans.
- In real code, use go.opentelemetry.io/otel.

## Rules / Expectations
- StartSpan creates a span
- FinishSpan marks it done
- Parent-child relationships are tracked

## What this kata is about (and why it matters)
- Core lesson: tracing is essential for understanding distributed systems.
- After this kata, you can add observability to any Go service.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
