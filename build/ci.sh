#!/bin/bash

function showMsg() {
  echo -e "\033[32m$1\033[0m"
}

cd ./cmd/safeserver

showMsg 'ci ci go build'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

cd ../../

showMsg 'ci docker build'
docker build . -t registry.cn-chengdu.aliyuncs.com/wmicroservice/hks-core:b0.1 -f ./build/LocalDockerfile

showMsg 'ci clean go build'

showMsg 'ci docker push'
docker push registry.cn-chengdu.aliyuncs.com/wmicroservice/hks-core:b0.1

# showMsg 'ci git push'
# git push


showMsg 'ci all complete'
