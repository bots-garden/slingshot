package main

import (
	"go-handler-plugin/slingshot"
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

	// Get the arguments passed by the host
	slingshot.CallHandler(func(param []byte) ([]byte, error) {
		res := `{"message":"👋 Hello ` + string(param) + `", "number":42, "message":"` + val1 + " - " + val2 + `"}`
		return []byte(res), nil
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
