# Kata Index

Organized by learning track and stage. Each kata includes its focus area and estimated time.

## Go Mastery: Junior to Lead

### Foundation (Junior)

**Go Setup & Toolchain** (13 katas)

- **000** — Go Setup & First Program _(focus: go mod init, main package, go run, go build)_
- **001** — Build Greeting _(focus: Functions, variables, conditionals, strings.TrimSpace)_
- **002** — Parse Whole Number _(focus: Input parsing, strconv.Atoi, boolean success flags)_
- **003** — Basic Calculator _(focus: Branching, arithmetic operators, error signaling with bool)_
- **004** — Even Odd Label _(focus: Conditionals, modulo arithmetic)_
- **005** — Go Vet & Lint _(focus: go vet, golangci-lint, static analysis, code quality)_
- **163** — Variable Scope & Lifetime _(focus: local vs package vs global scope, block scoping, short vs long declaration, ~25 min, 4 flashcards, 3 quiz questions)_
- **164** — Go Application Structure _(focus: cmd/, internal/, pkg/, directory layout, package naming conventions, ~30 min, 4 flashcards, 3 quiz questions)_
- **165** — Constants & iota _(focus: const declarations, iota enum pattern, typed vs untyped constants, ~20 min, 4 flashcards, 3 quiz questions)_
- **166** — init() Function & Package Lifecycle _(focus: init() execution order, package-level initialization, side effects, ~25 min, 4 flashcards, 3 quiz questions)_
- **167** — Exported vs Unexported Identifiers _(focus: uppercase/lowercase naming, visibility rules, API design, ~20 min, 4 flashcards, 3 quiz questions)_
- **168** — Makefile Basics _(focus: Make targets, variables, dependencies, phony targets, build/test/clean patterns, ~30 min, 4 flashcards, 3 quiz questions)_
- **169** — Docker Development Environment _(focus: Dockerfile, docker-compose, dev containers, volume mounts, healthchecks, ~35 min, 4 flashcards, 3 quiz questions)_

**Language Basics** (19 katas)

- **006** — FizzBuzz _(focus: Basics: loops, conditionals, slices, strconv)_
- **007** — Countdown Slice _(focus: Loops, slice construction, edge-case handling)_
- **008** — Reverse String _(focus: Strings, runes vs bytes)_
- **009** — Min/Max _(focus: Errors, edge cases)_
- **010** — Average (Integer) _(focus: Slices, integer arithmetic, validation)_
- **011** — Clamp Value _(focus: Comparisons, boundary logic, defensive argument handling)_
- **012** — Sum First N _(focus: For-loops, accumulation patterns)_
- **013** — Repeat Text _(focus: String construction, loops, guard clauses)_
- **014** — Grade Classifier _(focus: Condition ranges, input validation)_
- **015** — Count True Values _(focus: Boolean logic, loop counting)_
- **016** — Budget Status _(focus: Arithmetic comparison, decision outcomes)_
- **017** — Parse Positive Int _(focus: strconv.Atoi, strings.TrimSpace, input validation)_
- **018** — Factorial _(focus: Recursion vs loop, overflow checks)_
- **019** — Word Lengths _(focus: strings.Fields, maps, normalization decisions)_
- **170** — Error Handling Idioms _(focus: if err != nil, sentinel errors, error wrapping, custom types, Is/As, ~35 min, 4 flashcards, 3 quiz questions)_
- **171** — Naming & Stuttering _(focus: No stuttering, clear names, receiver consistency, package naming, ~25 min, 4 flashcards, 3 quiz questions)_
- **172** — Composition Over Inheritance _(focus: Struct embedding, small interfaces, accept interfaces return structs, ~30 min, 4 flashcards, 3 quiz questions)_
- **173** — Defer Patterns & Cleanup _(focus: defer, resource cleanup, panic recovery, named returns, ~25 min, 4 flashcards, 3 quiz questions)_
- **174** — Nil vs Empty Slices & Maps _(focus: nil vs empty, initialization patterns, when to use which, ~20 min, 4 flashcards, 3 quiz questions)_

**Data Structures** (18 katas)

