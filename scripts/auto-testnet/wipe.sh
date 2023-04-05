#!/bin/bash
set -e 

docker rm -f $(docker ps -a -q)
docker rmi $(docker images -a -q) -f

rm -rf /root/common_store/genesis.json /root/common_store/gentx