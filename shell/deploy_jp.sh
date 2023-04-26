#!/bin/sh
#set -x

if [[ "$(dirname $0)" == "." ]]; then
  cd ..
fi

basedir=$(
    cd $(dirname $0)
    pwd
)

cd ./code

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -v -o ./target/chatgptcode
cp -rf ./etc/code-api-prod.yaml ./target/etc/code-api.yaml

local_path="/Users/Jerry/Documents/my/chatgpt/chatgpt-server/code/target"
remote_path="/home/proj/chatgptserver/"

cd $local_path
hosts=(54.249.196.123)
#per host start
i=0
j=${#hosts[@]}
# i < j
while [ "$i" -lt "$j" ]
do


ssh root@${hosts[$i]} "cd /home/proj/chatgptserver && killall chatgptcode"
# scp files
scp -r ${local_path}/* root@${hosts[$i]}:${remote_path}
ssh root@${hosts[$i]} "cd /home/proj/chatgptserver && nohup ./chatgptcode & && exit"


echo ${hosts[$i]} done!
let i=i+1
done
