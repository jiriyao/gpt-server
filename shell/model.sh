#!/bin/sh
#set -x

if [[ "$(dirname $0)" == "." ]]; then
  cd ..
fi

basedir=$(
  cd $(dirname $0)
  pwd
)

cd ./pay/internal/model

#table name
TBNAME=$1
SRCURL=$2

if [ "$TBNAME" = "" ]; then
  echo "error: table name is empty"
  exit 0
fi

if [ -e "$TBNAME" ]; then
  rm -rf "$TBNAME"/*_gen.go
  mkdir "$TBNAME"
  # shellcheck disable=SC2086
  chmod 777 $TBNAME
  echo '已删除并重新创建目录'
else
  mkdir "$TBNAME"
  chmod 777 "$TBNAME"
  echo '目录不存在，已创建'
fi

cd ./"$TBNAME"

if [ "$SRCURL" = "" ]; then
  SRCURL="root:b36da6b4eb0f3@tcp(51.159.14.254:30060)/db_pay"
fi


echo "当前数据库连接：$SRCURL 表：$TBNAME"
goctl model mysql -i id datasource -url="$SRCURL" -table="$TBNAME" -c -dir .

sed -i '' "s/cache:/pay:cache:/g" *_gen.go
