#!/bin/bash

# --- Configuration ---
APP_NAME="EvernightMoments"
ICON_SOURCE="EvernightMoments.png"
CONFIG_SOURCE="EvernightMoments.json"

echo "===== $APP_NAME Uninstall ====="

# Paths
BIN_DEST="$HOME/.local/bin/$APP_NAME"
CONFIG_DEST="$HOME/.local/bin/$CONFIG_SOURCE"
ICON_DEST="$HOME/.local/share/icons/$ICON_SOURCE"
MENU_DEST="$HOME/.local/share/applications/$APP_NAME.desktop"

USER_DESKTOP=$(xdg-user-dir DESKTOP 2>/dev/null)
if [ -z "$USER_DESKTOP" ] || [ ! -d "$USER_DESKTOP" ]; then
    USER_DESKTOP="$HOME/Desktop"
fi
USER_SHORTCUT="$USER_DESKTOP/$APP_NAME.desktop"

# 1. Remove Files
for FILE in "$BIN_DEST" "$CONFIG_DEST" "$ICON_DEST" "$MENU_DEST" "$USER_SHORTCUT"; do
    if [ -f "$FILE" ]; then
        echo "Removing: $FILE"
        rm -f "$FILE"
    else
        echo "Not found: $FILE"
    fi
done

# 2. Update Database
if command -v update-desktop-database >/dev/null 2>&1; then
    echo "Updating user desktop database..."
    update-desktop-database "$HOME/.local/share/applications"
fi

echo "Uninstallation complete."
