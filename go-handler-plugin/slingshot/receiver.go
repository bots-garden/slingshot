package slingshot

import (
	"github.com/extism/go-pdk"
)

// Execute the function
func CallHandler(function func(param []byte) []byte)  {

	functionParameters := pdk.Input()

	//Print("🎃 " + string(functionParameters))

	value := function(functionParameters)

	mem := pdk.AllocateBytes(value)
	// copy output to host memory
	pdk.OutputMemory(mem)
	//return 0

}
