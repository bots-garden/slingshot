package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func PublishHandler(input []byte) []byte {

	natsURL := slingshot.GetEnv("NATS_URL")
	slingshot.Print("💜 NATS_URL: " + natsURL)
	idNatsConnection, errInit := slingshot.InitNatsConnection("natsconn01", natsURL)
	if errInit != nil {
		slingshot.Print("😡 " + errInit.Error())
	} else {
		slingshot.Print("🙂 " + idNatsConnection)
	}

	res, err := slingshot.NatsPublish("natsconn01", "news", string(input))

	if err != nil {
		slingshot.Print("😡 " + err.Error())
	} else {
		slingshot.Print("🙂 " + res)
	}
	return nil
}

func main() {
	slingshot.SetHandler(PublishHandler)
}

/* with the slingshot pdk, always call `callHandler`

    ./slingshot run --wasm=./natspub.wasm \
        --handler=callHandler \
        --input="I 💜 Wasm ✨"

*/
