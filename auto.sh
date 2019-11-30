#!/bin/bash
function showMsg() {
  echo -e "\033[32m$1\033[0m"
}
showMsg '开始自动构建镜像并推送到阿里云仓库'
sh ./build/ci.sh
showMsg '自动部署并清除本地构建产物'
sh ./build/cd.sh