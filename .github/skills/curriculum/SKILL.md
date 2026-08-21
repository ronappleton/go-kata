---
name: curriculum
description: Use when the user wants to design, revise, audit, or sequence a learning curriculum or track — the stage-by-stage progression a learner follows from beginner to expert. Covers competency modeling (junior/mid/senior/lead or similar phases), stage and category design, learning objectives, prerequisite ordering, crossover knowledge across languages and disciplines, difficulty calibration, spiral/revisit design, and quality audits of an existing track. Language-agnostic: works for Go, Rust, Java, C#, C++, Python, and any other language or domain. Not for writing a single exercise (use the katas skill for that).
version: 1.0.0
user-invocable: true
argument-hint: "[design|audit|phase|crossover] [target]"
license: Apache 2.0
---

A curriculum is a map of how a learner's competence changes over time. This skill gives you the research-backed models to design that map well — and to audit one that exists.

## Core principles

- **Competence is a progression, not a list.** The Dreyfus model (novice → advanced beginner → competent → proficient → expert) describes how learners shift from following rules, to recognizing situations, to choosing goals, to acting intuitively. A curriculum must move learners through those shifts, not just add facts. Map each stage to where the learner *is*, not what topics exist.
- **Bloom's taxonomy sets the cognitive bar per stage.** Remember → Understand → Apply → Analyze → Evaluate → Create. Junior work is mostly Remember/Understand/Apply; senior work is Analyze/Evaluate; lead work is Evaluate/Create. Every stage's objectives should use verbs at the right Bloom level, and katas within the stage should exercise at least the claimed level.
- **Shu-ha-ri keeps the discipline honest.** Shu: follow the rules exactly. Ha: break the rules with understanding. Ri: transcend the rules. A stage that never leaves Shu produces rule-followers, not engineers; a stage that leaves Shu too early produces cargo-culters. The junior stage is Shu-heavy; the lead stage is Ri-heavy.
- **Crossover knowledge is real and must be planned.** Transfer of learning research (Thorndike & Woodworth's identical-elements theory; Perkins & Salomon's near/far and low-road/high-road transfer) shows skills transfer between contexts when the learner abstracts the underlying principle. A good curriculum names the cross-language concepts (concurrency, error handling, testing, memory/ownership, composition, protocol design) and deliberately revisits them in each language so near-transfer becomes automatic and far-transfer is practiced. Never silo a curriculum into "language X only" — that leaves the learner unable to transfer to language Y or to a new domain.
- **Depth beats breadth at every stage.** A stage with a handful of thoroughly-practiced concepts beats a sprawling list of lightly-touched topics. Each stage should have a small set of explicit learning objectives; every category and kata in that stage serves one of them.
- **The body of knowledge is language-independent.** SWEBOK's knowledge areas (requirements, architecture, design, construction, testing, operations, maintenance, configuration management, process, quality, security, professional practice, economics, foundations) are the canonical map of what software professionals know. Map your curriculum's stages onto them so no major area silently vanishes.

## Research grounding

- **Dreyfus model of skill acquisition** (Dreyfus & Dreyfus, 1980/1986): five stages from rule-following to intuition. The anchor for stage design.
- **Bloom's taxonomy, revised** (Anderson & Krathwohl, 2001): Remember → Understand → Apply → Analyze → Evaluate → Create. The anchor for learning-objective verbs and assessment depth.
- **Shu-ha-ri** (Japanese martial arts, via Sen no Rikyū/Zeami): follow → break → transcend. The anchor for how strictly a stage enforces idiom and convention.
- **SWEBOK V4** (IEEE/ISO, 2024): 18 knowledge areas of software engineering. The anchor for coverage auditing.
- **Transfer of learning** (Thorndike & Woodworth, 1901; Perkins & Salomon): identical elements, near/far transfer, low-road/high-road transfer. The anchor for crossover knowledge.
- **Apprenticeship Patterns** (Hoover & Oshineye): journeyman→master progression, "sweep the floor," "expose your ignorance," "reading list." The anchor for stage culture and learning habits.
- **Deliberate practice** (Ericsson; also Yegge's "Practicing Programming"): focused, repeated, feedback-driven practice is what builds skill — not passive work experience. The anchor for why a curriculum needs *exercises with verification*, not just reading lists.

## Workflow

1. **Anchor the stages.** Before touching topics, decide the stage model (e.g. junior/mid/senior/lead, or novice/competent/expert). For each stage, write: who the learner is (Dreyfus), what cognitive level they operate at (Bloom), how strictly they follow idiom (shu-ha-ri), and the 3–5 headline objectives.
2. **Map coverage.** Lay the SWEBOK knowledge areas over the stages. Confirm every area that matters for the target role is *somewhere*, and that each area appears at the right cognitive depth for its stage.
3. **Design categories.** Each stage = 3–8 categories. Each category has one learning goal and a handful of katas that serve it. A category is a coherent *capability*, not a topic heading.
4. **Plan crossover.** Decide the 5–8 cross-language/cross-domain concepts. Ensure they're revisited at least twice across the curriculum (spiral), each revisit at a higher Bloom level. Name them explicitly so the learner knows the transfer is intended.
5. **Sequence with prerequisites.** Every kata that needs prior knowledge declares it. Ordering follows the dependency graph, not alphabetical or topical convenience.
6. **Calibrate difficulty.** Within a stage: entry katas are one-concept warm-ups; mid katas combine two concepts; exit katas combine three or more. The stage's exit kata should feel like a small real-world task.
7. **Audit.** Run the audit checklist (see reference/audit.md) against the design or an existing track.

## Commands

| Command | Purpose | Reference |
|---|---|---|
| `design [domain]` | Build a new curriculum from scratch | [reference/design.md](reference/design.md) |
| `phase [stage]` | Define or revise a single stage's objectives and shape | [reference/phase.md](reference/phase.md) |
| `crossover` | Design the cross-language/cross-domain concept map | [reference/crossover.md](reference/crossover.md) |
| `audit [track]` | Quality-check an existing curriculum against the model | [reference/audit.md](reference/audit.md) |

## Quality bar

A curriculum earns the label "well-designed" when:

- Each stage names the learner's Dreyfus position, Bloom ceiling, and shu-ha-ri posture — and the katas actually operate at that level.
- Every category has a measurable learning goal and every kata serves exactly one.
- Crossover knowledge is explicitly planned, not accidental, and is spiraled across stages.
- The prerequisite graph is acyclic and every kata's dependencies are declared.
- The exit point of each stage prepares the learner for the entry point of the next.
- SWEBOK coverage is deliberate: areas are either present at the right depth or consciously deferred with a reason.
- Any learner who completes the track can articulate what changed about their *judgment*, not just what they learned to type.
