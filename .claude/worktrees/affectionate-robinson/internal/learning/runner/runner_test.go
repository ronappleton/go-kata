package runner

import (
	"strings"
	"testing"
)

func TestParseTestOutputPass(t *testing.T) {
	output := strings.Join([]string{
		`{"Action":"run","Package":"kata001","Test":"TestFizzBuzz"}`,
		`{"Action":"pass","Package":"kata001","Test":"TestFizzBuzz","Elapsed":0.01}`,
		`{"Action":"pass","Package":"kata001","Elapsed":0.02}`,
	}, "\n")

	result := parseTestOutput(output)

	if len(result.FailedTests) != 0 {
		t.Fatalf("expected no failed tests, got %v", result.FailedTests)
	}
	if got := result.PackageState["kata001"]; got != "pass" {
		t.Fatalf("expected package pass, got %q", got)
	}
	if result.Elapsed <= 0 {
		t.Fatalf("expected elapsed > 0, got %v", result.Elapsed)
	}
}

func TestParseTestOutputFailingTests(t *testing.T) {
	output := strings.Join([]string{
		`{"Action":"run","Package":"kata001","Test":"TestFizzBuzz_Buzz"}`,
		`{"Action":"fail","Package":"kata001","Test":"TestFizzBuzz_Buzz","Elapsed":0.01}`,
		`{"Action":"fail","Package":"kata001","Elapsed":0.02}`,
	}, "\n")

	result := parseTestOutput(output)

	if len(result.FailedTests) != 1 || result.FailedTests[0] != "TestFizzBuzz_Buzz" {
		t.Fatalf("unexpected failing tests: %v", result.FailedTests)
	}
	if got := result.PackageState["kata001"]; got != "fail" {
		t.Fatalf("expected package fail, got %q", got)
	}
}
