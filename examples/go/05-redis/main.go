package main

import (
	"errors"
	"strings"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostPrint
func hostPrint(offset uint64) uint64

func Print(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrint(memoryText.Offset())
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

//export hostRedisSet
func hostRedisSet(offset uint64) uint64

func RedisSet(redisClientId string, key string, value string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name",
	//    "value": "Bob Morane"
	// }
	jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `","value":"` + value + `"}`
	
	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)

	// Call host function with the offset of the arguments
	offset := hostRedisSet(arguments.Offset())

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

//export hostRedisGet
func hostRedisGet(offset uint64) uint64

func RedisGet(redisClientId string, key string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name"
	// }
	jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)
	// Call the host function
	offset := hostRedisGet(arguments.Offset())

	// Get result (the value associated to the key) from shared memory
	// The host function (hostRedisGet) returns a JSON buffer:
	// {
	//   "success": "the value associated to the key",
	//   "failure": "error message if error, else empty"
	// }
	memoryValue := pdk.FindMemory(offset)
	buffer := make([]byte, memoryValue.Length())
	memoryValue.Load(buffer)

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

//export hostRedisDel
func hostRedisDel(offset uint64) uint64

func RedisDel(redisClientId string, key string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "name"
	// }
	jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)
	// Call the host function
	offset := hostRedisDel(arguments.Offset())

	// Get result (the value associated to the key) from shared memory
	// The host function (hostRedisDel) returns a JSON buffer:
	// {
	//   "success": "the value associated to the key",
	//   "failure": "error message if error, else empty"
	// }
	memoryValue := pdk.FindMemory(offset)
	buffer := make([]byte, memoryValue.Length())
	memoryValue.Load(buffer)

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

//export hostRedisFilter
func hostRedisFilter(offset uint64) uint64

func RedisFilter(redisClientId string, key string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the redis client",
	//    "key": "filter ex 00*"
	// }
	jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStr)
	// Call the host function
	offset := hostRedisFilter(arguments.Offset())

	// Get result (the value associated to the key) from shared memory
	// The host function (hostRedisDel) returns a JSON buffer:
	// {
	//   "success": "array of keys",
	//   "failure": "error message if error, else empty"
	// }
	memoryValue := pdk.FindMemory(offset)
	buffer := make([]byte, memoryValue.Length())
	memoryValue.Load(buffer)

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

//export hello
func hello() uint64 {
	redisURI := GetEnv("REDIS_URI")
	idRedisClient, errInit := InitRedisClient("redisDb", redisURI)
	if errInit!= nil {
		Print("😡 " + errInit.Error())
	} else {
		Print("🙂 " + idRedisClient)
	}

	k1, errSet1 := RedisSet("redisDb", "001", "Huey")
	k2, errSet2 := RedisSet("redisDb", "002", "Dewey")
	k3, errSet3 := RedisSet("redisDb", "003", "Louie")

	allSetErrs := errors.Join(errSet1, errSet2, errSet3) 
	if allSetErrs != nil {
		Print("😡 " + allSetErrs.Error())
	} else {
		Print("🙂 " + strings.Join([]string{k1,k2,k3}, ","))
	}

	v1, errGet1 := RedisGet("redisDb", "001")
	v2, errGet2 := RedisGet("redisDb", "002")
	v3, errGet3 := RedisGet("redisDb", "003")

	allGetErrs := errors.Join(errGet1, errGet2, errGet3) 
	if allGetErrs != nil {
		Print("😡 " + allSetErrs.Error())
	} else {
		Print("🙂 " + strings.Join([]string{v1,v2,v3}, ","))
	}

	key, errDel := RedisDel("redisDb", "002")
	if errDel != nil {
		Print("😡 " + errDel.Error())
	} else {
		Print("🙂 " + key)
	}

	keys, errKeys := RedisFilter("redisDb", "00*")
	if errKeys != nil {
		Print("😡 " + errKeys.Error())
	} else {
		Print("🙂 " + keys)
	}

	/* output:
		🙂 001,002,003
		🙂 Huey,Dewey,Louie
		🙂 002
		🙂 ["003","001"]
	*/

	return 0
}

func main() {}
