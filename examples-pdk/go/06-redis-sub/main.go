package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func MessageHandler(input []byte) []byte {

	slingshot.Print("👋 message: " + string(input))
	return nil

}

func main() {
	slingshot.SetHandler(MessageHandler)
}

/* with the slingshot pdk, always call `callHandler`

    ./slingshot redis subscribe \
    	--wasm=./redissub.wasm \
        --handler=callHandler \
        --uri=${REDIS_URI} \
        --client-id=pubsubcli \
        --channel=news

*/
