package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func helloHandler(input []byte) []byte {
	_, err := slingshot.MemorySet("bob", "Bob Morane")

	value, err := slingshot.MemoryGet("bob")
	if err != nil {
		slingshot.Println("😡 ouch! " + err.Error())
	} else {
		slingshot.Println("🙂 value: " + value)
	}

	value, err = slingshot.MemoryGet("bobby")
	if err != nil {
		slingshot.Println("😡 ouch! " + err.Error())
	} else {
		slingshot.Println("🙂 value: " + value)
	}
	return nil
}

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
	slingshot.ExecHandler(helloHandler)
}

func main() {}

/* 
	./slingshot run \
	--wasm=./memcache.wasm \
	--handler=callHandler

*/