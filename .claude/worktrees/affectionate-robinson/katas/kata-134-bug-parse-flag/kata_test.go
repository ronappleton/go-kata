package kata

import "testing"

func TestParseFlag(t *testing.T) {
	v, err := ParseFlag(" no ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v {
		t.Fatalf("ParseFlag should return false for 'no'")
	}
}
