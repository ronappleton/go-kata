package content

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// archiveURLFromRaw derives the codeload zipball URL from a raw GitHub base
// URL. Returns "" when the base URL is not a raw.githubusercontent.com URL.
//
//	https://raw.githubusercontent.com/OWNER/REPO/BRANCH
//	→ https://codeload.github.com/OWNER/REPO/zip/refs/heads/BRANCH
func archiveURLFromRaw(baseURL string) string {
	trimmed := strings.TrimPrefix(baseURL, "https://raw.githubusercontent.com/")
	if trimmed == baseURL {
		return ""
	}
	parts := strings.SplitN(trimmed, "/", 3)
	if len(parts) != 3 {
		return ""
	}
	return "https://codeload.github.com/" + parts[0] + "/" + parts[1] + "/zip/refs/heads/" + parts[2]
}

// DownloadArchive fetches the full content tree as a single zipball.
func (r *remoteStore) DownloadArchive(ctx context.Context) ([]byte, error) {
	url := archiveURLFromRaw(r.baseURL)
	if url == "" {
		return nil, fmt.Errorf("cannot derive archive URL from %s", r.baseURL)
	}
	return r.fetchURL(ctx, url)
}

// fetchURL fetches an absolute URL with the same retry behaviour as fetch.
func (r *remoteStore) fetchURL(ctx context.Context, url string) ([]byte, error) {
	delay := 400 * time.Millisecond
	var lastErr error
	for attempt := 0; attempt < maxFetchAttempts; attempt++ {
		data, permanent, err := r.fetchAttempt(ctx, url)
		if err == nil {
			return data, nil
		}
		lastErr = err
		if permanent || attempt == maxFetchAttempts-1 {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
		}
		delay *= 2
	}
	return nil, lastErr
}

// extractArchive writes a zipball's manifest.json and tracks/ tree into dir,
// stripping the archive's root folder and guarding against path traversal.
// progress, when non-nil, reports kata files written as (done, total).
func extractArchive(data []byte, dir string, progress func(done, total int)) (int, error) {
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return 0, fmt.Errorf("open archive: %w", err)
	}

	// Count kata files up front so progress has a total.
	total := 0
	for _, f := range zr.File {
		if isKataFile(f.Name) {
			total++
		}
	}

	added := 0
	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}
		name := stripArchiveRoot(f.Name)
		if name == "" || name == "README.md" || name == ".gitignore" {
			continue
		}
		// Reject path traversal outright, before any other filtering, so a
		// malicious archive cannot silently smuggle files outside the cache.
		if strings.Contains(name, "..") {
			return added, fmt.Errorf("unsafe archive path %q", f.Name)
		}
		if name != "manifest.json" && !strings.HasPrefix(name, "tracks/") {
			continue
		}

		dest, err := safeJoin(dir, name)
		if err != nil {
			return added, err
		}
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return added, fmt.Errorf("create dir for %s: %w", name, err)
		}

		rc, err := f.Open()
		if err != nil {
			return added, fmt.Errorf("open %s in archive: %w", f.Name, err)
		}
		content, readErr := io.ReadAll(rc)
		rc.Close()
		if readErr != nil {
			return added, fmt.Errorf("read %s from archive: %w", f.Name, readErr)
		}
		if err := atomicWriteFile(dest, content, 0o644); err != nil {
			return added, fmt.Errorf("write %s: %w", name, err)
		}
		if isKataFile(f.Name) {
			added++
			if progress != nil {
				progress(added, total)
			}
		}
	}
	return added, nil
}

// isKataFile reports whether an archive entry is a kata content file
// (tracks/<track>/katas/<id>.json).
func isKataFile(name string) bool {
	base := filepath.Base(name)
	if !strings.HasSuffix(base, ".json") {
		return false
	}
	parts := strings.Split(filepath.ToSlash(name), "/")
	return len(parts) >= 4 && parts[len(parts)-2] == "katas"
}

// stripArchiveRoot removes the top-level folder GitHub puts in zipballs
// (e.g. "gokatas-content-main/"). Archives without a root folder pass through.
func stripArchiveRoot(name string) string {
	first, rest := name, ""
	if idx := strings.IndexByte(name, '/'); idx >= 0 {
		first, rest = name[:idx], name[idx+1:]
	}
	switch first {
	case "manifest.json", "tracks":
		return name
	default:
		return rest
	}
}

// maxPathLen guards against absurdly long archive entries that would fail
// (or worse, overflow) filesystem operations.
const maxPathLen = 4096

// safeJoin joins dir and name, refusing any path that escapes dir or exceeds
// a sane length.
func safeJoin(dir, name string) (string, error) {
	if len(name) == 0 || len(name) > maxPathLen {
		return "", fmt.Errorf("unsafe archive path %q", name)
	}
	clean := filepath.Clean(filepath.Join(dir, name))
	rel, err := filepath.Rel(dir, clean)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("unsafe archive path %q", name)
	}
	return clean, nil
}
