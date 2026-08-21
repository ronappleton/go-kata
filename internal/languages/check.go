package languages

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// GoFmtChecker reports Go syntax errors by running gofmt -e on the source.
// It returns nil diagnostics when gofmt is unavailable.
type GoFmtChecker struct{}

var goErrorRe = regexp.MustCompile(`^.*?:(\d+):(\d+):\s*(.*)$`)

// Check runs gofmt -e over src and converts its diagnostics.
func (GoFmtChecker) Check(src string) []Diagnostic {
	cmd := exec.Command("gofmt", "-e")
	cmd.Stdin = strings.NewReader(src)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	_ = cmd.Run() // non-zero exit is expected when there are errors

	// gofmt writes parse errors to stdout (with -e) — some versions use stderr.
	all := out.String() + "\n" + errBuf.String()
	var diags []Diagnostic
	for _, line := range strings.Split(all, "\n") {
		m := goErrorRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		lineNo, _ := strconv.Atoi(m[1])
		col, _ := strconv.Atoi(m[2])
		lineNo-- // 0-based
		col--
		if col < 0 {
			col = 0
		}
		diags = append(diags, Diagnostic{
			Line:    lineNo,
			Col:     col,
			EndLine: lineNo,
			EndCol:  col + 1,
			Message: strings.TrimSpace(m[3]),
			IsError: true,
		})
	}
	return diags
}

// NoopChecker reports no diagnostics; used for languages without a checker.
type NoopChecker struct{}

func (NoopChecker) Check(string) []Diagnostic { return nil }
