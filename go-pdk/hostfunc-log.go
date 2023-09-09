package slingshot

import "github.com/extism/go-pdk"

// hostLog:
/*
  This exported function will call the host function callback `Log` of the slingshot application
*/
//export hostLog
func hostLog(offset uint64) uint64

// Log: print `text` (as a log entry)
/*
	- This helper call the `hostLog` exported function
	- It copies the `text` parameter to the shared memory
	- It calls the `Log` host function callback (when calling `hostLog`)
	- It will print (with a log format) the `text` argument
*/
func Log(text string) {
	memoryText := pdk.AllocateString(text)
	hostLog(memoryText.Offset())
}
