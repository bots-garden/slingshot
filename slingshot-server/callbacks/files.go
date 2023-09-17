package callbacks

import (
	"context"
	"log"
	"os"
	"slingshot-server/mem"
	"slingshot-server/slingshot"

	"github.com/extism/extism"
)

func ReadFile(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	var result = slingshot.StringResult{}

	// Read data from the shared memory
	filePath, errArg := mem.ReadStringFromMemory(plugin, stack)
	if errArg != nil {
		result.Failure = errArg.Error()
		result.Success = ""
	} else {
		// Construct the result
		data, errReadFile := os.ReadFile(string(filePath))
		if errReadFile != nil {
			result.Failure = errReadFile.Error()
			result.Success = ""
		}
		result.Failure = ""
		result.Success = string(data)
	}

	// Copy the result to the memory (== return value)
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 RedisDel, CopyJsonToMemory:", errResult.Error())
	}

}

func WriteFile(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	
}
