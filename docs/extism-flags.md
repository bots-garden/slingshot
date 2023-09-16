# Extism flags

The Slingshot runner can use these additional flags:

- `--log-level`
- `--allow-hosts`
- `--config`
- `--allow-paths`

These flags allow to use some specific Extism **PDKs** API.

## `--log-level`
> Possible values: error, warn, info, debug, trace

Usage:
```bash
./slingshot run \
    --wasm=./print.wasm \
    --handler=callHandler \
    --input="🤓 I'm a geek" \
    --log-level info 
```

Plugin source code:
=== "Go"
    ```golang
    pdk.Log(pdk.LogInfo, "😀😃😄")
    ```


## `--allow-hosts`

Usage:
```bash
./slingshot run \
    --wasm=./print.wasm \
    --handler=callHandler \
    --input="🤓 I'm a geek" \
    --allow-hosts *,*.google.com,yo.com
```

Plugin source code:
=== "Go"
    ```golang
    req := pdk.NewHTTPRequest("GET", "https://jsonplaceholder.typicode.com/todos/1")
    ```

## `--config`

Usage:
```bash
./slingshot run \
    --wasm=./print.wasm \
    --handler=callHandler \
    --input="🤓 I'm a geek" \
    --config '{"firstName":"Borane","lastName":"Morane"}'
```

Plugin source code:
=== "Go"
    ```golang
	firstName, _ := pdk.GetConfig("firstName")
	lastName, _ := pdk.GetConfig("lastName")

	pdk.Log(pdk.LogInfo, firstName)
	pdk.Log(pdk.LogInfo, lastName)
    ```


## `--allow-paths`

> No implemented yet