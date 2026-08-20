# GoKatas — Linux Build Container
# Multi-stage: build in Ubuntu 24.04 with GTK4 dev libs
# Usage: docker build -f build/Dockerfile -t gokatas-builder .
#        docker cp $(docker create gokatas-builder):/app/dist/bin/gokatas .

# ── Stage 1: Build ─────────────────────────────────────────────
FROM ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive

# Install build dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential pkg-config ca-certificates \
    libgtk-4-dev libglib2.0-dev libcairo2-dev libpango1.0-dev \
    libgdk-pixbuf-2.0-dev \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ARG GO_VERSION=1.25.0
RUN curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/root/go"
ENV PATH="${GOPATH}/bin:${PATH}"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Version info from build args
ARG VERSION=dev
ARG COMMIT=unknown

# Build desktop app (GTK4)
RUN mkdir -p dist/bin && \
    CGO_ENABLED=1 go build -trimpath \
    -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \
    -tags gtk4 \
    -o dist/bin/gokatas ./apps/learner-desktop/

# Build CLI (no CGO)
RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \

# Build test
RUN go test ./...

# ── Stage 2: Runtime (minimal) ─────────────────────────────────
FROM ubuntu:24.04 AS runtime

RUN apt-get update && apt-get install -y --no-install-recommends \
    libgtk-4-1 libglib2.0-0 libcairo2 libpango-1.0-0 \
    libgdk-pixbuf-2.0-0 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/dist/bin/ /usr/local/bin/

ENTRYPOINT ["gokatas"]
