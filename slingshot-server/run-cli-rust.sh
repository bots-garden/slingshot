#!/bin/bash


go run main.go cli \
--wasm=../examples/rust/01-print/print.wasm \
--handler=hello \
--input="🤓 I'm a geek"

