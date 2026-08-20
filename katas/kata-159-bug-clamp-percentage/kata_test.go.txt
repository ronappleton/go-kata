package kata

import "testing"

func TestClampPercentage(t *testing.T) {
	if got := ClampPercentage(-10); got != 0 {
		t.Fatalf("ClampPercentage(-10) = %d, want 0", got)
	}
}
