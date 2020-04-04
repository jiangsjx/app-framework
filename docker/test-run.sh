#!/bin/bash

VER=$(cat ./VERSION)
CWD=$(pwd)

docker rm -f framework
docker run -d \
    --name=framework \
    --network=host \
    -v $CWD:/etc/framework/ \
    -v $CWD:/var/log/framework/ \
    framework:$VER
