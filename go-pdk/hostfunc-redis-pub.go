package slingshot

import (
	"errors"

	"github.com/extism/go-pdk"
)

// hostRedisPublish:
/*
  This exported function will call the host function callback `RedisPublish` of the slingshot application
*/
//export hostRedisPublish
func hostRedisPublish(offset uint64) uint64

// RedisPublish: publish a message on a channel of a Redis server
/*
	- This helper call the `hostRedisPublish` exported function
	- It copies the `redisClientId`, `channel` and `payload` parameters to the shared memory
	- It calls the `RedisPublish` host function callback (when calling `hostRedisPublish`)
	- It reads the result of the callback
	- And returns this result
*/
func RedisPublish(redisClientId string, channel string, payload string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "channel": "name",
	//    "payload": "Bob Morane"
	// }
	jsonStr := `{"id":"` + redisClientId + `","channel":"` + channel + `","payload":"` + payload + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)

	// Call host function with the offset of the arguments
	offset := hostRedisPublish(arguments.Offset())

	// Get result from the shared memory
	// The host function (hostMemorySet) returns a JSON buffer:
	// {
	//   "success": "the value associated to the key",
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
