package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func helloHandler(argHandler []byte) {

	name := slingshot.Input("🤖 What's your name? > ")

	slingshot.Println("👋 Hello " + name + " 😄")
	
}

//export callHandler
func callHandler() {
	slingshot.ExecVoidHandler(helloHandler)
}

func main() {}
