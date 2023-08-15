package main

import "github.com/extism/go-pdk"

//export hello
func hello() uint64 {
	// read function argument from the memory
	input := pdk.Input()
	output := "Hello " + string(input)
	mem := pdk.AllocateString(output)
	// copy output to host memory
	pdk.OutputMemory(mem)
	return 0
}

func main() {}
