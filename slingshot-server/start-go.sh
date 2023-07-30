#!/bin/bash
#LD_LIBRARY_PATH=/usr/local/lib go run main.go \
#../go-handler-plugin/simple.wasm \
#handle \
#8080


#LD_LIBRARY_PATH=/usr/local/lib \
./slingshot-v0.0.0-linux-arm64 \
../go-handler-plugin/simple.wasm \
handle \
8080
