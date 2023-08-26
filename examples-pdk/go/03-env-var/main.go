package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func Handler(input []byte) []byte {
	message := slingshot.GetEnv("MESSAGE")
	slingshot.Print("🤖 MESSAGE=" + message)
	
	//return []byte("hello") // will print hello
	return nil // will print nothing

}
// TODO: void handler?
// Or test if return something or not

func main() {
	slingshot.SetHandler(Handler)
}

/* with the slingshot pdk, always call `callHandler`

    MESSAGE="👋 Hello World 🌍" \
    ./slingshot run --wasm=./envvar.wasm \
	--handler=callHandler
  
*/