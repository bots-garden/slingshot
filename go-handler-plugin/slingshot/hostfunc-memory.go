package slingshot

import (
	"errors"
)

//export hostMemorySet
func hostMemorySet(offset uint64) uint64

func MemorySet(key string, value string) {
	// call host function with json argument
	jsonStr := `{"key":"` + key + `","value":"` + value + `"}`
	memoryJsonStr := CopyStringToMemory(jsonStr)

	offset := hostMemorySet(memoryJsonStr.Offset())

	// get result from shared memory
	buffResult := ReadBufferFromMemory(offset)
	// this is a test
	Print("🟣 MemorySet from guest: " + string(buffResult))
}

//export hostMemoryGet
func hostMemoryGet(offset uint64) uint64

func MemoryGet(key string) (string, error) {
	// Copy argument to memory
	memoryKeyStr := CopyStringToMemory(key)
	// Call the host function
	offset := hostMemoryGet(memoryKeyStr.Offset())

	// Get result from shared memory
	// it will be a JSON object
	JSONData, err := ReadJsonFromMemory(offset)
	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}
}

