package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/runner"
)

func (s *studioServer) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req runRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	if req.TimeoutSeconds <= 0 {
		req.TimeoutSeconds = defaultRunTimeoutSeconds
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if req.SaveBeforeRun {
		if req.Code != nil {
			if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), []byte(*req.Code), 0o644); err != nil {
				writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
				return
			}
		}
		if req.Tests != nil {
			if err := os.WriteFile(filepath.Join(kata.Dir, "kata_test.go"), []byte(*req.Tests), 0o644); err != nil {
				writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata_test.go: %v", err))
				return
			}
		}
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(req.TimeoutSeconds)*time.Second)
	defer cancel()

	start := time.Now()
	runResult, _ := runner.RunKataTests(ctx, kata.Dir)
	duration := runResult.Elapsed
	if duration <= 0 {
		duration = time.Since(start)
	}

	attemptedAt := time.Now().UTC()
	outputTail := tailLines(runResult.RawOutput, defaultOutputTailLines)
	failureInsights := extractFailureInsights(outputTail)
	state, err := s.store.RecordAttempt(kata.ID, progress.AttemptResult{
		Passed:      runResult.Passed,
		Duration:    duration,
		FailedTests: runResult.FailedTests,
		OutputTail:  outputTail,
		RanAt:       attemptedAt,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("record progress: %v", err))
		return
	}

	resp := runResponse{
		Passed:          runResult.Passed,
		CompileErr:      runResult.CompileErr,
		DurationMS:      duration.Milliseconds(),
		FailedTests:     append([]string(nil), runResult.FailedTests...),
		OutputTail:      outputTail,
		FailureInsights: failureInsights,
		CoachHint:       coachHint(runResult.Passed, runResult.FailedTests, failureInsights),
		NextRecommended: s.nextRecommendation(state),
		Progress:        state.Attempts[kata.ID],
	}
	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleFormat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req formatRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, _, ok := s.track.FindKata(req.KataID); !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
	defer cancel()

	code := derefOrEmpty(req.Code)
	tests := derefOrEmpty(req.Tests)

	formattedCode, codeErr := formatGoSource(ctx, code)
	formattedTests, testsErr := formatGoSource(ctx, tests)
	if codeErr != nil || testsErr != nil {
		errs := []string{}
		if codeErr != nil {
			errs = append(errs, fmt.Sprintf("code: %v", codeErr))
		}
		if testsErr != nil {
			errs = append(errs, fmt.Sprintf("tests: %v", testsErr))
		}
		writeError(w, http.StatusBadRequest, "format failed: "+strings.Join(errs, "; "))
		return
	}

	writeJSON(w, http.StatusOK, formatResponse{
		Code:  formattedCode,
		Tests: formattedTests,
	})
}

func coachHint(passed bool, failedTests []string, insights []failureInsight) string {
	if passed {
		return "Nice pass. Capture one sentence about what changed, then move to the next kata."
	}
	if len(failedTests) == 0 {
		return "Start with the first failure line. Make one small fix, run again, repeat."
	}
	if len(insights) > 0 {
		return "Focus on the first mismatch only. Get that passing before touching anything else."
	}
	return "Use the failing test names as your map. Fix in order, smallest behavior gap first."
}

func extractFailureInsights(output string) []failureInsight {
	lines := strings.Split(output, "\n")
	insights := make([]failureInsight, 0, 6)
	seen := map[string]bool{}

	add := func(item failureInsight) {
		key := item.Kind + "|" + item.Summary + "|" + item.Expected + "|" + item.Actual
		if seen[key] || item.Summary == "" {
			return
		}
		seen[key] = true
		insights = append(insights, item)
	}

	for i := 0; i < len(lines) && len(insights) < 6; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		normalized := strings.TrimSpace(goTestPrefixPattern.ReplaceAllString(line, ""))
		lower := strings.ToLower(normalized)

		if strings.HasPrefix(normalized, "--- FAIL:") {
			add(failureInsight{Kind: "test", Summary: clip(normalized, 140)})
			continue
		}
		if strings.Contains(lower, "panic:") {
			add(failureInsight{Kind: "panic", Summary: clip(normalized, 140)})
			continue
		}

		if strings.HasPrefix(lower, "got:") || strings.HasPrefix(lower, "actual:") {
			actual := strings.TrimSpace(strings.SplitN(normalized, ":", 2)[1])
			expected := ""
			if i+1 < len(lines) {
				next := strings.TrimSpace(lines[i+1])
				next = strings.TrimSpace(goTestPrefixPattern.ReplaceAllString(next, ""))
				nextLower := strings.ToLower(next)
				if strings.HasPrefix(nextLower, "want:") || strings.HasPrefix(nextLower, "expected:") {
					expected = strings.TrimSpace(strings.SplitN(next, ":", 2)[1])
				}
			}
			add(failureInsight{
				Kind:     "mismatch",
				Summary:  "Expected and actual values do not match.",
				Expected: clip(expected, 120),
				Actual:   clip(actual, 120),
			})
			continue
		}

		match := expectedGotPattern.FindStringSubmatch(normalized)
		if len(match) == 3 {
			add(failureInsight{
				Kind:     "mismatch",
				Summary:  "Expected and actual values do not match.",
				Expected: clip(strings.TrimSpace(match[1]), 120),
				Actual:   clip(strings.TrimSpace(match[2]), 120),
			})
		}
	}

	return insights
}
