#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-$ROOT_DIR/.gocache}"

echo "==> Running learner platform tests"
go test ./apps/learner-cli ./apps/learner-studio ./internal/learning/...

echo "==> Running CLI smoke checks"
LIST_OUT="$(go run ./apps/learner-cli list)"
echo "$LIST_OUT" | grep -q "Go Core 100"
echo "$LIST_OUT" | grep -q "Overall progress:"

SHOW_OUT="$(go run ./apps/learner-cli show --kata 001)"
echo "$SHOW_OUT" | grep -q "001 — FizzBuzz"
echo "$SHOW_OUT" | grep -q "Expected behavior:"

echo "Operability checks passed."
