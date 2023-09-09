package slingshot

import (
	"github.com/extism/go-pdk"
)

func ExecHandler(handlerFunction func(param []byte) []byte) {
	functionParameters := pdk.Input()

	value := handlerFunction(functionParameters)

	mem := pdk.AllocateBytes(value)
	// copy output to host memory
	pdk.OutputMemory(mem)
	//return 0
}
