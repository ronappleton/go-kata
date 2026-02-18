package progress

import (
	"path/filepath"
	"testing"
	"time"
)

func TestStoreRecordAttemptPersists(t *testing.T) {
	tempDir := t.TempDir()
	store := NewStore(filepath.Join(tempDir, "progress.json"))

	state, err := store.Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if len(state.Attempts) != 0 {
		t.Fatalf("expected empty state, got %v", state.Attempts)
	}

	first, err := store.RecordAttempt("001", AttemptResult{
		Passed:      false,
		Duration:    500 * time.Millisecond,
		FailedTests: []string{"TestFizzBuzz"},
		OutputTail:  "some output",
		RanAt:       time.Now().UTC(),
	})
	if err != nil {
		t.Fatalf("RecordAttempt(1) returned error: %v", err)
	}
	if first.Attempts["001"].Attempts != 1 {
		t.Fatalf("expected attempts=1, got %d", first.Attempts["001"].Attempts)
	}
	if first.Attempts["001"].Passes != 0 {
		t.Fatalf("expected passes=0, got %d", first.Attempts["001"].Passes)
	}

	second, err := store.RecordAttempt("001", AttemptResult{
		Passed:   true,
		Duration: 250 * time.Millisecond,
		RanAt:    time.Now().UTC(),
	})
	if err != nil {
		t.Fatalf("RecordAttempt(2) returned error: %v", err)
	}
	if second.Attempts["001"].Attempts != 2 {
		t.Fatalf("expected attempts=2, got %d", second.Attempts["001"].Attempts)
	}
	if second.Attempts["001"].Passes != 1 {
		t.Fatalf("expected passes=1, got %d", second.Attempts["001"].Passes)
	}
	if second.Attempts["001"].LastResult != "pass" {
		t.Fatalf("expected last result=pass, got %q", second.Attempts["001"].LastResult)
	}
}
