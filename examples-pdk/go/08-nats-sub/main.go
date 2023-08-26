package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func MessageHandler(input []byte) []byte {

	slingshot.Print("👋 NATS message: " + string(input))
	return nil

}

func main() {
	slingshot.SetHandler(MessageHandler)
}

/* with the slingshot pdk, always call `callHandler`

     ./slingshot nats subscribe \
        --wasm=./natssub.wasm \
        --handler=callHandler \
        --url=nats://0.0.0.0:4222 \
        --connection-id=natsconn01 \
        --subject=news

*/