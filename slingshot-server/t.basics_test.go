package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/plg"
	"testing"

	"github.com/extism/extism"
)

func TestStorePlugin(t *testing.T) {
	wasmFilePath := "../plugins/tests/some-functions/some-functions.wasm"
	wasmFunctionName := "hello"

	ctx := context.Background()
	manifest := plg.GetManifest(wasmFilePath, "*", "{}", "{}")
	config := plg.GetPluginConfig("info")

	pluginInst, err := extism.NewPlugin(ctx, manifest, config, nil)
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	plg.StorePlugin("helloPlugin", plg.ExtismPlugin{Plugin: pluginInst, MainFunction: false})

	plugin, err := plg.GetPlugin("helloPlugin")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}

	fmt.Println("🟠", plugin)
	fmt.Println("🟠", plugin.Plugin.FunctionExists(wasmFunctionName))

	if plugin.Plugin.FunctionExists(wasmFunctionName) != true {
		fmt.Println("🔴", "TestStorePlugin")
		t.Errorf("Error didn't find %q", wasmFunctionName)
	} else {
		fmt.Println("🟢", "TestStorePlugin")
	}
}

func TestGetPluginConfig(t *testing.T) {

	config := plg.GetPluginConfig("info")
	if config.EnableWasi != true {
		fmt.Println("🔴", "TestGetPluginConfig")
		t.Errorf("Error EnableWasi should be set to true")
	} else {
		fmt.Println("🟢", "TestGetPluginConfig")
	}
}

// Be sure that GetManifest returns the appropriate manifest
func TestGetManifest(t *testing.T) {
	wasmFilePath := "../plugins/tests/some-functions/some-functions.wasm"
	manifestForTest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
	}

	manifest := plg.GetManifest(wasmFilePath, "*", "{}", "{}")
	fmt.Println("🟠", manifest.Wasm[0])
	if manifest.Wasm[0] != manifestForTest.Wasm[0] {
		fmt.Println("🔴", "TestGetManifest")
		t.Errorf("Error with the manifest")
	} else {
		fmt.Println("🟢", "TestGetManifest")
	}

}

func TestInitializeWasmPlugin(t *testing.T) {
	wasmFilePath := "../plugins/tests/some-functions/some-functions.wasm"
	ctx := context.Background()

	config := plg.GetPluginConfig("info")
	manifest := plg.GetManifest(wasmFilePath, "*", "{}", "{}")

	err := plg.InitializePluging(ctx, "slingshotplug", manifest, config, nil)
	if err != nil {
		fmt.Println("🔴", "TestInitializeWasmPlugin")
		t.Errorf("Error when loading the plugin %q", err)
	} else {
		fmt.Println("🟢", "TestInitializeWasmPlugin")
	}
}
