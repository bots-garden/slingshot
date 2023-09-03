package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func helloHandler(input []byte) []byte {
	message := slingshot.GetEnv("MESSAGE")
	slingshot.Print("🤖 MESSAGE=" + message)
	
	//return []byte("hello") // will print hello
	return nil // will print nothing

}
// TODO: void handler?
// Or test if return something or not

//export callHandler
func callHandler() {
	slingshot.Print("👋 callHandler function")
	slingshot.ExecHandler(helloHandler)
}

func main() {}

/* 
    MESSAGE="👋 Hello World 🌍" \
    ./slingshot run --wasm=./envvar.wasm \
	--handler=callHandler
  
*/