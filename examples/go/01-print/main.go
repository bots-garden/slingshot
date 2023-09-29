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

//export hostLog
func hostLog(offset uint64) uint64

func Log(text string) {
	memoryText := pdk.AllocateString(text)
	hostLog(memoryText.Offset())
}


//export hello
func hello() uint64 {
	// read function argument from the memory
	input := pdk.Input()

	Println("👋 hello world 🌍 " + string(input))
	Log("🙂 have a nice day 🏖️")

	return 0
}

func main() {}
