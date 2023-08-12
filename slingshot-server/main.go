package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"slingshot-server/callbacks"
	"slingshot-server/slingshot"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var mutex sync.Mutex

/* TODO:
- Get and Set for the memory
- Download a plugin from a location
- OnStart
- OnStop
*/

func main() {

	wasmFilePath := os.Args[1:][0]
	wasmFunctionName := os.Args[1:][1]
	httpPort := os.Args[1:][2]

	// this is for tests
	var counter = 0

	ctx := context.Background()

	config := slingshot.GetPluginConfig()
	manifest := slingshot.GetManifest(wasmFilePath)

	print_string := slingshot.DefineHostFunctionCallBack(
		"hostPrint",
		callbacks.Print,
	)

	memory_get := slingshot.DefineHostFunctionCallBack(
		"hostMemoryGet",
		callbacks.MemoryGet,
	)

	slingshot.AppendHostFunction(memory_get)
	slingshot.AppendHostFunction(print_string)

	err := slingshot.InitializePluging(ctx, "slingshotplug", manifest, config, slingshot.GetHostFunctions())

	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

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

		plugin, err := slingshot.GetPlugin("slingshotplug")

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
