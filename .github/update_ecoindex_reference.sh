#!/usr/bin/env bash

set -euo pipefail

# Go back to repository root from .github directory
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TARGET_FILE="$ROOT_DIR/services/ecoindex_reference.json"
SOURCE_URL="https://raw.githubusercontent.com/cnumr/ecoindex_reference/refs/heads/main/ecoindex_reference.json"

echo "Updating ecoindex reference from:"
echo "  $SOURCE_URL"
echo "to:"
echo "  $TARGET_FILE"

curl -fsSL "$SOURCE_URL" -o "$TARGET_FILE"

echo "ecoindex_reference.json successfully updated."

