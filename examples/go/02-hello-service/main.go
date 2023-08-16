package main

import (
	"strings"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

var parser = fastjson.Parser{}

//export handle
func handle()  {
	// read function argument from the memory
	http_request_data := pdk.Input()

	var text string
	var code string

	JSONData, err := parser.ParseBytes(http_request_data)
	if err != nil {
		text = "😡 Error: " + err.Error()
		code = "500"
	} else {
		text = "🩵 Hello " + string(JSONData.GetStringBytes("body"))
		code = "200"
	}

	headers := []string{
		`"Content-Type": "application/json; charset=utf-8"`,
		`"X-Slingshot-version": "0.0.0"`,
	}

	headersStr := strings.Join(headers, ",")

	response := `{"headers":{` + headersStr + `}, "textBody": "` + text + `", "statusCode": `+ code +`}`

	mem := pdk.AllocateString(response)
	// copy output to host memory
	pdk.OutputMemory(mem)	
}

func main() {}

