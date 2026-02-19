#!/usr/bin/env bash

set -e

BIN="nmsurf"
INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/nmsurf"

echo "Uninstalling $BIN..."

if [ -f "$INSTALL_DIR/$BIN" ]; then
    sudo rm -f "$INSTALL_DIR/$BIN"
    echo "Removed binary: $INSTALL_DIR/$BIN"
else
    echo "Binary not found in $INSTALL_DIR"
fi

if [ -d "$CONFIG_DIR" ]; then
    echo
    read -rp "Remove config directory ($CONFIG_DIR)? [y/N]: " confirm
    case "$confirm" in
        [yY][eE][sS]|[yY])
            rm -rf "$CONFIG_DIR"
            echo "Removed config directory"
            ;;
        *)
            echo "Config directory preserved"
            ;;
    esac
fi

echo
echo "Uninstall complete."
