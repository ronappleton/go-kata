# Designing a curriculum

Build the map from first principles. Do not start from a topic list — start from the learner's trajectory.

## 1. Choose the stage model

The most common mapping for professional software tracks:

| Stage | Dreyfus | Bloom ceiling | Shu-ha-ri | Headline change |
|---|---|---|---|---|
| Junior (foundation) | Novice → Advanced Beginner | Apply | Shu (follow exactly) | Can write correct, idiomatic small programs from a clear spec |
| Mid (practitioner) | Advanced Beginner → Competent | Analyze | Shu → Ha | Can choose approaches, combine concepts, and debug systematically |
| Senior | Competent → Proficient | Evaluate | Ha (break with judgment) | Can design, review, trade off, and own non-trivial systems |
| Lead | Proficient → Expert/Master | Create | Ri (transcend) | Can set direction, teach, and raise the team's bar |

If the target is narrower (e.g. "learn Go for backend work"), compress to three stages — but keep the *progression of judgment* visible even in two stages. The stage is not the language level; it is the learner's relationship to the material.

## 2. Write stage anchors before anything else

For each stage, a short block that answers:

- **Who**: the learner entering this stage (prior knowledge, what they can already do).
- **Cognitive bar**: the Bloom verbs this stage's katas must exercise.
- **Posture**: how strictly idiom is enforced (shu) vs. how much latitude is allowed (ha).
- **3–5 objectives**: measurable, outcome-shaped ("can write a table-driven test suite", not "knows testing").

These anchors are the acceptance criteria for the whole stage. A kata that doesn't serve an anchor verb doesn't belong in the stage.

## 3. Design categories as capabilities

Each stage has 3–8 categories. A category is a *capability*: "Error Handling", "Concurrency", "Observability". Not "Useful Packages" — that's a topic heading with no learning shape.

For each category write:

- **Learning goal**: one sentence, outcome-shaped, at the stage's Bloom level.
- **Description**: what the learner will be able to do, in the learner's voice.
- **Kata list**: 3–12 katas, ordered entry → mid → exit.
- **Crossover markers**: which cross-language concept this category is an instance of (so the spiral is visible).

## 4. Order by dependency, not by index

- Every kata declares prerequisites (kata IDs). The track's ordering is the transitive closure of the prerequisite graph — acyclic by construction.
- Category order within a stage follows capability need: setup/toolchain before language basics, language basics before concurrency, etc.
- Stage boundaries are crossings: the exit katas of stage N must feed the entry katas of stage N+1.

## 5. Calibrate difficulty within a stage

- **Entry katas**: one concept, one function, spec fully given. Bloom: Remember/Apply. Time: minutes.
- **Mid katas**: two concepts combined (e.g. "error handling + file I/O"). Bloom: Apply/Analyze. The spec may leave one decision open.
- **Exit katas**: three or more concepts; the spec is a small real-world task, possibly with deliberate ambiguity ("handle the edge cases a reviewer would flag"). Bloom: Analyze/Evaluate.

A stage whose katas are all the same size has no ramp. A stage whose entry kata is as hard as its exit kata has no pedagogy.

## 6. Plan the spiral

Concepts that matter are revisited across stages, each time at a higher Bloom level and in a new context:

- Concurrency: junior = "what is a goroutine"; mid = "coordinate with channels"; senior = "design for contention, reason about deadlock".
- Error handling: junior = "return and check errors"; mid = "wrap with context"; senior = "design an error taxonomy for an API".
- Testing: junior = "write a table-driven test"; mid = "write tests first"; senior = "design the test surface of a package".

Name the concept explicitly in each kata's readme ("this is a revisit of X from stage N at a higher level"). Naming is what converts near-transfer into far-transfer (Perkins & Salomon's high-road transfer).

## 7. Coverage map

Draft a coverage table: SWEBOK areas × stages, marking each cell P (present), R (revisited), or — (deferred, with reason). Then design categories to fill the gaps. A junior-to-lead track that never touches security, operations, or economics has holes — say why they're holes or fill them.

## Output shape

A track definition with:

```json
{
  "id": "track-id",
  "title": "Track Title",
  "description": "one-line promise",
  "stages": [
    {
      "id": "foundation", "title": "Foundation", "level": "junior",
      "description": "stage anchor in the learner's voice",
      "objectives": ["...", "..."],
      "categories": [
        {
          "id": "cat-id", "title": "Category Title",
          "description": "...", "learning_goal": "...",
          "kata_ids": ["000", "001", "..."]
        }
      ]
    }
  ]
}
```

Each kata in `kata_ids` is then designed by the **katas** skill with its own metadata (prerequisites, tags, level, flashcards, quiz) — the track declares the map; the kata declares the terrain.
