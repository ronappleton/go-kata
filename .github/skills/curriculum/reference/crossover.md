# Crossover knowledge

Crossover knowledge is the set of concepts and skills that transfer across languages, frameworks, and domains. It is the reason "knows Go" and "knows Java" should not be two unrelated curricula.

## Why crossover is a design requirement

Transfer of learning research:

- **Identical-elements theory** (Thorndike & Woodworth, 1901): transfer is strongest when the new situation shares elements with the learning situation. If a curriculum teaches "channels" only inside Go-specific syntax, the learner transfers only the syntax, not the idea.
- **Near vs. far transfer** (Perkins & Salomon): near transfer is applying a skill in a similar context; far transfer is applying it in a very different one. Both are trainable, but differently: near transfer needs repeated practice in varied contexts; far transfer needs *abstraction* — naming the underlying principle.
- **Low-road vs. high-road transfer**: low-road is automatic, reflex-like application (built by lots of practice); high-road is deliberate, mindful application (built by abstraction and self-explanation). A curriculum should deliberately cultivate both.

**Design implication**: every major concept should appear in the curriculum in at least two forms — once embedded in one language's idiom (low-road, near transfer), and once abstracted and re-applied in another context (high-road, far transfer). The learner who only ever sees Go will conclude every language is Go. The learner who sees the *concept* behind Go's channel, Rust's channel, and Java's BlockingQueue becomes language-portable.

## The crossover concept map

Design a small, stable set of concepts (typically 5–8) that recur across all target languages and domains:

| Concept | Go | Rust | Java/C# | Python | Appears at stage |
|---|---|---|---|---|---|
| Values & types | structs, zero values | ownership, Copy/Clone | classes, records | duck typing, dataclasses | junior |
| Control & flow | if/for/switch | match, loop | switch/streams | comprehensions | junior |
| Functions & composition | funcs, closures | fn, traits | lambdas, interfaces | first-class funcs, decorators | junior→mid |
| Error handling | error, errors.Is/As | Result, ? | exceptions | exceptions | junior→senior |
| Testing | table tests, t.Run | #[cfg(test)] | JUnit/xUnit | pytest, fixtures | junior→senior |
| Data & serialization | encoding/json | serde | Jackson/System.Text.Json | json/pydantic | mid |
| Concurrency | goroutines, channels | threads, channels, Send/Sync | virtual threads, futures | asyncio, GIL | mid→senior |
| Resource mgmt | defer | Drop/RAII | try-with-resources | context managers | mid |
| Composition & design | interfaces, embedding | traits, generics | interfaces, inheritance | ABCs, protocols | senior |
| Performance | pprof, benchmarks | cargo bench | JMH | cProfile | senior |
| Security | crypto stdlib | crate ecosystem | OWASP patterns | hardening | senior |
| Architecture | package layout | crates | modules, DI | project structure | senior→lead |

This table is a template, not a law: adapt the concept set to the target domain (add "protocol design" for networking, "state machines" for systems work, etc.).

## How to design crossover into the curriculum

1. **Name the concept in the kata.** Each kata's readme carries a "Crossover" line: `Crossover: concurrency (revisit, Analyze)`. The learner should see the thread.
2. **Spiral, don't silo.** Each concept appears ≥ 2× at increasing Bloom levels. Track the spiral in a table so the curriculum author can verify no concept is planted but never revisited.
3. **Abstract after the concrete.** The first encounter is always concrete and idiomatic (shu). The revisit deliberately abstracts: "the same idea in another form".
4. **Build far-transfer katas.** At mid+ stages, include katas that require transferring a concept to a new domain — e.g. "apply the error-handling taxonomy you designed to a CLI tool's parsing layer". These are the high-road exercises.
5. **Cross-language katas where the product supports it.** If the app supports multiple languages (see the `languages` registry), a crossover kata can offer the same exercise in two languages and ask the learner to compare — the comparison itself is the learning.

## Anti-patterns

- **Siloing**: a curriculum whose categories are pure language features ("Go Standard Library") with no concept threads across them.
- **Sneaky repetition**: revisiting a concept without raising the Bloom level (see phase.md).
- **Jargon-as-knowledge**: using cross-language terms without teaching the underlying idea ("it's like a monad" is not teaching).
- **No naming**: concepts that transfer but are never named, so the learner never abstracts them.
- **Crossover theater**: claiming crossover in the readme but the kata is identical to its earlier instance.

## Audit check

For an existing curriculum, verify: (1) the concept map exists or can be inferred; (2) each concept appears ≥ 2× at rising Bloom levels; (3) at least one far-transfer kata per major concept; (4) the katas' metadata (tags, prerequisites) actually encodes the threads — tags must match the concept map, not just describe the function.
