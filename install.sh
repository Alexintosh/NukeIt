#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "Installing nuke - macOS app uninstaller..."

# Check if running on macOS
if [[ "$OSTYPE" != "darwin"* ]]; then
  echo -e "${RED}Error: This tool is designed for macOS only.${NC}"
  exit 1
fi

# Download binary
echo "Downloading nuke binary..."
TEMP_DIR=$(mktemp -d)
curl -fsSL https://github.com/Alexintosh/NukeIt/raw/refs/heads/main/nuke -o "${TEMP_DIR}/nuke"

if [ $? -ne 0 ]; then
  echo -e "${RED}Error: Failed to download the binary.${NC}"
  rm -rf "${TEMP_DIR}"
  exit 1
fi

# Make executable
chmod +x "${TEMP_DIR}/nuke"

# Move to /usr/local/bin (creating directory if it doesn't exist)
INSTALL_DIR="/usr/local/bin"
if [ ! -d "$INSTALL_DIR" ]; then
  sudo mkdir -p "$INSTALL_DIR"
fi

# Use sudo to move the file to system directory
echo "Installing nuke to ${INSTALL_DIR}..."
sudo mv "${TEMP_DIR}/nuke" "$INSTALL_DIR"

# Clean up
rm -rf "${TEMP_DIR}"

# Verify installation
if [ -x "$INSTALL_DIR/nuke" ]; then
  echo -e "${GREEN}nuke has been successfully installed!${NC}"
  echo -e "${GREEN}You can now use it by running 'nuke' in your terminal.${NC}"
else
  echo -e "${RED}Installation failed.${NC}"
  exit 1
fi

echo ""
echo "nuke is a command-line tool for macOS that helps users fully uninstall applications"
echo "by removing the main app bundle and associated files like caches, preferences, logs, etc."
echo ""
echo "Usage: nuke [application name]" 