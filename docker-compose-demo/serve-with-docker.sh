#!/bin/bash
DOCKER_USER="botsgarden"
IMAGE_NAME="slingshot"
IMAGE_TAG="0.0.5"
echo "🖼️ ${IMAGE_NAME}"
HTTP_PORT="7070"

#docker run \
#  -p ${HTTP_PORT}:${HTTP_PORT} \
#  -v $(pwd):/app --rm ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG} \
#  /slingshot start \
#  --wasm=./app/hello.wasm \
#  --handler=callHandler \
#  --http-port=${HTTP_PORT} 

docker run \
  -p ${HTTP_PORT}:${HTTP_PORT} \
  -v $(pwd):/app/wasm-plugins --rm ${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG} \
  /slingshot start \
  --wasm=./app/wasm-plugins/hello.wasm \
  --handler=handle \
  --http-port=${HTTP_PORT} 