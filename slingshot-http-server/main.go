package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"slingshot-http-server/slingshotplugin"
	"sync"

	"github.com/extism/extism"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	//"github.com/tetratelabs/wazero/sys"
)

var mutex sync.Mutex

var memoryMap = map[string]string{
	"hello":   "🖖 Hello World 🌍",
	"message": "I 💜 Extism 😍",
}

func main() {

	wasmFilePath := os.Args[1:][0]
	wasmFunctionName := os.Args[1:][1]
	httpPort := os.Args[1:][2]

	// this is for tests
	var counter = 0

	ctx := context.Background() 

	config := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
	}

	print_string := extism.HostFunction{
		Name:      "hostPrint",
		Namespace: "env",
		Callback: func(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

			offset := stack[0]
			bufferInput, err := plugin.ReadBytes(offset)

			if err != nil {
				fmt.Println("🥵", err.Error())
				panic(err)
			}

			stringToDisplay := string(bufferInput)
			fmt.Println(stringToDisplay)

			plugin.Free(offset)

			stack[0] = 0
		},
		Params:  []api.ValueType{api.ValueTypeI64},
		Results: []api.ValueType{api.ValueTypeI64},
	}

	memory_get := extism.HostFunction{
		Name:      "hostMemoryGet",
		Namespace: "env",
		Callback: func(ctx context.Context, plugin *extism.CurrentPlugin, userData interface{}, stack []uint64) {

			offset := stack[0]
			bufferInput, err := plugin.ReadBytes(offset)

			if err != nil {
				fmt.Println("🥵", err.Error())
				panic(err)
			}

			keyStr := string(bufferInput)
			fmt.Println("🟢 keyStr:", keyStr) // this is for test

			returnValue := memoryMap[keyStr]

			plugin.Free(offset)
			offset, err = plugin.WriteBytes([]byte(returnValue))
			if err != nil {
				fmt.Println("😡", err.Error())
				panic(err)
			}

			stack[0] = offset
		},
		Params:  []api.ValueType{api.ValueTypeI64},
		Results: []api.ValueType{api.ValueTypeI64},
	}

	hostFunctions := []extism.HostFunction{
		memory_get,
		print_string,
	}

	pluginInst, err := extism.NewPlugin(ctx, manifest, config, hostFunctions) // new

	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	slingshotplugin.StorePlugin("slingshotplug", pluginInst)

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

		mutex.Lock()
		// don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
		defer mutex.Unlock()


		plugin, err := slingshotplugin.GetPlugin("slingshotplug")

		if err != nil {
			log.Println("🔴 !!! Error when getting the plugin", err)
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		//out, err := plugin.Call(wasmFunctionName, params)

		_, out, err := plugin.Call(wasmFunctionName, params) // new

		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusConflict)
			return c.SendString(err.Error())
			//os.Exit(1)
		} else {
			c.Status(http.StatusOK)
			fmt.Println("🟢 ->", counter, ": ", string(out))
			counter++
			return c.SendString(string(out))
		}

	})

	fmt.Println("🌍 http server is listening on:", httpPort)
	app.Listen(":" + httpPort)
}
