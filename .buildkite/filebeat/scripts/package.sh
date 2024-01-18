#!/usr/bin/env bash

set -euo pipefail

source .buildkite/env-scripts/unix-env.sh

echo ":: Evaluate Filebeat Changes ::"

echo ":: Start Packaging ::"
cd filebeat
umask 0022
mage package
