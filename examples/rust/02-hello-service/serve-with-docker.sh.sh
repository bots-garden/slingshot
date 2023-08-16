#!/bin/bash

operating_system=$(uname -o) # Darwin GNU/Linux 
processor=$(uname -p) # x86_64 arm aarch64

if [ $operating_system == "GNU/Linux" ]; then
  operating_system="linux"
fi

if [ $operating_system == "Darwin" ]; then
  operating_system="darwin"
else
  operating_system="linux"
fi

if [ $processor == "x86_64" ]; then
  processor="amd64"
fi

if [ $processor == "arm" ]; then
  processor="arm64"
fi

if [ $processor == "aarch64" ]; then
  processor="arm64"
else
  processor="amd64"
fi

rm slingshot

IMAGE_NAME="slingshot-${operating_system}-${processor}"
IMAGE_TAG="0.0.0"
echo "🖼️ ${IMAGE_NAME}"
HTTP_PORT="7070"

docker run \
  -p ${HTTP_PORT}:${HTTP_PORT} \
  -v $(pwd):/app --rm ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG} \
  /slingshot start \
  --wasm=./app/hello.wasm \
  --handler=handle \
  --http-port=${HTTP_PORT} 


