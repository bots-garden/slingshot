package main

import (
	"errors"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostPrint
func hostPrint(offset uint64) uint64

func Print(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrint(memoryText.Offset())
}

var parser = fastjson.Parser{}

// GetJsonFromBytes
// Convert a buffer (`[]byte`) into a JSON value
func GetJsonFromBytes(buffer []byte) (*fastjson.Value, error) {
	return parser.ParseBytes(buffer)
}

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

//export hello
func hello() uint64 {

	content, err := ReadFile("./hello.txt")
	if err != nil {
		Print("😡 " + err.Error())
	} else {
		Print(content)
	}

	return 0
}

func main() {}
