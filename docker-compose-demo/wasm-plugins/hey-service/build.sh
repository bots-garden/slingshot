#!/bin/bash

tinygo build -scheduler=none --no-debug \
    -o hey.wasm \
    -target wasi main.go

ls -lh *.wasm