- **020** — Sum of Integers _(focus: Basics: loops, function, edge cases)_
- **021** — Character Frequency _(focus: Maps, rune iteration, Unicode-safe counting)_
- **022** — Count Vowels _(focus: Loops, unicode, maps)_
- **023** — Swap First and Last _(focus: Slice copying, mutation safety, indexing)_
- **024** — Remove Empty Strings _(focus: Filtering, loops, preserving order)_
- **025** — Find First Index _(focus: Linear search, return conventions)_
- **026** — Safe Slice Range _(focus: Indices, bounds clamping, non-panicking behavior)_
- **027** — Merge Alternating Slices _(focus: Two-pointer merge, slice appends)_
- **028** — Fibonacci _(focus: Iteration, slices)_
- **029** — Incomplete Task Filter _(focus: Structs, slice filtering, business-rule selection)_
- **030** — Find User Name _(focus: errors.Is, sentinel errors, strings.TrimSpace)_
- **031** — Unique Strings _(focus: Maps as sets, order preservation)_
- **032** — Word Count _(focus: Strings, fields, maps)_
- **033** — Anagram Check _(focus: Rune counts, normalization)_
- **034** — Palindrome Check _(focus: Strings, normalization basics)_
- **035** — Rotate Slice _(focus: Generics, indexing)_
- **036** — Binary Search _(focus: Algorithms, generics)_
- **037** — Merge Two Sorted Lists _(focus: Two pointers)_

**Error Handling** (3 katas)

- **038** — Filter Even _(focus: Slices, order, immutability)_
- **039** — Error Wrapping Patterns _(focus: error wrapping, fmt.Errorf, %w, errors.Is, errors.As)_
- **040** — Interface Basics _(focus: interfaces, implicit satisfaction, accept interfaces return structs)_

**Testing Fundamentals** (3 katas)

- **041** — Table-Driven Tests _(focus: table-driven tests, subtests, testing.T, test patterns)_
- **042** — Increment Concurrently _(focus: sync.WaitGroup, sync.Mutex, goroutine coordination)_
- **043** — Wait For Context Or Duration _(focus: context cancellation, time.Timer selection)_

**AI-Augmented Development** (2 katas)

- **160** — AI Prompt Engineering for Developers _(focus: effective prompting, context setting, constraint specification)_
- **161** — AI-Assisted Code Review _(focus: using AI for code review, verifying AI suggestions, critical thinking)_

### Practitioner (Mid)

**Standard Library Deep Dive** (25 katas)

- **044** — Normalize Tags _(focus: strings.Split, strings.TrimSpace, strings.ToLower)_
- **045** — Is Within Business Hours _(focus: time.Time, time.Location, weekday/hour checks)_
- **046** — Build Search URL _(focus: net/url parsing, query encoding)_
- **047** — Pretty JSON _(focus: encoding/json validation and formatting)_
- **048** — Safe Join _(focus: filepath.Join, filepath.Clean, traversal guard checks)_
- **049** — Run-Length Encoding _(focus: Strings.Builder, runes)_
- **050** — Run-Length Decoding _(focus: Parsing, errors)_
- **051** — Caesar Cipher _(focus: ASCII letters shifting)_
- **052** — ISBN-10 Validator _(focus: Checksums)_
- **053** — Roman Numerals _(focus: Greedy mapping)_
- **054** — Time Window Check _(focus: time package)_
- **055** — LRU Cache (single-thread) _(focus: Structs, maps, list)_
- **056** — Stack _(focus: Data structures, generics)_
- **057** — Longest Common Prefix _(focus: Strings)_
- **058** — Kebab/Snake → Camel _(focus: String transforms)_
- **059** — File Extension Counter _(focus: os, filepath)_
- **060** — Line-Oriented Reader _(focus: bufio)_
- **061** — Log Line Parser _(focus: Key=value parsing)_
- **062** — Top N Words _(focus: Sorting, ties)_
- **063** — CSV to JSON _(focus: I/O, encoding)_
- **064** — Temperature Converter _(focus: Parsing floats)_
- **065** — Count Lines _(focus: bufio.Scanner, io.Reader)_
- **066** — Rate Limiter (token bucket) _(focus: Time, structs)_
- **067** — INI Parser _(focus: Sections, maps)_
- **068** — Tar/Gzip Archiver _(focus: archive/tar, compress/gzip)_

**Web & Networking** (10 katas)

