#!/bin/bash
#
# Contains the main cross compiler, that individually sets up each target build
# platform, compiles all the C dependencies, then build the requested executable
# itself.
#
# Usage: build.sh <import path>
#
# Needed environment variables:
#   REPO_REMOTE - Optional VCS remote if not the primary repository is needed
#   REPO_BRANCH - Optional VCS branch to use, if not the master branch
#   DEPS        - Optional list of C dependency packages to build
#   PACK        - Optional sub-package, if not the import path is being built
#   OUT         - Optional output prefix to override the package name
#   FLAG_V      - Optional verbosity flag to set on the Go builder
#   FLAG_RACE   - Optional race flag to set on the Go builder
#   TARGETS        - Comma separated list of build targets to compile for



# Download the canonical import path (may fail, don't allow failures beyond)

SRC_FOLDER=$SOURCE
DST_FOLDER=$GOPATH/src/$1

if [ $1 = "github.com/elastic/beats" ]; then
        WORKING_DIRECTORY=$GOPATH/src/$1
else
        WORKING_DIRECTORY=$GOPATH/src/`dirname $1`
fi

echo "Working directory=$WORKING_DIRECTORY"

if [ "$SOURCE" != "" ]; then
        mkdir -p ${DST_FOLDER}
        echo "Copying main source folder ${SRC_FOLDER} to folder ${DST_FOLDER}"
        rsync --exclude ".git"  --exclude "build/" -a ${SRC_FOLDER}/ ${DST_FOLDER}
else
        mkdir -p $GOPATH/src/`dirname $1`
        cd $GOPATH/src/`dirname $1`
        echo "Fetching main git repository $1 in folder $GOPATH/src/`dirname $1`"
        git clone https://$1.git
fi

set -e

cd $WORKING_DIRECTORY
export GOPATH=$GOPATH:`pwd`/Godeps/_workspace
export GO15VENDOREXPERIMENT=1

# Switch over the code-base to another checkout if requested
if [ "$REPO_REMOTE" != "" ]; then
  echo "Switching over to remote $REPO_REMOTE..."
  if [ -d ".git" ]; then
    git remote set-url origin $REPO_REMOTE
    git pull
  elif [ -d ".hg" ]; then
    echo -e "[paths]\ndefault = $REPO_REMOTE\n" >> .hg/hgrc
    hg pull
  fi
fi

if [ "$REPO_BRANCH" != "" ]; then
  echo "Switching over to branch $REPO_BRANCH..."
  if [ -d ".git" ]; then
    git checkout $REPO_BRANCH
  elif [ -d ".hg" ]; then
    hg checkout $REPO_BRANCH
  fi
fi

