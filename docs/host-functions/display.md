# 🛠️ Host functions

## hostPrint

**`hostPrint`**: print a message (from the wasm module).

=== "Go"
    ```golang linenums="1"
    import (
        "strings"
        "github.com/extism/go-pdk"
    )

    //export hostPrint
    func hostPrint(offset uint64) uint64

    func Print(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrint(memoryText.Offset())
    }

    //export hello
    func hello() uint64 {
        Print("👋 hello world 🌍")
        return 0
    }
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
    pub fn hello(_input: String) -> FnResult<u64> {

        print("👋 hello world 🌍".to_string());
        Ok(0)
    }
    ```

## hostPrintln

**`hostPrintln`**: print a message with a carriage return (from the wasm module).

=== "Go"
    ```golang linenums="1"
    import (
        "strings"
        "github.com/extism/go-pdk"
    )

    //export hostPrintln
    func hostPrintln(offset uint64) uint64

    func Println(text string) {
        memoryText := pdk.AllocateString(text)
        hostPrintln(memoryText.Offset())
    }

    //export hello
    func hello() uint64 {
        Println("👋 hello world 🌍")
        return 0
    }
    ```

=== "Rust"
    ```rust linenums="1"
    use extism_pdk::*;

    extern "C" {
        fn hostPrintln(ptr: u64) -> u64;
    }

    pub fn println(text: String) {
        let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
        memory_text.store(text);
        unsafe { hostPrintln(memory_text.offset) };
    }

    #[plugin_fn]
    pub fn hello(_input: String) -> FnResult<u64> {

        println("👋 hello world 🌍".to_string());
        Ok(0)
    }
    ```

## hostLog

**`hostLog`**: log a message (from the wasm module).

=== "Go"
    ```golang linenums="1"
    import (
        "strings"
        "github.com/extism/go-pdk"
    )

    //export hostLog
    func hostLog(offset uint64) uint64

    func Log(text string) {
        memoryText := pdk.AllocateString(text)
        hostLog(memoryText.Offset())
    }

    //export hello
    func hello() uint64 {
        Log("🙂 have a nice day 🏖️")
        return 0
    }
    ```

=== "Rust"
    ```rust linenums="1"
    use extism_pdk::*;

    extern "C" {
        fn hostLog(ptr: u64) -> u64;
    }

    pub fn log(text: String) {
        let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
        memory_text.store(text);
        unsafe { hostLog(memory_text.offset) };
    }

    #[plugin_fn]
    pub fn hello(_input: String) -> FnResult<u64> {

        log("🙂 have a nice day 🏖️".to_string());
        Ok(0)
    }
    ```
