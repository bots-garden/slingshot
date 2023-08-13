package callbacks

// This one is for experiments and tests
import (
	"context"
	"fmt"
	"os"
	"slingshot-server/slingshot"

	"github.com/extism/extism"
)

func GetEnv(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

	// Read data from the shared memory
	variableName, errArg := slingshot.ReadStringFromMemory(plugin, stack)
	if errArg != nil {
		fmt.Println("🔴", errArg.Error())
		panic(errArg)
	}
	// Construct the result
	variableValue := os.Getenv(variableName)

	errRet := slingshot.CopyStringToMemory(plugin, stack, variableValue)
	if errRet != nil {
		fmt.Println("🔴", errRet.Error())
		panic(errRet)
	}
}
