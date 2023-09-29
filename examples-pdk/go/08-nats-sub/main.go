package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func messageHandler(input []byte) []byte {

	slingshot.Println("👋 NATS message: " + string(input))
	return nil

}

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
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