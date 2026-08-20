# Phase 2A: Curriculum and Migration Planning Artefacts

Generated: 2026-03-13
Status: Planning only — no kata file changes until Phase 2B approval.

---

## Artefact 1: Final Kata Map (001–140)

### Design Principles Applied

1. **Difficulty reorder**: Katas 021–030 (stdlib packages including context/sync) are harder than 031–040 (FizzBuzz, palindrome). Swap these two blocks so foundations come before packages.
2. **Broken katas replaced**: 18 broken katas (generic `Kata##() error` stubs) are replaced with properly specified katas that fit their category.
3. **Overlaps resolved**: kata-025 and kata-052 both implement `PrettyJSON` — the duplicate (025) is replaced. kata-055 (single-thread LRU) and kata-100 (thread-safe LRU) are kept: they represent deliberate progression from single-thread to concurrent.
4. **Too-academic katas replaced**: 107 (Event Sourcing), 108 (SQLite Repo), 111 (WebSocket Chat), 112 (SSE Stream) replaced with practical Go engineering katas.
5. **Bug-fix katas distributed**: 5 bug-fix katas moved from isolated block (131-135) into levels where they reinforce learned material.
6. **Testing curriculum woven in**: Testing-focused katas added starting at kata 006 and building progressively.
7. **Four-tier levels**: Associate (001–050), Engineer (051–090), Senior (091–120), Lead (121–140).

### Level: Associate (001–050)

Outcome: Implement small-to-medium Go functions with clear contracts, edge-case coverage, and readable tests.

#### Category: Programming Essentials (001–010)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 001 | kata-001-build-greeting | Build Greeting | KEEP | — |
| 002 | kata-002-parse-whole-number | Parse Whole Number | KEEP | — |
| 003 | kata-003-basic-calculator | Basic Calculator | KEEP | — |
| 004 | kata-004-even-odd-label | Even Odd Label | KEEP | — |
| 005 | kata-005-clamp-value | Clamp Value | KEEP | — |
| 006 | kata-006-sum-first-n | Sum First N | KEEP | — |
| 007 | kata-007-countdown-slice | Countdown Slice | KEEP | — |
| 008 | kata-008-repeat-text | Repeat Text | KEEP | — |
| 009 | kata-009-char-frequency | Character Frequency | KEEP | — |
| 010 | kata-010-average-int | Average Integer | KEEP | — |

#### Category: Data and Logic (011–020)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 011 | kata-011-swap-first-last | Swap First and Last | KEEP | — |
| 012 | kata-012-remove-empty-strings | Remove Empty Strings | KEEP | — |
| 013 | kata-013-grade-classifier | Grade Classifier | KEEP | — |
| 014 | kata-014-count-true | Count True Values | KEEP | — |
| 015 | kata-015-find-first-index | Find First Index | KEEP | — |
| 016 | kata-016-safe-slice-range | Safe Slice Range | KEEP | — |
| 017 | kata-017-budget-status | Budget Status | KEEP | — |
| 018 | kata-018-merge-alternating | Merge Alternating Slices | KEEP | — |
| 019 | kata-019-word-lengths | Word Lengths | KEEP | — |
| 020 | kata-020-incomplete-task-filter | Incomplete Task Filter | KEEP | — |

#### Category: Go Foundations (021–030) ← SWAPPED from old 031–040

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 021 | kata-031-fizzbuzz | FizzBuzz | MOVE from 031→021 | — |
| 022 | kata-032-sum-of-integers | Sum of Integers | MOVE from 032→022 | — |
| 023 | kata-033-reverse-string | Reverse String | MOVE from 033→023 | — |
| 024 | kata-034-palindrome-check | Palindrome Check | MOVE from 034→024 | — |
| 025 | kata-035-count-vowels | Count Vowels | MOVE from 035→025 | — |
| 026 | kata-036-factorial | Factorial | MOVE from 036→026 | — |
| 027 | kata-037-fibonacci | Fibonacci | MOVE from 037→027 | — |
| 028 | kata-038-filter-even | Filter Even | MOVE from 038→028 | — |
| 029 | kata-039-unique-strings | Unique Strings | MOVE from 039→029 | — |
| 030 | kata-040-word-count | Word Count | MOVE from 040→030 | — |

#### Category: Well-Known Packages (031–040) ← SWAPPED from old 021–030

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 031 | kata-021-strings-normalize-tags | Normalize Tags | MOVE from 021→031 | — |
| 032 | kata-022-strconv-parse-positive-int | Parse Positive Int | MOVE from 022→032 | — |
| 033 | kata-023-time-business-hours | Business Hours Check | MOVE from 023→033 | — |
| 034 | kata-024-url-build-search-url | Build Search URL | MOVE from 024→034 | — |
| 035 | kata-025-json-pretty-print | Pretty JSON | MOVE from 025→035 | — |
| 036 | kata-026-os-env-default | Read Env Or Default | MOVE from 026→036 | — |
| 037 | kata-027-filepath-safe-join | Safe Join | MOVE from 027→037 | — |
| 038 | kata-028-context-wait-result | Context Wait | MOVE from 028→038 | — |
| 039 | kata-029-sync-concurrent-increment | Concurrent Increment | MOVE from 029→039 | — |
| 040 | kata-030-errors-user-lookup | Error User Lookup | MOVE from 030→040 | — |

#### Category: Core Collections and Algorithms (041–050)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 041 | kata-041-anagram-check | Anagram Check | KEEP | — |
| 042 | kata-042-min-max | Min/Max | KEEP | — |
| 043 | kata-043-rotate-slice | Rotate Slice | KEEP | — |
| 044 | kata-044-binary-search | Binary Search | KEEP | — |
| 045 | kata-045-merge-two-sorted-lists | Merge Sorted Lists | KEEP | — |
| 046 | kata-046-run-length-encoding | Run-Length Encoding | KEEP | — |
| 047 | kata-047-run-length-decoding | Run-Length Decoding | KEEP | — |
| 048 | kata-048-caesar-cipher | Caesar Cipher | KEEP | — |
| 049 | kata-049-isbn-10-validator | ISBN-10 Validator | KEEP | — |
| 050 | kata-050-roman-numerals | Roman Numerals | KEEP | — |

### Level: Engineer (051–090)

Outcome: Build and debug boundary-heavy Go code with explicit error handling and reliable runtime behavior.

#### Category: Text, Data, and Transformations (051–062)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 051 | kata-051-parse-csv-line | Parse CSV Line | KEEP | — |
| 052 | kata-052-json-pretty-print | JSON Pretty Print | REPLACE (duplicate of 035) | Table-Driven Test Refactor |
| 053 | kata-053-http-status-classifier | HTTP Status Classifier | KEEP | — |
| 054 | kata-054-time-window-check | Time Window Check | KEEP | — |
| 055 | kata-055-lru-cache-single-thread | LRU Cache (single-thread) | KEEP | — |
| 056 | kata-056-stack | Stack | KEEP | — |
| 057 | kata-057-bracket-matcher | Bracket Matcher | KEEP | — |
| 058 | kata-058-markdown-heading-extractor | Markdown Heading Extractor | KEEP | — |
| 059 | kata-059-longest-common-prefix | Longest Common Prefix | KEEP | — |
| 060 | kata-060-kebab-snake-camel | Kebab/Snake to Camel | KEEP | — |
| 061 | kata-061-file-extension-counter | File Extension Counter | KEEP | — |
| 062 | kata-062-top-n-words | Top N Words | KEEP | — |

