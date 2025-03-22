#!/usr/bin/env bash
set -euo pipefail

EXTRA_ARGS=""
BIN_PATH="$INPUT_BIN_PATH"
REUSABLE="$INPUT_REUSABLE"
VERSION="$INPUT_VERSION"

# action.yml file
IFS=$'\n' read -rd '' -a INPUT_COLUMNS <<<"$INPUT_INPUT_COLUMNS" || true

if [[ ${#INPUT_COLUMNS[@]} -gt 0 ]]; then
  for input_column in "${INPUT_COLUMNS[@]}"; do
    EXTRA_ARGS="${EXTRA_ARGS} --inputColumns=${input_column}"
  done
fi

IFS=$'\n' read -rd '' -a OUTPUT_COLUMNS <<<"$INPUT_OUTPUT_COLUMNS" || true

if [[ ${#OUTPUT_COLUMNS[@]} -gt 0 ]]; then
  for output_column in "${OUTPUT_COLUMNS[@]}"; do
    EXTRA_ARGS="${EXTRA_ARGS} --outputColumns=${output_column}"
  done
fi

# reusable workflow
IFS=$'\n' read -rd '' -a REUSABLE_SECRET_COLUMNS <<<"$INPUT_REUSABLE_SECRET_COLUMNS" || true

if [[ ${#REUSABLE_SECRET_COLUMNS[@]} -gt 0 ]]; then
  for reusable_secret_column in "${REUSABLE_SECRET_COLUMNS[@]}"; do
    EXTRA_ARGS="${EXTRA_ARGS} --reusableSecretColumns=${reusable_secret_column}"
  done
fi

IFS=$'\n' read -rd '' -a REUSABLE_INPUT_COLUMNS <<<"$INPUT_REUSABLE_INPUT_COLUMNS" || true

if [[ ${#REUSABLE_INPUT_COLUMNS[@]} -gt 0 ]]; then
  for reusable_input_column in "${REUSABLE_INPUT_COLUMNS[@]}"; do
    EXTRA_ARGS="${EXTRA_ARGS} --reusableInputColumns=${reusable_input_column}"
  done
fi

IFS=$'\n' read -rd '' -a REUSABLE_OUTPUT_COLUMNS <<<"$INPUT_REUSABLE_OUTPUT_COLUMNS" || true

if [[ ${#REUSABLE_OUTPUT_COLUMNS[@]} -gt 0 ]]; then
  for reusable_output_column in "${REUSABLE_OUTPUT_COLUMNS[@]}"; do
    EXTRA_ARGS="${EXTRA_ARGS} --reusableOutputColumns=${reusable_output_column}"
  done
fi

if [[ ! -f "$INPUT_FILENAME" ]]; then
  echo "::warning::No file found at: $INPUT_FILENAME"
  exit 0
fi

if [[ ! -f "$BIN_PATH" ]]; then
  echo "::error::No binary found at: $BIN_PATH"
  exit 1
fi

# reusable workflow
if [[ "$REUSABLE" == "true" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --reusable"
fi

# markdown links
if [[ "$INPUT_MARKDOWN_LINKS" == "true" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --markdownLinks"
fi

# repository
if [[ -n "$INPUT_REPOSITORY" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --repository=${INPUT_REPOSITORY}"
fi

# token
if [[ -n "$INPUT_TOKEN" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --token=${INPUT_TOKEN}"
fi

# use_code_blocks
if [[ "$INPUT_USE_CODE_BLOCKS" == "true" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --useCodeBlocks"
fi

# use_major_version
if [[ "$INPUT_USE_MAJOR_VERSION" == "true" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --useMajorVersion"
fi

# use_tag_commit_hash
if [[ "$INPUT_USE_TAG_COMMIT_HASH" == "true" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --useTagCommitHash"
fi

echo "::debug::Generating documentation using ${BIN_PATH}..."
echo "::debug::Extra args: ${EXTRA_ARGS}"

# shellcheck disable=SC2086
$BIN_PATH --filename="$INPUT_FILENAME" --output="$INPUT_OUTPUT" \
  --colMaxWidth="$INPUT_COL_MAX_WIDTH" --colMaxWords="$INPUT_COL_MAX_WORDS" \
  ${EXTRA_ARGS} && exit_status=$? || exit_status=$?

if [[ $exit_status -ne 0 ]]; then
  echo "::warning::Error occurred running auto-doc"
  exit "$exit_status";
fi
