package rendering

import (
	"strings"
	"testing"
)

func TestInlineProseIsEscapedForPango(t *testing.T) {
	got := MarkdownToPango("Return <Name> & keep it **safe**")
	if strings.Contains(got, "<Name>") || strings.Contains(got, " & ") {
		t.Fatalf("prose leaked markup characters: %q", got)
	}
	if !strings.Contains(got, "&lt;Name&gt;") || !strings.Contains(got, "&amp;") {
		t.Fatalf("expected escaped prose, got %q", got)
	}
	if !strings.Contains(got, "<b>safe</b>") {
		t.Fatalf("expected bold markup, got %q", got)
	}
}
