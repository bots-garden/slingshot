package cmds

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slingshot-server/plg"
	"slingshot-server/slingshot"

	"github.com/gofiber/fiber/v2"
)

// Start the slingshot HTTP server (triggered by the `listen` command, from parseCommand)
func Start(wasmFilePath string, wasmFunctionName string, httpPort string, logLevel string, allowHosts string, allowPaths string, config string) {


	plg.Initialize("slingshotplug", wasmFilePath, logLevel, allowHosts, allowPaths, config)

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

		extismPlugin, err := plg.GetPlugin("slingshotplug")

		if err != nil {
			log.Println("🔴 Error when getting the plugin", err)
			c.Status(http.StatusInternalServerError)
			return c.SendString(err.Error())
		}

		/*
			if extismPlugin.MainFunction == true {
				_, _, err := extismPlugin.Plugin.Call("_start", nil)
				if err != nil {
					log.Println("🔴 Error with _start function", err)
				}
			}
		*/

		_, response, err := extismPlugin.Plugin.Call(wasmFunctionName, params)
		if err != nil {
			log.Println("🔴 Error when calling the function", err)
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
				// test encoding of TextBody
				decodedStrAsByteSlice, err := base64.StdEncoding.DecodeString(string(httpResponse.TextBody))
				if err != nil {
					return c.SendString(httpResponse.TextBody)
				}
				//return c.SendString(httpResponse.TextBody)
				return c.SendString(string(decodedStrAsByteSlice))
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
