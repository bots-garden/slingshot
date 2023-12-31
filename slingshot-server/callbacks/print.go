package callbacks

import (
	"context"
	"fmt"
	"log"
	"slingshot-server/mem"
	"time"

	extism "github.com/extism/go-sdk"

)

func Print(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {

	// The WASM plugin copied data into the shared memory
	// Red this data
	bufferToPrint, err := mem.ReadBytesFromMemory(plugin, stack)

	if err != nil {
		log.Println("🔴 Print, ReadBytesFromMemory", err.Error())
	}

	stringToDisplay := string(bufferToPrint)
	fmt.Print(stringToDisplay)

}

func Println(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {

	// The WASM plugin copied data into the shared memory
	// Red this data
	bufferToPrint, err := mem.ReadBytesFromMemory(plugin, stack)

	if err != nil {
		log.Println("🔴 Println, ReadBytesFromMemory", err.Error())
	}

	stringToDisplay := string(bufferToPrint)
	fmt.Println(stringToDisplay)

}

func Log(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {

	// The WASM plugin copied data into the shared memory
	// Red this data
	bufferToPrint, err := mem.ReadBytesFromMemory(plugin, stack)

	if err != nil {
		log.Println("🔴 Log, ReadBytesFromMemory", err.Error())
	}

	stringToDisplay := string(bufferToPrint)
	fmt.Println(time.Now(), ":", stringToDisplay)

}
