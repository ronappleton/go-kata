# Learner Desktop

Desktop launcher for Learner Studio.

## Modes

- `auto` (default):
  - uses embedded native webview when built with `desktop_webview` tag
  - otherwise falls back to external app-window launch
- `embedded`: requires native webview build (`-tags desktop_webview`)
- `external`: forces browser/app-window launcher path

## Run

From repo root:

```bash
# default build (external launcher unless webview tag build)
go run ./apps/learner-desktop

# force external mode
go run ./apps/learner-desktop -mode external

# native embedded webview mode
go run -tags desktop_webview ./apps/learner-desktop -mode embedded
```

## Native webview dependency

Embedded mode uses `github.com/webview/webview_go`.

If first run reports missing module, run:

```bash
go get github.com/webview/webview_go@latest
```

Platform notes:
- macOS: uses native WebKit framework
- Windows: requires WebView2 runtime
- Linux: requires GTK/WebKitGTK development/runtime libraries
