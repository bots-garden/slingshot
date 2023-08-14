#!/bin/bash
#LD_LIBRARY_PATH=/usr/local/lib go run main.go \
#../go-handler-plugin/simple.wasm \
#handle \
#8080

#./slingshot-v0.0.0-linux-arm64 \
#../go-handler-plugin/simple.wasm \
#handle \
#8080

#./slingshot-v0.0.0-linux-arm64 \
go run main.go start \
--wasm=../go-handler-plugin/simple.wasm \
--handler=handle \
--http-port=8080
