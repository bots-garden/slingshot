#!/bin/bash
extism call ./target/wasm32-wasi/release/rust_handler_plugin.wasm \
  handle --input "Bob"  \
  --wasi \

echo ""
