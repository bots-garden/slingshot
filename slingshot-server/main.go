package main

/*
  Slingshot is an Extism plugins launcher
*/

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slingshot-server/cmds"
	"slingshot-server/infos"

	"github.com/go-resty/resty/v2"
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

func downloadWasmFile(wasmFileURL, wasmFilePath, authenticationHeaderName, authenticationHeaderValue string) error {
	// authenticationHeader:
	// Example: "PRIVATE-TOKEN: ${GITLAB_WASM_TOKEN}"
	client := resty.New()

	if authenticationHeaderName != "" {
		client.SetHeader(authenticationHeaderName, authenticationHeaderValue)
	}

	resp, err := client.R().
		SetOutput(wasmFilePath).
		Get(wasmFileURL)

	if resp.IsError() {
		return errors.New("error while downloading the wasm file")
	}

	if err != nil {
		return err
	}
	return nil
}

/*
call ../01-simple-go-plugin/simple.wasm \
say_hello \
--input "Bob Morane" \
--log-level info \
--allow-hosts *,*.google.com,yo.com \
--config '{"firstName":"Philippe","lastName":"Charrière"}' \
--allow-paths '{"testdata":"./"}'
*/

func setCommonFlag(flagSet *flag.FlagSet) (*string, *string, *string, *string, *string, *string, *string, *string, *string) {
	handler := flagSet.String("handler", "handle", "wasm function name")
	wasmFile := flagSet.String("wasm", "*.wasm", "wasm file path (and name)")
	//fmt.Println("📝 file", *wasmFile)

	// Specific to Extism
	logLevel := flagSet.String("log-level", "", "Log level to print message")
	allowHosts := flagSet.String("allow-hosts", `["*"]`, "Hosts for HTTP request (json array)")
	allowPaths := flagSet.String("allow-paths", "{}", "Allowed paths to write and read files (json string)")
	config := flagSet.String("config", "{}", "Configuration data (json string)")

	authHeaderName := flagSet.String("auth-header-name", "", "ex: PRIVATE-TOKEN")
	authHeaderValue := flagSet.String("auth-header-value", "", "ex: IlovePandas")

	wasmUrl := flagSet.String("wasm-url", "", "url to download the wasm file")

	return handler, wasmFile, logLevel, allowHosts, allowPaths, config, wasmUrl, authHeaderName, authHeaderValue
}

