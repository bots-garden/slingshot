#!/bin/bash
extism call ./simple.wasm \
  handle --input "Bob Morane" \
  --wasi

echo ""

# _start allows to call the main function