# SlingShot

**SlingShot** is a **Wasm** runner to run or serve **[Extism](https://extism.org/)** **Wasm** plug-ins

```bash title="Run a wasm plug-in"
./slingshot run --wasm=./hello.wasm --handler=hello --input="Bob 🤓"
```

```bash title="Serve a wasm plug-in as a function"
./slingshot listen --wasm=./hello.wasm --handler=handle --http-port=7070
```

```bash title="Trigger a wasm plug-in with Redis messages"
./slingshot redis subscribe --wasm=./hello.wasm --handler=message \
--uri=${REDIS_URI} \
--client-id=007 \
--channel=news
```

```bash title="Trigger a wasm plug-in with NATS messages (✋ experimental 🚧 WIP)"
./slingshot nats subscribe --wasm=./hello.wasm --handler=message \
--url=${NATS_SERVER_URL} \
--connection-id=007 \
--subject=news
```

```bash title="Execute a remote wasm file"
./slingshot run \
--wasm-url="http://0.0.0.0:9000/print.wasm" \
--wasm=./print.wasm \
--handler=callHandler \
--input="🤓 I'm a geek"
```

```bash title="Run Slingshot with Docker (multi arch image 🐳)"
HTTP_PORT=8080
docker run \
    -p ${HTTP_PORT}:${HTTP_PORT} \
    -v $(pwd)/hello-service:/app --rm botsgarden/slingshot:0.0.5 \
    /slingshot start \
    --wasm=./app/hello.wasm \
    --handler=callHandler \
    --http-port=${HTTP_PORT} 
```

## How is Slingshot developed?

Slingshot is developed in Go with **[Wazero](https://wazero.io/)**[^1] as the Wasm runtime and **[Extism](https://extism.org/)**[^2], which offers a Wazero-based Go SDK and a Wasm plugin system.

[^1]: Wazero is a project from **[Tetrate](https://tetrate.io/)**
[^2]: Extism is a project from **[Dylibso](https://dylibso.com/)**

## Install SlingShot

- Download the latest release of SlingShot: [https://github.com/bots-garden/slingshot/releases](https://github.com/bots-garden/slingshot/releases) for your machine and OS.
- Rename it to `slingshot`
- Check by typing: `./slingshot version`

## Getting started
### 🚀 Write Extism plug-ins

- [Write & run a plug-in](write-plug-in.md)
- [Write & serve a plug-in as a nano-service](write-service.md)
- [Write a Redis subscriber plug-in](write-redis-subscriber.md)
- [Write a Redis publisher plug-in](write-redis-publisher.md)
- [Write a NATS subscriber plug-in](write-nats-subscriber.md)
- [Write a NATS publisher plug-in](write-nats-publisher.md)
