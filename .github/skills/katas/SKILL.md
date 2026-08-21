---
name: katas
description: Use when the user wants to design, write, revise, or audit a single kata (coding exercise) — the task, spec, starter code, tests, and metadata. Covers exercise design from the code-kata canon, one-concept-per-exercise discipline, test-first specifications, rules and examples, difficulty calibration, evaluator readiness, and quality audits of existing katas. Language-agnostic: works for Go, Rust, Java, C#, C++, Python, and any other language, including multi-language crossover exercises. Not for designing a whole track (use the curriculum skill for that).
version: 1.0.0
user-invocable: true
argument-hint: "[design|spec|audit] [target]"
license: Apache 2.0
---

A kata is a small, focused exercise practiced deliberately until the skill becomes reflex. This skill gives you the research-backed discipline for writing one that actually teaches.

## Core principles

- **One kata, one concept.** The kata canon (Dave Thomas, *The Pragmatic Programmer*, 1999; Robert C. Martin's "The Programming Dojo"; the software craftsmanship movement) treats a kata as a single, repeatable form drilled to build muscle memory. If an exercise teaches two unrelated concepts, it is two katas. The *focus* field must name exactly one concept.
- **The spec is the teaching.** A kata is a contract: signature, rules, examples, definition of done. The learner's job is to satisfy the contract. If the rules are ambiguous, the learner learns ambiguity — not the concept. Write rules as explicit, testable behaviors with examples.
- **Tests come first, always.** The learner should be able to verify progress without a human: starter code + tests + a runnable sandbox. An exercise without verification is a reading list, not a kata. The evaluator being `ready` is the difference between a kata and a suggestion.
- **Difficulty is deliberate practice, not a puzzle.** Yegge's "Practicing Programming" and Ericsson's deliberate-practice research agree: the exercise should be at the edge of the learner's ability — hard enough to require focus, easy enough to complete in one sitting. Katas are not interview puzzles or riddles; they are drills with a clear correct answer.
- **Every kata declares its place.** Stage, category, level, prerequisites, tags, estimated minutes. A kata without prerequisites either assumes too much (learners get stuck) or too little (learners get bored). Metadata is how the curriculum's spiral becomes visible (see the curriculum skill's crossover playbook).
- **The readme is the lesson.** It should explain *what* you're practicing, *why it matters*, *how it grows you*, and *when you're done*. The "rules / expectations" section is the source of truth the tests encode.

## Research grounding

- **The Programming Dojo** (Robert C. Martin, 2006) and **CodeKata** (Dave Thomas): the origin of the term; small repeatable exercises for deliberate practice; solo practice of form.
- **Practicing Programming** (Steve Yegge): programmers don't practice; drills in focused solitude beat passive experience; solve the problem yourself, don't watch the solution.
- **Deliberate practice** (Ericsson): practice must be focused, feedback-driven, and just beyond current ability to build expertise.
- **Transfer of learning** (Thorndike & Woodworth; Perkins & Salomon): a kata embeds a concept in one language's idiom; the curriculum's spiral abstracts it later. The kata should tag the concept so the transfer is visible.
- **Bloom's taxonomy** (Anderson & Krathwohl, 2001): the kata's verbs must match its stage's cognitive level (see curriculum skill, phase.md).
- **Test-driven development canon** (Beck; Kent Beck's red-green-refactor): write the failing test first, encode the rule as a test, implement to pass.

## Workflow

1. **Anchor the concept.** One focus, at the right Bloom level for the kata's stage. Write the learning goal as an outcome.
2. **Write the signature.** The exact function/method/entry the learner implements. Keep parameter and return contracts minimal and explicit.
3. **Write the rules.** Each rule is one testable behavior, with examples. Rules are the source of truth.
4. **Write the tests.** Table-driven or equivalent: each test encodes one rule, including edge cases (nil, empty, boundary, error). Tests must pass on a correct implementation and fail on an incorrect one — verify both.
5. **Write the starter code.** Enough shape to compile and signal the contract, with TODO markers. Never a solution hiding in comments.
6. **Write the readme.** Task, focus, learning goal (what/why/how/when-done), tips, rules, submission requirements.
7. **Declare metadata.** stage, category, level, tags, prerequisites, estimated_minutes, flashcards, quiz questions, language.
8. **Verify evaluator readiness.** The kata is `ready` only when the sandbox can run learner code against the trusted tests and produce a verdict. Mark it `incomplete` otherwise — an unverified kata is honestly labeled.

## Commands

| Command | Purpose | Reference |
|---|---|---|
| `design [concept]` | Create a kata from a concept | [reference/design.md](reference/design.md) |
| `spec [kata]` | Write/rewrite the contract: signature, rules, tests | [reference/spec.md](reference/spec.md) |
| `audit [kata]` | Quality-check an existing kata | [reference/audit.md](reference/audit.md) |

## Quality bar

A kata earns the label "well-designed" when:

- It has exactly one focus, and the readme, signature, rules, and tests all serve it.
- Every rule is testable, and the tests encode every rule including edge cases.
- A correct implementation passes; an incorrect one fails (verified, not assumed).
- The readme explains what/why/how/when-done in plain language.
- Its metadata (stage, level, tags, prerequisites, estimated minutes) is complete and honest.
- Its evaluator is `ready` — or it is explicitly marked incomplete.
- A learner who completes it can say what *skill* they drilled, not just what they typed.
