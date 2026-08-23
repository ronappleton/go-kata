/**
 * API client for the GoKatas Go sidecar HTTP server.
 *
 * The sidecar listens on a random localhost port. We discover it via
 * Tauri's get_port command, or fall back to probing ports 9100-9200.
 */

import type { Track, KataDetail, RunResult, ProgressState, AppStatus } from "./types";

let API_BASE = "";

async function getBaseUrl(): Promise<string> {
  if (API_BASE) return API_BASE;
  // Retry up to 15 times with 1s delay — sidecar may take time to start
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
    // Wait before retrying
    await new Promise((r) => setTimeout(r, 1000));
  }
  throw new Error("Go sidecar not found on any port");
}

async function api<T>(path: string, init?: RequestInit): Promise<T> {
  const base = await getBaseUrl();
  const resp = await fetch(`${base}${path}`, {
    headers: { "Content-Type": "application/json" },
    ...init,
  });
  if (!resp.ok) {
    const body = await resp.json().catch(() => ({ error: resp.statusText }));
    throw new Error(body.error || resp.statusText);
  }
  return resp.json();
}

export const catalog = {
  get: () => api<Track>("/api/catalog"),
};

export const kata = {
  get: (id: string) => api<KataDetail>(`/api/kata/${id}`),
  save: (id: string, data: { code: string; tests: string; sourceFilename?: string; testFilename?: string }) =>
    api<{ status: string }>(`/api/kata/${id}/save`, {
      method: "POST",
      body: JSON.stringify(data),
    }),
  run: (id: string, data: { code: string; tests: string }) =>
    api<RunResult>(`/api/kata/${id}/run`, {
      method: "POST",
      body: JSON.stringify(data),
    }),
};

export const progress = {
  get: () => api<ProgressState>("/api/progress"),
};

export const syncApi = {
  trigger: () => api<{ status: string }>("/api/sync", { method: "POST" }),
  stream: async function* () {
    const base = await getBaseUrl();
    const resp = await fetch(`${base}/api/sync/stream`);
    const reader = resp.body?.getReader();
    if (!reader) return;
    const decoder = new TextDecoder();
    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      const text = decoder.decode(value);
      const lines = text.split("\n").filter((l) => l.startsWith("data: "));
      for (const line of lines) {
        try {
          yield JSON.parse(line.slice(6));
        } catch {
          // skip
        }
      }
    }
  },
};

export const lint = {
  check: (code: string, language: string) =>
    api<{ diagnostics: Array<{ line: number; col: number; endLine: number; endCol: number; message: string; isError: boolean }> }>("/api/lint", {
      method: "POST",
      body: JSON.stringify({ code, language }),
    }),
};

export const status = {
  get: () => api<AppStatus>("/api/status"),
};
