package main

import (
	"errors"
	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostPrint
func hostPrint(offset uint64) uint64

func Print(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrint(memoryText.Offset())
}

var parser = fastjson.Parser{}

//export hostMemorySet
func hostMemorySet(offset uint64) uint64

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

//export hostMemoryGet
func hostMemoryGet(offset uint64) uint64

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

	_, err := MemorySet("bob", "Bob Morane")

	value, err := MemoryGet("bob")
	if err != nil {
		Print("😡 ouch! " + err.Error())
	} else {
		Print("value: " + value)
	}
	
	return 0
}

func main() {}
