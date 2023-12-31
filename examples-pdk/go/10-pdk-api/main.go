package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
	"github.com/extism/go-pdk"
)

func helloHandler(argHandler []byte) []byte {
	input := string(argHandler)
	slingshot.Println("👋 hello world 🌍 " + string(input))

	slingshot.Log("🙂 have a nice day 🏖️")

	pdk.Log(pdk.LogInfo, "😀😃😄")

	firstName, _ := pdk.GetConfig("firstName")
	lastName, _ := pdk.GetConfig("lastName")

	pdk.Log(pdk.LogInfo, firstName)
	pdk.Log(pdk.LogInfo, lastName)

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
