#!/usr/bin/env bash

set -euo pipefail

source .buildkite/env-scripts/util.sh

echo "--- Installing kind & kubectl"
retry -t 5 -- .buildkite/deploy/kubernetes/scripts/install-kind.sh
retry -t 5 -- .buildkite/deploy/kubernetes/scripts/install-kubectl.sh

echo "Setting up kind"
max_retries=3
timeout=5
retries=0

while true; do
  echo "Creating cluster"
  script_output=$(.buildkite/deploy/kubernetes/scripts/kind-setup.sh 2>&1)
  exit_code=$?

  echo "Script Output: $script_output"

  if [ $exit_code -eq 0 ]; then
    break
  else
    retries=$((retries + 1))

    if [ $retries -gt $max_retries ]; then
      kind delete cluster
      echo "Kind setup FAILED: $script_output"
      exit 1
    fi

    kind delete cluster

    sleep_time=$((timeout * retries))
    echo "Retry #$retries failed. Retrying after ${sleep_time}s..."
    sleep $sleep_time
  fi
done
