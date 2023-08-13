#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o some-functions.wasm \
  -target wasi main.go

ls -lh *.wasm
