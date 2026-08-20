//go:build gtk4

package main

import (
	"sync"
	"testing"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func TestThemeCSSParses(t *testing.T) {
	if !gtk.InitCheck() {
		t.Skip("no display available to initialize GTK")
	}

	provider := gtk.NewCSSProvider()
	var mu sync.Mutex
	var parseErrors []string
	provider.ConnectParsingError(func(_ *gtk.CSSSection, err error) {
		mu.Lock()
		defer mu.Unlock()
		parseErrors = append(parseErrors, err.Error())
	})

	provider.LoadFromData(themeCSS)

	mu.Lock()
	defer mu.Unlock()
	if len(parseErrors) > 0 {
		t.Fatalf("theme CSS failed to parse: %v", parseErrors)
	}
}
