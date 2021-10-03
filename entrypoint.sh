#!/usr/bin/env bash

set -e

IFS=" " read -r -a FILES <<< "$(echo "${INPUT_FILES[@]}" | sort -u | tr "\n" " ")"

make run PATHS="${FILES[*]}" ACTION="${INPUT_ACTION}"
