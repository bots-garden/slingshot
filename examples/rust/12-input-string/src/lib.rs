use extism_pdk::*;
use serde::{Serialize, Deserialize};
use thiserror::Error;


extern "C" {
    fn hostPrintln(ptr: u64) -> u64;
}

extern "C" {
    fn hostInput(ptr: u64) -> u64;
}

pub fn println(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrintln(memory_text.offset) };
}

pub fn input(prompt: String) -> String {
    // copy the prompt value to the shared memory
    let mut memory_prompt: Memory = extism_pdk::Memory::new(prompt.len());
    memory_prompt.store(prompt);

    // call the host function
    let offset: u64 = unsafe { hostInput(memory_prompt.offset) };

    // read the value of the result from the shared memory
    let input_value: Memory = extism_pdk::Memory::find(offset).unwrap();

    // return the value
    return input_value.to_string().unwrap()
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


#[plugin_fn]
pub fn hello(_: String) -> FnResult<u64> {

    let name: String = input("🤖 Name? > ".to_string());
    println("🦀 Hello ".to_string() + &name);

    Ok(0)
}
