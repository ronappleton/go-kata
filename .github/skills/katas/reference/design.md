# Designing a kata

Turn one concept into a complete, verifiable exercise.

## 1. Anchor the concept

- **Focus**: exactly one concept ("two-pointer merge", "error wrapping with context", "nil vs empty slices"). If you catch yourself writing two concepts joined by "and", split the kata.
- **Stage fit**: which curriculum stage/level does it serve? What Bloom verb does it demand? (Apply for junior, Analyze for mid, Evaluate for senior.)
- **Learning goal**: outcome-shaped, learner's voice — "write a function that satisfies N rules with explicit edge-case handling", not "learn about slices".

## 2. Write the contract

### Signature
The exact declaration the learner must implement. Minimal parameters; explicit return contract. For languages with files, this is the starter file's exported symbol.

### Rules
Each rule is one behavior, written as `condition → expectation` with an example:

```
- simple module => "example.com/myapp"
- with path => "github.com/user/repo"
- empty name after trim => "Hello, stranger!"
```

Rules must be:
- **Testable**: each maps to at least one assertion.
- **Exhaustive for the concept**: cover the normal case, edge cases (nil, empty, boundary), and the error case if the concept includes errors.
- **Unambiguous**: two readers agree on the same expectation.

### Definition of done
One plain-English sentence: "A reviewer can confirm this behavior from the tests: input X produces Y; empty input produces Z."

## 3. Write the tests

- Table-driven where the language supports it (Go: `[]struct{name, input, want}`; xUnit: parametrized; pytest: parametrize).
- **One test row per rule**, plus rows for the edge cases the rules imply even if not stated ("nil input", "empty string", "max boundary").
- **Tests are the spec.** A learner reading only the tests should reconstruct the rules.
- **Verify both directions**: tests pass on a correct implementation AND fail on an incorrect one (the classic trap is a test that passes trivially, e.g. a test that only checks the function exists).

## 4. Write the starter code

- Compiles, exposes the signature, and signals the contract with TODO markers.
- Comments state the signature and the intent — never the solution.
- Sized to the concept: ~5–20 lines. A 200-line starter is a codebase, not a kata.

## 5. Write the readme

Follow this structure (the app renders it as markdown):

```markdown
# Kata <ID> — <Title>

**Focus:** <one concept>

## Your task
Implement: `<signature>`

### Learning goal
- What you are practicing: ...
- Why this matters: ...
- How this grows your skills: ...
- When correct: ...

### Tips
- ...

## Rules / Expectations
- rule 1 with example
- rule 2 with example

## What you must submit for marking
- `kata.<ext>`
- `kata_test.<ext>`

## Run
<how to run the tests>
```

### Crossover line (mid+ stages)
For revisit katas, name the concept thread: `Crossover: concurrency (revisit, Analyze)`. This is what makes the curriculum's spiral visible (curriculum skill, crossover.md).

## 6. Metadata

Every kata declares:

```json
{
  "id": "NNN",
  "slug": "kebab-case-name",
  "title": "Human Title",
  "focus": "one concept",
  "signature": "func Name(args) result",
  "rules": ["rule 1", "rule 2"],
  "evaluator_status": "ready" | "incomplete",
  "stage": "foundation", "category": "cat-id", "level": "junior",
  "tags": ["concept-tags", "matching", "the", "concept-map"],
  "prerequisites": ["NNN", "..."],
  "estimated_minutes": 20,
  "language": "go",
  "flashcards": [...], "quiz_questions": [...]
}
```

- **tags** match the curriculum's crossover concept map — never just the surface function.
- **prerequisites** are the minimal set; a kata that needs prior knowledge and declares none breaks the learner.
- **estimated_minutes** is honest (deliberate practice is 15–60 min; if it's 2 hours it's not a kata, it's a project).
- **flashcards/quiz** reinforce the concept's vocabulary and recall — they support the spiral.

## 7. Evaluator readiness

`ready` means: the trusted tests are written, reviewed, and runnable in the sandbox, and the verdict is trustworthy. `incomplete` means any of that is missing. Never mark `ready` on faith — an unverified evaluator is a blocker (the learner can't complete the kata).
