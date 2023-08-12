#![no_main]

use extism_pdk::*;
use serde::Serialize;

#[derive(Serialize)]
struct Output {
    pub success: String,
    pub failure: String,
}

extern "C" {
    fn hostPrint(ptr: u64) -> u64;
}

pub fn print(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrint(memory_text.offset) };
}

extern "C" {
    fn hostMemoryGet(ptr: u64) -> u64;
}

pub fn memory_get(key: String) -> String {
    let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
    memory_key.store(key);
    let offset: u64 = unsafe { hostMemoryGet(memory_key.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();

    return memory_result.to_string().unwrap()
}


#[plugin_fn]
pub fn handle(input: String) -> FnResult<Json<Output>> {

    print("🟣 this is the wasm handle function".to_string());

    let val1: String = memory_get("hello".to_string());
    print(val1);

    let msg: String = "🦀 Hello ".to_string() + &input;

    let output = Output { success: msg, failure: "no error".to_string() };
    
    Ok(Json(output))
}

