package main

import (
	"strings"
	slingshot "github.com/bots-garden/slingshot/go-pdk"
	"github.com/valyala/fastjson"
)

var parser = fastjson.Parser{}

func Handler(http_request_data []byte) []byte {

	var text string
	var code string

	/*
		TODO: getHTTPResponse(http_request_data []byte)
		http_request_data.(HTTPResponse)
	*/

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

	response := `{"headers":{` + headersStr + `}, "textBody": "` + text + `", "statusCode": ` + code + `}`

	return []byte(response)

}

func main() {
	slingshot.Print("👋 main function")
	slingshot.SetHandler(Handler)
}

/* with the slingshot pdk, always call `callHandler`

    ./slingshot listen --wasm=./hello.wasm \
	--handler=callHandler \
	--http-port=7070

*/

/* TODO(?)
	1- make an helper with json (something like getHttpRequest)
	2- make an HTTP Handler
	3- make an helper to return something like httpResponse
*/