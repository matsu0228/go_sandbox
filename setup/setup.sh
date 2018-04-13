#!/bin/bash

# setup for cmd directory
# GOPATH/src/github/matsu0228 にシンボリックリンクをはる

FIRST_GOPATH=`echo $GOPATH | awk -F ":" '{ print $1 }'`
COMMAND_ROOT=$(cd $(dirname $0)/../; pwd)


if [ -d ${COMMAND_ROOT} ]; then
  ln -s ${COMMAND_ROOT} ${FIRST_GOPATH}/src/github.com/matsu0228
  echo "setupしました: ${FIRST_GOPATH}/src/github.com/matsu0228"
else
  echo "${COMMAND_ROOT} が存在しないため、setup完了していません"
  exit 1
fi
