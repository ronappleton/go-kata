# File, Time, and Config

**Kata range:** `063-070`

## What you are building
You are learning to work safely with filesystem boundaries, clocks, and configuration inputs.

## Why this matters in real Go work
I/O and time are unstable boundaries. Reliable systems treat them as contracts with explicit behavior under failure.

## What "correct" looks like by the end
- You handle missing files, invalid config, and parsing failures intentionally.
- You keep time-based behavior deterministic in tests.
- You avoid hidden assumptions about environment and local machine state.

## How to practice this category
- Test both happy path and failure path for each boundary.
- Keep parsing and validation near input boundaries.
- Return actionable errors instead of generic failures.
