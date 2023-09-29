# Write a Redis subscriber plug-in

=== "Go"
    ```golang linenums="1"
    package main

    import (
        "github.com/extism/go-pdk"
    )

    //export hostPrintln
    func hostPrintln(offset uint64) uint64

    func Println(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrintln(memoryText.Offset())
    }

    //export message
    func message() uint64 {
        // read function argument from the memory
        input := pdk.Input()

        Println("👋 message: " + string(input))
        
        return 0
    }

    func main() {}
    ```

## Build

=== "Go"
    ```bash linenums="1"
    #!/bin/bash
    tinygo build -scheduler=none --no-debug \
        -o redissub.wasm \
        -target wasi main.go
    ```
    
## Run

```bash linenums="1"
#!/bin/bash

./slingshot redis subscribe \
--wasm=./redissub.wasm \
--handler=message \
--uri=${REDIS_URI} \
--client-id=pubsubcli \
--channel=news

# Output:
🌍 redis URI      : *****
🌍 redis Client Id: pubsubcli
🚀 handler        : message
📦 wasm           : ./redissub.wasm
📺 channel        : news
```

### Trigger the plugin

Connect to the Redis server:
```bash linenums="1"
#!/bin/bash
redis-cli -u  ${REDIS_URI}
```

Publish message(s):
```bash linenums="1"
redis.aivencloud.com:17170> PUBLISH news "Hello World"
```

### Output

```bash linenums="1"
👋 message: {"id":"pubsubcli","channel":"news","payload":"Hello World"}
```
