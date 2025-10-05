#!/bin/bash

# --- Configuration ---
APP_NAME="EvernightMoments"
BIN_SOURCE="EvernightMoments"
ICON_SOURCE="EvernightMoments.png"
# Including the config file mentioned previously
CONFIG_SOURCE="EvernightMoments.json"

echo "===== $APP_NAME Install ====="

# User-specific Local Paths
BIN_DEST_DIR="$HOME/.local/bin"
ICON_DEST_DIR="$HOME/.local/share/icons"
MENU_DEST_DIR="$HOME/.local/share/applications"

# Target File Paths
BIN_DEST="$BIN_DEST_DIR/$APP_NAME"
CONFIG_DEST="$BIN_DEST_DIR/$CONFIG_SOURCE"
ICON_DEST="$ICON_DEST_DIR/$ICON_SOURCE"
MENU_DEST="$MENU_DEST_DIR/$APP_NAME.desktop"

# --- Directory Preparation ---
# Ensure local directories exist
mkdir -p "$BIN_DEST_DIR"
mkdir -p "$ICON_DEST_DIR"
mkdir -p "$MENU_DEST_DIR"

# --- File Existence Check ---
if [[ ! -f "$BIN_SOURCE" ]]; then
    echo "Error: Source executable '$BIN_SOURCE' not found in current directory."
    exit 1
fi

# --- Identify Desktop Path ---
USER_DESKTOP=$(xdg-user-dir DESKTOP 2>/dev/null)
if [ -z "$USER_DESKTOP" ] || [ ! -d "$USER_DESKTOP" ]; then
    USER_DESKTOP="$HOME/Desktop"
fi

echo "Target User: $USER"
echo "Target Desktop: $USER_DESKTOP"

# 1. Install Binary
echo "Copying binary: ./$BIN_SOURCE -> $BIN_DEST"
cp "$BIN_SOURCE" "$BIN_DEST"
chmod 755 "$BIN_DEST"
chmod +x "$BIN_DEST"

# 1.1 Install Config (if exists)
if [[ -f "$CONFIG_SOURCE" ]]; then
    echo "Copying config: ./$CONFIG_SOURCE -> $CONFIG_DEST"
    cp "$CONFIG_SOURCE" "$CONFIG_DEST"
    chmod 644 "$CONFIG_DEST"
fi

# 2. Install Icon
if [[ -f "$ICON_SOURCE" ]]; then
    echo "Copying icon: ./$ICON_SOURCE -> $ICON_DEST"
    cp "$ICON_SOURCE" "$ICON_DEST"
    chmod 644 "$ICON_DEST"
fi

# 3. Create .desktop File
TMP_DESKTOP="/tmp/$APP_NAME.desktop"
echo "Creating desktop entry file at $TMP_DESKTOP"

# Note: Using absolute path for Icon in local installation to ensure reliability
cat <<EOF > "$TMP_DESKTOP"
[Desktop Entry]
Version=1.0
Type=Application
Name=$APP_NAME Config
Comment=$APP_NAME is a utility that automatically renames your visual archives by extracting original capture timestamps.
Comment[zh_CN]=$APP_NAME 是一款通过提取照片原始拍摄时间，为您自动重命名的工具。
Comment[zh_TW]=$APP_NAME 是一款透過提取照片原始拍攝時間，為您自動重新命名的工具。
Comment[ja]=$APP_NAME は、写真の撮影日時を抽出し、ファイル名を自動でリネームするツールです。
Exec=$BIN_DEST
Icon=$ICON_DEST
Terminal=true
Categories=Graphics;
EOF

# 4. Install to User Menu
echo "Installing menu entry: $TMP_DESKTOP -> $MENU_DEST"
mv "$TMP_DESKTOP" "$MENU_DEST"
chmod 644 "$MENU_DEST"
chmod +x "$MENU_DEST"

# 5. Install to User's Desktop
USER_SHORTCUT="$USER_DESKTOP/$APP_NAME.desktop"
echo "Creating desktop shortcut: $MENU_DEST -> $USER_SHORTCUT"
cp "$MENU_DEST" "$USER_SHORTCUT"
chmod 644 "$USER_SHORTCUT"
chmod +x "$USER_SHORTCUT"

# 6. Refresh Desktop Database
if command -v update-desktop-database >/dev/null 2>&1; then
    echo "Updating user desktop database..."
    update-desktop-database "$MENU_DEST_DIR"
fi

echo "Binary: $BIN_DEST"
echo "Config: $CONFIG_DEST"
echo "Menu Entry: $MENU_DEST"
echo "Desktop Shortcut: $USER_SHORTCUT"
echo "Note: Please ensure $BIN_DEST_DIR is in your PATH."
echo "Installation complete for $APP_NAME."
