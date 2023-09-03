package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func messageHandler(input []byte) []byte {

	slingshot.Print("👋 message: " + string(input))
	return nil

}

//export callHandler
func callHandler() {
	slingshot.Print("👋 callHandler function")
	slingshot.ExecHandler(messageHandler)
}

func main() {}

/*
    ./slingshot redis subscribe \
    	--wasm=./redissub.wasm \
        --handler=callHandler \
        --uri=${REDIS_URI} \
        --client-id=pubsubcli \
        --channel=news

*/
