package content

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// provider is the composite ContentManager that resolves content from
// local cache → remote fetch → embedded fallback.
type provider struct {
	local    *localStore
	remote   *remoteStore
	embedded *embeddedStore

	mu       sync.RWMutex
	lastSync time.Time
	remoteOK bool // whether remote was reachable on last attempt
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
		local:    local,
		remote:   remote,
		embedded: NewEmbeddedStore(),
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

	// Fallback to embedded
	return p.embedded.GetManifest(ctx)
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

	// Try embedded (trackID doesn't matter for embedded)
	k, err = p.embedded.GetKata(ctx, trackID, kataID)
	if err != nil {
		return nil, fmt.Errorf("kata %s not found in any source", kataID)
	}
	return k, nil
}

// Sync downloads updated content from the remote source.
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


	// Save the new manifest
	if err := p.local.SaveManifest(remoteManifest); err != nil {
		result.Failed = append(result.Failed, fmt.Sprintf("save manifest: %v", err))
	}

	// Sync each track
	for _, trackEntry := range remoteManifest.Tracks {
		track, err := p.remote.GetTrack(ctx, trackEntry.ID)
		if err != nil {
			result.Failed = append(result.Failed, fmt.Sprintf("track %s: %v", trackEntry.ID, err))
			continue
		}
		if err := p.local.SaveTrack(trackEntry.ID, track); err != nil {
			result.Failed = append(result.Failed, fmt.Sprintf("save track %s: %v", trackEntry.ID, err))
			continue
		}

		// Sync katas for this track
		for _, stage := range track.Stages {
			for _, cat := range stage.Categories {
				for _, kataID := range cat.KataIDs {
					k, err := p.remote.GetKata(ctx, trackEntry.ID, kataID)
					if err != nil {
						result.Failed = append(result.Failed, fmt.Sprintf("kata %s/%s: %v", trackEntry.ID, kataID, err))
						continue
					}
					if err := p.local.SaveKata(trackEntry.ID, k); err != nil {
						result.Failed = append(result.Failed, fmt.Sprintf("save kata %s: %v", kataID, err))
					} else {
						result.Added++
					}
				}
			}
		}
	}

	p.mu.Lock()
	p.lastSync = result.SyncedAt
	p.remoteOK = true
	p.mu.Unlock()

	return result, nil
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
