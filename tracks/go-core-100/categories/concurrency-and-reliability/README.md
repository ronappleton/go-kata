# Concurrency and Reliability

**Kata range:** `091-110`

## What you are building
You are learning to coordinate goroutines safely and keep behavior correct under timing variance.

## Why this matters in real Go work
Concurrency bugs are expensive and hard to reproduce. Strong ownership and cancellation patterns prevent production incidents.

## What "correct" looks like by the end
- You avoid races, leaks, and deadlocks in common patterns.
- You use context and synchronization primitives deliberately.
- You can explain lifecycle ownership: who starts, who stops, who waits.

## How to practice this category
- Test with repeated runs and edge timing where possible.
- Keep goroutine lifetimes explicit in code.
- Prefer predictable coordination over clever channel tricks.
