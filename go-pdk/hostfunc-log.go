package slingshot

import "github.com/extism/go-pdk"

//export hostLog
func hostLog(offset uint64) uint64

/* Log 
	calls the hostLog callback
*/
func Log(text string) {
	memoryText := pdk.AllocateString(text)
	hostLog(memoryText.Offset())
}
