package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
	"github.com/extism/go-pdk"
)

//export hello
func hello() uint64 {
	// read function argument from the memory
	input := pdk.Input()
	
	slingshot.Print("👋 hello world 🌍 " + string(input))
	slingshot.Log("🙂 have a nice day 🏖️")

	return 0
}

func main() {}
