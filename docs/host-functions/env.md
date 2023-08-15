# 🛠️ Host functions

## hostGetEnv

**`hostGetEnv`**: read an environment variable.

=== "Go"
    ```golang linenums="1" hl_lines="15-32"
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

    //export hello
    func hello() uint64 {
        message := GetEnv("MESSAGE")
        Print("🤖 MESSAGE=" + message)

        return 0
    }

    func main() {}

    ```

=== "Rust"
    ```rust linenums="1" hl_lines="13-30"
    use extism_pdk::*;

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

    #[plugin_fn]
    pub fn hello(_: String) -> FnResult<u64> {

        let message : String = get_env("MESSAGE".to_string());

        print("🦀 MESSAGE=".to_string() + &message);
        
        Ok(0)
    }
    ```
