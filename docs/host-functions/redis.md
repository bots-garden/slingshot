# 🛠️ Host functions

## hostInitRedisClient

**`hostInitRedisClient`**: initialize a Redis client it it does not exist.

=== "Go"
    ```go linenums="1" hl_lines="40-78"
    package main

    import (
        "errors"
        "strings"

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

    //export hostInitRedisClient
    func hostInitRedisClient(offset uint64) uint64

    func InitRedisClient(redisClientId string, redisUri string) (string, error) {

        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        jsonStrArguments := `{"id":"` + redisClientId + `","uri":"` + redisUri + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostInitRedisClient(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
        // {
        //   "success": "the redis client id",
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

    //export hello
    func hello() uint64 {
        redisURI := GetEnv("REDIS_URI")
        idRedisClient, errInit := InitRedisClient("redisDb", redisURI)
        if errInit!= nil {
            Print("😡 " + errInit.Error())
        } else {
            Print("🙂 " + idRedisClient)
        }

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="56-96"
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
    struct RedisClientArguments {
        pub id: String,
        pub uri: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum RedisError {
        #[error("Redis Client issue")]
        ClientFailure,
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostInitRedisClient(offset: u64) -> u64;
    }

    pub fn init_redis_client(redis_client_id: String, redis_uri: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        let args = RedisClientArguments {
            id: redis_client_id,
            uri: redis_uri,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostInitRedisClient(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
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
            return Err(RedisError::ClientFailure.into());
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        let redis_uri : String = get_env("REDIS_URI".to_string());
        let redis_client : Result<String, Error> = init_redis_client("redisDb".to_string(), redis_uri);

        match redis_client {
            Ok(value) => print("🦀 redis client: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        Ok(0)
    }
    ```

## hostRedisSet

**`hostRedisSet`**: store a value with a key into the Redis database.

=== "Go"
    ```go linenums="1" hl_lines="81-120"
    package main

    import (
        "errors"
        "strings"

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

    //export hostInitRedisClient
    func hostInitRedisClient(offset uint64) uint64

    func InitRedisClient(redisClientId string, redisUri string) (string, error) {

        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        jsonStrArguments := `{"id":"` + redisClientId + `","uri":"` + redisUri + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostInitRedisClient(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
        // {
        //   "success": "the redis client id",
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

    //export hostRedisSet
    func hostRedisSet(offset uint64) uint64

    func RedisSet(redisClientId string, key string, value string) (string, error) {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name",
        //    "value": "Bob Morane"
        // }
        jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `","value":"` + value + `"}`
        
        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStr)

        // Call host function with the offset of the arguments
        offset := hostRedisSet(arguments.Offset())

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
        redisURI := GetEnv("REDIS_URI")
        InitRedisClient("redisDb", redisURI)

        k1, errSet1 := RedisSet("redisDb", "001", "Huey")
        k2, errSet2 := RedisSet("redisDb", "002", "Dewey")
        k3, errSet3 := RedisSet("redisDb", "003", "Louie")

        allSetErrs := errors.Join(errSet1, errSet2, errSet3) 
        if allSetErrs != nil {
            Print("😡 " + allSetErrs.Error())
        } else {
            Print("🙂 " + strings.Join([]string{k1,k2,k3}, ","))
        }

        /* output:
            🙂 001,002,003
        */

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="105-147"
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
    struct RedisClientArguments {
        pub id: String,
        pub uri: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct RedisArguments {
        pub id: String,
        pub key: String,
        pub value: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum RedisError {
        #[error("Redis Client issue")]
        ClientFailure,
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostInitRedisClient(offset: u64) -> u64;
    }

    pub fn init_redis_client(redis_client_id: String, redis_uri: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        let args = RedisClientArguments {
            id: redis_client_id,
            uri: redis_uri,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostInitRedisClient(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
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
            return Err(RedisError::ClientFailure.into());
        }
    }

    extern "C" {
        fn hostRedisSet(offset: u64) -> u64;
    }

    pub fn redis_set(redis_client_id: String, key: String, value: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name",
        //    "value": "Bob Morane"
        // }
        let args = RedisArguments {
            id: redis_client_id, 
            key: key,
            value: value,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostRedisSet(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostRedisSet) returns a JSON buffer:
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
            return Err(RedisError::StoreFailure.into());
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        let redis_uri : String = get_env("REDIS_URI".to_string());
        let redis_client : Result<String, Error> = init_redis_client("redisDb".to_string(), redis_uri);

        match redis_set("redisDb".to_string(), "100".to_string(), "Huey".to_string()) {
            Ok(value)  => print("🦀 saved value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }
        match redis_set("redisDb".to_string(), "200".to_string(), "Dewey".to_string()) {
            Ok(value)  => print("🦀 saved value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }
        match redis_set("redisDb".to_string(), "300".to_string(), "Louie".to_string()) {
            Ok(value)  => print("🦀 saved value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        Ok(0)
    }
    ```

## hostRedisGet

**`hostRedisGet`**: retrieve a value with a key from the Redis database.

=== "Go"
    ```go linenums="1" hl_lines="81-119"
    package main

    import (
        "errors"
        "strings"

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

    //export hostInitRedisClient
    func hostInitRedisClient(offset uint64) uint64

    func InitRedisClient(redisClientId string, redisUri string) (string, error) {

        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        jsonStrArguments := `{"id":"` + redisClientId + `","uri":"` + redisUri + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostInitRedisClient(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
        // {
        //   "success": "the redis client id",
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

    //export hostRedisGet
    func hostRedisGet(offset uint64) uint64

    func RedisGet(redisClientId string, key string) (string, error) {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name"
        // }
        jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStr)
        // Call the host function
        offset := hostRedisGet(arguments.Offset())

        // Get result (the value associated to the key) from shared memory
        // The host function (hostRedisGet) returns a JSON buffer:
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
        redisURI := GetEnv("REDIS_URI")
        InitRedisClient("redisDb", redisURI)

        v1, errGet1 := RedisGet("redisDb", "001")
        v2, errGet2 := RedisGet("redisDb", "002")
        v3, errGet3 := RedisGet("redisDb", "003")

        allGetErrs := errors.Join(errGet1, errGet2, errGet3) 
        if allGetErrs != nil {
            Print("😡 " + allSetErrs.Error())
        } else {
            Print("🙂 " + strings.Join([]string{v1,v2,v3}, ","))
        }

        /* output:
            🙂 Huey,Dewey,Louie
        */

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="105-146"
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
    struct RedisClientArguments {
        pub id: String,
        pub uri: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct RedisArguments {
        pub id: String,
        pub key: String,
        pub value: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum RedisError {
        #[error("Redis Client issue")]
        ClientFailure,
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostInitRedisClient(offset: u64) -> u64;
    }

    pub fn init_redis_client(redis_client_id: String, redis_uri: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        let args = RedisClientArguments {
            id: redis_client_id,
            uri: redis_uri,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostInitRedisClient(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
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
            return Err(RedisError::ClientFailure.into());
        }
    }

    extern "C" {
        fn hostRedisGet(offset: u64) -> u64;
    }

    pub fn redis_get(redis_client_id: String, key: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name",
        //    "value": ""
        // }
        let args = RedisArguments {
            id: redis_client_id, 
            key: key,
            value: String::new(),
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostRedisGet(memory_json_str.offset) };

        // Get result (the value associated to the key) from shared memory
        // The host function (hostRedisGet) returns a JSON buffer:
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
            return Err(RedisError::NotFound.into());
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        let redis_uri : String = get_env("REDIS_URI".to_string());
        let redis_client : Result<String, Error> = init_redis_client("redisDb".to_string(), redis_uri);

        match redis_get("redisDb".to_string(), "100".to_string()) {
            Ok(value)  => print("🦀 value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        match redis_get("redisDb".to_string(), "200".to_string()) {
            Ok(value)  => print("🦀 value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        match redis_get("redisDb".to_string(), "300".to_string()) {
            Ok(value)  => print("🦀 value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        match redis_get("redisDb".to_string(), "400".to_string()) {
            Ok(value)  => print("🦀 value: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        Ok(0)
    }
    ```

## hostRedisDel

**`hostRedisDel`**: delete a value with a key from the Redis database.

=== "Go"
    ```go linenums="1" hl_lines="81-119"
    package main

    import (
        "errors"
        "strings"

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

    //export hostInitRedisClient
    func hostInitRedisClient(offset uint64) uint64

    func InitRedisClient(redisClientId string, redisUri string) (string, error) {

        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        jsonStrArguments := `{"id":"` + redisClientId + `","uri":"` + redisUri + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostInitRedisClient(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
        // {
        //   "success": "the redis client id",
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

    //export hostRedisDel
    func hostRedisDel(offset uint64) uint64

    func RedisDel(redisClientId string, key string) (string, error) {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name"
        // }
        jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStr)
        // Call the host function
        offset := hostRedisDel(arguments.Offset())

        // Get result (the value associated to the key) from shared memory
        // The host function (hostRedisDel) returns a JSON buffer:
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
        redisURI := GetEnv("REDIS_URI")
        InitRedisClient("redisDb", redisURI)

        key, errDel := RedisDel("redisDb", "002")
        if errDel != nil {
            Print("😡 " + errDel.Error())
        } else {
            Print("🙂 " + key)
        }

        /* output:
            🙂 002
        */

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="105-146"
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
    struct RedisClientArguments {
        pub id: String,
        pub uri: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct RedisArguments {
        pub id: String,
        pub key: String,
        pub value: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum RedisError {
        #[error("Redis Client issue")]
        ClientFailure,
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostInitRedisClient(offset: u64) -> u64;
    }

    pub fn init_redis_client(redis_client_id: String, redis_uri: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        let args = RedisClientArguments {
            id: redis_client_id,
            uri: redis_uri,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostInitRedisClient(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
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
            return Err(RedisError::ClientFailure.into());
        }
    }

    extern "C" {
        fn hostRedisDel(offset: u64) -> u64;
    }

    pub fn redis_del(redis_client_id: String, key: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name",
        //    "value": ""
        // }
        let args = RedisArguments {
            id: redis_client_id, 
            key: key,
            value: String::new(),
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostRedisDel(memory_json_str.offset) };

        // Get result (the value associated to the key) from shared memory
        // The host function (hostRedisDel) returns a JSON buffer:
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
            return Err(RedisError::NotFound.into());
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        let redis_uri : String = get_env("REDIS_URI".to_string());
        let redis_client : Result<String, Error> = init_redis_client("redisDb".to_string(), redis_uri);

        match redis_del("redisDb".to_string(), "200".to_string()) {
            Ok(value)  => print("🦀 deleted key: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        Ok(0)
    }
    ```

## hostRedisFilter

**`hostRedisFilter`**: retrieve keys with a filter from the Redis database.

=== "Go"
    ```go linenums="1" hl_lines="81-119"
    package main

    import (
        "errors"
        "strings"

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

    //export hostInitRedisClient
    func hostInitRedisClient(offset uint64) uint64

    func InitRedisClient(redisClientId string, redisUri string) (string, error) {

        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        jsonStrArguments := `{"id":"` + redisClientId + `","uri":"` + redisUri + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostInitRedisClient(arguments.Offset())

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
        // {
        //   "success": "the redis client id",
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

    //export hostRedisFilter
    func hostRedisFilter(offset uint64) uint64

    func RedisFilter(redisClientId string, key string) (string, error) {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "filter ex 00*"
        // }
        jsonStr := `{"id":"` + redisClientId + `","key":"` + key + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStr)
        // Call the host function
        offset := hostRedisFilter(arguments.Offset())

        // Get result (the value associated to the key) from shared memory
        // The host function (hostRedisDel) returns a JSON buffer:
        // {
        //   "success": "array of keys",
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
        redisURI := GetEnv("REDIS_URI")
        InitRedisClient("redisDb", redisURI)

        keys, errKeys := RedisFilter("redisDb", "00*")
        if errKeys != nil {
            Print("😡 " + errKeys.Error())
        } else {
            Print("🙂 " + keys)
        }

        /* output:
            🙂 ["003","001"]
        */

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust linenums="1" hl_lines="105-146"
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
    struct RedisClientArguments {
        pub id: String,
        pub uri: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct RedisArguments {
        pub id: String,
        pub key: String,
        pub value: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum RedisError {
        #[error("Redis Client issue")]
        ClientFailure,
        #[error("Store issue")]
        StoreFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostInitRedisClient(offset: u64) -> u64;
    }

    pub fn init_redis_client(redis_client_id: String, redis_uri: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "uri": "redis uri"
        // }
        let args = RedisClientArguments {
            id: redis_client_id,
            uri: redis_uri,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostInitRedisClient(memory_json_str.offset) };

        // Get result from the shared memory
        // The host function (hostInitRedisClient) returns a JSON buffer:
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
            return Err(RedisError::ClientFailure.into());
        }
    }

    extern "C" {
        fn hostRedisFilter(offset: u64) -> u64;
    }

    pub fn redis_filter(redis_client_id: String, key: String) -> Result<String, Error> {
        // Prepare the arguments for the host function
        // with a JSON string:
        // {
        //    "id": "id of the redis client",
        //    "key": "name",
        //    "value": ""
        // }
        let args = RedisArguments {
            id: redis_client_id, 
            key: key,
            value: String::new(),
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostRedisFilter(memory_json_str.offset) };

        // Get result (the value associated to the key) from shared memory
        // The host function (hostRedisDel) returns a JSON buffer:
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
            return Err(RedisError::NotFound.into()); //???
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        let redis_uri : String = get_env("REDIS_URI".to_string());
        let redis_client : Result<String, Error> = init_redis_client("redisDb".to_string(), redis_uri);

        match redis_filter("redisDb".to_string(), "*00".to_string()) {
            Ok(value)  => print("🦀 keys: ".to_string() + &value),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        Ok(0)
    }
    ```

<!--
## hostMemoryGet

**`hostMemoryGet`**: get value with a key from a memory map.

=== "Go"
    ```golang linenums="1" hl_lines="15-32"

    ```

=== "Rust"
    ```rust linenums="1" hl_lines="13-30"

-->