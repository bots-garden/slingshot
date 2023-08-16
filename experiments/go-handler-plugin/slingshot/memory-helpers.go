package slingshot

import (
	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

var parser = fastjson.Parser{}

func ReadBufferFromMemory(offset uint64) []byte {
	// read the value into the memory
	// offset is the position and the length of the result
	// (2 values into only one value)
	// get the length and the position of the result in memory
	memoryResult := pdk.FindMemory(offset)
	// create a buffer from memoryResult
	// fill the buffer with memoryResult
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	return buffResult
}

func ReadStringFromMemory(offset uint64) string {
	return string(ReadBufferFromMemory(offset))
}

func ReadJsonFromMemory(offset uint64) (*fastjson.Value, error) {
	buffer := ReadBufferFromMemory(offset)
	JSONData, err := parser.ParseBytes(buffer)
	/*
		if err != nil {
			return nil, err
		}
	*/
	return JSONData, err
}

func CopyBufferToMemory(value []byte) pdk.Memory {
	// Copy the buffer value to the shared memory
	memoryValue := pdk.AllocateBytes(value)
	return memoryValue
}

func CopyStringToMemory(value string) pdk.Memory {
	// Copy the string value to the shared memory
	memoryValue := pdk.AllocateString(value)
	return memoryValue
}

func CopyJsonToMemory() {

}