#### Category: File, Time, and Config (063–070)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 063 | kata-063-csv-to-json | CSV to JSON | KEEP | — |
| 064 | kata-064-temperature-converter | Temperature Converter | KEEP | — |
| 065 | kata-065-log-line-parser | Log Line Parser | KEEP | — |
| 066 | kata-066-rate-limiter-token-bucket | Rate Limiter (token bucket) | KEEP | — |
| 067 | kata-067-ini-parser | INI Parser | KEEP | — |
| 068 | kata-068-mini-template-renderer | Mini Template Renderer | KEEP | — |
| 069 | kata-069-uuid-v4-generator | UUID v4 Generator | KEEP | — |
| 070 | kata-070-password-strength-scorer | Password Strength Scorer | KEEP | — |

#### Category: Services and Boundaries (071–082)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 071 | kata-071-http-query-builder | HTTP Query Builder | KEEP | — |
| 072 | kata-072-context-timeout-wrapper | Context Timeout Wrapper | KEEP | — |
| 073 | kata-073-retry-with-backoff | Retry with Backoff | KEEP | — |
| 074 | kata-074-concurrent-mapreduce | Concurrent MapReduce | KEEP | — |
| 075 | kata-075-worker-pool | Worker Pool | KEEP | — |
| 076 | kata-076-safe-counter | Safe Counter | KEEP | — |
| 077 | kata-077-debounce | Debounce | KEEP | — |
| 078 | kata-078-throttle | Throttle | KEEP | — |
| 079 | kata-079-json-patch-subset | JSON Patch (subset) | KEEP | — |
| 080 | kata-080-config-loader | Config Loader | KEEP | — |
| 081 | kata-081-http-client-with-retries | HTTP Client with Retries | KEEP | — |
| 082 | kata-082-tcp-echo-server | TCP Echo Server | KEEP | — |

#### Category: Service Patterns (083–090)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 083 | kata-083-line-oriented-reader | Line-Oriented Reader | KEEP | — |
| 084 | kata-084-json-lines-filter | JSON Lines Filter | KEEP | — |
| 085 | kata-085-mini-http-router | Mini HTTP Router | KEEP | — |
| 086 | kata-086-middleware-chain | Middleware Chain | KEEP | — |
| 087 | kata-087-in-memory-kv-store | In-Memory KV Store | KEEP | — |
| 088 | kata-088-simple-pub-sub | Simple Pub/Sub | KEEP | — |
| 089 | kata-089-fixed-window-metrics | Fixed-Window Metrics | KEEP | — |
| 090 | kata-090-trie-autocomplete | Trie Autocomplete | KEEP | — |

### Level: Senior (091–120)

Outcome: Design, debug, and evolve senior-level Go solutions with strong operational and maintainability standards.

#### Category: Concurrency and Reliability (091–100)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 091 | kata-091-http-server-with-graceful-shutdown | Graceful Shutdown Server | KEEP | — |
| 092 | kata-092-context-aware-logger | Context-Aware Logger | REWRITE (broken) | Context-Aware Logger |
| 093 | kata-093-structured-errors | Structured Errors | REWRITE (broken) | Structured Error Types |
| 094 | kata-094-circuit-breaker | Circuit Breaker | KEEP | — |
| 095 | kata-095-concurrent-file-downloader | Concurrent File Downloader | KEEP | — |
| 096 | kata-096-semaphore | Semaphore | KEEP | — |
| 097 | kata-097-barrier | Barrier | KEEP | — |
| 098 | kata-098-fan-in-fan-out-pipeline | Fan-in/Fan-out Pipeline | KEEP | — |
| 099 | kata-099-batcher | Batcher | KEEP | — |
| 100 | kata-100-lru-cache-thread-safe | LRU Cache (thread-safe) | REWRITE (broken) | LRU Cache (Thread-Safe) |

#### Category: Advanced Data Structures and Algorithms (101–108)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 101 | kata-101-token-bucket-thread-safe | Token Bucket (thread-safe) | REWRITE (broken) | Token Bucket (Thread-Safe) |
| 102 | kata-102-bloom-filter | Bloom Filter | REWRITE (broken) | Bloom Filter |
| 103 | kata-103-merkle-tree | Merkle Tree | REWRITE (broken) | Merkle Tree |
| 104 | kata-104-binary-heap-priority-queue | Binary Heap Priority Queue | REWRITE (broken) | Priority Queue (Binary Heap) |
| 105 | kata-105-dijkstra-shortest-path | Dijkstra Shortest Path | REWRITE (broken) | Shortest Path (Dijkstra) |
| 106 | kata-106-a-grid-pathfinding | A* Grid Pathfinding | REWRITE (broken) | A* Grid Pathfinding |
| 107 | kata-107-event-sourcing-mini | Event Sourcing Mini | REPLACE (too academic) | Bug Fix: Goroutine Leak |
| 108 | kata-108-sqlite-backed-repo-pure-go-driver | SQLite-backed Repo | REPLACE (external dep) | Bug Fix: Data Race |

#### Category: Tooling and Security (109–120)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 109 | kata-109-json-schema-validator-subset | JSON Schema Validator | REWRITE (broken) | JSON Schema Validator (Subset) |
| 110 | kata-110-rate-limited-http-scraper | Rate-Limited HTTP Scraper | REWRITE (broken) | Rate-Limited HTTP Client |
| 111 | kata-111-websocket-chat-basic | WebSocket Chat | REPLACE (too academic) | Bug Fix: Timeout Handling |
| 112 | kata-112-sse-stream | SSE Stream | REPLACE (too academic) | Custom Test Helper |
| 113 | kata-113-jwt-sign-verify | JWT Sign/Verify | KEEP | — |
| 114 | kata-114-hmac-request-signing | HMAC Request Signing | KEEP | — |
| 115 | kata-115-password-hashing-bcrypt-argon2 | Password Hashing | KEEP | — |
| 116 | kata-116-file-integrity-checker | File Integrity Checker | KEEP | — |
| 117 | kata-117-tar-gzip-archiver | Tar/Gzip Archiver | KEEP | — |
| 118 | kata-118-diff-line-based | Diff (line-based) | KEEP | — |

### Level: Lead (119–140)

Outcome: Apply production-grade quality standards and design tradeoff reasoning to professional Go solutions.

#### Category: CLI Tooling (119–124)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 119 | kata-119-mini-grep | Mini grep | KEEP | — |
| 120 | kata-120-mini-wc | Mini wc | KEEP | — |
| 121 | kata-121-mini-cut | Mini cut | KEEP | — |
| 122 | kata-122-mini-head-tail | Mini head/tail | KEEP | — |
| 123 | kata-123-mini-sort | Mini sort | KEEP | — |
| 124 | kata-124-mini-uniq | Mini uniq | KEEP | — |

