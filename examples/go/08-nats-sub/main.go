package main

import (
	"github.com/extism/go-pdk"
)

//export hostPrint
func hostPrint(offset uint64) uint64

func Print(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrint(memoryText.Offset())
}

//export message
func message() uint64 {
	// read function argument from the memory
	input := pdk.Input()

	Print("👋 message: " + string(input))
	
	return 0
}

func main() {}
