package cmds

import (
	"fmt"
	"log"
	"os"
	"slingshot-server/plg"
)

// Execute is triggered by the `run` command (from parseCommand)
func Execute(wasmFilePath string, wasmFunctionName string, data string) {
	plg.Initialize("slingshotplug", wasmFilePath)
	extismPlugin, err := plg.GetPlugin("slingshotplug")

	if err != nil {
		log.Println("🔴 Error when getting the plugin", err)
		os.Exit(1)
	}

	if extismPlugin.Plugin.FunctionExists(wasmFunctionName) != true {
		log.Println("🔴 Error:", wasmFunctionName, "does not exist")
		os.Exit(1)
	}

	if extismPlugin.MainFunction == true {
		_, _, err := extismPlugin.Plugin.Call("_start", nil)
		if err != nil {
			fmt.Println("🔴 Error:", err)
			os.Exit(1)
		}

	}

	_, output, err := extismPlugin.Plugin.Call(wasmFunctionName, []byte(data))
	if err != nil {
		fmt.Println("🔴 Error:", err)
		os.Exit(1)
	}

	// Display output content, only if the wasm plugin returns something
	if (len(output)) > 0 {
		fmt.Println(string(output))
	}

}