#### Category: Testing and Performance (125–130)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 125 | kata-125-concurrent-test-harness | Concurrent Test Harness | REWRITE (broken) | Concurrent Test Harness |
| 126 | kata-126-benchmarking-kata | Benchmarking Kata | KEEP | — |
| 127 | kata-127-generics-set-map-utilities | Generics: Set/Map Utilities | REWRITE (broken) | Generics: Set/Map Utilities |
| 128 | kata-128-generics-optional-result | Generics: Optional/Result | REWRITE (broken) | Generics: Optional/Result |
| 129 | kata-129-reflection-struct-tag-parser | Reflection: Struct Tag Parser | KEEP | — |
| 130 | kata-130-plugin-free-di-container | Plugin-Free DI Container | KEEP | — |

#### Category: Databases and I/O (131–135)

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 131 | kata-136-io-count-lines | Count Lines | MOVE from 136→131 | — |
| 132 | kata-137-io-copy-with-limit | Copy With Limit | MOVE from 137→132 | — |
| 133 | kata-138-csv-read-records | Read CSV Records | MOVE from 138→133 | — |
| 134 | kata-139-io-read-exactly-n | Read Exactly N | MOVE from 139→134 | — |
| 135 | kata-140-db-run-transaction | Run In Transaction | MOVE from 140→135 | — |

#### Category: Bug Fix Lab (136–140) ← MOVED to end, expanded

| # | Current Dir | Current Title | Disposition | New Title (if changed) |
|---|------------|---------------|-------------|----------------------|
| 136 | kata-131-bug-username-normalization | Bug: Username Normalization | MOVE from 131→136 | — |
| 137 | kata-132-bug-sum-positive | Bug: Sum Positive | MOVE from 132→137 | — |
| 138 | kata-133-bug-first-non-empty | Bug: First Non-Empty | MOVE from 133→138 | — |
| 139 | kata-134-bug-parse-flag | Bug: Parse Flag | MOVE from 134→139 | — |
| 140 | kata-135-bug-clamp-percentage | Bug: Clamp Percentage | MOVE from 135→140 | — |

### Disposition Summary

| Disposition | Count | Description |
|------------|-------|-------------|
| KEEP | 102 | No changes needed — correct position, proper implementation |
| MOVE | 30 | Directory renumbered due to block swap or category restructure |
| REWRITE | 14 | Broken stubs → proper implementation with correct signatures/tests |
| REPLACE | 4 | Wrong topic for level → replaced with practical engineering kata |

---

## Artefact 2: Full Old-to-New Mapping Table

Every current kata directory mapped to its new position.

