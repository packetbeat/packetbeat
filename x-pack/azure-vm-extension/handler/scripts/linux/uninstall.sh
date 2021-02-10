#!/usr/bin/env bash
set -euo pipefail
script_path=$(dirname $(realpath -s $0))
source $script_path/helper.sh

log "Uninstalling Elastic Agent" "INFO"
elastic-agent uninstall
log "Elastic Agent is uninstalled" "INFO"

