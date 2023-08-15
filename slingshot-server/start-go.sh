#!/bin/bash

go run main.go start \
--wasm=../go-handler-plugin/simple.wasm \
--handler=handle \
--http-port=8080