| Old # | Old Directory | New # | Action | Notes |
|-------|--------------|-------|--------|-------|
| 001 | kata-001-build-greeting | 001 | KEEP | — |
| 002 | kata-002-parse-whole-number | 002 | KEEP | — |
| 003 | kata-003-basic-calculator | 003 | KEEP | — |
| 004 | kata-004-even-odd-label | 004 | KEEP | — |
| 005 | kata-005-clamp-value | 005 | KEEP | — |
| 006 | kata-006-sum-first-n | 006 | KEEP | — |
| 007 | kata-007-countdown-slice | 007 | KEEP | — |
| 008 | kata-008-repeat-text | 008 | KEEP | — |
| 009 | kata-009-char-frequency | 009 | KEEP | — |
| 010 | kata-010-average-int | 010 | KEEP | — |
| 011 | kata-011-swap-first-last | 011 | KEEP | — |
| 012 | kata-012-remove-empty-strings | 012 | KEEP | — |
| 013 | kata-013-grade-classifier | 013 | KEEP | — |
| 014 | kata-014-count-true | 014 | KEEP | — |
| 015 | kata-015-find-first-index | 015 | KEEP | — |
| 016 | kata-016-safe-slice-range | 016 | KEEP | — |
| 017 | kata-017-budget-status | 017 | KEEP | — |
| 018 | kata-018-merge-alternating | 018 | KEEP | — |
| 019 | kata-019-word-lengths | 019 | KEEP | — |
| 020 | kata-020-incomplete-task-filter | 020 | KEEP | — |
| 021 | kata-021-strings-normalize-tags | 031 | MOVE | Swap: packages block moves to 031–040 |
| 022 | kata-022-strconv-parse-positive-int | 032 | MOVE | Swap: packages block moves to 031–040 |
| 023 | kata-023-time-business-hours | 033 | MOVE | Swap: packages block moves to 031–040 |
| 024 | kata-024-url-build-search-url | 034 | MOVE | Swap: packages block moves to 031–040 |
| 025 | kata-025-json-pretty-print | 035 | MOVE | Swap: packages block moves to 031–040 |
| 026 | kata-026-os-env-default | 036 | MOVE | Swap: packages block moves to 031–040 |
| 027 | kata-027-filepath-safe-join | 037 | MOVE | Swap: packages block moves to 031–040 |
| 028 | kata-028-context-wait-result | 038 | MOVE | Swap: packages block moves to 031–040 |
| 029 | kata-029-sync-concurrent-increment | 039 | MOVE | Swap: packages block moves to 031–040 |
| 030 | kata-030-errors-user-lookup | 040 | MOVE | Swap: packages block moves to 031–040 |
| 031 | kata-031-fizzbuzz | 021 | MOVE | Swap: foundations block moves to 021–030 |
| 032 | kata-032-sum-of-integers | 022 | MOVE | Swap: foundations block moves to 021–030 |
| 033 | kata-033-reverse-string | 023 | MOVE | Swap: foundations block moves to 021–030 |
| 034 | kata-034-palindrome-check | 024 | MOVE | Swap: foundations block moves to 021–030 |
| 035 | kata-035-count-vowels | 025 | MOVE | Swap: foundations block moves to 021–030 |
| 036 | kata-036-factorial | 026 | MOVE | Swap: foundations block moves to 021–030 |
| 037 | kata-037-fibonacci | 027 | MOVE | Swap: foundations block moves to 021–030 |
| 038 | kata-038-filter-even | 028 | MOVE | Swap: foundations block moves to 021–030 |
| 039 | kata-039-unique-strings | 029 | MOVE | Swap: foundations block moves to 021–030 |
| 040 | kata-040-word-count | 030 | MOVE | Swap: foundations block moves to 021–030 |
| 041 | kata-041-anagram-check | 041 | KEEP | — |
| 042 | kata-042-min-max | 042 | KEEP | — |
| 043 | kata-043-rotate-slice | 043 | KEEP | — |
| 044 | kata-044-binary-search | 044 | KEEP | — |
| 045 | kata-045-merge-two-sorted-lists | 045 | KEEP | — |
| 046 | kata-046-run-length-encoding | 046 | KEEP | — |
| 047 | kata-047-run-length-decoding | 047 | KEEP | — |
| 048 | kata-048-caesar-cipher | 048 | KEEP | — |
| 049 | kata-049-isbn-10-validator | 049 | KEEP | — |
| 050 | kata-050-roman-numerals | 050 | KEEP | — |
| 051 | kata-051-parse-csv-line | 051 | KEEP | — |
| 052 | kata-052-json-pretty-print | 052 | REPLACE | Duplicate of 025 → Table-Driven Test Refactor |
| 053 | kata-053-http-status-classifier | 053 | KEEP | — |
| 054 | kata-054-time-window-check | 054 | KEEP | — |
| 055 | kata-055-lru-cache-single-thread | 055 | KEEP | — |
| 056 | kata-056-stack | 056 | KEEP | — |
| 057 | kata-057-bracket-matcher | 057 | KEEP | — |
| 058 | kata-058-markdown-heading-extractor | 058 | KEEP | — |
| 059 | kata-059-longest-common-prefix | 059 | KEEP | — |
| 060 | kata-060-kebab-snake-camel | 060 | KEEP | — |
| 061 | kata-061-file-extension-counter | 061 | KEEP | — |
| 062 | kata-062-top-n-words | 062 | KEEP | — |
| 063 | kata-063-csv-to-json | 063 | KEEP | — |
| 064 | kata-064-temperature-converter | 064 | KEEP | — |
| 065 | kata-065-log-line-parser | 065 | KEEP | — |
| 066 | kata-066-rate-limiter-token-bucket | 066 | KEEP | — |
| 067 | kata-067-ini-parser | 067 | KEEP | — |
| 068 | kata-068-mini-template-renderer | 068 | KEEP | — |
| 069 | kata-069-uuid-v4-generator | 069 | KEEP | — |
| 070 | kata-070-password-strength-scorer | 070 | KEEP | — |
| 071 | kata-071-http-query-builder | 071 | KEEP | — |
| 072 | kata-072-context-timeout-wrapper | 072 | KEEP | — |
| 073 | kata-073-retry-with-backoff | 073 | KEEP | — |
| 074 | kata-074-concurrent-mapreduce | 074 | KEEP | — |
| 075 | kata-075-worker-pool | 075 | KEEP | — |
| 076 | kata-076-safe-counter | 076 | KEEP | — |
| 077 | kata-077-debounce | 077 | KEEP | — |
| 078 | kata-078-throttle | 078 | KEEP | — |
| 079 | kata-079-json-patch-subset | 079 | KEEP | — |
| 080 | kata-080-config-loader | 080 | KEEP | — |
| 081 | kata-081-http-client-with-retries | 081 | KEEP | — |
| 082 | kata-082-tcp-echo-server | 082 | KEEP | — |
| 083 | kata-083-line-oriented-reader | 083 | KEEP | — |
| 084 | kata-084-json-lines-filter | 084 | KEEP | — |
| 085 | kata-085-mini-http-router | 085 | KEEP | — |
| 086 | kata-086-middleware-chain | 086 | KEEP | — |
| 087 | kata-087-in-memory-kv-store | 087 | KEEP | — |
| 088 | kata-088-simple-pub-sub | 088 | KEEP | — |
| 089 | kata-089-fixed-window-metrics | 089 | KEEP | — |
| 090 | kata-090-trie-autocomplete | 090 | KEEP | — |
| 091 | kata-091-http-server-with-graceful-shutdown | 091 | KEEP | — |
| 092 | kata-092-context-aware-logger | 092 | REWRITE | Broken stub → proper implementation |
| 093 | kata-093-structured-errors | 093 | REWRITE | Broken stub → proper implementation |
| 094 | kata-094-circuit-breaker | 094 | KEEP | — |
| 095 | kata-095-concurrent-file-downloader | 095 | KEEP | — |
| 096 | kata-096-semaphore | 096 | KEEP | — |
| 097 | kata-097-barrier | 097 | KEEP | — |
| 098 | kata-098-fan-in-fan-out-pipeline | 098 | KEEP | — |
| 099 | kata-099-batcher | 099 | KEEP | — |
| 100 | kata-100-lru-cache-thread-safe | 100 | REWRITE | Broken stub → proper implementation |
| 101 | kata-101-token-bucket-thread-safe | 101 | REWRITE | Broken stub → proper implementation |
| 102 | kata-102-bloom-filter | 102 | REWRITE | Broken stub → proper implementation |
| 103 | kata-103-merkle-tree | 103 | REWRITE | Broken stub → proper implementation |
| 104 | kata-104-binary-heap-priority-queue | 104 | REWRITE | Broken stub → proper implementation |
| 105 | kata-105-dijkstra-shortest-path | 105 | REWRITE | Broken stub → proper implementation |
| 106 | kata-106-a-grid-pathfinding | 106 | REWRITE | Broken stub → proper implementation |
| 107 | kata-107-event-sourcing-mini | 107 | REPLACE | Too academic → Bug Fix: Goroutine Leak |
| 108 | kata-108-sqlite-backed-repo-pure-go-driver | 108 | REPLACE | External dep → Bug Fix: Data Race |
| 109 | kata-109-json-schema-validator-subset | 109 | REWRITE | Broken stub → proper implementation |
| 110 | kata-110-rate-limited-http-scraper | 110 | REWRITE | Broken stub → proper implementation |
| 111 | kata-111-websocket-chat-basic | 111 | REPLACE | Too academic → Bug Fix: Timeout Handling |
| 112 | kata-112-sse-stream | 112 | REPLACE | Too academic → Custom Test Helper |
| 113 | kata-113-jwt-sign-verify | 113 | KEEP | — |
| 114 | kata-114-hmac-request-signing | 114 | KEEP | — |
| 115 | kata-115-password-hashing-bcrypt-argon2 | 115 | KEEP | — |
| 116 | kata-116-file-integrity-checker | 116 | KEEP | — |
| 117 | kata-117-tar-gzip-archiver | 117 | KEEP | — |
| 118 | kata-118-diff-line-based | 118 | KEEP | — |
| 119 | kata-119-mini-grep | 119 | KEEP | — |
| 120 | kata-120-mini-wc | 120 | KEEP | — |
| 121 | kata-121-mini-cut | 121 | KEEP | — |
| 122 | kata-122-mini-head-tail | 122 | KEEP | — |
| 123 | kata-123-mini-sort | 123 | KEEP | — |
| 124 | kata-124-mini-uniq | 124 | KEEP | — |
| 125 | kata-125-concurrent-test-harness | 125 | REWRITE | Broken stub → proper implementation |
| 126 | kata-126-benchmarking-kata | 126 | KEEP | — |
| 127 | kata-127-generics-set-map-utilities | 127 | REWRITE | Broken stub → proper implementation |
| 128 | kata-128-generics-optional-result | 128 | REWRITE | Broken stub → proper implementation |
| 129 | kata-129-reflection-struct-tag-parser | 129 | KEEP | — |
| 130 | kata-130-plugin-free-di-container | 130 | KEEP | — |
| 131 | kata-131-bug-username-normalization | 136 | MOVE | Bug-fix lab moves to 136–140 |
| 132 | kata-132-bug-sum-positive | 137 | MOVE | Bug-fix lab moves to 136–140 |
| 133 | kata-133-bug-first-non-empty | 138 | MOVE | Bug-fix lab moves to 136–140 |
| 134 | kata-134-bug-parse-flag | 139 | MOVE | Bug-fix lab moves to 136–140 |
| 135 | kata-135-bug-clamp-percentage | 140 | MOVE | Bug-fix lab moves to 136–140 |
| 136 | kata-136-io-count-lines | 131 | MOVE | DB/IO moves to 131–135 |
| 137 | kata-137-io-copy-with-limit | 132 | MOVE | DB/IO moves to 131–135 |
| 138 | kata-138-csv-read-records | 133 | MOVE | DB/IO moves to 131–135 |
| 139 | kata-139-io-read-exactly-n | 134 | MOVE | DB/IO moves to 131–135 |
| 140 | kata-140-db-run-transaction | 135 | MOVE | DB/IO moves to 131–135 |

