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
    
    // TODO: encode content_file to 64b string
    let args = FileArguments {
        path: file_path,
        content: content_file,
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
