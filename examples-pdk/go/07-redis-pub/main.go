package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func publishHandler(input []byte) []byte {

	redisURI := slingshot.GetEnv("REDIS_URI")
	idRedisClient, errInit := slingshot.InitRedisClient("pubsubcli", redisURI)
	if errInit != nil {
		slingshot.Print("😡 " + errInit.Error())
	} else {
		slingshot.Print("🙂 " + idRedisClient)
	}

	slingshot.RedisPublish("pubsubcli", "news", string(input))

	return nil
}

//export callHandler
func callHandler() {
	slingshot.Print("👋 callHandler function")
	slingshot.ExecHandler(publishHandler)
}

func main() {}

/* with the slingshot pdk, always call `callHandler`

    ./slingshot run --wasm=./redispub.wasm \
        --handler=callHandler \
        --input="I 💜 Wasm ✨"

*/
