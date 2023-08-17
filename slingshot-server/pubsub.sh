#!/bin/bash
set -o allexport; source .env; set +o allexport
# you need these environment variables
# - REDIS_URI
#echo "REDIS_URI: ${REDIS_URI}"

go run main.go redis subscribe \
--wasm=../examples/go/06-redis-sub/redissub.wasm \
--handler=message \
--redis-uri=${REDIS_URI} \
--redis-client-id=pubsubcli \
--channel=news
