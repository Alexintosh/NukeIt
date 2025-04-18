# 🧹 nuke - CLI AppCleaner Alternative for macOS

`nuke` is a command-line tool written in Go for macOS that helps users fully uninstall applications by removing the main app bundle **and** associated files (like caches, preferences, logs, etc.).

## 📌 Purpose

Most macOS applications leave behind files after you drag them to the Trash. `nuke` helps you find and remove these leftover files, keeping your system clean.

## 🛠 Features

- Detects and deletes app support files across macOS system paths
- Reads app bundle ID from `.app` files
- Provides dry-run, verbose, and force-delete options
- Confirms file deletion before removing anything (unless `--force` is used)
- Interactive Terminal UI (TUI) for a more visual experience

### ✨ TUI Features

- **Spinner** during file scanning process
- **Interactive file selection** to choose which files to delete
- **Progress bar** showing deletion progress
- **Keyboard controls**:
  - `↑/↓` - Navigate through files
  - `Space` - Toggle selection of a file
  - `a` - Select all files
  - `n` - Deselect all files
  - `Enter` - Confirm and delete selected files
  - `Ctrl+C` - Quit

## 💻 Installation

### Install via curl
`curl -fsSL https://raw.githubusercontent.com/Alexintosh/NukeIt/main/install.sh | bash`

This will download and install the nuke binary to `/usr/local/bin`, making it available system-wide.

### From source

```bash
git clone https://github.com/alexintosh/gocleaner.git
cd gocleaner
go build -o nuke cmd/nuke/main.go
```

Then move the binary to your PATH:

```bash
mv nuke /usr/local/bin/
```

## 💻 Usage

```bash
nuke uninstall <AppName> [--dry-run] [--force] [--verbose] [--no-tui]
```

### Example

```bash
nuke uninstall Spotify
```

### Flags

- `--dry-run` – Show what would be deleted, but don't delete.
- `--force` – Skip confirmation prompt and delete files immediately.
- `--verbose` – Show detailed scanning and deletion info.
- `--no-tui` – Disable the interactive TUI and use the simple CLI interface.

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