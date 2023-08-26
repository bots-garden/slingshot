package slingshot

import (
	"github.com/extism/go-pdk"
)

var handlerFunction func([]byte) ([]byte, error)

func SetHandler(function func([]byte) ([]byte, error)) {
	handlerFunction = function
}

//export callHandler
func callHandler() {
	functionParameters := pdk.Input()

	value, err := handlerFunction(functionParameters)
	//TODO return error?
	if err != nil {
		Print("😡👋 we have something to do here: " + err.Error())
	}
	mem := pdk.AllocateBytes(value)
	// copy output to host memory
	pdk.OutputMemory(mem)
	//return 0
}
