# Write and run a plug-in
> Official documentation: [https://extism.org/docs/category/write-a-plug-in](https://extism.org/docs/category/write-a-plug-in)
=== "Go"
    ```golang linenums="1"
    package main

    import (
        "github.com/extism/go-pdk"
    )

    //export hello
    func hello() {
        // read function argument from the shared memory
        input := pdk.Input()
        output := "👋 Hello " + string(input)

        // copy output to shared memory
        mem := pdk.AllocateString(output)
        pdk.OutputMemory(mem)
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1"
    use extism_pdk::*;

    #[plugin_fn]
    pub fn hello(input: String) -> FnResult<String> {

        let output : String = "👋 Hello ".to_string() + &input;
        
        Ok(output)
    }
    ```

## Build

=== "Go"
    ```bash linenums="1"
    #!/bin/bash
    tinygo build -scheduler=none --no-debug \
    -o hello.wasm \
    -target wasi main.go
    ```
=== "Rust"
    ```bash linenums="1"
    #!/bin/bash
    cargo clean
    cargo build --release --target wasm32-wasi #--offline
    ls -lh ./target/wasm32-wasi/release/*.wasm
    cp ./target/wasm32-wasi/release/*.wasm .
    ```

## Run

```bash linenums="1"
#!/bin/bash
./slingshot cli --wasm=./hello.wasm --handler=hello --input="Bob 🤓"
```

### Output

```bash linenums="1"
👋 Hello Bob 🤓
```