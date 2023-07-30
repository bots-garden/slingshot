package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/extism/extism"
	"github.com/gofiber/fiber/v2"
)

// store all your plugins in a normal Go hash map, protected by a Mutex
var m sync.Mutex
var plugins = make(map[string]extism.Plugin)

func StorePlugin(plugin extism.Plugin) {
	// store all your plugins in a normal Go hash map, protected by a Mutex
	plugins["code"] = plugin
}

func GetPlugin() (extism.Plugin, error) {
	if plugin, ok := plugins["code"]; ok {
		return plugin, nil
	} else {
		return extism.Plugin{}, errors.New("🔴 no plugin")
	}
}


func main() {
	wasmFilePath := os.Args[1:][0]
	wasmFunctionName := os.Args[1:][1]
	httpPort := os.Args[1:][2]

	//var counter = 0

	ctx := extism.NewContext()

	defer ctx.Free() // this will free the context and all associated plugins

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
	}

	/*
	plugin, err := ctx.PluginFromManifest(manifest, []extism.Function{}, true)
	if err != nil {
		panic(err)
	}
	*/

	plugin, err := ctx.PluginFromManifest(manifest, []extism.Function{}, true)
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	StorePlugin(plugin)
	
	/*
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		DisableKeepalive:      true,
		Concurrency:           100000,
	})
	*/

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/", func(c *fiber.Ctx) error {

		params := c.Body()

		/*
		plugin, err := ctx.PluginFromManifest(manifest, []extism.Function{}, true)
		if err != nil {
			//panic(err)
			fmt.Println(err)
			c.Status(http.StatusConflict)
			return c.SendString(err.Error())
		}
		*/
		m.Lock()
		// don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
		defer m.Unlock()
		
		plugin, err := GetPlugin()

		if err != nil {
			log.Println("🔴 !!! Error when getting the plugin", err)
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}
		

		out, err := plugin.Call(wasmFunctionName, params)

		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusConflict)
			return c.SendString(err.Error())
			//os.Exit(1)
		} else {
			c.Status(http.StatusOK)
			//fmt.Println(counter, string(out))
			//counter ++
			return c.SendString(string(out))
		}

	})

	fmt.Println("🌍 http server is listening on:", httpPort)
	app.Listen(":" + httpPort)
}
