package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func helloHandler(argHandler []byte) []byte {
	input := string(argHandler)
	slingshot.Println("👋 hello world 🌍 " + string(input))
	
	slingshot.Log("🙂 have a nice day 🏖️")

	return []byte("👋 Hello World 🌍")
}

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
	slingshot.ExecHandler(helloHandler)
}

func main() {}
/*
    ./slingshot run --wasm=./println.wasm \
	--handler=callHandler \
	--input="🤓 I'm a geek"

*/