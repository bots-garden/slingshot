#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o read-env.wasm \
  -target wasi main.go

ls -lh *.wasm
