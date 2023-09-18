# SlingShot

**SlingShot** is a **Wasm** runner to run or serve **[Extism](https://extism.org/)** **Wasm** plug-ins

**Run a wasm plug-in**:
```bash title="Run a wasm plug-in"
./slingshot run --wasm=./hello.wasm --handler=hello --input="Bob 🤓"
```

**Serve a wasm plug-in as a function**:
```bash title="Serve a wasm plug-in as a function"
./slingshot listen --wasm=./hello.wasm --handler=handle --http-port=7070
```

**Trigger a wasm plug-in with Redis messages**:
```bash title="Trigger a wasm plug-in with Redis messages"
./slingshot redis subscribe --wasm=./hello.wasm --handler=message \
--uri=${REDIS_URI} \
--client-id=007 \
--channel=news
```

**title="Trigger a wasm plug-in with NATS messages (✋ experimental 🚧 WIP)**:
```bash title="Trigger a wasm plug-in with NATS messages (✋ experimental 🚧 WIP)"
./slingshot nats subscribe --wasm=./hello.wasm --handler=message \
--url=${NATS_SERVER_URL} \
--connection-id=007 \
--subject=news
```

**Execute a remote wasm file**:
```bash title="Execute a remote wasm file"
./slingshot run \
--wasm-url="http://0.0.0.0:9000/print.wasm" \
--wasm=./print.wasm \
--handler=callHandler \
--input="🤓 I'm a geek"
```

## How is Slingshot developed?

Slingshot is developed in Go with **[Wazero](https://wazero.io/)**[^1] as the Wasm runtime and **[Extism](https://extism.org/)**[^2], which offers a Wazero-based Go SDK and a Wasm plugin system.

[^1]: Wazero is a project from **[Tetrate](https://tetrate.io/)**
[^2]: Extism is a project from **[Dylibso](https://dylibso.com/)**

## Install SlingShot

- Download the latest release of SlingShot: [https://github.com/bots-garden/slingshot/releases](https://github.com/bots-garden/slingshot/releases) for your machine and OS.
- Rename it to `slingshot`
- Check by typing: `./slingshot version`
