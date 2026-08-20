//go:build gtk4

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/ronappleton/golang-katas-1-100/internal/updater"
)

// initUpdateCheck starts a background check for app updates.
// When an update is found, it shows a notification badge on the header.
func (n *nativeApp) initUpdateCheck() {
	mgr := updater.NewManager(version)

	go func() {
		// Wait a few seconds before checking (don't slow down startup)
		time.Sleep(5 * time.Second)

		info, err := mgr.CheckForUpdate(context.Background())
		if err != nil {
			log.Printf("update check failed: %v", err)
			return
		}
		if info == nil {
			return // Up to date
		}

		log.Printf("update available: %s (current: %s)", info.Version, mgr.CurrentVersion())

		// Show update notification on the GTK thread
		glib.MainContextDefault().InvokeFull(0, func() bool {
			n.showUpdateAvailable(mgr, info)
			return false
		})
	}()
}

// showUpdateAvailable adds an update badge to the header and shows a dialog.
func (n *nativeApp) showUpdateAvailable(mgr *updater.Manager, info *updater.UpdateInfo) {
	// Add update button to header
	updateBtn := gtk.NewButtonWithLabel(fmt.Sprintf("Update to %s", info.Version))
	updateBtn.AddCSSClass("suggested-action")
	updateBtn.AddCSSClass("update-btn")
	updateBtn.SetTooltipText(fmt.Sprintf("GoKatas %s is available\n%s", info.Version, info.Body))
	updateBtn.ConnectClicked(func() {
		n.applyUpdate(mgr, info)
	})

	// Insert before the status label in the header
	n.headerBox.Append(updateBtn)
}

// applyUpdate downloads and installs the update.
func (n *nativeApp) applyUpdate(mgr *updater.Manager, info *updater.UpdateInfo) {
	n.setStatus(fmt.Sprintf("Downloading update %s…", info.Version))
	n.runButton.SetSensitive(false)
	n.saveButton.SetSensitive(false)

	go func() {
		newPath, err := mgr.DownloadAndReplace(context.Background(), info)
		if err != nil {
			glib.MainContextDefault().InvokeFull(0, func() bool {
				n.setStatus(fmt.Sprintf("Update failed: %v", err))
				n.runButton.SetSensitive(n.selected.EvaluatorStatus == "ready")
				n.saveButton.SetSensitive(true)
				return false
			})
			return
		}

		glib.MainContextDefault().InvokeFull(0, func() bool {
		// Show restart dialog
		dialog := gtk.NewMessageDialog(
			&n.window.Window,
			gtk.DialogModal,
			gtk.MessageInfo,
			gtk.ButtonsNone,
		)
		dialog.SetTitle("Update Installed")
		dialog.SetMarkup(fmt.Sprintf(
			"<b>GoKatas %s</b> has been downloaded successfully.\n\nThe app will restart to apply the update.",
			info.Version,
		))

		dialog.AddButton("Restart Now", int(gtk.ResponseAccept))
		dialog.AddButton("Later", int(gtk.ResponseCancel))

		dialog.ConnectResponse(func(response int) {
			dialog.Destroy()
			if gtk.ResponseType(response) == gtk.ResponseAccept {
				n.setStatus("Restarting…")
				// Restart the process
				if err := updater.Restart(newPath); err != nil {
					n.setStatus(fmt.Sprintf("Restart failed: %v", err))
				}
			}
		})

		dialog.Show()
			return false
		})
	}()
}
