package main

import (
	"github.com/extism/go-pdk"
)

//export hello
func hello() {
	// read function argument from the shared memory
	input := pdk.Input()
	output := "👋 Hello " + string(input)

	// copy output to shared memory
	mem := pdk.AllocateString(output)
	pdk.OutputMemory(mem)
}

func main() {}
