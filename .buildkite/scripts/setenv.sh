#!/bin/bash

set -euo pipefail

SETUP_GVM_VERSION="v0.5.1"
DOCKER_COMPOSE_VERSION="1.21.0"
DOCKER_COMPOSE_VERSION_AARCH64="v2.21.0"
SETUP_WIN_PYTHON_VERSION="3.11.0"
GO_VERSION=$(cat .go-version)
ALLOW_EXTENDED_TESTS=${ALLOW_EXTENDED_TESTS:-false}
ALLOW_MANDATORY_TESTS=${ALLOW_MANDATORY_TESTS:-false}
ALLOW_MACOS_TESTS=${ALLOW_MACOS_TESTS:-false}
ALLOW_EXTENDED_WIN_TESTS=${ALLOW_EXTENDED_WIN_TESTS:-false}
ALLOW_PACKAGING=${ALLOW_PACKAGING:-false}
GITHUB_PR_TRIGGER_COMMENT=${GITHUB_PR_TRIGGER_COMMENT:-""}
ONLY_DOCS=${ONLY_DOCS:-"true"}
UI_MACOS_TESTS="$(buildkite-agent meta-data get UI_MACOS_TESTS --default ${UI_MACOS_TESTS:-"false"})"
runAllStages="$(buildkite-agent meta-data get runAllStages --default ${runAllStages:-"false"})"
metricbeat_changeset=(
  "^metricbeat/.*"
  "^go.mod"
  "^pytest.ini"
  "^dev-tools/.*"
  "^libbeat/.*"
  "^testing/.*"
  )
oss_changeset=(
  "^go.mod"
  "^pytest.ini"
  "^dev-tools/.*"
  "^libbeat/.*"
  "^testing/.*"
)
ci_changeset=(
  "^.buildkite/.*"
)
go_mod_changeset=(
  "^go.mod"
  )
docs_changeset=(
  ".*\\.(asciidoc|md)"
  "deploy/kubernetes/.*-kubernetes\\.yaml"
  )
packaging_changeset=(
  "^dev-tools/packaging/.*"
  ".go-version"
  )

if ! are_changed_only_paths "${docs_changeset[@]}" ; then
  ONLY_DOCS="false"
  echo "Changes include files outside the docs_changeset vairiabe. ONLY_DOCS=$ONLY_DOCS."
else
  echo "All changes are related to DOCS. ONLY_DOCS=$ONLY_DOCS."
fi

if are_paths_changed "${go_mod_changeset[@]}" ; then
  GO_MOD_CHANGES="true"
fi

if are_paths_changed "${packaging_changeset[@]}" ; then
  PACKAGING_CHANGES="true"
fi
