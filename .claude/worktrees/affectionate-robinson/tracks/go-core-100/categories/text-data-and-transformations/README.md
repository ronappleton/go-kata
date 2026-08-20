# Text, Data, and Transformations

**Kata range:** `051-062`

## What you are building
You are training robust parsing and normalization habits for text and structured payloads.

## Why this matters in real Go work
Most software consumes imperfect input. Your job is to parse strictly, fail safely, and return consistent output contracts.

## What "correct" looks like by the end
- You differentiate valid, invalid, and partial input paths clearly.
- You avoid silent data loss when transforming values.
- You can explain what happens when input format is malformed.

## How to practice this category
- Write tests for malformed and near-valid examples.
- Keep parser rules explicit rather than implicit.
- Return errors and fallback behavior exactly as the kata contract states.
