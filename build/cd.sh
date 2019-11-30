#!/bin/bash

function showMsg() {
  echo -e "\033[32m$1\033[0m"
}

showMsg 'cd ssh'

sshpass -p "password" ssh root@"ip" "cd /home/hks;docker pull "docker address" ;docker-compose up -d ;exit"
# docker logs --tail 100 -f dmsj_apigatecard_1


showMsg 'cd clean docker build'
docker image rm "docker address"

showMsg 'cd clean go build'

showMsg 'cd all complete'
