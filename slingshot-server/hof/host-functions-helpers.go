package hof

import (
	extism "github.com/extism/go-sdk"
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
func DefineHostFunctionCallBack(wasmName string, callBack extism.HostFunctionStackCallback) extism.HostFunction {
	
	host_function := extism.NewHostFunctionWithStack(
		wasmName,
		"env",
		callBack,
		[]api.ValueType{api.ValueTypeI64},
		[]api.ValueType{api.ValueTypeI64},
	)

	return host_function

}