func parseCommand(command string, args []string) error {

	switch command {
	case "start", "listen":

		flagSet := flag.NewFlagSet("listen", flag.ExitOnError)
		httpPort := flagSet.String("http-port", "8080", "http port")

		handler,
			wasmFile,
			logLevel,
			allowHosts,
			allowPaths,
			config,
			wasmUrl,
			authHeaderName,
			authHeaderValue := setCommonFlag(flagSet)

		flagSet.Parse(args)

		if *wasmUrl != "" { // we need to download the wasm file
			fmt.Println("🌍 downloading...", *wasmUrl)
			err := downloadWasmFile(*wasmUrl, *wasmFile, *authHeaderName, *authHeaderValue)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Println("🌍 http-port:", *httpPort)
		fmt.Println("🚀 handler  :", *handler)
		fmt.Println("📦 wasm     :", *wasmFile)

		cmds.Start(*wasmFile, *handler, *httpPort, *logLevel, *allowHosts, *allowPaths, *config)

		return nil

	case "redis":
		//redisCmd := flag.Args()[0]
		subCommand := flag.Args()[1]

		flagSet := flag.NewFlagSet("redis", flag.ExitOnError)

		// Allow to use an existing redis client
		redisUri := flagSet.String("uri", "rediss://default:pwd@redis-something:port", "Redis URI")
		redisClientId := flagSet.String("client-id", "something", "Redis client id")
		maskVariables := flagSet.Bool("mask-variables", true, "")

		handler,
			wasmFile,
			logLevel,
			allowHosts,
			allowPaths,
			config,
			wasmUrl,
			authHeaderName,
			authHeaderValue := setCommonFlag(flagSet)

		switch subCommand {
		case "subscribe":

			redisChannel := flagSet.String("channel", "knock-knock", "Redis channel")

			flagSet.Parse(args[1:]) // from 1, because of the subCommand

			if *wasmUrl != "" { // we need to download the wasm file
				fmt.Println("🌍 downloading...", *wasmUrl)
				err := downloadWasmFile(*wasmUrl, *wasmFile, *authHeaderName, *authHeaderValue)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			if *maskVariables != true {
				fmt.Println("🌍 redis URI      :", *redisUri)
			} else {
				fmt.Println("🌍 redis URI      :", "*****")
			}

			fmt.Println("🌍 redis Client Id:", *redisClientId)
			fmt.Println("🚀 handler        :", *handler)
			fmt.Println("📦 wasm           :", *wasmFile)
			fmt.Println("📺 channel        :", *redisChannel)

			cmds.RedisSubscribe(*wasmFile, *handler, *redisChannel, *redisUri, *redisClientId, *logLevel, *allowHosts, *allowPaths, *config)

		case "memdb":

			// could be use for a mono redis connection

		default:
			return fmt.Errorf("🔴 invalid subcommand: '%s'\n\n%s\n", subCommand, infos.Usage)
		}

		return nil

	case "nats":
		subCommand := flag.Args()[1]

		flagSet := flag.NewFlagSet("nats", flag.ExitOnError)

		// Allow to use an existing NATS client
		natsUrl := flagSet.String("url", "nats://somebody:secretpassword@demo.nats.io:4222", "Nats URL")
		natsConnectionId := flagSet.String("connection-id", "something", "NATS connection id")

		maskVariables := flagSet.Bool("mask-variables", true, "")

		handler,
			wasmFile,
			logLevel,
			allowHosts,
			allowPaths,
			config,
			wasmUrl,
			authHeaderName,
			authHeaderValue := setCommonFlag(flagSet)

		switch subCommand {
		case "subscribe":

			natsSubject := flagSet.String("subject", "knock-knock", "NATS subject")
			flagSet.Parse(args[1:]) // from 1, because of the subCommand

			if *wasmUrl != "" { // we need to download the wasm file
				fmt.Println("🌍 downloading...", *wasmUrl)
				err := downloadWasmFile(*wasmUrl, *wasmFile, *authHeaderName, *authHeaderValue)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			if *maskVariables != true {
				fmt.Println("🌍 NATS URL      :", *natsUrl)
			} else {
				fmt.Println("🌍 NATS URL      :", "*****")
			}

			fmt.Println("🌍 NATS Connection Id:", *natsConnectionId)
			fmt.Println("🚀 handler           :", *handler)
			fmt.Println("📦 wasm              :", *wasmFile)
			fmt.Println("📺 Subject           :", *natsSubject)

			cmds.NatsSubscribe(*wasmFile, *handler, *natsSubject, *natsUrl, *natsConnectionId, *logLevel, *allowHosts, *allowPaths, *config)

		default:
			return fmt.Errorf("🔴 invalid subcommand: '%s'\n\n%s\n", subCommand, infos.Usage)
		}

		return nil

	case "cli", "run":
		flagSet := flag.NewFlagSet("run", flag.ExitOnError)

		input := flagSet.String("input", "hello", "input data for the wasm plugin")

		handler,
			wasmFile,
			logLevel,
			allowHosts,
			allowPaths,
			config,
			wasmUrl,
			authHeaderName,
			authHeaderValue := setCommonFlag(flagSet)

		flagSet.Parse(args)

		if *wasmUrl != "" { // we need to download the wasm file
			fmt.Println("🌍 downloading...", *wasmUrl)
			err := downloadWasmFile(*wasmUrl, *wasmFile, *authHeaderName, *authHeaderValue)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		cmds.Execute(*wasmFile, *handler, *input, *logLevel, *allowHosts, *allowPaths, *config)

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
