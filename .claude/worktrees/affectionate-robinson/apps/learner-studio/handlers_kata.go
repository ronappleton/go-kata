package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (s *studioServer) handleKata(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	rawID := strings.TrimSpace(r.URL.Query().Get("id"))
	if rawID == "" {
		writeError(w, http.StatusBadRequest, "missing id query parameter")
		return
	}

	kata, category, ok := s.track.FindKata(rawID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	readme, err := os.ReadFile(kata.ReadmePath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read readme: %v", err))
		return
	}
	code, err := os.ReadFile(filepath.Join(kata.Dir, "kata.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata.go: %v", err))
		return
	}
	tests, err := os.ReadFile(filepath.Join(kata.Dir, "kata_test.go"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	state, err := s.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("load progress: %v", err))
		return
	}

	resp := kataResponse{
		ID:        kata.ID,
		Title:     kata.Title,
		Focus:     kata.Focus,
		Signature: kata.Signature,
		Rules:     append([]string(nil), kata.Rules...),
		Category: kataCategoryReference{
			ID:    category.ID,
			Title: category.Title,
		},
		Readme:   string(readme),
		Code:     string(code),
		Tests:    string(tests),
		Progress: state.Attempts[kata.ID],
	}

	writeJSON(w, http.StatusOK, resp)
}

func (s *studioServer) handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req saveRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), []byte(req.Code), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
		return
	}
	if err := os.WriteFile(filepath.Join(kata.Dir, "kata_test.go"), []byte(req.Tests), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata_test.go: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "saved"})
}

func (s *studioServer) handleResetBuggy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req resetBuggyRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	kata, _, ok := s.track.FindKata(req.KataID)
	if !ok {
		writeError(w, http.StatusNotFound, "kata not found")
		return
	}

	buggyCodePath := filepath.Join(kata.Dir, "buggy_kata.go")
	rawBuggyCode, err := os.ReadFile(buggyCodePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writeError(w, http.StatusBadRequest, "buggy starter not available for this kata")
			return
		}
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read buggy starter: %v", err))
		return
	}
	buggyCode := stripIgnoreBuildTags(rawBuggyCode)

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.WriteFile(filepath.Join(kata.Dir, "kata.go"), buggyCode, 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("write kata.go: %v", err))
		return
	}

	testsPath := filepath.Join(kata.Dir, "kata_test.go")
	tests, err := os.ReadFile(testsPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("read kata_test.go: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, formatResponse{
		Code:  string(buggyCode),
		Tests: string(tests),
	})
}

func stripIgnoreBuildTags(data []byte) []byte {
	const tagPrefix = "//go:build ignore\n// +build ignore\n\n"
	if strings.HasPrefix(string(data), tagPrefix) {
		return []byte(strings.TrimPrefix(string(data), tagPrefix))
	}
	return data
}
