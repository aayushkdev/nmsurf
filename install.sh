#!/usr/bin/env bash

set -e

BIN="nmsurf"
INSTALL_DIR="/usr/local/bin"

echo "Installing $BIN..."

# Check Go exists
if ! command -v go >/dev/null 2>&1; then
    echo "Error: Go is not installed."
    exit 1
fi

# Check runtime dependencies
check_dep() {
    if ! command -v "$1" >/dev/null 2>&1; then
        echo "Warning: dependency '$1' not found"
    fi
}

check_dep nmcli

if command -v wofi >/dev/null 2>&1; then
    :
elif command -v rofi >/dev/null 2>&1; then
    :
else
    echo "Warning: neither wofi nor rofi found"
fi

# Ensure modules are correct
echo "Resolving dependencies..."
go mod tidy

# Build binary
echo "Building $BIN..."
go build -ldflags="-s -w" -o "$BIN" ./cmd/nmsurf

# Install binary
echo "Installing to $INSTALL_DIR..."
sudo install -Dm755 "$BIN" "$INSTALL_DIR/$BIN"

echo
echo "Install complete."
echo "Binary location: $INSTALL_DIR/$BIN"
echo
echo "Run with:"
echo "  nmsurf"
