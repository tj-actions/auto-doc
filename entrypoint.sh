#!/usr/bin/env bash
set -exuo pipefail

EXTRA_ARGS=""
BIN_PATH="$INPUT_BIN_PATH"
REUSABLE="$INPUT_REUSABLE"
VERSION="$INPUT_VERSION"

# action.yml file
INPUT_COLUMNS=()
if [[ -n "$INPUT_INPUT_COLUMNS" ]]; then
  IFS=$'\n' read -rd '' -a INPUT_COLUMNS <<<"$INPUT_INPUT_COLUMNS"
fi

OUTPUT_COLUMNS=()
if [[ -n "$INPUT_OUTPUT_COLUMNS" ]]; then
  IFS=$'\n' read -rd '' -a OUTPUT_COLUMNS <<<"$INPUT_OUTPUT_COLUMNS"
fi

# reusable workflow
REUSABLE_SECRET_COLUMNS=()
if [[ -n "$INPUT_REUSABLE_SECRET_COLUMNS" ]]; then
  IFS=$'\n' read -rd '' -a REUSABLE_SECRET_COLUMNS <<<"$INPUT_REUSABLE_SECRET_COLUMNS"
fi

REUSABLE_INPUT_COLUMNS=()
if [[ -n "$INPUT_REUSABLE_INPUT_COLUMNS" ]]; then
  IFS=$'\n' read -rd '' -a REUSABLE_INPUT_COLUMNS <<<"$INPUT_REUSABLE_INPUT_COLUMNS"
fi

REUSABLE_OUTPUT_COLUMNS=()
if [[ -n "$INPUT_REUSABLE_OUTPUT_COLUMNS" ]]; then
  IFS=$'\n' read -rd '' -a REUSABLE_OUTPUT_COLUMNS <<<"$INPUT_REUSABLE_OUTPUT_COLUMNS"
fi

if [[ ! -f "$INPUT_FILENAME" ]]; then
  echo "::warning::No file found at: $INPUT_FILENAME"
  exit 0
fi

# action.yml file
for input_column in "${INPUT_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --inputColumns ${input_column}"
done

for output_column in "${OUTPUT_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --outputColumns ${output_column}"
done

# reusable workflow
if [[ "$REUSABLE" == "true" ]]; then
  EXTRA_ARGS="${EXTRA_ARGS} --reusable"
fi

for reusable_input_column in "${REUSABLE_INPUT_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --reusableInputColumns ${reusable_input_column}"
done

for reusable_output_column in "${REUSABLE_OUTPUT_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --reusableOutputColumns ${reusable_output_column}"
done

for reusable_secret_column in "${REUSABLE_SECRET_COLUMNS[@]}"; do
  EXTRA_ARGS="${EXTRA_ARGS} --reusableSecretColumns ${reusable_secret_column}"
done

if [[ -z "$BIN_PATH" ]]; then
  LATEST_VERSION=${VERSION:-v2.3.2}
  echo "Downloading auto-doc $LATEST_VERSION binary..."

  WINDOWS_TARGET=Windows_x86_64
  LINUX_TARGET=Linux_x86_64
  MACOS_TARGET=Darwin_x86_64
  ARCHIVE=zip
  TEMP_DIR=$(mktemp -d)

  if [[ $(uname -s) == "Linux" ]]; then
    TARGET=$LINUX_TARGET
    ARCHIVE=tar.gz
  elif [[ $(uname -s) == "Darwin" ]]; then
    TARGET=$MACOS_TARGET
    ARCHIVE=tar.gz
  else
    TARGET=$WINDOWS_TARGET
  fi

  DELAY=10
  OUTPUT_FILE="$TEMP_DIR"/auto-doc."$ARCHIVE"

  for i in $(seq 1 5); do
    curl --connect-timeout 300 -sLf https://github.com/tj-actions/auto-doc/releases/download/"$LATEST_VERSION"/auto-doc_"${LATEST_VERSION/v/}"_"$TARGET"."$ARCHIVE" -o "$OUTPUT_FILE" && break
    sleep $DELAY
    echo "$i retries"
  done

  if [[ "$ARCHIVE" == "zip" ]]; then
    unzip -q "$OUTPUT_FILE" -d "$TEMP_DIR"
  else
    tar -xzf "$OUTPUT_FILE" -C "$TEMP_DIR"
  fi

  chmod +x "$TEMP_DIR"/auto-doc
  BIN_PATH="$TEMP_DIR"/auto-doc
fi

echo "::debug::Generating documentation using ${BIN_PATH}..."
echo "::debug::Extra args: ${EXTRA_ARGS}"

# shellcheck disable=SC2086
$BIN_PATH --filename="$INPUT_FILENAME" --output="$INPUT_OUTPUT" \
  --colMaxWidth="$INPUT_COL_MAX_WIDTH" --colMaxWords="$INPUT_COL_MAX_WORDS" \
  ${EXTRA_ARGS} && exit_status=$? || exit_status=$?

rm -f "$BIN_PATH"

if [[ $exit_status -ne 0 ]]; then
  echo "::warning::Error occurred running auto-doc"
  exit "$exit_status";
fi
