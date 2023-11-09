package callbacks

// This one is for experiments and tests
import (
	"context"
	"fmt"
	"os"
	"slingshot-server/mem"

	extism "github.com/extism/go-sdk"

)

func GetEnv(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {

	// Read data from the shared memory
	variableName, errArg := mem.ReadStringFromMemory(plugin, stack)
	if errArg != nil {
		fmt.Println("🔴", errArg.Error())
		panic(errArg)
	}
	// Construct the result
	variableValue := os.Getenv(variableName)

	errRet := mem.CopyStringToMemory(plugin, stack, variableValue)
	if errRet != nil {
		fmt.Println("🔴", errRet.Error())
		panic(errRet)
	}
}
// TODO: return error (failure) instead of panic
// TODO: 👀 files.go
