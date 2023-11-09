package callbacks

import (
	"context"
	"log"
	"slingshot-server/clients"
	"slingshot-server/mem"
	"slingshot-server/slingshot"

	extism "github.com/extism/go-sdk"

)

func InitNatsConnection(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	/* Expected
	{
		id:""
		uri:""
	}
	*/
	var result = slingshot.StringResult{}
	var config slingshot.NatsConfig
	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &config)

	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		_, err := clients.CreateOrGetNatsConnection(config)
		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		}
		result.Failure = ""
		result.Success = config.Id
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 InitNatsConnection, CopyJsonToMemory:", err)
	}

}

func NatsPublish(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {

	/* Expected
	{ id: "", subject: "", data: "" }
	*/
	var result = slingshot.StringResult{}
	var message slingshot.NatsPublishMessage

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &message)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		natsConnection := clients.GetNatsConnection(message.Id)

		//natsConnection, errConn := clients.CreateOrGetNatsConnection(slingshot.NatsConfig{Id: message.Id})
		//natsConnection, errConn := clients.CreateNatsConnection(slingshot.NatsConfig{Url: message.Url})

		defer natsConnection.Close()

		//fmt.Println("🟣", message.Subject, message.Data)
		err := natsConnection.Publish(message.Subject, []byte(message.Data))

		if err != nil {
			result.Failure = err.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = "ok"
		}
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 NatsPublish, CopyJsonToMemory:", err)
	}
}
