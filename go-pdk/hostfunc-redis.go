package slingshot

import (
	"errors"

	"github.com/extism/go-pdk"
)

// hostInitRedisClient:
/*
  This exported function will call the host function callback `InitRedisClient` of the slingshot application
*/
//export hostInitRedisClient
func hostInitRedisClient(offset uint64) uint64

// InitRedisClient: initialize a Redis client
/*
	- This helper call the `hostInitRedisClient` exported function
	- It copies the `redisClientId` and `redisUri` parameters to the shared memory
	- It calls the `InitRedisClient` host function callback (when calling `hostInitRedisClient`)
	- It reads the result of the callback
	- And returns this result

  This function will create a Redis client and store it with an id (`redisClientId`),
  that means you can retrieve the connection by id (for example with the `RedisSet` function)
*/
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

// hostRedisSet:
/*
  This exported function will call the host function callback `RedisSet` of the slingshot application
*/
//export hostRedisSet
func hostRedisSet(offset uint64) uint64

// RedisSet: set/store a value in a Redis db
/*
	- This helper call the `hostRedisSet` exported function
	- It copies the `redisClientId`, `key` and `value` parameters to the shared memory
	- It calls the `RedisSet` host function callback (when calling `hostRedisSet`)
	- It reads the result of the callback
	- And returns this result
*/
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

// hostRedisGet:
/*
  This exported function will call the host function callback `RedisGet` of the slingshot application
*/
//export hostRedisGet
func hostRedisGet(offset uint64) uint64

// RedisGet: get a value from a Redis db
/*
	- This helper call the `hostRedisGet` exported function
	- It copies the `redisClientId` and `key` parameters to the shared memory
	- It calls the `RedisGet` host function callback (when calling `hostRedisGet`)
	- It reads the result of the callback
	- And returns this result
*/
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

// hostRedisDel:
/*
  This exported function will call the host function callback `RedisDel` of the slingshot application
*/
//export hostRedisDel
func hostRedisDel(offset uint64) uint64

// RedisDel: delete a value from a Redis db
/*
	- This helper call the `hostRedisDel` exported function
	- It copies the `redisClientId` and `key` parameters to the shared memory
	- It calls the `RedisDel` host function callback (when calling `hostRedisDel`)
	- It reads the result of the callback
	- And returns this result
*/
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

// hostRedisFilter:
/*
  This exported function will call the host function callback `RedisFilter` of the slingshot application
*/
//export hostRedisFilter
func hostRedisFilter(offset uint64) uint64

// RedisFilter: retrieve a list of keys using a filter
/*
	- This helper call the `hostRedisFilter` exported function
	- It copies the `redisClientId` and `key` parameters to the shared memory
	- It calls the `RedisFilter` host function callback (when calling `hostRedisFilter`)
	- It reads the result of the callback
	- And returns this result
*/
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
