#!/bin/bash
# build
cargo clean
cargo build --release --target wasm32-wasi #--offline
# ls -lh *.wasm
ls -lh ./target/wasm32-wasi/release/*.wasm

