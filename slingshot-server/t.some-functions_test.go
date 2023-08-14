package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/plg"
	"testing"
)

func TestCallHello(t *testing.T) {

	// /home/k33g/workspaces/slingshot/plugins/tests/some-functions
	wasmFilePath := "../plugins/tests/some-functions/some-functions.wasm"
	wasmFunctionName := "hello"
	wasmFunctionArgument := "Bob Morane"
	expected := "Hello Bob Morane"

	ctx := context.Background()

	config := plg.GetPluginConfig()
	manifest := plg.GetManifest(wasmFilePath)

	err := plg.InitializePluging(ctx, "slingshotplug", manifest, config, nil)
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	plugin, err := plg.GetPlugin("slingshotplug")

	_, out, err := plugin.Call(wasmFunctionName, []byte(wasmFunctionArgument))

	result := string(out)

	if result != expected {
		fmt.Println("🔴", "TestCallHello")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestCallHello")
	}
}

// TODO: test with JSON data
