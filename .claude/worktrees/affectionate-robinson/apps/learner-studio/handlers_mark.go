package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/marking"
)

func (s *studioServer) handleMark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req markRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, category, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	kataCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata.go: %v", err))
		return
	}
	testCode, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	last := state.Attempts[kata.ID]
	prompt := marking.BuildPrompt(marking.PromptInput{
		TrackTitle:      s.track.Title,
		CategoryTitle:   category.Title,
		KataID:          kata.ID,
		KataTitle:       kata.Title,
		Focus:           kata.Focus,
		Signature:       kata.Signature,
		Rules:           kata.Rules,
		KataCode:        string(kataCode),
		TestCode:        string(testCode),
		LastRunResult:   last.LastResult,
		LastFailedTests: last.LastFailedTests,
		LastOutputTail:  last.LastOutputTail,
	})

	target := marking.DefaultPromptPath(s.repoRoot, kata.ID, time.Now().UTC())
	if err := marking.WritePrompt(target, prompt); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write prompt: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, markResponse{
		PromptPath: target,
		Prompt:     prompt,
	})
}
