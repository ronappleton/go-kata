package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var version = "dev"

const defaultContentURL = "https://raw.githubusercontent.com/ronappleton/gokatas-content/main"

type desktopConfig struct {
	ContentRoot string
	ContentURL  string
	Image       string
	DevMode     bool
}

func main() {
	contentFlag := flag.String("content", "", "Path to immutable curriculum content (development override)")
	contentURLFlag := flag.String("content-url", strings.TrimSpace(os.Getenv("GOKATAS_CONTENT_URL")), "Remote curriculum base URL")
	imageFlag := flag.String("runner-image", defaultRunnerImage(), "Digest-pinned Podman runner image")
	devFlag := flag.Bool("dev", false, "Enable developer diagnostics")
	flag.Parse()

	contentRoot, err := resolveContentRoot(*contentFlag)
	if err != nil && strings.TrimSpace(*contentFlag) != "" {
		fatalf("resolve curriculum content: %v", err)
	}
	contentURL := strings.TrimSpace(*contentURLFlag)
	if contentURL == "" {
		contentURL = defaultContentURL
	}

	code := runNative(desktopConfig{
		ContentRoot: contentRoot,
		ContentURL:  contentURL,
		Image:       strings.TrimSpace(*imageFlag),
		DevMode:     *devFlag,
	})
	if code != 0 {
		os.Exit(code)
	}
}

func defaultRunnerImage() string {
	if image := strings.TrimSpace(os.Getenv("GOKATAS_RUNNER_IMAGE")); image != "" {
		return image
	}
	configHome := strings.TrimSpace(os.Getenv("XDG_CONFIG_HOME"))
	if configHome == "" {
		home, err := os.UserHomeDir()
		if err == nil {
			configHome = filepath.Join(home, ".config")
		}
	}
	for _, path := range []string{
		filepath.Join(configHome, "gokatas", "runner-image"),
		"/etc/gokatas/runner-image",
	} {
		data, err := os.ReadFile(path)
		if err == nil && strings.TrimSpace(string(data)) != "" {
			return strings.TrimSpace(string(data))
		}
	}
	return ""
}

func resolveContentRoot(requested string) (string, error) {
	candidates := []string{}
	if strings.TrimSpace(requested) != "" {
		candidates = append(candidates, requested)
	} else {
		candidates = append(candidates, ".", "/usr/share/gokatas")
	}

	for _, candidate := range candidates {
		root, err := filepath.Abs(candidate)
		if err != nil {
			continue
		}
		if _, err := os.Stat(filepath.Join(root, "tracks", "go-core-100", "track.json")); err == nil {
			return root, nil
		}
	}
	return "", errors.New("no curriculum found; use -content pointing at the installed content root")
}

func fatalf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
