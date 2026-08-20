package kata

import (
	"strings"
	"testing"
)

func TestFirstNonEmpty(t *testing.T) {
	got := FirstNonEmpty([]string{"", "   ", "go"}, "fallback")
	if strings.TrimSpace(got) != "go" {
		t.Fatalf("FirstNonEmpty mismatch: got %q want %q", got, "go")
	}
}
