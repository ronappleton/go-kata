package runner

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/evaluator"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/katas"
)

type Result struct {
	Passed       bool
	FailedTests  []string
	PackageState map[string]string
	Elapsed      time.Duration
	RawOutput    string
}

type testEvent struct {
	Action  string  `json:"Action"`
	Package string  `json:"Package"`
	Test    string  `json:"Test"`
	Elapsed float64 `json:"Elapsed"`
}

// RunKataTests is retained for CLI compatibility. All execution is delegated
// to the rootless Podman evaluator; no learner code is executed on the host.
func RunKataTests(ctx context.Context, kataID string, content katas.KataContent) (Result, error) {
	image := strings.TrimSpace(os.Getenv("GOKATAS_RUNNER_IMAGE"))
	if image == "" {
		return Result{}, errors.New("GOKATAS_RUNNER_IMAGE must contain a digest-pinned evaluator image")
	}

	runner, err := evaluator.NewRunner(image)
	if err != nil {
		return Result{}, err
	}
	run := runner.Run(ctx, evaluator.Request{
		KataID:       kataID,
		Module:       "kata" + kataID,
		Code:         content.KataGo,
		LearnerTests: "",
		TrustedTests: content.KataTest,
	})

	result := Result{
		Passed:       run.Passed,
		FailedTests:  append([]string(nil), run.FailedTests...),
		Elapsed:      run.Duration,
		RawOutput:    run.Output,
		PackageState: map[string]string{},
	}
	if run.Passed {
		result.PackageState["kata"+kataID] = "pass"
	} else {
		result.PackageState["kata"+kataID] = "fail"
	}
	if run.EvaluatorError != "" {
		return result, errors.New(run.EvaluatorError)
	}
	return result, nil
}

func parseTestOutput(output string) Result {
	result := Result{
		PackageState: map[string]string{},
		RawOutput:    output,
	}

	failedSet := map[string]bool{}
	maxElapsed := 0.0

	scanner := bufio.NewScanner(bytes.NewReader([]byte(output)))
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

		if event.Test != "" && event.Action == "fail" {
			failedSet[event.Test] = true
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
	return result
}
