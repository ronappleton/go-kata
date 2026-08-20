#!/usr/bin/env python3
"""Migrate GoKatas curriculum to 4-stage junior→lead progression.

This script:
1. Defines old_id → new_id mapping
2. Renames kata directories
3. Generates new track.json
4. Adds stage/category/tags/prerequisites/flashcards/quiz metadata to each kata.json
"""
import json
import os
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
KATAS_DIR = ROOT / "katas"
TRACK_DIR = ROOT / "tracks" / "go-core-100"

# ── Mapping: (new_id, old_id, stage, category) ──
# Where old_id is None, the kata is NEW and will be created later.
MAPPING = [
    # ═══════════════════════════════════════════════
    # STAGE 1: FOUNDATION (Junior)
    # ═══════════════════════════════════════════════

    # Category: setup-and-toolchain
    ("000", None,    "foundation", "setup-and-toolchain"),   # NEW: Go Setup & First Program
    ("001", "001",   "foundation", "setup-and-toolchain"),   # Build Greeting
    ("002", "002",   "foundation", "setup-and-toolchain"),   # Parse Whole Number
    ("003", "003",   "foundation", "setup-and-toolchain"),   # Basic Calculator
    ("004", "004",   "foundation", "setup-and-toolchain"),   # Even Odd Label
    ("005", None,    "foundation", "setup-and-toolchain"),   # NEW: Go Vet & Lint

    # Category: language-basics
    ("006", "031",   "foundation", "language-basics"),        # FizzBuzz
    ("007", "032",   "foundation", "language-basics"),        # Sum of Integers
    ("008", "033",   "foundation", "language-basics"),        # Reverse String
    ("009", "034",   "foundation", "language-basics"),        # Palindrome Check
    ("010", "035",   "foundation", "language-basics"),        # Count Vowels
    ("011", "005",   "foundation", "language-basics"),        # Clamp Value
    ("012", "006",   "foundation", "language-basics"),        # Sum First N
    ("013", "008",   "foundation", "language-basics"),        # Repeat Text
    ("014", "013",   "foundation", "language-basics"),        # Grade Classifier
    ("015", "014",   "foundation", "language-basics"),        # Count True Values
    ("016", "017",   "foundation", "language-basics"),        # Budget Status
    ("017", "022",   "foundation", "language-basics"),        # Parse Positive Int
    ("018", "036",   "foundation", "language-basics"),        # Factorial
    ("019", "037",   "foundation", "language-basics"),        # Fibonacci

    # Category: data-structures
    ("020", "007",   "foundation", "data-structures"),        # Countdown Slice
    ("021", "009",   "foundation", "data-structures"),        # Character Frequency
    ("022", "010",   "foundation", "data-structures"),        # Average Integer
    ("023", "011",   "foundation", "data-structures"),        # Swap First and Last
    ("024", "012",   "foundation", "data-structures"),        # Remove Empty Strings
    ("025", "015",   "foundation", "data-structures"),        # Find First Index
    ("026", "016",   "foundation", "data-structures"),        # Safe Slice Range
    ("027", "018",   "foundation", "data-structures"),        # Merge Alternating Slices
    ("028", "019",   "foundation", "data-structures"),        # Word Lengths
    ("029", "020",   "foundation", "data-structures"),        # Incomplete Task Filter
    ("030", "038",   "foundation", "data-structures"),        # Filter Even
    ("031", "039",   "foundation", "data-structures"),        # Unique Strings
    ("032", "040",   "foundation", "data-structures"),        # Word Count
    ("033", "041",   "foundation", "data-structures"),        # Anagram Check
    ("034", "042",   "foundation", "data-structures"),        # Min/Max
    ("035", "043",   "foundation", "data-structures"),        # Rotate Slice
    ("036", "044",   "foundation", "data-structures"),        # Binary Search
    ("037", "045",   "foundation", "data-structures"),        # Merge Two Sorted Lists

    # Category: error-handling
    ("038", "030",   "foundation", "error-handling"),         # Find User Name (errors.Is, sentinel)
    ("039", None,    "foundation", "error-handling"),         # NEW: Error Wrapping Patterns
    ("040", None,    "foundation", "error-handling"),         # NEW: Interface Basics

    # Category: testing-fundamentals
    ("041", None,    "foundation", "testing-fundamentals"),   # NEW: Table-Driven Tests
    ("042", "029",   "foundation", "testing-fundamentals"),   # Increment Concurrently (sync basics)
    ("043", "028",   "foundation", "testing-fundamentals"),   # Wait For Context Or Duration (context)

    # ═══════════════════════════════════════════════
    # STAGE 2: PRACTITIONER (Mid)
    # ═══════════════════════════════════════════════

    # Category: standard-library
    ("044", "021",   "practitioner", "standard-library"),     # Normalize Tags
    ("045", "023",   "practitioner", "standard-library"),     # Time Business Hours
    ("046", "024",   "practitioner", "standard-library"),     # Build Search URL
    ("047", "025",   "practitioner", "standard-library"),     # Pretty JSON
    ("048", "027",   "practitioner", "standard-library"),     # Safe Join
    ("049", "046",   "practitioner", "standard-library"),     # Run-Length Encoding
    ("050", "047",   "practitioner", "standard-library"),     # Run-Length Decoding
    ("051", "048",   "practitioner", "standard-library"),     # Caesar Cipher
    ("052", "049",   "practitioner", "standard-library"),     # ISBN-10 Validator
    ("053", "050",   "practitioner", "standard-library"),     # Roman Numerals
    ("054", "054",   "practitioner", "standard-library"),     # Time Window Check
    ("055", "057",   "practitioner", "standard-library"),     # Bracket Matcher
    ("056", "058",   "practitioner", "standard-library"),     # Markdown Heading Extractor
    ("057", "059",   "practitioner", "standard-library"),     # Longest Common Prefix
    ("058", "060",   "practitioner", "standard-library"),     # Kebab/Snake → Camel
    ("059", "061",   "practitioner", "standard-library"),     # File Extension Counter
    ("060", "064",   "practitioner", "standard-library"),     # Temperature Converter
    ("061", "065",   "practitioner", "standard-library"),     # Log Line Parser
    ("062", "069",   "practitioner", "standard-library"),     # UUID v4 Generator
    ("063", "070",   "practitioner", "standard-library"),     # Password Strength Scorer
    ("064", "083",   "practitioner", "standard-library"),     # Line-Oriented Reader
    ("065", "136",   "practitioner", "standard-library"),     # Count Lines
    ("066", "137",   "practitioner", "standard-library"),     # Copy With Limit
    ("067", "139",   "practitioner", "standard-library"),     # Read Exactly N
    ("068", "117",   "practitioner", "standard-library"),     # Tar/Gzip Archiver

    # Category: web-networking
    ("069", "053",   "practitioner", "web-networking"),        # HTTP Status Classifier
    ("070", "071",   "practitioner", "web-networking"),        # HTTP Query Builder
    ("071", "072",   "practitioner", "web-networking"),        # Context Timeout Wrapper
    ("072", "073",   "practitioner", "web-networking"),        # Retry with Backoff
    ("073", "081",   "practitioner", "web-networking"),        # HTTP Client with Retries
    ("074", "082",   "practitioner", "web-networking"),        # TCP Echo Server
    ("075", "085",   "practitioner", "web-networking"),        # Mini HTTP Router
    ("076", "086",   "practitioner", "web-networking"),        # Middleware Chain
    ("077", "091",   "practitioner", "web-networking"),        # HTTP Server Graceful Shutdown
    ("078", None,    "practitioner", "web-networking"),        # NEW: net/http Server Basics

    # Category: data-serialization
    ("079", "051",   "practitioner", "data-serialization"),    # Parse CSV Line
    ("080", "052",   "practitioner", "data-serialization"),    # JSON Pretty Print
    ("081", "063",   "practitioner", "data-serialization"),    # CSV to JSON
    ("082", "067",   "practitioner", "data-serialization"),    # INI Parser
    ("083", "079",   "practitioner", "data-serialization"),    # JSON Patch (subset)
    ("084", "080",   "practitioner", "data-serialization"),    # Config Loader
    ("085", "084",   "practitioner", "data-serialization"),    # JSON Lines Filter
    ("086", "062",   "practitioner", "data-serialization"),    # Top N Words
    ("087", "138",   "practitioner", "data-serialization"),    # Read CSV Records
    ("088", None,    "practitioner", "data-serialization"),    # NEW: YAML/TOML Parsing

    # Category: real-world-packages
    ("089", None,    "practitioner", "real-world-packages"),   # NEW: Structured Logging (slog)
    ("090", None,    "practitioner", "real-world-packages"),   # NEW: Cobra CLI
    ("091", None,    "practitioner", "real-world-packages"),   # NEW: testify Assertions
    ("092", None,    "practitioner", "real-world-packages"),   # NEW: pgx Database Basics
    ("093", None,    "practitioner", "real-world-packages"),   # NEW: Fuzz Testing

    # Category: concurrency
    ("094", "055",   "practitioner", "concurrency"),           # LRU Cache (single-thread)
    ("095", "056",   "practitioner", "concurrency"),           # Stack
    ("096", "066",   "practitioner", "concurrency"),           # Rate Limiter (token bucket)
    ("097", "074",   "practitioner", "concurrency"),           # Concurrent MapReduce
    ("098", "075",   "practitioner", "concurrency"),           # Worker Pool
    ("099", "076",   "practitioner", "concurrency"),           # Safe Counter
    ("100", "077",   "practitioner", "concurrency"),           # Debounce
    ("101", "078",   "practitioner", "concurrency"),           # Throttle
    ("102", "087",   "practitioner", "concurrency"),           # In-Memory KV Store
    ("103", "088",   "practitioner", "concurrency"),           # Simple Pub/Sub
    ("104", "089",   "practitioner", "concurrency"),           # Fixed-Window Metrics
    ("105", "090",   "practitioner", "concurrency"),           # Trie Autocomplete
    ("106", "095",   "practitioner", "concurrency"),           # Concurrent File Downloader
    ("107", "096",   "practitioner", "concurrency"),           # Semaphore
    ("108", "097",   "practitioner", "concurrency"),           # Barrier
    ("109", "098",   "practitioner", "concurrency"),           # Fan-in/Fan-out Pipeline
    ("110", "099",   "practitioner", "concurrency"),           # Batcher
    ("111", "100",   "practitioner", "concurrency"),           # LRU Cache (thread-safe)
    ("112", "101",   "practitioner", "concurrency"),           # Token Bucket (thread-safe)
    ("113", None,    "practitioner", "concurrency"),           # NEW: Race Detection

    # ═══════════════════════════════════════════════
    # STAGE 3: SENIOR
    # ═══════════════════════════════════════════════

    # Category: architecture-design
    ("114", "093",   "senior", "architecture-design"),         # Structured Errors
    ("115", "094",   "senior", "architecture-design"),         # Circuit Breaker
    ("116", "109",   "senior", "architecture-design"),         # JSON Schema Validator
    ("117", "110",   "senior", "architecture-design"),         # Rate-limited HTTP Scraper
    ("118", "111",   "senior", "architecture-design"),         # Websocket Chat
    ("119", "112",   "senior", "architecture-design"),         # SSE Stream
    ("120", "127",   "senior", "architecture-design"),         # Generics: Set/Map Utilities
    ("121", "128",   "senior", "architecture-design"),         # Generics: Optional/Result

    # Category: observability
    ("122", "092",   "senior", "observability"),               # Context-aware Logger
    ("123", None,    "senior", "observability"),               # NEW: OpenTelemetry Basics

    # Category: performance
    ("124", "126",   "senior", "performance"),                 # Benchmarking Kata
    ("125", None,    "senior", "performance"),                 # NEW: pprof Profiling
    ("126", None,    "senior", "performance"),                 # NEW: Benchmarking Deep Dive

    # Category: security
    ("127", "113",   "senior", "security"),                    # JWT Sign/Verify
    ("128", "114",   "senior", "security"),                    # HMAC Request Signing
    ("129", "115",   "senior", "security"),                    # Password Hashing
    ("130", "116",   "senior", "security"),                    # File Integrity Checker
    ("131", None,    "senior", "security"),                    # NEW: Security Hardening

    # Category: advanced-data
    ("132", "102",   "senior", "advanced-data"),               # Bloom Filter
    ("133", "103",   "senior", "advanced-data"),               # Merkle Tree
    ("134", "104",   "senior", "advanced-data"),               # Binary Heap Priority Queue
    ("135", "105",   "senior", "advanced-data"),               # Dijkstra Shortest Path
    ("136", "106",   "senior", "advanced-data"),               # A* Grid Pathfinding
    ("137", "108",   "senior", "advanced-data"),               # SQLite-backed Repo
    ("138", "140",   "senior", "advanced-data"),               # Run In Transaction

    # ═══════════════════════════════════════════════
    # STAGE 4: LEAD
    # ═══════════════════════════════════════════════

    # Category: code-quality-review
    ("139", "118",   "lead", "code-quality-review"),           # Diff (line-based)
    ("140", "119",   "lead", "code-quality-review"),           # Mini grep
    ("141", "120",   "lead", "code-quality-review"),           # Mini wc
    ("142", "121",   "lead", "code-quality-review"),           # Mini cut
    ("143", "122",   "lead", "code-quality-review"),           # Mini head/tail
    ("144", "123",   "lead", "code-quality-review"),           # Mini sort
    ("145", "124",   "lead", "code-quality-review"),           # Mini uniq
    ("146", "125",   "lead", "code-quality-review"),           # Concurrent Test Harness
    ("147", None,    "lead", "code-quality-review"),           # NEW: Code Review Kata

    # Category: build-deploy
    ("148", None,    "lead", "build-deploy"),                  # NEW: Docker Multi-Stage Build
    ("149", None,    "lead", "build-deploy"),                  # NEW: CI/CD Pipeline

    # Category: advanced-patterns
    ("150", "107",   "lead", "advanced-patterns"),             # Event Sourcing Mini
    ("151", "129",   "lead", "advanced-patterns"),             # Reflection: Struct Tag Parser
    ("152", "130",   "lead", "advanced-patterns"),             # Plugin-free DI Container

    # Category: leadership-communication
    ("153", None,    "lead", "leadership-communication"),      # NEW: Architecture Decision Records
    ("154", None,    "lead", "leadership-communication"),      # NEW: Mentoring Scenario

    # Category: bug-fix-lab
    ("155", "131",   "lead", "bug-fix-lab"),                   # Normalize Username Bug
    ("156", "132",   "lead", "bug-fix-lab"),                   # Sum Positive Bug
    ("157", "133",   "lead", "bug-fix-lab"),                   # First Non-Empty Bug
    ("158", "134",   "lead", "bug-fix-lab"),                   # Parse Flag Bug
    ("159", "135",   "lead", "bug-fix-lab"),                   # Clamp Percentage Bug
]


