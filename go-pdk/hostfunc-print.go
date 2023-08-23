package slingshot

import "github.com/extism/go-pdk"

//export hostPrint
func hostPrint(offset uint64) uint64

/* Print 
	calls the hostPrint callback
*/
func Print(text string) {
	memoryText := pdk.AllocateString(text)
	hostPrint(memoryText.Offset())
}
