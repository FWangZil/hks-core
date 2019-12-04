#!/bin/bash

function showMsg() {
  echo -e "\033[32m$1\033[0m"
}

showMsg 'cd ssh'

# sshpass -p Copell123 ssh root@47.92.39.81 "cd /home/hks;docker pull registry.cn-chengdu.aliyuncs.com/wmicroservice/hks-core:b0.1 ;docker-compose up -d ;exit"
sshpass -p 19951129wzl ssh root@119.3.6.213 "cd /home/hks;docker pull registry.cn-chengdu.aliyuncs.com/wmicroservice/hks-core:b0.1 ;docker-compose up -d ;exit"

showMsg 'cd clean docker build'
docker image rm registry.cn-chengdu.aliyuncs.com/wmicroservice/hks-core:b0.1

showMsg 'cd clean go build'
rm cmd/safeserver/safeserver

showMsg 'cd all complete'
