package main

/*
  Slingshot is an Extism plugins launcher
*/

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slingshot-server/hof"
	"slingshot-server/infos"
	"slingshot-server/initcbk"
	"slingshot-server/plg"
	"slingshot-server/slingshot"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var mutex sync.Mutex

/*
# TODO

- Download a plugin from a location
- OnStart
- OnStop
- Certificates (https)
- healthcheck
- monitoring
- PostGRESQL
- hostfunctions: at start we can choose to activate or deactivate some hostfunctions
- input host function?
- onkey host function?
*/

// Initialise the extism wasm plugin
func initialize(idPlugin string, wasmFilePath string) context.Context {

	ctx := context.Background()

	config := plg.GetPluginConfig()
	manifest := plg.GetManifest(wasmFilePath)

	// load all the host function callbacks
	initcbk.LoadHostFunctionCallBacks()

	errPlgInit := plg.InitializePluging(ctx, idPlugin, manifest, config, hof.GetHostFunctions())

	if errPlgInit != nil {
		log.Println("🔴 !!! Error when loading the plugin", errPlgInit)
		os.Exit(1)
	}
	return ctx
}

// Start the slingshot HTTP server
func start(wasmFilePath string, wasmFunctionName string, httpPort string) {

	// this is for tests
	//var counter = 0

	initialize("slingshotplug", wasmFilePath)

	/*
		app := fiber.New(fiber.Config{
			DisableStartupMessage: true,
			DisableKeepalive:      true,
			Concurrency:           100000,
		})
	*/

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	handler := func(c *fiber.Ctx, params []byte) error {

		mutex.Lock()
		// don't forget to release the lock on the Mutex, sometimes its best to `defer m.Unlock()` right after yout get the lock
		defer mutex.Unlock()

		plugin, err := plg.GetPlugin("slingshotplug")

		if err != nil {
			log.Println("🔴 Error when getting the plugin", err)
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		_, response, err := plugin.Call(wasmFunctionName, params)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusConflict)
			return c.SendString(err.Error())
			//os.Exit(1)
		}

		httpResponse := slingshot.HTTPResponse{}

		errMarshal := json.Unmarshal(response, &httpResponse)
		if errMarshal != nil {
			fmt.Println(errMarshal)
			c.Status(http.StatusConflict)
			return c.SendString(errMarshal.Error())
		} else {

			/*
				fmt.Println("🟢 ->", counter, ": ", string(response))
				fmt.Println("🟣", httpResponse)
				counter++
			*/

			c.Status(httpResponse.StatusCode)
			// set headers
			for key, value := range httpResponse.Headers {
				c.Set(key, value)
			}
			if len(httpResponse.TextBody) > 0 {
				return c.SendString(httpResponse.TextBody)
			}
			// send JSON body
			jsonBody, err := json.Marshal(httpResponse.JsonBody)
			if err != nil {
				log.Println("🔴 Error when marshal the content", err)
				c.Status(http.StatusInternalServerError) // .🤔
				return c.SendString(errMarshal.Error())
			}
			return c.Send(jsonBody)
		}
	}

	app.All("/", func(c *fiber.Ctx) error {

		request := slingshot.HTTPRequest{
			Method:  c.Method(),
			BaseUrl: c.BaseURL(),
			Body:    string(c.Body()),
			Headers: c.GetReqHeaders(),
		}

		jsonRequest, err := json.Marshal(request)

		if err != nil {
			log.Println("🔴 Error when marshal the request", err)
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}
		//fmt.Println("🖐️", string(jsonRequest))

		return handler(c, jsonRequest)
	})

	fmt.Println("🌍 slingshot server is listening on:", httpPort)
	app.Listen(":" + httpPort)
}

func execute(wasmFilePath string, wasmFunctionName string, data string) {
	initialize("slingshotplug", wasmFilePath)
	plugin, err := plg.GetPlugin("slingshotplug")

	if err != nil {
		log.Println("🔴 Error when getting the plugin", err)
		os.Exit(1)
	}

	_, output, err := plugin.Call(wasmFunctionName, []byte(data))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	// Display output content, only if the wasm plugin returns something
	if(len(output)) > 0 {
		fmt.Println(string(output))
	}

}

func parseCommand(command string, args []string) error {
	//fmt.Println("Command:", command)
	//fmt.Println("Args:", args)
	switch command {
	case "start":
		//fmt.Println("start")
		flagSet := flag.NewFlagSet("start", flag.ExitOnError)

		httpPort := flagSet.String("http-port", "8080", "http port")
		handler := flagSet.String("handler", "handle", "wasm function name")
		wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")

		flagSet.Parse(args)

		fmt.Println("🌍 http-port:", *httpPort)
		fmt.Println("🚀 handler  :", *handler)
		fmt.Println("📦 wasm     :", *wasmFile)

		start(*wasmFile, *handler, *httpPort)

		return nil

	case "cli":
		flagSet := flag.NewFlagSet("start", flag.ExitOnError)

		handler := flagSet.String("handler", "handle", "wasm function name")
		wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")
		input := flagSet.String("input", "hello", "input data for the wasm plugin")

		flagSet.Parse(args)
		execute(*wasmFile, *handler, *input)

		return nil

	case "version":
		fmt.Println(infos.GetVersion())
		//os.Exit(0)
		return nil
	case "help":
		fmt.Println(infos.Help)
		return nil
	case "about":
		fmt.Println(infos.About)
		return nil
	default:
		return fmt.Errorf("invalid command: '%s'\n\n%s\n", command, infos.Usage)
	}
}

func main() {

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println(infos.Usage)
		os.Exit(0)
	}

	command := flag.Args()[0]

	errCmd := parseCommand(command, flag.Args()[1:])
	if errCmd != nil {
		fmt.Println(errCmd)
		os.Exit(1)
	}

}
