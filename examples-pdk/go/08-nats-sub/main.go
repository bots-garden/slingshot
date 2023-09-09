package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func messageHandler(input []byte) []byte {

	slingshot.Print("👋 NATS message: " + string(input))
	return nil

}

//export callHandler
func callHandler() {
	slingshot.Print("👋 callHandler function")
	slingshot.ExecHandler(messageHandler)
}

func main() {}

/* 
     ./slingshot nats subscribe \
        --wasm=./natssub.wasm \
        --handler=callHandler \
        --url=nats://0.0.0.0:4222 \
        --connection-id=natsconn01 \
        --subject=news
*/