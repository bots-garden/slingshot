package slingshot

import (
	"github.com/extism/extism"
	"github.com/tetratelabs/wazero/api"
)

var hostFunctions = []extism.HostFunction{}

// Add a new host function
func AppendHostFunction(hostFunction extism.HostFunction) {
	hostFunctions = append(hostFunctions, hostFunction)
}

// Retrieve the list of the defined host functions
func GetHostFunctions() []extism.HostFunction {
	return hostFunctions
}

// Define a Host Function CallBack
func DefineHostFunctionCallBack(wasmName string, callBack extism.HostFunctionCallback) extism.HostFunction {
	host_function := extism.HostFunction{
		Name:      wasmName,
		Namespace: "env",
		Callback:  callBack,
		Params:    []api.ValueType{api.ValueTypeI64},
		Results:   []api.ValueType{api.ValueTypeI64},
	}
	return host_function

}
