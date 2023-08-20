# Write a NATS subscriber plug-in

> You need to install a NATS server: [https://docs.nats.io/running-a-nats-service/introduction/installation](https://docs.nats.io/running-a-nats-service/introduction/installation)

=== "Go"
    ```golang linenums="1"
    package main

    import (
        "github.com/extism/go-pdk"
    )

    //export hostPrint
    func hostPrint(offset uint64) uint64

    func Print(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrint(memoryText.Offset())
    }

    //export message
    func message() uint64 {
        // read function argument from the memory
        input := pdk.Input()

        Print("👋 message: " + string(input))
        
        return 0
    }

    func main() {}

    ```

=== "Rust"
    ```rust linenums="1"
    use extism_pdk::*;

    extern "C" {
        fn hostPrint(ptr: u64) -> u64;
    }

    pub fn print(text: String) {
        let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
        memory_text.store(text);
        unsafe { hostPrint(memory_text.offset) };
    }

    #[plugin_fn]
    pub fn message(input: String) -> FnResult<u64> {

        print("👋 message: ".to_string() + &input);
        
        Ok(0)
    }
    ```

## Build

=== "Go"
    ```bash linenums="1"
    #!/bin/bash
    tinygo build -scheduler=none --no-debug \
        -o natssub.wasm \
        -target wasi main.go
    ```

=== "Rust"
    ```bash linenums="1"
    #!/bin/bash
    cargo clean
    cargo build --release --target wasm32-wasi
    ls -lh ./target/wasm32-wasi/release/*.wasm
    cp ./target/wasm32-wasi/release/*.wasm .
    ```

## Run

```bash linenums="1"
#!/bin/bash

./slingshot nats subscribe \
--wasm=./natssub.wasm \
--handler=message \
--url=nats://0.0.0.0:4222 \
--connection-id=natsconn01 \
--subject=news

# Output:
🌍 NATS URL      : *****
🌍 NATS Connection Id: natsconn01
🚀 handler           : message
📦 wasm              : ./natssub.wasm
📺 Subject           : news
```

### Trigger the plugin

You need a NATS client. It's easy to write one with Go:
```golang linenums="1"
package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	//nc, err := nats.Connect("nats://0.0.0.0:4222")
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer nc.Close()

	err = nc.Publish("news", []byte("Hello World"))

	if err != nil {
		fmt.Println(err.Error())
	}
}

```

Publish message(s):
```bash linenums="1"
go run main.go
```

### Output

```bash linenums="1"
👋 message: {"id":"natscli","subject":"news","data":"Hello World"}
```
