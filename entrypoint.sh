#!/usr/bin/env bash

set -e

make build

for path in ${INPUT_FILES}
do
  ./auto_doc "$path"
done
