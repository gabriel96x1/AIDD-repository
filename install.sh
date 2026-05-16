#!/usr/bin/env bash
set -euo pipefail

REPO="https://github.com/gabriel96x1/AIDD-repository.git"
TMP_DIR=$(mktemp -d)

echo "▶ Cloning aidd-mcp..."
git clone --depth 1 "$REPO" "$TMP_DIR"

echo "▶ Building binary..."
cd "$TMP_DIR/aidd/mcp"
go build -o "$OLDPWD/aidd-mcp" .

echo "▶ Cleaning up..."
rm -rf "$TMP_DIR"

echo "✓ Done. Binary ready at: $(pwd)/aidd-mcp"
echo ""
echo "Add to your agent's mcp.json:"
echo "  \"aidd-mcp\": { \"command\": \"$(pwd)/aidd-mcp\" }"