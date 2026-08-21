# Curriculum & Kata Quality Audit — go-core-100 (August 2026)

Audited against the researched standards encoded in the `curriculum` and `katas` skills: Dreyfus skill model, revised Bloom's taxonomy, shu-ha-ri, SWEBOK V4 knowledge areas, transfer-of-learning (crossover) research, and the code-kata canon (Dave Thomas, Uncle Bob's Programming Dojo, Yegge's deliberate practice).

Source of truth: live content repo (`gokatas-content`), 175 katas, all fetched and verified.

---

## Coverage summary

| Metric | Value |
|---|---|
| Katas in track | 175 |
| Stages | 4 (foundation/junior, practitioner/mid, senior, lead) |
| Categories | 18 |
| Evaluators `ready` | **42 (24%)** |
| Evaluators `incomplete` | 133 (76%) |
| Katas with rich metadata (tags, prerequisites, level, flashcards, quiz, estimated_minutes) | **12 of 175 (7%)** |
| Katas whose README title number mismatches the kata ID | **~116 (66%)** |

---

## Findings

### Blockers

1. **A track where 76% of katas are `incomplete` cannot be completed.** The evaluator being `ready` is the difference between a kata and a suggestion (per the katas skill: an exercise without verification is a reading list). A learner who follows the track linearly will hit an `incomplete` kata almost immediately. This is the single highest-leverage fix: the curriculum's promise ("Junior to Lead") is unfulfillable until the evaluators exist.

2. **~116 of 175 katas have a README title number that doesn't match their ID** (e.g. kata `006`'s README says "# Kata 031", kata `027` says "# Kata 018", kata `058` says "# Kata 060"). This is systematic copy-paste drift from a renumbering pass. It destroys learner trust ("am I on 006 or 031?"), breaks search/continuity, and signals the content was bulk-generated rather than crafted. Fix is mechanical but must be done for every kata.

### Major

3. **93% of katas lack the rich metadata the app already supports.** Only katas 163–174 carry `tags`, `prerequisites`, `level`, `estimated_minutes`, `flashcards`, and `quiz_questions`. The first 162 katas have only id/slug/title/focus/signature/rules/evaluator_status. Consequences:
   - **Prerequisite graph is nearly empty** — the curriculum skill requires acyclic, declared dependencies; ours can't sequence correctly, and the app's linear-lock feature can't enforce sensible ordering for most katas.
   - **Crossover/spiral is invisible** — tags are the mechanism for naming cross-stage concept threads (concurrency, error handling, testing). With tags on 7% of katas, the curriculum has no way to plan or audit near/far transfer. The `crossover` reference's audit check fails wholesale.
   - **Flashcards/quiz don't exist for 93%** — the app's flashcard and quiz learning modes have almost nothing to draw from.

4. **Category quality is inconsistent.** The first stages read as capabilities ("Language Basics", "Concurrency Patterns") — good. But some senior/lead categories are thin or topic-like, and the distribution is lopsided (foundation: 58 katas, senior: 25, lead: 21) with no stated rationale. Per the curriculum skill, stage size should reflect cognitive lift, and each category needs an outcome-shaped learning goal — those goals exist in the track.json but aren't echoed in kata metadata (stage/category fields are missing from 93% of kata JSON).

5. **Crossover knowledge is not planned, it's accidental.** No concept map exists in the content. Concurrency, error handling, testing, serialization appear in category titles, but nothing ties e.g. "concurrency" across foundation → practitioner → senior at rising Bloom levels. The app now supports multiple languages (see `internal/languages`) — a genuine opportunity to build far-transfer katas — but content is 100% Go-only and single-context.

### Minor

6. **README learning goals drift across a template upgrade.** Katas 000–026 and 027+ use different "Learning goal" templates; both are fine individually, but the earlier ones are shallower ("what you are practicing") vs later ("definition of done in plain English"). A bulk pass should normalize all to the later standard (which includes "Definition of done").

7. **A few READMEs have empty/truncated rule lines** (e.g. kata 058's rules end with `- ` and an empty bullet). Minor but visible in the rendered docs view.

8. **`focus` sometimes over-promises** — "Two-pointer merge, slice appends" is two concepts for one kata (kata 027); the katas skill's one-concept rule says split or re-focus.

### Info

9. **Bloom level is implied, never stated.** No kata declares its cognitive level in metadata beyond junior/mid/senior/lead (12 katas do). Adding a `bloom` field or encoding it in level verbs would let the curriculum audit verify stage claims automatically.

10. **No far-transfer or cross-language katas exist yet.** With the language registry now live, this is the cheapest high-value addition: the same concept in two languages with a compare-and-contrast question.

---

## Verdict

**The track is a well-named topic list, not yet a curriculum.** It has the right skeleton (4 stages, sensible categories, a dependency graph that *could* be acyclic), but 76% incomplete evaluators, 93% missing metadata, and 66% ID mismatches mean it cannot currently be completed, sequenced, or audited. It fails the curriculum skill's core test: *a learner who completes it should be able to say what changed about their judgment, not just what they learned to type* — because most of it can't be completed.

**The single highest-leverage fix: make the track completable** — finish evaluators for the foundation stage first (the 42 `ready` are mostly later katas 163–174; the early katas learners meet first are largely incomplete). Then, in order: fix README ID mismatches, backfill metadata (prerequisites + tags at minimum) using the crossover concept map, and use the now-live language registry to add the first far-transfer katas.

## Suggested remediation order

1. **Completability**: build/verify evaluators for foundation-stage katas (000–162 backlog) — blockers from audit item 1.
2. **Correctness**: mechanical pass to fix README number drift (item 2).
3. **Metadata backfill**: prerequisites + tags for all katas, guided by a documented crossover concept map (item 3, 5).
4. **Spiral**: name concept threads in readmes and tags at rising Bloom levels (item 9).
5. **Crossover**: first far-transfer / multi-language katas via the language registry (item 10).
