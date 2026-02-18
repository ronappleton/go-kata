#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-$ROOT_DIR/.gocache}"

echo "==> Running learner platform tests"
go test ./apps/learner-cli ./apps/learner-studio ./apps/learner-desktop ./internal/learning/...

echo "==> Running CLI smoke checks"
LIST_OUT="$(go run ./apps/learner-cli list)"
echo "$LIST_OUT" | grep -q "go-core-100"
echo "$LIST_OUT" | grep -q "Overall progress:"

SHOW_OUT="$(go run ./apps/learner-cli show --kata 001)"
echo "$SHOW_OUT" | grep -q "Build Greeting"
echo "$SHOW_OUT" | grep -q "Expected behavior:"

echo "==> Running desktop launcher compile smoke check"
go run ./apps/learner-desktop -h >/dev/null 2>&1
echo "==> Checking native launcher script exists"
test -x ./scripts/run_desktop_native.sh

echo "Operability checks passed."
