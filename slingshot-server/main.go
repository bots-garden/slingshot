package main

/*
  Slingshot is an Extism plugins launcher
*/

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
- Download a plugin from a location
- OnStart
- OnStop
- Certificats
- Redis
- improve CLI XP : https://dev.to/cherryramatis/bonzai-and-how-to-create-a-personal-cli-to-rule-them-all-1bnl
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

	log_string := slingshot.DefineHostFunctionCallBack(
		"hostLog",
		callbacks.Log,
	)

	get_message := slingshot.DefineHostFunctionCallBack(
		"hostGetMessage",
		callbacks.GetMessage,
	)

	memory_set := slingshot.DefineHostFunctionCallBack(
		"hostMemorySet",
		callbacks.MemorySet,
	)

	memory_get := slingshot.DefineHostFunctionCallBack(
		"hostMemoryGet",
		callbacks.MemoryGet,
	)

	get_env := slingshot.DefineHostFunctionCallBack(
		"hostGetEnv",
		callbacks.GetEnv,
	)

	init_redis_cli := slingshot.DefineHostFunctionCallBack(
		"hostInitRedisClient",
		callbacks.InitRedisClient,
	)

	redis_set := slingshot.DefineHostFunctionCallBack(
		"hostRedisSet",
		callbacks.RedisSet,
	)

	slingshot.AppendHostFunction(get_message)
	slingshot.AppendHostFunction(print_string)
	slingshot.AppendHostFunction(log_string)
	slingshot.AppendHostFunction(memory_set)
	slingshot.AppendHostFunction(memory_get)
	slingshot.AppendHostFunction(get_env)
	slingshot.AppendHostFunction(init_redis_cli)
	slingshot.AppendHostFunction(redis_set)

	err := slingshot.InitializePluging(ctx, "slingshotplug", manifest, config, slingshot.GetHostFunctions())
	/*
	err = slingshot.InitializePluging(ctx, "slingshotplug0", manifest, config, slingshot.GetHostFunctions())
	err = slingshot.InitializePluging(ctx, "slingshotplug1", manifest, config, slingshot.GetHostFunctions())
	err = slingshot.InitializePluging(ctx, "slingshotplug2", manifest, config, slingshot.GetHostFunctions())
	err = slingshot.InitializePluging(ctx, "slingshotplug3", manifest, config, slingshot.GetHostFunctions())
	err = slingshot.InitializePluging(ctx, "slingshotplug4", manifest, config, slingshot.GetHostFunctions())
	*/

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
		
		//plugin, err := slingshot.SelectPlugin()

		if err != nil {
			log.Println("🔴 !!! Error when getting the plugin", err)
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		//out, err := plugin.Call(wasmFunctionName, params)

		_, out, err := plugin.Call(wasmFunctionName, params)

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
