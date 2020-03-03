#!/bin/bash

# builds and installs dockma globally
# tested on macOS only

if [ -z "$DOCKMA_HOME" ]
then
      echo "\$DOCKMA_HOME not set"
      exit 0
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd -P)"
cd $SCRIPT_DIR/..
DOCKMA_BIN="$(pwd -P)/builds/dockma"

make build

echo "generating completion code for bash"
./builds/dockma completion bash > $DOCKMA_HOME/bash_completion

echo "linking bin and completion"
ln -s $DOCKMA_BIN /usr/local/bin/dockma
ln -s $DOCKMA_HOME/bash_completion /usr/local/etc/bash_completion.d/dockma