---

## Artefact 3: Testing Curriculum Map

### Design

Testing skills are woven throughout the curriculum rather than isolated in a single block. Each testing kata builds on the previous ones. The "Test" column below shows which katas include explicit testing instruction.

### Progressive Testing Skills

| Level | Kata | Testing Skill Introduced | How It's Taught |
|-------|------|-------------------------|-----------------|
| Associate | 001–005 | Read and understand test files | Tests provided; learner reads them to understand contracts |
| Associate | 006–010 | Test-first thinking | README specifies: "read the tests before writing code" |
| Associate | 021 (FizzBuzz) | Writing basic test cases | Kata includes exercise: "add 2 edge-case tests" |
| Associate | 035 (PrettyJSON) | Test error paths | Tests include malformed input; learner must handle them |
| Associate | 041–050 | Table-driven test reading | Tests use `[]struct{name, input, want}` pattern throughout |
| Engineer | **052 (NEW)** | **Table-Driven Test Refactor** | Learner refactors a working kata's tests into table-driven style |
| Engineer | 063–070 | Testing with file I/O | Tests use `t.TempDir()`, test fixtures, cleanup |
| Engineer | 071–082 | Testing with interfaces | Tests inject mock/stub implementations |
| Engineer | 085–086 | HTTP handler testing | Tests use `httptest.NewRecorder()` and `httptest.NewServer()` |
| Senior | 092 | Testing with context | Tests verify context propagation and cancellation |
| Senior | **107 (NEW)** | **Bug Fix: Goroutine Leak** | Learner finds and fixes a goroutine leak, writes regression test |
| Senior | **108 (NEW)** | **Bug Fix: Data Race** | Learner finds data race with `-race`, writes concurrent regression test |
| Senior | **111 (NEW)** | **Bug Fix: Timeout Handling** | Learner fixes timeout bug, writes time-sensitive test |
| Senior | **112 (NEW)** | **Custom Test Helper** | Learner builds `t.Helper()` functions and test assertion utilities |
| Lead | **125** | **Concurrent Test Harness** | Learner builds a concurrent test runner with `t.Parallel()` |
| Lead | **126** | **Benchmarking** | Learner writes `BenchmarkXxx` functions and interprets results |

### Testing Concepts Coverage

| Concept | First Introduced | Reinforced At |
|---------|-----------------|---------------|
| Reading test files | 001 | Every kata |
| `t.Run` subtests | 041 | 052, 085 |
| Table-driven tests | 041 | 052 (explicit refactor kata) |
| `t.TempDir()` / fixtures | 061 | 063, 116, 117 |
| `httptest` | 085 | 086, 091 |
| `t.Helper()` | 112 (NEW) | 125 |
| `t.Parallel()` | 125 | — |
| `-race` flag | 108 (NEW) | 125 |
| `testing.B` benchmarks | 126 | — |
| Mock/stub injection | 071 | 073, 081 |
| Regression tests | 107 (NEW) | 108, 111, 136–140 |

### New Testing Katas: Specifications

#### Kata 052: Table-Driven Test Refactor (replaces duplicate PrettyJSON)
- **Focus**: Convert procedural tests to table-driven style
- **Input**: A working function + its procedural test file
- **Task**: Rewrite tests using `[]struct{name string; input X; want Y}` with `t.Run`
- **Signature**: `func Deduplicate(items []string) []string` (simple function, complex tests)
- **Tests provided**: Procedural tests that learner must refactor

#### Kata 107: Bug Fix — Goroutine Leak (replaces Event Sourcing)
- **Focus**: Detect and fix goroutine leak, write regression test
- **Input**: Buggy function that spawns goroutines without cleanup
- **Task**: Fix the leak, add `runtime.NumGoroutine()` regression test
- **Signature**: `func ProcessStream(ctx context.Context, ch <-chan string) []string`
- **Bug-fix mode**: Pre-broken `kata.go` with deliberate goroutine leak

#### Kata 108: Bug Fix — Data Race (replaces SQLite Repo)
- **Focus**: Find data race with `-race`, fix with proper synchronization
- **Input**: Buggy concurrent map access
- **Task**: Fix the race, write test that fails under `-race` without fix
- **Signature**: `func NewConcurrentMap() *ConcurrentMap`
- **Bug-fix mode**: Pre-broken `kata.go` with deliberate race condition

#### Kata 111: Bug Fix — Timeout Handling (replaces WebSocket Chat)
- **Focus**: Fix timeout-related bug in HTTP handler
- **Input**: Handler that doesn't respect context deadline
- **Task**: Fix timeout handling, write test with short deadline
- **Signature**: `func HandleWithTimeout(ctx context.Context, work func() (string, error)) (string, error)`
- **Bug-fix mode**: Pre-broken `kata.go` with missing deadline check

#### Kata 112: Custom Test Helper (replaces SSE Stream)
- **Focus**: Build reusable test helpers with `t.Helper()`
- **Input**: Specification for test assertion helpers
- **Task**: Implement `AssertEqual`, `AssertError`, `AssertPanics` as `t.Helper()` functions
- **Signature**: `func AssertEqual(t testing.TB, got, want any)` (plus others)
- **Tests provided**: Meta-tests that verify the helpers work correctly

---

## Artefact 4: Pathway and Category Structure

### New pathways.json

