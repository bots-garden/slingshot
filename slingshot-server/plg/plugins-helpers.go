package plg

import (
	"context"
	"errors"
	"log"
	"os"
	"slingshot-server/hof"
	"slingshot-server/initcbk"

	"github.com/extism/extism"
	"github.com/tetratelabs/wazero"
)

// store all your plugins in a normal Go hash map, protected by a Mutex
var plugins = make(map[string]*extism.Plugin)

func StorePlugin(key string, plugin *extism.Plugin) {
	// store all your plugins in a normal Go hash map, protected by a Mutex
	plugins[key] = plugin

}

func GetPlugin(key string) (extism.Plugin, error) {

	if plugin, ok := plugins[key]; ok {
		return *plugin, nil
	} else {
		return extism.Plugin{}, errors.New("🔴 no plugin")
	}
}

/*
func SelectPlugin() (extism.Plugin, error) {
	min := 0
	max := 4
	i := rand.Intn(max-min) + min

	if plugin, ok := plugins["slingshotplug"+strconv.Itoa(i)]; ok {
		return *plugin, nil
	} else {
		return extism.Plugin{}, errors.New("🔴 no plugin")
	}
}
*/

func GetPluginConfig() extism.PluginConfig {
	config := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
	}
	return config
}

func GetManifest(wasmFilePath string) extism.Manifest {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
	}
	return manifest
}

// Create a plugin and store it into the plugins map
func InitializePluging(ctx context.Context, pluginName string, manifest extism.Manifest, config extism.PluginConfig, hostsFunctions []extism.HostFunction) error {
	pluginInst, err := extism.NewPlugin(ctx, manifest, config, hof.GetHostFunctions())
	StorePlugin(pluginName, pluginInst)

	return err
}

// Initialise the extism wasm plugin
func Initialize(idPlugin string, wasmFilePath string) context.Context {

	ctx := context.Background()

	config := GetPluginConfig()
	manifest := GetManifest(wasmFilePath)

	// load all the host function callbacks
	initcbk.LoadHostFunctionCallBacks()

	errPlgInit := InitializePluging(ctx, idPlugin, manifest, config, hof.GetHostFunctions())

	if errPlgInit != nil {
		log.Println("🔴 Error when loading the plugin", errPlgInit)
		os.Exit(1)
	}
	return ctx
}
