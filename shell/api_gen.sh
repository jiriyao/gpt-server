#!/bin/sh
#set -x

if [[ "$(dirname $0)" == "." ]]; then
  cd ..
fi

basedir=$(
    cd $(dirname $0)
    pwd
)

cd ./code/apis
goctl api go -api code.api -dir ../