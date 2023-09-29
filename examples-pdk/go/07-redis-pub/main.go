package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func publishHandler(input []byte) []byte {

	redisURI := slingshot.GetEnv("REDIS_URI")
	idRedisClient, errInit := slingshot.InitRedisClient("pubsubcli", redisURI)
	if errInit != nil {
		slingshot.Println("😡 " + errInit.Error())
	} else {
		slingshot.Println("🙂 " + idRedisClient)
	}

	slingshot.RedisPublish("pubsubcli", "news", string(input))

	return nil
}

//export callHandler
func callHandler() {
	slingshot.Println("👋 callHandler function")
	slingshot.ExecHandler(publishHandler)
}

func main() {}

/*
    ./slingshot run --wasm=./redispub.wasm \
        --handler=callHandler \
        --input="I 💜 Wasm ✨"

*/
