#!/usr/bin/env bash
set -euo pipefail

REPO="https://github.com/gabriel96x1/AIDD-repository.git"
TMP_DIR=$(mktemp -d)

echo "▶ Cloning aidd-mcp..."
git clone --depth 1 "$REPO" "$TMP_DIR"

echo "▶ Building binary..."
cd "$TMP_DIR/aidd/mcp"
go build -o "$OLDPWD/aidd-mcp-server" .

echo "▶ Cleaning up..."
rm -rf "$TMP_DIR"

echo "✓ Done. Binary ready at: $(pwd)/aidd-mcp-server"
echo ""
echo "Add to your agent's mcp.json:"
echo " {
	"servers": {
		"aidd-server": {
			"type": "stdio",
			"command": "./aidd-mcp",
			"args": []
		}
	},
	"inputs": []
}"