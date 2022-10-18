#!/usr/bin/env bash

set -eu

EXTRA_ARGS=""

for input_column in "${INPUT_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --inputColumns ${input_column}"
done

for output_column in "${OUTPUT_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --outputColumns ${output_column}"
done

if [[ -z "$INPUT_BIN_PATH" ]]; then
  TEMP_DIR=$(mktemp -d)
  curl -sf https://gobinaries.com/github.com/tj-actions/auto-doc | PREFIX=$TEMP_DIR sh

  INPUT_BIN_PATH="$TEMP_DIR/auto-doc"

  # Remove the temp directory on exit.
  trap 'rm -rf "$TEMP_DIR"' EXIT
fi

chmod +x "$INPUT_BIN_PATH"

$INPUT_BIN_PATH --version

echo "::debug::Generating documentation using ${INPUT_BIN_PATH}..."
echo "::debug::Extra args: ${EXTRA_ARGS}"

$INPUT_BIN_PATH --action="$INPUT_ACTION" --output="$INPUT_OUTPUT" \
  --colMaxWidth="$INPUT_COL_MAX_WIDTH" --colMaxWords="$INPUT_COL_MAX_WORDS" \
  "${EXTRA_ARGS}" && exit_status=$? || exit_status=$?

# Remove the bin path if it still exists.
[[ -f "$INPUT_BIN_PATH" ]] && rm -f "$INPUT_BIN_PATH"

if [[ $exit_status -ne 0 ]]; then
  echo "::warning::Error occurred running auto-doc"
  exit $exit_status;
fi