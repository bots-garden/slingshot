package slingshot

import (
	"errors"

	"github.com/extism/go-pdk"
)

//export hostReadFile
func hostReadFile(offset uint64) uint64

func ReadFile(filePath string) (string, error) {
	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(filePath)

	// Call host function with the offset of the arguments
	offset := hostReadFile(arguments.Offset())

	// Get result from the shared memory
	memoryResult := pdk.FindMemory(offset)
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)
	JSONData, err := GetJsonFromBytes(buffResult)
	if err != nil {
		return "", err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return string(JSONData.GetStringBytes("success")), nil
	} else {
		return "", errors.New(string(JSONData.GetStringBytes("failure")))
	}

}
