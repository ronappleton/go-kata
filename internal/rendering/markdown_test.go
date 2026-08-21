package rendering

import (
	"strings"
	"testing"
)

func TestHeading1(t *testing.T) {
	got := MarkdownToPango("# Hello World")
	if !strings.Contains(got, `<span size="x-large" weight="800"`) {
		t.Fatalf("expected heading1 pango tag, got %q", got)
	}
	if !strings.Contains(got, "Hello World") {
		t.Fatalf("expected heading text, got %q", got)
	}
}

func TestHeading2(t *testing.T) {
	got := MarkdownToPango("## Subheading")
	if !strings.Contains(got, `<span size="large" weight="700"`) {
		t.Fatalf("expected heading2 pango tag, got %q", got)
	}
}

func TestBold(t *testing.T) {
	got := MarkdownToPango("This is **bold** text")
	if !strings.Contains(got, "<b>bold</b>") {
		t.Fatalf("expected <b>bold</b>, got %q", got)
	}
}

func TestInlineCode(t *testing.T) {
	got := MarkdownToPango("Use `go test` to run")
	if !strings.Contains(got, "go test") {
		t.Fatalf("expected inline code content, got %q", got)
	}
}

func TestBulletItem(t *testing.T) {
	got := MarkdownToPango("- First item")
	if !strings.Contains(got, "•") {
		t.Fatalf("expected bullet character, got %q", got)
	}
	if !strings.Contains(got, "First item") {
		t.Fatalf("expected bullet text, got %q", got)
	}
}

func TestFencedCodeBlock(t *testing.T) {
	input := "```go\nfunc main() {}\n```"
	got := MarkdownToPango(input)
	if !strings.Contains(got, `font_family="monospace"`) {
		t.Fatalf("expected monospace span open tag, got %q", got)
	}
	if !strings.Contains(got, "</span>") {
		t.Fatalf("expected </span> close tag, got %q", got)
	}
	if !strings.Contains(got, "func main()") {
		t.Fatalf("expected code content, got %q", got)
	}
}

func TestHorizontalRule(t *testing.T) {
	got := MarkdownToPango("---")
	if !strings.Contains(got, "─") {
		t.Fatalf("expected horizontal rule, got %q", got)
	}
}

func TestURLDetection(t *testing.T) {
	got := MarkdownToPango("Visit https://go.dev/tour")
	if !strings.Contains(got, "https://go.dev/tour") {
		t.Fatalf("expected URL in output, got %q", got)
	}
}

func TestEmptyInput(t *testing.T) {
	got := MarkdownToPango("")
	if got != "" {
		t.Fatalf("expected empty output for empty input, got %q", got)
	}
}

func TestMixedContent(t *testing.T) {
	input := `# Kata 001

**Focus:** testing

- item one
- item two

` + "```go" + `
func hello() string
` + "```"

	got := MarkdownToPango(input)
	if !strings.Contains(got, `<span size="x-large" weight="800"`) {
		t.Fatalf("expected heading, got %q", got)
	}
	if !strings.Contains(got, "<b>Focus:</b>") {
		t.Fatalf("expected bold, got %q", got)
	}
	if !strings.Contains(got, "•") {
		t.Fatalf("expected bullet, got %q", got)
	}
	if !strings.Contains(got, `font_family="monospace"`) {
		t.Fatalf("expected code block, got %q", got)
	}
}

func TestHTMLEscapingInCodeBlock(t *testing.T) {
	input := "```html\n<div>\n```"
	got := MarkdownToPango(input)
	if strings.Contains(got, "<div>") {
		t.Fatalf("expected HTML to be escaped, got %q", got)
	}
	if !strings.Contains(got, "&lt;div&gt;") {
		t.Fatalf("expected escaped HTML, got %q", got)
	}
}

func TestNestedInlineInHeading(t *testing.T) {
	got := MarkdownToPango("# Use `go test`")
	if !strings.Contains(got, `<span size="x-large" weight="800"`) {
		t.Fatalf("expected heading tag, got %q", got)
	}
	if !strings.Contains(got, "go test") {
		t.Fatalf("expected inline code in heading, got %q", got)
	}
}

func TestUnclosedCodeBlockClosesAutomatically(t *testing.T) {
	input := "```go\nfunc main() {}"
	got := MarkdownToPango(input)
	if !strings.Contains(got, "</span>") {
		t.Fatalf("expected auto-closed code block, got %q", got)
	}
}
