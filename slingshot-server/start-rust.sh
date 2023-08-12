#!/bin/bash
#LD_LIBRARY_PATH=/usr/local/lib go run main.go \
#../rust-handler-plugin/target/wasm32-wasi/release/rust_handler_plugin.wasm \
#handle \
#8080

go run main.go \
../rust-handler-plugin/target/wasm32-wasi/release/rust_handler_plugin.wasm \
handle \
8080
