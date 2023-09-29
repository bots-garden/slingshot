package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func publishHandler(input []byte) []byte {

	natsURL := slingshot.GetEnv("NATS_URL")
	slingshot.Println("💜 NATS_URL: " + natsURL)
	idNatsConnection, errInit := slingshot.InitNatsConnection("natsconn01", natsURL)
	if errInit != nil {
		slingshot.Println("😡 " + errInit.Error())
	} else {
		slingshot.Println("🙂 " + idNatsConnection)
	}

	res, err := slingshot.NatsPublish("natsconn01", "news", string(input))

	if err != nil {
		slingshot.Println("😡 " + err.Error())
	} else {
		slingshot.Println("🙂 " + res)
	}
	return nil
}

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
	slingshot.ExecHandler(publishHandler)
}

func main() {}

/*
    ./slingshot run --wasm=./natspub.wasm \
        --handler=callHandler \
        --input="I 💜 Wasm ✨"
*/
