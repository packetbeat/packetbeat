#!/usr/bin/env bash

set -euo pipefail

echo "--- Run Python Intergration Tests for $BEATS_PROJECT_NAME"
pushd "${BEATS_PROJECT_NAME}" > /dev/null

mage pythonIntegTest

popd > /dev/null
