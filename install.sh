#!/bin/bash

# ==============================================================================
# CONFIGURATION
# ==============================================================================
APP_NAME="EvernightMoments"
BIN_SOURCE="EvernightMoments"
ICON_SOURCE="EvernightMoments.png"

# User-specific Local Paths (XDG Standard)
BIN_DEST_DIR="$HOME/.local/bin"
ICON_DEST_DIR="$HOME/.local/share/icons"
MENU_DEST_DIR="$HOME/.local/share/applications"
CONFIG_DEST_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/$APP_NAME"

# Target File Paths
BIN_DEST="$BIN_DEST_DIR/$APP_NAME"
ICON_DEST="$ICON_DEST_DIR/$ICON_SOURCE"
MENU_DEST="$MENU_DEST_DIR/$APP_NAME.desktop"
CONFIG_DEST="$CONFIG_DEST_DIR/config.json"

echo "------------------------------------------------------------"
echo "Starting Installation for $APP_NAME on Linux"
echo "------------------------------------------------------------"

# ==============================================================================
# 1. DIRECTORY PREPARATION
# ==============================================================================
echo "[Step 1/6] Preparing directories..."
for dir in "$BIN_DEST_DIR" "$ICON_DEST_DIR" "$MENU_DEST_DIR" "$CONFIG_DEST_DIR"; do
    if [ ! -d "$dir" ]; then
        echo "Creating directory: $dir"
        mkdir -p "$dir"
    else
        echo "Directory exists: $dir"
    fi
done

# Identify Desktop Path
USER_DESKTOP=$(xdg-user-dir DESKTOP 2>/dev/null)
[[ -z "$USER_DESKTOP" ]] && USER_DESKTOP="$HOME/Desktop"

# Check if source binary exists
if [[ ! -f "$BIN_SOURCE" ]]; then
    echo "Error: Source binary '$BIN_SOURCE' not found in current directory."
    exit 1
fi

# ==============================================================================
# 2. INSTALL BINARY
# ==============================================================================
echo "[Step 2/6] Installing binary to local bin..."

echo "Copying binary:"
echo "  From: ./$BIN_SOURCE"
echo "  To:   $BIN_DEST"
cp "$BIN_SOURCE" "$BIN_DEST"
chmod 755 "$BIN_DEST"

# ==============================================================================
# 3. INSTALL ICON
# ==============================================================================
echo "[Step 3/6] Installing application icon..."
if [[ -f "$ICON_SOURCE" ]]; then
    echo "Copying icon:"
    echo "  From: ./$ICON_SOURCE"
    echo "  To:   $ICON_DEST"
    cp "$ICON_SOURCE" "$ICON_DEST"
    chmod 644 "$ICON_DEST"
else
    echo "Notice: Icon source not found, shortcut might show default icon."
fi

# ==============================================================================
# 4. CREATE DESKTOP ENTRY
# ==============================================================================
echo "[Step 4/6] Generating .desktop entry..."
TMP_DESKTOP="/tmp/$APP_NAME.desktop"

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

echo "Installing menu entry:"
echo "  To: $MENU_DEST"
mv "$TMP_DESKTOP" "$MENU_DEST"
chmod 644 "$MENU_DEST"

# Update desktop database to ensure it shows up in menus
update-desktop-database "$MENU_DEST_DIR" 2>/dev/null

# ==============================================================================
# 5. CREATE DESKTOP SHORTCUT
# ==============================================================================
echo "[Step 5/6] Creating desktop shortcut..."
USER_SHORTCUT="$USER_DESKTOP/$APP_NAME.desktop"

echo "Creating shortcut:"
echo "  Source: $MENU_DEST"
echo "  To:     $USER_SHORTCUT"
cp "$MENU_DEST" "$USER_SHORTCUT"
chmod +x "$USER_SHORTCUT"

# ==============================================================================
# 6. SHELL PATH CONFIGURATION
# ==============================================================================
echo "[Step 6/6] Checking Environment PATH..."
CURRENT_SHELL=$(basename "$SHELL")
PATH_LINE="export PATH=\"\$HOME/.local/bin:\$PATH\""
CONF_FILE=""

case "$CURRENT_SHELL" in
    zsh)  CONF_FILE="$HOME/.zshrc" ;;
    bash) CONF_FILE="$HOME/.bashrc" ;;
    fish)
        CONF_FILE="$HOME/.config/fish/config.fish"
        PATH_LINE="set -gx PATH \$HOME/.local/bin \$PATH"
        mkdir -p "$(dirname "$CONF_FILE")"
        ;;
    *)    CONF_FILE="$HOME/.profile" ;;
esac

echo "Detected Shell: $CURRENT_SHELL"
echo "Target Config:  $CONF_FILE"

if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
    if ! grep -qF "$PATH_LINE" "$CONF_FILE" 2>/dev/null; then
        echo "Updating PATH in $CONF_FILE..."
        echo -e "\n$PATH_LINE" >> "$CONF_FILE"
        PATH_UPDATED=true
    else
        echo "PATH entry already exists in $CONF_FILE."
    fi
else
    echo "PATH is already configured."
fi

# ==============================================================================
# 7. FINAL SUMMARY
# ==============================================================================
echo "------------------------------------------------------------"
echo "INSTALLATION COMPLETE"
echo "------------------------------------------------------------"
echo "Binary Path:      $BIN_DEST"
echo "Config Path:      $CONFIG_DEST"
echo "Icon Path:        $ICON_DEST"
echo "Menu Entry:       $MENU_DEST"
echo "Desktop Link:     $USER_SHORTCUT"
echo "------------------------------------------------------------"

if [ "$PATH_UPDATED" = true ]; then
    echo "IMPORTANT: Please run 'source $CONF_FILE' or restart terminal."
    echo "Then you can run the app by typing: $APP_NAME"
else
    echo "Installation successful. You can now run '$APP_NAME' from anywhere."
fi
