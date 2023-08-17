#!/bin/bash

go run main.go start \
--wasm=../experiments/js-handler-plugin/handler-js.wasm \
--handler=handle \
--http-port=8080
