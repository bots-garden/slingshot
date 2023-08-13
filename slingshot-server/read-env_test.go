package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/slingshot"
	"testing"
)

func initPlugin(wasmFilePath string, pluginId string) {
	ctx := context.Background()

	config := slingshot.GetPluginConfig()
	manifest := slingshot.GetManifest(wasmFilePath)

	// Add an host function
	get_env := slingshot.DefineHostFunctionCallBack(
		"hostGetEnv",
		callbacks.GetEnv,
	)
	slingshot.AppendHostFunction(get_env)

	err := slingshot.InitializePluging(ctx, pluginId, manifest, config, slingshot.GetHostFunctions())
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

	plugin, err := slingshot.GetPlugin("slingshotplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}

	_, out, err := plugin.Call(wasmFunctionName, nil)

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

	plugin, err := slingshot.GetPlugin("slingshotplug")
	if err != nil {
		log.Println("🔴 !!! Error when getting the plugin", err)
		os.Exit(1)
	}
	_, out, err := plugin.Call(wasmFunctionName, nil)

	result := string(out)
	fmt.Println("🟠", result)
	if result != expected {
		fmt.Println("🔴", "TestReadEmptyEnvVar")
		t.Errorf("Result %q, Expected %q", result, expected)
	} else {
		fmt.Println("🟢", "TestReadEmptyEnvVar")
	}
}
