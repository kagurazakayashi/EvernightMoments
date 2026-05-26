#!/bin/bash

# ==============================================================================
# CONFIGURATION
# ==============================================================================
APP_NAME="EvernightMoments"
ICON_SOURCE="EvernightMoments.png"

# User-specific Local Paths
BIN_DEST_DIR="$HOME/.local/bin"
ICON_DEST_DIR="$HOME/.local/share/icons"
MENU_DEST_DIR="$HOME/.local/share/applications"
CONFIG_DEST_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/$APP_NAME"

# Target File Paths
BIN_DEST="$BIN_DEST_DIR/$APP_NAME"
ICON_DEST="$ICON_DEST_DIR/$ICON_SOURCE"
MENU_DEST="$MENU_DEST_DIR/$APP_NAME.desktop"
CONFIG_LEGACY="$BIN_DEST_DIR/$APP_NAME.json"
CONFIG_DEST="$CONFIG_DEST_DIR/config.json"

echo "------------------------------------------------------------"
echo "Starting Uninstallation for $APP_NAME on Linux"
echo "------------------------------------------------------------"

# Identify Desktop Path
USER_DESKTOP=$(xdg-user-dir DESKTOP 2>/dev/null)
[[ -z "$USER_DESKTOP" ]] && USER_DESKTOP="$HOME/Desktop"
USER_SHORTCUT="$USER_DESKTOP/$APP_NAME.desktop"

# ==============================================================================
# 1. REMOVE DESKTOP SHORTCUT
# ==============================================================================
echo "[Step 1/4] Removing desktop shortcut..."
if [ -f "$USER_SHORTCUT" ]; then
    echo "Deleting shortcut:"
    echo "  Target: $USER_SHORTCUT"
    rm -f "$USER_SHORTCUT"
    echo "  Status: Removed"
else
    echo "  Status: Not found"
fi

# ==============================================================================
# 2. REMOVE MENU ENTRY & ICON
# ==============================================================================
echo "[Step 2/4] Removing menu entry and icon..."

if [ -f "$MENU_DEST" ]; then
    echo "Deleting .desktop file:"
    echo "  Target: $MENU_DEST"
    rm -f "$MENU_DEST"
    echo "  Status: Removed"
    
    # Update desktop database so it disappears from app launcher immediately
    echo "Refreshing desktop database..."
    update-desktop-database "$MENU_DEST_DIR" 2>/dev/null
else
    echo "  Status: Menu entry not found"
fi

if [ -f "$ICON_DEST" ]; then
    echo "Deleting icon file:"
    echo "  Target: $ICON_DEST"
    rm -f "$ICON_DEST"
    echo "  Status: Removed"
else
    echo "  Status: Icon not found"
fi

# ==============================================================================
# 3. REMOVE BINARY AND CONFIG
# ==============================================================================
echo "[Step 3/4] Removing binary and configuration..."

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
# 4. FINAL SUMMARY
# ==============================================================================
echo "------------------------------------------------------------"
echo "UNINSTALLATION COMPLETE"
echo "------------------------------------------------------------"
echo "All files related to '$APP_NAME' have been removed."
echo "------------------------------------------------------------"
