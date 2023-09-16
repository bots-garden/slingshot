package plg

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"slingshot-server/hof"
	"slingshot-server/initcbk"
	"strings"

	"github.com/extism/extism"
	"github.com/tetratelabs/wazero"
)

/*
With Wasm, the main function of a wasm module is considered
as a function names "_start".
If this function exists ExtismPlugin.MainFunction == true
And then, it will be triggered first
*/

type ExtismPlugin struct {
	Plugin       *extism.Plugin
	MainFunction bool
}

// store all your plugins in a normal Go hash map, protected by a Mutex
var plugins = make(map[string]ExtismPlugin)

func StorePlugin(key string, plugin ExtismPlugin) {
	// store all your plugins in a normal Go hash map, protected by a Mutex
	plugins[key] = plugin

}

func GetPlugin(key string) (ExtismPlugin, error) {

	if plugin, ok := plugins[key]; ok {
		return plugin, nil
	} else {
		return ExtismPlugin{}, errors.New("🔴 no plugin")
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

func getLevel(logLevel string) extism.LogLevel {
	level := extism.Off
	switch logLevel {
	case "error":
		level = extism.Error
	case "warn":
		level = extism.Warn
	case "info":
		level = extism.Info
	case "debug":
		level = extism.Debug
	case "trace":
		level = extism.Trace
	}
	return level
}

func GetPluginConfig(logLevel string) extism.PluginConfig {
	level := getLevel(logLevel)

	config := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
		LogLevel:     &level,
	}
	return config
}

func GetManifest(wasmFilePath string, allowHosts string, allowPaths string, config string) extism.Manifest {

	hosts := strings.Split(strings.ReplaceAll(allowHosts, " ", ""), ",")

	var paths map[string]string
	unmarshallError := json.Unmarshal([]byte(allowPaths), &paths)
	if unmarshallError != nil {
		log.Println("🔴 (GetManifest/paths) Error:", unmarshallError.Error())
		os.Exit(1)
	}

	var manifestConfig map[string]string
	unmarshallError = json.Unmarshal([]byte(config), &manifestConfig)
	if unmarshallError != nil {
		log.Println("🔴 (GetManifest/manifestConfig) Error:", unmarshallError.Error())
		os.Exit(1)
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
		AllowedHosts: hosts,
		AllowedPaths: paths,
		Config:       manifestConfig,
	}

	return manifest
}

// Create a plugin and store it into the plugins map
func InitializePluging(ctx context.Context, pluginName string, manifest extism.Manifest, config extism.PluginConfig, hostsFunctions []extism.HostFunction) error {
	pluginInst, err := extism.NewPlugin(ctx, manifest, config, hof.GetHostFunctions())

	// Here we test if there is a _start function into the wasm module
	// with Extism, this function does not exist with Rust (lib)
	StorePlugin(pluginName, ExtismPlugin{Plugin: pluginInst, MainFunction: pluginInst.FunctionExists("_start")})

	return err
}

// Initialise the extism wasm plugin
func Initialize(idPlugin string, wasmFilePath string, logLevel string, allowHosts string, allowPaths string, config string) context.Context {

	ctx := context.Background()

	pluginConfig := GetPluginConfig(logLevel) // LogLevel
	manifest := GetManifest(wasmFilePath, allowHosts, allowPaths, config)

	// load all the host function callbacks
	initcbk.LoadHostFunctionCallBacks()

	errPlgInit := InitializePluging(ctx, idPlugin, manifest, pluginConfig, hof.GetHostFunctions())

	if errPlgInit != nil {
		log.Println("🔴 Error when loading the plugin", errPlgInit)
		os.Exit(1)
	}
	return ctx
}
