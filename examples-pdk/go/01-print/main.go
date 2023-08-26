package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func Handler(argHandler []byte) []byte {
	input := string(argHandler)
	slingshot.Print("👋 hello world 🌍 " + string(input))
	
	slingshot.Log("🙂 have a nice day 🏖️")
	//TODO: set header
	return []byte("👋 Hello World 🌍")
}

func main() {
	slingshot.Print("👋 main function")
	slingshot.SetHandler(Handler)
}

/* with the slingshot pdk, always call `callHandler`
    ./slingshot run --wasm=./print.wasm \
	--handler=callHandler \
	--input="🤓 I'm a geek"

*/