# Write a NATS publisher plug-in

> We are going to use it to publish message to the [NATS subscriber](write-nats-subscriber.md)

=== "Go"
    ```golang linenums="1"
    package main

    import (
        "errors"

        "github.com/extism/go-pdk"
        "github.com/valyala/fastjson"
    )

    //export hostPrint
    func hostPrint(offset uint64) uint64

    func Print(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrint(memoryText.Offset())
    }

    //export hostGetEnv
    func hostGetEnv(offset uint64) uint64

    func GetEnv(name string) string {
        // copy the name of the environment variable to the shared memory
        variableName := pdk.AllocateString(name)
        // call the host function
        offset := hostGetEnv(variableName.Offset())

        // read the value of the result from the shared memory
        variableValue := pdk.FindMemory(offset)
        buffer := make([]byte, variableValue.Length())
        variableValue.Load(buffer)

        // cast the buffer to string and return the value
        envVarValue := string(buffer)
        return envVarValue
    }

    var parser = fastjson.Parser{}

    //export hostInitNatsConnection
    func hostInitNatsConnection(offset uint64) uint64

    func InitNatsConnection(natsConnectionId string, natsUrl string) (string, error) {
        jsonStrArguments := `{"id":"` + natsConnectionId + `","url":"` + natsUrl + `"}`
        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostInitNatsConnection(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostInitNatsConnection) returns a JSON buffer:
        // {
        //   "success": "the NATS connexion id",
        //   "failure": "error message if error, else empty"
        // }
        memoryResult := pdk.FindMemory(offset)
        buffer := make([]byte, memoryResult.Length())
        memoryResult.Load(buffer)

        JSONData, err := parser.ParseBytes(buffer)
        if err != nil {
            return "", err
        }
        if len(JSONData.GetStringBytes("failure")) == 0 {
            return string(JSONData.GetStringBytes("success")), nil
        } else {
            return "", errors.New(string(JSONData.GetStringBytes("failure")))
        }

    }


    //export hostNatsPublish
    func hostNatsPublish(offset uint64) uint64

    func NatsPublish(natsServer string, subject string, data string) (string, error) {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the NATS client",
        //    "subject": "name",
        //    "data": "Bob Morane"
        // }
        //jsonStr := `{"url":"` + natsServer + `","subject":"` + subject + `","data":"` + data + `"}`
        jsonStr := `{"id":"` + natsServer + `","subject":"` + subject + `","data":"` + data + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStr)

        // Call host function with the offset of the arguments
        offset := hostNatsPublish(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostNatsPublish) returns a JSON buffer:
        // {
        //   "success": "the message",
        //   "failure": "error message if error, else empty"
        // }
        memoryResult := pdk.FindMemory(offset)
        buffResult := make([]byte, memoryResult.Length())
        memoryResult.Load(buffResult)

        JSONData, err := parser.ParseBytes(buffResult)

        if err != nil {
            return "", err
        }
        if len(JSONData.GetStringBytes("failure")) == 0 {
            return string(JSONData.GetStringBytes("success")), nil
        } else {
            return "", errors.New(string(JSONData.GetStringBytes("failure")))
        }
    }

    //export publish
    func publish() uint64 {
        input := pdk.Input()

        natsURL := GetEnv("NATS_URL")
        Print("💜 NATS_URL: " + natsURL)
        idNatsConnection, errInit := InitNatsConnection("natsconn01", natsURL)
        if errInit != nil {
            Print("😡 " + errInit.Error())
        } else {
            Print("🙂 " + idNatsConnection)
        }

        res, err := NatsPublish("natsconn01", "news", string(input))

        if err != nil {
            Print("😡 " + err.Error())
        } else {
            Print("🙂 " + res)
        }
        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1"
    use extism_pdk::*;
    use serde::{Serialize, Deserialize};
    use thiserror::Error;

    extern "C" {
        fn hostPrint(ptr: u64) -> u64;
    }

    pub fn print(text: String) {
        let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
        memory_text.store(text);
        unsafe { hostPrint(memory_text.offset) };
    }

    extern "C" {
        fn hostGetEnv(ptr: u64) -> u64;
    }

    pub fn get_env(name: String) -> String {
        // copy the name of the environment variable to the shared memory
        let mut variable_name: Memory = extism_pdk::Memory::new(name.len());
        variable_name.store(name);

        // call the host function
        let offset: u64 = unsafe { hostGetEnv(variable_name.offset) };

        // read the value of the result from the shared memory
        let variable_value: Memory = extism_pdk::Memory::find(offset).unwrap();

        // return the value
        return variable_value.to_string().unwrap()
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct NatsConfig {
        pub id: String,
        pub url: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct NatsMessage {
        pub id: String,
        pub subject: String,
        pub data: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum NatsError {
        #[error("Nats Connection issue")]
        ConnectionFailure,
        #[error("Store issue")]
        MessageFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostInitNatsConnection(offset: u64) -> u64;
    }

    pub fn init_nats_connection(nats_connection_id: String, nats_url: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the NATS connection",
        //    "url": "URL of the NATS server"
        // }
        let args = NatsConfig {
            id: nats_connection_id,
            url: nats_url,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostInitNatsConnection(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostInitNatsConnection) returns a JSON buffer:
        // {
        //   "success": "id of the connection",
        //   "failure": "error message if error, else empty"
        // }
        let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
        
        let json_string:String = memory_result.to_string().unwrap();
        let result: StringResult = serde_json::from_str(&json_string).unwrap();

        if result.failure.is_empty()  {
            return Ok(result.success);
        } else {
            return Err(NatsError::ConnectionFailure.into());
        }
    }

    extern "C" {
        fn hostNatsPublish(offset: u64) -> u64;
    }

    pub fn nats_publish(nats_connection_id: String, subject: String, data: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the NATS client",
        //    "subject": "name",
        //    "data": "Bob Morane"
        // }
        let args = NatsMessage {
            id: nats_connection_id, 
            subject: subject,
            data: data,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostNatsPublish(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function returns a JSON buffer:
        // {
        //   "success": "OK",
        //   "failure": "error message if error, else empty"
        // }
        let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
        
        let json_string:String = memory_result.to_string().unwrap();
        let result: StringResult = serde_json::from_str(&json_string).unwrap();

        if result.failure.is_empty()  {
            return Ok(result.success);
        } else {
            return Err(NatsError::MessageFailure.into());
        }


    }

    #[plugin_fn]
    pub fn publish(input: String) -> FnResult<u64> {

        let nats_url : String = get_env("NATS_URL".to_string());
        let nats_connection : Result<String, Error> = init_nats_connection("natsconn01".to_string(), nats_url);

        match nats_connection {
            Ok(value) => print("🦀 nats connection: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        match nats_publish("natsconn01".to_string(), "news".to_string(), input.to_string()) {
            Ok(value)  => print("🦀 🙂 ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }
        
        Ok(0)
    }
    ```

## Build

=== "Go"
    ```golang linenums="1"
    tinygo build -scheduler=none --no-debug \
        -o natspub.wasm \
        -target wasi main.go
    ```

=== "Rust"
    ```rust linenums="1"
    cargo clean
    cargo build --release --target wasm32-wasi #--offline
    ls -lh ./target/wasm32-wasi/release/*.wasm
    cp ./target/wasm32-wasi/release/*.wasm .
    ```

## Run the plug-in to publish a message

```bash linenums="1"
export NATS_URL="nats://0.0.0.0:4222"
./slingshot run --wasm=./natspub.wasm \
--handler=publish \
--input="I 💜 Wasm ✨"
```

On the subscriber side, you shoul read:
```bash linenums="1"
👋 message: {"subject":"news","data":"I 💜 Wasm ✨"}
```