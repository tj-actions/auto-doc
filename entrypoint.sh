#!/usr/bin/env bash

set -e

make build

for path in ${INPUT_FILES}
do
  echo "$path"
  ./auto_doc --help
done