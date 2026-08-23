import { defineConfig } from "@playwright/test";
import { fileURLToPath } from "url";
import path from "path";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const SIDECAST_PORT = 9100;
const FRONTEND_DIR = path.resolve(__dirname, "dist");
const PROJECT_ROOT = path.resolve(__dirname, "../../..");

export default defineConfig({
  testDir: "./e2e",
  fullyParallel: false,
  retries: 1,
  use: {
    baseURL: `http://127.0.0.1:${SIDECAST_PORT}`,
    headless: true,
    screenshot: "only-on-failure",
    trace: "on-first-retry",
    timeout: 60_000,
  },
  webServer: {
    command: `GOKATAS_PORT=${SIDECAST_PORT} GOKATAS_FRONTEND_DIR="${FRONTEND_DIR}" go run ./apps/tauri-studio/gokatas-ui/src-tauri/sidecar/main.go`,
    url: `http://127.0.0.1:${SIDECAST_PORT}/api/status`,
    reuseExistingServer: true,
    cwd: PROJECT_ROOT,
  },
});
