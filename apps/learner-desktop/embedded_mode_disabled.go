//go:build !desktop_webview

package main

import "errors"

func embeddedModeAvailable() bool {
	return false
}

func openEmbeddedWindow(_ string) error {
	return errors.New("native webview support is not enabled")
}
