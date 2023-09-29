package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func helloHandler(input []byte) []byte {
	message := slingshot.GetEnv("MESSAGE")
	slingshot.Println("🤖 MESSAGE=" + message)
	
	//return []byte("hello") // will println hello
	return nil // will println nothing

}
// TODO: void handler?
// Or test if return something or not

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
	slingshot.ExecHandler(helloHandler)
}

func main() {}

/* 
    MESSAGE="👋 Hello World 🌍" \
    ./slingshot run --wasm=./envvar.wasm \
	--handler=callHandler
  
*/