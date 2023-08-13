package main

import (
	"github.com/extism/go-pdk"
)

//export hostGetEnv
func hostGetEnv(offset uint64) uint64

//export read_env_var
func read_env_var() int32 {

	// Call the host function
	// Copy the name of the environment variable to the shared memory
	varNameMemory := pdk.AllocateString("SLINGSHOT_MESSAGE")
	// Call the host function
	// varNameMemory.Offset() is the position and the length of "SLINGSHOT_MESSAGE" into the memory
	// (2 values into only one value)
	offset := hostGetEnv(varNameMemory.Offset())
	// Read the value into the memory
	// offset is the position and the length of the result (the value of the environment variable)
	// (2 values into only one value)
	// get the length and the position of the result in memory
	varValueMemory := pdk.FindMemory(offset)
	/*
		varValueMemory is a struct instance
		type Memory struct {
			offset uint64
			length uint64
		}
	*/

	// create a buffer from the varValueMemory
	// fill the buffer with varValueMemory
	buffer := make([]byte, varValueMemory.Length())
	varValueMemory.Load(buffer)

	envVarValue := string(buffer)

	if envVarValue == "" {
		// Allocate space into the memory
		mem := pdk.AllocateString("EMPTY")
		// copy output to host memory
		pdk.OutputMemory(mem)
	} else {
		// Allocate space into the memory
		mem := pdk.AllocateString(envVarValue)
		// copy output to host memory
		pdk.OutputMemory(mem)
	}

	return 0
}

func main() {}
