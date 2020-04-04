#!/bin/bash

VER=$(cat ./VERSION)

docker rmi -f framework:$VER
docker build --tag framework:$VER .
