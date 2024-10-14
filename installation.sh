#!/bin/bash


VERSION="v1.0.0"
URL="https://github.com/SwanHtetAungPhyo/Scache/releases/tag/v1.0.0/scachev1.0.0.tar.gz"
BIN_DIR="/usr/local/bin"


TEMP_DIR=$(mktemp -d)


echo "Downloading Scache $VERSION..."
if ! curl -L "$URL" -o "$TEMP_DIR/scache.tar.gz"; then
  echo "Error: Failed to download $URL"
  exit 1
fi


echo "Extracting files..."
if ! tar -xzf "$TEMP_DIR/scache.tar.gz" -C "$TEMP_DIR"; then
  echo "Error: Failed to extract the tar.gz file."
  exit 1
fi


echo "Installing Scache..."
if [ -f "$TEMP_DIR/scache/scache" ]; then
  mv "$TEMP_DIR/scache/scache" "$BIN_DIR/scache"
else
  echo "Error: scache binary not found in the extracted files."
  exit 1
fi


rm -rf "$TEMP_DIR"

echo "Scache installed successfully! You can run it using 'scache --version' to verify."
echo "Scache installed successfully! You can run it using 'scache --version' to verify."
echo "To get started with Scache, follow these instructions:"
echo
echo "1. Start the Caching Server:"
echo "   Run the following command to start the server:"
echo "   ./scache"
echo
echo "2. Open Another Terminal Window:"
echo "   Use 'nc' (netcat) to send requests to your caching server:"
echo "   Example command to set a value:"
echo "   echo '{\"command\": \"set\", \"key\": \"exampleKey\", \"value\": \"exampleValue\", \"expiration\": 60}' | nc localhost 8080"
echo
echo "   Example command to get a value:"
echo "   echo '{\"command\": \"get\", \"key\": \"exampleKey\"}' | nc localhost 8080"
echo
echo "   Example command to delete a value:"
echo "   echo '{\"command\": \"delete\", \"key\": \"exampleKey\"}' | nc localhost 8080"
echo
echo "3. Check Responses:"
echo "   You will see responses in the terminal where you ran the 'nc' command."
echo
echo "4. Closing the Connection:"
echo "   You can close the terminal running 'nc' when you are done."