```json
{
  "pathways": [
    {
      "id": "associate",
      "title": "Associate",
      "description": "Build consistent Go implementation habits from programming basics through standard library fluency.",
      "categories": [
        "programming-essentials",
        "data-and-logic",
        "go-foundations",
        "well-known-packages",
        "core-collections-and-algorithms"
      ],
      "recommended_modes": [
        "Documentation",
        "Flashcards",
        "Workbench",
        "Quiz",
        "Reflection"
      ],
      "level_outcome": "You can implement small-to-medium Go functions with clear contracts, edge-case coverage, and readable tests."
    },
    {
      "id": "engineer",
      "title": "Engineer",
      "description": "Shift from isolated functions to boundary-aware service code, testing discipline, and debugging skills.",
      "categories": [
        "text-data-and-transformations",
        "file-time-and-config",
        "services-and-boundaries",
        "service-patterns"
      ],
      "recommended_modes": [
        "Documentation",
        "Workbench",
        "Bug Hunt",
        "Quiz",
        "Reflection"
      ],
      "level_outcome": "You can build and debug boundary-heavy Go code with explicit error handling and reliable runtime behavior."
    },
    {
      "id": "senior",
      "title": "Senior",
      "description": "Master concurrency, reliability patterns, advanced data structures, and production security practices.",
      "categories": [
        "concurrency-and-reliability",
        "advanced-data-structures-and-algorithms",
        "tooling-and-security"
      ],
      "recommended_modes": [
        "Documentation",
        "Workbench",
        "Bug Hunt",
        "Flashcards",
        "Quiz",
        "Reflection"
      ],
      "level_outcome": "You can design, debug, and evolve senior-level Go solutions with strong operational and maintainability standards."
    },
    {
      "id": "lead",
      "title": "Lead",
      "description": "Apply production-grade quality standards, testing discipline, and architectural reasoning across complete systems.",
      "categories": [
        "cli-tooling",
        "testing-and-performance",
        "databases-and-io",
        "bug-fix-lab"
      ],
      "recommended_modes": [
        "Documentation",
        "Workbench",
        "Bug Hunt",
        "Flashcards",
        "Quiz",
        "Reflection"
      ],
      "level_outcome": "You can apply production-grade quality standards and design tradeoff reasoning to professional Go solutions."
    }
  ]
}
```

### New track.json categories

```json
{
  "id": "go-core-100",
  "title": "Go Core Path: Associate to Lead (140 Katas)",
  "description": "Structured Go learning from first programming principles to lead-level engineering, with guided pathways, testing discipline, bug-fix practice, and professional I/O patterns.",
  "categories": [
    {
      "id": "programming-essentials",
      "title": "Programming Essentials",
      "description": "Zero-to-code foundations for learners new to programming and Go.",
      "learning_goal": "Understand functions, conditionals, loops, and boundaries so you can turn plain rules into working code.",
      "kata_ranges": [{"start": 1, "end": 10}]
    },
    {
      "id": "data-and-logic",
      "title": "Data and Logic",
      "description": "Practical beginner problems using slices, maps, structs, and search/filter logic.",
      "learning_goal": "Build confidence with common data operations and predictable return contracts.",
      "kata_ranges": [{"start": 11, "end": 20}]
    },
    {
      "id": "go-foundations",
      "title": "Go Foundations",
      "description": "Core syntax, control flow, and deterministic behavior with everyday coding patterns.",
      "learning_goal": "Convert plain-language requirements into clear, testable Go functions that handle edge cases intentionally.",
      "kata_ranges": [{"start": 21, "end": 30}]
    },
    {
      "id": "well-known-packages",
      "title": "Well-Known Packages",
      "description": "Focused katas for core standard library packages used in everyday Go code.",
      "learning_goal": "Use common packages with confidence by implementing deterministic behavior contracts.",
      "kata_ranges": [{"start": 31, "end": 40}]
    },
    {
      "id": "core-collections-and-algorithms",
      "title": "Core Collections and Algorithms",
      "description": "Slices, maps, generics, and algorithmic thinking for correctness under constraints.",
      "learning_goal": "Use Go data structures and iteration patterns to solve problems without sacrificing readability or safety.",
      "kata_ranges": [{"start": 41, "end": 50}]
    },
    {
      "id": "text-data-and-transformations",
      "title": "Text, Data, and Transformations",
      "description": "Parsing, formatting, and data transformation across structured and unstructured inputs.",
      "learning_goal": "Design robust input/output behavior that is explicit about malformed data and contract guarantees.",
      "kata_ranges": [{"start": 51, "end": 62}]
    },
    {
      "id": "file-time-and-config",
      "title": "File, Time, and Config",
      "description": "File handling, time-based behavior, and configuration safety in local systems.",
      "learning_goal": "Handle operating-system boundaries and time-sensitive logic with predictable, testable behavior.",
      "kata_ranges": [{"start": 63, "end": 70}]
    },
    {
      "id": "services-and-boundaries",
      "title": "Services and Boundaries",
      "description": "HTTP, networking, middleware, and service composition across unstable external boundaries.",
      "learning_goal": "Treat I/O boundaries as contracts and return safe behavior under failure and partial success.",
      "kata_ranges": [{"start": 71, "end": 82}]
    },
    {
      "id": "service-patterns",
      "title": "Service Patterns",
      "description": "Reusable patterns for routing, middleware, storage, messaging, and observability.",
      "learning_goal": "Compose service building blocks into maintainable, testable application layers.",
      "kata_ranges": [{"start": 83, "end": 90}]
    },
    {
      "id": "concurrency-and-reliability",
      "title": "Concurrency and Reliability",
      "description": "Lifecycle ownership, cancellation, coordination primitives, and resilient runtime behavior.",
      "learning_goal": "Build concurrent systems that stay correct under load, interruption, and timing variance.",
      "kata_ranges": [{"start": 91, "end": 100}]
    },
    {
      "id": "advanced-data-structures-and-algorithms",
      "title": "Advanced Data Structures and Algorithms",
      "description": "Thread-safe collections, probabilistic structures, graph algorithms, and engineering debugging.",
      "learning_goal": "Implement and reason about advanced data structures with correctness proofs and performance awareness.",
      "kata_ranges": [{"start": 101, "end": 108}]
    },
    {
      "id": "tooling-and-security",
      "title": "Tooling and Security",
      "description": "Validation, security primitives, testing practices, and production hardening.",
      "learning_goal": "Apply security-first thinking and testing rigor to production Go code.",
      "kata_ranges": [{"start": 109, "end": 118}]
    },
    {
      "id": "cli-tooling",
      "title": "CLI Tooling",
      "description": "Build Unix-style command-line tools using Go's I/O and flag packages.",
      "learning_goal": "Design composable CLI tools that handle stdin/stdout, flags, and error reporting correctly.",
      "kata_ranges": [{"start": 119, "end": 124}]
    },
    {
      "id": "testing-and-performance",
      "title": "Testing and Performance",
      "description": "Advanced testing patterns, generics, reflection, and architectural design.",
      "learning_goal": "Master Go's testing package and type system for professional-grade code.",
      "kata_ranges": [{"start": 125, "end": 130}]
    },
    {
      "id": "databases-and-io",
      "title": "Databases and I/O",
      "description": "Persistence, streaming, filesystem, and query-boundary contracts.",
      "learning_goal": "Build confidence with database and I/O boundaries where correctness, error handling, and resource ownership are critical.",
      "kata_ranges": [{"start": 131, "end": 135}]
    },
    {
      "id": "bug-fix-lab",
      "title": "Bug Fix Lab",
      "description": "Deliberately buggy implementations where learners diagnose failures and repair behavior without rewriting.",
      "learning_goal": "Develop production debugging habits: reproduce, isolate, patch minimally, and confirm with regression tests.",
      "kata_ranges": [{"start": 136, "end": 140}]
    }
  ]
}
```

### Category Changes Summary