def build_track_json():
    """Generate the new track.json with 4 stages."""
    stages = {
        "foundation": {
            "id": "foundation",
            "title": "Foundation",
            "level": "junior",
            "description": "Write, test, and run Go programs. Build confidence with language fundamentals.",
            "categories": {}
        },
        "practitioner": {
            "id": "practitioner",
            "title": "Practitioner",
            "level": "mid",
            "description": "Build real systems with Go. Use standard library and popular packages confidently.",
            "categories": {}
        },
        "senior": {
            "id": "senior",
            "title": "Senior Developer",
            "level": "senior",
            "description": "Make architectural decisions. Build reliable, observable, performant systems.",
            "categories": {}
        },
        "lead": {
            "id": "lead",
            "title": "Lead Developer",
            "level": "lead",
            "description": "Lead teams, make technical decisions, mentor others, and ship with confidence.",
            "categories": {}
        }
    }

    category_meta = {
        "setup-and-toolchain": {
            "title": "Go Setup & Toolchain",
            "description": "Project setup, module system, toolchain basics, and development workflow.",
            "learning_goal": "Set up a Go project, use go fmt/vet/lint, and understand the module system."
        },
        "language-basics": {
            "title": "Language Basics",
            "description": "Variables, types, functions, conditionals, loops, and fundamental patterns.",
            "learning_goal": "Write clear, correct Go functions that handle edge cases and use the language idiomatically."
        },
        "data-structures": {
            "title": "Data Structures",
            "description": "Slices, maps, structs, pointers, and algorithmic thinking with collections.",
            "learning_goal": "Use Go data structures and iteration patterns to solve problems correctly."
        },
        "error-handling": {
            "title": "Error Handling",
            "description": "Return values, sentinel errors, error wrapping, custom types, and interface patterns.",
            "learning_goal": "Handle errors explicitly and design clean error boundaries."
        },
        "testing-fundamentals": {
            "title": "Testing Fundamentals",
            "description": "Table-driven tests, subtests, sync primitives, context, and test discipline.",
            "learning_goal": "Write tests that prove behavior and catch regressions."
        },
        "standard-library": {
            "title": "Standard Library Deep Dive",
            "description": "Core packages: strings, io, bytes, sort, filepath, encoding, crypto, and more.",
            "learning_goal": "Use common standard library packages with confidence and precision."
        },
        "web-networking": {
            "title": "Web & Networking",
            "description": "HTTP servers, clients, middleware, routing, context propagation, and network protocols.",
            "learning_goal": "Build robust web services that handle failures, timeouts, and partial states gracefully."
        },
        "data-serialization": {
            "title": "Data & Serialization",
            "description": "JSON, CSV, YAML, TOML, INI, config loading, and stream processing.",
            "learning_goal": "Parse and generate structured data formats with proper error handling."
        },
        "real-world-packages": {
            "title": "Real-World Packages",
            "description": "Production libraries: structured logging (slog), CLI frameworks (cobra), testing (testify), databases (pgx).",
            "learning_goal": "Use popular Go packages that appear in production codebases daily."
        },
        "concurrency": {
            "title": "Concurrency Patterns",
            "description": "Goroutines, channels, sync primitives, pipelines, fan-out/fan-in, and concurrent data structures.",
            "learning_goal": "Build concurrent systems that stay correct under load and timing variance."
        },
        "architecture-design": {
            "title": "Architecture & Design",
            "description": "Structured errors, circuit breakers, generics, websockets, SSE, and design patterns.",
            "learning_goal": "Design systems with clear boundaries, resilience patterns, and clean abstractions."
        },
        "observability": {
            "title": "Observability",
            "description": "Structured logging, metrics, tracing, and understanding runtime behavior.",
            "learning_goal": "Make systems observable so you can diagnose issues in production."
        },
        "performance": {
            "title": "Performance",
            "description": "Benchmarking, profiling with pprof, memory analysis, and optimization.",
            "learning_goal": "Measure before optimizing and use profiling to find real bottlenecks."
        },
        "security": {
            "title": "Security",
            "description": "JWT, HMAC, password hashing, file integrity, input validation, and secure coding.",
            "learning_goal": "Apply security best practices to protect data and systems."
        },
        "advanced-data": {
            "title": "Advanced Data Structures",
            "description": "Bloom filters, Merkle trees, heaps, graph algorithms, databases, and transactions.",
            "learning_goal": "Implement and use advanced data structures for real-world problems."
        },
        "code-quality-review": {
            "title": "Code Quality & Review",
            "description": "CLI tools, code review patterns, refactoring, and quality standards.",
            "learning_goal": "Write code that is maintainable, reviewable, and production-ready."
        },
        "build-deploy": {
            "title": "Build & Deploy",
            "description": "Docker, CI/CD, cross-compilation, release management, and deployment.",
            "learning_goal": "Ship Go applications reliably with automated build and deployment pipelines."
        },
        "advanced-patterns": {
            "title": "Advanced Patterns",
            "description": "Event sourcing, reflection, dependency injection, and architectural patterns.",
            "learning_goal": "Apply advanced design patterns to complex systems."
        },
        "leadership-communication": {
            "title": "Leadership & Communication",
            "description": "Architecture decision records, technical writing, mentoring, and team practices.",
            "learning_goal": "Lead technical decisions and mentor other developers effectively."
        },
        "bug-fix-lab": {
            "title": "Bug Fix Lab",
            "description": "Production debugging: reproduce, isolate, patch minimally, and confirm with regression tests.",
            "learning_goal": "Develop debugging habits that save hours in production."
        }
    }

    # Build the categories for each stage
    for new_id, old_id, stage, cat_id in MAPPING:
        s = stages[stage]
        if cat_id not in s["categories"]:
            meta = category_meta[cat_id]
            s["categories"][cat_id] = {
                "id": cat_id,
                "title": meta["title"],
                "description": meta["description"],
                "learning_goal": meta["learning_goal"],
                "kata_ids": []
            }
        s["categories"][cat_id]["kata_ids"].append(new_id)

    # Assemble the final track structure
    track = {
        "id": "go-core-100",
        "title": "Go Mastery: Junior to Lead",
        "description": "Structured Go learning from first principles to lead-level architecture, design, and team leadership.",
        "stages": []
    }

    for stage_key in ["foundation", "practitioner", "senior", "lead"]:
        s = stages[stage_key]
        stage_obj = {
            "id": s["id"],
            "title": s["title"],
            "level": s["level"],
            "description": s["description"],
            "categories": list(s["categories"].values())
        }
        track["stages"].append(stage_obj)

    return track


