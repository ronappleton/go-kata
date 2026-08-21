# Writing a kata spec

The spec is the kata's contract: signature, rules, and the tests that encode them. This playbook is for writing or rewriting one rigorously.

## Rules-first

Write the rules before the signature, before the tests, before the readme. The rules are the source of truth; everything else encodes them.

### Rule anatomy

Each rule is one behavior. The template:

```
<condition> => <expected behavior>   # example
```

- Condition: the input/state the rule applies to.
- Expected behavior: the observable outcome.
- Example: a concrete instance the learner can test mentally.

### Rule checklist

- [ ] One behavior per rule (a rule with "and" is two rules).
- [ ] The normal case is covered.
- [ ] Edge cases are covered: nil/empty, boundaries (min/max, zero), duplicates, the "obvious but wrong" case.
- [ ] If the concept includes errors, the error behavior is a rule, not a footnote.
- [ ] Rules are consistent (no two rules contradict; no rule is implied by another in a way that makes one redundant).
- [ ] Every rule is testable — a human reading it can write the test.

## Signature discipline

- Minimal parameters. If a parameter exists only to make a test pass, remove it.
- Explicit returns. Prefer returning values over mutating inputs; if mutation is the point, say so in a rule.
- Language-idiomatic naming (Go: `MergeAlternating`; Rust: snake_case functions; Java/C#: PascalCase methods).

## Tests as executable rules

For each rule, the test is the proof:

```go
// Go
{"normal interleave", []int{1, 2}, []int{3, 4}, []int{1, 3, 2, 4}},
{"a empty", nil, []int{3, 4}, []int{3, 4}},
{"both empty", nil, nil, nil},
```

```python
@pytest.mark.parametrize("a,b,want", [
    ([1, 2], [3, 4], [1, 3, 2, 4]),
    ([], [3, 4], [3, 4]),
    ([], [], []),
])
def test_merge_alternating(a, b, want): ...
```

### Test quality gates

- **Red-green verified**: the suite fails on a deliberately wrong implementation and passes on the reference implementation. Run both; do not ship on assumption.
- **No trivial passes**: a test that passes without exercising the function (wrong name, wrong signature, empty body that happens to match) is a trap.
- **Edge rows present**: nil/empty/boundary rows exist for every rule that has them, even when "obvious".
- **Deterministic**: no time, randomness, or environment dependence in the assertions.

## The readme's rules section

The readme's "Rules / Expectations" must be the same rules the tests encode — copy them verbatim from the spec, not paraphrased. If a rule lives only in the tests or only in the readme, the kata has two specs, and the learner will learn the wrong one.

## Rewriting a bad spec

When a kata's spec is weak (ambiguous rules, untestable claims, tests that don't match rules):

1. Extract every claim from the readme and the tests into rule form.
2. Deduplicate and reconcile (pick the readme's intent over the test's accident, or vice versa — decide which is right, then make both match).
3. Rewrite rules as `condition => expectation` with examples.
4. Regenerate tests from rules.
5. Verify red-green in the sandbox.
6. Update the readme's rules section verbatim.

Never ship a kata whose spec is "whatever the tests happen to assert." The learner deserves a contract, not a rumor.
