package languages

import (
	"os/exec"
	"testing"
)

func TestGoFmtCheckerFindsErrors(t *testing.T) {
	if _, err := exec.LookPath("gofmt"); err != nil {
		t.Skip("gofmt not installed")
	}
	diags := GoFmtChecker{}.Check("package kata\nfunc main( {\n")
	if len(diags) == 0 {
		t.Fatal("expected diagnostics for invalid Go")
	}
	first := diags[0]
	if !first.IsError || first.Message == "" {
		t.Fatalf("expected error diagnostic with message, got %+v", first)
	}
	// gofmt reports the parse error on the line with the bad function.
	if first.Line != 1 {
		t.Fatalf("expected line 1 (0-based) for the error, got %d", first.Line)
	}
}

func TestGoFmtCheckerClean(t *testing.T) {
	if _, err := exec.LookPath("gofmt"); err != nil {
		t.Skip("gofmt not installed")
	}
	src := "package kata\n\nfunc Main() string { return \"\" }\n"
	diags := GoFmtChecker{}.Check(src)
	if len(diags) != 0 {
		t.Fatalf("expected no diagnostics for valid Go, got %+v", diags)
	}
}

func TestNoopChecker(t *testing.T) {
	diags := NoopChecker{}.Check("anything at all {{{\n")
	if len(diags) != 0 {
		t.Fatal("noop checker should never report diagnostics")
	}
}