- **069** — HTTP Status Classifier _(focus: Switch ranges)_
- **070** — HTTP Query Builder _(focus: net/url)_
- **071** — Context Timeout Wrapper _(focus: context)_
- **072** — HTTP Client with Retries _(focus: net/http)_
- **073** — Retry with Backoff _(focus: time)_
- **074** — Concurrent MapReduce _(focus: Goroutines, channels)_
- **075** — Worker Pool _(focus: Concurrency patterns)_
- **076** — Middleware Chain _(focus: net/http)_
- **077** — HTTP Server with Graceful Shutdown _(focus: net/http, context)_
- **078** — net/http Server Basics _(focus: net/http, http.HandleFunc, http.ListenAndServe, request/response)_

**Data & Serialization** (10 katas)

- **079** — Parse CSV Line _(focus: Parsing quoted CSV)_
- **080** — JSON Pretty Print _(focus: encoding/json)_
- **081** — Password Strength Scorer _(focus: Rules)_
- **082** — Read Exactly N _(focus: io.ReadFull semantics)_
- **083** — JSON Patch (subset) _(focus: encoding/json)_
- **084** — Config Loader _(focus: env + file)_
- **085** — JSON Lines Filter _(focus: Stream processing)_
- **086** — UUID v4 Generator _(focus: crypto/rand)_
- **087** — Run In Transaction _(focus: transaction control flow)_
- **088** — YAML/TOML Parsing _(focus: encoding/json patterns, struct tags, config file parsing)_

**Real-World Packages** (5 katas)

- **089** — Structured Logging (slog) _(focus: log/slog, structured logging, key-value pairs, log levels)_
- **090** — Cobra CLI _(focus: github.com/spf13/cobra, CLI commands, flags, help text)_
- **091** — testify Assertions _(focus: github.com/stretchr/testify, assert/require, test helpers)_
- **092** — pgx Database Basics _(focus: database/sql, pgx, connection, query, scan)_
- **093** — Fuzz Testing _(focus: testing.F, fuzzing, property-based testing, crash detection)_

**Concurrency Patterns** (20 katas)

- **094** — Bracket Matcher _(focus: Stacks, runes)_
- **095** — Markdown Heading Extractor _(focus: Line scanning)_
- **096** — Copy With Limit _(focus: io.CopyN, boundary validation)_
- **097** — TCP Echo Server _(focus: net)_
- **098** — Mini HTTP Router _(focus: net/http)_
- **099** — Safe Counter _(focus: sync/atomic or mutex)_
- **100** — Debounce _(focus: Time, channels)_
- **101** — Throttle _(focus: Time, channels)_
- **102** — In-Memory KV Store _(focus: Maps, RWMutex)_
- **103** — Simple Pub/Sub _(focus: Fanout)_
- **104** — Fixed-Window Metrics _(focus: Time buckets)_
- **105** — Trie Autocomplete _(focus: Data structures)_
- **106** — Concurrent File Downloader _(focus: Concurrency, net/http)_
- **107** — Semaphore _(focus: Concurrency primitives)_
- **108** — Barrier _(focus: Concurrency coordination)_
- **109** — Fan-in/Fan-out Pipeline _(focus: Channels, composition)_
- **110** — Batcher _(focus: Time, channels)_
- **111** — LRU Cache (thread-safe) _(focus: Advanced practice)_
- **112** — Token Bucket (thread-safe) _(focus: Advanced practice)_
- **113** — Race Detection _(focus: go test -race, data races, sync.Mutex, race detector)_

**AI at Scale** (1 katas)

- **162** — AI as a Debugging Partner _(focus: using AI to understand errors, not just fix them, root cause analysis)_

### Senior Developer (Senior)

**Architecture & Design** (8 katas)

- **114** — Structured Errors _(focus: Advanced practice)_
- **115** — Circuit Breaker _(focus: State machines, time)_
- **116** — JSON Schema Validator (subset) _(focus: Advanced practice)_
- **117** — Rate-limited HTTP Scraper _(focus: Advanced practice)_
- **118** — Websocket Chat (basic) _(focus: Advanced practice)_
- **119** — SSE Stream _(focus: Advanced practice)_
- **120** — Mini wc _(focus: CLI, IO)_
- **121** — Mini cut _(focus: CLI, IO)_

**Observability** (2 katas)

- **122** — Context-aware Logger _(focus: Advanced practice)_
- **123** — OpenTelemetry Basics _(focus: tracing, spans, context propagation, observability)_

