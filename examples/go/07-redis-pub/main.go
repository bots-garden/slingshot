package main

import (
	"errors"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostPrintln
func hostPrintln(offset uint64) uint64

func Println(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrintln(memoryText.Offset())
}

//export hostGetEnv
func hostGetEnv(offset uint64) uint64

func GetEnv(name string) string {
	// copy the name of the environment variable to the shared memory
	variableName := pdk.AllocateString(name)
	// call the host function
	offset := hostGetEnv(variableName.Offset())

	// read the value of the result from the shared memory
	variableValue := pdk.FindMemory(offset)
	buffer := make([]byte, variableValue.Length())
	variableValue.Load(buffer)

	// cast the buffer to string and return the value
	envVarValue := string(buffer)
	return envVarValue
}

var parser = fastjson.Parser{}

//export hostInitRedisClient
func hostInitRedisClient(offset uint64) uint64

func InitRedisClient(redisClientId string, redisUri string) (string, error) {

	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "uri": "redis uri"
	// }
	jsonStrArguments := `{"id":"` + redisClientId + `","uri":"` + redisUri + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostInitRedisClient(arguments.Offset())

	// Get result from the shared memory
	// The host function (hostInitRedisClient) returns a JSON buffer:
	// {
	//   "success": "the redis client id",
	//   "failure": "error message if error, else empty"
	// }
	memoryResult := pdk.FindMemory(offset)
	buffer := make([]byte, memoryResult.Length())
	memoryResult.Load(buffer)

	JSONData, err := parser.ParseBytes(buffer)
	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}

}

//export hostRedisPublish
func hostRedisPublish(offset uint64) uint64

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

	JSONData, err := parser.ParseBytes(buffResult)

	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}
}

//export publish
func publish() uint64 {
	input := pdk.Input()

	redisURI := GetEnv("REDIS_URI")
	idRedisClient, errInit := InitRedisClient("pubsubcli", redisURI)
	if errInit != nil {
		Println("😡 " + errInit.Error())
	} else {
		Println("🙂 " + idRedisClient)
	}

	RedisPublish("pubsubcli", "news", string(input))

	

	return 0
}

func main() {}
