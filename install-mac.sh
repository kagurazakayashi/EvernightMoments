#!/bin/bash

# ==============================================================================
# CONFIGURATION
# ==============================================================================
APP_NAME="EvernightMoments"
BIN_SOURCE="EvernightMoments"
ICON_SOURCE="EvernightMoments.icns"

# User-specific Local Paths
BIN_DEST_DIR="$HOME/.local/bin"
APP_DEST_DIR="$HOME/Applications"
SHORTCUT="/Applications"
CONFIG_DEST_DIR="$HOME/Library/Application Support/$APP_NAME"

# Target File Paths
BIN_DEST="$BIN_DEST_DIR/$APP_NAME"
APP_BUNDLE="$APP_DEST_DIR/$APP_NAME Config.app"
CONFIG_DEST="$CONFIG_DEST_DIR/config.json"

echo "------------------------------------------------------------"
echo "Starting Installation for $APP_NAME on macOS"
echo "------------------------------------------------------------"

# ==============================================================================
# 1. DIRECTORY PREPARATION
# ==============================================================================
echo "[Step 1/5] Preparing directories..."
for dir in "$BIN_DEST_DIR" "$APP_DEST_DIR" "$CONFIG_DEST_DIR"; do
    if [ ! -d "$dir" ]; then
        echo "Creating directory: $dir"
        mkdir -p "$dir"
    else
        echo "Directory exists: $dir"
    fi
done

# Check if source binary exists
if [[ ! -f "$BIN_SOURCE" ]]; then
    echo "Error: Source binary '$BIN_SOURCE' not found in current directory."
    exit 1
fi

# ==============================================================================
# 2. INSTALL BINARY
# ==============================================================================
echo "[Step 2/5] Installing binary to local bin..."

echo "Copying binary:"
echo "  From: ./$BIN_SOURCE"
echo "  To:   $BIN_DEST"
cp "$BIN_SOURCE" "$BIN_DEST"
chmod 755 "$BIN_DEST"

# ==============================================================================
# 3. CREATE APP BUNDLE (Terminal Wrapper)
# ==============================================================================
echo "[Step 3/5] Creating macOS App Bundle..."
echo "Target App Bundle: $APP_BUNDLE"

# osacompile creates a native .app that runs the binary via Terminal
osacompile -e "tell application \"Terminal\"" \
           -e "    activate" \
           -e "    do script \"\\\"$BIN_DEST\\\"\"" \
           -e "end tell" \
           -o "$APP_BUNDLE" >/dev/null 2>&1

if [ $? -eq 0 ]; then
    echo "Successfully created application bundle."
else
    echo "Error: Failed to create App Bundle."
    exit 1
fi

# ==============================================================================
# 4. INSTALL ICON
# ==============================================================================
echo "[Step 4/5] Applying custom icon..."
APP_ICON_DEST="$APP_BUNDLE/Contents/Resources/applet.icns"

if [[ -f "$ICON_SOURCE" ]]; then
    echo "Copying icon:"
    echo "  From: $ICON_SOURCE"
    echo "  To:   $APP_ICON_DEST"
    cp "$ICON_SOURCE" "$APP_ICON_DEST"
    
    # Force Finder to refresh the icon cache
    touch "$APP_BUNDLE"
    echo "Icon applied successfully."
else
    echo "Warning: Icon file '$ICON_SOURCE' not found. App will use default icon."
fi

# ==============================================================================
# 5. CREATE DESKTOP SHORTCUT
# ==============================================================================
echo "[Step 5/5] Creating desktop shortcut..."
USER_SHORTCUT="$SHORTCUT/$APP_NAME Config.app"

echo "Creating symlink:"
echo "  Source: $APP_BUNDLE"
echo "  Link:   $USER_SHORTCUT"
ln -sf "$APP_BUNDLE" "$USER_SHORTCUT"

# ==============================================================================
# 6. SHELL PATH CONFIGURATION
# ==============================================================================
echo "[Shell] Checking Environment PATH..."
CURRENT_SHELL=$(basename "$SHELL")
PATH_LINE="export PATH=\"\$HOME/.local/bin:\$PATH\""
CONF_FILE=""

# Identify correct config file based on active shell
case "$CURRENT_SHELL" in
    zsh)
        CONF_FILE="$HOME/.zshrc"
        ;;
    bash)
        CONF_FILE="$HOME/.bash_profile"
        [[ ! -f "$CONF_FILE" ]] && CONF_FILE="$HOME/.bashrc"
        ;;
    fish)
        CONF_FILE="$HOME/.config/fish/config.fish"
        PATH_LINE="set -gx PATH \$HOME/.local/bin \$PATH"
        mkdir -p "$(dirname "$CONF_FILE")"
        ;;
    *)
        CONF_FILE="$HOME/.profile"
        ;;
esac

echo "Detected Shell: $CURRENT_SHELL"
echo "Target Config:  $CONF_FILE"

# Check if PATH already contains the directory
if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
    # Double check if the line exists in the file to avoid duplicates
    if ! grep -qF "$PATH_LINE" "$CONF_FILE" 2>/dev/null; then
        echo "Updating PATH in $CONF_FILE..."
        echo -e "\n$PATH_LINE" >> "$CONF_FILE"
        PATH_UPDATED=true
    else
        echo "PATH entry already exists in $CONF_FILE."
    fi
else
    echo "PATH is already configured in current session."
fi

# ==============================================================================
# FINAL SUMMARY
# ==============================================================================
echo "------------------------------------------------------------"
echo "INSTALLATION COMPLETE"
echo "------------------------------------------------------------"
echo "Binary:           $BIN_DEST"
echo "Configuration:    $CONFIG_DEST"
echo "Application:      $APP_BUNDLE"
echo "Desktop Link:     $USER_SHORTCUT"
echo "------------------------------------------------------------"

if [ "$PATH_UPDATED" = true ]; then
    echo "IMPORTANT: Please run 'source $CONF_FILE' or restart your terminal."
    echo "Then you can run the app by typing: $APP_NAME"
else
    echo "You can now run '$APP_NAME' directly from your terminal."
fi
