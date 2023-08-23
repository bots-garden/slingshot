package slingshot

import (
	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

/* TODO:
  - ReadHTTPRequest (?)

*/


var parser = fastjson.Parser{}

/*
ReadBufferFromMemory: (host -> guest)
	- read the value from the shared memory
	- `offset` arg represents the position and the length of the result (2 values into only one value)
	- return a buffer with the content (of the piece of memory)
*/
func ReadBufferFromMemory(offset uint64) []byte {

	// get the length and the position of the result in memory
	memoryResult := pdk.FindMemory(offset)
	
	// create a buffer from memoryResult
	// fill the buffer with memoryResult
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	return buffResult
}

/*
ReadStringFromMemory: (host -> guest)
	- read the value from the shared memory
	- `offset` arg represents the position and the length of the result (2 values into only one value)
	- return a string with the content (of the piece of memory)
*/
func ReadStringFromMemory(offset uint64) string {
	return string(ReadBufferFromMemory(offset))
}

/*
ReadJsonFromMemory: (host -> guest)
	- read the value from the shared memory
	- `offset` arg represents the position and the length of the result (2 values into only one value)
	- return a *fastjson.Value & error
*/
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

 
/*
CopyBufferToMemory: (guest -> host)
	Copy the buffer value to the shared memory
	This function does exactly the same thing as `pdk.AllocateBytes`
*/
func CopyBufferToMemory(value []byte) pdk.Memory {
	memoryValue := pdk.AllocateBytes(value)
	return memoryValue
}

/*
CopyStringToMemory: (guest -> host)
	Copy the string value to the shared memory
	This function does exactly the same thing as `pdk.AllocateString`

*/
func CopyStringToMemory(value string) pdk.Memory {
	// Copy the string value to the shared memory
	memoryValue := pdk.AllocateString(value)
	return memoryValue
}

/*
CopyJsonToMemory: (guest -> host)
	Copy the JSON value to the shared memory
	👋 I'm not sure if it's useful yet. 🤔
	I need to test it
*/
func CopyJsonToMemory(JSONData *fastjson.Value) pdk.Memory {
	var b []byte
	JSONData.MarshalTo(b)
	memoryValue := CopyBufferToMemory(b)
	return memoryValue
}
