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

//export message
func message() uint64 {
	// read function argument from the memory
	input := pdk.Input()

	Println("👋 message: " + string(input))
	
	return 0
}

func main() {}
