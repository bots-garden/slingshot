package main

import (
	"encoding/base64"
	"errors"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export hostPrintln
func hostPrintln(offset uint64) uint64

func Println(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrintln(memoryText.Offset())
}

var parser = fastjson.Parser{}

// GetJsonFromBytes
// Convert a buffer (`[]byte`) into a JSON value
func GetJsonFromBytes(buffer []byte) (*fastjson.Value, error) {
	return parser.ParseBytes(buffer)
}

//export hostWriteFile
func hostWriteFile(offset uint64) uint64

func WriteFile(filePath string, contentFile string) error {

	content := base64.StdEncoding.EncodeToString([]byte(contentFile))

	jsonStrArguments := `{"path":"` + filePath + `","content":"` + content + `"}`

	// Copy the string value to the shared memory
	arguments := pdk.AllocateString(jsonStrArguments)

	// Call the host function with Json string argument
	offset := hostWriteFile(arguments.Offset())

	// Get result from the shared memory
	memoryResult := pdk.FindMemory(offset)
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)
	JSONData, err := GetJsonFromBytes(buffResult)

	if err != nil {
		return err
	}
	if len(JSONData.GetStringBytes("failure")) == 0 {
		return nil
	} else {
		return errors.New(string(JSONData.GetStringBytes("failure")))
	}
}

//export hello
func hello() uint64 {

	text := `
	<html>
	  <h1>"Hello World!!!"</h1>
	</html>
	`

	err := WriteFile("./index.html", text)
	if err != nil {
		Println("😡 " + err.Error())
	}

	return 0
}

func main() {}
