/**
 * API client for GoKatas.
 *
 * Uses Tauri IPC (invoke) for all backend calls — no HTTP sidecar probing.
 * Falls back to HTTP fetch if running outside Tauri (e.g. Playwright tests).
 */

import type { Track, KataDetail, RunResult, ProgressState, AppStatus } from "./types";

// Check if we're running inside Tauri
const isTauri = typeof window !== "undefined" && "__TAURI_INTERNALS__" in window;

async function invoke<T>(cmd: string, args?: Record<string, unknown>): Promise<T> {
  if (isTauri) {
    const { invoke: tauriInvoke } = await import("@tauri-apps/api/core");
    return tauriInvoke<T>(cmd, args);
  }
  // Fallback: HTTP to sidecar (for Playwright / dev mode)
  return httpFallback<T>(cmd, args);
}

// ── HTTP fallback for tests/dev ──

let API_BASE = "";

async function getBaseUrl(): Promise<string> {
  if (API_BASE) return API_BASE;
  for (let attempt = 0; attempt < 15; attempt++) {
    for (let port = 9100; port <= 9200; port++) {
      try {
        const resp = await fetch(`http://127.0.0.1:${port}/api/status`, {
          signal: AbortSignal.timeout(1000),
        });
        if (resp.ok) {
          API_BASE = `http://127.0.0.1:${port}`;
          return API_BASE;
        }
      } catch {
        // continue
      }
    }
    await new Promise((r) => setTimeout(r, 1000));
  }
  throw new Error("Go sidecar not found on any port");
}

async function httpFallback<T>(cmd: string, args?: Record<string, unknown>): Promise<T> {
  const base = await getBaseUrl();
  // Map Tauri command names back to HTTP endpoints
  const routeMap: Record<string, { method: string; path: string }> = {
    get_catalog: { method: "GET", path: "/api/catalog" },
    get_kata: { method: "GET", path: "" },
    save_kata: { method: "POST", path: "" },
    run_kata: { method: "POST", path: "" },
    get_progress: { method: "GET", path: "/api/progress" },
    get_status: { method: "GET", path: "/api/status" },
    sync_content: { method: "POST", path: "/api/sync" },
    lint_code: { method: "POST", path: "/api/lint" },
  };
  const route = routeMap[cmd] || { method: "GET", path: "/api/status" };
  let path = route.path;
  if (args?.id) path = `/api/kata/${args.id}`;
  if (cmd === "save_kata") path = `/api/kata/${args?.id}/save`;
  if (cmd === "run_kata") path = `/api/kata/${args?.id}/run`;

  const resp = await fetch(`${base}${path}`, {
    method: route.method,
    headers: { "Content-Type": "application/json" },
    body: route.method === "POST" ? JSON.stringify(args) : undefined,
  });
  if (!resp.ok) {
    const body = await resp.json().catch(() => ({ error: resp.statusText }));
    throw new Error(body.error || resp.statusText);
  }
  return resp.json();
}

// ── Public API ──

export const catalog = {
  get: () => invoke<Track>("get_catalog"),
};

export const kata = {
  get: (id: string) => invoke<KataDetail>("get_kata", { id }),
  save: (id: string, data: { code: string; tests: string; sourceFilename?: string; testFilename?: string }) =>
    invoke<{ status: string }>("save_kata", { id, ...data }),
  run: (id: string, data: { code: string; tests: string; mode?: string }) =>
    invoke<RunResult>("run_kata", { id, ...data, mode: data.mode || "evaluate" }),
};

export const progress = {
  get: () => invoke<ProgressState>("get_progress"),
};

export const syncApi = {
  trigger: () => invoke<{ status: string }>("sync_content"),
};

export const lint = {
  check: (code: string, language: string) =>
    invoke<{ diagnostics: Array<{ line: number; col: number; endLine: number; endCol: number; message: string; isError: boolean }> }>(
      "lint_code",
      { code, language }
    ),
};

export const status = {
  get: () => invoke<AppStatus>("get_status"),
};
