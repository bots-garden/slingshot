package main

import (
	"github.com/extism/go-pdk"
)

//export hostPrint
func hostPrint(offset uint64) uint64

func Print(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrint(memoryText.Offset())
}

//export hostGetEnv
func hostGetEnv(offset uint64) uint64

func GetEnv(name string) string {
	// copy the name of the environment variable to the shared memory
	variableName := pdk.AllocateString(name)
	// call the host function
	offset := hostGetEnv(variableName.Offset())

	// read the value of the result from the shared memory
	variableValue := pdk.FindMemory(offset)
	buffer := make([]byte, variableValue.Length())
	variableValue.Load(buffer)

	// cast the buffer to string and return the value
	envVarValue := string(buffer)
	return envVarValue
}

//export hello
func hello() uint64 {
	message := GetEnv("MESSAGE")
	Print("🤖 MESSAGE=" + message)

	return 0
}

func main() {}
