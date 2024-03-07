#!/usr/bin/env bash

source .buildkite/scripts/install_tools.sh

set -euo pipefail

trap 'cleanup; unset_secrets' EXIT

echo "--- Run Cloud Tests for $BEATS_PROJECT_NAME"
pushd "${BEATS_PROJECT_NAME}" > /dev/null

mage build test

popd > /dev/null
