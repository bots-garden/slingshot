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
		Id:  natsClientId,
		Url: natsUrl,
	}

	natsClient, err := clients.CreateOrGetNatsClient(natsConfig)
	if err != nil {
		log.Println("🔴 Error when connecting with the NATS server", err)
		os.Exit(1)
	}
	//natsClient := clients.GetNatsClient("natscli")

	ctx := plg.Initialize("slingshotplug", wasmFilePath)

	// Close the connection when we are done.
	defer natsClient.Close()

	go func() {
		// Simple Async Subscriber: natsSubscription
		_, err := natsClient.Subscribe(natsSubject, func(msg *nats.Msg) {

			mutex.Lock()
			// don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
			defer mutex.Unlock()

			plugin, err := plg.GetPlugin("slingshotplug")
			if err != nil {
				log.Println("🔴 Error when getting the plugin", err)
				os.Exit(1)
			}

			natsMessage := slingshot.NatsMessage{
				Subject: msg.Subject,
				Data:    string(msg.Data),
				Id: natsClientId,
			}
			jsonBytes, err := json.Marshal(&natsMessage)
			if err != nil {
				fmt.Println("🔴 Error:", err)
			}

			_, output, err := plugin.Call(wasmFunctionName, jsonBytes)
			if err != nil {
				fmt.Println("🔴 Error:", err)
				//os.Exit(1)
			}
			// Display output content, only if the wasm plugin returns something
			if (len(output)) > 0 {
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
