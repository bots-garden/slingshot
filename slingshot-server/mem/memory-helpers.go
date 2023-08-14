package mem

import (
	"encoding/json"

	"github.com/extism/extism"
)

// Read bytes from the shared memory. The bytes were copied by the wasm plugin
func ReadBytesFromMemory(plugin *extism.CurrentPlugin, stack []uint64) ([]byte, error) {
	offset := stack[0] // <- always only one parameter
	bufferInput, err := plugin.ReadBytes(offset)
	if err != nil {
		return nil, err
	}
	plugin.Free(offset) //? should I do another function without this
	return bufferInput, nil
}

// Read string from the shared memory. The bytes were copied by the wasm plugin
func ReadStringFromMemory(plugin *extism.CurrentPlugin, stack []uint64) (string, error) {
	offset := stack[0] // <- always only one parameter
	stringInput, err := plugin.ReadString(offset)
	if err != nil {
		return "", err
	}
	plugin.Free(offset) //? should I do another function without this
	return stringInput, nil
}

// Read a Json buffer from the shared memory
// and unmarshall it on a model (structure)
func ReadJsonFromMemory(plugin *extism.CurrentPlugin, stack []uint64, model any) error {

	dataFromWasmModule, errReadBytes := ReadBytesFromMemory(plugin, stack)
	if errReadBytes != nil {
		return errReadBytes
	}
	errMarshal := json.Unmarshal(dataFromWasmModule, &model)
	if errMarshal != nil {
		return errMarshal
	}

	return nil
}

func CopyBytesToMemory(plugin *extism.CurrentPlugin, stack []uint64, value []byte) error {
	offset, err := plugin.WriteBytes(value)
	if err != nil {
		return err
	}
	stack[0] = offset
	return nil
}

func CopyStringToMemory(plugin *extism.CurrentPlugin, stack []uint64, value string) error {
	offset, err := plugin.WriteString(value)
	if err != nil {
		return err
	}
	stack[0] = offset
	return nil
}

func CopyJsonToMemory(plugin *extism.CurrentPlugin, stack []uint64, model any) error {
	jsonBytes, err := json.Marshal(&model)
	if err != nil {
		return err
	}
	return CopyBytesToMemory(plugin, stack, jsonBytes)

}
