package callbacks

import (
	"context"
	"fmt"

	"github.com/extism/extism"
)

var memoryMap = map[string]string{
	"hello":   "🖖 Hello World 🌍",
	"message": "I 💜 Extism 😍",
}

func MemoryGet(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	offset := stack[0]
	bufferInput, err := plugin.ReadBytes(offset)

	if err != nil {
		fmt.Println("🥵", err.Error())
		panic(err)
	}

	keyStr := string(bufferInput)
	fmt.Println("🟢🟠 keyStr:", keyStr) // this is for test

	returnValue := memoryMap[keyStr]

	plugin.Free(offset)
	offset, err = plugin.WriteBytes([]byte(returnValue))
	if err != nil {
		fmt.Println("😡", err.Error())
		panic(err)
	}

	stack[0] = offset
}