**Performance** (3 katas)

- **124** — Mini uniq _(focus: CLI, IO)_
- **125** — pprof Profiling _(focus: runtime/pprof, profiling, CPU/memory analysis, benchmarks)_
- **126** — Benchmarking Deep Dive _(focus: benchmark patterns, sub-benchmarks, comparison, regression)_

**Security** (5 katas)

- **127** — JWT Sign/Verify _(focus: crypto, encoding)_
- **128** — HMAC Request Signing _(focus: crypto/hmac)_
- **129** — Password Hashing (bcrypt/argon2) _(focus: crypto, security)_
- **130** — File Integrity Checker _(focus: io, crypto/sha256)_
- **131** — Security Hardening _(focus: input validation, SQL injection, XSS, secrets management)_

**Advanced Data Structures** (7 katas)

- **132** — Bloom Filter _(focus: Advanced practice)_
- **133** — Merkle Tree _(focus: Advanced practice)_
- **134** — Binary Heap Priority Queue _(focus: Advanced practice)_
- **135** — Dijkstra Shortest Path _(focus: Advanced practice)_
- **136** — A* Grid Pathfinding _(focus: Advanced practice)_
- **137** — SQLite-backed Repo (pure Go driver) _(focus: Advanced practice)_
- **138** — Read CSV Records _(focus: encoding/csv reader loops)_

### Lead Developer (Lead)

**Code Quality & Review** (9 katas)

- **139** — Diff (line-based) _(focus: Algorithms)_
- **140** — Mini grep _(focus: CLI, IO)_
- **141** — Generics: Set/Map Utilities _(focus: Advanced practice)_
- **142** — Generics: Optional/Result _(focus: Advanced practice)_
- **143** — Mini head/tail _(focus: CLI, IO)_
- **144** — Mini sort _(focus: CLI, IO)_
- **145** — Benchmarking Kata _(focus: testing/benchmark)_
- **146** — Concurrent Test Harness _(focus: Advanced practice)_
- **147** — Code Review Kata _(focus: code review patterns, refactoring, readability, maintainability)_

**Build & Deploy** (2 katas)

- **148** — Docker Multi-Stage Build _(focus: Dockerfile, multi-stage builds, Go binary, containerization)_
- **149** — CI/CD Pipeline _(focus: GitHub Actions, CI/CD, automated testing, deployment)_

**Advanced Patterns** (3 katas)

- **150** — Event Sourcing Mini _(focus: Advanced practice)_
- **151** — Reflection: Struct Tag Parser _(focus: reflect)_
- **152** — Plugin-free DI Container _(focus: Design)_

**Leadership & Communication** (2 katas)

- **153** — Architecture Decision Records _(focus: ADR, technical decision documentation, RFC process)_
- **154** — Mentoring Scenario _(focus: mentoring, code review feedback, teaching patterns, knowledge transfer)_

**Bug Fix Lab** (5 katas)

- **155** — Normalize Username Bug _(focus: strings.TrimSpace, strings.ToLower, strings.ReplaceAll)_
- **156** — Sum Positive Bug _(focus: loops, integer filtering)_
- **157** — First Non-Empty Bug _(focus: strings.TrimSpace, loop selection)_
- **158** — Parse Flag Bug _(focus: strings.ToLower, explicit parsing, error handling)_
- **159** — Clamp Percentage Bug _(focus: boundary logic)_


## Container Orchestration: Helm

### Helm Foundations (Junior)

**Helm Basics** (2 katas)

- **300** — Helm Chart Basics _(focus: Chart.yaml, values.yaml, templates, helm install, helm upgrade)_
- **301** — Helm Values & Templating _(focus: values.yaml, Go templates, conditional rendering, loops)_


## Security & CVE Awareness

### Security Fundamentals (Junior)

**Security Fundamentals** (1 katas)

- **400** — CVE Analysis & Response _(focus: CVE databases, vulnerability assessment, patching strategy, dependency scanning)_


## Infrastructure as Code: Terraform

### Terraform Foundations (Junior)

**Terraform Basics** (2 katas)

- **200** — Terraform Hello World _(focus: terraform init, plan, apply, destroy, HCL basics)_
- **201** — Terraform Modules & Reusability _(focus: modules, variables, outputs, reusability, composition)_

