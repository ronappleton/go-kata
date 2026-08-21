package content

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// provider is the composite ContentManager that resolves content from
// local cache → remote fetch. Embedded content remains available only as
// an explicit compatibility fallback for callers that opt into it.
type provider struct {
	local         *localStore
	remote        *remoteStore
	embedded      *embeddedStore
	allowEmbedded bool

	mu       sync.RWMutex
	lastSync time.Time
	remoteOK bool // whether remote was reachable on last attempt

	// progress reports sync progress as (completed, total) kata downloads.
	// It is called from sync worker goroutines and may be nil.
	progress func(completed, total int)
}

// SetProgress installs a callback invoked during Sync as katas are downloaded.
// The callback may be called from multiple goroutines concurrently.
func (p *provider) SetProgress(fn func(completed, total int)) {
	p.progress = fn
}

// ProviderConfig configures the content provider.
type ProviderConfig struct {
	ContentDir string // Local cache directory (e.g., ~/.local/share/gokatas/content)
	RemoteURL  string // Remote content source URL (empty = no remote)
}

// NewProvider creates a composite content manager.
// It tries local cache first, then remote, then embedded fallback.
func NewProvider(cfg ProviderConfig) (*provider, error) {
	local, err := NewLocalStore(cfg.ContentDir)
	if err != nil {
		return nil, fmt.Errorf("create local store: %w", err)
	}

	var remote *remoteStore
	if cfg.RemoteURL != "" {
		remote = NewRemoteStore(cfg.RemoteURL)
	}

	return &provider{
		local:         local,
		remote:        remote,
		embedded:      NewEmbeddedStore(),
		allowEmbedded: false,
	}, nil
}

// GetManifest returns the manifest from the best available source.
func (p *provider) GetManifest(ctx context.Context) (*Manifest, error) {
	// Try local cache first
	if p.local.HasManifest() {
		m, err := p.local.GetManifest()
		if err == nil {
			return m, nil
		}
		log.Printf("local manifest read failed: %v, trying remote", err)
	}

	// Try remote
	if p.remote != nil {
		m, err := p.remote.GetManifest(ctx)
		if err == nil {
			// Cache locally
			_ = p.local.SaveManifest(m)
			p.mu.Lock()
			p.remoteOK = true
			p.mu.Unlock()
			return m, nil
		}
		log.Printf("remote manifest fetch failed: %v", err)
		p.mu.Lock()
		p.remoteOK = false
		p.mu.Unlock()
	}

	if p.allowEmbedded {
		return p.embedded.GetManifest(ctx)
	}
	return nil, fmt.Errorf("curriculum is not cached and remote content is unavailable")
}

// GetTrack returns track metadata from the best available source.
func (p *provider) GetTrack(ctx context.Context, trackID string) (*TrackMeta, error) {
	// Try local cache
	t, err := p.local.GetTrack(trackID)
	if err == nil {
		return t, nil
	}

	// Try remote
	if p.remote != nil {
		t, err := p.remote.GetTrack(ctx, trackID)
		if err == nil {
			_ = p.local.SaveTrack(trackID, t)
			return t, nil
		}
	}

	return nil, fmt.Errorf("track %s not found in any source", trackID)
}

// GetKata returns kata content from the best available source.
func (p *provider) GetKata(ctx context.Context, trackID, kataID string) (*KataContent, error) {
	// Try local cache
	k, err := p.local.GetKata(trackID, kataID)
	if err == nil {
		return k, nil
	}

	// Try remote
	if p.remote != nil {
		k, err := p.remote.GetKata(ctx, trackID, kataID)
		if err == nil {
			_ = p.local.SaveKata(trackID, k)
			return k, nil
		}
	}

	if p.allowEmbedded {
		k, err = p.embedded.GetKata(ctx, trackID, kataID)
		if err == nil {
			return k, nil
		}
	}
	return nil, fmt.Errorf("kata %s not found in any source", kataID)
}

