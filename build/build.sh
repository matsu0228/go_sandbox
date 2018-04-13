#!/bin/bash

COMMAND_NAME=go-head
# build
go build -o ${COMMAND_NAME} main.go

# GOPATH/bin に格納
FIRST_GOPATH=`echo $GOPATH | awk -F ":" '{ print $1 }'`
COMMAND_ROOT=$(cd $(dirname $0)/../; pwd)
COMMAND=${COMMAND_ROOT}/${COMMAND_NAME}

if [ -f ${COMMAND} ]; then
  if [ -L ${FIRST_GOPATH}/bin/${COMMAND_NAME} ]; then
    unlink ${FIRST_GOPATH}/bin/${COMMAND_NAME}
  fi
  ln -s ${COMMAND} ${FIRST_GOPATH}/bin/
  echo "goコマンドを利用可能にしました: ${FIRST_GOPATH}/bin/"
else
  echo "${COMMAND} が存在しないため、setup完了していません"
  exit 1
fi
