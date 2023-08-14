package main

import (
	"go-handler-plugin/slingshot"
	"strings"

	"github.com/valyala/fastjson"
)

//export handle
func handle() {
	slingshot.Print("🟣 this is the handle() function")

	val1 := slingshot.GetMessage("hello")
	val2 := slingshot.GetMessage("message")

	slingshot.MemorySet("Bob", "Morane")

	name, err := slingshot.MemoryGet("Bob")
	if err != nil {
		slingshot.Print("🔴 " + err.Error())
	} else {
		slingshot.Print("🙂 " + name)
		slingshot.Log("😈" + name)
	}

	var parser = fastjson.Parser{}

	// Get the arguments passed by the host
	slingshot.CallHandler(func(param []byte) []byte {

		slingshot.Print("🎃🎃 " + string(param))
		//TODO getResponseFrom()

		JSONData, err := parser.ParseBytes(param)
		if err != nil {}
		body := JSONData.GetStringBytes("body")

		data := `{"salutation":"👋 Hello ` + string(body) + `", "number":42, "message":"` + val1 + " - " + val2 + `"}`

		headers := []string{
			`"Content-Type": "application/json; charset=utf-8"`,
			`"X-Slingshot-version": "0.0.0"`,
		}
		response := `{"headers":{` + strings.Join(headers, ",") + `}, "jsonBody": ` + data + `, "statusCode": 200}`

		return []byte(response)
	})
}

func init() {
	slingshot.Print("🟠 this is the init() function")
}

func main() {
	slingshot.Print("🔵 this is the main() function")
}

// 👋 see this example:
// https://github.com/GoogleCloudPlatform/go-templates/blob/main/functions/httpfn/httpfn.go

/*
func init() {
	functions.HTTP("HelloHTTP", helloHTTP)
}

// helloHTTP is an HTTP Cloud Function.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
*/
