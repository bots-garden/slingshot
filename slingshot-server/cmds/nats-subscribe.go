package cmds

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slingshot-server/clients"
	"slingshot-server/plg"
	"slingshot-server/slingshot"

	"github.com/nats-io/nats.go"
)

// NatsSubscribe is triggered by the `nats subscribe` command (from parseCommand)
func NatsSubscribe(wasmFilePath string, wasmFunctionName string, natsSubject string, natsUrl string, natsClientId string) {
	natsConfig := slingshot.NatsConfig{
		Url: natsUrl,
	}

	natsConnection, err := clients.CreateNatsConnection(natsConfig)
	// Close the connection when we are done.

	defer natsConnection.Close()

	if err != nil {
		log.Println("🔴 Error when connecting with the NATS server", err)
		os.Exit(1)
	}

	ctx := plg.Initialize("slingshotplug", wasmFilePath)

	go func() {
		// Simple Async Subscriber: natsSubscription
		_, err := natsConnection.Subscribe(natsSubject, func(msg *nats.Msg) {

			mutex.Lock()
			// don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
			defer mutex.Unlock()

			extismPlugin, err := plg.GetPlugin("slingshotplug")
			if err != nil {
				log.Println("🔴 Error when getting the plugin", err)
				os.Exit(1)
			}

			natsMessage := slingshot.NatsSubscribeMessage{
				Subject: msg.Subject,
				Data:    string(msg.Data),
			}
			jsonBytes, err := json.Marshal(&natsMessage)
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

		})
		if err != nil {
			log.Println("🔴 Error when subscribing", err)
			os.Exit(1)
		}

	}()
	<-ctx.Done()
}
