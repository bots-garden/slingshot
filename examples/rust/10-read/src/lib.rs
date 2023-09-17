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
