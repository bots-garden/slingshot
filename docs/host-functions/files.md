# 🛠️ Host functions

## hostReadFile

**`hostReadFile`**: read the content of a file.

=== "Go"
    ```golang
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

    var parser = fastjson.Parser{}

    // GetJsonFromBytes
    // Convert a buffer (`[]byte`) into a JSON value
    func GetJsonFromBytes(buffer []byte) (*fastjson.Value, error) {
        return parser.ParseBytes(buffer)
    }

    //export hostReadFile
    func hostReadFile(offset uint64) uint64

    func ReadFile(filePath string) (string, error) {
        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(filePath)

        // Call host function with the offset of the arguments
        offset := hostReadFile(arguments.Offset())

        // Get result from the shared memory
        memoryResult := pdk.FindMemory(offset)
        buffResult := make([]byte, memoryResult.Length())
        memoryResult.Load(buffResult)
        JSONData, err := GetJsonFromBytes(buffResult)
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

        content, err := ReadFile("./hello.txt")
        if err != nil {
            Print("😡 " + err.Error())
        } else {
            Print(content)
        }

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust
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

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Error, Debug)]
    pub enum FileError {
        #[error("Read issue")]
        ReadFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostReadFile(ptr: u64) -> u64;
    }

    pub fn read_file(file_path: String) -> Result<String, Error> {
        // Copy the string value to the shared memory
        let mut memory_str: Memory = extism_pdk::Memory::new(file_path.len());
        memory_str.store(file_path);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostReadFile(memory_str.offset) };

        // Get result
        let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
        let json_string:String = memory_result.to_string().unwrap();
        let result: StringResult = serde_json::from_str(&json_string).unwrap();
        
        if result.failure.is_empty()  {
            return Ok(result.success);
        } else {
            return Err(FileError::ReadFailure.into());
        }
    }

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<String> {

        match read_file("./hello.txt".to_string()) {
            Ok(value) => print(value.to_string()),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        let output : String = "👋 Hello ".to_string();
        
        Ok(output)
    }
    ```

## hostWriteFile

**`hostWriteFile`**: write a content to a file.

=== "Go"
    ```golang
    package main

    import (
        "encoding/base64"
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

    var parser = fastjson.Parser{}

    // GetJsonFromBytes
    // Convert a buffer (`[]byte`) into a JSON value
    func GetJsonFromBytes(buffer []byte) (*fastjson.Value, error) {
        return parser.ParseBytes(buffer)
    }

    //export hostWriteFile
    func hostWriteFile(offset uint64) uint64

    func WriteFile(filePath string, contentFile string) error {

        content := base64.StdEncoding.EncodeToString([]byte(contentFile))

        jsonStrArguments := `{"path":"` + filePath + `","content":"` + content + `"}`

        // Copy the string value to the shared memory
        arguments := pdk.AllocateString(jsonStrArguments)

        // Call the host function with Json string argument
        offset := hostWriteFile(arguments.Offset())

        // Get result from the shared memory
        memoryResult := pdk.FindMemory(offset)
        buffResult := make([]byte, memoryResult.Length())
        memoryResult.Load(buffResult)
        JSONData, err := GetJsonFromBytes(buffResult)

        if err != nil {
            return err
        }
        if len(JSONData.GetStringBytes("failure")) == 0 {
            return nil
        } else {
            return errors.New(string(JSONData.GetStringBytes("failure")))
        }
    }

    //export hello
    func hello() uint64 {

        text := `
        <html>
        <h1>"Hello World!!!"</h1>
        </html>
        `

        err := WriteFile("./index.html", text)
        if err != nil {
            Print("😡 " + err.Error())
        }

        return 0
    }

    func main() {}
    ```

=== "Rust"
    ```rust
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

    #[derive(Serialize, Deserialize, Debug)]
    struct StringResult {
        pub success: String,
        pub failure: String,
    }

    #[derive(Serialize, Deserialize, Debug)]
    struct FileArguments {
        pub path: String,
        pub content: String,
    }

    #[derive(Error, Debug)]
    pub enum FileError {
        #[error("Write issue")]
        WriteFailure,
        #[error("Not found")]
        NotFound,
    }

    extern "C" {
        fn hostWriteFile(ptr: u64) -> u64;
    }

    pub fn write_file(file_path: String, content_file: String) -> Result<String, Error> {
        
        let args = FileArguments {
            path: file_path,
            content:  content_file,
        };
        let json_str: String = serde_json::to_string(&args).unwrap();

        // Copy the string value to the shared memory
        let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
        memory_json_str.store(json_str);

        // Call host function with the offset of the arguments
        let offset: u64 = unsafe { hostWriteFile(memory_json_str.offset) };

        // Get result
        let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
        let json_string:String = memory_result.to_string().unwrap();
        let result: StringResult = serde_json::from_str(&json_string).unwrap();
        
        if result.failure.is_empty()  {
            return Ok(result.success);
        } else {
            return Err(FileError::WriteFailure.into());
        }

    }


    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<String> {

        let text = "<html><h1>Hello World!!!</h1></html>".to_string();

        match write_file("./index.html".to_string(), text) {
            Ok(value) => print(value.to_string()),
            Err(error) => print("😡 error: ".to_string() + &error.to_string()),
        }

        let output : String = "👋 Hello ".to_string();
        
        Ok(output)
    }
    ```
