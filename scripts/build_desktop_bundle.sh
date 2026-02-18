#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

export GOCACHE="${GOCACHE:-$ROOT_DIR/.gocache}"

TARGETS=(
  "darwin/arm64"
  "darwin/amd64"
  "linux/amd64"
  "windows/amd64"
)

for target in "${TARGETS[@]}"; do
  IFS="/" read -r GOOS GOARCH <<<"$target"
  OUT_DIR="$ROOT_DIR/dist/learner-studio/$GOOS-$GOARCH"
  rm -rf "$OUT_DIR"
  mkdir -p "$OUT_DIR"

  SERVER_BIN="learner-studio-server"
  DESKTOP_BIN="learner-studio"
  if [[ "$GOOS" == "windows" ]]; then
    SERVER_BIN+=".exe"
    DESKTOP_BIN+=".exe"
  fi

  echo "==> Building $GOOS/$GOARCH"
  GOOS="$GOOS" GOARCH="$GOARCH" CGO_ENABLED=0 go build -o "$OUT_DIR/$SERVER_BIN" ./apps/learner-studio
  GOOS="$GOOS" GOARCH="$GOARCH" CGO_ENABLED=0 go build -o "$OUT_DIR/$DESKTOP_BIN" ./apps/learner-desktop

  cp README.md "$OUT_DIR/README.md"
  cp -R tracks "$OUT_DIR/"
  cp -R katas "$OUT_DIR/"
  cp -R assets "$OUT_DIR/"

done

echo "Desktop bundles written to dist/learner-studio/"
