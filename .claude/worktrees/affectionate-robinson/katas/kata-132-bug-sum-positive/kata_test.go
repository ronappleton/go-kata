package kata

import "testing"

func TestSumPositive(t *testing.T) {
	got := SumPositive([]int{5, -3, 0, 2, -1})
	if got != 7 {
		t.Fatalf("SumPositive mismatch: got %d want 7", got)
	}
}
