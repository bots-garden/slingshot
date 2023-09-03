# Slingshot PDK
> Plug-in development kit

## Slingshot Go PDK

### Structure of a plugin

```golang
package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func helloHandler(argHandler []byte) []byte {
	input := string(argHandler)
	slingshot.Print("👋 hello world 🌍 " + string(input))
	
	return []byte("👋 Hello World 🌍")
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(helloHandler)
}

func main() {}
```

1. You must export a function (it will be called at the execution, here, it's `callHandler`).
2. `slingshot.ExecHandler` is an helper to execute the function (`helloHandler`). This helper will read the wasm shared memory to get the argument of the function (`argHandler []byte`).
3. Then the helper executes the function.
4. And finally write to the wasm shared memory to return the result of the function (`[]byte("👋 Hello World 🌍")`).

#### Build the Slingshot plugin

```bash
tinygo build -scheduler=none --no-debug \
    -o print.wasm \
    -target wasi main.go
```

#### Run the Slingshot plugin

```bash
./slingshot run --wasm=./print.wasm \
--handler=callHandler \
--input="🤓 I'm a geek"
```

### Slingshot host functions
> 🚧 This is a work in progress

- `slingshot.Print(text string)`
- `slingshot.Log(text string)`
- `slingshot.MemorySet(key string, value string) (string, error)`
- `slingshot.MemoryGet(key string) (string, error)`
- `slingshot.InitRedisClient(redisClientId string, redisUri string) (string, error)`
- `slingshot.RedisSet(redisClientId string, key string, value string) (string, error)`
- `slingshot.RedisGet(redisClientId string, key string) (string, error)`
- `slingshot.RedisDel(redisClientId string, key string) (string, error)`
- `slingshot.RedisFilter(redisClientId string, key string) (string, error)`
- `slingshot.RedisPublish(redisClientId string, channel string, payload string) (string, error)`
- `slingshot.InitNatsConnection(natsConnectionId string, natsUrl string) (string, error)`
- `slingshot.NatsPublish(natsConnectionId string, subject string, data string) (string, error)`

## Other examples

### HTTP nano-service plug-in

```golang
package main

import (
	"strings"
	slingshot "github.com/bots-garden/slingshot/go-pdk"
	"github.com/valyala/fastjson"
)

var parser = fastjson.Parser{}

func helloHandler(http_request_data []byte) []byte {

	var text string
	var code string

	JSONData, err := parser.ParseBytes(http_request_data)
	if err != nil {
		text = "😡 Error: " + err.Error()
		code = "500"
	} else {
		text = "🩵 Hello " + string(JSONData.GetStringBytes("body"))
		code = "200"
	}

	headers := []string{
		`"Content-Type": "application/json; charset=utf-8"`,
		`"X-Slingshot-version": "0.0.0"`,
	}

	headersStr := strings.Join(headers, ",")

	response := `{"headers":{` + headersStr + `}, "textBody": "` + text + `", "statusCode": ` + code + `}`

	return []byte(response)

}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(helloHandler)
}

func main() {}
```

#### Run

```bash
./slingshot listen --wasm=./hello.wasm \
--handler=callHandler \
--http-port=7070
```

### Redis subscriber plug-in

```golang
package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func messageHandler(input []byte) []byte {

	slingshot.Print("👋 message: " + string(input))
	return nil

}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(messageHandler)
}

func main() {}
```

#### Run

```bash
./slingshot redis subscribe \
    --wasm=./redissub.wasm \
    --handler=callHandler \
    --uri=${REDIS_URI} \
    --client-id=pubsubcli \
    --channel=news
```

### Redis publisher plug-in

```golang
package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func publishHandler(input []byte) []byte {

	redisURI := slingshot.GetEnv("REDIS_URI")
	idRedisClient, errInit := slingshot.InitRedisClient("pubsubcli", redisURI)
	if errInit != nil {
		slingshot.Print("😡 " + errInit.Error())
	} else {
		slingshot.Print("🙂 " + idRedisClient)
	}

	slingshot.RedisPublish("pubsubcli", "news", string(input))

	return nil
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(publishHandler)
}

func main() {}
```

#### Run

```bash
./slingshot run --wasm=./redispub.wasm \
    --handler=callHandler \
    --input="I 💜 Wasm ✨"
```

### Nats subscriber plug-in

```golang
package main

import slingshot "github.com/bots-garden/slingshot/go-pdk"

func messageHandler(input []byte) []byte {
	slingshot.Print("👋 NATS message: " + string(input))
	return nil
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(messageHandler)
}

func main() {}
```

#### Run

```bash
./slingshot nats subscribe \
--wasm=./natssub.wasm \
--handler=callHandler \
--url=nats://0.0.0.0:4222 \
--connection-id=natsconn01 \
--subject=news
```

### Nats publisher plug-in

```golang
package main

import (
	slingshot "github.com/bots-garden/slingshot/go-pdk"
)

func publishHandler(input []byte) []byte {

	natsURL := slingshot.GetEnv("NATS_URL")
	slingshot.Print("💜 NATS_URL: " + natsURL)
	idNatsConnection, errInit := slingshot.InitNatsConnection("natsconn01", natsURL)
	if errInit != nil {
		slingshot.Print("😡 " + errInit.Error())
	} else {
		slingshot.Print("🙂 " + idNatsConnection)
	}

	res, err := slingshot.NatsPublish("natsconn01", "news", string(input))

	if err != nil {
		slingshot.Print("😡 " + err.Error())
	} else {
		slingshot.Print("🙂 " + res)
	}
	return nil
}

//export callHandler
func callHandler() {
	slingshot.ExecHandler(publishHandler)
}

func main() {}
```

#### Run

```bash
./slingshot run --wasm=./natspub.wasm \
    --handler=callHandler \
    --input="I 💜 Wasm ✨"
```
