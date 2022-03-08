#!/bin/bash

case "$1" in
  "start"|"up")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server up -d;;
  "stop"|"down")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server down;;
  "restart")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server down
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server up -d
   ;;
  "ps")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server ps;;
  "logs")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server logs ${@:2};;
  "exec")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server exec ${2} ${@:3};;
  "bash")
    docker-compose -f docker-compose.yml -p shuoshuo-daxigua-server exec ${2} bash;;
  default)
    echo "不支持的命令";;
esac