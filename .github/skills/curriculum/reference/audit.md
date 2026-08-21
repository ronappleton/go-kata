# Auditing a curriculum

Audit an existing track against the model. Produce a findings report with severity (blocker / major / minor / info) and concrete fixes — not just a score.

## How to audit

1. Read the track definition (stages, categories, kata IDs, levels, descriptions).
2. Sample katas across every stage — at minimum the entry, middle, and exit kata of each stage, plus any kata whose metadata looks anomalous. When a content provider is available, pull the live content (manifest → track → katas) rather than trusting a stale local copy.
3. Check claims against evidence: read the actual kata readmes and tests, not just the titles.
4. Score each checklist item below as pass / fail / partial with a one-line reason.

## Checklist

### A. Stage model
- [ ] Each stage names its Dreyfus position, Bloom ceiling, and shu-ha-ri posture (or the equivalent is inferable and consistent).
- [ ] Stages are ordered by competence, and the exit of one stage feeds the entry of the next.
- [ ] Stage sizes are proportionate to their cognitive lift (a "senior" stage smaller than the "junior" stage is suspicious unless justified).

### B. Objectives and Bloom level
- [ ] Every category has an outcome-shaped learning goal, not a topic heading.
- [ ] Kata verbs match the stage's claimed Bloom ceiling (junior katas shouldn't demand evaluation; senior katas shouldn't be pure recall).
- [ ] Katas within a stage ramp: entry < mid < exit difficulty.

### C. Crossover and spiral
- [ ] A concept map is inferable from categories and tags (or explicitly documented).
- [ ] Major concepts appear ≥ 2× at rising Bloom levels across stages.
- [ ] At least one far-transfer kata per major concept exists.
- [ ] Kata tags match the concept map, not just the surface function.

### D. Prerequisites and ordering
- [ ] The prerequisite graph is acyclic (no kata transitively requires itself).
- [ ] Every kata that needs prior knowledge declares it.
- [ ] Ordering in the track follows the dependency graph.

### E. SWEBOK coverage
- [ ] Map categories to SWEBOK knowledge areas. Any area material to the target role that is absent is either consciously deferred (with reason) or a gap.
- [ ] Security, testing, quality, and operations appear at *some* stage at a depth appropriate to the role.

### F. Metadata completeness
- [ ] Every kata has: title, focus, signature, rules, evaluator status, stage, category, level, tags, prerequisites, estimated minutes.
- [ ] "Ready" evaluator coverage is high enough that the advertised track is actually completable (a track where most katas are `incomplete` is a track that can't be finished — flag as a blocker).
- [ ] Titles and numbers are internally consistent (readme title matches the kata's ID — mismatches are a blocker because they destroy learner trust and searchability).

### G. Kata-level quality (summary pass)
Spot-check readmes: do they follow the kata skill's quality floor (task, learning goal, rules, tips, definition of done)? One systematic miss (e.g. copy-pasted boilerplate) is a major finding, not a minor one.

## Severity guide

- **Blocker**: the track can't be completed as advertised (most evaluators incomplete), or a systematic defect like mismatched IDs/titles across the majority of katas.
- **Major**: a whole stage or category missing its learning shape; crossover never named; prerequisite graph broken.
- **Minor**: individual kata metadata gaps; one kata's readme boilerplate.
- **Info**: suggestions beyond correctness (spiral could go deeper, far-transfer katas could be added).

## Output shape

A findings report:

```markdown
## Curriculum audit: <track>

**Coverage**: X stages · Y categories · Z katas · W with ready evaluators · V with full metadata

### Blockers
- ...

### Major
- ...

### Minor
- ...

### Info
- ...

### Summary
One paragraph: is this track a curriculum (a progression of competence) or a topic list? What's the single highest-leverage fix?
```

The single highest-leverage fix is almost always: **make the spiral visible** (name cross-stage concepts in metadata and readmes) or **complete the evaluators** (a track of incomplete katas teaches nothing because it can't be verified).
