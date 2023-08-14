package callbacks

import (
	"context"
	"fmt"
	"log"
	"slingshot-server/mem"
	"slingshot-server/slingshot"
	"sync"

	"github.com/extism/extism"
)

var memCache sync.Map

type memoryRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func MemorySet(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {
	/* Expected
	{ key: "", value: "" }
	*/
	var result = slingshot.StringResult{}
	var record memoryRecord

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &record)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		// Remove this after the tests
		fmt.Println("🟡 MemorySet from host:", record.Key, record.Value)
		// Read the memCache map
		memCache.Store(record.Key, record.Value)
		result.Failure = ""
		result.Success = record.Value
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}

}

func MemoryGet(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

	var result = slingshot.StringResult{}
	// The expected argument is a key (string)
	keyFromWasmModule, err := mem.ReadStringFromMemory(plugin, stack)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {
		keyStr := string(keyFromWasmModule)
		fmt.Println("🟡 MemoryGet from host:", keyStr)
		value, ok := memCache.Load(keyStr)

		if ok {
			result.Failure = ""
			result.Success = value.(string)
		} else {
			result.Failure = "Not found"
			result.Success = ""
		}
	}

	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 MemorySet, CopyJsonToMemory:", err)
	}

}
