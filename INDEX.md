# Kata Index

This index follows the recommended learning order from first programming concepts to senior-level Go design and debugging.

## Starter: Programming Essentials (`001-010`)

- **001** — Build Greeting _(focus: Functions, variables, conditionals, strings.TrimSpace)_
- **002** — Parse Whole Number _(focus: Input parsing, strconv.Atoi, boolean success flags)_
- **003** — Basic Calculator _(focus: Branching, arithmetic operators, error signaling with bool)_
- **004** — Even Odd Label _(focus: Conditionals, modulo arithmetic)_
- **005** — Clamp Value _(focus: Comparisons, boundary logic, defensive argument handling)_
- **006** — Sum First N _(focus: For-loops, accumulation patterns)_
- **007** — Countdown Slice _(focus: Loops, slice construction, edge-case handling)_
- **008** — Repeat Text _(focus: String construction, loops, guard clauses)_
- **009** — Character Frequency _(focus: Maps, rune iteration, Unicode-safe counting)_
- **010** — Average (Integer) _(focus: Slices, integer arithmetic, validation)_

## Starter: Data and Logic (`011-020`)

- **011** — Swap First and Last _(focus: Slice copying, mutation safety, indexing)_
- **012** — Remove Empty Strings _(focus: Filtering, loops, preserving order)_
- **013** — Grade Classifier _(focus: Condition ranges, input validation)_
- **014** — Count True Values _(focus: Boolean logic, loop counting)_
- **015** — Find First Index _(focus: Linear search, return conventions)_
- **016** — Safe Slice Range _(focus: Indices, bounds clamping, non-panicking behavior)_
- **017** — Budget Status _(focus: Arithmetic comparison, decision outcomes)_
- **018** — Merge Alternating Slices _(focus: Two-pointer merge, slice appends)_
- **019** — Word Lengths _(focus: strings.Fields, maps, normalization decisions)_
- **020** — Incomplete Task Filter _(focus: Structs, slice filtering, business-rule selection)_

## Well-Known Packages (`021-030`)

- **021** — Normalize Tags _(focus: strings.Split, strings.TrimSpace, strings.ToLower)_
- **022** — Parse Positive Int _(focus: strconv.Atoi, strings.TrimSpace, input validation)_
- **023** — Is Within Business Hours _(focus: time.Time, time.Location, weekday/hour checks)_
- **024** — Build Search URL _(focus: net/url parsing, query encoding)_
- **025** — Pretty JSON _(focus: encoding/json validation and formatting)_
- **026** — Read Env Or Default _(focus: os.LookupEnv, strings.TrimSpace)_
- **027** — Safe Join _(focus: filepath.Join, filepath.Clean, traversal guard checks)_
- **028** — Wait For Context Or Duration _(focus: context cancellation, time.Timer selection)_
- **029** — Increment Concurrently _(focus: sync.WaitGroup, sync.Mutex, goroutine coordination)_
- **030** — Find User Name _(focus: errors.Is, sentinel errors, strings.TrimSpace)_

## Go Foundations (`031-040`)

- **031** — FizzBuzz _(focus: Basics: loops, conditionals, slices, strconv)_
- **032** — Sum of Integers _(focus: Basics: loops, function, edge cases)_
- **033** — Reverse String _(focus: Strings, runes vs bytes)_
- **034** — Palindrome Check _(focus: Strings, normalization basics)_
- **035** — Count Vowels _(focus: Loops, unicode, maps)_
- **036** — Factorial _(focus: Recursion vs loop, overflow checks)_
- **037** — Fibonacci _(focus: Iteration, slices)_
- **038** — Filter Even _(focus: Slices, order, immutability)_
- **039** — Unique Strings _(focus: Maps as sets, order preservation)_
- **040** — Word Count _(focus: Strings, fields, maps)_

## Core Collections and Algorithms (`041-050`)

- **041** — Anagram Check _(focus: Rune counts, normalization)_
- **042** — Min/Max _(focus: Errors, edge cases)_
- **043** — Rotate Slice _(focus: Generics, indexing)_
- **044** — Binary Search _(focus: Algorithms, generics)_
- **045** — Merge Two Sorted Lists _(focus: Two pointers)_
- **046** — Run-Length Encoding _(focus: Strings.Builder, runes)_
- **047** — Run-Length Decoding _(focus: Parsing, errors)_
- **048** — Caesar Cipher _(focus: ASCII letters shifting)_
- **049** — ISBN-10 Validator _(focus: Checksums)_
- **050** — Roman Numerals _(focus: Greedy mapping)_

## Text, Data, and Transformations (`051-062`)

- **051** — Parse CSV Line _(focus: Parsing quoted CSV)_
- **052** — JSON Pretty Print _(focus: encoding/json)_
- **053** — HTTP Status Classifier _(focus: Switch ranges)_
- **054** — Time Window Check _(focus: time package)_
- **055** — LRU Cache (single-thread) _(focus: Structs, maps, list)_
- **056** — Stack _(focus: Data structures, generics)_
- **057** — Bracket Matcher _(focus: Stacks, runes)_
- **058** — Markdown Heading Extractor _(focus: Line scanning)_
- **059** — Longest Common Prefix _(focus: Strings)_
- **060** — Kebab/Snake → Camel _(focus: String transforms)_
- **061** — File Extension Counter _(focus: os, filepath)_
- **062** — Top N Words _(focus: Sorting, ties)_

## File, Time, and Config (`063-070`)

