//go:build desktop_webview

package main

import (
	"errors"

	webview "github.com/webview/webview_go"
)

func embeddedModeAvailable() bool {
	return true
}

func openEmbeddedWindow(url string) error {
	window := webview.New(true)
	if window == nil {
		return errors.New("failed to initialize native webview")
	}
	defer window.Destroy()

	window.SetTitle("Learner Studio")
	window.SetSize(1440, 900, webview.HintNone)
	window.Navigate(url)
	window.Run()
	return nil
}
