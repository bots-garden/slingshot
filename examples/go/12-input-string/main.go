package main

import (
	"github.com/extism/go-pdk"
)

//export hostPrintln
func hostPrintln(offset uint64) uint64

func Println(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrintln(memoryText.Offset())
}

//export hostInput
func hostInput(offset uint64) uint64

func Input(prompt string) string {
	promptMemory := pdk.AllocateString(prompt)
	offset := hostInput(promptMemory.Offset())

	memoryResult := pdk.FindMemory(offset)
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	return string(buffResult)
}

//export hello
func hello() {
	name := Input("🤖 Name? > ")
	Println("👋 Hello " + name + " 😃")

}

func main() {}
