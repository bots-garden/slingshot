package cmds

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slingshot-server/clients"
	"slingshot-server/plg"
	"slingshot-server/slingshot"
)

// RedisSubscribe is triggered by the `redis subscribe` command (from parseCommand)
func RedisSubscribe(wasmFilePath string, wasmFunctionName string, redisChannel string, redisUri string, redisClientId string) {

	redisConfig := slingshot.RedisConfig{
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

			extismPlugin, err := plg.GetPlugin("slingshotplug")
			if err != nil {
				log.Println("🔴 Error when getting the plugin", err)
				os.Exit(1)
			}

			redisMessage := slingshot.RedisMessage{
				Channel: msg.Channel,
				Payload: msg.Payload,
				Id: redisClientId,
			}
			jsonBytes, err := json.Marshal(&redisMessage)
			if err != nil {
				log.Println("🔴 Error:", err)
			}

			if extismPlugin.MainFunction == true {
				_, _, err := extismPlugin.Plugin.Call("_start", nil)
				if err != nil {
					log.Println("🔴 Error with _start function", err)
				}
			}

			_, output, err := extismPlugin.Plugin.Call(wasmFunctionName, jsonBytes)
			if err != nil {
				log.Println("🔴 Error:", err)
				//os.Exit(1)
			}
			// Display output content, only if the wasm plugin returns something
			if (len(output)) > 0 {
				// CLI output
				fmt.Println(string(output))
			}

		}
	}()
	<-ctx.Done()

}
