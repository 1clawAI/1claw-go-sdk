#!/usr/bin/env bash
# Regenerate internal/openapi from the OpenAPI spec.
# Run from repo root. Expects spec at ../1claw-openapi-spec/openapi.yaml by default.

set -e
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
SPEC_PATH="${SPEC_PATH:-$REPO_ROOT/../1claw-openapi-spec/openapi.yaml}"
OUTPUT_PATH="$REPO_ROOT/internal/openapi"

if [[ ! -f "$SPEC_PATH" ]]; then
  echo "Error: OpenAPI spec not found at $SPEC_PATH"
  echo "Set SPEC_PATH to override."
  exit 1
fi

cd "$REPO_ROOT"
npx @openapitools/openapi-generator-cli generate \
  -i "$SPEC_PATH" \
  -g go \
  -o "$OUTPUT_PATH" \
  --additional-properties=packageName=openapi,outputAsLibrary=true

# Remove generated go.mod so internal/openapi stays part of parent module
rm -f "$OUTPUT_PATH/go.mod" "$OUTPUT_PATH/go.sum"
echo "Generated client at $OUTPUT_PATH"
