package main

import (
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ronappleton/golang-katas-1-100/internal/learning/catalog"
	"github.com/ronappleton/golang-katas-1-100/internal/learning/progress"
)

const (
	defaultTrackConfigRelative = "tracks/go-core-100/track.json"
	defaultPathwaysConfigPath  = "tracks/go-core-100/pathways.json"
	defaultRunTimeoutSeconds   = 90
	defaultOutputTailLines     = 100
)

//go:embed static/*
var embeddedStatic embed.FS

type studioServer struct {
	repoRoot string
	track    catalog.Track
	pathways []pathwayDefinition
	store    *progress.Store
	files    fs.FS
	mu       sync.Mutex
}

func main() {
	repoRootFlag := flag.String("repo", ".", "Path to kata repository root")
	addrFlag := flag.String("addr", "127.0.0.1:7777", "HTTP listen address")
	flag.Parse()

	server, err := newStudioServer(*repoRootFlag)
	if err != nil {
		fatalf("init learner studio: %v", err)
	}

	fmt.Printf("Learner Studio running at http://%s\n", *addrFlag)
	if err := http.ListenAndServe(*addrFlag, server.routes()); err != nil {
		fatalf("listen: %v", err)
	}
}

func newStudioServer(repoRoot string) (*studioServer, error) {
	absRepoRoot, err := filepath.Abs(repoRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve repo root: %w", err)
	}

	track, err := catalog.LoadTrack(filepath.Join(absRepoRoot, defaultTrackConfigRelative), absRepoRoot)
	if err != nil {
		return nil, err
	}

	pathways, err := loadPathways(filepath.Join(absRepoRoot, defaultPathwaysConfigPath))
	if err != nil {
		return nil, err
	}

	store := progress.NewStore(filepath.Join(absRepoRoot, ".learning", "progress.json"))
	if _, err := store.Load(); err != nil {
		return nil, fmt.Errorf("load progress: %w", err)
	}

	staticFS, err := fs.Sub(embeddedStatic, "static")
	if err != nil {
		return nil, fmt.Errorf("load embedded static files: %w", err)
	}

	return &studioServer{
		repoRoot: absRepoRoot,
		track:    track,
		pathways: pathways,
		store:    store,
		files:    staticFS,
	}, nil
}

func loadPathways(path string) ([]pathwayDefinition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read pathways config: %w", err)
	}

	var cfg pathwaysConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse pathways config: %w", err)
	}

	if len(cfg.Pathways) == 0 {
		return nil, errors.New("pathways config must include at least one pathway")
	}

	for _, item := range cfg.Pathways {
		if strings.TrimSpace(item.ID) == "" || strings.TrimSpace(item.Title) == "" {
			return nil, errors.New("pathways config contains item with empty id/title")
		}
		if len(item.Categories) == 0 {
			return nil, fmt.Errorf("pathway %q must include categories", item.ID)
		}
	}

	return cfg.Pathways, nil
}

func (s *studioServer) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/app.js", s.handleStaticFile("app.js", "text/javascript; charset=utf-8"))
	mux.HandleFunc("/styles.css", s.handleStaticFile("styles.css", "text/css; charset=utf-8"))

	mux.HandleFunc("/api/track", s.handleTrack)
	mux.HandleFunc("/api/pathways", s.handlePathways)
	mux.HandleFunc("/api/kata", s.handleKata)
	mux.HandleFunc("/api/learn", s.handleLearn)
	mux.HandleFunc("/api/save", s.handleSave)
	mux.HandleFunc("/api/reset-buggy", s.handleResetBuggy)
	mux.HandleFunc("/api/format", s.handleFormat)
	mux.HandleFunc("/api/run", s.handleRun)
	mux.HandleFunc("/api/mark", s.handleMark)

	return mux
}

func (s *studioServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	s.serveFile(w, "index.html", "text/html; charset=utf-8")
}

func (s *studioServer) handleStaticFile(name, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.serveFile(w, name, contentType)
	}
}

func (s *studioServer) serveFile(w http.ResponseWriter, name, contentType string) {
	data, err := fs.ReadFile(s.files, name)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", contentType)
	_, _ = w.Write(data)
}

func fatalf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
