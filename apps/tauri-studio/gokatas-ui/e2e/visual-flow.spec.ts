/**
 * Visual flow tests for GoKatas.
 *
 * Captures screenshots at every step of the user journey:
 *   1. App launch / splash screen
 *   2. Sidebar populated with curriculum
 *   3. Click a kata → readme renders
 *   4. Switch to Code tab → editor visible
 *   5. Run in sandbox → output shows result
 *
 * Screenshots saved to e2e/screenshots/ for manual review.
 */
import { test, expect } from "@playwright/test";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const SCREENSHOT_DIR = path.resolve(__dirname, "screenshots");

test.describe("Visual flow: open → kata → editor → run", () => {
  test("full user journey with screenshots at every step", async ({
    page,
  }) => {
    // ── Step 1: App launch / splash screen ──
    await page.goto("/");
    // Capture the splash screen immediately (may be brief)
    await page.waitForTimeout(1500);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "01-app-launch-splash.png"),
      fullPage: false,
    });

    // ── Step 2: Sidebar populated with curriculum ──
    // Wait for JUNIOR stage to appear (means sidebar loaded)
    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 30_000 });

    // Give the UI a moment to fully settle
    await page.waitForTimeout(500);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "02-sidebar-curriculum.png"),
      fullPage: false,
    });

    // ── Step 3: Click a kata → readme renders ──
    const kataRow = page.locator(".kata-row").first();
    await expect(kataRow).toBeVisible({ timeout: 10_000 });
    await kataRow.click();

    // Wait for markdown readme to appear
    await expect(page.locator(".markdown-body")).toBeVisible({
      timeout: 10_000,
    });
    // Give markdown time to render fully
    await page.waitForTimeout(500);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "03-kata-readme-rendered.png"),
      fullPage: false,
    });

    // ── Step 4: Switch to Code tab → editor visible ──
    const codeTab = page
      .locator("button.tab-btn", { hasText: /Code/ })
      .first();
    await codeTab.click();

    // Wait for Monaco editor to load
    await expect(page.locator(".monaco-editor").first()).toBeVisible({
      timeout: 15_000,
    });
    // Give Monaco time to render syntax highlighting
    await page.waitForTimeout(1000);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "04-code-editor-monaco.png"),
      fullPage: false,
    });

    // ── Step 5: Type some code in the editor ──
    // Click inside the Monaco editor to focus it
    const editor = page.locator(".monaco-editor .view-lines").first();
    await editor.click();
    await page.waitForTimeout(300);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "05-editor-focused.png"),
      fullPage: false,
    });

    // ── Step 6: Click Run in sandbox ──
    const runBtn = page.locator("button:has-text('Run')").first();
    await expect(runBtn).toBeVisible();

    // Check if Run is enabled (depends on evaluator status)
    const isDisabled = await runBtn.isDisabled();
    if (!isDisabled) {
      await runBtn.click();

      // Wait for the output tab to show
      await page.waitForTimeout(1000);
      await page.screenshot({
        path: path.join(SCREENSHOT_DIR, "06-running-in-sandbox.png"),
        fullPage: false,
      });

      // Wait for run to complete (up to 60s for sandbox execution)
      // The output area should show pass/fail or error
      await page.waitForTimeout(5000);
      await page.screenshot({
        path: path.join(SCREENSHOT_DIR, "07-run-complete.png"),
        fullPage: false,
      });
    } else {
      // Evaluator not ready — capture that state
      await page.screenshot({
        path: path.join(SCREENSHOT_DIR, "06-run-disabled-no-evaluator.png"),
        fullPage: false,
      });
    }

    // ── Step 7: Switch back to Readme tab ──
    const readmeTab = page
      .locator("button.tab-btn", { hasText: /Readme/ })
      .first();
    await readmeTab.click();
    await expect(page.locator(".markdown-body")).toBeVisible({
      timeout: 5_000,
    });
    await page.waitForTimeout(300);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "08-back-to-readme.png"),
      fullPage: false,
    });
  });

  test("search filter flow with screenshot", async ({ page }) => {
    await page.goto("/");

    // Wait for sidebar to load
    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 30_000 });
    await page.waitForTimeout(500);

    // ── Step 1: Search for a kata ──
    const searchInput = page.locator('input[placeholder*="Search"]');
    await expect(searchInput).toBeVisible();
    await searchInput.fill("greeting");
    await page.waitForTimeout(500);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "09-search-filtered.png"),
      fullPage: false,
    });

    // ── Step 2: Click filtered result ──
    const filteredKata = page.locator(".kata-row").first();
    if (await filteredKata.isVisible()) {
      await filteredKata.click();
      await expect(page.locator(".markdown-body")).toBeVisible({
        timeout: 10_000,
      });
      await page.waitForTimeout(500);
      await page.screenshot({
        path: path.join(SCREENSHOT_DIR, "10-filtered-kata-selected.png"),
        fullPage: false,
      });
    }
  });

  test("output tab with empty state", async ({ page }) => {
    await page.goto("/");

    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 30_000 });

    // Select a kata
    const kataRow = page.locator(".kata-row").first();
    await kataRow.click();
    await expect(page.locator(".markdown-body")).toBeVisible({
      timeout: 10_000,
    });

    // Switch to output tab (before running anything)
    const outputTab = page
      .locator("button.tab-btn", { hasText: /Output/ })
      .first();
    await outputTab.click();
    await page.waitForTimeout(500);
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "11-output-empty-state.png"),
      fullPage: false,
    });
  });

  test("workbench with description banner", async ({ page }) => {
    await page.goto("/");

    await expect(
      page.locator("button:has-text('JUNIOR')").first()
    ).toBeVisible({ timeout: 30_000 });

    // Select a kata
    const kataRow = page.locator(".kata-row").first();
    await kataRow.click();
    await expect(page.locator(".markdown-body")).toBeVisible({
      timeout: 10_000,
    });

    // Switch to code tab
    const codeTab = page
      .locator("button.tab-btn", { hasText: /Code/ })
      .first();
    await codeTab.click();

    await expect(page.locator(".monaco-editor").first()).toBeVisible({
      timeout: 15_000,
    });
    await page.waitForTimeout(1000);

    // Capture the workbench showing the description banner + editor + tabs
    await page.screenshot({
      path: path.join(SCREENSHOT_DIR, "12-workbench-description-banner.png"),
      fullPage: false,
    });
  });
});
