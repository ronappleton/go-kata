package runner

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"sort"
	"strings"
	"time"
)

type Result struct {
	Passed       bool
	CompileErr   bool
	FailedTests  []string
	PackageState map[string]string
	Elapsed      time.Duration
	RawOutput    string

	testEventsSeen bool // unexported: used internally to detect compile-only failures
}

type testEvent struct {
	Action  string  `json:"Action"`
	Package string  `json:"Package"`
	Test    string  `json:"Test"`
	Elapsed float64 `json:"Elapsed"`
}

func RunKataTests(ctx context.Context, kataDir string) (Result, error) {
	cmd := exec.CommandContext(ctx, "go", "test", "-json", "./...")
	cmd.Dir = kataDir

	outputBytes, execErr := cmd.CombinedOutput()
	output := string(outputBytes)

	result := parseTestOutput(output)

	// Prefer structured package/test actions to decide pass/fail when available.
	switch {
	case len(result.PackageState) == 0:
		result.Passed = execErr == nil
	default:
		result.Passed = true
		for _, action := range result.PackageState {
			if action != "pass" {
				result.Passed = false
				break
			}
		}
		if len(result.FailedTests) > 0 {
			result.Passed = false
		}
	}

	// Detect compile errors: package failed but no individual test events were
	// recorded (no tests ran because the code didn't compile).
	if !result.Passed && len(result.FailedTests) == 0 && !result.testEventsSeen {
		result.CompileErr = true
	}

	return result, execErr
}

func parseTestOutput(output string) Result {
	result := Result{
		PackageState: map[string]string{},
		RawOutput:    output,
	}

	failedSet := map[string]bool{}
	maxElapsed := 0.0
	sawTestEvent := false

	scanner := bufio.NewScanner(bytes.NewReader([]byte(output)))
	// Keep scanner robust for larger output lines.
	buf := make([]byte, 0, 1024*64)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0] != '{' {
			continue
		}

		var event testEvent
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			continue
		}

		if event.Elapsed > maxElapsed {
			maxElapsed = event.Elapsed
		}

		if event.Test != "" {
			sawTestEvent = true
			if event.Action == "fail" {
				failedSet[event.Test] = true
			}
			continue
		}

		if event.Test == "" && (event.Action == "pass" || event.Action == "fail") {
			result.PackageState[event.Package] = event.Action
		}
	}

	if maxElapsed > 0 {
		result.Elapsed = time.Duration(maxElapsed * float64(time.Second))
	}

	for testName := range failedSet {
		result.FailedTests = append(result.FailedTests, testName)
	}
	sort.Strings(result.FailedTests)

	result.testEventsSeen = sawTestEvent
	return result
}
