#!/usr/bin/env bash

set -e

make run PATHS="${INPUT_FILES}" ACTION="${INPUT_ACTION}"