| Old Category | Old Range | New Category | New Range | Change |
|-------------|-----------|-------------|-----------|--------|
| Starter: Programming Essentials | 001–010 | Programming Essentials | 001–010 | Renamed (drop "Starter:") |
| Starter: Data and Logic | 011–020 | Data and Logic | 011–020 | Renamed |
| Well-Known Packages | 021–030 | Go Foundations | 021–030 | **SWAPPED** — foundations first |
| Go Foundations | 031–040 | Well-Known Packages | 031–040 | **SWAPPED** — packages second |
| Core Collections and Algorithms | 041–050 | Core Collections and Algorithms | 041–050 | No change |
| Text, Data, and Transformations | 051–062 | Text, Data, and Transformations | 051–062 | No change |
| File, Time, and Config | 063–070 | File, Time, and Config | 063–070 | No change |
| Services and Boundaries | 071–090 | Services and Boundaries | 071–082 | **SPLIT** — 20→12 katas |
| — | — | Service Patterns | 083–090 | **NEW** — split from above |
| Concurrency and Reliability | 091–110 | Concurrency and Reliability | 091–100 | **SHRUNK** — 20→10 katas |
| — | — | Advanced DS & Algorithms | 101–108 | **NEW** — carved from above |
| Tooling, Security, Advanced Design | 111–130 | Tooling and Security | 109–118 | **SHRUNK/MOVED** |
| — | — | CLI Tooling | 119–124 | **NEW** — carved from above |
| — | — | Testing and Performance | 125–130 | **NEW** — carved from above |
| Bug Fix Lab | 131–135 | Bug Fix Lab | 136–140 | **MOVED** to end |
| Databases and I/O | 136–140 | Databases and I/O | 131–135 | **MOVED** before bugs |

---

## Artefact 5: Migration Impact Analysis

### Files Affected

#### 1. Directory Renames (30 moves)

The block swap (021–030 ↔ 031–040) requires renaming 20 directories.
The tail swap (131–135 ↔ 136–140) requires renaming 10 directories.

**Risk**: Directory renames are the highest-risk operation. Must be done atomically per block to avoid ID collisions.

**Approach**: Use a temporary staging directory to avoid collisions:
```
katas/kata-021-* → katas/.staging/kata-031-*   (first pass)
katas/kata-031-* → katas/kata-021-*            (second pass)
katas/.staging/kata-031-* → katas/kata-031-*   (third pass)
```

Affected directories:
- `kata-021-strings-normalize-tags` → `kata-031-strings-normalize-tags`
- `kata-022-strconv-parse-positive-int` → `kata-032-strconv-parse-positive-int`
- ... (all 10 in the 021→031 direction)
- `kata-031-fizzbuzz` → `kata-021-fizzbuzz`
- `kata-032-sum-of-integers` → `kata-022-sum-of-integers`
- ... (all 10 in the 031→021 direction)
- `kata-131-bug-*` → `kata-136-bug-*` (5 directories)
- `kata-136-io-*` through `kata-140-*` → `kata-131-*` through `kata-135-*` (5 directories)

#### 2. Kata Rewrites (14 files)

Each broken kata needs:
- `kata.go`: New function signature matching the kata title, TODO body
- `kata_test.go`: Proper test structure (not just `t.Skip`)
- `README.md`: Full specification with contract, examples, edge cases, rules

Affected katas: 092, 093, 100, 101, 102, 103, 104, 105, 106, 109, 110, 125, 127, 128

#### 3. Kata Replacements (4 files)

Complete new kata content needed:
- 052: Table-Driven Test Refactor (replacing duplicate PrettyJSON)
- 107: Bug Fix — Goroutine Leak (replacing Event Sourcing)
- 108: Bug Fix — Data Race (replacing SQLite Repo)
- 111: Bug Fix — Timeout Handling (replacing WebSocket Chat)
- 112: Custom Test Helper (replacing SSE Stream)

Wait — that's 5 replacements. The 052 replacement plus 4 topic replacements (107, 108, 111, 112).

Each replacement needs: new `go.mod`, `kata.go`, `kata_test.go`, `README.md`, and directory rename.

#### 4. Configuration File Updates

- `tracks/go-core-100/track.json`: New category structure (16 categories, new ranges)
- `tracks/go-core-100/pathways.json`: 4 pathways replacing 3, new category IDs

#### 5. Internal `go.mod` Files

Every kata has its own `go.mod` with `module kata-NNN-slug`. Directory renames must update the module path in `go.mod` to match the new directory name.

Affected: 30 moved directories × 1 `go.mod` each = 30 `go.mod` updates.

### Progress State Impact

The progress store (`/.learning/progress.json`) tracks kata IDs like `"kata-031-fizzbuzz"`. After renumbering:

- **Safe approach**: Include a `migrations` map in the progress store, or do a one-time migration script that remaps old IDs to new IDs.
- **Simplest approach**: Wipe progress on migration. Acceptable for a pre-release product.
- **Recommended**: Build an ID migration map into the progress loading code that translates old→new IDs on first load.

Migration map (subset):
```
"kata-031-fizzbuzz" → "kata-021-fizzbuzz"
"kata-021-strings-normalize-tags" → "kata-031-strings-normalize-tags"
"kata-131-bug-username-normalization" → "kata-136-bug-username-normalization"
"kata-136-io-count-lines" → "kata-131-io-count-lines"
```

### Edge Cases

1. **Test runner references**: `scripts/test_operability.sh` runs `go test ./apps/learner-studio ./apps/learner-desktop ./internal/learning/...` — no kata-specific references, safe.

2. **Catalog loader**: `internal/learning/catalog` discovers katas by scanning `katas/kata-NNN-*` directories. The naming convention is maintained, so this continues to work.

3. **Git history**: Directory renames will show as delete+create in git unless `git mv` is used. Using `git mv` preserves history tracking.

4. **In-flight learner state**: If someone is mid-kata during migration, their saved code in `kata.go` will be at the old path. The migration script should preserve file contents for KEEP/MOVE katas.

5. **Replaced katas (107, 108, 111, 112)**: Any learner progress on these is lost. Acceptable — the old implementations were broken stubs with no real work possible.

6. **Duplicate resolution (052)**: The old kata-052 has `PrettyJSON` with a compile error (`return nil` for a two-value return). Replacing it loses nothing.

### Workload Estimate

| Task | Count | Effort Per | Total |
|------|-------|-----------|-------|
| Directory renames | 30 | Script-automated | ~1 hour |
| go.mod updates | 30 | Script-automated | Included above |
| Broken kata rewrites | 14 | ~30 min each | ~7 hours |
| Kata replacements | 5 | ~45 min each | ~4 hours |
| track.json rewrite | 1 | ~30 min | 30 min |
| pathways.json rewrite | 1 | ~20 min | 20 min |
| Progress migration code | 1 | ~1 hour | 1 hour |
| Verification/testing | 1 | ~2 hours | 2 hours |
| **Total** | | | **~16 hours** |

---

## Artefact 6: Phase 2B Execution Plan

### Guiding Principles

1. Every step must leave the repo in a compilable, test-passing state
2. Commits are atomic — one logical change per commit
3. Directory renames use `git mv` for history preservation
4. Verification runs after every batch

### Step 1: Configuration Updates (no kata changes)

