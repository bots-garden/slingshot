package callbacks

import (
	"context"
	"encoding/base64"
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
		log.Println("🔴 ReadFile, CopyJsonToMemory:", errResult.Error())
	}

}

func WriteFile(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	var result = slingshot.StringResult{}
	var arguments slingshot.FileRecord

	// Read data from the shared memory
	err := mem.ReadJsonFromMemory(plugin, stack, &arguments)

	// Construct the result
	if err != nil {
		result.Failure = err.Error()
		result.Success = ""
	} else {

		// Decoding
		var content string
		decodedStrAsByteSlice, err := base64.StdEncoding.DecodeString(arguments.Content)
		if err != nil {
			content = arguments.Content
		} else {
			content = string(decodedStrAsByteSlice)
		}

		// Write file
		errWriteFile := os.WriteFile(string(arguments.Path), []byte(content), 0644)

		if errWriteFile != nil {
			result.Failure = errWriteFile.Error()
			result.Success = ""
		} else {
			result.Failure = ""
			result.Success = "saved"
		}
	}
	// Copy the result to the memory
	errResult := mem.CopyJsonToMemory(plugin, stack, result)

	if errResult != nil {
		log.Println("🔴 WriteFile, CopyJsonToMemory:", errResult.Error())
	}

}