def rename_kata_dirs():
    """Rename kata directories according to the mapping."""
    # Build reverse mapping: old_id -> new_id
    rename_map = {}
    for new_id, old_id, stage, cat_id in MAPPING:
        if old_id is not None and old_id != new_id:
            rename_map[old_id] = new_id

    # Sort by old_id descending to avoid conflicts during rename
    for old_id, new_id in sorted(rename_map.items(), key=lambda x: -int(x[0])):
        old_dir = None
        for d in KATAS_DIR.iterdir():
            if d.is_dir() and d.name.startswith(f"kata-{old_id}-"):
                old_dir = d
                break
        if old_dir is None:
            print(f"WARNING: No directory found for old_id {old_id}", file=sys.stderr)
            continue

        slug = old_dir.name[len(f"kata-{old_id}-"):]
        new_dir = KATAS_DIR / f"kata-{new_id}-{slug}"

        if new_dir.exists():
            print(f"WARNING: {new_dir} already exists, skipping", file=sys.stderr)
            continue

        print(f"  {old_dir.name} -> {new_dir.name}")
        old_dir.rename(new_dir)


def add_metadata_to_katas():
    """Add stage, category, tags, prerequisites, flashcards, quiz to each kata.json."""
    # Build lookup: new_id -> (stage, category)
    meta_lookup = {}
    for new_id, old_id, stage, cat_id in MAPPING:
        meta_lookup[new_id] = (stage, cat_id)

    # Read existing katas and enrich them
    for new_id, (stage, cat_id) in meta_lookup.items():
        # Find the directory
        kata_dir = None
        for d in KATAS_DIR.iterdir():
            if d.is_dir() and d.name.startswith(f"kata-{new_id}-"):
                kata_dir = d
                break
        if kata_dir is None:
            print(f"WARNING: No directory for new_id {new_id}", file=sys.stderr)
            continue

        json_path = kata_dir / "kata.json"
        if json_path.exists():
            with open(json_path) as f:
                meta = json.load(f)
        else:
            meta = {"id": new_id, "slug": kata_dir.name[len(f"kata-{new_id}-"):], "title": kata_dir.name}

        # Add new fields
        meta["stage"] = stage
        meta["category"] = cat_id
        meta["level"] = {"foundation": "junior", "practitioner": "mid", "senior": "senior", "lead": "lead"}[stage]

        # Auto-generate tags from focus
        focus = meta.get("focus", "")
        tags = []
        tag_keywords = {
            "testing": ["test", "testing", "benchmark", "fuzz"],
            "concurrency": ["goroutine", "channel", "mutex", "waitgroup", "sync", "concurrent", "atomic"],
            "error-handling": ["error", "errors", "panic", "recover"],
            "http": ["http", "net", "tcp", "websocket", "sse"],
            "data-structures": ["slice", "map", "struct", "stack", "queue", "heap", "trie", "bloom", "merkle"],
            "algorithms": ["search", "sort", "path", "dijkstra", "a*", "binary"],
            "io": ["file", "io", "reader", "writer", "stream", "csv", "json", "yaml", "tar", "gzip"],
            "crypto": ["crypto", "jwt", "hmac", "hash", "bcrypt", "argon"],
            "string": ["string", "rune", "text", "markdown", "cipher"],
            "time": ["time", "timer", "rate", "throttle", "debounce"],
            "database": ["database", "sql", "sqlite", "transaction", "repository"],
            "cli": ["cli", "grep", "wc", "cut", "sort", "uniq", "head", "tail"],
            "architecture": ["design", "pattern", "architecture", "di", "container", "event", "sourcing"],
            "generics": ["generic"],
            "reflection": ["reflect"],
        }
        focus_lower = focus.lower()
        for tag, keywords in tag_keywords.items():
            if any(kw in focus_lower for kw in keywords):
                tags.append(tag)
        meta["tags"] = tags

        # Add prerequisites (first 5 katas have none)
        new_id_int = int(new_id)
        prereqs = []
        if new_id_int > 5:
            # Simple heuristic: previous katas in same category
            for prev_id, prev_old, prev_stage, prev_cat in MAPPING:
                if prev_stage == stage and prev_cat == cat_id and int(prev_id) < new_id_int:
                    prereqs.append(prev_id)
                    if len(prereqs) >= 2:
                        break
        meta["prerequisites"] = prereqs

        # Estimated minutes (rough heuristic based on stage)
        stage_minutes = {"foundation": 15, "practitioner": 25, "senior": 40, "lead": 30}
        meta["estimated_minutes"] = stage_minutes.get(stage, 20)

        # Write back
        with open(json_path, "w") as f:
            json.dump(meta, f, indent=2)
            f.write("\n")


def main():
    print("=== GoKatas Curriculum Migration ===\n")

    print("1. Generating new track.json...")
    track = build_track_json()
    track_path = TRACK_DIR / "track.json"
    with open(track_path, "w") as f:
        json.dump(track, f, indent=2)
        f.write("\n")
    print(f"   Wrote {track_path}")

    # Count katas per stage
    for stage in track["stages"]:
        total = sum(len(cat["kata_ids"]) for cat in stage["categories"])
        print(f"   {stage['title']} ({stage['level']}): {total} katas, {len(stage['categories'])} categories")

    print(f"\n2. Renaming kata directories...")
    rename_kata_dirs()

    print(f"\n3. Adding metadata to kata.json files...")
    add_metadata_to_katas()

    print("\n=== Migration complete ===")
    print(f"Total katas: {len(MAPPING)}")
    new_katas = sum(1 for _, old, _, _ in MAPPING if old is None)
    print(f"New katas to create: {new_katas}")


if __name__ == "__main__":
    main()
