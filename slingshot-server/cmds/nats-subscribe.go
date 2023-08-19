package cmds

import "fmt"

// NatsSubscribe is triggered by the `nats subscribe` command (from parseCommand)
func NatsSubscribe(wasmFilePath string, wasmFunctionName string, natsSubject string, natsUrl string, natClientId string) {
	fmt.Println("I 💜 Nats")
	
}
