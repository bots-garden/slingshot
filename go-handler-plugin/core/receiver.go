package receiver

import (
	"github.com/extism/go-pdk"
)

// Execute the function
func SetHandler(function func(param []byte) ([]byte, error))  {

	functionParameters := pdk.Input()

	value, err := function(functionParameters)

	stringValue := string(value)

	var stringError string
	if err != nil {
		stringError = err.Error()
	} else {
		stringError = ""
	}

	jsonResult := `{"success":"`+stringValue+`", failure:"`+stringError+`"}`

	mem := pdk.AllocateString(jsonResult)
	// copy output to host memory
	pdk.OutputMemory(mem)
	//return 0

}
