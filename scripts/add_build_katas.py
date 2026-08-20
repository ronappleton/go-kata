#!/usr/bin/env python3
"""Add katas for Makefile and Docker development setup."""
import json, os

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
KATAS_DIR = os.path.join(ROOT, "katas")
TRACK_FILE = os.path.join(ROOT, "tracks", "go-core-100", "track.json")

NEW_KATAS = [
    {
        "id": "168",
        "slug": "makefile-basics",
        "title": "Makefile Basics",
        "focus": "Make targets, variables, dependencies, phony targets, build/test/clean patterns",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["makefile", "build", "targets", "variables", "phony"],
        "prerequisites": ["000", "005"],
        "estimated_minutes": 30,
        "flashcards": [
            {"front": "What is a Makefile target?", "back": "A named rule that specifies how to build a file or perform an action. Format: target: dependencies\\n\\tcommand"},
            {"front": "What does .PHONY mean in a Makefile?", "back": "Declares targets that don't produce files. Without .PHONY, Make checks if a file with that name exists and skips the rule if it does."},
            {"front": "What is $@ in a Makefile recipe?", "back": "The automatic variable for the target name. $< is the first prerequisite. $? is prerequisites newer than the target."},
            {"front": "What is the difference between = and := in Makefile variables?", "back": "= is recursive (evaluated when used). := is simple (evaluated immediately). Use := for performance and predictable behavior."}
        ],
        "quiz_questions": [
            {"question": "What .PHONY prevents Make from doing?", "options": ["Running commands", "Checking for file existence", "Using variables", "Including other files"], "answer": "Checking for file existence"},
            {"question": "What does 'make clean' typically do?", "options": ["Delete build artifacts", "Run tests", "Install dependencies", "Format code"], "answer": "Delete build artifacts"},
            {"question": "Which variable holds the first prerequisite?", "options": ["$@", "$<", "$?", "$^"], "answer": "$<"}
        ],
        "kata_go": '''package kata

import (
\t"strings"
)

// ParseMakefile reads a Makefile string and returns all target names.
// Rules:
// - A target line has format: "target:" at the start (possibly with tabs/spaces before)
// - Ignore lines starting with # (comments)
// - Ignore .PHONY declarations
// - Ignore variable assignments (lines with = but no :)
// - Return target names in order of appearance
func ParseMakefile(content string) []string {
\t// TODO: Parse target names from Makefile content
\treturn nil
}

// ExtractVariables reads a Makefile and returns variable assignments.
// Rules:
// - Variable lines have format: "VAR = value" or "VAR := value"
// - Return map of variable name to value
// - Trim whitespace from both name and value
// - Ignore recursive vs simple assignment differences
func ExtractVariables(content string) map[string]string {
\t// TODO: Parse variable assignments
\treturn nil
}

// ValidateTargets checks if a Makefile has required targets.
// Rules:
// - required is a list of target names that must exist
// - Return missing target names (empty slice means all present)
// - Case-sensitive comparison
func ValidateTargets(content string, required []string) []string {
\t// TODO: Check which required targets are missing
\treturn nil
}

// GenerateMakefile creates a simple Makefile for a Go project.
// Rules:
// - Include .PHONY for all targets
// - Include build, test, clean, and lint targets
// - build: go build -o bin/app ./cmd/...
// - test: go test ./...
// - clean: rm -rf bin/
// - lint: go vet ./...
// - Format with proper tabs for recipes
func GenerateMakefile() string {
\tvar b strings.Builder
\t// TODO: Generate a proper Makefile
\treturn b.String()
}
''',
        "kata_test_go": '''package kata

import (
\t"strings"
\t"testing"
)

func TestParseMakefile(t *testing.T) {
\tinput := `
# Build targets
.PHONY: build test clean

build: dep1 dep2
\\tgo build -o bin/app .

test:
\\tgo test ./...

clean:
\\tram -rf bin/
`
\tgot := ParseMakefile(input)
\twant := []string{"build", "test", "clean"}
\tif len(got) != len(want) {
\t\tt.Fatalf("ParseMakefile() returned %d targets, want %d: %v", len(got), len(want), got)
\t}
\tfor i, w := range want {
\t\tif got[i] != w {
\t\t\tt.Errorf("ParseMakefile()[%d] = %q, want %q", i, got[i], w)
\t\t}
\t}
}

func TestExtractVariables(t *testing.T) {
\tinput := `
APP = myapp
VERSION := 1.0.0
GOBIN := ./bin
`
\tgot := ExtractVariables(input)
\tif got["APP"] != "myapp" {
\t\tt.Errorf("ExtractVariables()[APP] = %q, want %q", got["APP"], "myapp")
\t}
\tif got["VERSION"] != "1.0.0" {
\t\tt.Errorf("ExtractVariables()[VERSION] = %q, want %q", got["VERSION"], "1.0.0")
\t}
\tif got["GOBIN"] != "./bin" {
\t\tt.Errorf("ExtractVariables()[GOBIN] = %q, want %q", got["GOBIN"], "./bin")
\t}
}

func TestValidateTargets(t *testing.T) {
\tcontent := `
build:
\\tgo build .
test:
\\tgo test ./...
`
\tmissing := ValidateTargets(content, []string{"build", "test", "clean"})
\tif len(missing) != 1 || missing[0] != "clean" {
\t\tt.Errorf("ValidateTargets() = %v, want [clean]", missing)
\t}
\tmissing = ValidateTargets(content, []string{"build", "test"})
\tif len(missing) != 0 {
\t\tt.Errorf("ValidateTargets() = %v, want []", missing)
\t}
}

func TestGenerateMakefile(t *testing.T) {
\tmf := GenerateMakefile()
\tif !strings.Contains(mf, ".PHONY:") {
\t\tt.Error("GenerateMakefile() missing .PHONY")
\t}
\tif !strings.Contains(mf, "build:") {
\t\tt.Error("GenerateMakefile() missing build target")
\t}
\tif !strings.Contains(mf, "test:") {
\t\tt.Error("GenerateMakefile() missing test target")
\t}
\tif !strings.Contains(mf, "clean:") {
\t\tt.Error("GenerateMakefile() missing clean target")
\t}
\tif !strings.Contains(mf, "lint:") {
\t\tt.Error("GenerateMakefile() missing lint target")
\t}
}
''',
        "readme": '''# Kata 168 — Makefile Basics

**Focus:** Make targets, variables, dependencies, phony targets, build/test/clean patterns

## Your task

Parse and generate Makefiles for Go projects.

### Learning goal
- What you are practicing: understanding Make syntax, targets, variables, and common Go build patterns.
- Why this matters: every Go project needs a consistent build interface. Makefiles standardize build/test/lint/clean.
- How this grows your Go skills: you'll create reproducible build systems that work on any machine.

## Rules / Expectations
- ParseMakefile extracts target names in order
- ExtractVariables returns variable assignments
- ValidateTargets finds missing required targets
- GenerateMakefile creates a complete Go project Makefile

## What this kata is about (and why it matters)
- Core lesson: Makefiles are the universal build system. Understanding targets, dependencies, and recipes is essential.
- After this kata, you should write Makefiles for any Go project.

## What you must submit for marking
- `kata.go`
'''
    },
    {
        "id": "169",
        "slug": "docker-dev-setup",
        "title": "Docker Development Environment",
        "focus": "Dockerfile, docker-compose, dev containers, volume mounts, healthchecks",
        "stage": "foundation",
        "category": "setup-and-toolchain",
        "level": "junior",
        "tags": ["docker", "dockerfile", "docker-compose", "containers", "dev-environment"],
        "prerequisites": ["000", "148"],
        "estimated_minutes": 35,
        "flashcards": [
            {"front": "What is the difference between COPY and ADD in Dockerfile?", "back": "COPY copies files from host to container. ADD also supports URLs and auto-extracts tar files. Use COPY unless you need ADD's extras."},
            {"front": "What is a multi-stage Dockerfile?", "back": "Using multiple FROM statements to build in one stage and copy only the final binary to a slim runtime stage. Reduces image size dramatically."},
            {"front": "What does docker-compose up -d do?", "back": "Starts all services defined in docker-compose.yml in detached mode (background). Creates containers, networks, and volumes as defined."},
            {"front": "What is a Docker healthcheck?", "back": "A command that Docker runs periodically to check if the container is healthy. Unhealthy containers can be restarted or have traffic routed away."}
        ],
        "quiz_questions": [
            {"question": "Which Dockerfile instruction runs at build time?", "options": ["CMD", "ENTRYPOINT", "RUN", "EXPOSE"], "answer": "RUN"},
            {"question": "What does EXPOSE do in a Dockerfile?", "options": ["Opens a port on the host", "Documents which ports the container uses", "Creates a network bridge", "Enables port forwarding"], "answer": "Documents which ports the container uses"},
            {"question": "How do you reduce Docker image size for Go?", "options": ["Use alpine base image", "Multi-stage build with scratch/distroless", "Remove go.mod", "Use go mod vendor"], "answer": "Multi-stage build with scratch/distroless"}
        ],
        "kata_go": '''package kata

import (
\t"fmt"
\t"strings"
)

// GenerateDockerfile creates a multi-stage Dockerfile for a Go project.
// Rules:
// - Stage 1 (builder): use golang:1.24-alpine, copy source, build binary
// - Stage 2 (runtime): use gcr.io/distroless/static-debian12, copy binary
// - Binary path: /app/server
// - Final stage exposes port 8080
// - Use proper WORKDIR and COPY ordering
func GenerateDockerfile() string {
\t// TODO: Generate multi-stage Dockerfile
\treturn ""
}

// GenerateCompose creates a docker-compose.yml for a Go app with Redis.
// Rules:
// - Service "app": build from current dir, port 8080:8080, depends on redis
// - Service "redis": image redis:7-alpine, port 6379:6379
// - Network: app-network (bridge)
// - Volume: redis-data for redis persistence
func GenerateCompose() string {
\t// TODO: Generate docker-compose.yml content
\treturn ""
}

// ParseDockerfile extracts all RUN commands from a Dockerfile.
// Rules:
// - Return the command part of each RUN instruction
// - Handle both exec form RUN ["cmd"] and shell form RUN cmd
// - Ignore comments and blank lines
func ParseDockerfile(content string) []string {
\t// TODO: Extract RUN commands
\treturn nil
}

// ValidateCompose checks a docker-compose.yml for required fields.
// Rules:
// - Check for "services:" key
// - Check each service has either "image:" or "build:"
// - Check for "ports:" in each service
// - Return list of issues (empty = valid)
func ValidateCompose(content string) []string {
\t// TODO: Validate compose file structure
\treturn nil
}

// GenerateHealthcheck creates a HEALTHCHECK instruction.
// Rules:
// - For HTTP services: curl -f http://localhost:PORT/health || exit 1
// - For TCP services: nc -z localhost PORT || exit 1
// - Include interval=30s, timeout=5s, retries=3
func GenerateHealthcheck(port int, checkType string) string {
\tif checkType == "http" {
\t\treturn fmt.Sprintf("HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD curl -f http://localhost:%d/health || exit 1", port)
\t}
\treturn fmt.Sprintf("HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD nc -z localhost %d || exit 1", port)
}
''',
        "kata_test_go": '''package kata

import (
\t"strings"
\t"testing"
)

func TestGenerateDockerfile(t *testing.T) {
\tdf := GenerateDockerfile()
\tif !strings.Contains(df, "FROM golang:") {
\t\tt.Error("Dockerfile missing golang base stage")
\t}
\tif !strings.Contains(df, "FROM") && strings.Count(df, "FROM") < 2 {
\t\tt.Error("Dockerfile should have multiple stages")
\t}
\tif !strings.Contains(df, "COPY") {
\t\tt.Error("Dockerfile missing COPY instruction")
\t}
\tif !strings.Contains(df, "EXPOSE") {
\t\tt.Error("Dockerfile missing EXPOSE")
\t}
\tif !strings.Contains(df, "go build") {
\t\tt.Error("Dockerfile missing go build command")
\t}
}

func TestGenerateCompose(t *testing.T) {
\tcomp := GenerateCompose()
\tif !strings.Contains(comp, "services:") {
\t\tt.Error("Compose missing services key")
\t}
\tif !strings.Contains(comp, "redis") {
\t\tt.Error("Compose missing redis service")
\t}
\tif !strings.Contains(comp, "8080") {
\t\tt.Error("Compose missing port 8080")
\t}
\tif !strings.Contains(comp, "6379") {
\t\tt.Error("Compose missing port 6379")
\t}
}

func TestParseDockerfile(t *testing.T) {
\tinput := `FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .
FROM alpine:3.19
COPY --from=builder /app/server /server
RUN chmod +x /server
CMD ["/server"]
`
\tgot := ParseDockerfile(input)
\t// Should find: "go mod download", "go build -o server .", "chmod +x /server"
\tif len(got) < 2 {
\t\tt.Errorf("ParseDockerfile() returned %d commands, want at least 2: %v", len(got), got)
\t}
\tfound := false
\tfor _, cmd := range got {
\t\tif strings.Contains(cmd, "go build") {
\t\t\tfound = true
\t\t\tbreak
\t\t}
\t}
\tif !found {
\t\tt.Errorf("ParseDockerfile() didn't find go build command: %v", got)
\t}
}

func TestValidateCompose(t *testing.T) {
\tvalid := `
services:
  app:
    build: .
    ports:
      - "8080:8080"
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
`
\tissues := ValidateCompose(valid)
\tif len(issues) != 0 {
\t\tt.Errorf("ValidateCompose() on valid content = %v, want []", issues)
\t}

\tinvalid := `
app:
  build: .
`
\tissues = ValidateCompose(invalid)
\tif len(issues) == 0 {
\t\tt.Error("ValidateCompose() on invalid content = [], want issues")
\t}
}

func TestGenerateHealthcheck(t *testing.T) {
\thc := GenerateHealthcheck(8080, "http")
\tif !strings.Contains(hc, "HEALTHCHECK") {
\t\tt.Error("GenerateHealthcheck() missing HEALTHCHECK")
\t}
\tif !strings.Contains(hc, "8080") {
\t\tt.Error("GenerateHealthcheck() missing port")
\t}
\tif !strings.Contains(hc, "curl") {
\t\tt.Error("GenerateHealthcheck() missing curl for http type")
\t}

\thc = GenerateHealthcheck(6379, "tcp")
\tif !strings.Contains(hc, "nc") {
\t\tt.Error("GenerateHealthcheck() missing nc for tcp type")
\t}
}
''',
        "readme": '''# Kata 169 — Docker Development Environment

**Focus:** Dockerfile, docker-compose, dev containers, volume mounts, healthchecks

## Your task

Generate and validate Docker configurations for Go projects.

### Learning goal
- What you are practicing: writing Dockerfiles, docker-compose files, and understanding container workflows.
- Why this matters: Docker standardizes development environments and deployments. Every team uses it.
- How this grows your Go skills: you'll containerize Go apps efficiently with multi-stage builds.

## Rules / Expectations
- GenerateDockerfile creates a proper multi-stage build
- GenerateCompose creates a working docker-compose.yml
- ParseDockerfile extracts RUN commands
- ValidateCompose checks structure

## What this kata is about (and why it matters)
- Core lesson: Dockerfiles and docker-compose are the standard for Go deployment. Multi-stage builds keep images small.
- After this kata, you should containerize any Go application.

## What you must submit for marking
- `kata.go`
'''
    }
]

