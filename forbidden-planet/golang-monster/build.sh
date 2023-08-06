#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o golang-monster.wasm \
  -target wasi main.go

ls -lh *.wasm