# Download all the C dependencies
echo "Fetching dependencies..."
BUILD_DEPS=/build_deps.sh
DEPS_FOLDER=/deps
LIST_DEPS=""
mkdir -p $DEPS_FOLDER
DEPS=($DEPS) && for dep in "${DEPS[@]}"; do
  dep_filename=${dep##*/}
  echo "Downloading $dep to $DEPS_FOLDER/$dep_filename"
  wget -q $dep --directory-prefix=$DEPS_FOLDER
  dep_name=$(tar --list --no-recursion --file=$DEPS_FOLDER/$dep_filename  --exclude="*/*" | sed 's/\///g')
  LIST_DEPS="${LIST_DEPS} ${dep_name}"
  if [ "${dep_filename##*.}" == "tar" ]; then tar -xf  $DEPS_FOLDER/$dep_filename --directory $DEPS_FOLDER/  ; fi
  if [ "${dep_filename##*.}" == "gz"  ]; then tar -xzf $DEPS_FOLDER/$dep_filename --directory $DEPS_FOLDER/  ; fi
  if [ "${dep_filename##*.}" == "bz2" ]; then tar -xj  $DEPS_FOLDER/$dep_filename --directory $DEPS_FOLDER/  ; fi
done

# Configure some global build parameters
NAME=`basename $1/$PACK`
if [ "$OUT" != "" ]; then
  NAME=$OUT
fi


if [ "$FLAG_V" == "true" ]; then V=-v; fi
if [ "$FLAG_RACE" == "true" ]; then R=-race; fi
if [ "$STATIC" == "true" ]; then LDARGS=--ldflags\ \'-extldflags\ \"-static\"\'; fi

if [ -n $BEFORE_BUILD ]; then
	chmod +x /scripts/$BEFORE_BUILD
	echo "Execute /scripts/$BEFORE_BUILD ${1}"
	/scripts/$BEFORE_BUILD ${1}
fi


# If no build targets were specified, inject a catch all wildcard
if [ "$TARGETS" == "" ]; then
  TARGETS="./."
fi


built_targets=0
for TARGET in $TARGETS; do
	# Split the target into platform and architecture
	XGOOS=`echo $TARGET | cut -d '/' -f 1`
	XGOARCH=`echo $TARGET | cut -d '/' -f 2`

	# Check and build for Linux targets
	if ([ $XGOOS == "." ] || [ $XGOOS == "linux" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "amd64" ]); then
		echo "Compiling $PACK for linux/amd64..."
		HOST=x86_64-linux PREFIX=/usr/local $BUILD_DEPS /deps $LIST_DEPS
		GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go get -d ./$PACK
		sh -c "GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build $V $R $LDARGS -o $NAME-linux-amd64$R ./$PACK"
		built_targets=$((built_targets+1))
	fi
	if ([ $XGOOS == "." ] || [ $XGOOS == "linux" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "386" ]); then
		echo "Compiling $PACK for linux/386..."
		HOST=i686-linux PREFIX=/usr/local $BUILD_DEPS /deps $LIST_DEPS
		GOOS=linux GOARCH=386 CGO_ENABLED=1 go get -d ./$PACK
		sh -c "GOOS=linux GOARCH=386 CGO_ENABLED=1 go build $V $R $LDARGS -o $NAME-linux-386$R ./$PACK"
		built_targets=$((built_targets+1))
	fi	
	if ([ $XGOOS == "." ] || [ $XGOOS == "linux" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "arm" ]); then
		echo "Compiling $PACK for linux/arm..."
		CC=arm-linux-gnueabi-gcc HOST=arm-linux PREFIX=/usr/local/arm $BUILD_DEPS /deps $LIST_DEPS
		CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm CGO_ENABLED=1 GOARM=5 go get -d ./$PACK
		CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm CGO_ENABLED=1 GOARM=5 go build $V -o $NAME-linux-arm ./$PACK
		built_targets=$((built_targets+1))
	fi
	
	# Check and build for Windows targets
	if ([ $XGOOS == "." ] || [ $XGOOS == "windows" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "amd64" ]); then
		echo "Compiling $PACK for windows/amd64..."
		CC=x86_64-w64-mingw32-gcc HOST=x86_64-w64-mingw32 PREFIX=/usr/x86_64-w64-mingw32 $BUILD_DEPS /deps $LIST_DEPS
		CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go get -d ./$PACK
		CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build  $V $R -o $NAME-windows-amd64$R.exe ./$PACK
		built_targets=$((built_targets+1))
	fi
	
	if ([ $XGOOS == "." ] || [ $XGOOS == "windows" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "386" ]); then
		echo "Compiling $PACK for windows/386..."
		CC=i686-w64-mingw32-gcc HOST=i686-w64-mingw32 PREFIX=/usr/i686-w64-mingw32 $BUILD_DEPS /deps $LIST_DEPS
		CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 CGO_ENABLED=1 go get -d ./$PACK
		CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 CGO_ENABLED=1 go build $V -o $NAME-windows-386.exe ./$PACK
		built_targets=$((built_targets+1))
	fi
	
	# Check and build for OSX targets
	if ([ $XGOOS == "." ] || [ $XGOOS == "darwin" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "amd64" ]); then
		echo "Compiling $PACK for darwin/amd64..."
		CC=o64-clang HOST=x86_64-apple-darwin10 PREFIX=/usr/local $BUILD_DEPS /deps $LIST_DEPS
		CC=o64-clang GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go get -d ./$PACK
		CC=o64-clang GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -ldflags=-s $V $R -o $NAME-darwin-amd64$R ./$PACK
		built_targets=$((built_targets+1))
	fi
	if ([ $XGOOS == "." ] || [ $XGOOS == "darwin" ]) && ([ $XGOARCH == "." ] || [ $XGOARCH == "386" ]); then
		echo "Compiling for darwin/386..."
		CC=o32-clang HOST=i386-apple-darwin10 PREFIX=/usr/local $BUILD_DEPS /deps $LIST_DEPS
		CC=o32-clang GOOS=darwin GOARCH=386 CGO_ENABLED=1 go get -d ./$PACK
		CC=o32-clang GOOS=darwin GOARCH=386 CGO_ENABLED=1 go build $V -o $NAME-darwin-386 ./$PACK
		built_targets=$((built_targets+1))
	fi
done


# The binary files are the last created files
echo "Moving $built_targets $PACK binaries to host folder..."
ls -t | head -n $built_targets 
cp `ls -t | head -n $built_targets ` /build

echo "Build process completed"
