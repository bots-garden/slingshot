package callbacks

import (
	"context"
	"fmt"

	"github.com/extism/extism"
)

func Print(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	offset := stack[0]
	bufferInput, err := plugin.ReadBytes(offset)

	if err != nil {
		fmt.Println("🥵", err.Error())
		panic(err)
	}

	stringToDisplay := string(bufferInput)
	fmt.Println(stringToDisplay)

	plugin.Free(offset)

	//stack[0] = 0
}
