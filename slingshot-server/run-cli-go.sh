#!/bin/bash

go run main.go cli \
--wasm=../examples/go/01-print/print.wasm \
--handler=hello \
--input="🤓 I'm a geek"

