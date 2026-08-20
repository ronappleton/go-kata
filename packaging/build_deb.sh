#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
VERSION="${1:-0.1.0}"
ARCH="amd64"
BUILD_DIR="$ROOT_DIR/.build/deb"
PACKAGE_DIR="$BUILD_DIR/gokatas_${VERSION}_${ARCH}"
OUT="$ROOT_DIR/dist/gokatas_${VERSION}_${ARCH}.deb"

cd "$ROOT_DIR"
rm -rf "$BUILD_DIR" "$ROOT_DIR/dist"
mkdir -p "$PACKAGE_DIR/DEBIAN" "$PACKAGE_DIR/usr/bin" \
  "$PACKAGE_DIR/usr/share/applications" "$PACKAGE_DIR/usr/share/icons/hicolor/scalable/apps" \
  "$PACKAGE_DIR/usr/share/gokatas" "$PACKAGE_DIR/usr/lib/gokatas"

CGO_ENABLED=1 go build -trimpath -tags gtk4 -ldflags "-s -w -X main.version=$VERSION" \
  -o "$PACKAGE_DIR/usr/bin/gokatas" ./apps/learner-desktop

cp -R tracks katas "$PACKAGE_DIR/usr/share/gokatas/"
cp packaging/gokatas.desktop "$PACKAGE_DIR/usr/share/applications/"
cp packaging/gokatas.svg "$PACKAGE_DIR/usr/share/icons/hicolor/scalable/apps/gokatas.svg"
cp packaging/runner/setup.sh "$PACKAGE_DIR/usr/lib/gokatas/setup-runner.sh"
cp packaging/runner/Containerfile "$PACKAGE_DIR/usr/lib/gokatas/Containerfile"
chmod 0755 "$PACKAGE_DIR/usr/bin/gokatas" "$PACKAGE_DIR/usr/lib/gokatas/setup-runner.sh"

cat > "$PACKAGE_DIR/DEBIAN/control" <<CONTROL
Package: gokatas
Version: $VERSION
Section: devel
Priority: optional
Architecture: $ARCH
Maintainer: GoKatas Maintainers <maintainers@example.invalid>
Depends: libgtk-4-1, podman
Description: Native Go learning workbench
 GoKatas is a native Ubuntu GTK4 application for learning Go with
 rootless Podman evaluation of learner code.
CONTROL

mkdir -p "$ROOT_DIR/dist"
dpkg-deb --build --root-owner-group "$PACKAGE_DIR" "$OUT" >/dev/null
echo "$OUT"