package callbacks

// This one is for experiments and tests
import (
	"context"
	"fmt"
	"slingshot-server/mem"

	"github.com/extism/extism"
)

var messagesMap = map[string]string{
	"hello":   "👋 Hello World 🌍",
	"message": "I 💜 Extism 😍",
	"vulcan":  "🖖 peace and long life",
}

func GetMessage(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

	// Read data from the shared memory
	keyStr, errArg := mem.ReadStringFromMemory(plugin, stack)
	if errArg != nil {
		fmt.Println("😡", errArg.Error())
		panic(errArg)
	}
	fmt.Println("🟡 GetMessage from host, keyStr:", keyStr) // this is for test

	// Construct the result
	returnValue := messagesMap[keyStr]

	errRet := mem.CopyStringToMemory(plugin, stack, returnValue)
	if errRet != nil {
		fmt.Println("😡", errRet.Error())
		panic(errRet)
	}
}
