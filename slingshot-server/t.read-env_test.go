package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/hof"
	"slingshot-server/plg"
	"testing"
)

func initPlugin(wasmFilePath string, pluginId string) {
	ctx := context.Background()

	config := plg.GetPluginConfig()
	manifest := plg.GetManifest(wasmFilePath)

	// Add an host function
	get_env := hof.DefineHostFunctionCallBack(
		"hostGetEnv",
		callbacks.GetEnv,
	)
	hof.AppendHostFunction(get_env)

	err := plg.InitializePluging(ctx, pluginId, manifest, config, hof.GetHostFunctions())
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

}

func TestReadEnvVar(t *testing.T) {
	// SLINGSHOT_MESSAGE
	wasmFilePath := "../plugins/tests/read-env/read-env.wasm"
	wasmFunctionName := "read_env_var"
	//expected := "Hello Bobby Morane"
	expected := "Hello Bob Morane"

	os.Setenv("SLINGSHOT_MESSAGE", "Hello Bob Morane")

	initPlugin(wasmFilePath, "slingshotplug")

	plugin, err := plg.GetPlugin("slingshotplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}

	_, out, err := plugin.Plugin.Call(wasmFunctionName, nil)

	result := string(out)
	fmt.Println("🟠", result)
	if result != expected {
		fmt.Println("🔴", "TestReadEnvVar")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestReadEnvVar")
	}
}

func TestReadEmptyEnvVar(t *testing.T) {
	// SLINGSHOT_MESSAGE
	wasmFilePath := "../plugins/tests/read-env/read-env.wasm"
	wasmFunctionName := "read_env_var"
	expected := "EMPTY"

	//os.Setenv("SLINGSHOT_MESSAGE", "")
	os.Unsetenv("SLINGSHOT_MESSAGE")

	initPlugin(wasmFilePath, "slingshotplug")

	plugin, err := plg.GetPlugin("slingshotplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	_, out, err := plugin.Plugin.Call(wasmFunctionName, nil)

	result := string(out)
	fmt.Println("🟠", result)
	if result != expected {
		fmt.Println("🔴", "TestReadEmptyEnvVar")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestReadEmptyEnvVar")
	}
}
