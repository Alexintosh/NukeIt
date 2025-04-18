
# 🧹 gocleaner - CLI AppCleaner Alternative (Go)

## 📌 Purpose

`gocleaner` is a command-line tool written in Go for macOS that helps users fully uninstall applications by removing the main app bundle **and** associated files (like caches, preferences, logs, etc.).

---

## 🛠 Features

- Detects and deletes app support files across macOS system paths.
- Reads app bundle ID from `.app` files.
- Provides dry-run, verbose, and force-delete options.
- Confirms file deletion before removing anything (unless `--force` is used).

---

## 💻 Command Usage

```bash
gocleaner uninstall <AppName> [--dry-run] [--force] [--verbose]
```

### Example
```bash
gocleaner uninstall Spotify
```

### Flags

- `--dry-run` – Show what would be deleted, but don’t delete.
- `--force` – Skip confirmation prompt and delete files immediately.
- `--verbose` – Show detailed scanning and deletion info.

---

## 📂 macOS Paths to Scan

Files related to the app (based on name or bundle ID) will be searched in:

- `~/Library/Application Support/`
- `~/Library/Preferences/`
- `~/Library/Caches/`
- `~/Library/Logs/`
- `~/Library/Containers/`
- `~/Library/Saved Application State/`

---

## 🧠 How It Works

1. **Input**: App name is provided via CLI.
2. **Find `.app` bundle**:
   - Search `/Applications/` and `~/Applications/`.
   - If found, open `Info.plist` to get the bundle ID.
3. **If not found**:
   - Use the app name to scan for related files.
4. **Search associated folders** using:
   - Bundle ID match
   - App name (case-insensitive)
5. **Display results** for user review.
6. **Prompt for confirmation**, unless `--force` is used.
7. **Delete files** and show status.

---

## ⚠️ Caution & Safeguards

- Prevent deletion of system-critical files.
- Never touch user documents (e.g., files in `Documents/`, `Downloads/`).
- Always request confirmation unless `--force` is passed.

---

## 📋 Logging and Errors

- Use clear output for what’s happening.
- Verbose mode prints full scan paths and deletion results.
- Handle permission errors gracefully (e.g., warn user, skip file).

---

## 🏁 Summary

`gocleaner` gives power users and developers a fast, scriptable way to cleanly uninstall macOS apps via the terminal using Go.
