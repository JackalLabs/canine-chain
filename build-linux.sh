#!/bin/bash
# Script to build canined for Linux using Docker with full CGO/ledger support

set -e

echo "Building canined for Linux with ledger support using Docker..."

# Get the directory of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# Parent directory (Github folder)
PARENT_DIR="$(dirname "$SCRIPT_DIR")"

# Architecture (default: amd64)
ARCH="${1:-amd64}"

echo "Build context: $PARENT_DIR"
echo "Target architecture: $ARCH"

# Build the Docker image and extract the binary
docker build \
  --build-arg GOARCH="$ARCH" \
  --build-arg GOOS=linux \
  --build-arg LEDGER_ENABLED=true \
  -f "$SCRIPT_DIR/Dockerfile.linux-build" \
  -t canined-linux-builder \
  "$PARENT_DIR"

# Create a temporary container to copy files out
CONTAINER_ID=$(docker create canined-linux-builder)

# Copy the binary from the container
docker cp "$CONTAINER_ID:/workspace/canine-chain/build/canined" "$SCRIPT_DIR/build/canined-linux-$ARCH"

# Copy libwasmvm.so if it exists
if docker cp "$CONTAINER_ID:/workspace/canine-chain/build/libwasmvm.so" "$SCRIPT_DIR/build/libwasmvm-linux-$ARCH.so" 2>/dev/null; then
    echo "✓ Extracted libwasmvm.so"
    WASMVM_PRESENT=true
else
    echo "⚠ Warning: libwasmvm.so not found in build"
    WASMVM_PRESENT=false
fi

# Clean up the container
docker rm "$CONTAINER_ID"

echo ""
echo "✓ Build complete:"
echo "  Binary: $SCRIPT_DIR/build/canined-linux-$ARCH"
if [ "$WASMVM_PRESENT" = true ]; then
    echo "  Library: $SCRIPT_DIR/build/libwasmvm-linux-$ARCH.so"
fi
echo ""
echo "To verify the binary:"
echo "  file build/canined-linux-$ARCH"
echo ""
if [ "$WASMVM_PRESENT" = true ]; then
    echo "Deployment instructions:"
    echo "  1. Copy both files to your Linux server:"
    echo "     scp build/canined-linux-$ARCH user@server:/usr/local/bin/canined"
    echo "     scp build/libwasmvm-linux-$ARCH.so user@server:/usr/lib/libwasmvm.so"
    echo ""
    echo "  2. Or place libwasmvm.so in the same directory as canined and set:"
    echo "     export LD_LIBRARY_PATH=/path/to/canined:\$LD_LIBRARY_PATH"
else
    echo "Note: The binary should be ready for deployment on Linux $ARCH systems."
fi
