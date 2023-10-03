package slingshot

import "github.com/extism/go-pdk"

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
