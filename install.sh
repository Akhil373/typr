#!/usr/bin/env bash
set -euo pipefail

# typr installer
URL="https://github.com/Akhil373/typr/releases/download/v1.0.0/typr-linux"
BIN_NAME="typr"
INSTALL_DIR="${HOME}/.local/bin"
BIN_PATH="${INSTALL_DIR}/${BIN_NAME}"

echo "Installing ${BIN_NAME}..."

# 1. Create install dir
mkdir -p "${INSTALL_DIR}"

# 2. Detect downloader
if command -v curl >/dev/null 2>&1; then
    DOWNLOADER="curl -fsSL -o"
elif command -v wget >/dev/null 2>&1; then
    DOWNLOADER="wget -qO"
else
    echo "Error: curl or wget is required to install ${BIN_NAME}" >&2
    exit 1
fi

# 3. Download
echo "Downloading from ${URL}"
TMP_FILE="$(mktemp)"
trap 'rm -f "$TMP_FILE"' EXIT

if [[ $DOWNLOADER == curl* ]]; then
    curl -fsSL -o "$TMP_FILE" "$URL"
else
    wget -qO "$TMP_FILE" "$URL"
fi

# 4. Install
mv "$TMP_FILE" "$BIN_PATH"
chmod +x "$BIN_PATH"
trap - EXIT

echo "Installed to ${BIN_PATH}"

# 5. Check PATH
case ":$PATH:" in
*":${INSTALL_DIR}:"*)
    echo "You're all set! Run: ${BIN_NAME}"
    ;;
*)
    echo "Add it to your PATH by running:"
    echo ""
    echo "  echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ${RC_FILE}"
    echo ""
    echo "Or to do it automatically now, run:"
    echo ""
    echo "  grep -q '\.local/bin' ${RC_FILE} 2>/dev/null || echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ${RC_FILE} && source ${RC_FILE}"
    echo ""
    ;;
esac
