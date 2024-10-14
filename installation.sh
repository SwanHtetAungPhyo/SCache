#!/bin/bash

# Define version and URL
VERSION="v1.0.0"
URL="https://github.com/SwanHtetAungPhyo/Scache/releases/tag/v1.0.0/scachev1.0.0.tar.gz"
BIN_DIR="/usr/local/bin"

# Create a temporary directory
TEMP_DIR=$(mktemp -d)

# Download the tar.gz file
echo "Downloading Scache $VERSION..."
if ! curl -L "$URL" -o "$TEMP_DIR/scache.tar.gz"; then
  echo "Error: Failed to download $URL"
  exit 1
fi

# Extract the tar.gz file
echo "Extracting files..."
if ! tar -xzf "$TEMP_DIR/scache.tar.gz" -C "$TEMP_DIR"; then
  echo "Error: Failed to extract the tar.gz file."
  exit 1
fi

# Move the binary to the desired directory
echo "Installing Scache..."
if [ -f "$TEMP_DIR/scache/scache" ]; then
  mv "$TEMP_DIR/scache/scache" "$BIN_DIR/scache"
else
  echo "Error: scache binary not found in the extracted files."
  exit 1
fi

# Clean up
rm -rf "$TEMP_DIR"

echo "Scache installed successfully! You can run it using 'scache --version' to verify."
