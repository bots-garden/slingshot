package main

import (
	receiver "go-handler-plugin/core"
	"github.com/extism/go-pdk"
)

//go:wasm-module env
//export hostMemoryGet
func hostMemoryGet(offset uint64) uint64

func MemoryGet(key string) string {

	// Call the host function
	// 1- copy the key to the shared memory
	memoryKey := pdk.AllocateString(key)
	// call the host function
	// memoryKey.Offset() is the position and the length of memoryKey into the memory (2 values into only one value)
	offset := hostMemoryGet(memoryKey.Offset())
	// read the value into the memory
	// offset is the position and the length of the result (2 values into only one value)
	// get the length and the position of the result in memory
	memoryResult := pdk.FindMemory(offset)
	/*
		memoryResult is a struct instance
		type Memory struct {
			offset uint64
			length uint64
		}
	*/	
	// create a buffer from memoryResult
	// fill the buffer with memoryResult
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	return string(buffResult)

}

//export handle
func handle() {

	val1 := MemoryGet("hello")
	val2 := MemoryGet("message")
	
	receiver.SetHandler(func(param []byte) ([]byte, error) {
		res := `{"message":"👋 Hello `+ string(param) + `", "number":42, "message":"`+ val1 + " - " + val2 +`"}`
		return []byte(res), nil
	})
}

func main() {

}

// 👋 see this example:
// https://github.com/GoogleCloudPlatform/go-templates/blob/main/functions/httpfn/httpfn.go

/*
func init() {
	functions.HTTP("HelloHTTP", helloHTTP)
}

// helloHTTP is an HTTP Cloud Function.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
*/