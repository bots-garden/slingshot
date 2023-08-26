package slingshot

import (
	"github.com/extism/go-pdk"
)

var handlerFunction func([]byte) []byte

// SetHandler:
/*
  - Define the handler function that will be called/triggered.
  - This handler will be called by the exported function named `callHandler`.
  - ✋ this function is used in the `main` function.

*/
func SetHandler(function func([]byte) []byte) {
	handlerFunction = function
}

// callHandler:
/*
  This exported function will call the handler defined with `SetHandler`
*/
//export callHandler
func callHandler() {
	functionParameters := pdk.Input()

	value := handlerFunction(functionParameters)

	mem := pdk.AllocateBytes(value)
	// copy output to host memory
	pdk.OutputMemory(mem)
	//return 0
}
