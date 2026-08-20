# GoKatas - Cross-Platform Build System
# GTK4 requires CGO so each platform must be built natively.
# Prerequisites:
#   Linux:   sudo apt install libgtk-4-dev build-essential pkg-config
#   macOS:   brew install gtk4 pkg-config
#   Windows: MSYS2 MinGW64 + pacman -S mingw-w64-x86_64-gtk4
#
# Quick start:
#   make build          - build for current platform
#   make test           - run all tests
#   make docker-build   - reproducible Linux build in Docker

VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GOOS     := $(shell go env GOOS)
GOARCH   := $(shell go env GOARCH)
LD       := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildDate=$(BUILD)
DIST     := dist
BIN      := $(DIST)/bin

.PHONY: all build test test-race test-cover vet fmt lint clean help
.PHONY: build-linux-amd64 build-darwin-arm64 build-darwin-amd64
.PHONY: build-windows-amd64 build-windows-arm64 package-windows
.PHONY: docker-build package package-darwin install
.PHONY: install-deps install-deps-macos install-deps-windows
.PHONY: gen-katas sync-content

all: build

# --- Build: Desktop App (GTK4) ---

build: ## Build desktop app for current platform
	CGO_ENABLED=1 go build -trimpath -ldflags "$(LD)" -tags gtk4 -o $(BIN)/gokatas ./apps/learner-desktop/

# --- Linux ---

build-linux-amd64: ## Build for Linux amd64
	mkdir -p $(DIST)/gokatas-linux-amd64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -trimpath -ldflags "$(LD)" -tags gtk4 -o $(DIST)/gokatas-linux-amd64/gokatas ./apps/learner-desktop/

# --- macOS ---

build-darwin-arm64: ## Build for macOS arm64 (must run on macOS)
	mkdir -p $(DIST)/gokatas-darwin-arm64
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -trimpath -ldflags "$(LD)" -tags gtk4 -o $(DIST)/gokatas-darwin-arm64/gokatas ./apps/learner-desktop/

build-darwin-amd64: ## Build for macOS amd64 (must run on macOS)
	mkdir -p $(DIST)/gokatas-darwin-amd64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -trimpath -ldflags "$(LD)" -tags gtk4 -o $(DIST)/gokatas-darwin-amd64/gokatas ./apps/learner-desktop/

# --- Windows (run from MSYS2 MinGW64 shell) ---

build-windows-amd64: ## Build for Windows amd64
	mkdir -p $(DIST)/gokatas-windows-amd64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
	  go build -trimpath -ldflags "$(LD)" -tags gtk4 \
	  -o $(DIST)/gokatas-windows-amd64/gokatas.exe ./apps/learner-desktop/

build-windows-arm64: ## Build for Windows arm64 (run from MSYS2 CLANGARM64 shell)
	mkdir -p $(DIST)/gokatas-windows-arm64
	GOOS=windows GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-w64-mingw32-gcc \
	  go build -trimpath -ldflags "$(LD)" -tags gtk4 \
	  -o $(DIST)/gokatas-windows-arm64/gokatas.exe ./apps/learner-desktop/

package-windows: build-windows-amd64 ## Package Windows as zip
	mkdir -p $(DIST)/gokatas-windows
	cp $(DIST)/gokatas-windows-amd64/gokatas.exe $(DIST)/gokatas-windows/
	cd $(DIST) && zip -r gokatas-$(VERSION)-windows-amd64.zip gokatas-windows/
	@echo "  -> $(DIST)/gokatas-$(VERSION)-windows-amd64.zip"

# --- Docker (Linux, reproducible) ---

docker-build: ## Build Linux amd64 in Docker
	docker build -f build/Dockerfile --build-arg VERSION=$(VERSION) --build-arg COMMIT=$(COMMIT) -t gokatas-builder .
	mkdir -p $(DIST)/gokatas-linux-amd64
	docker cp $$(docker create gokatas-builder):/app/dist/bin/gokatas $(DIST)/gokatas-linux-amd64/gokatas
	@echo "  -> $(DIST)/gokatas-linux-amd64/gokatas"

# --- Test ---

test: ## Run all tests
	go test ./...

test-race: ## Run tests with race detector
	go test -race ./...

test-cover: ## Run tests with coverage
	go test -coverprofile=$(DIST)/coverage.out ./...
	go tool cover -html=$(DIST)/coverage.out -o $(DIST)/coverage.html

# --- Lint ---

vet: ## Run go vet
	go vet ./...

fmt: ## Format all Go files
	gofmt -w .

lint: vet fmt ## Run vet + fmt

# --- Content ---

gen-katas: ## Regenerate embedded kata data
	go run scripts/gen_kata_data.go katas internal/learning/katas/katas.go

sync-content: ## Generate content repo for remote distribution
	go run scripts/sync_content.go . $(DIST)/gokatas-content

# --- Package ---

package: build ## Package current platform as tarball
	mkdir -p $(DIST)/release
	cp $(BIN)/gokatas $(DIST)/release/
	tar -czf $(DIST)/gokatas-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz -C $(DIST)/release .
	@echo "  -> $(DIST)/gokatas-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz"

package-darwin: build ## Package macOS as .app bundle
	mkdir -p $(DIST)/GoKatas.app/Contents/MacOS
	cp $(BIN)/gokatas $(DIST)/GoKatas.app/Contents/MacOS/
	tar -czf $(DIST)/gokatas-$(VERSION)-darwin-universal.tar.gz -C $(DIST) GoKatas.app
	@echo "  -> $(DIST)/gokatas-$(VERSION)-darwin-universal.tar.gz"

# --- Install ---

install: build ## Install to /usr/local/bin
	install -m 755 $(BIN)/gokatas /usr/local/bin/gokatas

install-deps: ## Install Linux build dependencies
	apt-get update && apt-get install -y build-essential pkg-config libgtk-4-dev libglib2.0-dev

install-deps-macos: ## Install macOS build dependencies
	brew install gtk4 pkg-config glib cairo pango gdk-pixbuf

install-deps-windows: ## Install Windows build dependencies (run in MSYS2 MinGW64)
	MSYSTEM=MINGW64 pacman -S --noconfirm mingw-w64-x86_64-gtk4 mingw-w64-x86_64-pkg-config mingw-w64-x86_64-gcc

# --- Clean ---

clean: ## Remove build artifacts
	rm -rf $(DIST)

# --- Help ---

help: ## Show available targets
	@grep -E "^[a-zA-Z_-]+:.*?## " $(MAKEFILE_LIST) | sed "s/:.*## /\t/"
