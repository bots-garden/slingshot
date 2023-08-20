package main

/*
  Slingshot is an Extism plugins launcher
*/

import (
	"flag"
	"fmt"
	"os"
	"slingshot-server/cmds"
	"slingshot-server/infos"
)

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

func parseCommand(command string, args []string) error {
	//fmt.Println("Command:", command)
	//fmt.Println("Args:", args)
	switch command {
	case "start", "listen":
		//fmt.Println("start")
		flagSet := flag.NewFlagSet("listen", flag.ExitOnError)

		httpPort := flagSet.String("http-port", "8080", "http port")
		handler := flagSet.String("handler", "handle", "wasm function name")
		wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")

		flagSet.Parse(args)

		fmt.Println("🌍 http-port:", *httpPort)
		fmt.Println("🚀 handler  :", *handler)
		fmt.Println("📦 wasm     :", *wasmFile)

		cmds.Start(*wasmFile, *handler, *httpPort)

		return nil

	// TODO: MQTT, Nats,...

	case "redis":
		//redisCmd := flag.Args()[0]
		subCommand := flag.Args()[1]

		flagSet := flag.NewFlagSet("redis", flag.ExitOnError)

		// Allow to use an existing redis client
		redisUri := flagSet.String("uri", "rediss://default:pwd@redis-something:port", "Redis URI")
		redisClientId := flagSet.String("client-id", "something", "Redis client id")
		handler := flagSet.String("handler", "handle", "wasm function name")
		wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")

		maskVariables := flagSet.Bool("mask-variables", true, "")

		switch subCommand {
		case "subscribe":

			redisChannel := flagSet.String("channel", "knock-knock", "Redis channel")

			flagSet.Parse(args[1:]) // from 1, because of the subCommand
			if *maskVariables != true {
				fmt.Println("🌍 redis URI      :", *redisUri)
			} else {
				fmt.Println("🌍 redis URI      :", "*****")
			}

			fmt.Println("🌍 redis Client Id:", *redisClientId)
			fmt.Println("🚀 handler        :", *handler)
			fmt.Println("📦 wasm           :", *wasmFile)
			fmt.Println("📺 channel        :", *redisChannel)

			cmds.RedisSubscribe(*wasmFile, *handler, *redisChannel, *redisUri, *redisClientId)

		case "memdb":

			// could be use for a mono redis connection

		default:
			return fmt.Errorf("🔴 invalid subcommand: '%s'\n\n%s\n", subCommand, infos.Usage)
		}

		return nil

	case "nats":
		subCommand := flag.Args()[1]

		flagSet := flag.NewFlagSet("nats", flag.ExitOnError)

		// NATS flags
		// TODO, id, conn string...

		handler := flagSet.String("handler", "handle", "wasm function name")
		wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")

		// Allow to use an existing NATS client
		natsUrl := flagSet.String("url", "nats://somebody:secretpassword@demo.nats.io:4222", "Nats URL")
		natsConnectionId := flagSet.String("connection-id", "something", "NATS connection id")

		maskVariables := flagSet.Bool("mask-variables", true, "")

		switch subCommand {
		case "subscribe":
			
			natsSubject := flagSet.String("subject", "knock-knock", "NATS subject")
			flagSet.Parse(args[1:]) // from 1, because of the subCommand
			if *maskVariables != true {
				fmt.Println("🌍 NATS URL      :", *natsUrl)
			} else {
				fmt.Println("🌍 NATS URL      :", "*****")
			}

			fmt.Println("🌍 NATS Connection Id:", *natsConnectionId)
			fmt.Println("🚀 handler           :", *handler)
			fmt.Println("📦 wasm              :", *wasmFile)
			fmt.Println("📺 Subject           :", *natsSubject)

			cmds.NatsSubscribe(*wasmFile, *handler, *natsSubject, *natsUrl, *natsConnectionId)

		default:
			return fmt.Errorf("🔴 invalid subcommand: '%s'\n\n%s\n", subCommand, infos.Usage)
		}

		return nil

	case "cli", "run":
		flagSet := flag.NewFlagSet("run", flag.ExitOnError)

		handler := flagSet.String("handler", "handle", "wasm function name")
		wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")
		input := flagSet.String("input", "hello", "input data for the wasm plugin")

		flagSet.Parse(args)
		cmds.Execute(*wasmFile, *handler, *input)

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
		return fmt.Errorf("🔴 invalid command: '%s'\n\n%s\n", command, infos.Usage)
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
