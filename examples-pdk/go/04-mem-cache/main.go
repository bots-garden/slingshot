package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func Handler(input []byte) []byte {
	_, err := slingshot.MemorySet("bob", "Bob Morane")

	value, err := slingshot.MemoryGet("bob")
	if err != nil {
		slingshot.Print("😡 ouch! " + err.Error())
	} else {
		slingshot.Print("🙂 value: " + value)
	}

	value, err = slingshot.MemoryGet("bobby")
	if err != nil {
		slingshot.Print("😡 ouch! " + err.Error())
	} else {
		slingshot.Print("🙂 value: " + value)
	}
	return nil
}

func main() {
	slingshot.SetHandler(Handler)
}

/* with the slingshot pdk, always call `callHandler`
	
	./slingshot run \
	--wasm=./memcache.wasm \
	--handler=callHandler

*/