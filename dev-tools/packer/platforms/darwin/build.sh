#!/bin/sh

set -e

BASEDIR=$(dirname "$0")

# executed from the top directory
runid=darwin-$BEAT-$ARCH

cat beats/$BEAT.yml archs/$ARCH.yml version.yml > build/settings-$runid.yml
gotpl ${BASEDIR}/run.sh.j2 < build/settings-$runid.yml > build/run-$runid.sh
chmod +x build/run-$runid.sh

docker run -v `pwd`/build:/build -e BUILDID=$BUILDID -e RUNID=$runid --name build-image tudorg/fpm /build/run-$runid.sh
docker cp build-image:/build/upload `pwd`/build/binary
docker rm -v build-image

rm build/settings-$runid.yml build/run-$runid.sh
