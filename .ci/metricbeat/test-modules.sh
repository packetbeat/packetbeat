#!/bin/sh

readonly MODULES_PATH=$1
readonly FILE_NAME=$2

source ./test-module.sh

for file in $(find ${MODULES_PATH} -name "${FILE_NAME}"); do
  parent=$(dirname "$file")
  moduleDir=$(dirname "$parent")
  module=$(basename ${moduleDir})

  test_module $module
done
