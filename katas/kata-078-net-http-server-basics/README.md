# Kata 078 — net/http Server Basics

**Focus:** net/http, http.HandleFunc, http.ListenAndServe, request/response

## Your task
Implement:

```go
func CreateMux() *http.ServeMux
```

### Learning goal
- What you are practicing: building HTTP handlers with Go's standard library.
- Why this matters: `net/http` is the foundation of every Go web service. Understanding handlers, muxes, and request/response is essential.
- How this grows your Go skills: you learn the handler interface, request parsing, and response writing.

### Tips
- Use `http.NewRequest` and `httptest.NewRecorder` for testing.
- Register handlers with `mux.HandleFunc`.

## Rules / Expectations
- /health => 200 OK
- /hello?name=X => "Hello, X!"

## What this kata is about (and why it matters)
- Core lesson: Go's net/http is powerful enough for production use.
- After this kata, you can build HTTP services without any framework.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
