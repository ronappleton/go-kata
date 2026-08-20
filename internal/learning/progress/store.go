package progress

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/workspace"
)

const currentSchemaVersion = 2

type Store struct {
	path string
	mu   sync.Mutex
}

type State struct {
	SchemaVersion int                     `json:"schema_version"`
	Attempts      map[string]KataProgress `json:"attempts"`
}

// KataProgress tracks a user's progress on a specific kata version.
type KataProgress struct {
	Version         string    `json:"version"`                     // Kata version when last completed (e.g., "1.0.0")
	Attempts        int       `json:"attempts"`                    // Total attempts on this version
	Passes          int       `json:"passes"`                      // Passes on this version
	LastResult      string    `json:"last_result"`                 // "pass" or "fail"
	LastRunAt       time.Time `json:"last_run_at"`                 // When last attempted
	LastDurationMS  int64     `json:"last_duration_ms"`            // Duration of last attempt
	LastFailedTests []string  `json:"last_failed_tests,omitempty"` // Failed test names
	LastOutputTail  string    `json:"last_output_tail,omitempty"`  // Last lines of output
}

// IsCompletedForVersion returns true if the kata was passed for the given version.
func (kp KataProgress) IsCompletedForVersion(version string) bool {
	return kp.Passes > 0 && kp.Version == version
}

// NeedsRecompletion returns true if the kata version has changed since last completion.
func (kp KataProgress) NeedsRecompletion(newVersion string) bool {
	return kp.Passes > 0 && kp.Version != "" && kp.Version != newVersion
}

type AttemptResult struct {
	Passed      bool
	Version     string      // Kata version being attempted
	Duration    time.Duration
	FailedTests []string
	OutputTail  string
	RanAt       time.Time
}

func NewStore(path string) *Store {
	return &Store{path: path}
}

func (s *Store) Load() (State, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.load()
}

func (s *Store) load() (State, error) {
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
	if state.SchemaVersion == 0 {
		// Migrate v1 -> v2: add empty version fields
		state.SchemaVersion = currentSchemaVersion
	}
	if state.SchemaVersion > currentSchemaVersion {
		return State{}, errors.New("progress state was created by a newer application version")
	}
	return state, nil
}

func (s *Store) Save(state State) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.save(state)
}

func (s *Store) save(state State) error {
	if state.Attempts == nil {
		state.Attempts = map[string]KataProgress{}
	}
	state.SchemaVersion = currentSchemaVersion

	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')

	return workspace.AtomicWrite(s.path, data, 0o600)
}

// RecordAttempt records an attempt on a kata with its version.
func (s *Store) RecordAttempt(kataID string, result AttemptResult) (State, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	state, err := s.load()
	if err != nil {
		return State{}, err
	}

	entry := state.Attempts[kataID]

	// If the kata version changed, reset progress for the new version
	if result.Version != "" && entry.Version != "" && entry.Version != result.Version {
		// Version changed — start fresh for new version
		entry.Attempts = 0
		entry.Passes = 0
		entry.Version = result.Version
	} else if entry.Version == "" && result.Version != "" {
		entry.Version = result.Version
	}

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

	if err := s.save(state); err != nil {
		return State{}, err
	}
	return state, nil
}

// IsCompleted checks if a kata is completed for a specific version.
func IsCompleted(state State, kataID, version string) bool {
	p, ok := state.Attempts[kataID]
	if !ok {
		return false
	}
	return p.IsCompletedForVersion(version)
}

// NeedsUpdate checks if a kata needs re-completion due to version change.
func NeedsUpdate(state State, kataID, currentVersion string) bool {
	p, ok := state.Attempts[kataID]
	if !ok {
		return false
	}
	return p.NeedsRecompletion(currentVersion)
}

// CompletedCount counts katas completed for their current versions.
func CompletedCount(state State, kataVersions map[string]string) int {
	completed := 0
	for id, version := range kataVersions {
		if IsCompleted(state, id, version) {
			completed++
		}
	}
	return completed
}

// CompletedCountLegacy counts katas with any pass (backward compat).
func CompletedCountLegacy(state State, kataIDs []string) int {
	completed := 0
	for _, id := range kataIDs {
		if state.Attempts[id].Passes > 0 {
			completed++
		}
	}
	return completed
}
