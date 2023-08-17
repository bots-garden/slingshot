#!/bin/bash

go run main.go start \
--wasm=../experiments/rust-handler-plugin/target/wasm32-wasi/release/rust_handler_plugin.wasm \
--handler=handle \
--http-port=8080