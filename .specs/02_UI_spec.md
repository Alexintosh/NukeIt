
# ✨ gocleaner v2 - Enhanced CLI AppCleaner with TUI

## 🚀 Goal

Enhance the existing `gocleaner` CLI app with a terminal user interface (TUI) using [Charmbracelet Bubbles](https://github.com/charmbracelet/bubbles). This new version will provide a more interactive and visual experience for scanning, reviewing, and cleaning up macOS app-related files.

---

## 🔧 New Features

### 1. 🔄 Spinner During File Scanning

- **Purpose:** Indicate to the user that the tool is actively searching for related files.
- **Implementation:** Use the `spinner` Bubble during the scanning process.
- **Behavior:**
  - Spinner appears after entering the app name.
  - Displays status messages like `Scanning /Library/Preferences/...`
  - Ends with either a success message or "no related files found".

### 2. 📊 Progress Bar During File Deletion

- **Purpose:** Give users a visual cue of progress during file deletions.
- **Implementation:** Use the `progress` Bubble.
- **Behavior:**
  - Each file deletion updates the progress bar.
  - Errors or skipped files are shown in a summary afterward.

### 3. 🗂 Interactive File Review & Exclusion UI

- **Purpose:** Allow users to review and selectively exclude files from deletion.
- **Implementation:** Use a list Bubble with keyboard navigation.
- **Behavior:**
  - Files are listed with checkboxes.
  - Navigation: `↑ ↓` to move, `Space` to toggle selection.
  - Press `Enter` to confirm and proceed with deletion.
  - Optional: Add `a` to select all, `n` to select none.

---

## 🛠 Command Usage (Same as v1)

```bash
gocleaner uninstall <AppName> [--dry-run] [--force] [--verbose]
```

### Additional Interactive Steps in v2

- User enters app name.
- Spinner shows while scanning.
- Interactive file list appears for selection.
- User confirms files.
- Progress bar appears while deleting.

---

## 📦 Libraries Required

- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss) (for styling)
- [bubbletea](https://github.com/charmbracelet/bubbletea) (core TUI framework)

---

## 📋 Summary

This version of `gocleaner` introduces a sleek terminal UI powered by Charmbracelet's Bubbletea ecosystem. It improves usability, visibility, and control over which files get removed — all within a fast, keyboard-driven interface.
