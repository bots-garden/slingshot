package callbacks

import (
	"context"
	"fmt"
	"log"
	"slingshot-server/slingshot"
	"time"

	"github.com/extism/extism"
)

func Print(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

	// The WASM plugin copied data into the shared memory
	// Red this data
	bufferToPrint, err := slingshot.ReadBytesFromMemory(plugin, stack)

	if err != nil {
		log.Println("🔴 Print, ReadBytesFromMemory", err.Error())
	}

	stringToDisplay := string(bufferToPrint)
	fmt.Println(stringToDisplay)

}

func Log(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

	// The WASM plugin copied data into the shared memory
	// Red this data
	bufferToPrint, err := slingshot.ReadBytesFromMemory(plugin, stack)

	if err != nil {
		log.Println("🔴 Log, ReadBytesFromMemory", err.Error())
	}

	stringToDisplay := string(bufferToPrint)
	fmt.Println(time.Now(), ":", stringToDisplay)

}