// Sync downloads updated content from the remote source.
//
// The preferred path downloads the whole curriculum as a single zipball and
// extracts it atomically into the cache. When the archive is unavailable it
// falls back to per-kata fetches from a bounded worker pool. A sync where the
// manifest checksums are unchanged and every track is fully cached completes
// without any content download. Individual failures are recorded in the
// result and do not abort the whole sync.
func (p *provider) Sync(ctx context.Context) (*SyncResult, error) {
	if p.remote == nil {
		return &SyncResult{SyncedAt: time.Now()}, nil
	}

	result := &SyncResult{SyncedAt: time.Now()}

	// Fetch remote manifest
	remoteManifest, err := p.remote.GetManifest(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch remote manifest: %w", err)
	}

	// Nothing changed and everything is cached: skip the download entirely.
	if p.curriculumCurrent(remoteManifest) {
		p.mu.Lock()
		p.lastSync = result.SyncedAt
		p.remoteOK = true
		p.mu.Unlock()
		return result, nil
	}

	// Preferred path: one zipball instead of hundreds of kata requests.
	if data, aerr := p.remote.DownloadArchive(ctx); aerr == nil {
		added, xerr := extractArchive(data, p.local.Dir(), func(done, total int) {
			if p.progress != nil {
				p.progress(done, total)
			}
		})
		if xerr != nil {
			result.Failed = append(result.Failed, fmt.Sprintf("extract archive: %v", xerr))
		} else {
			result.Added = added
			if err := p.local.SaveManifest(remoteManifest); err != nil {
				result.Failed = append(result.Failed, fmt.Sprintf("save manifest: %v", err))
			}
			p.mu.Lock()
			p.lastSync = result.SyncedAt
			p.remoteOK = true
			p.mu.Unlock()
			return result, nil
		}
	} else {
		result.Failed = append(result.Failed, fmt.Sprintf("archive download failed, falling back to per-kata sync: %v", aerr))
	}

	// Fallback: per-kata concurrent fetches.
	added, failed := p.syncPerKata(ctx, remoteManifest)
	result.Added += added
	result.Failed = append(result.Failed, failed...)
	return result, nil
}

// curriculumCurrent reports whether every track in the remote manifest matches
// the previously synced checksums and is fully cached.
func (p *provider) curriculumCurrent(remote *Manifest) bool {
	prev, err := p.local.GetManifest()
	if err != nil {
		return false
	}
	byID := make(map[string]TrackEntry, len(prev.Tracks))
	for _, t := range prev.Tracks {
		byID[t.ID] = t
	}
	for _, t := range remote.Tracks {
		old, ok := byID[t.ID]
		if !ok || old.Checksum == "" || old.Checksum != t.Checksum {
			return false
		}
		track, err := p.local.GetTrack(t.ID)
		if err != nil || !p.trackFullyCached(t.ID, track) {
			return false
		}
	}
	return true
}

// syncPerKata downloads each kata individually with a bounded worker pool.
// It is the fallback when the zipball cannot be downloaded.
func (p *provider) syncPerKata(ctx context.Context, remoteManifest *Manifest) (added int, failed []string) {
	type kataJob struct {
		trackID string
		kataID  string
	}
	var jobs []kataJob

	// Sync each track
	for _, trackEntry := range remoteManifest.Tracks {
		track, err := p.remote.GetTrack(ctx, trackEntry.ID)
		if err != nil {
			failed = append(failed, fmt.Sprintf("track %s: %v", trackEntry.ID, err))
			continue
		}
		if err := p.local.SaveTrack(trackEntry.ID, track); err != nil {
			failed = append(failed, fmt.Sprintf("save track %s: %v", trackEntry.ID, err))
			continue
		}

		for _, stage := range track.Stages {
			for _, cat := range stage.Categories {
				for _, kataID := range cat.KataIDs {
					jobs = append(jobs, kataJob{trackID: trackEntry.ID, kataID: kataID})
				}
			}
		}
	}

	const workers = 8
	jobCh := make(chan kataJob)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var completed int
	report := func() {
		if p.progress != nil {
			p.progress(completed, len(jobs))
		}
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				k, err := p.remote.GetKata(ctx, job.trackID, job.kataID)
				if err != nil {
					mu.Lock()
					failed = append(failed, fmt.Sprintf("kata %s/%s: %v", job.trackID, job.kataID, err))
					completed++
					mu.Unlock()
					report()
					continue
				}
				if err := p.local.SaveKata(job.trackID, k); err != nil {
					mu.Lock()
					failed = append(failed, fmt.Sprintf("save kata %s: %v", job.kataID, err))
					completed++
					mu.Unlock()
				} else {
					mu.Lock()
					added++
					completed++
					mu.Unlock()
				}
				report()
			}
		}()
	}

	for _, job := range jobs {
		jobCh <- job
	}
	close(jobCh)
	wg.Wait()

	return added, failed
}

// trackFullyCached reports whether every kata referenced by the track is
// already present in the local cache.
func (p *provider) trackFullyCached(trackID string, track *TrackMeta) bool {
	for _, stage := range track.Stages {
		for _, cat := range stage.Categories {
			for _, kataID := range cat.KataIDs {
				if _, err := p.local.GetKata(trackID, kataID); err != nil {
					return false
				}
			}
		}
	}
	return true
}

// IsRemote returns true if the remote source is reachable.
func (p *provider) IsRemote() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.remote != nil && p.remoteOK
}

// LastSync returns when content was last synced.
func (p *provider) LastSync() time.Time {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if !p.lastSync.IsZero() {
		return p.lastSync
	}
	return p.local.LastSyncTime()
}

// ContentDir returns the local cache directory.
func (p *provider) ContentDir() string {
	return p.local.Dir()
}
