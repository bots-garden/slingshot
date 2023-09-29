# 🛠️ Host functions

!!! info "Similar helpers are already provided by the Extism PDK(s) and are more complete"
    
    - **Go**: [https://github.com/extism/go-pdk/blob/main/extism_pdk.go](https://github.com/extism/go-pdk/blob/main/extism_pdk.go)
    - **Rust**: [https://github.com/extism/rust-pdk/blob/main/src/var.rs](https://github.com/extism/rust-pdk/blob/main/src/var.rs)

    `hostMemorySet` and `hostMemoryGet` have been developed to validate our understanding of how the **Extism PDKs** work.

## hostMemorySet

**`hostMemorySet`**: store a key/value into a memory map.

=== "Go"
    ```golang linenums="1" hl_lines="19-57"
    package main

    import (
        "errors"
        "github.com/extism/go-pdk"
        "github.com/valyala/fastjson"
    )

    //export hostPrintln
    func hostPrintln(offset uint64) uint64

    func Println(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrintln(memoryText.Offset())
    }

    var parser = fastjson.Parser{}

    //export hostMemorySet
    func hostMemorySet(offset uint64) uint64

    func MemorySet(key string, value string) (string, error) {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "key": "name",
        //    "value": "Bob Morane"
        // }
        jsonStr := `{"key":"` + key + `","value":"` + value + `"}`

        // Copy the string value to the shared memory
        keyAndValue := pdk.AllocateString(jsonStr)

        // Call host function with the offset of the arguments
        offset := hostMemorySet(keyAndValue.Offset())

        // Get result from the shared memory
        // The host function (hostMemorySet) returns a JSON buffer:
        // {
        //   "success": "the value associated to the key",
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

    //export hello
    func hello() uint64 {

        _, err := MemorySet("bob", "Bob Morane")
        
        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="35-75"
    use extism_pdk::*;
    use serde::{Serialize, Deserialize};
    use thiserror::Error;

    extern "C" {
        fn hostPrintln(ptr: u64) -> u64;
    }

    pub fn println(text: String) {
        let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
        memory_text.store(text);
        unsafe { hostPrintln(memory_text.offset) };
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct MemArguments {
        pub key: String,
        pub value: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum MemError {
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostMemorySet(offset: u64) -> u64;
    }

    pub fn memory_set(key: String, value: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "key": "name",
        //    "value": "Bob Morane"
        // }
        let record = MemArguments {
            key: key,
            value: value,
        };
        let json_str: String = serde_json::to_string(&record).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostMemorySet(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostMemorySet) returns a JSON buffer:
        // {
        //   "success": "the value associated to the key",
        //   "failure": "error message if error, else empty"
        // }
        let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
        
        let json_string:String = memory_result.to_string().unwrap();
        let result: StringResult = serde_json::from_str(&json_string).unwrap();

        if result.failure.is_empty()  {
            return Ok(result.success);
        } else {
            return Err(MemError::StoreFailure.into());
        }
    }


    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        match memory_set("bob".to_string(), "Bob Morane".to_string()) {
            Ok(value)  => println("🦀 saved value: ".to_string() + &value),
            Err(error) => println("😡 error: ".to_string() + &error.to_string()),
        }
        
        Ok(0)
    }
    ```

## hostMemoryGet

**`hostMemoryGet`**: get value with a key from a memory map.

=== "Go"
    ```golang linenums="1" hl_lines="19-49"
    package main

    import (
        "errors"
        "github.com/extism/go-pdk"
        "github.com/valyala/fastjson"
    )

    //export hostPrintln
    func hostPrintln(offset uint64) uint64

    func Println(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrintln(memoryText.Offset())
    }

    var parser = fastjson.Parser{}

    //export hostMemoryGet
    func hostMemoryGet(offset uint64) uint64

    func MemoryGet(key string) (string, error) {
        // Copy argument to memory
        memoryKey := pdk.AllocateString(key)
        // Call the host function
        offset := hostMemoryGet(memoryKey.Offset())

        // Get result (the value associated to the key) from shared memory
        // The host function (hostMemoryGet) returns a JSON buffer:
        // {
        //   "success": "the value associated to the key",
        //   "failure": "error message if error, else empty"
        // }
        memoryValue := pdk.FindMemory(offset)
        buffer := make([]byte, memoryValue.Length())
        memoryValue.Load(buffer)

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

    //export hello
    func hello() uint64 {

        value, err := MemoryGet("bob")
        if err != nil {
            Println("😡 ouch! " + err.Error())
        } else {
            Println("value: " + value)
        }
        
        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="29-56"
    use extism_pdk::*;
    use serde::{Serialize, Deserialize};
    use thiserror::Error;

    extern "C" {
        fn hostPrintln(ptr: u64) -> u64;
    }

    pub fn println(text: String) {
        let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
        memory_text.store(text);
        unsafe { hostPrintln(memory_text.offset) };
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum MemError {
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostMemoryGet(offset: u64) -> u64;
    }

    pub fn memory_get(key: String) -> Result<String, Error> {
        // Copy argument to memory
        let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
        memory_key.store(key);

        // Call the host function
        let offset: u64 = unsafe { hostMemoryGet(memory_key.offset) };

        // Get result (the value associated to the key) from shared memory
        // The host function (hostMemoryGet) returns a JSON buffer:
        // {
        //   "success": "the value associated to the key",
        //   "failure": "error message if error, else empty"
        // }
        let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
        let json_string:String = memory_result.to_string().unwrap();
        let result: StringResult = serde_json::from_str(&json_string).unwrap();
        
        if result.failure.is_empty()  {
            return Ok(result.success);
        } else {
            return Err(MemError::NotFound.into());
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        match memory_get("bob".to_string()) {
            Ok(value)  => println("🦀 value: ".to_string() + &value),
            Err(error) => println("😡 error: ".to_string() + &error.to_string()),
        }

        match memory_get("sam".to_string()) {
            Ok(value)  => println("🦀 value: ".to_string() + &value),
            Err(error) => println("😡 error: ".to_string() + &error.to_string()),
        }
        
        Ok(0)
    }
    ```

