# Auditing a kata

Check one kata against the quality floor. Produce findings with severity and a concrete fix per finding.

## How to audit

1. Read the readme, starter code, tests, and metadata JSON.
2. Verify the tests actually run: pass on the reference implementation, fail on a broken one (run both in the sandbox when available).
3. Check the readme's claims against the tests — they must be the same spec.
4. Check the kata's place in the curriculum: stage/level fit, prerequisites, tags matching the concept map.

## Checklist

### Contract
- [ ] Exactly one focus; readme, signature, rules, and tests all serve it.
- [ ] Rules are `condition => expectation` with examples.
- [ ] Every rule is testable and every rule is encoded in a test.
- [ ] The readme's "Rules / Expectations" matches the tests verbatim (no drift).
- [ ] Definition of done is present and plain-English.

### Tests
- [ ] Tests pass on a correct implementation (run it).
- [ ] Tests fail on an incorrect implementation (run it).
- [ ] Edge rows exist: nil/empty, boundary, error case as applicable.
- [ ] No trivial passes (tests that don't actually exercise the contract).

### Starter code
- [ ] Compiles, exposes the signature, no solution hidden in comments.
- [ ] Sized to the concept (tens of lines, not hundreds).

### Readme
- [ ] Structure: task, learning goal (what/why/how/when-done), tips, rules, submission, run.
- [ ] Learning goal explains why this matters and how it grows the learner.
- [ ] Title/number is internally consistent (readme title matches kata ID — mismatches are a blocker).
- [ ] No copy-pasted boilerplate from another kata (an entire genre of low-quality katas).

### Metadata
- [ ] stage, category, level, tags, prerequisites, estimated_minutes present and honest.
- [ ] tags match the curriculum concept map (crossover threads), not just the function.
- [ ] prerequisites are minimal and correct.
- [ ] language declared; workspace filenames match the language.
- [ ] flashcards/quiz reinforce the concept (if the curriculum expects them).

### Evaluator
- [ ] `evaluator_status: ready` only if the sandbox verdict is verified trustworthy.
- [ ] `incomplete` is an honest label, not a failure — but a track with mostly-incomplete katas is a curriculum blocker (flag it to the curriculum audit).

## Severity guide

- **Blocker**: tests don't verify the contract (red-green unverified, trivial passes), or readme ID/title mismatch, or the kata can't be completed in the sandbox.
- **Major**: readme/tests drift (two specs), missing edge cases for a rule that has them, missing prerequisites, evaluator marked ready but unverified.
- **Minor**: missing metadata fields, one unclear rule, readme lacks a learning-goal section.
- **Info**: suggestions beyond correctness (better examples, a quiz question, a far-transfer variant).

## Output shape

```markdown
## Kata audit: <id> — <title>

**Focus:** <declared> · **Stage/level:** <declared> · **Evaluator:** <ready|incomplete>

### Findings
- [blocker] ...
- [major] ...

### Verdict
Pass / Pass with fixes / Rework — with the single highest-leverage fix.
```

The single highest-leverage fix for a bad kata is almost always: **make the tests the spec** — reconcile readme and tests, add the missing edge rows, and verify red-green in the sandbox.
