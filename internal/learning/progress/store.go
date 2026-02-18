package progress

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	path string
}

type State struct {
	Attempts map[string]KataProgress `json:"attempts"`
}

type KataProgress struct {
	Attempts        int       `json:"attempts"`
	Passes          int       `json:"passes"`
	LastResult      string    `json:"last_result"`
	LastRunAt       time.Time `json:"last_run_at"`
	LastDurationMS  int64     `json:"last_duration_ms"`
	LastFailedTests []string  `json:"last_failed_tests,omitempty"`
	LastOutputTail  string    `json:"last_output_tail,omitempty"`
}

type AttemptResult struct {
	Passed      bool
	Duration    time.Duration
	FailedTests []string
	OutputTail  string
	RanAt       time.Time
}

func NewStore(path string) *Store {
	return &Store{path: path}
}

func (s *Store) Load() (State, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return State{Attempts: map[string]KataProgress{}}, nil
		}
		return State{}, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return State{}, err
	}
	if state.Attempts == nil {
		state.Attempts = map[string]KataProgress{}
	}
	return state, nil
}

func (s *Store) Save(state State) error {
	if state.Attempts == nil {
		state.Attempts = map[string]KataProgress{}
	}

	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0o644)
}

func (s *Store) RecordAttempt(kataID string, result AttemptResult) (State, error) {
	state, err := s.Load()
	if err != nil {
		return State{}, err
	}

	entry := state.Attempts[kataID]
	entry.Attempts++
	if result.Passed {
		entry.Passes++
		entry.LastResult = "pass"
	} else {
		entry.LastResult = "fail"
	}

	entry.LastRunAt = result.RanAt
	entry.LastDurationMS = result.Duration.Milliseconds()
	entry.LastFailedTests = append([]string(nil), result.FailedTests...)
	entry.LastOutputTail = result.OutputTail

	state.Attempts[kataID] = entry

	if err := s.Save(state); err != nil {
		return State{}, err
	}
	return state, nil
}

func CompletedCount(state State, kataIDs []string) int {
	completed := 0
	for _, id := range kataIDs {
		if state.Attempts[id].Passes > 0 {
			completed++
		}
	}
	return completed
}
