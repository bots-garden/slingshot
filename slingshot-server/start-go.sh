#!/bin/bash

go run main.go listen \
--wasm=../experiments/go-handler-plugin/simple.wasm \
--handler=handle \
--http-port=8080

# start and listen are the same command
