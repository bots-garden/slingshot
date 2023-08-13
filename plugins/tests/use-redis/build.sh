#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o use-redis.wasm \
  -target wasi main.go

ls -lh *.wasm