for kata in NEW_KATAS:
    kid = kata["id"]
    slug = kata["slug"]
    dirname = f"kata-{kid.zfill(3)}-{slug}"
    dirpath = os.path.join(KATAS_DIR, dirname)
    os.makedirs(dirpath, exist_ok=True)
    
    meta = {
        "id": kid,
        "title": kata["title"],
        "slug": slug,
        "focus": kata["focus"],
        "stage": kata["stage"],
        "category": kata["category"],
        "level": kata["level"],
        "tags": kata["tags"],
        "prerequisites": kata["prerequisites"],
        "estimated_minutes": kata["estimated_minutes"],
        "evaluator_status": "ready",
        "flashcards": kata["flashcards"],
        "quiz_questions": kata["quiz_questions"]
    }
    with open(os.path.join(dirpath, "kata.json"), "w") as f:
        json.dump(meta, f, indent=2)
    
    with open(os.path.join(dirpath, "kata.go.txt"), "w") as f:
        f.write(kata["kata_go"])
    
    with open(os.path.join(dirpath, "kata_test.go.txt"), "w") as f:
        f.write(kata["kata_test_go"])
    
    with open(os.path.join(dirpath, "README.md"), "w") as f:
        f.write(kata["readme"])
    
    print(f"Created: {dirname}")

# Update track.json
with open(TRACK_FILE) as f:
    track = json.load(f)

new_ids = [k["id"] for k in NEW_KATAS]
for stage in track["stages"]:
    if stage["id"] == "foundation":
        for cat in stage["categories"]:
            if cat["id"] == "setup-and-toolchain":
                # Append after existing
                cat["kata_ids"].extend(new_ids)
                print(f"Updated setup-and-toolchain: {len(cat['kata_ids'])} katas")
                break

with open(TRACK_FILE, "w") as f:
    json.dump(track, f, indent=2)

print("Done!")
