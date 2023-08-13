#![no_main]

//use std::ptr::null;

use extism_pdk::*;
//use serde::Serialize;
use serde::{Serialize, Deserialize};


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
    fn hostGetMessage(ptr: u64) -> u64;
}

pub fn get_message(key: String) -> String {
    let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
    memory_key.store(key);
    let offset: u64 = unsafe { hostGetMessage(memory_key.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();

    return memory_result.to_string().unwrap()
}

extern "C" {
    fn hostMemorySet(offset: u64) -> u64;
}

#[derive(Serialize, Deserialize, Debug)]
struct MemRecord {
    pub key: String,
    pub value: String,
}

pub fn memory_set(key: String, value: String) {
    let record = MemRecord {
        key: key,
        value: value,
    };
    let json_str = serde_json::to_string(&record).unwrap();

    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);
    let offset: u64 = unsafe { hostMemorySet(memory_json_str.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();

    let msg: String = "🟣 from 🦀 guest: ".to_string() + &memory_result.to_string().unwrap();

    print(msg)
}

extern "C" {
    fn hostMemoryGet(offset: u64) -> u64;
}

#[derive(Serialize, Deserialize, Debug)]
struct StringResult {
    pub success: String,
    pub failure: String,
}

pub fn memory_get(key: String) -> String {

    let mut memory_key: Memory = extism_pdk::Memory::new(key.len());
    memory_key.store(key);
    let offset: u64 = unsafe { hostMemoryGet(memory_key.offset) };

    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
    let json_string:String = memory_result.to_string().unwrap();
    let result: StringResult = serde_json::from_str(&json_string).unwrap();

    // TODO: return Result
    return result.success
    
}



#[plugin_fn]
pub fn handle(input: String) -> FnResult<Json<Output>> {

    print("🟣 this is the wasm handle function".to_string());

    let val1: String = get_message("hello".to_string());
    print(val1);

    let msg: String = "🦀 Hello ".to_string() + &input;

    memory_set("bill".to_string(), "🦊 ballantines".to_string());

    print(memory_get("bill".to_string()));

    let output = Output { success: msg, failure: "no error".to_string() };
    
    Ok(Json(output))
}

