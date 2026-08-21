# Defining a stage (phase)

A stage is a *state of competence*, not a collection of topics. This playbook defines one stage properly.

## Stage template

For the stage being designed, produce:

1. **Position**: Dreyfus stage + one line on where the learner is ("has completed ~30 guided katas; can read specs; still reaches for examples before principle").
2. **Bloom ceiling**: the highest cognitive level the stage's katas may require, with the verb set for that level.
3. **Posture**: shu (strict idiom, exact rules), ha (judgment within constraints), or ri (design your own constraints) — and where in the stage the posture shifts.
4. **Objectives** (3–5): each an outcome, each starting with a Bloom verb appropriate to the ceiling.
5. **Crossover inventory**: the cross-language concepts this stage plants or revisits, with Bloom level of this visit.
6. **Entry/exit test**: the concrete skill a learner proves to enter and to leave.

## Bloom verb sets by level

| Level | Verbs for objectives | What katas should demand |
|---|---|---|
| Remember | list, name, identify, recall | fill-in-the-blank, "what does X do" |
| Understand | explain, describe, summarize, compare | "explain the difference between X and Y" |
| Apply | implement, use, calculate, demonstrate | "write code that does X given spec" |
| Analyze | compare, contrast, debug, trace, refactor | "find the bug", "explain why this fails" |
| Evaluate | assess, review, critique, judge, trade off | "which approach and why", code review katas |
| Create | design, build, compose, architect, teach | design-a-small-system, "write the plan" katas |

A stage that says "Analyze" but whose katas only say "implement" is lying about its level. Fix the katas or lower the claim.

## Stage-specific guidance

### Junior / foundation
- **Posture**: strict shu. Enforce idiom (gofmt, naming, error returns) hard — this is the only stage where rote correctness is the point.
- **Crossover to plant**: the universal concepts learners will need everywhere: types & values, control flow, functions, data structures, testing, basic error handling.
- **Entry test**: writes "hello world" and runs it. **Exit test**: completes a 3-concept kata (e.g. parse input → transform → validate) with passing tests, alone.

### Mid / practitioner
- **Posture**: shu → ha. Idiom still enforced; multiple valid approaches now accepted with justification.
- **Crossover to revisit**: concurrency, error taxonomy, testing strategy, serialization — each now Apply/Analyze instead of Remember/Understand.
- **Entry test**: combines two junior concepts without help. **Exit test**: debugs an unfamiliar failure and refactors a working but messy solution.

### Senior
- **Posture**: ha. Rules exist to be broken with judgment; review, trade-off, and design-in-the-large.
- **Crossover to revisit**: performance, security, observability, architecture — the "make it fast, safe, observable, and shapeable" layer.
- **Entry test**: reviews a junior's PR with specific, prioritized feedback. **Exit test**: designs a small system's architecture and defends it against a stated constraint.

### Lead
- **Posture**: ri. The learner now sets conventions and teaches; katas are often review/teach/design rather than "write function".
- **Crossover to revisit**: economics, process, professional practice — "how do teams make good software and good decisions".
- **Entry test**: identifies a process or quality gap in a mock team. **Exit test**: designs the curriculum *itself* — writes a mini-track teaching a concept they just mastered.

## Revisiting vs. repeating

A crossover revisit must be a *new Bloom level in a new context*. "Another kata about slices" is repeating; "use slices to implement a bounded buffer under concurrency, then defend the design" is revisiting. If the revisit kata could be dropped into the earlier stage without anyone noticing, it's not a revisit — raise it.
