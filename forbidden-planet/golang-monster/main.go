package main

import (
	"github.com/extism/go-pdk"
)

//export hostSendMessage
func hostSendMessage(offset uint64) uint64

func sendMessageToHost(message string, action string) string {
	//"👋 hello world 🌍"

	data := `{"avatar":"🥶", "name":"GoMonster", "x":1, "y":2, "message":"`+message+`", "action":"`+action+`"}`

	dataMemory := pdk.AllocateString(data)
	offsetResult := hostSendMessage(dataMemory.Offset())

	resultMemory := pdk.FindMemory(offsetResult)

	resultBuffer := make([]byte, resultMemory.Length())
	resultMemory.Load(resultBuffer)
	
	return string(resultBuffer)
}

//export hostDisplay
func hostDisplay(offset uint64) uint64

func display(text string) {
	dataMemory := pdk.AllocateString(text)
	hostDisplay(dataMemory.Offset())
	//offsetResult := hostSendMessage(dataMemory.Offset())

}



//export hey
func hey() int32 {
	response := sendMessageToHost("👋 hello world 🌍", "toctoc")
	display("display called by guest: " + response)

	//display("PHILIPPE CHARRIERE")

	return 0
}

//export hello
func hello() int32 {
	

	// read function argument from the memory
	input := pdk.Input()

	output := "👋 Hello " + string(input)

	mem := pdk.AllocateString(output)
	// copy output to host memory
	pdk.OutputMemory(mem)

	return 0
}


//export getName
func getName() int32 {
	name := "GoMonster"
	pdk.OutputMemory(pdk.AllocateString(name))
	return 0
}

//export getAvatar
func getAvatar() int32 {
	avatar := "🥶"
	pdk.OutputMemory(pdk.AllocateString(avatar))
	return 0
}

func main() {}
