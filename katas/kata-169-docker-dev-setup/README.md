# Kata 169 — Docker Development Environment

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
