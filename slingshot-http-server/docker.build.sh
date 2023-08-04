#!/bin/bash
IMAGE_NAME="demo-slingshot"
docker build -t ${IMAGE_NAME} . 

docker images | grep ${IMAGE_NAME}
