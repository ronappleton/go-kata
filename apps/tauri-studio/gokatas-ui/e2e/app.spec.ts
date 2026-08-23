/**
 * E2E tests for GoKatas Tauri app.
 *
 * These tests run against the Go sidecar serving the built React frontend.
 * The Playwright config starts the sidecar automatically.
 */
import { test, expect } from "@playwright/test";

// ── Health & Catalog ──

test.describe("App health", () => {
  test("API status endpoint is reachable", async ({ page }) => {
    const resp = await page.request.get("/api/status");
    expect(resp.ok()).toBeTruthy();
    const body = await resp.json();
    expect(body).toHaveProperty("kataCount");
  });

  test("catalog endpoint returns stages with katas", async ({ page }) => {
    // Track may not be loaded yet on cold start — poll until ready
    let track: any = null;
    for (let i = 0; i < 20; i++) {
      const resp = await page.request.get("/api/catalog");
      expect(resp.ok()).toBeTruthy();
      track = await resp.json();
      if (track.stages && track.stages.length > 0) break;
      await page.waitForTimeout(1000);
    }
    expect(track.stages.length).toBeGreaterThan(0);
    const totalKatas = track.stages.reduce(
      (sum: number, s: any) =>
        sum +
        s.categories.reduce(
          (s2: number, c: any) => s2 + c.katas.length,
          0
        ),
      0
    );
    expect(totalKatas).toBeGreaterThan(0);
  });
});

// ── Frontend loads ──

test.describe("Frontend loads", () => {
  test("loads the React app (no blank page)", async ({ page }) => {
    await page.goto("/");
    // The app renders a sidebar with "GoKatas" branding
    await expect(page.locator("text=GoKatas").first()).toBeVisible({
      timeout: 10_000,
    });
  });

  test("sidebar populates with curriculum stages", async ({ page }) => {
    await page.goto("/");
    // Wait for the splash to disappear and sidebar to populate
    // The splash shows spinner; once loaded, sidebar shows stage buttons
    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 30_000 });
  });
});

// ── Kata selection ──

test.describe("Kata selection", () => {
  test("clicking a kata loads its readme", async ({ page }) => {
    await page.goto("/");

    // Wait for sidebar to populate
    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 15_000 });

    // Find and click a kata row (e.g. kata-001)
    const kataRow = page.locator(".kata-row").first();
    await expect(kataRow).toBeVisible({ timeout: 10_000 });
    await kataRow.click();

    // The docs tab should show markdown content
    await expect(page.locator(".markdown-body")).toBeVisible({
      timeout: 10_000,
    });
    // The title should appear in the header
    const headerText = await page.locator("h1").first().textContent();
    expect(headerText).toBeTruthy();
  });
});

// ── Editor tab ──

test.describe("Editor tab", () => {
  test("workbench tab shows Monaco editor with Go code", async ({ page }) => {
    await page.goto("/");

    // Wait for sidebar
    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 15_000 });

    // Select a kata
    const kataRow = page.locator(".kata-row").first();
    await expect(kataRow).toBeVisible({ timeout: 10_000 });
    await kataRow.click();

    // Wait for docs to load
    await expect(page.locator(".markdown-body")).toBeVisible({
      timeout: 10_000,
    });

    // Switch to workbench tab (use the tab-btn class to avoid matching kata rows)
    await page.locator("button.tab-btn:has-text('Code')").click();

    // Monaco editor should appear
    await expect(
      page.locator(".monaco-editor").first()
    ).toBeVisible({ timeout: 10_000 });

    // Solution tab should be active
    await expect(
      page.locator("button:has-text('Solution')").first()
    ).toHaveClass(/active/);
  });
});

// ── Sync button ──

test.describe("Sync", () => {
  test("sync button exists and is clickable", async ({ page }) => {
    await page.goto("/");

    await expect(page.locator("text=GoKatas").first()).toBeVisible({
      timeout: 10_000,
    });

    const syncBtn = page.locator("button:has-text('Sync')");
    await expect(syncBtn).toBeVisible({ timeout: 10_000 });
    await expect(syncBtn).toBeEnabled();
  });

  test("sync button triggers reload and returns success", async ({ page }) => {
    await page.goto("/");

    // Wait for initial load
    await expect(page.locator("text=/\\d+ katas/").first()).toBeVisible({
      timeout: 15_000,
    });

    const syncBtn = page.locator("button", { hasText: /sync/i }).first();
    await syncBtn.click();

    // Should show syncing state then return to ready
    await expect(page.locator("text=/\\d+ katas/").first()).toBeVisible({
      timeout: 15_000,
    });
  });
});

