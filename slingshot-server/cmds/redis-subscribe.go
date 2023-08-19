package cmds

import (
	"fmt"
	"log"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/plg"
)

// Subscribe is triggered by the `redis subscribe` command (from parseCommand)
func Subscribe(wasmFilePath string, wasmFunctionName string, redisChannel string, redisUri string, redisClientId string) {

	redisClientRecord := callbacks.RedisClientRecord{
		Id:  redisClientId,
		Uri: redisUri,
	}
	redisClient, err := callbacks.CreateOrGetRedisClient(redisClientRecord)
	if err != nil {
		log.Println("🔴 Error when connecting with the redis database", err)
		os.Exit(1)
	}

	ctx := plg.Initialize("slingshotplug", wasmFilePath)
	plugin, err := plg.GetPlugin("slingshotplug")

	if err != nil {
		log.Println("🔴 Error when getting the plugin", err)
		os.Exit(1)
	}

	if plugin.FunctionExists(wasmFunctionName) != true {
		log.Println("🔴 Error:", wasmFunctionName, "does not exist")
		os.Exit(1)
	}

	// There is no error because go-redis
	// automatically reconnects on error.
	// 🤔
	pubsub := redisClient.Subscribe(ctx, redisChannel)

	// Close the subscription when we are done.
	defer pubsub.Close()

	go func() {
		ch := pubsub.Channel()

		for msg := range ch {
			//fmt.Println(msg.Channel, msg.Payload)

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
