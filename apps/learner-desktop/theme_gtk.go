//go:build gtk4

package main

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func installTheme() {
	gtk.SettingsGetDefault().SetObjectProperty("gtk-application-prefer-dark-theme", true)
	provider := gtk.NewCSSProvider()
	provider.LoadFromData(themeCSS)
	if display := gdk.DisplayGetDefault(); display != nil {
		gtk.StyleContextAddProviderForDisplay(display, provider, gtk.STYLE_PROVIDER_PRIORITY_USER)
	}
}

const themeCSS = `
/* ==== GoKatas design tokens ==== */
@define-color bg          #0d1117;
@define-color surface     #151b24;
@define-color surface_hi  #1c2430;
@define-color border      #26303e;
@define-color text        #e7ecf3;
@define-color text_dim    #9aa4b4;
@define-color text_faint  #6b7686;
@define-color accent      #14b8a6;
@define-color accent_hi   #2dd4bf;
@define-color accent_bg   rgba(20,184,166,0.14);
@define-color success     #34d399;
@define-color danger      #f87171;
@define-color warning     #fbbf24;

window.gokatas {
  background-color: @bg;
  color: @text;
}

/* ==== Header ==== */
.header {
  background-color: @surface;
  border-bottom: 1px solid @border;
  padding: 12px 20px;
}
.brand {
  font-size: 21px;
  font-weight: 800;
  color: @text;
}
.brand-sub {
  font-size: 12px;
  color: @text_dim;
  margin-top: 1px;
}
.status-pill {
  background-color: @surface_hi;
  color: @text_dim;
  border: 1px solid @border;
  border-radius: 999px;
  padding: 5px 12px;
  font-size: 12px;
}
.status-pill.ok {
  color: @success;
  border-color: rgba(52,211,153,0.35);
}
.status-pill.warn {
  color: @warning;
  border-color: rgba(251,191,36,0.35);
}

/* ==== Buttons ==== */
button {
  border-radius: 8px;
  padding: 7px 14px;
  min-height: 34px;
  background-color: @surface_hi;
  color: @text;
  border: 1px solid @border;
  font-weight: 600;
}
button:hover {
  background-color: @border;
  color: @text;
}
button:active {
  background-color: @surface;
}
button:disabled {
  background-color: @surface;
  color: @text_faint;
  border-color: @border;
}
button.suggested-action {
  background-color: @accent;
  color: #071014;
  border: 1px solid transparent;
  font-weight: 700;
}
button.suggested-action:hover {
  background-color: @accent_hi;
}
button.suggested-action:disabled {
  background-color: @surface_hi;
  color: @text_faint;
  border: 1px solid @border;
}

/* ==== Sidebar ==== */
.sidebar {
  background-color: @surface;
  border-right: 1px solid @border;
}
.sidebar-category {
  color: @text_faint;
  font-size: 11px;
  font-weight: 700;
  margin-top: 16px;
  margin-bottom: 4px;
  padding-left: 10px;
}

/* Stage headers */
.stage-label {
  color: @text;
  font-size: 12px;
  font-weight: 800;
  padding-left: 10px;
}

/* Level badges */
.level-badge {
  font-size: 9px;
  font-weight: 700;
  padding: 2px 8px;
  border-radius: 999px;
  color: #071014;
}
.level-junior {
  background-color: @success;
}
.level-mid {
  background-color: @accent;
}
.level-senior {
  background-color: @warning;
}
.level-lead {
  background-color: @danger;
}

/* Progress bars */
.stage-progress {
  min-height: 4px;
  min-width: 60px;
  margin-left: 8px;
}
.stage-progress progress,
.stage-progress trough {
  min-height: 4px;
  border-radius: 2px;
}
.stage-progress progress {
  background-color: @accent;
}
.stage-progress trough {
  background-color: @border;
}

.cat-progress {
  min-height: 3px;
  min-width: 40px;
  margin-left: 4px;
}
.cat-progress progress,
.cat-progress trough {
  min-height: 3px;
  border-radius: 2px;
}
.cat-progress progress {
  background-color: @accent;
  opacity: 0.6;
}
.cat-progress trough {
  background-color: @border;
}

/* Kata rows */
.kata-row {
  background-color: transparent;
  border: 1px solid transparent;
  border-radius: 8px;
  padding: 7px 10px;
  min-height: 0;
  font-weight: 500;
  color: @text_dim;
}
.kata-row:hover {
  background-color: @surface_hi;
  color: @text;
}
.kata-row.active {
  background-color: @accent_bg;
  color: @text;
  border-color: rgba(20,184,166,0.28);
}
.kata-row.locked {
  opacity: 0.4;
}
.kata-dot {
  font-size: 10px;
}

/* ==== Tabs (StackSwitcher) ==== */
stackswitcher {
  background-color: @surface;
  border: 1px solid @border;
  border-radius: 10px;
  padding: 4px;
}
stackswitcher > button {
  border-radius: 7px;
  border: 1px solid transparent;
  background-color: transparent;
  padding: 6px 14px;
  min-height: 28px;
  font-weight: 600;
  color: @text_dim;
}
stackswitcher > button:hover {
  background-color: @surface_hi;
  color: @text;
}
stackswitcher > button:checked {
  background-color: @surface_hi;
  color: @text;
  border-color: @border;
}

/* ==== Surfaces ==== */
.panel {
  background-color: @surface;
  border: 1px solid @border;
  border-radius: 12px;
  padding: 14px;
}
.panel-title {
  font-size: 13px;
  font-weight: 700;
  color: @text_dim;
  margin-bottom: 8px;
}
.heading {
  font-size: 18px;
  font-weight: 800;
  color: @text;
}
.heading-sub {
  font-size: 13px;
  color: @text_dim;
}

/* ==== Editors & console ==== */
textview.mono,
textview.mono text {
  font-family: "JetBrains Mono", "Ubuntu Mono", "DejaVu Sans Mono", monospace;
  font-size: 13px;
}
.editor {
  background-color: @bg;
  border: 1px solid @border;
  border-radius: 10px;
}
.editor text {
  background-color: transparent;
  color: @text;
}
.console {
  background-color: #0a0d12;
  border: 1px solid @border;
  border-radius: 10px;
}
.console text {
  background-color: transparent;
  color: @text_dim;
}

/* ==== Body copy ==== */
.body {
  font-size: 14px;
  color: @text_dim;
}
.footnote {
  font-size: 12px;
  color: @text_faint;
}

/* ==== Flashcards ==== */
.flash-side {
  font-size: 12px;
  font-weight: 700;
  color: @accent;
  text-transform: uppercase;
  letter-spacing: 1px;
}
.confidence-again {
  background-color: rgba(248,113,113,0.2);
  color: @danger;
  border-color: rgba(248,113,113,0.3);
}
.confidence-again:hover {
  background-color: rgba(248,113,113,0.35);
}
.confidence-hard {
  background-color: rgba(251,191,36,0.2);
  color: @warning;
  border-color: rgba(251,191,36,0.3);
}
.confidence-hard:hover {
  background-color: rgba(251,191,36,0.35);
}
.confidence-good {
  background-color: rgba(20,184,166,0.2);
  color: @accent;
  border-color: rgba(20,184,166,0.3);
}
.confidence-good:hover {
  background-color: rgba(20,184,166,0.35);
}
.confidence-easy {
  background-color: rgba(52,211,153,0.2);
  color: @success;
  border-color: rgba(52,211,153,0.3);
}
.confidence-easy:hover {
  background-color: rgba(52,211,153,0.35);
}

/* ==== Quiz feedback ==== */
.quiz-feedback {
  color: @accent;
  font-weight: 600;
}

/* ==== Break reminder (ADHD) ==== */
.break-reminder {
  color: @warning;
  font-size: 13px;
  font-weight: 600;
  padding: 12px 10px;
  margin-top: 8px;
  background-color: rgba(251,191,36,0.1);
  border-radius: 8px;
  border: 1px solid rgba(251,191,36,0.2);
}

/* ==== Mode buttons ==== */
.mode-btn {
  font-size: 11px;
  padding: 4px 10px;
  min-height: 24px;
  border-radius: 6px;
  font-weight: 600;
}
.mode-btn.active {
  background-color: @accent_bg;
  color: @accent;
  border-color: rgba(20,184,166,0.3);
}

/* ==== Separator ==== */
paned > separator {
  background-color: @border;
  min-width: 1px;
  min-height: 1px;
}

/* ==== Scrollbars ==== */
scrolledwindow undershoot.top,
scrolledwindow undershoot.bottom,
scrolledwindow undershoot.left,
scrolledwindow undershoot.right {
  background-color: transparent;
}
`