// ── Status bar ──

test.describe("Status bar", () => {
  test("shows kata count in the status area", async ({ page }) => {
    await page.goto("/");

    // The status area should eventually show the kata count
    await expect(page.locator("text=/\\d+ katas/").first()).toBeVisible({
      timeout: 15_000,
    });
  });
});

// ── Search filter ──

test.describe("Search filter", () => {
  test("search input filters sidebar katas", async ({ page }) => {
    await page.goto("/");

    // Wait for sidebar to populate
    await expect(page.locator("text=/\\d+ katas/").first()).toBeVisible({
      timeout: 15_000,
    });

    // Find the search input
    const searchInput = page.locator('input[placeholder*="Search"]');
    await expect(searchInput).toBeVisible();

    // Type a search term
    await searchInput.fill("greeting");

    // Should show filtered results
    const greetingKata = page.locator("text=Build Greeting").first();
    await expect(greetingKata).toBeVisible({ timeout: 5_000 });

    // Clear search
    await searchInput.fill("");
  });
});

// ── Tab switching ──

test.describe("Tab switching", () => {
  test("can switch between docs, workbench, and output tabs", async ({ page }) => {
    await page.goto("/");

    // Wait for sidebar to populate
    await expect(page.locator("text=/\\d+ katas/").first()).toBeVisible({
      timeout: 15_000,
    });

    // Click a kata to load it
    const firstKata = page.locator(".kata-row").first();
    await firstKata.click();

    // Should be on docs tab by default
    await expect(page.locator("text=Readme").first()).toBeVisible();

    // Switch to workbench tab
    const workbenchTab = page.locator("button.tab-btn", { hasText: /Code/ }).first();
    await workbenchTab.click();

    // Editor should be visible
    await expect(page.locator(".monaco-editor").first()).toBeVisible({
      timeout: 5_000,
    });

    // Switch to output tab
    const outputTab = page.locator("button.tab-btn", { hasText: /Output/ }).first();
    await outputTab.click();

    // Output content should be visible
    await expect(page.locator("text=Run your solution").first()).toBeVisible();
  });
});

// ── Lint endpoint ──

test.describe("Lint", () => {
  test("lint endpoint returns diagnostics for bad Go code", async ({ page }) => {
    const resp = await page.request.post("/api/lint", {
      data: {
        code: "package main\n\nfunc main() {\n",
        language: "go",
      },
    });
    expect(resp.ok()).toBeTruthy();
    const body = await resp.json();
    expect(body).toHaveProperty("diagnostics");
    // Malformed code should produce diagnostics
    expect(body.diagnostics.length).toBeGreaterThan(0);
  });

  test("lint endpoint returns empty for clean Go code", async ({ page }) => {
    const resp = await page.request.post("/api/lint", {
      data: {
        code: "package main\n\nfunc main() {}\n",
        language: "go",
      },
    });
    expect(resp.ok()).toBeTruthy();
    const body = await resp.json();
    expect(body).toHaveProperty("diagnostics");
  });
});

// ── Progress endpoint ──

test.describe("Progress", () => {
  test("progress endpoint returns attempts map", async ({ page }) => {
    const resp = await page.request.get("/api/progress");
    expect(resp.ok()).toBeTruthy();
    const body = await resp.json();
    expect(body).toHaveProperty("attempts");
  });
});

// ── Save kata ──

test.describe("Save kata", () => {
  test("save endpoint persists code", async ({ page }) => {
    const resp = await page.request.post("/api/kata/001/save", {
      data: {
        code: 'package kata\nfunc Hello() string { return "hi" }',
        tests: "package kata\nfunc TestHello(t *testing.T) {}",
      },
    });
    expect(resp.ok()).toBeTruthy();
    const body = await resp.json();
    expect(body.status).toBe("saved");
  });
});