- **063** — CSV to JSON _(focus: I/O, encoding)_
- **064** — Temperature Converter _(focus: Parsing floats)_
- **065** — Log Line Parser _(focus: Key=value parsing)_
- **066** — Rate Limiter (token bucket) _(focus: Time, structs)_
- **067** — INI Parser _(focus: Sections, maps)_
- **068** — Mini Template Renderer _(focus: Placeholders)_
- **069** — UUID v4 Generator _(focus: crypto/rand)_
- **070** — Password Strength Scorer _(focus: Rules)_

## Services and Boundaries (`071-090`)

- **071** — HTTP Query Builder _(focus: net/url)_
- **072** — Context Timeout Wrapper _(focus: context)_
- **073** — Retry with Backoff _(focus: time)_
- **074** — Concurrent MapReduce _(focus: Goroutines, channels)_
- **075** — Worker Pool _(focus: Concurrency patterns)_
- **076** — Safe Counter _(focus: sync/atomic or mutex)_
- **077** — Debounce _(focus: Time, channels)_
- **078** — Throttle _(focus: Time, channels)_
- **079** — JSON Patch (subset) _(focus: encoding/json)_
- **080** — Config Loader _(focus: env + file)_
- **081** — HTTP Client with Retries _(focus: net/http)_
- **082** — TCP Echo Server _(focus: net)_
- **083** — Line-Oriented Reader _(focus: bufio)_
- **084** — JSON Lines Filter _(focus: Stream processing)_
- **085** — Mini HTTP Router _(focus: net/http)_
- **086** — Middleware Chain _(focus: net/http)_
- **087** — In-Memory KV Store _(focus: Maps, RWMutex)_
- **088** — Simple Pub/Sub _(focus: Fanout)_
- **089** — Fixed-Window Metrics _(focus: Time buckets)_
- **090** — Trie Autocomplete _(focus: Data structures)_

## Concurrency and Reliability (`091-110`)

- **091** — HTTP Server with Graceful Shutdown _(focus: net/http, context)_
- **092** — Context-aware Logger _(focus: Advanced practice)_
- **093** — Structured Errors _(focus: Advanced practice)_
- **094** — Circuit Breaker _(focus: State machines, time)_
- **095** — Concurrent File Downloader _(focus: Concurrency, net/http)_
- **096** — Semaphore _(focus: Concurrency primitives)_
- **097** — Barrier _(focus: Concurrency coordination)_
- **098** — Fan-in/Fan-out Pipeline _(focus: Channels, composition)_
- **099** — Batcher _(focus: Time, channels)_
- **100** — LRU Cache (thread-safe) _(focus: Advanced practice)_
- **101** — Token Bucket (thread-safe) _(focus: Advanced practice)_
- **102** — Bloom Filter _(focus: Advanced practice)_
- **103** — Merkle Tree _(focus: Advanced practice)_
- **104** — Binary Heap Priority Queue _(focus: Advanced practice)_
- **105** — Dijkstra Shortest Path _(focus: Advanced practice)_
- **106** — A* Grid Pathfinding _(focus: Advanced practice)_
- **107** — Event Sourcing Mini _(focus: Advanced practice)_
- **108** — SQLite-backed Repo (pure Go driver) _(focus: Advanced practice)_
- **109** — JSON Schema Validator (subset) _(focus: Advanced practice)_
- **110** — Rate-limited HTTP Scraper _(focus: Advanced practice)_

## Tooling, Security, and Advanced Design (`111-130`)

- **111** — Websocket Chat (basic) _(focus: Advanced practice)_
- **112** — SSE Stream _(focus: Advanced practice)_
- **113** — JWT Sign/Verify _(focus: crypto, encoding)_
- **114** — HMAC Request Signing _(focus: crypto/hmac)_
- **115** — Password Hashing (bcrypt/argon2) _(focus: crypto, security)_
- **116** — File Integrity Checker _(focus: io, crypto/sha256)_
- **117** — Tar/Gzip Archiver _(focus: archive/tar, compress/gzip)_
- **118** — Diff (line-based) _(focus: Algorithms)_
- **119** — Mini grep _(focus: CLI, IO)_
- **120** — Mini wc _(focus: CLI, IO)_
- **121** — Mini cut _(focus: CLI, IO)_
- **122** — Mini head/tail _(focus: CLI, IO)_
- **123** — Mini sort _(focus: CLI, IO)_
- **124** — Mini uniq _(focus: CLI, IO)_
- **125** — Concurrent Test Harness _(focus: Advanced practice)_
- **126** — Benchmarking Kata _(focus: testing/benchmark)_
- **127** — Generics: Set/Map Utilities _(focus: Advanced practice)_
- **128** — Generics: Optional/Result _(focus: Advanced practice)_
- **129** — Reflection: Struct Tag Parser _(focus: reflect)_
- **130** — Plugin-free DI Container _(focus: Design)_

## Bug Fix Lab (`131-135`)

- **131** — Normalize Username Bug _(focus: strings.TrimSpace, strings.ToLower, strings.ReplaceAll)_
- **132** — Sum Positive Bug _(focus: loops, integer filtering)_
- **133** — First Non-Empty Bug _(focus: strings.TrimSpace, loop selection)_
- **134** — Parse Flag Bug _(focus: strings.ToLower, explicit parsing, error handling)_
- **135** — Clamp Percentage Bug _(focus: boundary logic)_

## Databases and I/O (`136-140`)

- **136** — Count Lines _(focus: bufio.Scanner, io.Reader)_
- **137** — Copy With Limit _(focus: io.CopyN, boundary validation)_
- **138** — Read CSV Records _(focus: encoding/csv reader loops)_
- **139** — Read Exactly N _(focus: io.ReadFull semantics)_
- **140** — Run In Transaction _(focus: transaction control flow)_

