# 🧹 gocleaner - CLI AppCleaner Alternative for macOS

`gocleaner` is a command-line tool written in Go for macOS that helps users fully uninstall applications by removing the main app bundle **and** associated files (like caches, preferences, logs, etc.).

## 📌 Purpose

Most macOS applications leave behind files after you drag them to the Trash. `gocleaner` helps you find and remove these leftover files, keeping your system clean.

## 🛠 Features

- Detects and deletes app support files across macOS system paths
- Reads app bundle ID from `.app` files
- Provides dry-run, verbose, and force-delete options
- Confirms file deletion before removing anything (unless `--force` is used)

## 💻 Installation

### From source

```bash
git clone https://github.com/alexintosh/gocleaner.git
cd gocleaner
go build -o gocleaner cmd/gocleaner/main.go
```

Then move the binary to your PATH:

```bash
mv gocleaner /usr/local/bin/
```

## 💻 Usage

```bash
gocleaner uninstall <AppName> [--dry-run] [--force] [--verbose]
```

### Example

```bash
gocleaner uninstall Spotify
```

### Flags

- `--dry-run` – Show what would be deleted, but don't delete.
- `--force` – Skip confirmation prompt and delete files immediately.
- `--verbose` – Show detailed scanning and deletion info.

## 📂 macOS Paths Scanned

Files related to the app (based on name or bundle ID) will be searched in:

- `~/Library/Application Support/`
- `~/Library/Preferences/`
- `~/Library/Caches/`
- `~/Library/Logs/`
- `~/Library/Containers/`
- `~/Library/Saved Application State/`

## ⚠️ Caution & Safeguards

- The tool prevents deletion of system-critical files
- User documents (e.g., files in `Documents/`, `Downloads/`) are never touched
- Always requests confirmation unless `--force` is passed

## 📝 License

MIT 