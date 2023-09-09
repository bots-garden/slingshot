package slingshot

import (
	"errors"

	"github.com/extism/go-pdk"
)

// hostInitNatsConnection:
/*
  This exported function will call the host function callback `InitNatsConnection` of the slingshot application
*/
//export hostInitNatsConnection
func hostInitNatsConnection(offset uint64) uint64

// InitNatsConnection: initialize a NATS connection
/*
	- This helper call the `hostInitNatsConnection` exported function
	- It copies the `natsConnectionId` and `natsUrl` parameters to the shared memory
	- It calls the `InitNatsConnection` host function callback (when calling `hostInitNatsConnection`)
	- It reads the result of the callback
	- And returns this result

  This function will create a connection and store it with an id (`natsConnectionId`),
  that means you can retrieve the connection by id (for example with the `NatsPublish` function)
*/
func InitNatsConnection(natsConnectionId string, natsUrl string) (string, error) {
	jsonStrArguments := `{"id":"` + natsConnectionId + `","url":"` + natsUrl + `"}`
	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostInitNatsConnection(arguments.Offset())

	// Get result from the shared memory
	// The host function (hostInitNatsConnection) returns a JSON buffer:
	// {
	//   "success": "the NATS connexion id",
	//   "failure": "error message if error, else empty"
	// }
	memoryResult := pdk.FindMemory(offset)
	buffer := make([]byte, memoryResult.Length())
	memoryResult.Load(buffer)

	JSONData, err := GetJsonFromBytes(buffer)
	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}

}

// hostNatsPublish:
/*
  This exported function will call the host function callback `NatsPublish` of the slingshot application
*/
//export hostNatsPublish
func hostNatsPublish(offset uint64) uint64

// NatsPublish: publish a message on a subject of a NATS server
/*
	- This helper call the `hostNatsPublish` exported function
	- It copies the `natsConnectionId`, `subject` and `data` parameters to the shared memory
	- It calls the `NatsPublish` host function callback (when calling `hostNatsPublish`)
	- It reads the result of the callback
	- And returns this result
*/
func NatsPublish(natsConnectionId string, subject string, data string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the NATS client",
	//    "subject": "name",
	//    "data": "Bob Morane"
	// }
	//jsonStr := `{"url":"` + natsServer + `","subject":"` + subject + `","data":"` + data + `"}`
	jsonStr := `{"id":"` + natsConnectionId + `","subject":"` + subject + `","data":"` + data + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)

	// Call host function with the offset of the arguments
	offset := hostNatsPublish(arguments.Offset())

	// Get result from the shared memory
	// The host function (hostNatsPublish) returns a JSON buffer:
	// {
	//   "success": "the message",
	//   "failure": "error message if error, else empty"
	// }
	memoryResult := pdk.FindMemory(offset)
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	JSONData, err := GetJsonFromBytes(buffResult)

	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}
}
