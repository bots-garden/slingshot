package slingshot

import "github.com/extism/go-pdk"

// hostGetEnv:
/*
  This exported function will call the host function callback `GetEnv` of the slingshot application
*/
//export hostGetEnv
func hostGetEnv(offset uint64) uint64

// GetEnv: get the value of an environment variable
/*
	- This helper call the `hostGetEnv` exported function
	- It copies the `name` parameter to the shared memory
	- It calls the `GetEnv` host function callback (when calling `hostGetEnv`)
	- It reads the result of the callback
	- And returns this result
*/
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
