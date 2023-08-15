#!/bin/bash


go run main.go cli \
--wasm=../00-print-rust/print.wasm \
--handler=hello \
--input="🤓 I'm a geek"

