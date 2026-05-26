#!/bin/bash

# ==============================================================================
# CONFIGURATION
# ==============================================================================
APP_NAME="EvernightMoments"

# User-specific Local Paths
BIN_DEST_DIR="$HOME/.local/bin"
APP_DEST_DIR="$HOME/Applications"
SHORTCUT="/Applications"
CONFIG_DEST_DIR="$HOME/Library/Application Support/$APP_NAME"

# Target File Paths
BIN_DEST="$BIN_DEST_DIR/$APP_NAME"
APP_BUNDLE="$APP_DEST_DIR/$APP_NAME Config.app"
USER_SHORTCUT="$SHORTCUT/$APP_NAME Config.app"
CONFIG_LEGACY="$BIN_DEST_DIR/$APP_NAME.json"
CONFIG_DEST="$CONFIG_DEST_DIR/config.json"

echo "------------------------------------------------------------"
echo "Starting Uninstallation for $APP_NAME on macOS"
echo "------------------------------------------------------------"

# ==============================================================================
# 1. REMOVE APP BUNDLE
# ==============================================================================
echo "[Step 1/3] Removing macOS App Bundle..."
if [ -d "$APP_BUNDLE" ]; then
    echo "Deleting directory:"
    echo "  Target: $APP_BUNDLE"
    rm -rf "$APP_BUNDLE"
    echo "  Status: Removed"
else
    echo "  Status: Not found (already removed or never installed)"
fi

# ==============================================================================
# 2. REMOVE DESKTOP SHORTCUT
# ==============================================================================
echo "[Step 2/3] Removing desktop shortcut..."
if [ -L "$USER_SHORTCUT" ] || [ -f "$USER_SHORTCUT" ]; then
    echo "Deleting symlink:"
    echo "  Target: $USER_SHORTCUT"
    rm -f "$USER_SHORTCUT"
    echo "  Status: Removed"
else
    echo "  Status: Not found"
fi

# ==============================================================================
# 3. REMOVE BINARY AND CONFIG
# ==============================================================================
echo "[Step 3/3] Removing binary and configuration..."

if [ -f "$BIN_DEST" ]; then
    echo "Deleting binary:"
    echo "  Target: $BIN_DEST"
    rm -f "$BIN_DEST"
    echo "  Status: Removed"
else
    echo "  Status: Binary not found"
fi

if [ -f "$CONFIG_DEST" ]; then
    echo "Deleting configuration (standard):"
    echo "  Target: $CONFIG_DEST"
    rm -f "$CONFIG_DEST"
    echo "  Status: Removed"
else
    echo "  Status: Standard configuration not found"
fi

if [ -f "$CONFIG_LEGACY" ]; then
    echo "Deleting configuration (legacy):"
    echo "  Target: $CONFIG_LEGACY"
    rm -f "$CONFIG_LEGACY"
    echo "  Status: Removed"
else
    echo "  Status: Legacy configuration not found"
fi

# ==============================================================================
# FINAL SUMMARY
# ==============================================================================
echo "------------------------------------------------------------"
echo "UNINSTALLATION COMPLETE"
echo "------------------------------------------------------------"
echo "All files related to '$APP_NAME' have been removed."
echo "------------------------------------------------------------"
