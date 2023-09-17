package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func helloHandler(argHandler []byte) []byte {
	input := string(argHandler)
	slingshot.Print("👋 hello world 🌍 " + string(input))
	
	slingshot.Log("🙂 have a nice day 🏖️")

	content, err := slingshot.ReadFile("./hello.txt")
	if err != nil {
		slingshot.Log("😡 " + err.Error())
	}
	slingshot.Print(content)

	return []byte("👋 Hello World 🌍")
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(helloHandler)
}

func main() {}
/*
    ./slingshot run --wasm=./print.wasm \
	--handler=callHandler \
	--input="🤓 I'm a geek"

*/