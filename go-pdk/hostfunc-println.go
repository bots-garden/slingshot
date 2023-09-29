package slingshot

import "github.com/extism/go-pdk"

// hostPrint:
/*
  This exported function will call the host function callback `Print` of the slingshot application
*/
//export hostPrintln
func hostPrintln(offset uint64) uint64

// Println: print `text`
/*
	- This helper call the `hostPrint` exported function
	- It copies the `text` parameter to the shared memory
	- It calls the `Print` host function callback (when calling `hostPrint`)
	- It will print the `text` argument
*/
func Println(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrintln(memoryText.Offset())
}

