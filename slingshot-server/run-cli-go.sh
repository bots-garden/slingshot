#!/bin/bash

go run main.go cli \
--wasm=../00-print-go/print.wasm \
--handler=hello \
--input="🤓 I'm a geek"

