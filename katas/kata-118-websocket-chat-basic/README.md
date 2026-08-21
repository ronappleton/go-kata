# Kata 118 — WebSocket Chat (basic)

**Focus:** net, protocol, encoding

## Your task
Implement:

```go
func HandleUpgrade(conn net.Conn) error; func ReadTextFrame(conn net.Conn) (string, error); func WriteTextFrame(conn net.Conn, text string) error; func HandleEcho(conn net.Conn) error
```

### Learning goal
- What you are building: func HandleUpgrade(conn net.Conn) error; func ReadTextFrame(conn net.Conn) (string, error); func WriteTextFrame(conn net.Conn, text string) error; func HandleEcho(conn net.Conn) error as a reliable contract. Focus: net, protocol, encoding.
- Why this matters in real projects: real systems depend on exact, testable behavior here.
- How this grows your Go skills: you practice invariants, edge cases, and test-first discipline.
- Definition of done (plain English): HandleUpgrade validates the HTTP upgrade and writes a correct 101 with the Sec-WebSocket-Accept key; ReadTextFrame decodes a masked client text frame; WriteTextFrame sends an unmasked server text frame; HandleEcho upgrades then echoes text frames until a close frame.

### Tips
- Write tests from the rules before implementation.
- Name edge cases explicitly: nil, empty, min, max.
- Keep logic linear; branch only when a rule requires it.

## Rules / Expectations
- correct 101 upgrade with valid accept key
- read masked text frames
- write unmasked text frames
- echo loop ends on close frame

## Prior reading
- [RFC 6455 (WebSocket)](https://www.rfc-editor.org/rfc/rfc6455)
- [Go net package](https://pkg.go.dev/net)
- [Go crypto/sha1 package](https://pkg.go.dev/crypto/sha1)

## What this kata is about (and why it matters)
- Core lesson: protocols are byte contracts: every header, mask, and opcode must be exact.
- After this kata, you should be able to justify every rule with a test.

## What you must submit for marking
- `kata.go`
- `kata_test.go`

## Run
```bash
go test ./...
```
