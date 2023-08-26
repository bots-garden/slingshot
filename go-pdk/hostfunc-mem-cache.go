package slingshot

import (
	"errors"

	"github.com/extism/go-pdk"
)

// hostMemorySet:
/*
  This exported function will call the host function callback `MemorySet` of the slingshot application
*/
//export hostMemorySet
func hostMemorySet(offset uint64) uint64

// MemorySet: set/store a value in a memory map (handled by the host application)
/*
	- This helper call the `hostMemorySet` exported function
	- It copies the `key` and `value` parameters to the shared memory
	- It calls the `MemorySet` host function callback (when calling `hostMemorySet`)
	- It reads the result of the callback
	- And returns this result
*/
func MemorySet(key string, value string) (string, error) {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "key": "name",
	//    "value": "Bob Morane"
	// }
	jsonStr := `{"key":"` + key + `","value":"` + value + `"}`

	// Copy the string value to the shared memory
	keyAndValue := pdk.AllocateString(jsonStr)

	// Call host function with the offset of the arguments
	offset := hostMemorySet(keyAndValue.Offset())

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

// hostMemoryGet:
/*
  This exported function will call the host function callback `MemoryGet` of the slingshot application
*/
//export hostMemoryGet
func hostMemoryGet(offset uint64) uint64

// MemoryGet: get a value from a memory map (handled by the host application)
/*
	- This helper call the `hostMemoryGet` exported function
	- It copies the `key` parameter to the shared memory
	- It calls the `MemoryGet` host function callback (when calling `hostMemoryGet`)
	- It reads the result of the callback
	- And returns this result
*/
func MemoryGet(key string) (string, error) {
	// Copy argument to memory
	memoryKey := pdk.AllocateString(key)
	// Call the host function
	offset := hostMemoryGet(memoryKey.Offset())

	// Get result (the value associated to the key) from shared memory
	// The host function (hostMemoryGet) returns a JSON buffer:
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
