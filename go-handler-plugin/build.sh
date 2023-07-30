#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o simple.wasm \
  -target wasi main.go

ls -lh *.wasm
