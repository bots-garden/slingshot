package cmds

import (
	"fmt"
	"log"
	"os"
	"slingshot-server/clients"
	"slingshot-server/plg"
	"slingshot-server/slingshot"
)

// RedisSubscribe is triggered by the `redis subscribe` command (from parseCommand)
func RedisSubscribe(wasmFilePath string, wasmFunctionName string, redisChannel string, redisUri string, redisClientId string) {

	redisConfig := slingshot.RedisClientConfig{
		Id:  redisClientId,
		Uri: redisUri,
	}
	redisClient, err := clients.CreateOrGetRedisClient(redisConfig)
	if err != nil {
		log.Println("🔴 Error when connecting with the redis database", err)
		os.Exit(1)
	}

	ctx := plg.Initialize("slingshotplug", wasmFilePath)

	// There is no error because go-redis
	// automatically reconnects on error.
	// 🤔
	pubsub := redisClient.Subscribe(ctx, redisChannel)

	// Close the subscription when we are done.
	defer pubsub.Close()

	go func() {
		ch := pubsub.Channel()

		for msg := range ch { // this is synchronous, no need mutex (apparently)

			plugin, err := plg.GetPlugin("slingshotplug")
			if err != nil {
				log.Println("🔴 Error when getting the plugin", err)
				os.Exit(1)
			}

			// ? how to stress Redis pub sub?
			// TODO: Create a Json Payload
			_, output, err := plugin.Call(wasmFunctionName, []byte(msg.Channel+" "+msg.Payload))
			if err != nil {
				fmt.Println("🔴 Error:", err)
				//os.Exit(1)
			}
			// Display output content, only if the wasm plugin returns something
			if (len(output)) > 0 {
				fmt.Println(string(output))
			}

		}
	}()
	<-ctx.Done()

}
