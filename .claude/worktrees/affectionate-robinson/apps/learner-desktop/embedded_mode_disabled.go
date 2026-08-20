//go:build !desktop_webview

package main

import "errors"

func embeddedModeAvailable() bool {
	return false
}

func openEmbeddedWindow(_ string) error {
	return errors.New("this build does not include native webview support; rebuild with: go build -tags desktop_webview ./apps/learner-desktop")
}
