#!/bin/bash


BIN_PATH=".cli-unit/cli-unit"

if [ ! -f $BIN_PATH ]; then
	mkdir -p $(dirname $BIN_PATH)
	

    wget -qO- https://github.com/benschw/cli-unit/releases/download/0.1.0/cli-unit-`uname`-`uname -m`-0.1.0.gz | \
        gunzip > $BIN_PATH


    chmod +x $BIN_PATH
fi

ARGS=( "$@" )

./$BIN_PATH ${ARGS[@]}
