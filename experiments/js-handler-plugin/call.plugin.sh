#!/bin/bash
extism call ./handler-js.wasm \
  handle --input "😀 Hello World 🌍! (from JavaScript)" \
  --wasi \
  --log-level info

echo ""