**Commit 1a**: Update `track.json` with new 16-category structure
**Commit 1b**: Update `pathways.json` with 4-pathway structure
**Commit 1c**: Update `internal/learning/catalog` if category IDs changed (verify loader works)

**Verification**: `go test ./internal/learning/... ./apps/learner-studio/...`

### Step 2: Block Swap — Foundations ↔ Packages (021–040)

This is the riskiest operation. Use staging directory approach.

**Commit 2a**: Move kata-021 through kata-030 to staging
```bash
mkdir -p katas/.staging
for i in $(seq 21 30); do
  old=$(ls -d katas/kata-$(printf "%03d" $i)-*)
  new_num=$(printf "%03d" $((i + 10)))
  slug=$(basename "$old" | sed "s/kata-$(printf "%03d" $i)/kata-${new_num}/")
  git mv "$old" "katas/.staging/$slug"
done
```

**Commit 2b**: Move kata-031 through kata-040 to their new positions (021–030)
```bash
for i in $(seq 31 40); do
  old=$(ls -d katas/kata-$(printf "%03d" $i)-*)
  new_num=$(printf "%03d" $((i - 10)))
  slug=$(basename "$old" | sed "s/kata-$(printf "%03d" $i)/kata-${new_num}/")
  git mv "$old" "katas/$slug"
done
```

**Commit 2c**: Move staging to final positions (031–040)
```bash
for f in katas/.staging/kata-*; do
  git mv "$f" "katas/$(basename "$f")"
done
rmdir katas/.staging
```

**Commit 2d**: Update `go.mod` module paths in all 20 renamed directories

**Verification**: `go test ./apps/learner-studio/... && ./scripts/test_operability.sh`

### Step 3: Tail Swap — DB/IO ↔ Bug-Fix (131–140)

Same staging approach as Step 2 but for 10 directories.

**Commit 3a**: Move kata-131 through kata-135 (bug-fix) to staging as kata-136 through kata-140
**Commit 3b**: Move kata-136 through kata-140 (DB/IO) to kata-131 through kata-135
**Commit 3c**: Move staging to final positions kata-136 through kata-140
**Commit 3d**: Update `go.mod` module paths in all 10 renamed directories

**Verification**: `go test ./apps/learner-studio/... && ./scripts/test_operability.sh`

### Step 4: Fix Broken Katas — Batch 1 (Concurrency, 092–093 + 100)

Three katas per batch to keep commits reviewable.

**Commit 4a**: Rewrite kata-092 (Context-Aware Logger)
- New signature: `func NewContextLogger() *ContextLogger`
- Proper README, kata.go stub, kata_test.go with real tests

**Commit 4b**: Rewrite kata-093 (Structured Error Types)
- New signature: `func NewAppError(code string, msg string, cause error) *AppError`
- Proper README, kata.go stub, kata_test.go with real tests

**Commit 4c**: Rewrite kata-100 (LRU Cache Thread-Safe)
- New signature: `func NewThreadSafeLRU(capacity int) *ThreadSafeLRU`
- Proper README, kata.go stub, kata_test.go with real tests

**Verification**: `go test ./katas/kata-092-*/... ./katas/kata-093-*/... ./katas/kata-100-*/...`

### Step 5: Fix Broken Katas — Batch 2 (Data Structures, 101–106)

**Commit 5a**: Rewrite kata-101 (Token Bucket Thread-Safe)
**Commit 5b**: Rewrite kata-102 (Bloom Filter)
**Commit 5c**: Rewrite kata-103 (Merkle Tree)
**Commit 5d**: Rewrite kata-104 (Priority Queue)
**Commit 5e**: Rewrite kata-105 (Dijkstra Shortest Path)
**Commit 5f**: Rewrite kata-106 (A* Pathfinding)

**Verification**: `go test ./katas/kata-10{1,2,3,4,5,6}-*/...`

### Step 6: Replace Academic Katas (107, 108, 111, 112)

**Commit 6a**: Replace kata-107 → Bug Fix: Goroutine Leak
- Rename directory, new go.mod, kata.go (pre-broken), kata_test.go, README.md

**Commit 6b**: Replace kata-108 → Bug Fix: Data Race
**Commit 6c**: Replace kata-111 → Bug Fix: Timeout Handling
**Commit 6d**: Replace kata-112 → Custom Test Helper

**Verification**: `go test ./katas/kata-107-*/... ./katas/kata-108-*/... ./katas/kata-111-*/... ./katas/kata-112-*/...`

### Step 7: Fix Remaining Broken Katas (109, 110, 125, 127, 128)

**Commit 7a**: Rewrite kata-109 (JSON Schema Validator)
**Commit 7b**: Rewrite kata-110 (Rate-Limited HTTP Client)
**Commit 7c**: Rewrite kata-125 (Concurrent Test Harness)
**Commit 7d**: Rewrite kata-127 (Generics: Set/Map Utilities)
**Commit 7e**: Rewrite kata-128 (Generics: Optional/Result)

**Verification**: `go test ./katas/kata-10{9,10}-*/... ./katas/kata-12{5,7,8}-*/...`

### Step 8: Replace Duplicate (052)

**Commit 8a**: Replace kata-052 → Table-Driven Test Refactor
- New directory name: `kata-052-table-driven-test-refactor`
- New go.mod, kata.go, kata_test.go, README.md

**Verification**: `go test ./katas/kata-052-*/...`

### Step 9: Progress Migration

**Commit 9a**: Add ID migration map to `internal/learning/progress`
- On load, translate old kata IDs to new ones
- Write migrated state back to file
- Log migration actions for debugging

**Verification**: Unit test with fixture progress.json containing old IDs

### Step 10: Full Verification

**Final check**:
```bash
./scripts/test_operability.sh
go test ./...
```

Verify:
- All 140 kata directories exist with correct numbering
- All `go.mod` files have correct module paths
- All `kata.go` files have proper function signatures (no `Kata##()` stubs)
- All `kata_test.go` files have real test functions (no `t.Skip` only)
- track.json categories cover 001–140 completely with no gaps
- pathways.json references only valid category IDs
- Progress store can load/migrate old format
- Studio app starts and serves all katas correctly

### Batching Summary

| Step | Commits | Katas Touched | Risk |
|------|---------|--------------|------|
| 1. Config | 3 | 0 | Low |
| 2. Block swap 021↔031 | 4 | 20 | **High** |
| 3. Tail swap 131↔136 | 4 | 10 | Medium |
| 4. Fix broken batch 1 | 3 | 3 | Low |
| 5. Fix broken batch 2 | 6 | 6 | Low |
| 6. Replace academic | 4 | 4 | Low |
| 7. Fix broken batch 3 | 5 | 5 | Low |
| 8. Replace duplicate | 1 | 1 | Low |
| 9. Progress migration | 1 | 0 | Medium |
| 10. Full verification | 0 | 0 | — |
| **Total** | **31** | **49** | |

### Rollback Strategy

Each step is independently revertable via `git revert`. The block swap (Step 2) is the only operation requiring multi-commit rollback — if it fails partway through, revert all commits in that step.

The staging directory approach means there is never a state where two katas have the same number simultaneously.
