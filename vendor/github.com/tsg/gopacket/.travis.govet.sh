#!/bin/bash

cd "$(dirname $0)"
DIRS=". layers pcap pcapgo pfring tcpassembly tcpassembly/tcpreader routing bytediff macs"
set -e
for subdir in $DIRS; do
  pushd $subdir
  go vet
  popd
done
