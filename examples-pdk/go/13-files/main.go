package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func helloHandler(argHandler []byte) []byte {
	input := string(argHandler)
	slingshot.Println("👋 hello world 🌍 " + string(input))
	
	slingshot.Log("🙂 have a nice day 🏖️")

	content, err := slingshot.ReadFile("./hello.txt")
	if err != nil {
		slingshot.Log("😡 " + err.Error())
	}
	slingshot.Println(content)

	text := `
	<html>
	  <h1>"Hello World!!!"</h1>
	</html>
	`

	errWrite := slingshot.WriteFile("./index.html", text)
	if errWrite != nil {
		slingshot.Log("😡 " + errWrite.Error())
	}

	return []byte("👋 Hello World 🌍")
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(helloHandler)
}

func main() {}
